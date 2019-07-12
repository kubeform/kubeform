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

type AwsMediaStoreContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMediaStoreContainerSpec   `json:"spec,omitempty"`
	Status            AwsMediaStoreContainerStatus `json:"status,omitempty"`
}

type AwsMediaStoreContainerSpec struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	Arn      string `json:"arn"`
}

type AwsMediaStoreContainerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMediaStoreContainerList is a list of AwsMediaStoreContainers
type AwsMediaStoreContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMediaStoreContainer CRD objects
	Items []AwsMediaStoreContainer `json:"items,omitempty"`
}
