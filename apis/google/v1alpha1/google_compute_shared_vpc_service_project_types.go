package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSharedVpcServiceProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSharedVpcServiceProjectSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSharedVpcServiceProjectStatus `json:"status,omitempty"`
}

type GoogleComputeSharedVpcServiceProjectSpec struct {
	ServiceProject string `json:"service_project"`
	HostProject    string `json:"host_project"`
}



type GoogleComputeSharedVpcServiceProjectStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSharedVpcServiceProjectList is a list of GoogleComputeSharedVpcServiceProjects
type GoogleComputeSharedVpcServiceProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSharedVpcServiceProject CRD objects
	Items []GoogleComputeSharedVpcServiceProject `json:"items,omitempty"`
}