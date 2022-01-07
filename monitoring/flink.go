package monitoring

import (
	"encoding/json"
	"fmt"
	c "github.com/Lukpier/flinkmonitoring/config"
	"io/ioutil"
	"log"
	"net/http"
)

type FlinkClient struct {
	endpoint string
	client   *http.Client
}

func NewFlinkClient(config c.FlinkConfig) *FlinkClient {
	return &FlinkClient{
		endpoint: config.Endpoint,
		client:   &http.Client{},
	}
}

var _ IFlinkClient = (*FlinkClient)(nil)

func (fc *FlinkClient) GetJobs() (Jobs, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/jobs", fc.endpoint), nil)
	res, err := fc.client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error requesting Flink jobs")
		return Jobs{}, err
	}

	var data Jobs
	json.Unmarshal(body, &data)
	return data, nil
}

func (fc *FlinkClient) LookupExceptions(jobId string) (FlinkExceptions, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/jobs/%s/exceptions", fc.endpoint, jobId), nil)
	res, err := fc.client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error requesting Flink exceptions for job: %s", jobId)
		return FlinkExceptions{}, err
	}

	var data FlinkExceptions
	json.Unmarshal(body, &data)
	return data, nil
}
