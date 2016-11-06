package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	data string
	_    = fmt.Println
)

func init() {
	flag.StringVar(&data, "data", "", "challenge data base64 encoded")
}

func main() {
	flag.Parse()
	data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("error decoding base64:", err)
		os.Exit(1)
	}

	challenges := []*Challenge{}
	if err = json.Unmarshal(data, &challenges); err != nil {
		fmt.Println("error decoding challenge json:", err)
		os.Exit(1)
	}

	var solutions = []string{}
	var total int
	for _, c := range challenges {
		c.ParseEncoded()
		state := NewState(c)
		path := state.
			GetTargets().
			WalkToPickaxe().
			ConsumeClusters().
			LimitMoves()

		points, moves := result(c, path)
		solutions = append(solutions, strings.Join(moves, ""))
		total += points
	}

	fmt.Println("Points:", total)
	fmt.Println("Solutions:", strings.Join(solutions, ","))
}

func result(c *Challenge, path []Point) (int, []string) {
	var (
		pickaxes, points int
		previous         *Point
	)
	moves := make([]string, c.Moves)

	for i, v := range path {
		item := c.ItemAt(v)
		if item.IsPickaxe {
			pickaxes++
		}
		points += item.Score * (pickaxes + 1)

		if previous != nil {
			if previous.X < v.X {
				moves[i-1] = "r"
			}
			if previous.X > v.X {
				moves[i-1] = "l"
			}
			if previous.Y < v.Y {
				moves[i-1] = "d"
			}
			if previous.Y > v.Y {
				moves[i-1] = "u"
			}
		}
		copy := v
		previous = &copy
	}

	return points, moves
}
