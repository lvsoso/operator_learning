package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	//config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	pod := v1.Pod{}
	err = restClient.Get().Namespace("default").Resource("pods").Name("pi-4svs8").Do(context.TODO()).Into(&pod)
	if err != nil {
		println(err.Error())
	} else {
		println(pod.GetName())
	}
}
