package routes

import "github.com/gin-gonic/gin"

// InitializeRoutes given a certain gin route.
func InitializeRoutes(r *gin.Engine) {

	r.POST("/jsontoxml", JsonToXmlRoute)

	addPerson := r.Group("/addPerson")
	{
		addPerson.POST("/json", AddPersonJson)
		addPerson.POST("/xml", AddPersonXML)
		addPerson.POST("/jsontoxml", AddPersonJsonToXml)
	}
}
