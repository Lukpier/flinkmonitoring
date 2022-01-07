FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY config/ config
COPY monitoring/ monitoring
COPY *.go /app/

RUN go build -o /flinkmonitoring

ENTRYPOINT [ "/flinkmonitoring" ]