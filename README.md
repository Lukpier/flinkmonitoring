# Flink Monitoring

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Lukpier/flinkmonitoring?style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Lukpier/flinkmonitoring/Go?style=for-the-badge)

Flink exceptions alerting made simple (with go!).

# Build

* `go install`
* `go build -o flinkmonitoring`

## Usage

```
usage: flinkmonitoring [-h|--help] -c|--config "<value>" [-j|--jobid "<value>"]

                       Listens for exceptions of all / specified flink jobs and
                       send them to the specified receiver address

Arguments:

  -h  --help    Print help information
  -c  --config  path to config file
  -j  --jobid   optional job id. Default:
```

* Configure sender and receivers mail info in related config section.
* Specify flink endpoint
* Customize poll and run time as you prefer

## Config

```yaml
poll: 5 # seconds
for: 15 # seconds

mail:
  sender: "sendermailaddress"
  password: "senderpassword"
  receivers: 
    - "receiver1"
    - "receiver2
  smtphost: "smtp.gmail.com"
  smtpport: "587"

flink:
  endpoint: http://localhost:8081
```

## Run with Docker
*`docker build --tag flinkmonitoring .`
*`docker run --network=host -v  $PWD/config.yml:/app/config.yml flinkmonitoring -config config.yml`

