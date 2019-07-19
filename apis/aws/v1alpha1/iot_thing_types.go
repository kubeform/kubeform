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

type IotThing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotThingSpec   `json:"spec,omitempty"`
	Status            IotThingStatus `json:"status,omitempty"`
}

type IotThingSpec struct {
	// +optional
	Attributes map[string]string `json:"attributes,omitempty" tf:"attributes,omitempty"`
	Name       string            `json:"name" tf:"name"`
	// +optional
	ThingTypeName string                    `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`
	ProviderRef   core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IotThingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotThingList is a list of IotThings
type IotThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotThing CRD objects
	Items []IotThing `json:"items,omitempty"`
}
