package main

import (
	"fmt"
	c "github.com/Lukpier/flinkmonitoring/config"
	"github.com/Lukpier/flinkmonitoring/monitoring"
	"github.com/akamensky/argparse"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func loadConfig(configPath string) c.Config {
	// Set the file name of the configurations file
	viper.SetConfigFile(configPath)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return configuration
}

func main() {

	parser := argparse.NewParser("flinkmonitoring", "Listens for exceptions of all / specified flink jobs and send them to the specified receiver address")
	configPath := parser.String("c", "config", &argparse.Options{Required: true, Help: "path to config file"})
	jobId := parser.String("j", "jobid", &argparse.Options{Required: false, Default: "", Help: "optional job id"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Fatal(parser.Usage(err))
	}

	config := loadConfig(*configPath)

	lookuper := monitoring.NewExceptionLookuper(config)

	done := make(chan bool)
	ticker := time.NewTicker(time.Second * time.Duration(config.Poll))

	fmt.Println("Starting lookup for exception at ", config.Flink.Endpoint)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				var err error
				if *jobId != "" {
					err = lookuper.LookupJobAndSend(*jobId)
				} else {
					err = lookuper.LookupAllAndSend()
				}
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}()
	time.Sleep(time.Second * time.Duration(config.For))
	done <- true
}
