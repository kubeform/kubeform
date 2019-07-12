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

type AwsGlacierVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlacierVaultSpec   `json:"spec,omitempty"`
	Status            AwsGlacierVaultStatus `json:"status,omitempty"`
}

type AwsGlacierVaultSpecNotification struct {
	Events   []string `json:"events"`
	SnsTopic string   `json:"sns_topic"`
}

type AwsGlacierVaultSpec struct {
	Tags         map[string]string     `json:"tags"`
	Name         string                `json:"name"`
	Location     string                `json:"location"`
	Arn          string                `json:"arn"`
	AccessPolicy string                `json:"access_policy"`
	Notification []AwsGlacierVaultSpec `json:"notification"`
}

type AwsGlacierVaultStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlacierVaultList is a list of AwsGlacierVaults
type AwsGlacierVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlacierVault CRD objects
	Items []AwsGlacierVault `json:"items,omitempty"`
}
