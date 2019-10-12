package persona_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/stretchr/testify/assert"
)

func TestUnaPersonaTieneSueldo(t *testing.T) {
	ana := persona.New(1.0, persona.Senior, persona.Desarrollo)

	assert.Equal(t, 1.0, ana.Sueldo(), "El sueldo de Ana no es el esperado")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	juan := persona.New(1.0, persona.Junior, persona.Operaciones)
	ana := persona.New(1.0, persona.Senior, persona.Desarrollo)

	assert.Less(t, int(juan.Seniority()), int(ana.Seniority()), "El seniority de Juan debe ser menor que el de Ana")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSemiSenior(t *testing.T) {
	juan := persona.New(1.0, persona.Junior, persona.Operaciones)
	maria := persona.New(1.0, persona.SemiSenior, persona.Diseño)

	assert.Less(t, int(juan.Seniority()), int(maria.Seniority()), "El seniority de Juan debe ser menor que el de María")
}

func TestUnaPersonaSemiSeniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	maria := persona.New(1.0, persona.SemiSenior, persona.Diseño)
	ana := persona.New(1.0, persona.Senior, persona.Desarrollo)

	assert.Less(t, int(maria.Seniority()), int(ana.Seniority()), "El seniority de María debe ser menor que el de Ana")
}

func TestDistintasPersonasTienenDistintosHardSkills(t *testing.T) {
	ana := persona.New(1.0, persona.Senior, persona.Desarrollo)
	maria := persona.New(1.0, persona.SemiSenior, persona.Diseño)
	juan := persona.New(1.0, persona.Junior, persona.Operaciones)

	assert.Equal(t, persona.Desarrollo, ana.HardSkill(), "El HardSkill de Ana debería ser Desarrollo")
	assert.Equal(t, persona.Diseño, maria.HardSkill(), "El HardSkill de María debería ser Diseño")
	assert.Equal(t, persona.Operaciones, juan.HardSkill(), "El HardSkill de Juan debería ser Operaciones")

}
