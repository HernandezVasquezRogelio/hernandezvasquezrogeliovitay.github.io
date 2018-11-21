package PerfilModel

import "time"

type Perfil struct {
	IdPerfil 			     string 	    `gorm:"PRIMARY_KEY"`
	Tipo                     string
	Estado 				     bool
	FechaRegistro 		     time.Time
}