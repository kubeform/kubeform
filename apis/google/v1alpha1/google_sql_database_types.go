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

type GoogleSqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlDatabaseSpec   `json:"spec,omitempty"`
	Status            GoogleSqlDatabaseStatus `json:"status,omitempty"`
}

type GoogleSqlDatabaseSpec struct {
	SelfLink  string `json:"self_link"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
	Name      string `json:"name"`
	Instance  string `json:"instance"`
	Project   string `json:"project"`
}

type GoogleSqlDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSqlDatabaseList is a list of GoogleSqlDatabases
type GoogleSqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlDatabase CRD objects
	Items []GoogleSqlDatabase `json:"items,omitempty"`
}
