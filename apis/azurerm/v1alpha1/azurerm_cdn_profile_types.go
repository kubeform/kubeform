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

type AzurermCdnProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCdnProfileSpec   `json:"spec,omitempty"`
	Status            AzurermCdnProfileStatus `json:"status,omitempty"`
}

type AzurermCdnProfileSpec struct {
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
	Sku               string            `json:"sku"`
	Tags              map[string]string `json:"tags"`
}

type AzurermCdnProfileStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCdnProfileList is a list of AzurermCdnProfiles
type AzurermCdnProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCdnProfile CRD objects
	Items []AzurermCdnProfile `json:"items,omitempty"`
}
