package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleLoggingFolderExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingFolderExclusionSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingFolderExclusionStatus `json:"status,omitempty"`
}

type GoogleLoggingFolderExclusionSpec struct {
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	Folder      string `json:"folder"`
	Filter      string `json:"filter"`
	Name        string `json:"name"`
}

type GoogleLoggingFolderExclusionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleLoggingFolderExclusionList is a list of GoogleLoggingFolderExclusions
type GoogleLoggingFolderExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingFolderExclusion CRD objects
	Items []GoogleLoggingFolderExclusion `json:"items,omitempty"`
}
