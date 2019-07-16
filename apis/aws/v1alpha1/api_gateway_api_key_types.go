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

type ApiGatewayApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayApiKeySpec   `json:"spec,omitempty"`
	Status            ApiGatewayApiKeyStatus `json:"status,omitempty"`
}

type ApiGatewayApiKeySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Enabled bool   `json:"enabled,omitempty"`
	Name    string `json:"name"`
}

type ApiGatewayApiKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayApiKeyList is a list of ApiGatewayApiKeys
type ApiGatewayApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayApiKey CRD objects
	Items []ApiGatewayApiKey `json:"items,omitempty"`
}
