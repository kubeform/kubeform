package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type SignalrService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignalrServiceSpec   `json:"spec,omitempty"`
	Status            SignalrServiceStatus `json:"status,omitempty"`
}

type SignalrServiceSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
}

type SignalrServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	Location  string `json:"location" tf:"location"`
	Name      string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PublicPort        int    `json:"publicPort,omitempty" tf:"public_port,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	ServerPort int `json:"serverPort,omitempty" tf:"server_port,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []SignalrServiceSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SignalrServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SignalrServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SignalrServiceList is a list of SignalrServices
type SignalrServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SignalrService CRD objects
	Items []SignalrService `json:"items,omitempty"`
}
