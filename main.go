package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
    LogLevel string `mapstructure:"loglevel"`
    DB       struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Database string `mapstructure:"database"`
    } `mapstructure:"db"`
    Port int `mapstructure:"port"`
}

func main() {
	// Load configuration from config.yaml
    var config Config
    if err := loadConfig(&config, "config.yaml"); err != nil {
        log.Fatalf("Error loading config: %s", err)
    }
	
	// Crée une instance du routeur Gin
	router := gin.Default()

	// Use configuration values
    log.Printf("Log Level: %s", config.LogLevel)
    log.Printf("DB Host: %s", config.DB.Host)
    log.Printf("DB Port: %d", config.DB.Port)
    log.Printf("Server Port: %d", config.Port)

	// Définit une route simple
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	// Définit une route pour /healthz
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	// Définit une route pour /hello/:name
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := fmt.Sprintf("hello %s", name)
		c.String(200, message)
	})

	// Lance le serveur sur le port 8080
	router.Run(fmt.Sprintf(":%d", config.Port))
}

func loadConfig(cfg interface{}, path string) error {
    viper.SetConfigFile(path)
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    if err := viper.Unmarshal(cfg); err != nil {
        return err
    }

    return nil
}