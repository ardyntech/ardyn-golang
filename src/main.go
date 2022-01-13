package main

import (
	"flag"
	"log"

	"ardyngolang/src/env"

	"ardyngolang/src/certsec"

	eureka "github.com/xuanbo/eureka-client"
	//"github.com/hudl/fargo"
)

//--------------------------------------------------------------------

func registerWithSoothsayer() {

	// Connect to Eureka
	// https://github.com/xuanbo/eureka-client/blob/master/examples/main.go

	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           env.Config.EurekaDefaultZone,
		App:                   env.Config.ApplicationName,
		Port:                  env.Config.Port,
		RenewalIntervalInSecs: 20,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})

	// start client, register, heartbeat, refresh
	client.Start()

}

//--------------------------------------------------------------------

func main() {

	log.Println("=================================")
	log.Println("Ardyn-Golang Example Service v0.1")
	log.Println("=================================")
	log.Println()

	// Get the command line flags
	configPtr := flag.String("config", "", "Config file location")

	// Parse the command line parameters
	flag.Parse()

	// Parse the properties file
	env.ParseConfig(*configPtr)

	// Register this service with Soothsayer
	registerWithSoothsayer()

	// Initialize the security
	if !certsec.InitSecurity() {

		// Something went wrong
		log.Fatalln("*** Error initializing certificate ***")
	}

	// Run the HTTP server
	startHttp()

}

//--------------------------------------------------------------------
