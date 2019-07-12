package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53ZoneAssociationSpec   `json:"spec,omitempty"`
	Status            AwsRoute53ZoneAssociationStatus `json:"status,omitempty"`
}

type AwsRoute53ZoneAssociationSpec struct {
	ZoneId    string `json:"zone_id"`
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
}



type AwsRoute53ZoneAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneAssociationList is a list of AwsRoute53ZoneAssociations
type AwsRoute53ZoneAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53ZoneAssociation CRD objects
	Items []AwsRoute53ZoneAssociation `json:"items,omitempty"`
}