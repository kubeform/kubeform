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

type AwsWafregionalWebAclAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalWebAclAssociationSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalWebAclAssociationStatus `json:"status,omitempty"`
}

type AwsWafregionalWebAclAssociationSpec struct {
	ResourceArn string `json:"resource_arn"`
	WebAclId    string `json:"web_acl_id"`
}



type AwsWafregionalWebAclAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalWebAclAssociationList is a list of AwsWafregionalWebAclAssociations
type AwsWafregionalWebAclAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalWebAclAssociation CRD objects
	Items []AwsWafregionalWebAclAssociation `json:"items,omitempty"`
}