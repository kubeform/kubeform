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

type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseStatus `json:"status,omitempty"`
}

type SqlDatabaseSpecImport struct {
	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	AuthenticationType         string `json:"authentication_type"`
	// +optional
	OperationMode  string `json:"operation_mode,omitempty"`
	StorageKey     string `json:"storage_key"`
	StorageKeyType string `json:"storage_key_type"`
	StorageUri     string `json:"storage_uri"`
}

type SqlDatabaseSpec struct {
	// +optional
	CreateMode string `json:"create_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Import   *[]SqlDatabaseSpec `json:"import,omitempty"`
	Location string             `json:"location"`
	Name     string             `json:"name"`
	// +optional
	ReadScale         bool   `json:"read_scale,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
}

type SqlDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseList is a list of SqlDatabases
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabase CRD objects
	Items []SqlDatabase `json:"items,omitempty"`
}
