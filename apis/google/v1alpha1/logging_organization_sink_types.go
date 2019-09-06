package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Destination string `json:"destination" tf:"destination"`
	// +optional
	Filter string `json:"filter,omitempty" tf:"filter,omitempty"`
	// +optional
	IncludeChildren bool   `json:"includeChildren,omitempty" tf:"include_children,omitempty"`
	Name            string `json:"name" tf:"name"`
	OrgID           string `json:"orgID" tf:"org_id"`
	// +optional
	WriterIdentity string `json:"writerIdentity,omitempty" tf:"writer_identity,omitempty"`
}

type LoggingOrganizationSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LoggingOrganizationSinkSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
