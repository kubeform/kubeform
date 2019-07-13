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

type AwsIamInstanceProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamInstanceProfileSpec   `json:"spec,omitempty"`
	Status            AwsIamInstanceProfileStatus `json:"status,omitempty"`
}

type AwsIamInstanceProfileSpec struct {
	Roles      []string `json:"roles"`
	Role       string   `json:"role"`
	Arn        string   `json:"arn"`
	CreateDate string   `json:"create_date"`
	UniqueId   string   `json:"unique_id"`
	Name       string   `json:"name"`
	NamePrefix string   `json:"name_prefix"`
	Path       string   `json:"path"`
}



type AwsIamInstanceProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamInstanceProfileList is a list of AwsIamInstanceProfiles
type AwsIamInstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamInstanceProfile CRD objects
	Items []AwsIamInstanceProfile `json:"items,omitempty"`
}