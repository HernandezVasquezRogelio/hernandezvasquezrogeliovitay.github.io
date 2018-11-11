package AgenteModel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rogelio/Agentes/Conexion"
)

//Consultar todos los agentes
func GetAllAgentes(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var agentes []Agente
	db.Find(&agentes)
	c.JSON(200, agentes)
}

//Agregar un agente
func PostAgentes(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var agentes Agente
	c.Bind(&agentes)
	if agentes.IdAgente != "" && agentes.IdPersona != "" && agentes.CodigoInvitacion != "" {
		db.Create(&agentes)
		c.JSON(201, gin.H{"Exito": agentes})
	} else {
		c.JSON(422, gin.H{"error": "Algun dato hace falta"})
	}
}
//actualizar datos del agente
func UpdateAgentes(c *gin.Context) {
	var newAgente Agente
	c.BindJSON(&newAgente)
	fmt.Println(newAgente)
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	idAgente := c.Param("id")
	var agentes Agente
	db.First(&agentes, idAgente)
	fmt.Println("Estos estan en la tabla", agentes)
	if newAgente.CodigoInvitacion != "" {
		if agentes.IdAgente != "" {

			result := Agente{
				IdAgente:         agentes.IdAgente,
				IdPersona:        agentes.IdPersona,
				CodigoInvitacion: newAgente.CodigoInvitacion,
				Estado:           agentes.Estado,
				FechaRegistro:    agentes.FechaRegistro,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"Actualizacion": result})
		} else {
			c.JSON(404, gin.H{"error": "Agente no encontrada"})
		}
	} else {
		c.JSON(422, gin.H{"Error": "Falta algun dato"})
	}

}
//inhabilitar agente
func EliminarAgente(c *gin.Context) {
	var newAgente Agente
	c.BindJSON(&newAgente)
	fmt.Println(newAgente)
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	idAgente := c.Param("id")
	var agentes Agente
	db.First(&agentes, idAgente)
	fmt.Println("Estos estan en la tabla", agentes)
	if newAgente.Estado != true {
		if agentes.IdAgente != "" {

			result := Agente{
				IdAgente:         agentes.IdAgente,
				IdPersona:        agentes.IdPersona,
				CodigoInvitacion: agentes.CodigoInvitacion,
				Estado:           newAgente.Estado,
				FechaRegistro:    agentes.FechaRegistro,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"Inhabilitado": result})
		} else {
			c.JSON(404, gin.H{"error": "Agente no encontrada"})
		}
	} else {
		c.JSON(422, gin.H{"Error": "Falta algun dato"})
	}

}
//Consultar todos las personas
func GetAllPersona(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var personas []Persona
	db.Find(&personas)
	c.JSON(200, personas)
}

//Agregar una persona
func PostPersona(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var personas Persona
	c.Bind(&personas)
	if personas.IdPersona != "" && personas.ApellidoMat != "" && personas.ApellidoPat != "" {
		db.Create(&personas)
		c.JSON(201, gin.H{"Exito": personas})
	} else {
		c.JSON(422, gin.H{"error": "Algun dato hace falta"})
	}

}
func UpdatePersona(c *gin.Context) {
	var newPersonas Persona
	c.BindJSON(&newPersonas)
	fmt.Println(newPersonas)
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	idPersona := c.Param("id")
	var personas Persona
	db.First(&personas, idPersona)
	fmt.Println(personas)
	if newPersonas.Nombre != "" && newPersonas.ApellidoPat != "" && newPersonas.ApellidoMat != "" &&
		newPersonas.Email != "" {
		if personas.IdPersona != "" {

			result := Persona{
				IdPersona:       personas.IdPersona,
				Nombre:          newPersonas.Nombre,
				ApellidoPat:     newPersonas.ApellidoPat,
				ApellidoMat:     newPersonas.ApellidoMat,
				FechaNacimiento: personas.FechaNacimiento,
				Sexo:            personas.Sexo,
				Email:           newPersonas.Email,
				Telefono:        personas.Telefono,
				Estado:          personas.Estado,
				FechaRegistro:   personas.FechaRegistro,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"Actualizacion": result})
		} else {
			c.JSON(404, gin.H{"error": "Persona no encontrada"})
		}
	} else {
		c.JSON(422, gin.H{"Error": "Falta algun dato"})
	}

}

//consultar todas las direcciones
func GetAllDirecciones(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var direcciones []CrudDirecciones
	db.Find(&direcciones)
	c.JSON(200, direcciones)
}

//Agregar una direccion
func PostDirecciones(c *gin.Context) {
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	var direcciones CrudDirecciones
	c.Bind(&direcciones)
	if direcciones.IdDirecciones != "" && direcciones.Calle != "" && direcciones.CodigoPostal != 0 &&
		direcciones.Colonia != "" {
		db.Create(&direcciones)
		c.JSON(201, gin.H{"Correcto": direcciones})
	} else {
		c.JSON(422, gin.H{"error": "Algun dato hace falta"})
	}

}

//Actualizar direcciones
func UpdateDirecciones(c *gin.Context) {
	var newDirecciones CrudDirecciones
	c.BindJSON(&newDirecciones)
	fmt.Println(newDirecciones)
	db, err := Conexion.ConexionBDPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	idDireccion := c.Param("id")
	var direcciones CrudDirecciones
	db.First(&direcciones, idDireccion)
	fmt.Println(direcciones)
	if newDirecciones.Colonia != "" && newDirecciones.Municipio != "" && newDirecciones.CodigoPostal != 0 &&
		newDirecciones.Calle != "" && newDirecciones.NumeroInterior != "" && newDirecciones.NumeroExterior != "" {
		if direcciones.IdDirecciones != "" {

			result := CrudDirecciones{
				IdDirecciones:       direcciones.IdDirecciones,
				Calle: 				 newDirecciones.Calle,
				Colonia: 			 newDirecciones.Colonia,
				Municipio: 			 newDirecciones.Municipio,
				NumeroInterior: 	 newDirecciones.NumeroInterior,
				NumeroExterior: 	 newDirecciones.NumeroExterior,
				CodigoPostal: 	     newDirecciones.CodigoPostal,
				Estado: 		     direcciones.Estado,
				FechaRegistro: 		 direcciones.FechaRegistro,

			}
			db.Save(&result)
			c.JSON(200, gin.H{"Actualizado": result})
		} else {
			c.JSON(404, gin.H{"error": "Direccion no encontrada"})
		}
	} else {
		c.JSON(422, gin.H{"Error": "Falta algun dato"})
	}

}
