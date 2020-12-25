An in-progress package to explore bringing text segmentation to [x/text](https://github.com/golang/text) and perhaps bufio.

The top-level package implements a bidirectonal segmenter, meant to address use cases similar to the [ICU's BreakIterator](https://github.com/unicode-org/icu/blob/master/icu4j/eclipse-build/plugins.template/com.ibm.icu.base/src/com/ibm/icu/text/BreakIterator.java).

Here is an initial [design exploration](https://docs.google.com/document/d/1MXgcwq22ySUUTStKOtrDnCYAOA7ARLrwWzCKg_K5sEA/edit).

The `words` package implements [word boundaries](https://unicode.org/reports/tr29/#Word_Boundaries) from UAX #29. Derived from [this package](https://github.com/clipperhouse/uax29).

There is also a `whitespace` segmenter, primarily for testing.