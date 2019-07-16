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

type LoggingOrganizationSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingOrganizationSinkSpec   `json:"spec,omitempty"`
	Status            LoggingOrganizationSinkStatus `json:"status,omitempty"`
}

type LoggingOrganizationSinkSpec struct {
	Destination string `json:"destination"`
	// +optional
	Filter string `json:"filter,omitempty"`
	// +optional
	IncludeChildren bool   `json:"include_children,omitempty"`
	Name            string `json:"name"`
	OrgId           string `json:"org_id"`
}

type LoggingOrganizationSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingOrganizationSinkList is a list of LoggingOrganizationSinks
type LoggingOrganizationSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingOrganizationSink CRD objects
	Items []LoggingOrganizationSink `json:"items,omitempty"`
}
