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

type AwsShieldProtection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsShieldProtectionSpec   `json:"spec,omitempty"`
	Status            AwsShieldProtectionStatus `json:"status,omitempty"`
}

type AwsShieldProtectionSpec struct {
	Name        string `json:"name"`
	ResourceArn string `json:"resource_arn"`
}

type AwsShieldProtectionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsShieldProtectionList is a list of AwsShieldProtections
type AwsShieldProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsShieldProtection CRD objects
	Items []AwsShieldProtection `json:"items,omitempty"`
}
