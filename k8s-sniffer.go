package main

// Adapted from https://www.youtube.com/watch?v=qiB4RxCDC8o
import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Error configuring kube client: %s", err)
		os.Exit(1)
	}
	cl, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Printf("Error creatinga new client: %s", err)
		os.Exit(1)
	}
}
