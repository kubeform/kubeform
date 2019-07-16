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

type LoggingBillingAccountSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingBillingAccountSinkSpec   `json:"spec,omitempty"`
	Status            LoggingBillingAccountSinkStatus `json:"status,omitempty"`
}

type LoggingBillingAccountSinkSpec struct {
	BillingAccount string `json:"billing_account"`
	Destination    string `json:"destination"`
	// +optional
	Filter string `json:"filter,omitempty"`
	Name   string `json:"name"`
}

type LoggingBillingAccountSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingBillingAccountSinkList is a list of LoggingBillingAccountSinks
type LoggingBillingAccountSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingBillingAccountSink CRD objects
	Items []LoggingBillingAccountSink `json:"items,omitempty"`
}
