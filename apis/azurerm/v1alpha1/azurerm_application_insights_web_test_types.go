package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationInsightsWebTestSpec   `json:"spec,omitempty"`
	Status            AzurermApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

type AzurermApplicationInsightsWebTestSpec struct {
	Name                  string            `json:"name"`
	Description           string            `json:"description"`
	Tags                  map[string]string `json:"tags"`
	Location              string            `json:"location"`
	Frequency             int               `json:"frequency"`
	Timeout               int               `json:"timeout"`
	Enabled               bool              `json:"enabled"`
	ResourceGroupName     string            `json:"resource_group_name"`
	ApplicationInsightsId string            `json:"application_insights_id"`
	Kind                  string            `json:"kind"`
	GeoLocations          []string          `json:"geo_locations"`
	Configuration         string            `json:"configuration"`
	RetryEnabled          bool              `json:"retry_enabled"`
	SyntheticMonitorId    string            `json:"synthetic_monitor_id"`
}

type AzurermApplicationInsightsWebTestStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApplicationInsightsWebTestList is a list of AzurermApplicationInsightsWebTests
type AzurermApplicationInsightsWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationInsightsWebTest CRD objects
	Items []AzurermApplicationInsightsWebTest `json:"items,omitempty"`
}
