package segmenter

type Forward interface {
	SetText(data []byte)
	Next() bool
	Start() int
	End() int
	Err() error
}

type Bidirectional interface {
	Forward
	Previous() bool
}
