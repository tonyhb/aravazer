package main

func NewCluster() *Cluster {
	return &Cluster{
		Items: map[string]Item{},
	}
}

type Cluster struct {
	Items         map[string]Item
	TotalScore    int
	TotalPickaxes int
	TotalEmpties  int

	center Point
}

func (c *Cluster) Add(i Item) {
	if c.Contains(i) {
		return
	}

	c.Items[i.ToString()] = i
	c.TotalScore += i.Score
	if i.IsPickaxe {
		c.TotalPickaxes++
	}
	if i.IsEmpty() {
		c.TotalEmpties++
	}
}

func (c *Cluster) CalculateCenter() Point {
	// memoize
	if c.center.X != 0 && c.center.Y != 0 {
		return c.center
	}

	var x, y int
	for _, v := range c.Items {
		x += v.X
		y += v.Y
	}

	if x > 0 {
		x = x / len(c.Items)
	}
	if y > 0 {
		y = y / len(c.Items)
	}

	c.center = Point{x, y}
	return c.center
}

func (c *Cluster) Contains(i Item) bool {
	_, ok := c.Items[i.ToString()]
	return ok
}

func (c Cluster) Score() int {
	return c.TotalScore * c.TotalPickaxes
}

type ClusterList []Cluster

// TODO: optimize this...
func (cl ClusterList) Contains(i Item) bool {
	for _, list := range cl {
		if list.Contains(i) {
			return true
		}
	}
	return false
}

func (cl ClusterList) Best() Cluster {
	var highestCluster Cluster
	var maxScore int

	for _, v := range cl {
		// TODO: Find healthiest cluster
		if v.Score() > maxScore {
			maxScore = v.Score()
			highestCluster = v
		}
	}

	return highestCluster
}
