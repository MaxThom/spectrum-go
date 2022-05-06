package main

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	//"github.com/go-logr/logr"
	display "github.com/maxthom/spectrum-go/internal/spectrum"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

var (
	log  logr.Logger
	quit = make(chan struct{})
)

const ()

func main() {
	zapLogger, err := zap.NewDevelopment()
	utils.CheckError(err)
	defer zapLogger.Sync()
	log = zapr.NewLogger(zapLogger).WithName("spectrum")

	fmt.Println("----------")
	fmt.Println(" Spectrum ")
	fmt.Println("----------")

	args := os.Args[1:]
	log.V(0).Info("Args", "args", args)

	display.Run(log)

	// Then blocking (waiting for quit signal):
	<-quit

	// And in another goroutine if you want to quit:
	// close(quit)
}
