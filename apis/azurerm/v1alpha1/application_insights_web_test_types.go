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

type ApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsWebTestSpec   `json:"spec,omitempty"`
	Status            ApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

type ApplicationInsightsWebTestSpec struct {
	ApplicationInsightsId string `json:"application_insights_id"`
	Configuration         string `json:"configuration"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Frequency int `json:"frequency,omitempty"`
	// +kubebuilder:validation:MinItems=1
	GeoLocations      []string `json:"geo_locations"`
	Kind              string   `json:"kind"`
	Location          string   `json:"location"`
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	// +optional
	RetryEnabled bool `json:"retry_enabled,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty"`
}

type ApplicationInsightsWebTestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
