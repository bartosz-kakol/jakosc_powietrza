package base

type JudgementValue uint8

const (
	JudgementNone     JudgementValue = iota // Indicates "Room ventilation is not needed"
	JudgementMedium                         // Indicates "Room ventilation is recommended"
	JudgementCritical                       // Indicates "Room ventilation is required"
)

func (j JudgementValue) String() string {
	switch j {
	case JudgementNone:
		return "None"
	case JudgementMedium:
		return "Medium"
	case JudgementCritical:
		return "Critical"
	default:
		return "Unknown"
	}
}

type Judgement struct {
	CO2  JudgementValue
	TVOC JudgementValue
}

type Judge interface {
	JudgeReadout(readout *SensorReadout) *Judgement
}
