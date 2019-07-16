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

type DxConnectionAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxConnectionAssociationSpec   `json:"spec,omitempty"`
	Status            DxConnectionAssociationStatus `json:"status,omitempty"`
}

type DxConnectionAssociationSpec struct {
	ConnectionId string `json:"connection_id"`
	LagId        string `json:"lag_id"`
}

type DxConnectionAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxConnectionAssociationList is a list of DxConnectionAssociations
type DxConnectionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxConnectionAssociation CRD objects
	Items []DxConnectionAssociation `json:"items,omitempty"`
}
