package UsuariosModel

import "time"

type Usuario struct {
	IdUsuario			     string         `gorm:"PRIMARY_KEY"`
	IdPersona 				 string
	IdPerfil 				 string
	Estado    	             bool
	FechaRegistro	         time.Time
}

