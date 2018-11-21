package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rogelio/Agentes/Controladores/AgentesControlador"
	"github.com/rogelio/Agentes/Controladores/DireccionesControlador"
	"github.com/rogelio/Agentes/Controladores/PersonaControlador"
	"github.com/rogelio/Agentes/Controladores/PreRegistroControlador"
	"github.com/rogelio/Agentes/Controladores/UsuariosControlador"
	"github.com/rogelio/Agentes/Init"
)

func main() {

	Init.InitAgentes()
	Init.InitUsuarios()
	Init.InitPreRegistros()
	r := gin.Default()
	//EndPoint Agentes
	agente := r.Group("agente")
	{
		agente.GET("", AgentesControlador.GetAllAgentes)
		agente.GET(":id", AgentesControlador.GetAgente)
		agente.POST("", AgentesControlador.PostAgente)
		agente.PUT(":id", AgentesControlador.UpdateAgente)
		agente.DELETE(":id", AgentesControlador.DeleteAgente)

	}
	//EndPoint PreRegistro
	preregistro := r.Group("preregistro")
	{
		preregistro.GET("", PreRegistroControlador.GetAllPreRegistro)
		preregistro.GET(":id", PreRegistroControlador.GetPreRegistro)
		preregistro.POST("", PreRegistroControlador.PostPreRegistro)
		preregistro.PUT(":id", PreRegistroControlador.UpdatePreRegistro)
		preregistro.DELETE(":id", PreRegistroControlador.DeletePreRegistro)
	}
	//EndPoint Usuarios
	usuario := r.Group("usuario")
	{
		usuario.GET("", UsuariosControlador.GetAllUsuarios)
		usuario.GET(":id", UsuariosControlador.GetUsuario)
		usuario.POST("", UsuariosControlador.PostUsuario)
		usuario.PUT(":id", UsuariosControlador.UpdateUsuario)
		usuario.DELETE(":id", UsuariosControlador.DeleteUsuario)
	}
	//EndPoint Persona
	persona := r.Group("persona")
	{
		persona.GET("", PersonaControlador.GetAllPersonas)
		persona.GET(":id", PersonaControlador.GetPersona)
		persona.POST("", PersonaControlador.PostPersona)
		persona.PUT(":id", PersonaControlador.UpdatePersona)
		persona.DELETE(":id", PersonaControlador.DeletePersona)
	}
	//EndPoint Direcciones
	direcciones := r.Group("direccion")
	{
		direcciones.GET("", DireccionesControlador.GetAllDirecciones)
		direcciones.GET(":id", DireccionesControlador.GetDireccion)
		direcciones.POST("", DireccionesControlador.PostDireccion)
		direcciones.PUT(":id", DireccionesControlador.UpdateDirecciones)
		direcciones.DELETE(":id", DireccionesControlador.DeleteDirecciones)

	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

/*
agente.GET("/direccion", AgenteModel.GetAllDirecciones)
agente.POST("/direccion", AgenteModel.PostDirecciones)
agente.PUT("/direccion/update/:id", AgenteModel.UpdateDirecciones)
agente.PUT("/direccion/delete/:id", AgenteModel.EliminarDirecciones)*/
