// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/intelops/tarian-detector/pkg/detector"
	"github.com/intelops/tarian-detector/pkg/k8s"
	"github.com/intelops/tarian-detector/tarian"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	// NotInClusterErrMsg is an error message for when the Kubernetes environment is not detected.
	NotInClusterErrMsg string = "Kubernetes environment not detected. The Kubernetes context has been disabled."
)

// main is the entry point of the application. It sets up the necessary components
// and starts the main event loop.
func main() {
	// Create a channel to listen for interrupt signals (Ctrl+C or SIGTERM)
	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	// Initialize and start the Kubernetes watcher
	watcher, err := K8Watcher()
	if err != nil {
		log.Print(err)
	}

	// Initialize Tarian eBPF module
	tarianEbpfModule, err := tarian.GetModule()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the Tarian detector by creating map readers
	tarianDetector, err := tarianEbpfModule.Prepare()
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate the event detectors
	eventsDetector := detector.NewEventsDetector()

	// Add the eBPF module to the detectors
	eventsDetector.Add(tarianDetector)
	eventsDetector.SetPodWatcher(watcher)

	// Start the event detectors and defer their closure
	err = eventsDetector.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer eventsDetector.Close()

	// Attaches tarian module programs to the kernel
	err = tarianEbpfModule.Attach(tarianDetector)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d probes running...\n\n", eventsDetector.Count())

	go func() {
		<-stopper // Wait for an interrupt signal

		eventsDetector.Close()
		log.Printf("Total records captured : %d\n", eventsDetector.GetTotalCount())
		count := 1
		for ky, vl := range eventsDetector.GetProbeCount() {
			fmt.Printf("%d. %s: %d\n", count, ky, vl)
			count++
		}
		os.Exit(0)
	}()

	// Continuously read events
	go func() {
		for {
			e, err := eventsDetector.ReadAsInterface()
			if err != nil {
				log.Print(err)
				continue
			}

			fmt.Println(e)
		}
	}()

	// Only for avoiding deadlock detection
	for {
		time.Sleep(1 * time.Minute)
	}
}

// K8Watcher initializes and returns a new PodWatcher for the current Kubernetes cluster.
func K8Watcher() (*k8s.PodWatcher, error) {
	// Get the in-cluster configuration.
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("%v. %s", err, NotInClusterErrMsg)
	}

	// Create a new Kubernetes client set.
	clientSet := kubernetes.NewForConfigOrDie(config)

	// Return a new PodWatcher for the current Kubernetes cluster.
	return k8s.NewPodWatcher(clientSet)
}
