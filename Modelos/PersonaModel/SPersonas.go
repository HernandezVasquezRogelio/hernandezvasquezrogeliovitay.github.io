package PersonaModel

import "time"

type Persona struct {
	IdPersona       string `gorm:"PRIMARY_KEY"`
	Nombre          string
	ApellidoPat     string
	ApellidoMat     string
	FechaNacimiento time.Time
	Sexo            string
	FotoPerfil      string
	Email           string
	Telefono        string
	Estado          bool
	FechaRegistro   time.Time
}
