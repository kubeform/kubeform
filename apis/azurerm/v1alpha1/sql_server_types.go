package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	AdministratorLogin string `json:"administratorLogin" tf:"administrator_login"`
	Location           string `json:"location" tf:"location"`
	Name               string `json:"name" tf:"name"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags    map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Version string            `json:"version" tf:"version"`
}

type SqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
