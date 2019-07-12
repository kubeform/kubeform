package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Sku                       []AzurermRelayNamespaceSpec `json:"sku"`
	PrimaryConnectionString   string                      `json:"primary_connection_string"`
	SecondaryKey              string                      `json:"secondary_key"`
	SecondaryConnectionString string                      `json:"secondary_connection_string"`
	PrimaryKey                string                      `json:"primary_key"`
	Tags                      map[string]string           `json:"tags"`
	Location                  string                      `json:"location"`
	ResourceGroupName         string                      `json:"resource_group_name"`
	SkuName                   string                      `json:"sku_name"`
	MetricId                  string                      `json:"metric_id"`
}

type AzurermRelayNamespaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRelayNamespaceList is a list of AzurermRelayNamespaces
type AzurermRelayNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRelayNamespace CRD objects
	Items []AzurermRelayNamespace `json:"items,omitempty"`
}
