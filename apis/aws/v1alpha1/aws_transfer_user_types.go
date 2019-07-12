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

type AwsTransferUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsTransferUserSpec   `json:"spec,omitempty"`
	Status            AwsTransferUserStatus `json:"status,omitempty"`
}

type AwsTransferUserSpec struct {
	ServerId      string            `json:"server_id"`
	Tags          map[string]string `json:"tags"`
	UserName      string            `json:"user_name"`
	Arn           string            `json:"arn"`
	HomeDirectory string            `json:"home_directory"`
	Policy        string            `json:"policy"`
	Role          string            `json:"role"`
}

type AwsTransferUserStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsTransferUserList is a list of AwsTransferUsers
type AwsTransferUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsTransferUser CRD objects
	Items []AwsTransferUser `json:"items,omitempty"`
}
