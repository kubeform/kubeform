package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DbInstanceRoleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceRoleAssociationSpec   `json:"spec,omitempty"`
	Status            DbInstanceRoleAssociationStatus `json:"status,omitempty"`
}

type DbInstanceRoleAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DbInstanceIdentifier string `json:"dbInstanceIdentifier" tf:"db_instance_identifier"`
	FeatureName          string `json:"featureName" tf:"feature_name"`
	RoleArn              string `json:"roleArn" tf:"role_arn"`
}



type DbInstanceRoleAssociationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *DbInstanceRoleAssociationSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbInstanceRoleAssociationList is a list of DbInstanceRoleAssociations
type DbInstanceRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbInstanceRoleAssociation CRD objects
	Items []DbInstanceRoleAssociation `json:"items,omitempty"`
}