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

type CloudwatchLogGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogGroupSpec   `json:"spec,omitempty"`
	Status            CloudwatchLogGroupStatus `json:"status,omitempty"`
}

type CloudwatchLogGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	RetentionInDays int `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudwatchLogGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudwatchLogGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchLogGroupList is a list of CloudwatchLogGroups
type CloudwatchLogGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchLogGroup CRD objects
	Items []CloudwatchLogGroup `json:"items,omitempty"`
}
