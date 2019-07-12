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

type AzurermServicebusNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusNamespaceSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusNamespaceStatus `json:"status,omitempty"`
}

type AzurermServicebusNamespaceSpec struct {
	Location                         string            `json:"location"`
	Capacity                         int               `json:"capacity"`
	DefaultSecondaryConnectionString string            `json:"default_secondary_connection_string"`
	Tags                             map[string]string `json:"tags"`
	DefaultSecondaryKey              string            `json:"default_secondary_key"`
	Name                             string            `json:"name"`
	ResourceGroupName                string            `json:"resource_group_name"`
	Sku                              string            `json:"sku"`
	DefaultPrimaryConnectionString   string            `json:"default_primary_connection_string"`
	DefaultPrimaryKey                string            `json:"default_primary_key"`
}

type AzurermServicebusNamespaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusNamespaceList is a list of AzurermServicebusNamespaces
type AzurermServicebusNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusNamespace CRD objects
	Items []AzurermServicebusNamespace `json:"items,omitempty"`
}
