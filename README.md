# Learning Go - Dealing with JSON and XML in Gin-Gonic

Just a POC of an API to understand how Go & gin-gonic works with JSON and XML, binding to structs and grouping routes for an API.

To lift the server up: `go run server.go` 

## Routes:

### POST `/jsonToXML`

Receives any JSON and return same JSON as XML in a `<MAP></MAP>` tag.

### POST `/addPerson/...`

Receives and "person" Object and bind it to a struct

```go
type Person struct {
	Name       string `json:"name" xml:"name" binding:"required"`
	MiddleName string `json:"middleName" xml:"middleName" `
	LastName   string `json:"lastName" xml:"lastName" binding:"required"`
	Age        int    `json:"age" xml:"age" binding:"required"`
}
```
You can pass the object as json, xml, or json and convert it to xml using: 
`addPerson/json` , `addPerson/xml` or `addPerson/jsontoxml `