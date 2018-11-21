package PreRegistro

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rogelio/Agentes/Conexion"
)

//********************************************************************************************************
//-----------------------Consultar todos los preregistros

func ConsultarAllPreRegistro() (error, *[] PreRegistro) {
	var datos [] PreRegistro
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
//-----------------------Insertar un preregistros

func (p *PreRegistro) InsertaPreRegistro() error {

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
//-----------------------Consultar un preregistro

func ConsultaPreRegistro(ID string) (error, *PreRegistro) {
	var item PreRegistro
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	res := db.
		Where("id = ?", ID).
		Find(&item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}
		return res.Error, &item
	}
	return nil, &item
}

//********************************************************************************************************
//-----------------------Actualizar un preregistro

func (a PreRegistro) ActualizarPreRegistro(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var preregistro PreRegistro
	tx := db.Begin()
	err = tx.Where("id = ?", id).First(&preregistro).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el preregistro")
		}
		return err
	}
	err = db.Model(&a).Update(PreRegistro{Nombre: a.Nombre, Apellido: a.Apellido, CorreoElectronico:
	a.CorreoElectronico, Telefono: a.Telefono, Contrasena: a.Contrasena, CodigoInvitacion: a.CodigoInvitacion}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
//-----------------------Eliminar un preregistro

func (a PreRegistro) EliminarPreRegistro(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var preregistro PreRegistro
	tx := db.Begin()
	err = tx.Where("id = ?", id).First(&preregistro).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el preregistro")
		}
		return err
	}
	if !preregistro.Estado {
		tx.Rollback()
		return errors.New("Este preregistro ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Where("estado = ?", true).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
