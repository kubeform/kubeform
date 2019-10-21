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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MacieS3BucketAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MacieS3BucketAssociationSpec   `json:"spec,omitempty"`
	Status            MacieS3BucketAssociationStatus `json:"status,omitempty"`
}

type MacieS3BucketAssociationSpecClassificationType struct {
	// +optional
	Continuous string `json:"continuous,omitempty" tf:"continuous,omitempty"`
	// +optional
	OneTime string `json:"oneTime,omitempty" tf:"one_time,omitempty"`
}

type MacieS3BucketAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClassificationType []MacieS3BucketAssociationSpecClassificationType `json:"classificationType,omitempty" tf:"classification_type,omitempty"`
	// +optional
	MemberAccountID string `json:"memberAccountID,omitempty" tf:"member_account_id,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type MacieS3BucketAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MacieS3BucketAssociationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MacieS3BucketAssociationList is a list of MacieS3BucketAssociations
type MacieS3BucketAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MacieS3BucketAssociation CRD objects
	Items []MacieS3BucketAssociation `json:"items,omitempty"`
}