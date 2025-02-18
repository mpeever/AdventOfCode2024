package lib

import (
	"errors"
	"fmt"
	"log/slog"
)

type Orientation int

const (
	NORTH Orientation = iota
	EAST
	SOUTH
	WEST
)

const (
	FACE_NORTH string = "^"
	FACE_EAST  string = ">"
	FACE_SOUTH string = "v"
	FACE_WEST  string = "<"
)

type GuardState struct {
	Position    Point
	Orientation Orientation
}

func (gs *GuardState) String() string {
	var orientation string
	switch gs.Orientation {
	case NORTH:
		orientation = FACE_NORTH
	case EAST:
		orientation = FACE_EAST
	case SOUTH:
		orientation = FACE_SOUTH
	case WEST:
		orientation = FACE_WEST
	}
	return fmt.Sprintf("(%d, %d), %s", gs.Position.X, gs.Position.Y, orientation)
}

type Guard struct {
	Orientation Orientation
	Grid        CharacterGrid
	Position    Point
	Path        []Point
	Icon        string
	Cache       map[string]GuardState
}

func (gs *Guard) State() GuardState {
	return GuardState{
		Orientation: gs.Orientation,
		Position:    gs.Position,
	}
}

func (g *Guard) TurnRight() Orientation {
	g.Orientation = (g.Orientation + 1) % 4
	return g.Orientation
}

func (g *Guard) TurnLeft() Orientation {
	g.Orientation = (g.Orientation + 3) % 4
	return g.Orientation
}

func (g *Guard) NextPosition() (p Point, err error) {
	switch g.Orientation {
	case NORTH:
		p, err = g.Grid.NextPoint(g.Position, UP)

	case EAST:
		p, err = g.Grid.NextPoint(g.Position, RIGHT)

	case SOUTH:
		p, err = g.Grid.NextPoint(g.Position, DOWN)

	case WEST:
		p, err = g.Grid.NextPoint(g.Position, LEFT)

	default:
		err = errors.New("Invalid Orientation")
	}

	return
}

// Move a move is either a change in location, or a change in orientation.
func (g *Guard) Move() (p Point, o Orientation, err error) {
	p, err = g.NextPosition()
	if err != nil {
		// we've probably stepped off the grid here
		return
	}

	if g.Grid.Char(p) == "#" {
		slog.Debug("found obstacle, turning RIGHT", "orientation", g.Orientation, "current position", g.Position)
		g.TurnRight()
	}

	if g.Grid.Char(p) == "." {
		slog.Debug("found no obstacle, stepping forward", "orientation", g.Orientation, "current position", g.Position, "next point", p)
		g.Path = append(g.Path, p)
		g.Position = p
	}

	return
}

// WalkToEdge - keep going until you step off the map.
// onLoopDetected will fire if we detect a loop, we then abort unless onLoopDetected returns true
func (g *Guard) WalkToEdge(onLoopDetected func(state *GuardState) bool) {
	var err error

	posn, or, err := g.Move()
	if err != nil {
		// we got an error, we probably stepped off the map
		return
	}

	slog.Debug("guard has moved", "orientation", or, "current position", posn)
	status := g.State()
	key := status.String()

	if _, ok := g.Cache[key]; ok {
		ok = onLoopDetected(&status)
		if !ok {
			return
		}
	}

	g.Cache[key] = status

	g.WalkToEdge(onLoopDetected)
}
