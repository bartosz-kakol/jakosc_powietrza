package display

import (
	"gobot.io/x/gobot/platforms/raspi"
	"jakosc_powietrza/base"
	"log"
)

type MockDisplay struct {
}

//goland:noinspection GoUnusedParameter
func (m *MockDisplay) Init(adaptor *raspi.Adaptor) error {
	log.Println("Init MockDisplay")

	return nil
}

func (m *MockDisplay) Show(data base.SensorReadout, judgement base.Judgement) error {
	log.Printf(
		"eCO2: %d ppm, TVOC: %d ppb\nJudgement - CO2: %s, TVOC: %s",
		data.ECO2, data.TVOC, judgement.CO2.String(), judgement.TVOC.String(),
	)

	return nil
}
