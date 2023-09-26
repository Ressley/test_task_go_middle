package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ressley/test_task_go_middle/http-grpc-client/http"
)

func main() {
	server := http.StartServer()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server")
	server.Close()
}
