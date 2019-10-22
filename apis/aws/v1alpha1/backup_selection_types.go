package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type BackupSelection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSelectionSpec   `json:"spec,omitempty"`
	Status            BackupSelectionStatus `json:"status,omitempty"`
}

type BackupSelectionSpecSelectionTag struct {
	Key   string `json:"key" tf:"key"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"value" tf:"value"`
}

type BackupSelectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	IamRoleArn string `json:"iamRoleArn" tf:"iam_role_arn"`
	Name       string `json:"name" tf:"name"`
	PlanID     string `json:"planID" tf:"plan_id"`
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources,omitempty"`
	// +optional
	SelectionTag []BackupSelectionSpecSelectionTag `json:"selectionTag,omitempty" tf:"selection_tag,omitempty"`
}

type BackupSelectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BackupSelectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BackupSelectionList is a list of BackupSelections
type BackupSelectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupSelection CRD objects
	Items []BackupSelection `json:"items,omitempty"`
}
