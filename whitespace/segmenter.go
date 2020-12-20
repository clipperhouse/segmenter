package whitespace

import (
	"errors"
	"unicode"
	"unicode/utf8"

	"github.com/clipperhouse/segmenter"
)

func NewSegmenter(p []byte) *segmenter.Segmenter {
	segment := segmenter.New(SegmentFunc)
	segment.SetText(p)
	return segment
}

var ErrInvalidUTF8 = errors.New("invalid UTF8")

func SegmentFunc(data []byte, atEOF bool) (start int, end int, err error) {
	pos := 0

	// Skip leading whitespace
	for pos < len(data) {
		r, w := utf8.DecodeRune(data[pos:])

		if r == utf8.RuneError {
			return start, end, ErrInvalidUTF8
		}

		if !unicode.IsSpace(r) {
			break
		}

		pos += w
	}

	start = pos

	// Continue through non-whitespace
	for pos < len(data) {
		r, w := utf8.DecodeRune(data[pos:])

		if r == utf8.RuneError {
			return start, end, ErrInvalidUTF8
		}

		if unicode.IsSpace(r) {
			break
		}

		pos += w
	}

	end = pos

	return start, end, nil
}
