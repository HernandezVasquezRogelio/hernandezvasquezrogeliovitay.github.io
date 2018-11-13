package PreRegistro

import "time"

type PreRegistro struct {
	Id                string `gorm:"PRIMARY_KEY"`
	Nombre            string
	Apellido          string
	CorreoElectronico string
	Telefono          string
	Contrasena        string
	CodigoInvitacion  string
	Estado            bool
	FechaRegistro     time.Time
}
