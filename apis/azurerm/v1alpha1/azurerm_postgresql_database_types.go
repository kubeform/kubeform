package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPostgresqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPostgresqlDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermPostgresqlDatabaseStatus `json:"status,omitempty"`
}

type AzurermPostgresqlDatabaseSpec struct {
	Charset           string `json:"charset"`
	Collation         string `json:"collation"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
}

type AzurermPostgresqlDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPostgresqlDatabaseList is a list of AzurermPostgresqlDatabases
type AzurermPostgresqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPostgresqlDatabase CRD objects
	Items []AzurermPostgresqlDatabase `json:"items,omitempty"`
}
