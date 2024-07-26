# BASE IMAGE
FROM golang:1.22.5-alpine

# SET ENV VARIABLES
ENV MQTT_BROKER_URL="tcp://host.docker.internal:1883" \
    MQTT_CREDENTIALS_SECRET="vehicle/vehicle-simulator/mqtt/credentials" \
    AWS_ACCESS_KEY_ID="test" \
    AWS_SECRET_ACCESS_KEY="test" \
    AWS_DEFAULT_REGION="us-east-1" \
    AWS_ENDPOINT_URL="http://host.docker.internal:4566"

# COPY
WORKDIR /app
COPY . ./

# DOWNLAOD DEPENDENCIES
RUN go mod tidy

# BUILD IMAGE
WORKDIR /app/cmd/simulator
RUN go build -o simulator .

# START SERVICE
ENTRYPOINT [ "./simulator" ]