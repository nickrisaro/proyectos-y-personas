package persona_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/stretchr/testify/assert"
)

func TestUnaPersonaTieneSueldo(t *testing.T) {
	ana := persona.New(1.0, persona.Senior)

	assert.Equal(t, 1.0, ana.Sueldo(), "El sueldo de Ana no es el esperado")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	juan := persona.New(1.0, persona.Junior)
	ana := persona.New(1.0, persona.Senior)

	assert.Less(t, int(juan.Seniority()), int(ana.Seniority()), "El seniority de Juan debe ser menor que el de Ana")
}

func TestUnaPersonaJuniorTieneMenosSeniorityQueUnaSemiSenior(t *testing.T) {
	juan := persona.New(1.0, persona.Junior)
	maria := persona.New(1.0, persona.SemiSenior)

	assert.Less(t, int(juan.Seniority()), int(maria.Seniority()), "El seniority de Juan debe ser menor que el de María")
}

func TestUnaPersonaSemiSeniorTieneMenosSeniorityQueUnaSenior(t *testing.T) {
	maria := persona.New(1.0, persona.SemiSenior)
	ana := persona.New(1.0, persona.Senior)

	assert.Less(t, int(maria.Seniority()), int(ana.Seniority()), "El seniority de María debe ser menor que el de Ana")
}
