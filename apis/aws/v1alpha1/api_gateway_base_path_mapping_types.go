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

type ApiGatewayBasePathMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayBasePathMappingSpec   `json:"spec,omitempty"`
	Status            ApiGatewayBasePathMappingStatus `json:"status,omitempty"`
}

type ApiGatewayBasePathMappingSpec struct {
	ApiId string `json:"api_id"`
	// +optional
	BasePath   string `json:"base_path,omitempty"`
	DomainName string `json:"domain_name"`
	// +optional
	StageName string `json:"stage_name,omitempty"`
}

type ApiGatewayBasePathMappingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayBasePathMappingList is a list of ApiGatewayBasePathMappings
type ApiGatewayBasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayBasePathMapping CRD objects
	Items []ApiGatewayBasePathMapping `json:"items,omitempty"`
}
