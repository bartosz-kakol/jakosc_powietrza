package sensor

import (
	"gobot.io/x/gobot/platforms/raspi"
	"jakosc_powietrza/base"
	"log"
	"math/rand"
)

type MockSensor struct {
}

//goland:noinspection GoUnusedParameter
func (s *MockSensor) Init(adaptor *raspi.Adaptor) error {
	log.Println("Init MockSensor")

	return nil
}

func (s *MockSensor) Read() (*base.SensorReadout, error) {
	readout := &base.SensorReadout{
		ECO2: 1400 + rand.Intn(800) - 400,
		TVOC: 400 + rand.Intn(400) - 200,
	}

	return readout, nil
}
