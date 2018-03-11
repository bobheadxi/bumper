package models

import (
	"math"
)

// Junk related constants
const (
	JunkFriction = 0.99
	MinimumBump  = 1.5
	BumpFactor   = 1.05
	JunkSize     = 15
)

// Junk a position and velocity struct describing it's state and player struct to identify rewarding points
type Junk struct {
	Position Position `json:"position"`
	Velocity Velocity `json:"velocity"`
	Color    string   `json:"color"`
	ID       int      `json:"int"` // player id who last hit the junk
}

//Update Junk's position based on calculations of position/velocity
func (j *Junk) updatePosition(height float64, width float64) {
	const r = JunkSize / 2
	if j.Position.X+j.Velocity.Dx > width-r || j.Position.X+j.Velocity.Dx < r {
		j.Velocity.Dx = -j.Velocity.Dx
	}
	if j.Position.Y+j.Velocity.Dy > height-r || j.Position.Y+j.Velocity.Dy < r {
		j.Velocity.Dy = -j.Velocity.Dy
	}

	j.Velocity.Dx *= JunkFriction
	j.Velocity.Dy *= JunkFriction

	j.Position.X += j.Velocity.Dx
	j.Position.Y += j.Velocity.Dy
}

//Update Junks's velocity based on calculations of being hit by a player
func (j *Junk) hitBy(p *Player) {
	j.Color = p.Color

	if p.Velocity.Dx < 0 {
		j.Velocity.Dx = math.Max(p.Velocity.Dx*BumpFactor, -MinimumBump)
	} else {
		j.Velocity.Dx = math.Max(p.Velocity.Dx*BumpFactor, MinimumBump)
	}

	if p.Velocity.Dy < 0 {
		j.Velocity.Dy = math.Max(p.Velocity.Dy*BumpFactor, -MinimumBump)
	} else {
		j.Velocity.Dy = math.Max(p.Velocity.Dy*BumpFactor, MinimumBump)
	}

	p.hitJunk()
}
