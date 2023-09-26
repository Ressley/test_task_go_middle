package app

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Error any         `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (g *Gin) Response(httpCode int, data interface{}, err any) {
	response := Response{}

	switch x := err.(type) {
	case string:
		response.Error = &x
		break
	case error:
		if x == nil {
			break
		}
		temp := x.Error()
		response.Error = &temp
		break
	case nil:
		response.Data = data
		break
	default:
		response.Error = x
	}
	g.C.Writer.Header().Set("Content-Type", g.C.Request.Header.Get("Content-Type"))
	g.C.Writer.WriteHeader(httpCode)
	encoder := json.NewEncoder(g.C.Writer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	encoder.Encode(response)
}
