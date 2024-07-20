package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Sohail-9098/simulator/internal/mqtt"
	"github.com/Sohail-9098/simulator/internal/telemetry"
	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE_PATH = "../../config/config.yaml"
)

var (
	isPublishing bool
	stopCh       chan struct{}
	mu           sync.Mutex
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
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go HandleUserInput()
	<-sigs
	fmt.Println("exiting")
}

func HandleUserInput() {
	for {
		var input int
		fmt.Println("Press 1 and Enter to start publish")
		fmt.Scanln(&input)

		mu.Lock()
		if input == 1 {
			if isPublishing {
				fmt.Println("Already Publishing")
			} else {
				stopCh = make(chan struct{})
				isPublishing = true
				go func() {
					StartPublishing()
					mu.Lock()
					isPublishing = false
					mu.Unlock()
				}()
			}
		} else {
			if isPublishing {
				close(stopCh)
				isPublishing = false
				fmt.Println("Publish Stop")
			} else {
				fmt.Println("Currently not publishing")
			}
		}
		mu.Unlock()
	}
}

func StartPublishing() {
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
