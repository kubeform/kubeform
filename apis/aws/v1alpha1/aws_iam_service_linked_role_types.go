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

type AwsIamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamServiceLinkedRoleSpec   `json:"spec,omitempty"`
	Status            AwsIamServiceLinkedRoleStatus `json:"status,omitempty"`
}

type AwsIamServiceLinkedRoleSpec struct {
	Path           string `json:"path"`
	Arn            string `json:"arn"`
	CreateDate     string `json:"create_date"`
	UniqueId       string `json:"unique_id"`
	CustomSuffix   string `json:"custom_suffix"`
	Description    string `json:"description"`
	AwsServiceName string `json:"aws_service_name"`
	Name           string `json:"name"`
}



type AwsIamServiceLinkedRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamServiceLinkedRoleList is a list of AwsIamServiceLinkedRoles
type AwsIamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamServiceLinkedRole CRD objects
	Items []AwsIamServiceLinkedRole `json:"items,omitempty"`
}