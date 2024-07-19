package telemetry

import (
	"math/rand/v2"

	"github.com/Sohail-9098/simulator/internal/protobufs/vehicle"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateTelemetry(vehicleID string) *vehicle.Telemetry {
	return &vehicle.Telemetry{
		VehicleId: vehicleID,
		Timestamp: timestamppb.Now(),
		Latitude:  rand.Float64()*180 - 90,
		Longitude: rand.Float64()*360 - 180,
		Speed:     rand.Float64() * 200,
		FuelLevel: rand.Float64() * 100,
	}
}
