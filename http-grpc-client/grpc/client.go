package grpc

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ressley/test_task_go_middle/pkg/ciphers"
	service "github.com/ressley/test_task_go_middle/pkg/eventBus_v1"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn    *grpc.ClientConn
	service service.EventServiceClient
}

func (c *GrpcClient) Close() {
	c.conn.Close()
}

func NewGrpcClient(address string) (*GrpcClient, error) {
	log.Printf("[info] Creating new connection with grpc client on address: %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("[error] Failed to connect address: %s, Error: %v", address, err)
		return nil, err
	}

	client := service.NewEventServiceClient(conn)

	return &GrpcClient{
		conn:    conn,
		service: client,
	}, nil
}

func (c *GrpcClient) SendChallenge() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("[info] Generating random bytes")
	challengeData := make([]byte, 32)
	_, err := rand.Read(challengeData)
	if err != nil {
		log.Printf("[error] Error on generating random bytes: %v", err)
		return errors.New(fmt.Sprintf("Error on generating random bytes %v", err))
	}

	log.Printf("[info] Sending bytes to gRPC server")
	r, err := c.service.EventBus(ctx, &service.Event{
		Data: challengeData,
	})
	if err != nil {
		log.Printf("[error] Error while sending bytes to gRPC server: %v", err)
		return errors.New(fmt.Sprintf("Error while sending bytes to gRPC server: %v", err))
	}

	log.Printf("[info] Response from server %v", r)
	return nil
}

func (c *GrpcClient) SendMessage(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("[info] Generating rsa keys")
	_, publicKey, err := ciphers.GenerateKeyPair(2048)
	if err != nil {
		log.Printf("[error] Error generating RSA key: %v", err)
		return errors.New(fmt.Sprintf("Error generating RSA key %v", err))
	}

	log.Printf("[info] Encripting with public key")
	encryptedData, err := ciphers.EncryptWithPublicKey(data, publicKey)
	if err != nil {
		log.Printf("[error] Error encrypting data: %v", err)
		return errors.New(fmt.Sprintf("Error encrypting data %v", err))
	}

	log.Printf("[info] Sending bytes to gRPC server")
	r, err := c.service.EventBus(ctx, &service.Event{
		Data: encryptedData,
	})

	if err != nil {
		log.Printf("[error] Error generating challenge: %v", err)
		return errors.New(fmt.Sprintf("Error generating challenge %v", err))
	}

	log.Printf("[info] Response from server %v", r)
	return nil
}
