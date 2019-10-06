package proyecto_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/stretchr/testify/assert"
)

func TestUnProyectoSinPersonasRequeridasNoTieneFitness(t *testing.T) {

	unProyecto := proyecto.New(0, 1.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoSinPersonasAsignadasTieneFitnessNegativo(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConIgualCantidadDePersonasRequeridasQueAsignadasTieneFitnessMayorOIgualQueCero(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)
	unProyecto.AsignarPersona(1.0)

	fitness, _ := unProyecto.Fitness()

	assert.GreaterOrEqual(t, fitness, 0.0, "El fitness del proyecto debería ser mayor o igual que cero")
}

func TestUnProyectoSinPresupuestoAsignadoNoTieneFitness(t *testing.T) {

	unProyecto := proyecto.New(1, 0.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoQueSeExcedeDelPresupuestoTieneFitnessNegativo(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)
	unProyecto.AsignarPersona(2.0)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConMenorGastoDeSueldosTieneMejorFitnessQueUnoConMayorGastoDeSueldos(t *testing.T) {

	proyectoQueGastaMenos := proyecto.New(1, 1.0)
	proyectoQueGastaMenos.AsignarPersona(0.7)
	proyectoQueGastaMas := proyecto.New(1, 1.0)
	proyectoQueGastaMas.AsignarPersona(0.9)

	fitnessDelProyectoQueGastaMenos, _ := proyectoQueGastaMenos.Fitness()
	fitnessDelProyectoQueGastaMas, _ := proyectoQueGastaMas.Fitness()

	assert.Greater(t, fitnessDelProyectoQueGastaMenos, fitnessDelProyectoQueGastaMas, "El fitness del proyecto más barato debería ser mayor que el del más caro")

}
