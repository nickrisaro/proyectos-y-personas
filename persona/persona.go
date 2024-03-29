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
	Desarrollo HardSkill = iota
	// Diseño habilidad de una persona para realizar el diseño de una web
	Diseño
	// Operaciones habilidad de una persona para asegurarse de que una aplicación funcione correctamente
	Operaciones
)

// SoftSkill representa las habilidades técnicas que puede tener una persona
type SoftSkill int

const (
	// Negociacion habilidad de una persona para negociar con otra persona
	Negociacion SoftSkill = iota
	// Mentoreo habilidad de una persona para enseñar a otra persona
	Mentoreo
	// Investigacion habilidad de una persona para investigar nuevas tecnologías
	Investigacion
)

// Persona representa a una persona que trabaja en la empresa
// tiene un sueledo
type Persona struct {
	ID         int       `json:"id"`
	Anombre    string    `json:"nombre"`
	Asueldo    float64   `json:"sueldo"`
	Aseniority Seniority `json:"seniority"`
	AhardSkill HardSkill `json:"hardSkill"`
	AsoftSkill SoftSkill `json:"softSkill"`
}

// New construye una nueva persona
func New(nombre string, sueldo float64, seniority Seniority, hardSkill HardSkill, softSkill SoftSkill) *Persona {
	return &Persona{
		Anombre:    nombre,
		Asueldo:    sueldo,
		Aseniority: seniority,
		AhardSkill: hardSkill,
		AsoftSkill: softSkill,
	}
}

// Sueldo representa cuanto gana una persona
func (p *Persona) Sueldo() float64 {
	return p.Asueldo
}

// Seniority indica cuánto sabe una persona sobre su hard skill
func (p *Persona) Seniority() Seniority {
	return p.Aseniority
}

// HardSkill inidica cuál es la habilidad técnica de la persona
func (p *Persona) HardSkill() HardSkill {
	return p.AhardSkill
}

// SoftSkill inidica cuál es la habilidad blanda de la persona
func (p *Persona) SoftSkill() SoftSkill {
	return p.AsoftSkill
}

// Nombre devuelve el nombre de la persona
func (p *Persona) Nombre() string {
	return p.Anombre
}
