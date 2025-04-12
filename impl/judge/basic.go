package judge

import (
	"jakosc_powietrza/base"
)

type BasicJudge struct {
	MediumCO2Level    int
	CriticalCO2Level  int
	MediumTVOCLevel   int
	CriticalTVOCLevel int
}

func (j *BasicJudge) JudgeReadout(readout *base.SensorReadout) *base.Judgement {
	judgement := &base.Judgement{}

	if readout.ECO2 >= j.CriticalCO2Level {
		judgement.CO2 = base.JudgementCritical
	} else if readout.ECO2 >= j.MediumCO2Level {
		judgement.CO2 = base.JudgementMedium
	} else {
		judgement.CO2 = base.JudgementNone
	}

	if readout.TVOC >= j.CriticalTVOCLevel {
		judgement.TVOC = base.JudgementCritical
	} else if readout.TVOC >= j.MediumTVOCLevel {
		judgement.TVOC = base.JudgementMedium
	} else {
		judgement.TVOC = base.JudgementNone
	}

	return judgement
}
