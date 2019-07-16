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

type SpannerDatabaseIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerDatabaseIamMemberSpec   `json:"spec,omitempty"`
	Status            SpannerDatabaseIamMemberStatus `json:"status,omitempty"`
}

type SpannerDatabaseIamMemberSpec struct {
	Database string `json:"database"`
	Instance string `json:"instance"`
	Member   string `json:"member"`
	Role     string `json:"role"`
}

type SpannerDatabaseIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerDatabaseIamMemberList is a list of SpannerDatabaseIamMembers
type SpannerDatabaseIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerDatabaseIamMember CRD objects
	Items []SpannerDatabaseIamMember `json:"items,omitempty"`
}
