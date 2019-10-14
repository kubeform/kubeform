package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ApiGatewayDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDeploymentSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDeploymentStatus `json:"status,omitempty"`
}

type ApiGatewayDeploymentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExecutionArn string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`
	// +optional
	InvokeURL string `json:"invokeURL,omitempty" tf:"invoke_url,omitempty"`
	RestAPIID string `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	StageDescription string `json:"stageDescription,omitempty" tf:"stage_description,omitempty"`
	// +optional
	StageName string `json:"stageName,omitempty" tf:"stage_name,omitempty"`
	// +optional
	Variables map[string]string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type ApiGatewayDeploymentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayDeploymentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
