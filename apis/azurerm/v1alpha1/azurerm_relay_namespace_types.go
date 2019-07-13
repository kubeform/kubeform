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

type AzurermRelayNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRelayNamespaceSpec   `json:"spec,omitempty"`
	Status            AzurermRelayNamespaceStatus `json:"status,omitempty"`
}

type AzurermRelayNamespaceSpecSku struct {
	Name string `json:"name"`
}

type AzurermRelayNamespaceSpec struct {
	Name                      string                      `json:"name"`
	SkuName                   string                      `json:"sku_name"`
	PrimaryConnectionString   string                      `json:"primary_connection_string"`
	Tags                      map[string]string           `json:"tags"`
	Location                  string                      `json:"location"`
	ResourceGroupName         string                      `json:"resource_group_name"`
	Sku                       []AzurermRelayNamespaceSpec `json:"sku"`
	MetricId                  string                      `json:"metric_id"`
	SecondaryConnectionString string                      `json:"secondary_connection_string"`
	PrimaryKey                string                      `json:"primary_key"`
	SecondaryKey              string                      `json:"secondary_key"`
}



type AzurermRelayNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermRelayNamespaceList is a list of AzurermRelayNamespaces
type AzurermRelayNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRelayNamespace CRD objects
	Items []AzurermRelayNamespace `json:"items,omitempty"`
}