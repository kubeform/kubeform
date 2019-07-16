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

type VpcEndpointServiceAllowedPrincipal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointServiceAllowedPrincipalSpec   `json:"spec,omitempty"`
	Status            VpcEndpointServiceAllowedPrincipalStatus `json:"status,omitempty"`
}

type VpcEndpointServiceAllowedPrincipalSpec struct {
	PrincipalArn         string `json:"principal_arn"`
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}

type VpcEndpointServiceAllowedPrincipalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointServiceAllowedPrincipalList is a list of VpcEndpointServiceAllowedPrincipals
type VpcEndpointServiceAllowedPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointServiceAllowedPrincipal CRD objects
	Items []VpcEndpointServiceAllowedPrincipal `json:"items,omitempty"`
}
