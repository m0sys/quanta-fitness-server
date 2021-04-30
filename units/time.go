package units

type Second float64
type Minute float64

func (s Second) SecondToMinute() Minute {
	return Minute(s / 60)
}

func (m Minute) MinuteToSecond() Second {
	return Second(m * 60)
}
