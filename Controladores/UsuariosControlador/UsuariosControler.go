package UsuariosControlador

import (
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Modelos/UsuariosModel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func GetAllUsuarios(c *gin.Context) {
	err, resultado := UsuariosModel.ConsultarAllUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//********************************************************************************************

func GetUsuario(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Debes ingresar el id para buscar", "Data": nil})
		return
	}
	err, resultado := UsuariosModel.ConsultaUsuario(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al consultar los datos", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos consultados correctamente", "Data": resultado})
	return
}

//*****************************************************************************************************************
func PostUsuario(c *gin.Context) {
	var usuarios UsuariosModel.Usuario
	err := c.BindJSON(&usuarios)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al castear tu estructura", "Data": err.Error()})
		return
	}
	if usuarios.IdPersona == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el idPersona", "Data": nil})
		return
	}
	if usuarios.IdPerfil == "" {

		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Falta el idPerfil", "Data": nil})
		return
	}
	usuarios.FechaRegistro = time.Now().UTC()
	usuarios.IdUsuario = bson.NewObjectId().Hex()
	usuarios.Estado = true
	err = usuarios.InsertaUsuario()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al insertar registro", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos agregados correctamnete", "Data": usuarios})
	return
}

func UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuarios UsuariosModel.Usuario
	err := c.BindJSON(&usuarios)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar la estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo id esta vacio", "Data": nil})
		return
	}
	if usuarios.IdPersona == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo idPersona esta vacio", "Data": nil})
		return
	}
	if usuarios.IdPerfil == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo idPerfil esta vacio", "Data": nil})
		return
	}
	err = usuarios.ActualizarUsuario(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al actualizar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Datos actualizados correctamente", "Data": usuarios})
	return
}

//********************************************************************************************
func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuarios UsuariosModel.Usuario
	err := c.BindJSON(&usuarios)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error al decodificar estructura", "Data": err.Error()})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Mensaje": "Error el campo idUsuario esta vacio", "Data": nil})
		return
	}

	err = usuarios.EliminarUsuario(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al eliminar", "Data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Usuario eliminado correctamnete", "Data": usuarios})
	return
}
