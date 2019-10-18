package proyecto_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/stretchr/testify/assert"
)

func TestUnProyectoTieneNombre(t *testing.T) {

	unProyecto := proyecto.New("Proyecto uno", nil, 1.0)

	assert.Equal(t, "Proyecto uno", unProyecto.Nombre(), "El nombre del proyecto es incorrecto")
}

func TestUnProyectoSePuedeClonar(t *testing.T) {
	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)

	clonDelProyecto := unProyecto.Clonar()

	assert.Equal(t, unProyecto, clonDelProyecto, "El clon no es igual al padre")
}

func TestUnProyectoSinPersonasRequeridasNoTieneFitness(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoSinPersonasAsignadasTieneFitnessNegativo(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConIgualCantidadDePersonasRequeridasQueAsignadasTieneFitnessMayorOIgualQueCero(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	unaPersona := persona.New("María", 1.0, persona.Senior, persona.Desarrollo, persona.Investigacion)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.GreaterOrEqual(t, fitness, 0.0, "El fitness del proyecto debería ser mayor o igual que cero")
}

func TestUnProyectoSinPresupuestoAsignadoNoTieneFitness(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 0.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoQueSeExcedeDelPresupuestoTieneFitnessNegativo(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 0.5)
	unaPersona := persona.New("María", 3.0, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConMenorGastoDeSueldosTieneMejorFitnessQueUnoConMayorGastoDeSueldos(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	proyectoQueGastaMenos := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	unaPersona := persona.New("María", 0.7, persona.Senior, persona.Desarrollo, persona.Investigacion)
	proyectoQueGastaMenos.AsignarPersona(unaPersona)
	proyectoQueGastaMas := proyecto.New("Proyecto dos", personasRequeridas, 1.0)
	otraPersona := persona.New("Ana", 0.9, persona.Senior, persona.Desarrollo, persona.Investigacion)
	proyectoQueGastaMas.AsignarPersona(otraPersona)

	fitnessDelProyectoQueGastaMenos, _ := proyectoQueGastaMenos.Fitness()
	fitnessDelProyectoQueGastaMas, _ := proyectoQueGastaMas.Fitness()

	assert.Greater(t, fitnessDelProyectoQueGastaMenos, fitnessDelProyectoQueGastaMas, "El fitness del proyecto más barato debería ser mayor que el del más caro")

}

func TestUnProyectoConPersonasDeMasSeniorityTieneMejorFitnessQueUnoConPersonasDeMenorSeniority(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	proyectoConMenosSeniority := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	unaPersona := persona.New("María", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	proyectoConMenosSeniority.AsignarPersona(unaPersona)
	proyectoConMAyorSeniority := proyecto.New("Proyecto dos", personasRequeridas, 1.0)
	otraPersona := persona.New("Ana", 0.9, persona.Senior, persona.Desarrollo, persona.Investigacion)
	proyectoConMAyorSeniority.AsignarPersona(otraPersona)

	fitnessDelProyectoConMenorSeniority, _ := proyectoConMenosSeniority.Fitness()
	fitnessDelProyectoConMayorSeniority, _ := proyectoConMAyorSeniority.Fitness()

	assert.Less(t, fitnessDelProyectoConMenorSeniority, fitnessDelProyectoConMayorSeniority, "El fitness del proyecto de más seniority debería ser mayor que el del menos")

}

func TestUnProyectoQueCubreTodosLosHardSkillsRequeridosTieneMejorFitnessQueUnoQueNoCubreTodos(t *testing.T) {

	unDesarrolladorJunior := persona.New("Mario", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaDesarrolladoraJunior := persona.New("Ana", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	otroDesarrolladorJunior := persona.New("Juan", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaDiseñadoraJunior := persona.New("Clara", 0.9, persona.Junior, persona.Diseño, persona.Investigacion)
	unaOperadoraJunior := persona.New("Lucía", 0.9, persona.Junior, persona.Operaciones, persona.Investigacion)

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	personasRequeridas.Diseño(1)
	personasRequeridas.Operaciones(1)
	proyectoQueCubreTodosLosHardSkills := proyecto.New("Proyecto uno", personasRequeridas, 3.0)
	proyectoQueNoCubreTodosLosHardSkills := proyecto.New("Proyecto dos", personasRequeridas, 3.0)

	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorJunior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaDiseñadoraJunior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaOperadoraJunior)

	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unaDesarrolladoraJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(otroDesarrolladorJunior)

	fitnessDelProyectoQueCubreTodosLosHardSkills, _ := proyectoQueCubreTodosLosHardSkills.Fitness()
	fitnessDelProyectoQueNoCubreTodosLosHardSkills, _ := proyectoQueNoCubreTodosLosHardSkills.Fitness()

	assert.Less(t, fitnessDelProyectoQueNoCubreTodosLosHardSkills, fitnessDelProyectoQueCubreTodosLosHardSkills, "El fitness del proyecto que cubre todos los  HardSkills debería ser mayor que el del que no los cubre")
}

func TestUnProyectoQueCubreTodosLosSoftSkillsTieneMejorFitnessQueUnoQueNoCubreTodos(t *testing.T) {

	unInvestigador := persona.New("Mario", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaInvestigadora := persona.New("Ana", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	otroInvestigador := persona.New("Juan", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaMentora := persona.New("Clara", 0.9, persona.Junior, persona.Desarrollo, persona.Mentoreo)
	unaNegociadora := persona.New("Lucía", 0.9, persona.Junior, persona.Desarrollo, persona.Negociacion)

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(3)
	proyectoQueCubreTodosLosSoftSkills := proyecto.New("Proyecto uno", personasRequeridas, 3.0)
	proyectoQueNoCubreTodosLosSoftSkills := proyecto.New("Proyecto dos", personasRequeridas, 3.0)

	proyectoQueCubreTodosLosSoftSkills.AsignarPersona(unInvestigador)
	proyectoQueCubreTodosLosSoftSkills.AsignarPersona(unaMentora)
	proyectoQueCubreTodosLosSoftSkills.AsignarPersona(unaNegociadora)

	proyectoQueNoCubreTodosLosSoftSkills.AsignarPersona(unInvestigador)
	proyectoQueNoCubreTodosLosSoftSkills.AsignarPersona(unaInvestigadora)
	proyectoQueNoCubreTodosLosSoftSkills.AsignarPersona(otroInvestigador)

	fitnessDelProyectoQueCubreTodosLosSoftSkills, _ := proyectoQueCubreTodosLosSoftSkills.Fitness()
	fitnessDelProyectoQueNoCubreTodosLosSoftSkills, _ := proyectoQueNoCubreTodosLosSoftSkills.Fitness()

	assert.Less(t, fitnessDelProyectoQueNoCubreTodosLosSoftSkills, fitnessDelProyectoQueCubreTodosLosSoftSkills, "El fitness del proyecto que cubre todos los  SoftSkills debería ser mayor que el del que no los cubre")
}
