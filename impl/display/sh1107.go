package display

import (
	"fmt"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"jakosc_powietrza/base"
	"log"
)

type SH1107Display struct {
	instance *i2c.Connection
}

func (d *SH1107Display) Init(adaptor *raspi.Adaptor) error {
	bus := adaptor.GetDefaultBus()
	const address = 0x3C
	oled, err := adaptor.GetConnection(address, bus)

	if err != nil {
		log.Fatalf("Failed to connect to OLED display: %v", err)
		return err
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
			return err
		}
	}

	d.instance = &oled
	fmt.Println("OLED initialized")

	return nil
}

func (d *SH1107Display) Show(data base.SensorReadout, judgement base.Judgement) error {
	oledCommands := [][]byte{
		append([]byte{0x40}, []byte(fmt.Sprintf("eCO2: %d ppm", data.ECO2))...),
		append([]byte{0x40}, []byte(fmt.Sprintf("TVOC: %d ppb", data.TVOC))...),
	}

	for _, cmd := range oledCommands {
		if _, err := (*d.instance).Write(cmd); err != nil {
			return err
		}
	}

	return nil
}
