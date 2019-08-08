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

type RelayNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RelayNamespaceSpec   `json:"spec,omitempty"`
	Status            RelayNamespaceStatus `json:"status,omitempty"`
}

type RelayNamespaceSpecSku struct {
	Name string `json:"name" tf:"name"`
}

type RelayNamespaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"kubeFormSecret,omitempty" tf:"-"`

	Location string `json:"location" tf:"location"`
	// +optional
	MetricID string `json:"metricID,omitempty" tf:"metric_id,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryKey        string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	Sku []RelayNamespaceSpecSku `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	SkuName string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RelayNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RelayNamespaceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RelayNamespaceList is a list of RelayNamespaces
type RelayNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RelayNamespace CRD objects
	Items []RelayNamespace `json:"items,omitempty"`
}
