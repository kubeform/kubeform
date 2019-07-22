package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DefaultSubnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSubnetSpec   `json:"spec,omitempty"`
	Status            DefaultSubnetStatus `json:"status,omitempty"`
}

type DefaultSubnetSpec struct {
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
	AvailabilityZone string                    `json:"availabilityZone" tf:"availability_zone"`
	// +optional
	MapPublicIPOnLaunch bool `json:"mapPublicIPOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultSubnetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultSubnetList is a list of DefaultSubnets
type DefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultSubnet CRD objects
	Items []DefaultSubnet `json:"items,omitempty"`
}
