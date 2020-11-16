package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen              string
	TargetServer        string
	AllowedFingerprints map[string]bool
	AllowedCommands     map[byte]bool
}

type rawConfig struct {
	Listen              string   `yaml:"listen"`
	TargetServer        string   `yaml:"target_server"`
	AllowedFingerprints []string `yaml:"allowed_fingerprints"`
	AllowedCommands     []string `yaml:"allowed_commands"`
}

var (
	config     *Config
	configLock = new(sync.RWMutex)
)

func loadConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	rawConfig := rawConfig{}
	if err = yaml.Unmarshal(file, &rawConfig); err != nil {
		log.Println("parse config: ", err)
		return fmt.Errorf("failed to parse config: %w", err)
	}

	tempConfig := new(Config)
	tempConfig.AllowedFingerprints = make(map[string]bool)
	tempConfig.AllowedCommands = make(map[byte]bool)

	tempConfig.Listen = rawConfig.Listen
	tempConfig.TargetServer = rawConfig.TargetServer
	for _, fingerprint := range rawConfig.AllowedFingerprints {
		tempConfig.AllowedFingerprints[fingerprint] = true
	}
	for _, command := range rawConfig.AllowedCommands {
		if len(command) != 1 {
			return fmt.Errorf("command must be one char: %s", command)
		}
		tempConfig.AllowedCommands[command[0]] = true
	}

	configLock.Lock()
	config = tempConfig
	configLock.Unlock()

	return nil
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}
