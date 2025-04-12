package base

import "gobot.io/x/gobot/platforms/raspi"

type SensorReadout struct {
	ECO2 int
	TVOC int
}

type ISensor interface {
	Init(adaptor *raspi.Adaptor) error
	Read() (*SensorReadout, error)
}

type IDisplay interface {
	Init(adaptor *raspi.Adaptor) error
	Show(data SensorReadout) error
}
