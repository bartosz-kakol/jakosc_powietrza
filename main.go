package main

import (
	"jakosc_powietrza/base"
	displaypkg "jakosc_powietrza/impl/display"
	judgepkg "jakosc_powietrza/impl/judge"
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
		var judge base.Judge
		var display base.IDisplay

		// Define used hardware
		sensor = &sensorpkg.MockSensor{}
		display = &displaypkg.MockDisplay{}

		// Configure readout data judge
		judge = &judgepkg.BasicJudge{
			MediumCO2Level:    1000,
			CriticalCO2Level:  2000,
			MediumTVOCLevel:   200,
			CriticalTVOCLevel: 500,
		}

		// Sensor setup
		if err := sensor.Init(r); err != nil {
			log.Fatalf("Failed to init sensor: %v", err)
			return
		}

		// OLED display setup
		if err := display.Init(r); err != nil {
			log.Fatalf("Failed to init display: %v", err)
			return
		}

		for {
			readout, err := sensor.Read()
			if err != nil {
				log.Printf("Error reading air quality data: %v", err)
			}

			judgement := judge.JudgeReadout(readout)

			if err := display.Show(*readout, *judgement); err != nil {
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
