package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Sohail-9098/simulator/internal/config"
	"github.com/Sohail-9098/simulator/internal/mqtt"
	"github.com/Sohail-9098/simulator/internal/telemetry"
)

const (
	CONFIG_FILE_PATH = "../../config/config.yaml"
)

func startPublish() {
	// Load config
	config, err := config.LoadConfig(CONFIG_FILE_PATH)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Connect to MQTT
	mqttClient := mqtt.NewClient(config.MQTT.Broker, config.MQTT.ClientID, config.MQTT.Username, config.MQTT.Password)
	mqttClient.Connect()
	defer mqttClient.Disconnect()
	for {
		select {
		case <-stopCh:
			log.Println("Stopping Publish")
			return
		default:
			for _, vehicleID := range config.Vehicles {
				telemetryData := telemetry.GenerateTelemetry(vehicleID)
				topic := fmt.Sprintf("vehicles/%s/telemetry", vehicleID)
				fmt.Println("publish : ", vehicleID)
				err := mqttClient.PublishTelemetry(topic, telemetryData)
				if err != nil {
					log.Printf("failed to publish telemetry data for %s: %v", vehicleID, err)
				}
			}
			time.Sleep(time.Second)
		}
	}
}
