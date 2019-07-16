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

type ApiGatewayDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDeploymentSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDeploymentStatus `json:"status,omitempty"`
}

type ApiGatewayDeploymentSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	RestApiId   string `json:"rest_api_id"`
	// +optional
	StageDescription string `json:"stage_description,omitempty"`
	// +optional
	StageName string `json:"stage_name,omitempty"`
	// +optional
	Variables map[string]string `json:"variables,omitempty"`
}

type ApiGatewayDeploymentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayDeploymentList is a list of ApiGatewayDeployments
type ApiGatewayDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayDeployment CRD objects
	Items []ApiGatewayDeployment `json:"items,omitempty"`
}
