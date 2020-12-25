package segmenter_test

import (
	"fmt"
	"testing"

	"github.com/clipperhouse/segmenter"
	"github.com/clipperhouse/segmenter/whitespace"
)

func TestUnicodeWords(t *testing.T) {
	segment := segmenter.New(whitespace.SegmentFunc)
	segment.SetText([]byte("hi   how are you!!  \nand more\r"))

	for segment.Next() {
		fmt.Printf("%q\n", segment.Bytes())
	}

	if err := segment.Err(); err != nil {
		t.Error(err)
	}

	segment.SetText([]byte("Let's try previous"))
	if segment.Next() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Next() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Next() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Previous() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Previous() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Previous() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Previous() {
		fmt.Printf("%q\n", segment.Bytes())
	}
	if segment.Next() {
		fmt.Printf("%q\n", segment.Bytes())
	}
}
