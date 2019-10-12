package persona

// Persona representa a una persona que trabaja en la empresa
// tiene un sueledo
type Persona struct {
	sueldo float64
}

// New construye una nueva persona
func New(sueldo float64) *Persona {
	return &Persona{sueldo: sueldo}
}

func (p *Persona) Sueldo() float64 {
	return p.sueldo
}
