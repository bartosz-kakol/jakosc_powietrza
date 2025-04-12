package sensor

import (
	"fmt"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"jakosc_powietrza/base"
)

type SGP30Sensor struct {
	instance *i2c.Connection
}

func (s *SGP30Sensor) Init(adaptor *raspi.Adaptor) error {
	const address = 0x58
	bus := adaptor.GetDefaultBus()
	sgp30, err := adaptor.GetConnection(address, bus)

	if err != nil {
		return err
	}

	// Initialize SGP30
	if _, err := sgp30.Write([]byte{0x20, 0x03}); err != nil {
		return err
	}

	s.instance = &sgp30
	fmt.Println("SGP30 initialized")

	return nil
}

func (s *SGP30Sensor) Read() (*base.SensorReadout, error) {
	data := make([]byte, 6)

	if _, err := (*s.instance).Read(data); err != nil {
		return nil, err
	}

	readout := &base.SensorReadout{
		ECO2: int(data[0])<<8 | int(data[1]),
		TVOC: int(data[2])<<8 | int(data[3]),
	}

	return readout, nil
}
