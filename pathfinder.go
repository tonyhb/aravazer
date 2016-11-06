package main

import (
	"container/heap"
)

type Pathfinder struct {
	Queue     *Queue
	List      map[Point]Point
	Scores    map[Point]int
	Challenge *Challenge

	From Point
	To   Point
}

func NewPathfinder(from, to Point, visited []Point, c *Challenge) *Pathfinder {
	list := map[Point]Point{}
	for _, v := range visited {
		list[v] = Point{}
	}

	queue := &Queue{&Member{from, 0, 0}}
	heap.Init(queue)

	return &Pathfinder{
		Queue:     queue,
		List:      list,
		Scores:    map[Point]int{from: 0},
		Challenge: c,
		From:      from,
		To:        to,
	}
}

func (p *Pathfinder) AStar() []Point {
	var current Point

	// Pathfind
	for p.Queue.Len() > 0 {
		current = p.Queue.Pop().(*Member).Item.(Point)

		if current == p.To {
			break
		}

		neighbours := current.Expand()
		// A typical A* search algorith adds all neighbours to the priorityqueue
		// and pops off the highest priority, which accounts for backtrack searching
		// while also handling cases for greedy breadth first search.
		//
		// Our algorith doesn't necessarily need this; we know the cluster in which
		// we want to attack, and we know the waypoint we need to go. In this case
		// we can just choose the path with the highest "score" based off of the
		// given priority list.
		for _, v := range neighbours {
			if !p.Challenge.IsPointValid(v) {
				continue
			}

			// Also, if we've already visited this particular square we
			// need to halt
			_, ok := p.List[v]
			if ok {
				continue
			}

			// We should only consider this square if it's a new square
			// or it beats the previous cost of that square.
			_, ok = p.Scores[v]
			score := p.Scores[current] + p.Score(v)

			if !ok || score >= p.Scores[v] {
				priority := p.Priority(v, p.To)
				heap.Push(p.Queue, &Member{v, priority, 0})
				p.List[v] = current
				p.Scores[v] = score
			}
		}

	}

	path := []Point{current}
	for current != p.From {
		current = p.List[current]
		path = append(path, current)
	}

	return p.ReversePoints(path)
}

// Greedy takes the nearest neighbours of the current point and chooses a path
// which nets the most points. It does not necessarily choose the shortest path,
// though it is directional.
func (p *Pathfinder) Greedy() []Point {
	var current Point

	// Pathfind
	for p.Queue.Len() > 0 {
		current = p.Queue.Pop().(*Member).Item.(Point)

		if current == p.To {
			break
		}
		neighbours := current.Expand()
		// A typical A* search algorith adds all neighbours to the priorityqueue
		// and pops off the highest priority, which accounts for backtrack searching
		// while also handling cases for greedy breadth first search.
		//
		// Our algorith doesn't necessarily need this; we know the cluster in which
		// we want to attack, and we know the waypoint we need to go. In this case
		// we can just choose the path with the highest "score" based off of the
		// given priority list.
		var highestPriority, highestScore int
		var next Point
		highestPriority = -1 * (p.Challenge.Width * p.Challenge.Height)
		for _, v := range neighbours {
			if !p.Challenge.IsPointValid(v) {
				continue
			}

			// Also, if we've already visited this particular square we
			// need to halt
			_, ok := p.List[v]
			if ok {
				continue
			}

			// We should only consider this square if it's a new square
			// or it beats the previous cost of that square.
			_, ok = p.Scores[v]
			score := p.Scores[current] + p.Score(v)

			if !ok || score >= p.Scores[v] {
				priority := p.Priority(v, p.To)
				if priority > highestPriority {
					highestPriority = priority
					highestScore = score
					next = v
				}
			}
		}

		heap.Push(p.Queue, &Member{next, highestPriority, 0})
		p.List[next] = current
		p.Scores[next] = highestScore

	}

	path := []Point{current}
	for current != p.From {
		current = p.List[current]
		path = append(path, current)
	}

	return p.ReversePoints(path)
}

func (Pathfinder) ReversePoints(points []Point) []Point {
	result := make([]Point, len(points), len(points))
	for i, v := range points {
		result[len(points)-i-1] = v
	}
	return result
}

func (p *Pathfinder) Score(to Point) int {
	// Cost should potentially return different integers based on the
	// point value of this square.
	// A larger point or pickaxe multiplies the priority of this particular
	// point's score.
	item := p.Challenge.ItemAt(to)
	if item.IsPickaxe {
		return 5
	}
	return item.Score
}

// Priority is our heuristic function for moving to a square. The larger
// the number the better.
func (p *Pathfinder) Priority(next, target Point) int {
	distance := next.ManhattanDistance(target) * -1
	// A larger point or pickaxe multiplies the priority of this particular
	// point's score.
	item := p.Challenge.ItemAt(next)
	if item.IsPickaxe {
		return distance + 5
	}
	return distance + item.Score
}
