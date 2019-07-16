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

type EksCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksClusterSpec   `json:"spec,omitempty"`
	Status            EksClusterStatus `json:"status,omitempty"`
}

type EksClusterSpecVpcConfig struct {
	// +optional
	EndpointPrivateAccess bool `json:"endpoint_private_access,omitempty"`
	// +optional
	EndpointPublicAccess bool `json:"endpoint_public_access,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
}

type EksClusterSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledClusterLogTypes []string `json:"enabled_cluster_log_types,omitempty"`
	Name                   string   `json:"name"`
	RoleArn                string   `json:"role_arn"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	VpcConfig []EksClusterSpec `json:"vpc_config"`
}

type EksClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EksClusterList is a list of EksClusters
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EksCluster CRD objects
	Items []EksCluster `json:"items,omitempty"`
}
