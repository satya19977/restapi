package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/satya19977/restapi/db"
	"github.com/satya19977/restapi/routes"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
  
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
    log.Println("We are logging in Golang! using logrus")
  
	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
  }

func Logger() gin.HandlerFunc{
	return func(c *gin.Context){
		// Log Incoming Requests
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"ip":     c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Info("Incoming request")
	}
  }

func main(){
	db.InitDB()
	server := gin.Default()
	// Logging middleware
	server.Use(Logger())
	routes.RegisterRoutes(server) // points to the original server
	server.Run(":8080") //localhost:8080
	
}
