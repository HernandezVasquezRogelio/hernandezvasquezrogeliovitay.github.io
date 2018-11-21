package PersonaControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/PersonaModel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

//*****************************************************************************************************************
func GetAllPersonas(c *gin.Context) {
	err, resultado := PersonaModel.ConsultarAllPersonas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func GetPersona(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Debes ingresar el id para buscar", "Data": nil})
		return
	}
	err, resultado := PersonaModel.ConsultaPersona(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func PostPersona(c *gin.Context) {
	var personas PersonaModel.Persona
	err := c.BindJSON(&personas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear tu estructura", "Data": err.Error()})
		return
	}
	if personas.Nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo nombre", "Data": nil})
		return
	}
	if personas.ApellidoPat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo apellido paterno", "Data": nil})
		return
	}
	if personas.ApellidoMat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo apellido materno", "Data": nil})
		return
	}
	if personas.FechaNacimiento.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo fecha de nacimiento ", "Data": nil})
		return
	}

	if personas.Sexo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo sexo", "Data": nil})
		return
	}

	if personas.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo email", "Data": nil})
		return
	}

	if personas.Telefono == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo telefono", "Data": nil})
		return
	}
	personas.FechaRegistro = time.Now().UTC()
	personas.IdPersona = bson.NewObjectId().Hex()
	personas.Estado = true
	err = personas.InsertaPersona()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al insertar registro", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": personas})
	return
}

//*****************************************************************************************************************
func UpdatePersona(c *gin.Context) {
	id := c.Param("id")
	var personas PersonaModel.Persona
	err := c.BindJSON(&personas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar la estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo id", "Data": nil})
		return
	}
	if personas.Nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo nombre", "Data": nil})
		return
	}
	if personas.ApellidoPat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo apellido paterno", "Data": nil})
		return
	}
	if personas.ApellidoMat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo apellido materno", "Data": nil})
		return
	}
	if personas.FechaNacimiento.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo fecha de nacimiento ", "Data": nil})
		return
	}

	if personas.Sexo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo sexo", "Data": nil})
		return
	}

	if personas.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo email", "Data": nil})
		return
	}

	if personas.Telefono == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el campo telefono", "Data": nil})
		return
	}
	err = personas.ActualizarPersona(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al actualizar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos actualizados correctamente", "Data": personas})
	return
}

//********************************************************************************************
func DeletePersona(c *gin.Context) {
	id := c.Param("id")
	var personas PersonaModel.Persona
	err := c.BindJSON(&personas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error se necesita el id", "Data": nil})
		return
	}

	err = personas.EliminarPersona(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al eliminar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "persona eliminada correctamente", "Data": personas})
	return
}
