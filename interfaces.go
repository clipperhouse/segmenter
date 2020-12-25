package segmenter

type Forward interface {
	Next() bool
	Start() int
	End() int
	Err() error
}

type Bidirectional interface {
	Forward
	Previous() bool
}
