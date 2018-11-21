package UsuariosModel

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rogelio/Agentes/Conexion"
)

//********************************************************************************************************
//-----------------------Consultar todos los usuarios

func ConsultarAllUsuarios() (error, *[] Usuario) {
	var datos [] Usuario
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
//-----------------------Consultar un usuario

func ConsultaUsuario(ID string) (error, *Usuario) {
	var item Usuario
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	res := db.Where("id_usuario = ?", ID).Find(&item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errors.New("No se encontro este usuario"), nil
		}
		return res.Error, &item
	}
	return nil, &item
}

//********************************************************************************************************
//-----------------------Insertar un usuario

func (p *Usuario) InsertaUsuario() error {
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
//-----------------------Actualizar un Usuario

func (a Usuario) ActualizarUsuario(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var usuarios Usuario
	tx := db.Begin()
	err = tx.Where("id_usuario = ?", id).First(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el usuario")
		}
		return err
	}
	err = tx.Model(&a).Update(Usuario{IdPersona: a.IdPersona, IdPerfil: a.IdPerfil}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
//-----------------------Eliminar un preregistro

func (a Usuario) EliminarUsuario( id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var usuarios Usuario
	tx := db.Begin()
	err = tx.Where("id = ?", id).First(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el usuario")
		}
		return err
	}
	if !usuarios.Estado {
		tx.Rollback()
		return errors.New("Este usuario ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Where("estado = ?", true).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
