package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnActivity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSfnActivitySpec   `json:"spec,omitempty"`
	Status            AwsSfnActivityStatus `json:"status,omitempty"`
}

type AwsSfnActivitySpec struct {
	Name         string            `json:"name"`
	CreationDate string            `json:"creation_date"`
	Tags         map[string]string `json:"tags"`
}

type AwsSfnActivityStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSfnActivityList is a list of AwsSfnActivitys
type AwsSfnActivityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSfnActivity CRD objects
	Items []AwsSfnActivity `json:"items,omitempty"`
}
