package words_test

import (
	"reflect"
	"testing"

	"github.com/clipperhouse/segmenter/words"
)

func TestSegments(t *testing.T) {
	var passed, failed int
	for _, test := range unicodeTests {

		var got [][]byte
		segment := words.NewSegmenter(test.input)

		for segment.Next() {
			got = append(got, segment.Bytes())
		}

		if err := segment.Err(); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, test.expected) {
			failed++
			t.Errorf(`
for input %v
expected  %v
got       %v
spec      %s`, test.input, test.expected, got, test.comment)
		} else {
			passed++
		}
	}
	t.Logf("passed %d, failed %d", passed, failed)
}
