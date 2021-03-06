package AgenteModel

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rogelio/Agentes/Conexion"
)

//********************************************************************************************************

//-----------------------Consultar todos los agentes
func ConsultarAllAgentes() (error, *[] Agente) {
	var datos [] Agente
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
//-----------------------Consultar un agente

func ConsultaAgente(ID string) (error, *Agente) {
	var item Agente
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	res := db.
		Where("id_agente = ?", ID).
		Find(&item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errors.New("No se encontro el agente"), nil
		}
		return res.Error, &item
	}
	return nil, &item
}

//********************************************************************************************************

//-----------------------Agregar un agente
func (a *Agente) AgregarAgente() error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.Begin()
	err = tx.Create(a).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************

//----------------------actualizar datos del agente
func (a Agente) ActualizarAgentes(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var agentes Agente
	tx := db.Begin()
	err = tx.Where("id_agente = ?", id).First(&agentes).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el agente")
		}
		return err

	}
	err = tx.Model(&a).Update(Agente{IdPersona: a.IdPersona, CodigoInvitacion: a.CodigoInvitacion}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************

//--------------------------inhabilitar un agente
func (a Agente) EliminarAgentes(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var agentes Agente
	tx := db.Begin()
	err = tx.Where("id_agente = ?", id).First(&agentes).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Este Agente no se encontro")
		}
		return err

	}
	if !agentes.Estado {
		tx.Rollback()
		return errors.New("Este agente ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}