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

type CloudwatchLogGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogGroupSpec   `json:"spec,omitempty"`
	Status            CloudwatchLogGroupStatus `json:"status,omitempty"`
}

type CloudwatchLogGroupSpec struct {
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	RetentionInDays int `json:"retention_in_days,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type CloudwatchLogGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
