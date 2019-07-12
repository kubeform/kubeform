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

type AzurermEventgridDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventgridDomainSpec   `json:"spec,omitempty"`
	Status            AzurermEventgridDomainStatus `json:"status,omitempty"`
}

type AzurermEventgridDomainSpecInputMappingDefaultValues struct {
	EventType   string `json:"event_type"`
	Subject     string `json:"subject"`
	DataVersion string `json:"data_version"`
}

type AzurermEventgridDomainSpecInputMappingFields struct {
	Id          string `json:"id"`
	Topic       string `json:"topic"`
	EventTime   string `json:"event_time"`
	EventType   string `json:"event_type"`
	Subject     string `json:"subject"`
	DataVersion string `json:"data_version"`
}

type AzurermEventgridDomainSpec struct {
	InputMappingDefaultValues []AzurermEventgridDomainSpec `json:"input_mapping_default_values"`
	Endpoint                  string                       `json:"endpoint"`
	Name                      string                       `json:"name"`
	Location                  string                       `json:"location"`
	ResourceGroupName         string                       `json:"resource_group_name"`
	Tags                      map[string]string            `json:"tags"`
	InputSchema               string                       `json:"input_schema"`
	InputMappingFields        []AzurermEventgridDomainSpec `json:"input_mapping_fields"`
}

type AzurermEventgridDomainStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermEventgridDomainList is a list of AzurermEventgridDomains
type AzurermEventgridDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventgridDomain CRD objects
	Items []AzurermEventgridDomain `json:"items,omitempty"`
}
