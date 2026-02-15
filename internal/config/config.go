package config

import (
	"errors"
	"fmt"
)

var ErrInvalidAppEnv = errors.New("app env is invalid or not supported")

type AppEnv string

const (
	Development AppEnv = "development"
	Release     AppEnv = "release"
)

func parseEnvVariable(env string) (AppEnv, error) {
	switch env {
	case string(Development):
		return Development, nil
	case string(Release):
		return Release, nil
	default:
		return "", ErrInvalidAppEnv
	}
}

type HTTPServerConfig struct {
	Port uint16
}

type GreenAPIConfig struct {
	URL string
}

type AppConfig struct {
	HTTPServer *HTTPServerConfig
	GreenAPI   *GreenAPIConfig
	Env        AppEnv
}

func New() (*AppConfig, error) {

	const op = "config.New"

	httpServerPort, err := getEnvAs[uint16]("HTTP_SERVER_PORT")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	envString, err := getEnv("APP_ENV")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	env, err := parseEnvVariable(envString)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	greenAPIURL, err := getEnv("GREEN_API_URL")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &AppConfig{
		HTTPServer: &HTTPServerConfig{
			Port: httpServerPort,
		},
		GreenAPI: &GreenAPIConfig{
			URL: greenAPIURL,
		},
		Env: env,
	}, nil

}
