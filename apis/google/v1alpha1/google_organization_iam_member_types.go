package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleOrganizationIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleOrganizationIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleOrganizationIamMemberStatus `json:"status,omitempty"`
}

type GoogleOrganizationIamMemberSpec struct {
	Role   string `json:"role"`
	Member string `json:"member"`
	Etag   string `json:"etag"`
	OrgId  string `json:"org_id"`
}

type GoogleOrganizationIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleOrganizationIamMemberList is a list of GoogleOrganizationIamMembers
type GoogleOrganizationIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleOrganizationIamMember CRD objects
	Items []GoogleOrganizationIamMember `json:"items,omitempty"`
}
