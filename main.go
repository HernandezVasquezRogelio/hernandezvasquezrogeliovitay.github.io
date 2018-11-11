package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rogelio/Agentes/Init"
	"github.com/rogelio/Agentes/Modelos/AgenteModel"
)

func main() {

	Init.InitAgentes()
	Init.InitUsuarios()
	r := gin.Default()
	//EndPoint Agentes
	agente := r.Group("agente")
	{
		agente.GET("/persona", AgenteModel.GetAllPersona)
		agente.POST("/persona", AgenteModel.PostPersona)
		agente.PUT("/persona/:id", AgenteModel.UpdatePersona)
		agente.GET("", AgenteModel.GetAllAgentes)
		agente.POST("", AgenteModel.PostAgentes)
		agente.PUT("/update/:id", AgenteModel.UpdateAgentes)
		agente.PUT("/delete/:id", AgenteModel.EliminarAgente)
		agente.GET("/direccion", AgenteModel.GetAllDirecciones)
		agente.POST("/direccion", AgenteModel.PostDirecciones)
		agente.PUT("/direccion/:id", AgenteModel.UpdateDirecciones)
	}
	//EndPoint Usuarios
	/*usuarios := r.Group("usuarios")
	{
		usuarios.GET("")
	}*/
	r.Run() // listen and serve on 0.0.0.0:8080
}
