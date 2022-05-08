package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	//"github.com/go-logr/logr"
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
	r.GET("/discovery", spectrum.GetDiscovery)       // Discover all animations with options
	r.GET("/animation", spectrum.GetAnimation)       // Get list of all running animation
	r.POST("/animation", spectrum.PostAnimation)     // Start a new anination
	r.DELETE("/animation", spectrum.DeleteAnimation) // Stop a new animation
	r.GET("/settings", spectrum.GetSettings)         // Get all settings
	r.POST("/settings", spectrum.PostSettings)       // Set all settings
	r.GET("/brightness", spectrum.GetBrightness)     // Get brightness
	r.POST("/brightness", spectrum.PostBrightness)   // Set brightness
	r.GET("/wifi", spectrum.GetWifi)                 // Set wifi settings
	r.POST("/wifi", spectrum.PostWifi)               // Set wifi settings

	// Get program args
	args := os.Args[1:]
	log.V(0).Info("Args", "args", args)

	// Spectrum
	spectrum.Run(log, args)
	// Gin
	r.Run(":8080")
	// Then blocking (waiting for quit signal):
	<-quit

	// And in another goroutine if you want to quit:
	// close(quit)
}

func printWelcomeMessage() {
	fmt.Println(`
--------------------------------------------------------
 _____                     _                           
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
