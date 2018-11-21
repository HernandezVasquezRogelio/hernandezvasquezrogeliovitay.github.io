package DireccionesModel

import "time"

type Direcciones struct {
	IdDirecciones  string `gorm:"PRIMARY_KEY"`
	Calle          string
	Colonia        string
	Municipio      string
	NumeroInterior string
	NumeroExterior string
	CodigoPostal   int
	Estado         bool
	FechaRegistro  time.Time
}
