package main

import (
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	work := func() {
		bus := r.GetDefaultBus()

		// SGP30 sensor setup
		const sgp30Address = 0x58
		sgp30, err := r.GetConnection(sgp30Address, bus)

		if err != nil {
			log.Fatalf("Failed to connect to SGP30 sensor: %v", err)
			return
		}

		// Initialize SGP30
		if _, err := sgp30.Write([]byte{0x20, 0x03}); err != nil {
			log.Fatalf("Failed to initialize SGP30: %v", err)
			return
		}

		fmt.Println("SGP30 initialized")

		// SH1107 OLED setup
		const oledAddress = 0x3C
		oled, err := r.GetConnection(oledAddress, bus)

		if err != nil {
			log.Fatalf("Failed to connect to OLED display: %v", err)
			return
		}

		// Initialize SH1107 OLED
		oledInitCommands := [][]byte{
			{0xAE},       // Display off
			{0xD5, 0x80}, // Set display clock divide ratio/oscillator frequency
			{0xA8, 0x3F}, // Set multiplex ratio
			{0xD3, 0x00}, // Set display offset
			{0x40},       // Set display start line
			{0x8D, 0x14}, // Enable charge pump regulator
			{0xAF},       // Display on
		}

		for _, cmd := range oledInitCommands {
			if _, err := oled.Write(cmd); err != nil {
				log.Fatalf("Failed to initialize OLED: %v", err)
				return
			}
		}

		fmt.Println("OLED initialized")

		for {
			data := make([]byte, 6)

			if _, err := sgp30.Read(data); err != nil {
				log.Printf("Error reading air quality data: %v", err)
				continue
			}

			eCO2 := int(data[0])<<8 | int(data[1])
			tVOC := int(data[2])<<8 | int(data[3])
			fmt.Printf("eCO2: %d ppm, TVOC: %d ppb\n", eCO2, tVOC)

			oledCommands := [][]byte{
				append([]byte{0x40}, []byte(fmt.Sprintf("eCO2: %d ppm", eCO2))...),
				append([]byte{0x40}, []byte(fmt.Sprintf("TVOC: %d ppb", tVOC))...),
			}

			for _, cmd := range oledCommands {
				if _, err := oled.Write(cmd); err != nil {
					log.Printf("Error writing to OLED: %v", err)
					continue
				}
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
