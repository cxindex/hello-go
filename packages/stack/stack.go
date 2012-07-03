package stack

type Stack struct {
	I int
	A [10]int
}

func (st *Stack) Push(x int) int {
	if st.I < len(st.A) {
		st.A[st.I] = x
		st.I++
		return 0
	} else {
		println("Stack is full")
	}
	return 1
}

func (st *Stack) Pop() int {
	if st.I > 0 {
		st.I--
		st.A[st.I] = 0
	}
	return 0
}

