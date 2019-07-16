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

type ConfigConfigurationAggregator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationAggregatorSpec   `json:"spec,omitempty"`
	Status            ConfigConfigurationAggregatorStatus `json:"status,omitempty"`
}

type ConfigConfigurationAggregatorSpecAccountAggregationSource struct {
	// +kubebuilder:validation:MinItems=1
	AccountIds []string `json:"account_ids"`
	// +optional
	AllRegions bool `json:"all_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Regions []string `json:"regions,omitempty"`
}

type ConfigConfigurationAggregatorSpecOrganizationAggregationSource struct {
	// +optional
	AllRegions bool `json:"all_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Regions []string `json:"regions,omitempty"`
	RoleArn string   `json:"role_arn"`
}

type ConfigConfigurationAggregatorSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccountAggregationSource *[]ConfigConfigurationAggregatorSpec `json:"account_aggregation_source,omitempty"`
	Name                     string                               `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OrganizationAggregationSource *[]ConfigConfigurationAggregatorSpec `json:"organization_aggregation_source,omitempty"`
}

type ConfigConfigurationAggregatorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationAggregatorList is a list of ConfigConfigurationAggregators
type ConfigConfigurationAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationAggregator CRD objects
	Items []ConfigConfigurationAggregator `json:"items,omitempty"`
}
