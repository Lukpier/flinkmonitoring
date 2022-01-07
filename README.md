# GOFLINK-EXCEPTION-ALERTING

Flink exceptions alerting made simple (with go!).

# Build

* `go install`
* `go build -o flinkmonitoring`

## Usage

`flinkmonitoring -config $CONFIG_PATH -job $JOBID <optional>`

* Configure sender and receivers mail info in related config section.
* Specify flink endpoint
* Customize poll and run time as you prefer

## Config

```
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

# Run with Docker
`docker build --tag flinkmonitoring .`
`docker run --network=host -v  $PWD/config.yml:/app/config.yml flinkmonitoring -config config.yml`

