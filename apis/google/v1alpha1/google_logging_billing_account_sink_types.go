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

type GoogleLoggingBillingAccountSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingBillingAccountSinkSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingBillingAccountSinkStatus `json:"status,omitempty"`
}

type GoogleLoggingBillingAccountSinkSpec struct {
	Destination    string `json:"destination"`
	Filter         string `json:"filter"`
	WriterIdentity string `json:"writer_identity"`
	BillingAccount string `json:"billing_account"`
	Name           string `json:"name"`
}

type GoogleLoggingBillingAccountSinkStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingBillingAccountSinkList is a list of GoogleLoggingBillingAccountSinks
type GoogleLoggingBillingAccountSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingBillingAccountSink CRD objects
	Items []GoogleLoggingBillingAccountSink `json:"items,omitempty"`
}
