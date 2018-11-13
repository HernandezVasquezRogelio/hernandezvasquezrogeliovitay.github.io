package Init

import (
	"github.com/rogelio/Agentes/Conexion"
	"github.com/rogelio/Agentes/Modelos/AgenteModel"

)

func InitAgentes() error  {
	db , err:= Conexion.ConexionBDPostgres()

	//Comprobando que no hay error
	if err != nil {
		return err
	}
	//Creando Tablas
	db.SingularTable(true)
	if !db.HasTable(&AgenteModel.Persona{}){
		if err := db.CreateTable(&AgenteModel.Persona{}).Error; err != nil{
			return  err
		}
	}
	if !db.HasTable(&AgenteModel.CrudAcreditaciones{}){
		if err := db.CreateTable(&AgenteModel.CrudAcreditaciones{}).Error; err != nil{
			return  err
		}
	}
	if !db.HasTable(&AgenteModel.CrudDirecciones{}){
		if err := db.CreateTable(&AgenteModel.CrudDirecciones{}).Error; err != nil{
			return  err
		}
	}

	if !db.HasTable(&AgenteModel.Agente{}){
		if err := db.CreateTable(&AgenteModel.Agente{}).AddForeignKey("id_persona",
			"persona(id_persona)","RESTRICT","CASCADE").Error; err != nil{
			return  err
		}
	}
	if !db.HasTable(&AgenteModel.AcreditacionAgentes{}){
		if err := db.CreateTable(&AgenteModel.AcreditacionAgentes{}).AddForeignKey("id_acreditacion",
			"crud_acreditaciones(id_acreditacion)","RESTRICT","CASCADE").
			AddForeignKey("id_agente","agente(id_agente)","RESTRICT",
				"CASCADE").Error; err != nil{
			return  err
		}
	}
	if !db.HasTable(&AgenteModel.DireccionPersona{}){
		if err := db.CreateTable(&AgenteModel.DireccionPersona{}).AddForeignKey("id_direcciones",
			"crud_direcciones(id_direcciones)","RESTRICT","CASCADE").
			AddForeignKey("id_persona","persona(id_persona)","RESTRICT",
				"CASCADE").Error; err != nil{
			return  err
		}
	}


	return nil
}