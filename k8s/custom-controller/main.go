package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

const kubeconfig = "/home/murtaza/.kube/config"

func main() {
	conf, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		log.Fatal(err)
	}

	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			fmt.Println("a pod is added")
		},
		UpdateFunc: func(oldObj, newObj any) {
			fmt.Println("a pod is updated")
		},
		DeleteFunc: func(obj any) {
			fmt.Println("a pod is deleted")
		},
	})

	stop := make(chan struct{})
	go informer.Run(stop)

	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)

	<-sigint
	close(stop)
	time.Sleep(time.Millisecond * 10)
}
