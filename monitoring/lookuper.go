package monitoring

import (
	"encoding/json"
	"fmt"
	c "github.com/Lukpier/flinkmonitoring/config"
	"log"
)

type ExceptionLookuper struct {
	poolPeriod  int
	mailClient  *MailClient
	flinkClient *FlinkClient
}

func NewExceptionLookuper(config c.Config) *ExceptionLookuper {
	return &ExceptionLookuper{
		poolPeriod:  config.Poll,
		mailClient:  NewMailClient(config.Mail),
		flinkClient: NewFlinkClient(config.Flink),
	}
}

func (el *ExceptionLookuper) LookupAllAndSend() error {
	jobs, err := el.flinkClient.GetJobs()

	if err != nil {
		log.Fatal("cannot lookup for jobs")
	}

	if len(jobs.Jobs) == 0 {
		log.Fatal("No Flink job running!")
	}

	for _, job := range jobs.Jobs {
		if err := el.lookupJobAndSend(*job.Id, *job.Status); err != nil {
			return err
		}
	}

	return nil
}

func (el *ExceptionLookuper) LookupJobAndSend(id string) error {
	jobInfo, err := el.flinkClient.GetJob(id)
	if err != nil {
		log.Fatalf("cannot lookup for job %s", id)
	}
	return el.lookupJobAndSend(*jobInfo.Id, *jobInfo.Properties.State)

}

func (el *ExceptionLookuper) lookupJobAndSend(id string, status JobStatus) error {

	exceptions, err := el.flinkClient.LookupExceptions(id)

	if err != nil {
		fmt.Printf("cannot lookup for exceptions for job %s", id)
		return err
	}

	if len(exceptions.AllExceptions) > 0 {
		fmt.Printf("Found exceptions for Job: %s", id)
		formatted := formatExceptions(id, status, &exceptions)

		if err := el.mailClient.SendMail(id, formatted); err != nil {
			fmt.Println("Error sendimg mail", err)
			return err
		}
	} else {
		fmt.Println("No exceptions for Job: ", id)
	}

	return nil

}

func formatExceptions(job string, status JobStatus, exceptions *FlinkExceptions) string {
	s, err := json.MarshalIndent(exceptions, "", "  ")
	if err != nil {
		log.Println("cannot marshal: ", exceptions)
	}
	return fmt.Sprintf(`Dear maintainer,

			The Flink job %v has encountered exceptions. Its current status is %v.
			Below the exception list:

			%v
			`, job, status, s)
}
