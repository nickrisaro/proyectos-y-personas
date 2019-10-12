package persona

// Seniority representa la experiencia de una persona
type Seniority int

const (
	// Junior es una persona con poca experiencia
	Junior Seniority = iota + 1
	// SemiSenior es una persona con un poco de experiencia
	SemiSenior
	// Senior es una persona con mucha experiencia
	Senior
)

// Persona representa a una persona que trabaja en la empresa
// tiene un sueledo
type Persona struct {
	sueldo    float64
	seniority Seniority
}

// New construye una nueva persona
func New(sueldo float64, seniority Seniority) *Persona {
	return &Persona{sueldo: sueldo, seniority: seniority}
}

func (p *Persona) Sueldo() float64 {
	return p.sueldo
}

func (p *Persona) Seniority() Seniority {
	return p.seniority
}
