package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sohail-9098/simulator/internal/mqtt"
	"github.com/Sohail-9098/simulator/internal/telemetry"
	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE_PATH = "../../config/config.yaml"
)

type Config struct {
	MQTT struct {
		Broker   string `yaml:"broker"`
		ClientID string `yaml:"client_id"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mqtt"`
	Vehicles []string `yaml:"vehicles"`
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	var config *Config
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func main() {
	// Load config
	config, err := loadConfig(CONFIG_FILE_PATH)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Connect to MQTT
	mqttClient := mqtt.NewClient(config.MQTT.Broker, config.MQTT.ClientID, config.MQTT.Username, config.MQTT.Password)
	mqttClient.Connect()
	defer mqttClient.Disconnect()

	for {
		for _, vehicleID := range config.Vehicles {
			telemetryData := telemetry.GenerateTelemetry(vehicleID)
			topic := fmt.Sprintf("vehicles/%s/telemetry", vehicleID)
			log.Printf("publishing telemetry data for %s", vehicleID)
			err := mqttClient.PublishTelemetry(topic, telemetryData)
			if err != nil {
				log.Printf("failed to publish telemetry data for %s: %v", vehicleID, err)
			}
		}
		time.Sleep(time.Second)
	}
}
