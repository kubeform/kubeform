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

type AzurermMysqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMysqlDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermMysqlDatabaseStatus `json:"status,omitempty"`
}

type AzurermMysqlDatabaseSpec struct {
	ServerName        string `json:"server_name"`
	Charset           string `json:"charset"`
	Collation         string `json:"collation"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type AzurermMysqlDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMysqlDatabaseList is a list of AzurermMysqlDatabases
type AzurermMysqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMysqlDatabase CRD objects
	Items []AzurermMysqlDatabase `json:"items,omitempty"`
}
