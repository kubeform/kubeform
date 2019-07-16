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

type SqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlServerSpec   `json:"spec,omitempty"`
	Status            SqlServerStatus `json:"status,omitempty"`
}

type SqlServerSpec struct {
	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	Location                   string `json:"location"`
	Name                       string `json:"name"`
	ResourceGroupName          string `json:"resource_group_name"`
	Version                    string `json:"version"`
}

type SqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlServerList is a list of SqlServers
type SqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlServer CRD objects
	Items []SqlServer `json:"items,omitempty"`
}
