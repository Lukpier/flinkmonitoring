package monitoring

import (
	"encoding/json"
	"fmt"
	c "github.com/Lukpier/goflink-exception-alerting/config"
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

	for _, job := range jobs.Jobs {
		if err := el.LookupJobAndSend(*job.Id); err != nil {
			return err
		}
	}

	return nil
}

func (el *ExceptionLookuper) LookupJobAndSend(jobId string) error {
	jobs, err := el.flinkClient.GetJobs()

	if err != nil {
		log.Fatal("cannot lookup for jobs")
	}

	for _, job := range jobs.Jobs {

		exceptions, err := el.flinkClient.LookupExceptions(*job.Id)

		if err != nil {
			fmt.Printf("cannot lookup for exceptions for job %s", *job.Id)
			return err
		}

		if len(exceptions.AllExceptions) > 0 {
			fmt.Printf("Found exceptions for Job: %s", *job.Id)
			formatted := formatExceptions(&exceptions)

			if err := el.mailClient.SendMail(*job.Id, formatted); err != nil {
				fmt.Println("Error sendimg mail", err)
				return err
			}
		} else {
			fmt.Println("No exceptions for Job: ", *job.Id)
		}

	}

	return nil

}

func formatExceptions(exceptions *FlinkExceptions) string {
	s, err := json.MarshalIndent(exceptions, "", "  ")
	if err != nil {
		log.Println("cannot marshal: ", exceptions)
	}
	return string(s)
}
