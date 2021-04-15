package record

type SetManger interface {
	Start()
	Complete(actualRepCount int) Set
}
