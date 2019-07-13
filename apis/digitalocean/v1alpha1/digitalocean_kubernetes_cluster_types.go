package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DigitaloceanKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanKubernetesClusterSpec   `json:"spec,omitempty"`
	Status            DigitaloceanKubernetesClusterStatus `json:"status,omitempty"`
}

type DigitaloceanKubernetesClusterSpecKubeConfig struct {
	Host                 string `json:"host"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	ClientKey            string `json:"client_key"`
	ClientCertificate    string `json:"client_certificate"`
	RawConfig            string `json:"raw_config"`
}

type DigitaloceanKubernetesClusterSpecNodePoolNodes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DigitaloceanKubernetesClusterSpecNodePool struct {
	Name      string                                      `json:"name"`
	Size      string                                      `json:"size"`
	NodeCount int                                         `json:"node_count"`
	Tags      []string                                    `json:"tags"`
	Nodes     []DigitaloceanKubernetesClusterSpecNodePool `json:"nodes"`
	Id        string                                      `json:"id"`
}

type DigitaloceanKubernetesClusterSpec struct {
	UpdatedAt     string                              `json:"updated_at"`
	KubeConfig    []DigitaloceanKubernetesClusterSpec `json:"kube_config"`
	Region        string                              `json:"region"`
	Version       string                              `json:"version"`
	Endpoint      string                              `json:"endpoint"`
	NodePool      []DigitaloceanKubernetesClusterSpec `json:"node_pool"`
	Status        string                              `json:"status"`
	CreatedAt     string                              `json:"created_at"`
	Name          string                              `json:"name"`
	ClusterSubnet string                              `json:"cluster_subnet"`
	ServiceSubnet string                              `json:"service_subnet"`
	Ipv4Address   string                              `json:"ipv4_address"`
	Tags          []string                            `json:"tags"`
}



type DigitaloceanKubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanKubernetesClusterList is a list of DigitaloceanKubernetesClusters
type DigitaloceanKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanKubernetesCluster CRD objects
	Items []DigitaloceanKubernetesCluster `json:"items,omitempty"`
}