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
	Name                  string            `json:"name"`
	GeoLocations          []string          `json:"geo_locations"`
	Tags                  map[string]string `json:"tags"`
	Frequency             int               `json:"frequency"`
	Timeout               int               `json:"timeout"`
	Enabled               bool              `json:"enabled"`
	Configuration         string            `json:"configuration"`
	SyntheticMonitorId    string            `json:"synthetic_monitor_id"`
	ResourceGroupName     string            `json:"resource_group_name"`
	Location              string            `json:"location"`
	Kind                  string            `json:"kind"`
	ApplicationInsightsId string            `json:"application_insights_id"`
	RetryEnabled          bool              `json:"retry_enabled"`
	Description           string            `json:"description"`
}

type AzurermApplicationInsightsWebTestStatus struct {
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
