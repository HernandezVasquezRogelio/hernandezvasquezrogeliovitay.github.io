package AgenteModel

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
type CrudAcreditaciones struct {
	IdAcreditacion    string `gorm:"PRIMARY_KEY"`
	Profesion         string
	CedulaProfesional string
	Titulo            string
	RFC               string
	Estado            string
	FechaRegistro     time.Time
}
type CrudDirecciones struct {
	IdDirecciones  string `gorm:"PRIMARY_KEY"`
	Calle          string
	Colonia        string
	Municipio      string
	NumeroInterior string
	NumeroExterior string
	CodigoPostal   int
	Estado         bool
	FechaRegistro  time.Time
}
type Agente struct {
	IdAgente         string `gorm:"PRIMARY_KEY"`
	IdPersona        string
	CodigoInvitacion string
	Estado           bool
	FechaRegistro    time.Time
}
type AcreditacionAgentes struct {
	IdAcreditacionAgente string `gorm:"PRIMARY_KEY"`
	IdAcreditacion       string
	IdAgente             string
	Estado               bool
	FechaRegistro        time.Time
}
type DireccionPersona struct {
	IdDireccionPersona string `gorm:"PRIMARY_KEY"`
	IdDirecciones      string
	IdPersona          string
	TipoDireccion      string
	Estado             bool
	FechaRegistro      time.Time
}
