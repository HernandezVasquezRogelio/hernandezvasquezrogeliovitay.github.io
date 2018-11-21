package PersonaModel

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rogelio/Agentes/Conexion"
)

//********************************************************************************************************
//-----------------------Consultar todos las personas

func ConsultarAllPersonas() (error, *[] Persona) {
	var datos [] Persona
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	err = db.Where("Estado=?", true).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
//-----------------------Consultar una persona

func ConsultaPersona(ID string) (error, *Persona) {
	var item Persona
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	res := db.Where("id_persona = ?", ID).Find(&item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errors.New("No se encontro esta persona"), nil
		}
		return res.Error, &item
	}
	return nil, &item
}

//********************************************************************************************************
//-----------------------Insertar una persona

func (p *Persona) InsertaPersona() error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.Begin()
	err = tx.Create(p).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

//********************************************************************************************************
//-----------------------Actualizar una persona

func (a Persona) ActualizarPersona(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var personas Persona
	tx := db.Begin()
	err = tx.Where("id_persona = ?", id).First(&personas).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la persona")
		}
		return err
	}
	err = tx.Model(&a).Update(Persona{Nombre: a.Nombre, ApellidoPat: a.ApellidoPat, ApellidoMat: a.ApellidoMat,
		FechaNacimiento: a.FechaNacimiento, Sexo: a.Sexo, Email: a.Email, Telefono: a.Telefono}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
//-----------------------Eliminar una persona

func (a Persona) EliminarPersona(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var personas Persona
	tx := db.Begin()
	err = tx.Where("id_persona = ?", id).First(&personas).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la persona")
		}
		return err
	}
	if !personas.Estado {
		tx.Rollback()
		return errors.New("Esta persona ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Where("estado = ?", true).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
