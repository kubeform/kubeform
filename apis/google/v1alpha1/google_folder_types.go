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

type GoogleFolder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleFolderSpec   `json:"spec,omitempty"`
	Status            GoogleFolderStatus `json:"status,omitempty"`
}

type GoogleFolderSpec struct {
	Parent         string `json:"parent"`
	DisplayName    string `json:"display_name"`
	Name           string `json:"name"`
	LifecycleState string `json:"lifecycle_state"`
	CreateTime     string `json:"create_time"`
}

type GoogleFolderStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleFolderList is a list of GoogleFolders
type GoogleFolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleFolder CRD objects
	Items []GoogleFolder `json:"items,omitempty"`
}
