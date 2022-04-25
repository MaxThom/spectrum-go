package main

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	//"github.com/go-logr/logr"
	display "github.com/maxthom/spectrum-go/internal/spectrum"
)

var (
	log logr.Logger
)

const ()

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	zapLogger, err := zap.NewDevelopment()
	checkError(err)
	defer zapLogger.Sync()
	log = zapr.NewLogger(zapLogger).WithName("spectrum")

	fmt.Println("----------")
	fmt.Println(" Spectrum ")
	fmt.Println("----------")

	args := os.Args[1:]
	log.V(0).Info("Args", "args", args)

	display.Run(log)
}
