package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermIothubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermIothubSharedAccessPolicySpec   `json:"spec,omitempty"`
	Status            AzurermIothubSharedAccessPolicyStatus `json:"status,omitempty"`
}

type AzurermIothubSharedAccessPolicySpec struct {
	Name                      string `json:"name"`
	ResourceGroupName         string `json:"resource_group_name"`
	IothubName                string `json:"iothub_name"`
	RegistryRead              bool   `json:"registry_read"`
	RegistryWrite             bool   `json:"registry_write"`
	ServiceConnect            bool   `json:"service_connect"`
	DeviceConnect             bool   `json:"device_connect"`
	PrimaryKey                string `json:"primary_key"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	SecondaryKey              string `json:"secondary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
}



type AzurermIothubSharedAccessPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermIothubSharedAccessPolicyList is a list of AzurermIothubSharedAccessPolicys
type AzurermIothubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermIothubSharedAccessPolicy CRD objects
	Items []AzurermIothubSharedAccessPolicy `json:"items,omitempty"`
}