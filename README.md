# Vehicle Data Simulator

The Vehicle Data Simulator is a powerful tool designed to simulate the telemetry data of multiple vehicles and publish it to an MQTT broker. It's ideal for testing and developing applications that consume MQTT messages.

## Features

- Simulates telemetry data for multiple vehicles.
- Publishes data to an MQTT broker.
- Configurable via AWS secrets.
- Simple command-line interface for starting and stopping data publishing.
- Dockerfile included for containerized deployment.

## Prerequisites

- **Go 1.18+**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
- **EMQX MQTT Broker**: Ensure you have a running instance of the EMQX broker. You can download and start it from [EMQX](https://www.emqx.io/).
- **Mosquitto Client**: Required for subscribing to MQTT topics. Install it from [Mosquitto](https://mosquitto.org/download/).

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

   Ensure your EMQX broker is running and accessible. Default settings are usually sufficient for local development.

## Configuration

Configure your application using AWS secrets instead of a local configuration file. Ensure your AWS secrets are properly set up to store the necessary MQTT broker information.

## Running the Simulator

1. **Start the Application**

   Run the application with the following command:

   ```sh
   go run cmd/simulator/main.go
   ```

2. **Control Publishing**

   The application will prompt you to control data publishing:

   ```plaintext
   Press 1 and Enter to start generating and sending telemetry data
   Press 2 to stop generating and sending
   Press Ctrl+C to exit
   ```

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
- **Configuration Issues**: Verify your AWS secrets configuration.
- **Logging**: Enable detailed logging in the `main.go` file to trace connection and publishing issues.

## Containerized Deployment

A Dockerfile is included to facilitate containerized deployment. You can build and run the Docker container as follows:

1. **Edit the Dockerfile to Configure AWS Environment Variables**

   Update the Dockerfile with your AWS credentials and MQTT broker details:

   ```Dockerfile
   ENV MQTT_BROKER_URL="tcp://host.docker.internal:1883" \
       MQTT_CREDENTIALS_SECRET="vehicle/vehicle-simulator/mqtt/credentials" \
       AWS_ACCESS_KEY_ID="test" \
       AWS_SECRET_ACCESS_KEY="test" \
       AWS_DEFAULT_REGION="us-east-1" \
       AWS_ENDPOINT_URL="http://host.docker.internal:4566"
   ```

2. **Build the Docker Image**

   ```sh
   docker build -t vehicle-simulator .
   ```

3. **Run the Docker Container**

   ```sh
   docker run --it vehicle-simulator
   ```

---

For more information or if you encounter any issues, feel free to reach out to me at sohailkhan9098@gmail.com.