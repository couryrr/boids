package objects

type Factors struct {
	Fov              float64
	Separation       float64
	BoundaryDistance float32
	BoundaryFactor   float32
	BoundaryScale    float32
	SeparationScale  float32
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
			Fov:              100,
			Separation:       50,
			BoundaryDistance: 20,
			BoundaryFactor:   1,
			BoundaryScale:    0.15,
			SeparationScale:  0.15,
			AlignmentScale:   0.25,
			CohesionScale:    0.25,
		},
		ShouldSeparate: true,
		ShouldAlign:    true,
		ShouldCohesion: true,
	}
}

