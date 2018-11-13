package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rogelio/Agentes/Controladores/PreRegistroControlador"
	"github.com/rogelio/Agentes/Init"
	"github.com/rogelio/Agentes/Modelos/AgenteModel"
)

func main() {

	Init.InitAgentes()
	Init.InitUsuarios()
	Init.InitPreRegistros()
	r := gin.Default()
	//EndPoint Agentes
	agente := r.Group("agente")
	{
		agente.GET("/persona", AgenteModel.GetAllPersona)
		agente.POST("/persona", AgenteModel.PostPersona)
		agente.PUT("/persona/update/:id", AgenteModel.UpdatePersona)
		agente.PUT("/persona/delete/:id", AgenteModel.EliminarPersona)

		agente.GET("", AgenteModel.GetAllAgentes)
		agente.POST("", AgenteModel.PostAgentes)
		agente.PUT("/update/:id", AgenteModel.UpdateAgentes)
		agente.PUT("/delete/:id", AgenteModel.EliminarAgente)

		agente.GET("/direccion", AgenteModel.GetAllDirecciones)
		agente.POST("/direccion", AgenteModel.PostDirecciones)
		agente.PUT("/direccion/update/:id", AgenteModel.UpdateDirecciones)
		agente.PUT("/direccion/delete/:id", AgenteModel.EliminarDirecciones)

	}
	//EndPoint PreRegistro
	usuarios := r.Group("preregistro")
	{
		usuarios.GET("")
		usuarios.POST("", PreRegistroControlador.AgregarUnPreRegistro)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
