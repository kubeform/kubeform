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

type AwsCloudwatchLogStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogStreamSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogStreamStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogStreamSpec struct {
	Arn          string `json:"arn"`
	Name         string `json:"name"`
	LogGroupName string `json:"log_group_name"`
}



type AwsCloudwatchLogStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchLogStreamList is a list of AwsCloudwatchLogStreams
type AwsCloudwatchLogStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogStream CRD objects
	Items []AwsCloudwatchLogStream `json:"items,omitempty"`
}