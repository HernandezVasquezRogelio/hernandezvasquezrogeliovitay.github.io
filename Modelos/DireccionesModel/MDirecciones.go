package DireccionesModel

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rogelio/Agentes/Conexion"
)

//********************************************************************************************************
//-----------------------Consultar todas las direcciones

func ConsultarAllDirecciones() (error, *[] Direcciones) {
	var datos [] Direcciones
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
//-----------------------Consultar una direccion

func ConsultaDireccion(ID string) (error, *Direcciones) {
	var item Direcciones
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err, nil
	}
	defer db.Close()
	res := db.Where("id_direcciones = ?", ID).Find(&item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return errors.New("No se encontro esta direccion"), nil
		}
		return res.Error, &item
	}
	return nil, &item
}

//********************************************************************************************************
//-----------------------Insertar una direccion

func (p *Direcciones) InsertaDireccion() error {
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
//-----------------------Actualizar una direccion

func (a Direcciones) ActualizarDireccion(id string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var direcciones Direcciones
	tx := db.Begin()
	err = tx.Where("id_direcciones = ?", id).First(&direcciones).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la direccion")
		}
		return err
	}
	err = tx.Model(&a).Update(Direcciones{Calle: a.Calle, Colonia: a.Colonia, Municipio: a.Municipio,
		NumeroInterior: a.NumeroInterior, NumeroExterior: a.NumeroExterior, CodigoPostal: a.CodigoPostal}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
//-----------------------Eliminar una direccion

func (a Direcciones) EliminarDireccion(ID string) error {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	var direcciones Direcciones
	tx := db.Begin()
	err = tx.Where("id_direcciones = ?", ID).First(&direcciones).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la direccion")
		}
		return err
	}
	if !direcciones.Estado {
		tx.Rollback()
		return errors.New("Esta direccion ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Where("estado = ?", true).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
