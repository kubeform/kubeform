package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsWebTestSpec   `json:"spec,omitempty"`
	Status            ApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

type ApplicationInsightsWebTestSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApplicationInsightsID string `json:"applicationInsightsID" tf:"application_insights_id"`
	Configuration         string `json:"configuration" tf:"configuration"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Frequency int `json:"frequency,omitempty" tf:"frequency,omitempty"`
	// +kubebuilder:validation:MinItems=1
	GeoLocations      []string `json:"geoLocations" tf:"geo_locations"`
	Kind              string   `json:"kind" tf:"kind"`
	Location          string   `json:"location" tf:"location"`
	Name              string   `json:"name" tf:"name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RetryEnabled bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`
	// +optional
	SyntheticMonitorID string `json:"syntheticMonitorID,omitempty" tf:"synthetic_monitor_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ApplicationInsightsWebTestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApplicationInsightsWebTestSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApplicationInsightsWebTestList is a list of ApplicationInsightsWebTests
type ApplicationInsightsWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationInsightsWebTest CRD objects
	Items []ApplicationInsightsWebTest `json:"items,omitempty"`
}
