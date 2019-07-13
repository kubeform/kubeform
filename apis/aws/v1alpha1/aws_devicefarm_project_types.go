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

type AwsDevicefarmProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDevicefarmProjectSpec   `json:"spec,omitempty"`
	Status            AwsDevicefarmProjectStatus `json:"status,omitempty"`
}

type AwsDevicefarmProjectSpec struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}



type AwsDevicefarmProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDevicefarmProjectList is a list of AwsDevicefarmProjects
type AwsDevicefarmProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDevicefarmProject CRD objects
	Items []AwsDevicefarmProject `json:"items,omitempty"`
}