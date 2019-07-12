package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamServiceLinkedRoleSpec   `json:"spec,omitempty"`
	Status            AwsIamServiceLinkedRoleStatus `json:"status,omitempty"`
}

type AwsIamServiceLinkedRoleSpec struct {
	Description    string `json:"description"`
	AwsServiceName string `json:"aws_service_name"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Arn            string `json:"arn"`
	CreateDate     string `json:"create_date"`
	UniqueId       string `json:"unique_id"`
	CustomSuffix   string `json:"custom_suffix"`
}



type AwsIamServiceLinkedRoleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamServiceLinkedRoleList is a list of AwsIamServiceLinkedRoles
type AwsIamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamServiceLinkedRole CRD objects
	Items []AwsIamServiceLinkedRole `json:"items,omitempty"`
}