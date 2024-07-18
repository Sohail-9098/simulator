package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Sohail-9098/simulator/internal/proto/vehicle"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func generateTelemetry(vehicleId string) *vehicle.Telemetry {
	return &vehicle.Telemetry{
		VehicleId: vehicleId,
		Timestamp: timestamppb.Now(),
		Latitude:  rand.Float64()*180 - 90,
		Longitude: rand.Float64()*360 - 180,
		Speed:     rand.Float64() * 200,
		FuelLevel: rand.Float64() * 100,
	}
}

func main() {
	// Set options for MQTT Client
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("vehicle_simulator")
	opts.SetUsername("")
	opts.SetPassword("")

	// Create a new MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect to emqx broker: %v", token.Error().Error())
	}

	// Generate telemetry data
	vehicle_ids := []string{"vehicle1", "vehicle2", "vehicle3"}
	for {
		for _, vehicleId := range vehicle_ids {
			telemetry := generateTelemetry(vehicleId)
			// serialize the data
			data, err := proto.Marshal(telemetry)
			if err != nil {
				fmt.Printf("failed to marshal telemetry: %v\n", err.Error())
				continue
			}
			topic := fmt.Sprintf("vehicles/%s/telemetry", vehicleId)
			token := client.Publish(topic, 0, false, data)
			token.Wait()
			if token.Error() != nil {
				fmt.Printf("failed to publish telemetry: %v\n", token.Error().Error())
			}
		}
		time.Sleep(time.Second)
	}
}
