package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSharedVpcHostProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSharedVpcHostProjectSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSharedVpcHostProjectStatus `json:"status,omitempty"`
}

type GoogleComputeSharedVpcHostProjectSpec struct {
	Project string `json:"project"`
}



type GoogleComputeSharedVpcHostProjectStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSharedVpcHostProjectList is a list of GoogleComputeSharedVpcHostProjects
type GoogleComputeSharedVpcHostProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSharedVpcHostProject CRD objects
	Items []GoogleComputeSharedVpcHostProject `json:"items,omitempty"`
}