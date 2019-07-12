package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRouteSpec   `json:"spec,omitempty"`
	Status            AzurermRouteStatus `json:"status,omitempty"`
}

type AzurermRouteSpec struct {
	AddressPrefix      string `json:"address_prefix"`
	NextHopType        string `json:"next_hop_type"`
	NextHopInIpAddress string `json:"next_hop_in_ip_address"`
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
	RouteTableName     string `json:"route_table_name"`
}

type AzurermRouteStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRouteList is a list of AzurermRoutes
type AzurermRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRoute CRD objects
	Items []AzurermRoute `json:"items,omitempty"`
}
