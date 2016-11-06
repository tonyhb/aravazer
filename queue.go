package main

type Member struct {
	Item     interface{}
	Priority int
	Index    int
}

type Queue []*Member

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	// In our queue we prioritize points based on the benefit of moving to them
	// as detected by our heuristic in state.go.  The larger this number the
	// greater the priority, therefore ensure i is greater than J to pop the
	// higher priority from the beginning of the queue.
	return q[i].Priority > q[j].Priority
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].Index = i
	q[j].Index = j
}

func (q *Queue) Push(i interface{}) {
	n := len(*q)
	member := i.(*Member)
	member.Index = n
	*q = append(*q, member)
}

func (q *Queue) Pop() interface{} {
	old := *q
	member := old[0]
	member.Index = -1
	*q = old[1:]
	return member
}
