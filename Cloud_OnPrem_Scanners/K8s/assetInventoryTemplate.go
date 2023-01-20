package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)
/*
Use the Kubernetes Go client library to connect to a Kubernetes cluster using a kubeconfig file. Then you can list all the pods, services, deployments, statefulsets and configmaps and print them out. This can be used to get an inventory of all the assets running in a kubernetes cluster.
*/
func main() {
	// Create a config for the Kubernetes client
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Perform an inventory of pods
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for _, pod := range pods.Items {
		fmt.Println("Pod:", pod.Name)
	}

	// Perform an inventory of services
	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
if err != nil {
fmt.Println("Error", err)
os.Exit(1)
}
for _, service := range services.Items {
fmt.Println("Service:", service.Name)
}
// Perform an inventory of deployments
deployments, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})
if err != nil {
	fmt.Println("Error", err)
	os.Exit(1)
}
for _, deployment := range deployments.Items {
	fmt.Println("Deployment:", deployment.Name)
}

// Perform an inventory of statefulsets
statefulsets, err := clientset.AppsV1().StatefulSets("").List(metav1.ListOptions{})
if err != nil {
	fmt.Println("Error", err)
	os.Exit(1)
}
for _, statefulset := range statefulsets.Items {
	fmt.Println("StatefulSet:", statefulset.Name)
}

// Perform an inventory of configmaps
configmaps, err := clientset.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
if err != nil {
	fmt.Println("Error", err)
	os.Exit(1)
}
for _, configmap := range configmaps.Items {
	fmt.Println("ConfigMap:", configmap.Name)
}
}

