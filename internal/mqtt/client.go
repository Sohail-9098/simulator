package mqtt

import (
	"log"

	"github.com/Sohail-9098/simulator/internal/protobufs/vehicle"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	client MQTT.Client
}

func NewClient(broker, clientID, username, password string) *Client {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	return &Client{client: MQTT.NewClient(opts)}
}

func (c *Client) Connect() {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect to MQTT broker: %v", token.Error().Error())
	}
}

func (c *Client) Disconnect() {
	c.client.Disconnect(250)
}

func (c *Client) PublishTelemetry(topic string, telemetry *vehicle.Telemetry) error {
	data, err := proto.Marshal(telemetry)
	if err != nil {
		return err
	}
	token := c.client.Publish(topic, 0, false, data)
	token.Wait()
	return token.Error()
}
