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
	ClientCertificate    string `json:"clientCertificate" tf:"client_certificate"`
	ClientKey            string `json:"clientKey" tf:"client_key"`
	ClusterCaCertificate string `json:"clusterCaCertificate" tf:"cluster_ca_certificate"`
	Host                 string `json:"host" tf:"host"`
	RawConfig            string `json:"rawConfig" tf:"raw_config"`
}

type KubernetesClusterSpecNodePoolNodes struct {
	CreatedAt string `json:"createdAt" tf:"created_at"`
	ID        string `json:"ID" tf:"id"`
	Name      string `json:"name" tf:"name"`
	Status    string `json:"status" tf:"status"`
	UpdatedAt string `json:"updatedAt" tf:"updated_at"`
}

type KubernetesClusterSpecNodePool struct {
	ID        string                               `json:"ID" tf:"id"`
	Name      string                               `json:"name" tf:"name"`
	NodeCount int                                  `json:"nodeCount" tf:"node_count"`
	Nodes     []KubernetesClusterSpecNodePoolNodes `json:"nodes" tf:"nodes"`
	Size      string                               `json:"size" tf:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ClusterSubnet string                            `json:"clusterSubnet" tf:"cluster_subnet"`
	CreatedAt     string                            `json:"createdAt" tf:"created_at"`
	Endpoint      string                            `json:"endpoint" tf:"endpoint"`
	Ipv4Address   string                            `json:"ipv4Address" tf:"ipv4_address"`
	KubeConfig    []KubernetesClusterSpecKubeConfig `json:"kubeConfig" tf:"kube_config"`
	Name          string                            `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	NodePool      []KubernetesClusterSpecNodePool `json:"nodePool" tf:"node_pool"`
	Region        string                          `json:"region" tf:"region"`
	ServiceSubnet string                          `json:"serviceSubnet" tf:"service_subnet"`
	Status        string                          `json:"status" tf:"status"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags      []string `json:"tags,omitempty" tf:"tags,omitempty"`
	UpdatedAt string   `json:"updatedAt" tf:"updated_at"`
	Version   string   `json:"version" tf:"version"`
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
