package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMysqlConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMysqlConfigurationSpec   `json:"spec,omitempty"`
	Status            AzurermMysqlConfigurationStatus `json:"status,omitempty"`
}

type AzurermMysqlConfigurationSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	Value             string `json:"value"`
}



type AzurermMysqlConfigurationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMysqlConfigurationList is a list of AzurermMysqlConfigurations
type AzurermMysqlConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMysqlConfiguration CRD objects
	Items []AzurermMysqlConfiguration `json:"items,omitempty"`
}