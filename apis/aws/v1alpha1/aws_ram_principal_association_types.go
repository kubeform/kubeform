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

type AwsRamPrincipalAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRamPrincipalAssociationSpec   `json:"spec,omitempty"`
	Status            AwsRamPrincipalAssociationStatus `json:"status,omitempty"`
}

type AwsRamPrincipalAssociationSpec struct {
	ResourceShareArn string `json:"resource_share_arn"`
	Principal        string `json:"principal"`
}



type AwsRamPrincipalAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRamPrincipalAssociationList is a list of AwsRamPrincipalAssociations
type AwsRamPrincipalAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRamPrincipalAssociation CRD objects
	Items []AwsRamPrincipalAssociation `json:"items,omitempty"`
}