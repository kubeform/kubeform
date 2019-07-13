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

type AwsCloudwatchEventPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchEventPermissionSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchEventPermissionStatus `json:"status,omitempty"`
}

type AwsCloudwatchEventPermissionSpecCondition struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCloudwatchEventPermissionSpec struct {
	Principal   string                             `json:"principal"`
	StatementId string                             `json:"statement_id"`
	Action      string                             `json:"action"`
	Condition   []AwsCloudwatchEventPermissionSpec `json:"condition"`
}



type AwsCloudwatchEventPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchEventPermissionList is a list of AwsCloudwatchEventPermissions
type AwsCloudwatchEventPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchEventPermission CRD objects
	Items []AwsCloudwatchEventPermission `json:"items,omitempty"`
}