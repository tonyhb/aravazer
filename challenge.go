package main

import (
	"strings"
)

type Challenge struct {
	Encoded string
	Layout  [][]Item // row (y) * column (x) grid
	Start   Point
	Width   int
	Height  int
	Moves   int

	// Here we're going to store some basic information used for optimizing our
	// algorithm. We're going to attempt to find the shortest path to a pickaxe
	// (or multiple) to icrease our score. Therefore, we're going to store the
	// location of all pickaxes here for fast lookups.
	PickaxePoints []Point
}

func NewChallenge(layout string) *Challenge {
	c := &Challenge{
		Encoded:       layout,
		PickaxePoints: []Point{},
	}
	c.parseEncoded()
	return c
}

func (c *Challenge) parseEncoded() {
	rows := strings.Split(c.Encoded, ",")

	c.Width = len(rows[0])
	c.Height = len(rows)
	c.Layout = make([][]Item, c.Height, c.Height)

	for i := range rows {
		c.Layout[i] = c.parseRow(rows[i], i)
	}
}

func (c *Challenge) parseRow(row string, y int) []Item {
	result := make([]Item, len(row), len(row))
	for x, val := range row {
		item := NewItemFromASCII(val)
		item.Point = Point{x, y}
		result[x] = item

		if item.IsStart {
			c.Start = item.Point
		}

		if item.IsPickaxe {
			c.PickaxePoints = append(c.PickaxePoints, item.Point)
		}
	}
	return result
}

// WalkFrom is used to do a point-by-point scan of the Challenge
func (c *Challenge) Scan(ingester func(next Item)) {
	for y := range c.Layout {
		for _, item := range c.Layout[y] {
			ingester(item)
		}
	}
}

func (c *Challenge) IsPointValid(p Point) bool {
	if p.X < 0 || p.X >= c.Width {
		return false
	}
	if p.Y < 0 || p.Y >= c.Height {
		return false
	}
	return true
}

func (c *Challenge) ItemAt(p Point) Item {
	return c.Layout[p.Y][p.X]
}

func (c *Challenge) Cluster() ClusterList {
	// Designing a basic clustering algorithm:
	// 1. Iterate through the points on a map;
	// 2. Expand outwards from a point;
	// 3. If expanding outwards includes more points, add themt o cluster
	// 4. Drop values which are too far out from the center (outliers) - TBD
	// 5. Expand outwards until values dropped are below threshold
	//
	// Continue onwards to a new point.
	// We'll use circular-ish clusters to represent our data for speed.
	clusters := ClusterList{}

	// var i int
	c.Scan(func(curr Item) {
		// Ignore these boring old corridors
		if curr.IsEmpty() {
			return
		}

		// If we have no clusters start one
		if len(clusters) == 0 || !clusters.Contains(curr) {
			cluster := NewCluster()
			cluster.Add(curr)
			expandCluster(c, cluster, curr)
			clusters = append(clusters, *cluster)
		}
	})

	return clusters
}

// Potential improvement:
// Use diagonal neighbours and expand each cluster one side at a time,
// then determine whether the added squares increase the cluster score
// enough and push out any old squares too far away from the new center.
// Compare cluster health of each direction and take the best one.
func expandCluster(c *Challenge, cluster *Cluster, from Item) {
	curr := from

	// Expand out of this point in a radius
	neighbours := curr.Expand()
	for _, next := range neighbours {
		if !c.IsPointValid(next) {
			continue
		}

		// We want connected items, so only add the item to the cluster
		// if it's non-empty
		item := c.ItemAt(next)
		if item.IsEmpty() || cluster.Contains(item) {
			continue
		}

		// Only add this to the clsuter if it's new
		cluster.Add(item)
		expandCluster(c, cluster, item)
	}
}
