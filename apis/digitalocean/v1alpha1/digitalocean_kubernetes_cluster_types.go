package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanKubernetesClusterSpec   `json:"spec,omitempty"`
	Status            DigitaloceanKubernetesClusterStatus `json:"status,omitempty"`
}

type DigitaloceanKubernetesClusterSpecNodePoolNodes struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

type DigitaloceanKubernetesClusterSpecNodePool struct {
	Nodes     []DigitaloceanKubernetesClusterSpecNodePool `json:"nodes"`
	Id        string                                      `json:"id"`
	Name      string                                      `json:"name"`
	Size      string                                      `json:"size"`
	NodeCount int                                         `json:"node_count"`
	Tags      []string                                    `json:"tags"`
}

type DigitaloceanKubernetesClusterSpecKubeConfig struct {
	ClientCertificate    string `json:"client_certificate"`
	RawConfig            string `json:"raw_config"`
	Host                 string `json:"host"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	ClientKey            string `json:"client_key"`
}

type DigitaloceanKubernetesClusterSpec struct {
	Region        string                              `json:"region"`
	ClusterSubnet string                              `json:"cluster_subnet"`
	ServiceSubnet string                              `json:"service_subnet"`
	Endpoint      string                              `json:"endpoint"`
	NodePool      []DigitaloceanKubernetesClusterSpec `json:"node_pool"`
	Status        string                              `json:"status"`
	UpdatedAt     string                              `json:"updated_at"`
	Name          string                              `json:"name"`
	Ipv4Address   string                              `json:"ipv4_address"`
	Tags          []string                            `json:"tags"`
	CreatedAt     string                              `json:"created_at"`
	KubeConfig    []DigitaloceanKubernetesClusterSpec `json:"kube_config"`
	Version       string                              `json:"version"`
}



type DigitaloceanKubernetesClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanKubernetesClusterList is a list of DigitaloceanKubernetesClusters
type DigitaloceanKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanKubernetesCluster CRD objects
	Items []DigitaloceanKubernetesCluster `json:"items,omitempty"`
}