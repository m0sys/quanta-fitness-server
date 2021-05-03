package units

type Kilogram float64
type Pound float64

const (
	KGToLBConvFact = 2.20462
	LBToKGConvFact = 0.453592
)

func (kg Kilogram) KGToLB() Pound {
	return Pound(kg * KGToLBConvFact)
}

func (lb Pound) LBToKG() Kilogram {
	return Kilogram(lb * LBToKGConvFact)
}
