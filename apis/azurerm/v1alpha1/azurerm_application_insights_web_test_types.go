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

type AzurermApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationInsightsWebTestSpec   `json:"spec,omitempty"`
	Status            AzurermApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

type AzurermApplicationInsightsWebTestSpec struct {
	Configuration         string            `json:"configuration"`
	Tags                  map[string]string `json:"tags"`
	Timeout               int               `json:"timeout"`
	Enabled               bool              `json:"enabled"`
	GeoLocations          []string          `json:"geo_locations"`
	Description           string            `json:"description"`
	ResourceGroupName     string            `json:"resource_group_name"`
	ApplicationInsightsId string            `json:"application_insights_id"`
	Kind                  string            `json:"kind"`
	RetryEnabled          bool              `json:"retry_enabled"`
	SyntheticMonitorId    string            `json:"synthetic_monitor_id"`
	Name                  string            `json:"name"`
	Location              string            `json:"location"`
	Frequency             int               `json:"frequency"`
}



type AzurermApplicationInsightsWebTestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApplicationInsightsWebTestList is a list of AzurermApplicationInsightsWebTests
type AzurermApplicationInsightsWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationInsightsWebTest CRD objects
	Items []AzurermApplicationInsightsWebTest `json:"items,omitempty"`
}