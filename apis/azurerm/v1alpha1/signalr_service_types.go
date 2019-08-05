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

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Hostname                  string `json:"hostname" tf:"hostname"`
	IpAddress                 string `json:"ipAddress" tf:"ip_address"`
	Location                  string `json:"location" tf:"location"`
	Name                      string `json:"name" tf:"name"`
	PrimaryAccessKey          string `json:"-" sensitive:"true" tf:"primary_access_key"`
	PrimaryConnectionString   string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	PublicPort                int    `json:"publicPort" tf:"public_port"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	SecondaryAccessKey        string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	ServerPort                int    `json:"serverPort" tf:"server_port"`
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
	State *apis.State `json:"state,omitempty"`
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
