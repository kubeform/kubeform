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

type GoogleLoggingOrganizationSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingOrganizationSinkSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingOrganizationSinkStatus `json:"status,omitempty"`
}

type GoogleLoggingOrganizationSinkSpec struct {
	Destination     string `json:"destination"`
	Filter          string `json:"filter"`
	WriterIdentity  string `json:"writer_identity"`
	OrgId           string `json:"org_id"`
	IncludeChildren bool   `json:"include_children"`
	Name            string `json:"name"`
}

type GoogleLoggingOrganizationSinkStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingOrganizationSinkList is a list of GoogleLoggingOrganizationSinks
type GoogleLoggingOrganizationSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingOrganizationSink CRD objects
	Items []GoogleLoggingOrganizationSink `json:"items,omitempty"`
}
