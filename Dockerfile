FROM golang:1.22.5-alpine
ENV MQTT_BROKER_URL="tcp://localhost:1883"
ENV MQTT_CREDENTIALS_SECRET="vehicle/vehicle-simulator/mqtt/credentials"
WORKDIR /app
COPY . ./
RUN go mod tidy
WORKDIR /app/cmd/simulator
RUN go build -o simulator .
ENTRYPOINT [ "./simulator" ]