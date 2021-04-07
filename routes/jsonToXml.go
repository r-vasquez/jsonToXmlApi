package routes

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r-vasquez/jsonToXmlApi/codec"
)

// JsonToXmlRoute
// Convert any given JSON to a basic XML Response
func JsonToXmlRoute(c *gin.Context) {
	var jsonMap codec.Map
	var xmlResponse codec.Map

	// Get the bytes from the body
	incBytes, e := ioutil.ReadAll(c.Request.Body)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}

	// Unmarshal Incoming JSON to a Map
	if e := json.Unmarshal(incBytes, &jsonMap); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}

	// Marshal Map to a Valid XML Encoding
	xmlBytes, e := xml.MarshalIndent(jsonMap, "", " ")

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}

	// Finally Unmarshal the XML Bytes to a XML Map[string]string for the response
	if e := xml.Unmarshal(xmlBytes, &xmlResponse); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}

	c.XML(http.StatusOK, xmlResponse)
}
