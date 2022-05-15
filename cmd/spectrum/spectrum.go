package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	"github.com/maxthom/spectrum-go/internal/spectrum"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

var (
	log  logr.Logger
	quit = make(chan struct{})
)

const ()

func main() {
	printWelcomeMessage()

	// Setup logger
	zapLogger, err := zap.NewDevelopment()
	utils.CheckError(err)
	defer zapLogger.Sync()
	log = zapr.NewLogger(zapLogger).WithName("spectrum")

	// Setup api
	r := gin.Default()
	isReady := false
	health := r.Group("")
	{
		health.GET("/alive", func(c *gin.Context) {
			c.String(http.StatusOK, "alive")
		})
		health.GET("/ready", func(c *gin.Context) {
			c.String(http.StatusOK, "%v", isReady)
		})
	}
	v1 := r.Group("/api")
	{
		v1.GET("/discovery", spectrum.GetDiscovery)       // Discover all animations with options
		v1.GET("/animation", spectrum.GetAnimation)       // Get list of all running animation
		v1.POST("/animation", spectrum.PostAnimation)     // Start a new anination
		v1.DELETE("/animation", spectrum.DeleteAnimation) // Stop an animation
		v1.GET("/settings", spectrum.GetSettings)         // Get all settings
		v1.POST("/settings", spectrum.PostSettings)       // Set all settings
		v1.GET("/brightness", spectrum.GetBrightness)     // Get brightness
		v1.POST("/brightness", spectrum.PostBrightness)   // Set brightness
		v1.GET("/wifi", spectrum.GetWifi)                 // Set wifi settings
		v1.POST("/wifi", spectrum.PostWifi)               // Set wifi settings
	}

	// Get program args
	args := os.Args[1:]
	log.V(0).Info("Args", "args", args)

	// Spectrum
	spectrum.Run(log)

	// Gin
	isReady = true
	r.Run(":8080")
	// Then blocking (waiting for quit signal):
	//<-quit

	// And in another goroutine if you want to quit:
	// close(quit)
}

func printWelcomeMessage() {
	fmt.Println(`
--------------------------------------------------------
 ______                    __                          
 /  ___|                   | |                          
 \ '--.  _ __    ___   ___ | |_  _ __  _   _  _ __ ___  
  '--. \| '_ \  / _ \ / __|| __|| '__|| | | || '_ '  _ \ 
 /\__/ /| |_) ||  __/| (__ | |_ | |   | |_| || | | | | |
 \____/ | .__/  \___| \___| \__||_|    \__,_||_| |_| |_|
		| |                                             
		|_|                                             
-------------------------------------------------------- 
		`)
}
