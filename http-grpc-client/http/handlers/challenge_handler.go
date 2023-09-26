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

func SendChallenge(c *gin.Context) {
	gin := app.Gin{C: c}
	var request requests.PostChallengeRequest

	log.Printf("[info] Reading body of challenge request")
	if has_error := json.FromBody(&request, c.Request.Body); has_error != nil {
		gin.Response(http.StatusBadRequest, nil, has_error)
		return
	}
	log.Printf("[info] Destination: %s", *request.Destination)

	client, err := gprc.NewGrpcClient(*request.Destination)
	if err != nil {
		gin.Response(http.StatusBadRequest, nil, err)
		return
	}
	defer client.Close()

	log.Printf("[info] Calling SendChallenge of grpc-client")
	err = client.SendChallenge()
	if err != nil {
		gin.Response(http.StatusBadRequest, nil, err)
		return
	}

	gin.Response(http.StatusOK, nil, nil)
}
