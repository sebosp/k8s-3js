package main

// Adapted from https://www.youtube.com/watch?v=qiB4RxCDC8o
import (
	"fmt"
	"log"
	"os"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func printPods(pods *v1.PodList) {
	template := "%-32s%-8s\n"
	fmt.Printf(template, "NAME", "STATUS")
	for _, pod := range pods.Items {
		fmt.Printf(
			template,
			pod.Name,
			string(pod.Status.Phase))
	}
}

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Error configuring kube client: %s", err)
		os.Exit(1)
	}
	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Printf("Error creatinga new client: %s", err)
		os.Exit(1)
	}
	api := clientSet.CoreV1()
	listOptions := metav1.ListOptions{}
	pods, err := api.Pods("default").List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	printPods(pods)
}
