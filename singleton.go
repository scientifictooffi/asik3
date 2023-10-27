package main

import (
	"fmt"
	"sync"
)

type ConfigManager struct {
	config map[string]string
}

var instance *ConfigManager
var once sync.Once

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			config: make(map[string]string),
		}
		instance.loadConfiguration()
	})
	return instance
}

func (cm *ConfigManager) loadConfiguration() {
	cm.config["databaseHost"] = "localhost"
	cm.config["databasePort"] = "5432"
	cm.config["appName"] = "MyApp"
}
func (cm *ConfigManager) GetConfig(key string) string {
	return cm.config[key]
}

func main() {
	configManager := GetConfigManager()

	fmt.Printf("Database Host: %s\n", configManager.GetConfig("databaseHost"))
	fmt.Printf("Database Port: %s\n", configManager.GetConfig("databasePort"))
	fmt.Printf("Application Name: %s\n", configManager.GetConfig("appName"))

	configManager2 := GetConfigManager()
	fmt.Println("Are both instances the same?", configManager == configManager2)
}
