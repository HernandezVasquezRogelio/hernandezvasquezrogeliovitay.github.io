package DireccionesControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/DireccionesModel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

//*****************************************************************************************************************
func GetAllDirecciones(c *gin.Context) {
	err, resultado := DireccionesModel.ConsultarAllDirecciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func GetDireccion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Debes ingresar el id para buscar", "Data": nil})
		return
	}
	err, resultado := DireccionesModel.ConsultaDireccion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func PostDireccion(c *gin.Context) {
	var direcciones DireccionesModel.Direcciones
	err := c.BindJSON(&direcciones)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear tu estructura", "Data": err.Error()})
		return
	}
	if direcciones.Calle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Calle", "Data": nil})
		return
	}
	if direcciones.Colonia == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Colonia", "Data": nil})
		return
	}
	if direcciones.Municipio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Municipios", "Data": nil})
		return
	}
	if direcciones.NumeroInterior == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Numero interior ", "Data": nil})
		return
	}

	if direcciones.NumeroExterior == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Numero exterior", "Data": nil})
		return
	}

	if direcciones.CodigoPostal == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Codigo postal", "Data": nil})
		return
	}
	direcciones.FechaRegistro = time.Now().UTC()
	direcciones.IdDirecciones = bson.NewObjectId().Hex()
	direcciones.Estado = true
	err = direcciones.InsertaDireccion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al insertar registro", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": direcciones})
	return
}

//*****************************************************************************************************************
func UpdateDirecciones(c *gin.Context) {
	id := c.Param("id")
	var direcciones DireccionesModel.Direcciones
	err := c.BindJSON(&direcciones)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar la estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo id", "Data": nil})
		return
	}
	if direcciones.Calle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Calle", "Data": nil})
		return
	}
	if direcciones.Colonia == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Colonia", "Data": nil})
		return
	}
	if direcciones.Municipio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Municipios", "Data": nil})
		return
	}
	if direcciones.NumeroInterior == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Numero interior ", "Data": nil})
		return
	}

	if direcciones.NumeroExterior == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Numero exterior", "Data": nil})
		return
	}

	if direcciones.CodigoPostal == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo Codigo postal", "Data": nil})
		return
	}
	err = direcciones.ActualizarDireccion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al actualizar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos actualizados correctamente", "Data": direcciones})
	return
}

//********************************************************************************************
func DeleteDirecciones(c *gin.Context) {
	id := c.Param("id")
	var direcciones DireccionesModel.Direcciones
	err := c.BindJSON(&direcciones)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error se necesita el id", "Data": nil})
		return
	}

	err = direcciones.EliminarDireccion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al eliminar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "persona eliminada correctamente", "Data": direcciones})
	return
}
