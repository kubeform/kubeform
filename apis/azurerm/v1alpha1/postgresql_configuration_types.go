package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PostgresqlConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlConfigurationSpec   `json:"spec,omitempty"`
	Status            PostgresqlConfigurationStatus `json:"status,omitempty"`
}

type PostgresqlConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string `json:"serverName" tf:"server_name"`
	Value             string `json:"value" tf:"value"`
}

type PostgresqlConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PostgresqlConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PostgresqlConfigurationList is a list of PostgresqlConfigurations
type PostgresqlConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresqlConfiguration CRD objects
	Items []PostgresqlConfiguration `json:"items,omitempty"`
}
