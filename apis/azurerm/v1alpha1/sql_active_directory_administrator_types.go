package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SqlActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlActiveDirectoryAdministratorSpec   `json:"spec,omitempty"`
	Status            SqlActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

type SqlActiveDirectoryAdministratorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Login             string `json:"login" tf:"login"`
	ObjectID          string `json:"objectID" tf:"object_id"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string `json:"serverName" tf:"server_name"`
	TenantID          string `json:"tenantID" tf:"tenant_id"`
}

type SqlActiveDirectoryAdministratorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlActiveDirectoryAdministratorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlActiveDirectoryAdministratorList is a list of SqlActiveDirectoryAdministrators
type SqlActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlActiveDirectoryAdministrator CRD objects
	Items []SqlActiveDirectoryAdministrator `json:"items,omitempty"`
}
