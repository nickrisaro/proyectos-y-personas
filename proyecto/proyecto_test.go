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

func TestUnProyectoConMenorGastoDeSueldosTieneIgualFitnessQueUnoConMayorGastoDeSueldos(t *testing.T) {

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

	assert.InDelta(t, fitnessDelProyectoQueGastaMenos, fitnessDelProyectoQueGastaMas, 0.1, "El fitness del proyecto más barato debería ser mayor que el del más caro")

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

func TestSePuedeObtenerUnResumenDelProyecto(t *testing.T) {

	unInvestigador := persona.New("Mario", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaInvestigadora := persona.New("Ana", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	otroInvestigador := persona.New("Juan", 0.9, persona.SemiSenior, persona.Desarrollo, persona.Investigacion)
	unaMentora := persona.New("Clara", 0.9, persona.SemiSenior, persona.Diseño, persona.Mentoreo)
	unaNegociadora := persona.New("Lucía", 0.9, persona.Senior, persona.Operaciones, persona.Negociacion)

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(3)
	personasRequeridas.Diseño(1)
	personasRequeridas.Operaciones(1)

	miProyecto := proyecto.New("Proyecto uno", personasRequeridas, 3.0)

	miProyecto.AsignarPersona(unInvestigador)
	miProyecto.AsignarPersona(unaMentora)
	miProyecto.AsignarPersona(unaNegociadora)
	miProyecto.AsignarPersona(unaInvestigadora)
	miProyecto.AsignarPersona(otroInvestigador)

	resumen := miProyecto.ObtenerResumen()

	hardSkills := make(map[persona.HardSkill]int)
	hardSkills[persona.Desarrollo] = 3
	hardSkills[persona.Diseño] = 1
	hardSkills[persona.Operaciones] = 1

	softSkills := make(map[persona.SoftSkill]int)
	softSkills[persona.Investigacion] = 3
	softSkills[persona.Mentoreo] = 1
	softSkills[persona.Negociacion] = 1

	seniorities := make(map[persona.Seniority]int)
	seniorities[persona.Junior] = 2
	seniorities[persona.SemiSenior] = 2
	seniorities[persona.Senior] = 1

	assert.NotEmpty(t, resumen, "El proyecto debe tener un resumen")
	assert.Equal(t, "Proyecto uno", resumen.Nombre, "El resumen no contiene información del nombre del proyecto")
	assert.Equal(t, personasRequeridas, resumen.PersonasRequeridas, "El resumen no contiene información de las personas requeridas")
	assert.Equal(t, 3.0, resumen.Presupuesto, "El resumen no contiene información del presupuesto")
	assert.Equal(t, 4.5, resumen.Sueldos, "El resumen no contiene información de los sueldos")
	assert.Equal(t, hardSkills, resumen.HardSkills, "El resumen no contiene información de los hard skills")
	assert.Equal(t, softSkills, resumen.SoftSkills, "El resumen no contiene información de los soft skills")
	assert.Equal(t, seniorities, resumen.Seniorities, "El resumen no contiene información de los seniorities")
}

func TestUnProyectoQueTieneLasPersonasExactasTieneMejorFitnessQueUnoQueNo(t *testing.T) {

	unDesarrolladorJunior := persona.New("Mario", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unaDesarrolladoraJunior := persona.New("Ana", 0.9, persona.Junior, persona.Desarrollo, persona.Investigacion)
	unDesarrolladorSenior := persona.New("Juan", 0.9, persona.Senior, persona.Desarrollo, persona.Investigacion)
	unaDiseñadoraJunior := persona.New("Clara", 0.9, persona.Junior, persona.Diseño, persona.Investigacion)
	unaOperadoraJunior := persona.New("Lucía", 0.9, persona.Junior, persona.Operaciones, persona.Investigacion)

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	personasRequeridas.Diseño(1)
	personasRequeridas.Operaciones(1)
	proyectoQueCubreTodosLosHardSkills := proyecto.New("Proyecto uno", personasRequeridas, 3.0)
	proyectoQueNoCubreTodosLosHardSkills := proyecto.New("Proyecto dos", personasRequeridas, 5.0)

	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorSenior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaDiseñadoraJunior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaOperadoraJunior)

	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unaDesarrolladoraJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unaDiseñadoraJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unaOperadoraJunior)

	fitnessDelProyectoQueCubreTodosLosHardSkills, _ := proyectoQueCubreTodosLosHardSkills.Fitness()
	fitnessDelProyectoQueNoCubreTodosLosHardSkills, _ := proyectoQueNoCubreTodosLosHardSkills.Fitness()

	assert.Less(t, fitnessDelProyectoQueNoCubreTodosLosHardSkills, fitnessDelProyectoQueCubreTodosLosHardSkills, "El fitness del proyecto que tiene exactamente los HardSkills debería ser mayor que el del que no los tiene exactamente")
}
