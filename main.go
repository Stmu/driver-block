package main

import (
	"github.com/davecgh/go-spew/spew"
  "github.com/stmu/driver-block/arduino"
)

func main() {

	var path = "/dev/ttyO1"
	var speed = 9600

	port, err := arduino.Connect(path, speed)

	if err != nil {
		log.Fatalf("Couldn't connect to arduino: %s", err)
	}

	version, err := port.GetVersion()

	if err != nil {
		log.Warningf("Failed to get version from arduino. Continuing anyway. #YOLO.")
	}

	if version != requiredVersion {
		log.Warningf("Unknown arduino version. Expected:%s Got: %s", requiredVersion, version)
	}

	// NewLight(d, 1007, "Nina's Eyes", port)
	// NewLight(d, 999, "Status Light", port)

	go func() {
		for message := range port.Incoming {
			spew.Dump("incoming", message)
		}
	}()
}
