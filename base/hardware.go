package base

import "gobot.io/x/gobot/platforms/raspi"

type SensorReadout struct {
	ECO2 int // Estimated CO2 levels (ppm units)
	TVOC int // Total volatile organic compounds levels (ppb units)
}

type ISensor interface {
	Init(adaptor *raspi.Adaptor) error
	Read() (*SensorReadout, error)
}

type IDisplay interface {
	Init(adaptor *raspi.Adaptor) error
	Show(data SensorReadout, judgement Judgement) error
}
