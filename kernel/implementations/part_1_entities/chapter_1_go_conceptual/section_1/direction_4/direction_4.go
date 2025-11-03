package direction_4

// Directions4 - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた
type Directions4 int

const (
	East Directions4 = iota
	West
	South
	North
)
