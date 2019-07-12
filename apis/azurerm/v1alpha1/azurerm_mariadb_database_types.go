package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMariadbDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMariadbDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermMariadbDatabaseStatus `json:"status,omitempty"`
}

type AzurermMariadbDatabaseSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	Charset           string `json:"charset"`
	Collation         string `json:"collation"`
}



type AzurermMariadbDatabaseStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMariadbDatabaseList is a list of AzurermMariadbDatabases
type AzurermMariadbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMariadbDatabase CRD objects
	Items []AzurermMariadbDatabase `json:"items,omitempty"`
}