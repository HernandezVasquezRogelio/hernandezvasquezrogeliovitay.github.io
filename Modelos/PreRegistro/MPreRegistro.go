package PreRegistro

import "github.com/rogelio/Agentes/Conexion"

func (p *PreRegistro) InsertaPreRegistro() error {

	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Create(p).Error
	if err != nil {
		return err
	}
	return nil

}

func (p *PreRegistro) ConsultarAllPreRegistro() error {

	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return err
	}
	defer db.Close()
	err=db.Where("Estado=?",true).Find(p).Error
	if err != nil {
		return err
	}
	return nil

}