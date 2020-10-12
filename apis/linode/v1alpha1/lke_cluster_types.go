/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LkeCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LkeClusterSpec   `json:"spec,omitempty"`
	Status            LkeClusterStatus `json:"status,omitempty"`
}

type LkeClusterSpecPoolNodes struct {
	// The ID of the node.
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// The ID of the underlying Linode instance.
	// +optional
	InstanceID int64 `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// The status of the node.
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
}

type LkeClusterSpecPool struct {
	// The number of nodes in the Node Pool.
	Count int64 `json:"count" tf:"count"`
	// The ID of the Node Pool.
	// +optional
	ID int64 `json:"ID,omitempty" tf:"id,omitempty"`
	// The nodes in the node pool.
	// +optional
	Nodes []LkeClusterSpecPoolNodes `json:"nodes,omitempty" tf:"nodes,omitempty"`
	// A Linode Type for all of the nodes in the Node Pool.
	Type string `json:"type" tf:"type"`
}

type LkeClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// The API endpoints for the cluster.
	// +optional
	ApiEndpoints []string `json:"apiEndpoints,omitempty" tf:"api_endpoints,omitempty"`
	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	K8sVersion string `json:"k8sVersion" tf:"k8s_version"`
	// The Base64-encoded Kubeconfig for the cluster.
	// +optional
	Kubeconfig string `json:"-" sensitive:"true" tf:"kubeconfig,omitempty"`
	// The unique label for the cluster.
	Label string `json:"label" tf:"label"`
	// A node pool in the cluster.
	// +kubebuilder:validation:MinItems=1
	Pool []LkeClusterSpecPool `json:"pool" tf:"pool"`
	// This cluster's location.
	Region string `json:"region" tf:"region"`
	// The status of the cluster.
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LkeClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LkeClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LkeClusterList is a list of LkeClusters
type LkeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LkeCluster CRD objects
	Items []LkeCluster `json:"items,omitempty"`
}
