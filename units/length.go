package units

type Meter float64
type Mile float64
type Foot float64

const (
	MeterToFootConvFact = 3.28084
	FootToMeterConvFact = 0.3048
	MeterToMileConvFact = 0.000621371
	MileToMeterConvFact = 1609.34
)

func (m Meter) MeterToMile() Mile {
	return Mile(m * MeterToMileConvFact)
}

func (m Mile) MileToMeter() Meter {
	return Meter(m * MileToMeterConvFact)
}

func (m Meter) MeterToFoot() Foot {
	return Foot(m * MeterToFootConvFact)
}

func (f Foot) FootToMeter() Meter {
	return Meter(f * FootToMeterConvFact)
}
