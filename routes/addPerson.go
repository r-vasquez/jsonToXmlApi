package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string `json:"name" xml:"name" binding:"required"`
	MiddleName string `json:"middleName" xml:"middleName" `
	LastName   string `json:"lastName" xml:"lastName" binding:"required"`
	Age        int    `json:"age" xml:"age" binding:"required"`
}

// AddPersonJson
// bind the JSON from the request Body to a pre-defined person struct.
func AddPersonJson(c *gin.Context) {
	var person Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(person)

	c.JSON(http.StatusOK, gin.H{
		"message": "Person added",
		"person": gin.H{
			"name":       person.Name,
			"middleName": person.MiddleName,
			"lastname":   person.LastName,
			"age":        person.Age,
		},
	})
}

// AddPersonXML
// bind the XML from the request Body to a pre-defined person struct
func AddPersonXML(c *gin.Context) {
	var personXml Person

	if err := c.ShouldBindXML(&personXml); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.XML(http.StatusOK, personXml)
}

// AddPersonJsonToXml
// just passing from Json to XML using standard marshalling features when it has a
// pre defined struct.
func AddPersonJsonToXml(c *gin.Context) {
	var personResponse Person

	// Get bytes from request body
	bytes, _ := ioutil.ReadAll(c.Request.Body)

	// Unmarshal and bind to a Person struct
	if e := json.Unmarshal(bytes, &personResponse); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}

	// With a given struct, gin allow to parse it to XML or JSON.
	c.XML(http.StatusOK, personResponse)
}
