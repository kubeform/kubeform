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

type AwsDbInstanceRoleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbInstanceRoleAssociationSpec   `json:"spec,omitempty"`
	Status            AwsDbInstanceRoleAssociationStatus `json:"status,omitempty"`
}

type AwsDbInstanceRoleAssociationSpec struct {
	DbInstanceIdentifier string `json:"db_instance_identifier"`
	FeatureName          string `json:"feature_name"`
	RoleArn              string `json:"role_arn"`
}



type AwsDbInstanceRoleAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbInstanceRoleAssociationList is a list of AwsDbInstanceRoleAssociations
type AwsDbInstanceRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbInstanceRoleAssociation CRD objects
	Items []AwsDbInstanceRoleAssociation `json:"items,omitempty"`
}