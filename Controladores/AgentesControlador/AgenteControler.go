package AgentesControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/AgenteModel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func GetAllAgentes(c *gin.Context) {
	err, resultado := AgenteModel.ConsultarAllAgentes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//********************************************************************************************

func GetAgente(c *gin.Context) {
	id := c.Param("id")
	err, resultado := AgenteModel.ConsultaAgente(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//********************************************************************************************
func PostAgente(c *gin.Context) {
	var agentes AgenteModel.Agente
	err := c.BindJSON(&agentes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear estructura", "Data": err.Error()})
		return
	}
	if agentes.IdPersona == "" || agentes.CodigoInvitacion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Ningun campo debe estar vacio", "Data": nil})
		return
	}
	agentes.FechaRegistro = time.Now().UTC()
	agentes.IdAgente = bson.NewObjectId().Hex()
	agentes.Estado = true
	err = agentes.AgregarAgente()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al insertar registro", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": agentes})
	return
}

//********************************************************************************************
func UpdateAgente(c *gin.Context) {
	id := c.Param("id")
	var agentes AgenteModel.Agente
	err := c.BindJSON(&agentes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar solicitud", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo id esta vacio", "Data": nil})
	}
	err = agentes.ActualizarAgentes(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al actualizar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": agentes})
	return
}

//********************************************************************************************
func DeleteAgente(c *gin.Context) {
	id := c.Param("id")
	var agentes AgenteModel.Agente
	err := c.BindJSON(&agentes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar solicitud", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo id esta vacio", "Data": nil})
	}
	err = agentes.EliminarAgentes(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al eliminar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Agente eliminado correctamnete", "Data": agentes})
	return
}
