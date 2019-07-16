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

type IotThingType struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotThingTypeSpec   `json:"spec,omitempty"`
	Status            IotThingTypeStatus `json:"status,omitempty"`
}

type IotThingTypeSpecProperties struct {
	// +optional
	Description string `json:"description,omitempty"`
}

type IotThingTypeSpec struct {
	// +optional
	Deprecated bool   `json:"deprecated,omitempty"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Properties *[]IotThingTypeSpec `json:"properties,omitempty"`
}

type IotThingTypeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotThingTypeList is a list of IotThingTypes
type IotThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotThingType CRD objects
	Items []IotThingType `json:"items,omitempty"`
}
