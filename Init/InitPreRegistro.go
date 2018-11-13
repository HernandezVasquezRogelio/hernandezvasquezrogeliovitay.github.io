package Init

import (
	"github.com/rogelio/Agentes/Conexion"
	"github.com/rogelio/Agentes/Modelos/PreRegistro"
)

func InitPreRegistros() error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()

	if !db.HasTable(&PreRegistro.PreRegistro{}){
		if err := db.CreateTable(&PreRegistro.PreRegistro{}).Error; err != nil{
			return  err
		}
	}
	return nil
}
