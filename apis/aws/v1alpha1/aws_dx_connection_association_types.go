package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnectionAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxConnectionAssociationSpec   `json:"spec,omitempty"`
	Status            AwsDxConnectionAssociationStatus `json:"status,omitempty"`
}

type AwsDxConnectionAssociationSpec struct {
	LagId        string `json:"lag_id"`
	ConnectionId string `json:"connection_id"`
}



type AwsDxConnectionAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxConnectionAssociationList is a list of AwsDxConnectionAssociations
type AwsDxConnectionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxConnectionAssociation CRD objects
	Items []AwsDxConnectionAssociation `json:"items,omitempty"`
}