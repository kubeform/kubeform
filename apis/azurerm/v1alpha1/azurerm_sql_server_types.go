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

type AzurermSqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSqlServerSpec   `json:"spec,omitempty"`
	Status            AzurermSqlServerStatus `json:"status,omitempty"`
}

type AzurermSqlServerSpec struct {
	AdministratorLogin         string            `json:"administrator_login"`
	AdministratorLoginPassword string            `json:"administrator_login_password"`
	FullyQualifiedDomainName   string            `json:"fully_qualified_domain_name"`
	Tags                       map[string]string `json:"tags"`
	Name                       string            `json:"name"`
	Location                   string            `json:"location"`
	ResourceGroupName          string            `json:"resource_group_name"`
	Version                    string            `json:"version"`
}

type AzurermSqlServerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSqlServerList is a list of AzurermSqlServers
type AzurermSqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSqlServer CRD objects
	Items []AzurermSqlServer `json:"items,omitempty"`
}
