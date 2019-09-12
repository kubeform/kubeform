package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecKubeConfig struct {
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"clientKey,omitempty" tf:"client_key,omitempty"`
	// +optional
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	RawConfig string `json:"rawConfig,omitempty" tf:"raw_config,omitempty"`
}

type KubernetesClusterSpecNodePoolNodes struct {
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	UpdatedAt string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type KubernetesClusterSpecNodePool struct {
	// +optional
	ID        string `json:"ID,omitempty" tf:"id,omitempty"`
	Name      string `json:"name" tf:"name"`
	NodeCount int    `json:"nodeCount" tf:"node_count"`
	// +optional
	Nodes []KubernetesClusterSpecNodePoolNodes `json:"nodes,omitempty" tf:"nodes,omitempty"`
	Size  string                               `json:"size" tf:"size"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ClusterSubnet string `json:"clusterSubnet,omitempty" tf:"cluster_subnet,omitempty"`
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	Ipv4Address string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`
	// +optional
	KubeConfig []KubernetesClusterSpecKubeConfig `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`
	Name       string                            `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	NodePool []KubernetesClusterSpecNodePool `json:"nodePool" tf:"node_pool"`
	Region   string                          `json:"region" tf:"region"`
	// +optional
	ServiceSubnet string `json:"serviceSubnet,omitempty" tf:"service_subnet,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UpdatedAt string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
	Version   string `json:"version" tf:"version"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KubernetesClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
