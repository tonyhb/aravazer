package main

import (
	"fmt"
	"time"
)

var (
	_ = fmt.Println
	_ = time.Second
)

func main() {
	c := NewChallenge(".2.1.33.s.7.p,.7.71.6.2.63.,3551...9.....,.745668.7.2..,5..45...p..4.,8..4p.529..99,85...2.5....9,..1..69.....3,28.16p1.83.5.,2....7.....9.,.1.p.......58,.pp3...p5..61,...2...6p.769,....3.7.633.p,.......585..5")
	c.Moves = 11

	state := State{
		Challenge: c,
		Clusters:  c.Cluster(),
	}

	state.
		GetTargets().
		WalkToPickaxe().
		ConsumeClusters().
		LimitMoves()

	points := state.Points()
	moves := state.Moves()

	fmt.Println(moves, points)
}
