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

type EventgridDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridDomainSpec   `json:"spec,omitempty"`
	Status            EventgridDomainStatus `json:"status,omitempty"`
}

type EventgridDomainSpecInputMappingDefaultValues struct {
	// +optional
	DataVersion string `json:"data_version,omitempty"`
	// +optional
	EventType string `json:"event_type,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty"`
}

type EventgridDomainSpecInputMappingFields struct {
	// +optional
	DataVersion string `json:"data_version,omitempty"`
	// +optional
	EventTime string `json:"event_time,omitempty"`
	// +optional
	EventType string `json:"event_type,omitempty"`
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty"`
	// +optional
	Topic string `json:"topic,omitempty"`
}

type EventgridDomainSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingDefaultValues *[]EventgridDomainSpec `json:"input_mapping_default_values,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingFields *[]EventgridDomainSpec `json:"input_mapping_fields,omitempty"`
	// +optional
	InputSchema       string `json:"input_schema,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type EventgridDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
