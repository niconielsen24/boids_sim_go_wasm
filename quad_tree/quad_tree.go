package quadtree

type Point struct {
	X uint64
	Y uint64
}

type Node struct {
	Position Point
	Data     int // boid??
}

type Quad struct {
	TopLeft     Point
	BottomRight Point
	Node        *Node

	TopLeftTree  *Node
	TopRightTree *Node
	BotLeftTree  *Node
	BotRightTree *Node
}
