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

type AwsOpsworksPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksPermissionSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksPermissionStatus `json:"status,omitempty"`
}

type AwsOpsworksPermissionSpec struct {
	StackId   string `json:"stack_id"`
	AllowSsh  bool   `json:"allow_ssh"`
	AllowSudo bool   `json:"allow_sudo"`
	UserArn   string `json:"user_arn"`
	Level     string `json:"level"`
}

type AwsOpsworksPermissionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksPermissionList is a list of AwsOpsworksPermissions
type AwsOpsworksPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksPermission CRD objects
	Items []AwsOpsworksPermission `json:"items,omitempty"`
}
