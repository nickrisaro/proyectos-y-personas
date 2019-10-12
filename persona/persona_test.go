package persona_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/stretchr/testify/assert"
)

func TestUnaPersonaTieneSueldo(t *testing.T) {
	ana := persona.New(1.0)

	assert.Equal(t, 1.0, ana.Sueldo(), "El sueldo de Ana no es el esperado")
}
