package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	Attributes      map[string]string `json:"attributes,omitempty" tf:"attributes,omitempty"`
	DefaultClientID string            `json:"defaultClientID" tf:"default_client_id"`
	Name            string            `json:"name" tf:"name"`
	// +optional
	ThingTypeName string `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`
	Version       int    `json:"version" tf:"version"`
}

type IotThingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IotThingSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
