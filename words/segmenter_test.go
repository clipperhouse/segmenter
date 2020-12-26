package words_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/clipperhouse/segmenter/words"
)

func TestSegmenter(t *testing.T) {
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

func BenchmarkSegmenter(b *testing.B) {
	file, err := ioutil.ReadFile("testdata/sample.txt")

	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	b.SetBytes(int64(len(file)))

	seg := words.NewSegmenter(file)

	for i := 0; i < b.N; i++ {
		seg.SetText(file)

		c := 0
		for seg.Next() {
			c++
		}
		if err := seg.Err(); err != nil {
			b.Error(err)
		}

		b.ReportMetric(float64(c), "tokens")
	}
}
