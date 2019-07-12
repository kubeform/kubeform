package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodedeployAppSpec   `json:"spec,omitempty"`
	Status            AwsCodedeployAppStatus `json:"status,omitempty"`
}

type AwsCodedeployAppSpec struct {
	Name            string `json:"name"`
	ComputePlatform string `json:"compute_platform"`
	UniqueId        string `json:"unique_id"`
}

type AwsCodedeployAppStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployAppList is a list of AwsCodedeployApps
type AwsCodedeployAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodedeployApp CRD objects
	Items []AwsCodedeployApp `json:"items,omitempty"`
}
