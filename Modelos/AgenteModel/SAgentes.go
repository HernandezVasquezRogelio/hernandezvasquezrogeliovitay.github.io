package AgenteModel

import "time"

type CrudAcreditaciones struct {
	IdAcreditacion    string `gorm:"PRIMARY_KEY"`
	Profesion         string
	CedulaProfesional string
	Titulo            string
	RFC               string
	Estado            string
	FechaRegistro     time.Time
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
