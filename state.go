package main

import (
	"container/heap"
	"fmt"
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
	Clusters ClusterList

	clusterQueue  *Queue
	targetPickaxe *Point
}

// Given we only have a limited number of moves, calculate the distance
// to the top cluster of points and its overall size so that we know
// how many moves we have remaining, then plan if we can pick up any
// axes.
func (s *State) GetTargets() *State {
	var targetCluster Cluster
	var score, distance int

	// Add our clusters to a priority queue
	s.clusterQueue = &Queue{}
	for _, cluster := range s.Clusters {
		// Is the distance greater than the number of moves? If so we discard
		// this cluster
		center := cluster.CalculateCenter()
		dist := center.ManhattanDistance(s.Start)
		if dist > s.Challenge.Moves {
			continue
		}
		if cluster.Score() > score {
			targetCluster = cluster
			score = cluster.Score()
			distance = dist
		}
		s.clusterQueue.Push(&Member{cluster, cluster.Score(), 0})
	}
	heap.Init(s.clusterQueue)

	// Figure out if there are any pickaxes nearby.
	distance = 0
	var target *Point
	for _, v := range s.PickaxePoints {
		// If this pickaxe is in our target cluster we'll circle back to it.
		if targetCluster.Contains(s.ItemAt(v)) {
			fmt.Printf("Pickaxe in target cluster at %v; skipping\n", v)
			continue
		}

		// Attempt to find the closest pickaxe to the cluster that we can get
		// en-route.
		// TODO: We could probably cluster pickaxes together and attack clusters
		// of pickaxes here.
		dist := v.ManhattanDistance(s.Start) + targetCluster.CalculateCenter().ManhattanDistance(v)
		if distance == 0 || dist < distance {
			distance = dist
			fmt.Printf("Target pickaxe: %v\n", v)
			copy := v
			target = &copy
		}
	}
	s.targetPickaxe = target

	return s
}

func (s *State) WalkToPickaxe() *State {
	return s
}

func (s *State) ConsumeClusters() *State {
	return s
}

func (s *State) LimitMoves() {
}

func (s *State) Points() int {
	return 0
}

func (s *State) Moves() []string {
	return []string{}
}
