package main

import (
	"fmt"
	"math"
)

// State stores the state of the current challenge, moving our character
// to solve the problem.
//
// Step 1: Run cluster analysis so that all points are organized into
// clusters, allowing us to choose which cluster to visit with our
// limited number of moves
//
// Step 2: Choose which clusters we could visit without pickaxes,
// and an approx. average of the points we could get
//
// Step 3: Calculate the estimated number of moves by gathering N
// pickaxes before attacking a cluster
//
// Step 4: Once we've found the best estimate run A* pathfinding
// on this badboy to grab the pickaxe(s) then visit our clusters
// to garner all those points.
type State struct {
	*Challenge

	CurrentPoint   Point
	Moves          []string
	RemainingMoves int
	Clusters       ClusterList

	Targets []Point
}

// Given we only have a limited number of moves, calculate the distance
// to the top cluster of points and its overall size so that we know
// how many moves we have remaining, then plan if we can pick up any
// axes.
func (s *State) GetTargets() {
	var currentCluster Cluster
	var score, distance int

	for _, cluster := range s.Clusters {
		// Is the distance greater than the number of moves? If so we discard
		// this cluster
		center := cluster.CalculateCenter()
		dist := center.ManhattanDistance(s.Start)
		if dist > s.RemainingMoves {
			continue
		}

		if cluster.Score() > score {
			currentCluster = cluster
			score = cluster.Score()
			distance = dist
		}
	}

	// This cluster is our target.
	fmt.Println(currentCluster, score, distance)

	// Figure out if there are any pickaxes nearby.

}
