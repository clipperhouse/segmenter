package segmenter

import "errors"

// SegmentFunc is like bufio.SplitFunc, but without an error return value
type SegmentFunc func(data []byte, atEOF bool) (start int, end int, err error)

var ErrIncompleteRune = errors.New("incomplete rune")
var ErrIncompleteToken = errors.New("incomplete token")

func AsSplitFunc(f SegmentFunc, data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}

	start, end, err := f(data, atEOF)

	if errors.Is(err, ErrIncompleteRune) && !atEOF {
		// Rune extends past current data, request more
		return 0, nil, nil
	}

	if errors.Is(err, ErrIncompleteToken) && !atEOF {
		// Token extends past current data, request more
		return 0, nil, nil
	}

	return end, data[start:end], err
}

type span struct {
	start, end int
}

// Segmenter is an iterator for byte arrays
type Segmenter struct {
	data    []byte
	segment SegmentFunc
	// a stack of spans
	spans []span
	err   error
}

func New(segment SegmentFunc) *Segmenter {
	return &Segmenter{
		segment: segment,
	}
}

func (sc *Segmenter) Previous() bool {
	if len(sc.spans) < 2 {
		return false
	}

	sc.spans = sc.spans[:len(sc.spans)-1]

	return true
}

// Next advances the Segmenter to the next token. It returns false when the
// scan reaches the end of the input.
func (sc *Segmenter) Next() bool {
	if sc.current().end == len(sc.data) {
		return false
	}

	start, end, err := sc.segment(sc.data[sc.current().end:], true)

	current := span{
		start: sc.current().end + start,
		end:   sc.current().end + end,
	}

	sc.spans = append(sc.spans, current)
	sc.err = err

	return sc.err == nil && end != start
}

func (sc *Segmenter) SetText(data []byte) {
	sc.data = data
	sc.spans = nil
	sc.err = nil
}

var empty = span{0, 0}

func (sc *Segmenter) current() span {
	if len(sc.spans) == 0 {
		return empty
	}
	last := len(sc.spans) - 1
	return sc.spans[last]
}

func (sc *Segmenter) Start() int {
	return sc.current().start
}

func (sc *Segmenter) End() int {
	return sc.current().end
}

func (sc *Segmenter) Err() error {
	return sc.err
}

// Bytes returns the most recent token generated by a call to Scan.
func (sc *Segmenter) Bytes() []byte {
	return sc.data[sc.current().start:sc.current().end]
}