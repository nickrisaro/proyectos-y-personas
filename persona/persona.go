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

// HardSkill representa las habilidades técnicas que puede tener una persona
type HardSkill int

const (
	// Desarrollo habilidad de una persona para desarrollar una funcionalidad
	Desarrollo HardSkill = iota + 1
	// Diseño habilidad de una persona para realizar el diseño de una web
	Diseño
	// Operaciones habilidad de una persona para asegurarse de que una aplicación funcione correctamente
	Operaciones
)

// Persona representa a una persona que trabaja en la empresa
// tiene un sueledo
type Persona struct {
	sueldo    float64
	seniority Seniority
	hardSkill HardSkill
}

// New construye una nueva persona
func New(sueldo float64, seniority Seniority, hardSkill HardSkill) *Persona {
	return &Persona{
		sueldo:    sueldo,
		seniority: seniority,
		hardSkill: hardSkill,
	}
}

// Sueldo representa cuanto gana una persona
func (p *Persona) Sueldo() float64 {
	return p.sueldo
}

// Seniority indica cuánto sabe una persona sobre su hard skill
func (p *Persona) Seniority() Seniority {
	return p.seniority
}

// HardSkill inidica cuál es la habilidad técnica de la persona
func (p *Persona) HardSkill() HardSkill {
	return p.hardSkill
}
