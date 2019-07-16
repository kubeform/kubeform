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

type Route53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverEndpointSpec   `json:"spec,omitempty"`
	Status            Route53ResolverEndpointStatus `json:"status,omitempty"`
}

type Route53ResolverEndpointSpecIpAddress struct {
	SubnetId string `json:"subnet_id"`
}

type Route53ResolverEndpointSpec struct {
	Direction string `json:"direction"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=2
	// +kubebuilder:validation:UniqueItems=true
	IpAddress []Route53ResolverEndpointSpec `json:"ip_address"`
	// +optional
	Name string `json:"name,omitempty"`
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type Route53ResolverEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ResolverEndpointList is a list of Route53ResolverEndpoints
type Route53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ResolverEndpoint CRD objects
	Items []Route53ResolverEndpoint `json:"items,omitempty"`
}
