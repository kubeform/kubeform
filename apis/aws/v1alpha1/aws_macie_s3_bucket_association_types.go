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

type AwsMacieS3BucketAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMacieS3BucketAssociationSpec   `json:"spec,omitempty"`
	Status            AwsMacieS3BucketAssociationStatus `json:"status,omitempty"`
}

type AwsMacieS3BucketAssociationSpecClassificationType struct {
	Continuous string `json:"continuous"`
	OneTime    string `json:"one_time"`
}

type AwsMacieS3BucketAssociationSpec struct {
	ClassificationType []AwsMacieS3BucketAssociationSpec `json:"classification_type"`
	BucketName         string                            `json:"bucket_name"`
	Prefix             string                            `json:"prefix"`
	MemberAccountId    string                            `json:"member_account_id"`
}

type AwsMacieS3BucketAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMacieS3BucketAssociationList is a list of AwsMacieS3BucketAssociations
type AwsMacieS3BucketAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMacieS3BucketAssociation CRD objects
	Items []AwsMacieS3BucketAssociation `json:"items,omitempty"`
}
