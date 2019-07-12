package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleOrganizationIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleOrganizationIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleOrganizationIamBindingStatus `json:"status,omitempty"`
}

type GoogleOrganizationIamBindingSpec struct {
	Etag    string   `json:"etag"`
	Role    string   `json:"role"`
	Members []string `json:"members"`
	OrgId   string   `json:"org_id"`
}



type GoogleOrganizationIamBindingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleOrganizationIamBindingList is a list of GoogleOrganizationIamBindings
type GoogleOrganizationIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleOrganizationIamBinding CRD objects
	Items []GoogleOrganizationIamBinding `json:"items,omitempty"`
}