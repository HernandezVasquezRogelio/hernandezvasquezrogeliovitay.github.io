package Init

import (
	"github.com/rogelio/Agentes/Conexion"
	"github.com/rogelio/Agentes/Modelos/UsuariosModel"
)

func InitUsuarios() error  {
	db , err:= Conexion.ConexionBDPostgres()

	//Comprobando que no hay error
	if err != nil {
		return err
	}
	//Creando Tablas
	db.SingularTable(true)
	if !db.HasTable(&UsuariosModel.Perfil{}){
		if err := db.CreateTable(&UsuariosModel.Perfil{}).Error; err != nil{
			return  err
		}
	}

	if !db.HasTable(&UsuariosModel.Usuario{}){
		if err := db.CreateTable(&UsuariosModel.Usuario{}).AddForeignKey("id_persona",
			"persona(id_persona)","RESTRICT","CASCADE").
			AddForeignKey("id_perfil","perfil(id_perfil)","RESTRICT",
				"CASCADE").Error; err != nil{
			return  err
		}
	}
	return nil
}
