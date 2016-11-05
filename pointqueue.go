package main

import (
	"container/heap"
)

type Member struct {
	Point    Point
	Priority int
	Index    int
}

type Queue []*Member

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
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
	n := len(old)
	member := old[n-1]
	member.Index = -1
	*q = old[0 : n-1]
	return member
}

func (q *Queue) update(member *Member, point Point, priority int) {
	member.Point = point
	member.Priority = priority
	heap.Fix(q, member.Index)
}
