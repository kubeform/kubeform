package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServicebusNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusNamespaceSpec   `json:"spec,omitempty"`
	Status            ServicebusNamespaceStatus `json:"status,omitempty"`
}

type ServicebusNamespaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Capacity int `json:"capacity,omitempty" tf:"capacity,omitempty"`
	// +optional
	DefaultPrimaryConnectionString string `json:"-" sensitive:"true" tf:"default_primary_connection_string,omitempty"`
	// +optional
	DefaultPrimaryKey string `json:"-" sensitive:"true" tf:"default_primary_key,omitempty"`
	// +optional
	DefaultSecondaryConnectionString string `json:"-" sensitive:"true" tf:"default_secondary_connection_string,omitempty"`
	// +optional
	DefaultSecondaryKey string `json:"-" sensitive:"true" tf:"default_secondary_key,omitempty"`
	Location            string `json:"location" tf:"location"`
	Name                string `json:"name" tf:"name"`
	ResourceGroupName   string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                 string `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type ServicebusNamespaceStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ServicebusNamespaceSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusNamespaceList is a list of ServicebusNamespaces
type ServicebusNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusNamespace CRD objects
	Items []ServicebusNamespace `json:"items,omitempty"`
}