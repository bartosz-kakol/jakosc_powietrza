package main

import (
	"jakosc_powietrza/base"
	displaypkg "jakosc_powietrza/impl/display"
	sensorpkg "jakosc_powietrza/impl/sensor"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	work := func() {
		var sensor base.ISensor
		var display base.IDisplay

		// Define used hardware
		sensor = &sensorpkg.SGP30Sensor{}
		display = &displaypkg.SH1107Display{}

		// SGP30 sensor setup
		if err := sensor.Init(r); err != nil {
			log.Fatalf("Failed to init sensor: %v", err)
			return
		}

		// SH1107 OLED setup
		if err := display.Init(r); err != nil {
			log.Fatalf("Failed to init display: %v", err)
			return
		}

		for {
			readout, err := sensor.Read()
			if err != nil {
				log.Printf("Error reading air quality data: %v", err)
			}

			if err := display.Show(*readout); err != nil {
				log.Printf("Error writing to OLED: %v", err)
			}

			time.Sleep(1 * time.Second)
		}
	}

	rbt := gobot.NewRobot(
		"StaleAirDetector",
		[]gobot.Connection{r},
		work,
	)

	if err := rbt.Start(); err != nil {
		log.Fatalf("Failed to start robot: %v", err)
	}
}
