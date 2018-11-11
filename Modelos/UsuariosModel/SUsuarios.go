package UsuariosModel

import "time"

type Perfil struct {
	IdPerfil 			     string 	    `gorm:"PRIMARY_KEY"`
	Tipo                     string
	Estado 				     bool
	FechaRegistro 		     time.Time
}
type Usuario struct {
	IdUsuario			     string         `gorm:"PRIMARY_KEY"`
	IdPersona 				 string
	IdPerfil 				 string
	Estado    	             bool
	FechaRegistro	         time.Time
}

