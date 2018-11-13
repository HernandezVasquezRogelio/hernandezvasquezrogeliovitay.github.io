package PreRegistroControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/PreRegistro"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func AgregarUnPreRegistro(c *gin.Context) {
	var preRegistro PreRegistro.PreRegistro
	err := c.BindJSON(&preRegistro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear estructura", "Data": err.Error()})
		return
	}
	if preRegistro.Nombre == "" || preRegistro.Apellido == "" || preRegistro.CorreoElectronico == "" ||
		preRegistro.Telefono == "" || preRegistro.Contrasena == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Ningun campo debe estar vacio", "Data": nil})
		return
	}
	preRegistro.FechaRegistro = time.Now().UTC()
	preRegistro.Id = bson.NewObjectId().Hex()
	preRegistro.Estado = true
    err= preRegistro.InsertaPreRegistro()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"Mensaje":"Error al insertar registro","Data":err.Error()})
		return
	}
    c.JSON(http.StatusOK ,gin.H{"Mensaje":"Datos agregados correctamnete","Data":preRegistro})
	return
}

func GetAllPreRegistro(c *gin.Context) {
	var preRegistro PreRegistro.PreRegistro
	err := c.BindJSON(&preRegistro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear estructura", "Data": err.Error()})
		return
	}
	if preRegistro.Nombre == "" || preRegistro.Apellido == "" || preRegistro.CorreoElectronico == "" ||
		preRegistro.Telefono == "" || preRegistro.Contrasena == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Ningun campo debe estar vacio", "Data": nil})
		return
	}
	preRegistro.FechaRegistro = time.Now().UTC()
	preRegistro.Id = bson.NewObjectId().Hex()
	preRegistro.Estado = true
	err= preRegistro.InsertaPreRegistro()
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"Mensaje":"Error al insertar registro","Data":err.Error()})
		return
	}
	c.JSON(http.StatusOK ,gin.H{"Mensaje":"Datos agregados correctamnete","Data":preRegistro})
	return
}
