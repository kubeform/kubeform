package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type EventgridDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridDomainSpec   `json:"spec,omitempty"`
	Status            EventgridDomainStatus `json:"status,omitempty"`
}

type EventgridDomainSpecInputMappingDefaultValues struct {
	// +optional
	DataVersion string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`
	// +optional
	EventType string `json:"eventType,omitempty" tf:"event_type,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type EventgridDomainSpecInputMappingFields struct {
	// +optional
	DataVersion string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`
	// +optional
	EventTime string `json:"eventTime,omitempty" tf:"event_time,omitempty"`
	// +optional
	EventType string `json:"eventType,omitempty" tf:"event_type,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
	// +optional
	Topic string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type EventgridDomainSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingDefaultValues []EventgridDomainSpecInputMappingDefaultValues `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingFields []EventgridDomainSpecInputMappingFields `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`
	// +optional
	InputSchema       string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventgridDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventgridDomainSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridDomainList is a list of EventgridDomains
type EventgridDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventgridDomain CRD objects
	Items []EventgridDomain `json:"items,omitempty"`
}
