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

type LambdaPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaPermissionSpec   `json:"spec,omitempty"`
	Status            LambdaPermissionStatus `json:"status,omitempty"`
}

type LambdaPermissionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action string `json:"action" tf:"action"`
	// +optional
	EventSourceToken string `json:"eventSourceToken,omitempty" tf:"event_source_token,omitempty"`
	FunctionName     string `json:"functionName" tf:"function_name"`
	Principal        string `json:"principal" tf:"principal"`
	// +optional
	Qualifier string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
	// +optional
	SourceAccount string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`
	// +optional
	SourceArn string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
	// +optional
	StatementID string `json:"statementID,omitempty" tf:"statement_id,omitempty"`
	// +optional
	StatementIDPrefix string `json:"statementIDPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

type LambdaPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LambdaPermissionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaPermissionList is a list of LambdaPermissions
type LambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaPermission CRD objects
	Items []LambdaPermission `json:"items,omitempty"`
}
