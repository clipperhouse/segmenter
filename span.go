package segmenter

type span struct {
	start, end int
}

var zero = span{}

type stack struct {
	items []span
}

func (st *stack) len() int {
	return len(st.items)
}

func (st *stack) push(sp span) {
	st.items = append(st.items, sp)
}

func (st *stack) pop() (span, bool) {
	if len(st.items) == 0 {
		return zero, false
	}

	last := len(st.items) - 1
	result := st.items[last]
	st.items = st.items[:last]

	return result, true
}

func (st *stack) clear() {
	st.items = nil
}
