package PreRegistroControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/PreRegistro"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func GetAllPreRegistro(c *gin.Context) {
	err, resultado := PreRegistro.ConsultarAllPreRegistro()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//********************************************************************************************

func GetPreRegistro(c *gin.Context) {
	id := c.Param("id")
	err, resultado := PreRegistro.ConsultaPreRegistro(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func PostPreRegistro(c *gin.Context) {
	var preRegistro PreRegistro.PreRegistro
	err := c.BindJSON(&preRegistro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear estructura", "Data": err.Error()})
		return
	}
	if preRegistro.Nombre == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "El campo nombre no debe estar vacio", "Data": nil})
		return
	}
	if preRegistro.Apellido == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "El campo Apellido no debe estar vacio", "Data": nil})
		return
	}
	if preRegistro.CorreoElectronico == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "El campo Correo no debe estar vacio", "Data": nil})
		return
	}
	if preRegistro.Telefono == "" || preRegistro.Contrasena == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "El campo Telefono no debe estar vacio", "Data": nil})
		return
	}
	if preRegistro.Contrasena == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "El campo Conraseña no debe estar vacio", "Data": nil})
		return
	}
	preRegistro.FechaRegistro = time.Now().UTC()
	preRegistro.Id = bson.NewObjectId().Hex()
	preRegistro.Estado = true
	err = preRegistro.InsertaPreRegistro()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al insertar registro", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": preRegistro})
	return
}
func UpdatePreRegistro(c *gin.Context) {
	id := c.Param("id")
	var preregistro PreRegistro.PreRegistro
	err := c.BindJSON(&preregistro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar la estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo id esta vacio", "Data": nil})
	}
	if preregistro.Nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo nombre esta vacio", "Data": nil})
	}
	if preregistro.Apellido == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo apellido esta vacio", "Data": nil})
	}
	if preregistro.CorreoElectronico == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo correo esta vacio", "Data": nil})
	}
	if preregistro.Telefono == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo telefono esta vacio", "Data": nil})
	}
	if preregistro.Contrasena == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo contraseña esta vacio", "Data": nil})
	}
	if preregistro.CodigoInvitacion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo codigo esta vacio", "Data": nil})
	}
	err = preregistro.ActualizarPreRegistro(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al actualizar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos actualizados correctamente", "Data": preregistro})
	return
}

//********************************************************************************************
func DeletePreRegistro(c *gin.Context) {
	id := c.Param("id")
	var preregistro PreRegistro.PreRegistro
	err := c.BindJSON(&preregistro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar estructura", "Data": err.Error()})
		return
	}
	if id=="" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo id esta vacio", "Data": nil})
		return
	}

	err = preregistro.EliminarPreRegistro(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al eliminar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Preregistro eliminado correctamnete", "Data": preregistro})
	return
}
