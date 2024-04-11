// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

import (
	"github.com/intelops/tarian-detector/pkg/err"
	"github.com/intelops/tarian-detector/pkg/k8s"
)

var enricherErr = err.New("kubernetes")

// GetK8sContext returns the Kubernetes context for a given process ID.
func GetK8sContext(watcher *k8s.PodWatcher, processId uint32) (Kubernetes, error) {
	k8sCtx := Kubernetes{}
	// Get the container ID for the given process ID.
	containerId, err := k8s.ProcsContainerID(processId)
	if err != nil {
		return k8sCtx, enricherErr.Throwf("%v", err)
	}

	// If the container ID is missing, return an error.
	if len(containerId) == 0 {
		return k8sCtx, enricherErr.Throw("missing container id")
	}

	// Find the pod associated with the container ID.
	pod, err := watcher.FindPod(containerId)
	if err != nil {
		return k8sCtx, enricherErr.Throwf("%v: unable to find the pod associated with the container ID: %s", err, containerId)
	}

	// Set the pod information in the Kubernetes context.
	k8sCtx.PodUid = string(pod.ObjectMeta.UID)
	k8sCtx.PodName = pod.ObjectMeta.Name
	k8sCtx.PodGeneratedName = pod.ObjectMeta.GenerateName
	k8sCtx.PodKind = pod.Kind
	k8sCtx.PodAPIVersion = pod.APIVersion
	k8sCtx.PodLabels = pod.ObjectMeta.Labels
	k8sCtx.PodAnnotations = pod.ObjectMeta.Annotations

	// Set the container information in the Kubernetes context.
	k8sCtx.ContainerID = containerId

	// Set the namespace information in the Kubernetes context.
	k8sCtx.Namespace = pod.ObjectMeta.Namespace

	return k8sCtx, nil
}
