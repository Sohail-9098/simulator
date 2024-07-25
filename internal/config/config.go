package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Config struct {
	MQTT struct {
		Broker   string `json:"broker"`
		ClientID string `json:"client_id"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
}

func New() (*Config, error) {
	config := &Config{}

	secret := os.Getenv("MQTT_CREDENTIALS_SECRET")
	if secret == "" {
		errStr := "environment variable MQTT_CREDENTIALS_SECRET not set"
		return nil, errors.New(errStr)
	}

	credentials, err := getSecret(secret)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(credentials), &config.MQTT)
	if err != nil {
		errStr := fmt.Sprintf("failed to decode config json: %v", err)
		return nil, errors.New(errStr)
	}

	broker := os.Getenv("MQTT_BROKER_URL")
	if broker == "" {
		errStr := "environment variable MQTT_BROKER_URL not set"
		return nil, errors.New(errStr)
	}
	config.MQTT.Broker = broker

	return config, nil
}

func getSecret(secretName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		errStr := fmt.Sprintf("failed to load cloud config: %v", err)
		return "", errors.New(errStr)
	}
	svc := secretsmanager.NewFromConfig(cfg)
	resp, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	if err != nil {
		errStr := fmt.Sprintf("failed to get secret: %s, %v", secretName, err)
		return "", errors.New(errStr)
	}
	return *resp.SecretString, nil
}
