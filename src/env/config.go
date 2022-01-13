package env

//-------------------------------------------------------------

import (
	"log"
	"strconv"
	"strings"

	"github.com/magiconair/properties"

	"ardyngolang/src/helpers"
)

//-------------------------------------------------------------

type EnvConfig struct {
	Port int

	ApplicationName string

	EurekaDefaultZone string

	EurekaInstanceID string

	EurekaPreferIP bool

	PublicKeyFile string

	KafkaGroupId string

	KafkaServer string

	KafkaTopicLogbook string
}

//-------------------------------------------------------------

// Global to store our configuration
var Config EnvConfig

//-------------------------------------------------------------

func ParseConfig(configFile string) {

	log.Println("Config location: ", configFile)

	p := properties.MustLoadURL(configFile)

	// General config stuff
	Config.Port = p.GetInt("server.port", 8080)
	Config.ApplicationName = p.GetString("application.name", "server")
	Config.EurekaInstanceID = p.GetString("eureka.instance.instance-id", "server:1")
	Config.EurekaDefaultZone = p.GetString("eureka.client.serviceUrl.defaultZone", "http://localhost:8761/eureka/")
	Config.EurekaPreferIP = p.GetBool("eureka.instance.preferIpAddress", true)

	Config.PublicKeyFile = p.GetString("ardyn.security.publicKey", "publickey.der")

	// Kafka config
	Config.KafkaGroupId = p.GetString("kafka.groupId", "mygroup")
	Config.KafkaServer = p.GetString("kafka.server", "http://localhost:9092")
	Config.KafkaTopicLogbook = p.GetString("kafka.topics.logbook", "myapp.logbook")

	// Replace the string "{{random}}" in the properties file with a random number
	if strings.Contains(strings.TrimSpace(Config.EurekaInstanceID), "{{random}}") {

		n := helpers.GenerateRandomNumber(1, 999999)

		// Generate a random ID for the instance
		Config.EurekaInstanceID = strings.Replace(Config.EurekaInstanceID, "{{random}}", strconv.Itoa(n), -1)

	}

	// If the port is configured as 0, generate a random port number
	if Config.Port == 0 {

		p := helpers.GenerateRandomNumber(9000, 65565)

		Config.Port = p
	}

	log.Println("-------------------------------------------------")
	log.Println("Server Port: ", Config.Port)
	log.Println("Application Name: ", Config.ApplicationName)
	log.Println("Eureka Default Zone: ", Config.EurekaDefaultZone)
	log.Println("Eureka Instance ID: ", Config.EurekaInstanceID)
	log.Println("Eureka Prefer IP: ", Config.EurekaPreferIP)
	log.Println("Public Key File: ", Config.PublicKeyFile)
	log.Println("-------------------------------------------------")
	log.Println()

}

//-------------------------------------------------------------
