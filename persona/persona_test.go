package persona_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/stretchr/testify/assert"
)

func TestUnaPersonaTieneNombre(t *testing.T) {
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)

	assert.Equal(t, "Ana", ana.Nombre(), "El nombre de Ana no es el esperado")
}

func TestUnaPersonaTieneSueldo(t *testing.T) {
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)

	assert.Equal(t, 1.0, ana.Sueldo(), "El sueldo de Ana no es el esperado")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)

	assert.Less(t, int(juan.Seniority()), int(ana.Seniority()), "El seniority de Juan debe ser menor que el de Ana")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSemiSenior(t *testing.T) {
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	maria := persona.New("María", 1.0, persona.SemiSenior, persona.Diseño, persona.Mentoreo)

	assert.Less(t, int(juan.Seniority()), int(maria.Seniority()), "El seniority de Juan debe ser menor que el de María")
}

func TestUnaPersonaSemiSeniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	maria := persona.New("María", 1.0, persona.SemiSenior, persona.Diseño, persona.Mentoreo)
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)

	assert.Less(t, int(maria.Seniority()), int(ana.Seniority()), "El seniority de María debe ser menor que el de Ana")
}

func TestDistintasPersonasTienenDistintosHardSkills(t *testing.T) {
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	maria := persona.New("María", 1.0, persona.SemiSenior, persona.Diseño, persona.Mentoreo)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)

	assert.Equal(t, persona.Desarrollo, ana.HardSkill(), "El HardSkill de Ana debería ser Desarrollo")
	assert.Equal(t, persona.Diseño, maria.HardSkill(), "El HardSkill de María debería ser Diseño")
	assert.Equal(t, persona.Operaciones, juan.HardSkill(), "El HardSkill de Juan debería ser Operaciones")

}

func TestDistintasPersonasTienenDistintosSoftSkills(t *testing.T) {
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	maria := persona.New("María", 1.0, persona.SemiSenior, persona.Diseño, persona.Mentoreo)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)

	assert.Equal(t, persona.Negociacion, ana.SoftSkill(), "El SoftSkill de Ana debería ser Negociacion")
	assert.Equal(t, persona.Mentoreo, maria.SoftSkill(), "El SoftSkill de María debería ser Mentoreo")
	assert.Equal(t, persona.Investigacion, juan.SoftSkill(), "El SoftSkill de Juan debería ser Investigacion")

}
