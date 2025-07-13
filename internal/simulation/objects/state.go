package objects

type Factors struct {
	Fov             float64
	Separation      float64
	BoundaryDistance float32
	BoundaryFactor  float32
	BoundaryScale   float32
	SeparationScale float32
	AlignmentScale  float32
	CohesionScale   float32
}

type State struct {
	Factors          Factors
	ShouldSeparate   bool
	ShouldAlign      bool
	ShouldCohesion   bool
}

func CreateState() State {
	return State{
		Factors: Factors{
			Fov:             125,
			Separation:      50,
			BoundaryDistance: 80,
			BoundaryFactor:  1,
			BoundaryScale:   50,
			SeparationScale: 0.15,
			AlignmentScale:  0.25,
			CohesionScale:   0.25,
		},
		ShouldSeparate: true,
		ShouldAlign:    true,
		ShouldCohesion: true,
	}
}

/*
BoundaryDistance float32 = 80
	Fov              float64 = 125
	Separation       float64 = 50
	BoundaryFactor   float32 = 1
	BoundaryScale    float32 = 50
	SeparationScale  float32 = 0.15
	AlignmentScale   float32 = 0.25
	CohesionScale    float32 = 0.25
	ShouldSeparate   bool    = true
	ShouldAlign      bool    = true
	ShouldCohesion   bool    = true
*/
