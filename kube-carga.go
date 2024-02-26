package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file")
	flag.Parse()

	// Use the in-cluster config if kubeconfig path is not provided
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a new Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Create a context for the test
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to handle OS signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start load testing
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// Perform some workload here
				// For example, create a deployment
				deployment := &appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-" + strconv.Itoa(rand.Intn(100)),
						Namespace: "default",
					},
					Spec: appsv1.DeploymentSpec{
						Replicas: int32Ptr(1),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"app": "test",
							},
						},
						Template: corev1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"app": "test",
								},
							},
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name:  "nginx",
										Image: "nginx",
									},
								},
							},
						},
					},
				}
				_, err := clientset.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
				if err != nil {
					log.Printf("Error creating deployment: %v\n", err)
				} else {
					log.Printf("Deployment %s created successfully\n", deployment.Name)
				}
			}
		}
	}()

	// Wait for termination signal
	<-sigCh
	log.Println("Received termination signal. Stopping test.")
}

func int32Ptr(i int32) *int32 { return &i }
