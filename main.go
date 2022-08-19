package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	zapLogger, err := zap.NewProduction()
	// Application MUSTN'T start without
	// log setup.
	if err != nil {
		log.Fatal(err.Error())
	}

	// ReplaceGlobals isn't recommended, but in this
	// project, I'm using it to avoid injecting it
	// in every struct and/or passing it as param.
	undoLogger := zap.ReplaceGlobals(zapLogger)

	defer undoLogger()
}
