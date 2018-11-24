package main

// Adapted from https://www.youtube.com/watch?v=qiB4RxCDC8o
import (
	"errors"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8s struct {
	clientSet kubernetes.Interface
}

func newK8s() (*k8s, error) {
	var cfg *restclient.Config
	var err error
	if os.Getenv("INSIDE_K8S") == "" {
		path := os.Getenv("HOME") + "/.kube/config"
		cfg, err = clientcmd.BuildConfigFromFlags("", path)
		if err != nil {
			os.Exit(1)
		}
	} else {
		cfg, err = restclient.InClusterConfig()
		if err != nil {
			log.Printf("Error configuring kube client: %s", err)
			os.Exit(1)
		}
	}
	client := k8s{}
	client.clientSet, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Printf("Error creating a new client: %s", err)
		os.Exit(1)
	}
	return &client, nil
}

func (o *k8s) listPods() ([]string, error) {
	var res []string
	api := o.clientSet.CoreV1()
	listOptions := metav1.ListOptions{}
	pods, err := api.Pods("kube-system").List(listOptions)
	if err != nil {
		log.Fatal(err)
		return res, errors.New("Unable to list pods")
	}
	template := "%-32s%-8s\n"
	fmt.Printf(template, "NAME", "STATUS")
	for _, pod := range pods.Items {
		res = append(res, fmt.Sprintf(
			template,
			pod.Name,
			string(pod.Status.Phase)))
	}
	return res, nil
}

func main() {
	k8s, err := newK8s()
	if err != nil {
		fmt.Println(err)
		return
	}
	if pods, err := k8s.listPods(); err == nil {
		for _, pod := range pods {
			fmt.Println(pod)
		}
	}
}
