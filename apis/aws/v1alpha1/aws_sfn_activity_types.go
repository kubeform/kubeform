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

type AwsSfnActivity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSfnActivitySpec   `json:"spec,omitempty"`
	Status            AwsSfnActivityStatus `json:"status,omitempty"`
}

type AwsSfnActivitySpec struct {
	CreationDate string            `json:"creation_date"`
	Tags         map[string]string `json:"tags"`
	Name         string            `json:"name"`
}



type AwsSfnActivityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSfnActivityList is a list of AwsSfnActivitys
type AwsSfnActivityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSfnActivity CRD objects
	Items []AwsSfnActivity `json:"items,omitempty"`
}