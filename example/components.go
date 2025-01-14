package example

// Data
type PositionComponent struct {
	X float64
	Y float64
}

type VelocityComponent struct {
	X float64
	Y float64
}

// Component types
func (c *PositionComponent) Type() string {
	return "position"
}

func (c *VelocityComponent) Type() string {
	return "velocity"
}

// Constructors
func MakePositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		X: x,
		Y: y,
	}
}

func MakeVelocityComponent(x, y float64) *VelocityComponent {
	return &VelocityComponent{
		X: x,
		Y: y,
	}
}
