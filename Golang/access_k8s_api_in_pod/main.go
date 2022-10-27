/*
Copyright 2022 Wade Xu
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	var namespaces []string
	if len(os.Args) > 1 {
		namespaces = os.Args[1:]
	} else {
		namespaces = []string{"default"}
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for _, ns := range namespaces {
		log.Printf("namespace: %s", ns)
		GetAllDeployName(clientset, ns)
		GetRunningContainerImage(clientset, ns)
	}
}

func GetAllDeployName(clientset *kubernetes.Clientset, namespace string) {
	deployment, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for id, deploy := range deployment.Items {
		log.Printf("%d: %s\n", id, deploy.Name)
	}

}

func GetRunningContainerImage(clientset *kubernetes.Clientset, namespace string) {
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("There are %d pods in namespace %s of the cluster\n", len(pods.Items), namespace)

	for _, pod := range pods.Items {
		if pod.Status.Phase == "Running" {
			for i, container := range pod.Spec.Containers {
				log.Printf("Name: %s, Version: %s\n", pod.ObjectMeta.Name+"-"+container.Name, pod.Spec.Containers[i].Image)
			}
		}
	}
}

func CheckPodIfExist(clientset *kubernetes.Clientset, podName string) {
	// error handling:
	_, err := clientset.CoreV1().Pods("app").Get(context.TODO(), podName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s not found in app namespace\n", podName)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found %s pod in app namespace\n", podName)
	}
}
