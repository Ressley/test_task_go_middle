package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gprc "github.com/ressley/test_task_go_middle/http-grpc-client/grpc"
	"github.com/ressley/test_task_go_middle/http-grpc-client/http/requests"
	"github.com/ressley/test_task_go_middle/pkg/app"
	"github.com/ressley/test_task_go_middle/pkg/serializers/json"
)

func SendMessage(c *gin.Context) {
	gin := app.Gin{C: c}
	var request requests.PostMessageRequest

	log.Printf("[info] Reading body of message request")
	if has_error := json.FromBody(&request, c.Request.Body); has_error != nil {
		gin.Response(http.StatusBadRequest, nil, has_error)
		return
	}
	log.Printf("[info] Destination: %s", *request.Destination)
	log.Printf("[info] Data: %v", *&request.Data)

	client, err := gprc.NewGrpcClient(*request.Destination)
	if err != nil {
		gin.Response(http.StatusBadRequest, nil, err)
		return
	}
	defer client.Close()

	log.Printf("[info] Calling SendMessage of grpc-client")
	err = client.SendMessage(request.Data)
	if err != nil {
		gin.Response(http.StatusBadRequest, nil, err)
		return
	}

	gin.Response(http.StatusOK, nil, nil)
}
