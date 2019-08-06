package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudwatchEventPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventPermissionSpec   `json:"spec,omitempty"`
	Status            CloudwatchEventPermissionStatus `json:"status,omitempty"`
}

type CloudwatchEventPermissionSpecCondition struct {
	Key   string `json:"key" tf:"key"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"value" tf:"value"`
}

type CloudwatchEventPermissionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Action string `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Condition   []CloudwatchEventPermissionSpecCondition `json:"condition,omitempty" tf:"condition,omitempty"`
	Principal   string                                   `json:"principal" tf:"principal"`
	StatementID string                                   `json:"statementID" tf:"statement_id"`
}

type CloudwatchEventPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudwatchEventPermissionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchEventPermissionList is a list of CloudwatchEventPermissions
type CloudwatchEventPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchEventPermission CRD objects
	Items []CloudwatchEventPermission `json:"items,omitempty"`
}
