package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerInstanceIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerInstanceIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerInstanceIamMemberStatus `json:"status,omitempty"`
}

type GoogleSpannerInstanceIamMemberSpec struct {
	Instance string `json:"instance"`
	Project  string `json:"project"`
	Member   string `json:"member"`
	Etag     string `json:"etag"`
	Role     string `json:"role"`
}



type GoogleSpannerInstanceIamMemberStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerInstanceIamMemberList is a list of GoogleSpannerInstanceIamMembers
type GoogleSpannerInstanceIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerInstanceIamMember CRD objects
	Items []GoogleSpannerInstanceIamMember `json:"items,omitempty"`
}