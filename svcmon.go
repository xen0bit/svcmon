package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sys/windows/svc"
)

const svcName = "MSMQ"

func getStatus(c *gin.Context) {
	status, err := statusService(svcName)
	if err != nil {
		log.Println(err)
		c.String(http.StatusPreconditionFailed, "fuckingwinapi")
	} else {
		if status.State == svc.Running {
			c.String(http.StatusOK, "Running")
		} else {
			c.String(http.StatusOK, "Stopped")
		}
	}
}

func getRestart(c *gin.Context) {
	err := controlService(svcName, svc.Stop, svc.Stopped)
	if err != nil {
		log.Println(err)
		c.String(http.StatusPreconditionFailed, "fuckingwinapi")
	}
	err = startService(svcName)
	if err != nil {
		log.Println(err)
		c.String(http.StatusPreconditionFailed, "fuckingwinapi")
	}
	c.String(http.StatusOK, "Success")
}

func main() {

	router := gin.Default()
	router.GET("/status", getStatus)
	router.GET("/restart", getRestart)

	router.Run("0.0.0.0:8080")
}
