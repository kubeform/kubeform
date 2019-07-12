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

type AwsAmiLaunchPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAmiLaunchPermissionSpec   `json:"spec,omitempty"`
	Status            AwsAmiLaunchPermissionStatus `json:"status,omitempty"`
}

type AwsAmiLaunchPermissionSpec struct {
	ImageId   string `json:"image_id"`
	AccountId string `json:"account_id"`
}

type AwsAmiLaunchPermissionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAmiLaunchPermissionList is a list of AwsAmiLaunchPermissions
type AwsAmiLaunchPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAmiLaunchPermission CRD objects
	Items []AwsAmiLaunchPermission `json:"items,omitempty"`
}
