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

type IothubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSharedAccessPolicySpec   `json:"spec,omitempty"`
	Status            IothubSharedAccessPolicyStatus `json:"status,omitempty"`
}

type IothubSharedAccessPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	DeviceConnect           bool   `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`
	IothubName              string `json:"iothubName" tf:"iothub_name"`
	Name                    string `json:"name" tf:"name"`
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	PrimaryKey              string `json:"-" sensitive:"true" tf:"primary_key"`
	// +optional
	RegistryRead bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`
	// +optional
	RegistryWrite             bool   `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	SecondaryKey              string `json:"-" sensitive:"true" tf:"secondary_key"`
	// +optional
	ServiceConnect bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

type IothubSharedAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubSharedAccessPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubSharedAccessPolicyList is a list of IothubSharedAccessPolicys
type IothubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IothubSharedAccessPolicy CRD objects
	Items []IothubSharedAccessPolicy `json:"items,omitempty"`
}
