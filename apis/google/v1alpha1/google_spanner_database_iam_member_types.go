package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerDatabaseIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerDatabaseIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerDatabaseIamMemberStatus `json:"status,omitempty"`
}

type GoogleSpannerDatabaseIamMemberSpec struct {
	Role     string `json:"role"`
	Member   string `json:"member"`
	Etag     string `json:"etag"`
	Instance string `json:"instance"`
	Database string `json:"database"`
	Project  string `json:"project"`
}



type GoogleSpannerDatabaseIamMemberStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerDatabaseIamMemberList is a list of GoogleSpannerDatabaseIamMembers
type GoogleSpannerDatabaseIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerDatabaseIamMember CRD objects
	Items []GoogleSpannerDatabaseIamMember `json:"items,omitempty"`
}