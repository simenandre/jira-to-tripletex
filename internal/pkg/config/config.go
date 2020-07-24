package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Host          string
	JiraToken     string
	JiraEmail     string
	ConsumerToken string
	EmployeeToken string
	ActivityCode  string
}

func Init() (*viper.Viper, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := fmt.Sprintf("%s/.JIRATripletex", userDir)
	_ = os.Mkdir(configPath, os.ModePerm)

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file does not exists. Creating!")
		} else {
			return nil, err
		}
	}
	return v, nil
}

func SetConfig(config Config) error {
	v, err := Init()
	if err != nil {
		return err
	}

	userDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := fmt.Sprintf("%s/.JIRATripletex/config.yaml", userDir)

	v.Set("host", config.Host)
	v.Set("jiraToken", config.JiraToken)
	v.Set("jiraEmail", config.JiraEmail)
	v.Set("consumerToken", config.ConsumerToken)
	v.Set("employeeToken", config.EmployeeToken)
	v.Set("activityCode", config.ActivityCode)
	err = v.WriteConfigAs(configPath)
	if err != nil {
		return err
	}
	return nil
}

func values() (*Config, error) {
	v, err := Init()
	if err != nil {
		return nil, err
	}
	return &Config{
		Host:          v.GetString("host"),
		JiraToken:     v.GetString("jiraToken"),
		JiraEmail:     v.GetString("jiraEmail"),
		ConsumerToken: v.GetString("consumerToken"),
		EmployeeToken: v.GetString("employeeToken"),
		ActivityCode:  v.GetString("activityCode"),
	}, nil
}

func GetConfig() (*Config, error) {
	return values()
}

func Print() {
	cfg, err := values()
	if err != nil {
		log.Fatal("Something happened when parsing config")
	}
	fmt.Println("Host:", cfg.Host)
}
