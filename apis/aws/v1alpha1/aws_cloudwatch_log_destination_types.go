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

type AwsCloudwatchLogDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogDestinationSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogDestinationStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogDestinationSpec struct {
	Name      string `json:"name"`
	RoleArn   string `json:"role_arn"`
	TargetArn string `json:"target_arn"`
	Arn       string `json:"arn"`
}



type AwsCloudwatchLogDestinationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchLogDestinationList is a list of AwsCloudwatchLogDestinations
type AwsCloudwatchLogDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogDestination CRD objects
	Items []AwsCloudwatchLogDestination `json:"items,omitempty"`
}