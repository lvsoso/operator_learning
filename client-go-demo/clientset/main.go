package main

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// get data
	coreV1 := clientset.CoreV1()
	pod, err := coreV1.Pods("default").Get(context.TODO(), "pi-4svs8", v1.GetOptions{})
	if err != nil {
		println(err.Error())
	} else {
		println(pod.GetName())
	}
}
