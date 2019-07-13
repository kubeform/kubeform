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

type AwsIamUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserSpec   `json:"spec,omitempty"`
	Status            AwsIamUserStatus `json:"status,omitempty"`
}

type AwsIamUserSpec struct {
	Arn                 string            `json:"arn"`
	UniqueId            string            `json:"unique_id"`
	Name                string            `json:"name"`
	Path                string            `json:"path"`
	PermissionsBoundary string            `json:"permissions_boundary"`
	ForceDestroy        bool              `json:"force_destroy"`
	Tags                map[string]string `json:"tags"`
}



type AwsIamUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamUserList is a list of AwsIamUsers
type AwsIamUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUser CRD objects
	Items []AwsIamUser `json:"items,omitempty"`
}