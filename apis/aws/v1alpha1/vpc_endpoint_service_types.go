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

type VpcEndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointServiceSpec   `json:"spec,omitempty"`
	Status            VpcEndpointServiceStatus `json:"status,omitempty"`
}

type VpcEndpointServiceSpec struct {
	AcceptanceRequired bool `json:"acceptance_required"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	NetworkLoadBalancerArns []string `json:"network_load_balancer_arns"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type VpcEndpointServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointServiceList is a list of VpcEndpointServices
type VpcEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointService CRD objects
	Items []VpcEndpointService `json:"items,omitempty"`
}
