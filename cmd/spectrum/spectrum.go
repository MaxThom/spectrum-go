package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.uber.org/zap"

	"github.com/maxthom/spectrum-go/internal/spectrum"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

var (
	log logr.Logger
	//quit = make(chan struct{})
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
	r := mux.NewRouter().StrictSlash(true)

	isReady := false

	r.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "alive")
	})
	r.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if isReady {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "ready")
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "not ready")
		}
	})
	r.HandleFunc("/api/discovery", spectrum.GetDiscovery).Methods("GET")       // Discover all animations with options
	r.HandleFunc("/api/animation", spectrum.GetAnimation).Methods("GET")       // Get list of all running animation
	r.HandleFunc("/api/animation", spectrum.PostAnimation).Methods("POST")     // Start a new anination
	r.HandleFunc("/api/animation", spectrum.DeleteAnimation).Methods("DELETE") // Stop an animation
	r.HandleFunc("/api/settings", spectrum.GetSettings).Methods("GET")         // Get all settings
	r.HandleFunc("/api/settings", spectrum.PostSettings).Methods("POST")       // Set all settings
	r.HandleFunc("/api/brightness", spectrum.GetBrightness).Methods("GET")     // Get brightness
	r.HandleFunc("/api/brightness", spectrum.PostBrightness).Methods("POST")   // Set brightness
	r.HandleFunc("/api/wifi", spectrum.GetWifi).Methods("GET")                 // Set wifi settings
	r.HandleFunc("/api/wifi", spectrum.PostWifi).Methods("POST")               // Set wifi settings

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// Get program args
	args := os.Args[1:]
	log.V(0).Info("Args", "args", args)

	// Spectrum
	spectrum.Run(log, args)

	// Mux
	isReady = true
	log.V(0).Info("Serving Mux on :8080", "mux", r)
	log.Error(http.ListenAndServe(":8080", handler), "Unexpected error in mux.")
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
