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

type ApiGatewayVpcLink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayVpcLinkSpec   `json:"spec,omitempty"`
	Status            ApiGatewayVpcLinkStatus `json:"status,omitempty"`
}

type ApiGatewayVpcLinkSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	TargetArns []string `json:"target_arns"`
}

type ApiGatewayVpcLinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayVpcLinkList is a list of ApiGatewayVpcLinks
type ApiGatewayVpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayVpcLink CRD objects
	Items []ApiGatewayVpcLink `json:"items,omitempty"`
}
