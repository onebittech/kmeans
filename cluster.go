package kmeans

// A Cluster which data points gravitate around
type Cluster struct {
	Center Point
	Points Points
}

// Clusters is a slice of clusters
type Clusters []Cluster

// recenter recenters a cluster
func (c *Cluster) recenter() {
	center, err := c.Points.Mean()
	if err != nil {
		return
	}

	c.Center = center
}

// recenter recenters all clusters
func (c Clusters) recenter() {
	for i := 0; i < len(c); i++ {
		c[i].recenter()
	}
}

// reset clears all point assignments
func (c Clusters) reset() {
	for i := 0; i < len(c); i++ {
		c[i].Points = Points{}
	}
}

// Nearest returns the index of the cluster nearest to point
func (c Clusters) Nearest(point Point) int {
	var dist float64
	var ci int

	// Find the nearest cluster for this data point
	for i, cluster := range c {
		d := point.Distance(cluster.Center)
		if dist == 0 || d < dist {
			dist = d
			ci = i
		}
	}

	return ci
}
