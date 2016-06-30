package main

import (
	"fmt"
	"os"
	"time"

	"github.com/coreos/rkt/api/v1alpha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:15441", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := v1alpha.NewPublicAPIClient(conn)
	defer conn.Close()

	for {
		// List pods.
		resp, err := c.ListPods(context.Background(), &v1alpha.ListPodsRequest{Detail: true})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		time.Sleep(time.Second)
		fmt.Println("num of pods", len(resp.Pods))

	}
}
