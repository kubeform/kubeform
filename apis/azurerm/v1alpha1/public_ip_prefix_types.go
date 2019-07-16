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

type PublicIpPrefix struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIpPrefixSpec   `json:"spec,omitempty"`
	Status            PublicIpPrefixStatus `json:"status,omitempty"`
}

type PublicIpPrefixSpec struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	// +optional
	PrefixLength      int    `json:"prefix_length,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type PublicIpPrefixStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PublicIpPrefixList is a list of PublicIpPrefixs
type PublicIpPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PublicIpPrefix CRD objects
	Items []PublicIpPrefix `json:"items,omitempty"`
}