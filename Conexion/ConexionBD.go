package Conexion

import (
	"github.com/jinzhu/gorm"
)

func ConexionBDPostgres() (*gorm.DB, error) {
	//conexion ala base de datos
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test "+
		"password=rogelioSCC2610 sslmode=disable")

	db.LogMode(true)
	db.SingularTable(true)
	//Error
	if err != nil {
		return nil, err
	}
	return db, err
}

