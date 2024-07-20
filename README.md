# Vehicle Data Simulator

The Vehicle Data Simulator is a tool designed to simulate the telemetry data of multiple vehicles and publish it to an MQTT broker. This can be useful for testing and developing applications that consume MQTT messages.

## Features

- Simulates telemetry data for multiple vehicles.
- Publishes data to an MQTT broker.
- Configurable via a YAML configuration file.
- Handles starting and stopping of data publishing via a simple command-line interface.

## Prerequisites

- **Go 1.18+**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
- **EMQX MQTT Broker**: A running instance of the EMQX broker. You can download and start it from [EMQX](https://www.emqx.io/).
- **Mosquitto Client**: For subscribing to MQTT topics, you need the `mosquitto` client. Install it from [Mosquitto](https://mosquitto.org/download/).

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/Sohail-9098/vehicle-simulator.git
   cd vehicle-simulator
   ```

2. **Install Dependencies**

   Navigate to the project directory and run:

   ```sh
   go mod tidy
   ```

3. **Configure EMQX**

   Ensure your EMQX broker is running and accessible. Default settings are usually fine for local development.

## Configuration

Create a `config.yaml` file in the `config` directory with the following structure:

```yaml
mqtt:
  broker: "tcp://localhost:1883"
  client_id: "vehicle_simulator"
  username: ""
  password: ""
vehicles:
  - vehicle1
  - vehicle2
  - vehicle3
```

## Running the Simulator

1. **Start the Application**

   Run the application with the following command:

   ```sh
   go run cmd/simulator/main.go
   ```

2. **Control Publishing**

   The application will prompt you to start or stop publishing:

   ```plaintext
   Press 1 and Enter to start publish
   ```

   Type `1` to start publishing data. Any other key to stop publishing. Press `Ctrl+C` to stop the application.

## Subscribing to Data

Use `mosquitto_sub` to subscribe to the telemetry topics:

```sh
mosquitto_sub -h localhost -t 'vehicles/#'
```

This will subscribe to all topics under `vehicles/` and display incoming messages.

## Example Command to Subscribe

```sh
mosquitto_sub -h localhost -t 'vehicles/#'
```

## Troubleshooting

- **No Messages Received**: Ensure the EMQX broker is running and the topic matches exactly.
- **Configuration Issues**: Verify `config.yaml` and ensure all fields are correctly set.
- **Logging**: Enable detailed logging in the `main.go` file to trace connection and publishing issues.

---

If you need more information or run into any issues, kindly reach out to me sohailkhan9098@gmail.com