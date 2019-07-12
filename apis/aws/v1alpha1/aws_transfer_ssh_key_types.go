package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsTransferSshKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsTransferSshKeySpec   `json:"spec,omitempty"`
	Status            AwsTransferSshKeyStatus `json:"status,omitempty"`
}

type AwsTransferSshKeySpec struct {
	Body     string `json:"body"`
	ServerId string `json:"server_id"`
	UserName string `json:"user_name"`
}

type AwsTransferSshKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsTransferSshKeyList is a list of AwsTransferSshKeys
type AwsTransferSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsTransferSshKey CRD objects
	Items []AwsTransferSshKey `json:"items,omitempty"`
}
