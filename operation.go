package gomen

type Mino int

const (
	M__     = 0
	M_EMPTY = 0
	M_I     = 1
	M_L     = 2
	M_O     = 3
	M_Z     = 4
	M_T     = 5
	M_J     = 6
	M_S     = 7
	M_X     = 8
	M_G     = 8
	M_GRAY  = 8
)

type Rotation int

const (
	R_REVERSE = 0
	R_RIGHT   = 1
	R_R       = 1
	R_CW      = 1
	R_SPAWN   = 2
	R_LEFT    = 3
	R_L       = 3
	R_CCW     = 3
)

var shapes map[Mino]map[Rotation][][2]int

type Operation struct {
	mino     Mino
	rotation Rotation
	x        int
	y        int
}

func NewOperation(mino Mino, rotation Rotation, x int, y int) Operation {
	return Operation{
		mino:     mino,
		rotation: rotation,
		x:        x,
		y:        y,
	}
}

func (op Operation) Shape() [][2]int {
	shape := [][2]int{}
	for _, offs := range shapes[op.mino][op.rotation] {
		shape = append(shape, [2]int{op.x + offs[0], op.y + offs[1]})
	}
	return shape
}

func minoMirror(m Mino) Mino {
	switch m {
	case M_L:
		return M_J
	case M_Z:
		return M_S
	case M_J:
		return M_L
	case M_S:
		return M_Z
	default:
		return m
	}
}

func init() {
	shapes = map[Mino]map[Rotation][][2]int{
		M__: {},
		M_I: {
			R_REVERSE: {{0, 0}, {1, 0}, {-1, 0}, {-2, 0}},
			R_RIGHT:   {{0, 0}, {0, 1}, {0, -1}, {0, -2}},
			R_SPAWN:   {{0, 0}, {-1, 0}, {1, 0}, {2, 0}},
			R_LEFT:    {{0, 0}, {0, -1}, {0, 1}, {0, 2}}},
		M_L: {
			R_REVERSE: {{0, 0}, {1, 0}, {-1, 0}, {-1, -1}},
			R_RIGHT:   {{0, 0}, {0, 1}, {0, -1}, {1, -1}},
			R_SPAWN:   {{0, 0}, {-1, 0}, {1, 0}, {1, 1}},
			R_LEFT:    {{0, 0}, {0, -1}, {0, 1}, {-1, 1}}},
		M_O: {
			R_REVERSE: {{0, 0}, {-1, 0}, {0, -1}, {-1, -1}},
			R_RIGHT:   {{0, 0}, {0, -1}, {1, 0}, {1, -1}},
			R_SPAWN:   {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			R_LEFT:    {{0, 0}, {0, 1}, {-1, 0}, {-1, 1}}},
		M_Z: {
			R_REVERSE: {{0, 0}, {-1, 0}, {0, -1}, {1, -1}},
			R_RIGHT:   {{0, 0}, {0, -1}, {1, 0}, {1, 1}},
			R_SPAWN:   {{0, 0}, {1, 0}, {0, 1}, {-1, 1}},
			R_LEFT:    {{0, 0}, {0, 1}, {-1, 0}, {-1, -1}}},
		M_T: {
			R_REVERSE: {{0, 0}, {1, 0}, {-1, 0}, {0, -1}},
			R_RIGHT:   {{0, 0}, {0, 1}, {0, -1}, {1, 0}},
			R_SPAWN:   {{0, 0}, {-1, 0}, {1, 0}, {0, 1}},
			R_LEFT:    {{0, 0}, {0, -1}, {0, 1}, {-1, 0}}},
		M_J: {
			R_REVERSE: {{0, 0}, {1, 0}, {-1, 0}, {1, -1}},
			R_RIGHT:   {{0, 0}, {0, 1}, {0, -1}, {1, 1}},
			R_SPAWN:   {{0, 0}, {-1, 0}, {1, 0}, {-1, 1}},
			R_LEFT:    {{0, 0}, {0, -1}, {0, 1}, {-1, -1}}},
		M_S: {
			R_REVERSE: {{0, 0}, {1, 0}, {0, -1}, {-1, -1}},
			R_RIGHT:   {{0, 0}, {0, 1}, {1, 0}, {1, -1}},
			R_SPAWN:   {{0, 0}, {-1, 0}, {0, 1}, {1, 1}},
			R_LEFT:    {{0, 0}, {0, -1}, {-1, 0}, {-1, 1}}},
		M_X: {},
	}
}
