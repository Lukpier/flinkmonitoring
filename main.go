package main

import (
	"flag"
	"fmt"
	c "github.com/Lukpier/flinkmonitoring/config"
	"github.com/Lukpier/flinkmonitoring/monitoring"
	"github.com/spf13/viper"
	"log"
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

	configPath := *flag.String("config", "./config.yml", "path to the config file")
	jobId := *flag.String("job", "", "specific job to follow")

	config := loadConfig(configPath)

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
				if jobId != "" {
					err = lookuper.LookupJobAndSend(jobId)
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
