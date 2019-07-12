package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRamResourceAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRamResourceAssociationSpec   `json:"spec,omitempty"`
	Status            AwsRamResourceAssociationStatus `json:"status,omitempty"`
}

type AwsRamResourceAssociationSpec struct {
	ResourceArn      string `json:"resource_arn"`
	ResourceShareArn string `json:"resource_share_arn"`
}



type AwsRamResourceAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRamResourceAssociationList is a list of AwsRamResourceAssociations
type AwsRamResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRamResourceAssociation CRD objects
	Items []AwsRamResourceAssociation `json:"items,omitempty"`
}