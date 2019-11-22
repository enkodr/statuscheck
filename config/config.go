package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	// Config represents the base configuration
	Config struct {
		Port         string `yaml:"port"`
		Endpoint     string `yaml:"endpoint"`
		OKMessage    string `yaml:"okMessage"`
		JSON         bool   `yaml:"json"`
		ErrorMessage string `yaml:"errorMessage"`
		Check        Check  `yaml:"check"`
	}

	// Check represents the endpoint's configuration to test/check
	Check struct {
		Type string `yaml:"type"`
		// HTTP
		URL             string `yaml:"url"`
		FollowRedirects bool   `yaml:"followRedirects"`
		StatusCode      int    `yaml:"statusCode"`
		// Port
		Port     int    `yaml:"port"`
		Address  string `yaml:"address"`
		Protocol string `yaml:"protocol"`
	}
)

// Load loads the configuration file from filesystem
func Load(path string) (Config, error) {
	c := Config{
		Port:         "8008",
		Endpoint:     "/status",
		JSON:         false,
		OKMessage:    "ok",
		ErrorMessage: "fail",
		Check: Check{
			Type:            "http",
			URL:             "http://localhost",
			FollowRedirects: false,
			StatusCode:      200,
			Port:            80,
			Address:         "127.0.0.1",
			Protocol:        "tcp",
		},
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
