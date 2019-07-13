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

type AwsBackupSelection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBackupSelectionSpec   `json:"spec,omitempty"`
	Status            AwsBackupSelectionStatus `json:"status,omitempty"`
}

type AwsBackupSelectionSpecSelectionTag struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AwsBackupSelectionSpec struct {
	SelectionTag []AwsBackupSelectionSpec `json:"selection_tag"`
	Resources    []string                 `json:"resources"`
	Name         string                   `json:"name"`
	PlanId       string                   `json:"plan_id"`
	IamRoleArn   string                   `json:"iam_role_arn"`
}



type AwsBackupSelectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsBackupSelectionList is a list of AwsBackupSelections
type AwsBackupSelectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBackupSelection CRD objects
	Items []AwsBackupSelection `json:"items,omitempty"`
}