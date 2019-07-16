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

type ApiGatewayModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayModelSpec   `json:"spec,omitempty"`
	Status            ApiGatewayModelStatus `json:"status,omitempty"`
}

type ApiGatewayModelSpec struct {
	ContentType string `json:"content_type"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	RestApiId   string `json:"rest_api_id"`
	// +optional
	Schema string `json:"schema,omitempty"`
}

type ApiGatewayModelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayModelList is a list of ApiGatewayModels
type ApiGatewayModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayModel CRD objects
	Items []ApiGatewayModel `json:"items,omitempty"`
}
