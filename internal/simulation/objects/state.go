package objects

type Factors struct {
	MinSpeed         int64
	MaxSpeed         int64
	Fov              float64
	Separation       float64
	BoundaryDistance float32
	BoundaryFactor   float32
	BoundaryScale    float32
	AvoidanceScale   float32
	AlignmentScale   float32
	CohesionScale    float32
}

type State struct {
	Factors        Factors
	ShouldSeparate bool
	ShouldAlign    bool
	ShouldCohesion bool
}

func CreateState() State {
	return State{
		Factors: Factors{
			MinSpeed:         2,
			MaxSpeed:         6,
			Fov:              100,
			Separation:       50,
			BoundaryDistance: 10,
			BoundaryFactor:   150,
			BoundaryScale:    7.00,
			AvoidanceScale:   4.50,
			AlignmentScale:   2.50,
			CohesionScale:    2.70,
		},
		ShouldSeparate: true,
		ShouldAlign:    true,
		ShouldCohesion: true,
	}
}
