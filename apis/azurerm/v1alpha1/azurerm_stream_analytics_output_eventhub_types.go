package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStreamAnalyticsOutputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsOutputEventhubSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsOutputEventhubStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsOutputEventhubSpecSerialization struct {
	Type           string `json:"type"`
	FieldDelimiter string `json:"field_delimiter"`
	Encoding       string `json:"encoding"`
	Format         string `json:"format"`
}

type AzurermStreamAnalyticsOutputEventhubSpec struct {
	ServicebusNamespace    string                                     `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                     `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                     `json:"shared_access_policy_name"`
	Serialization          []AzurermStreamAnalyticsOutputEventhubSpec `json:"serialization"`
	Name                   string                                     `json:"name"`
	StreamAnalyticsJobName string                                     `json:"stream_analytics_job_name"`
	ResourceGroupName      string                                     `json:"resource_group_name"`
	EventhubName           string                                     `json:"eventhub_name"`
}

type AzurermStreamAnalyticsOutputEventhubStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStreamAnalyticsOutputEventhubList is a list of AzurermStreamAnalyticsOutputEventhubs
type AzurermStreamAnalyticsOutputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsOutputEventhub CRD objects
	Items []AzurermStreamAnalyticsOutputEventhub `json:"items,omitempty"`
}
