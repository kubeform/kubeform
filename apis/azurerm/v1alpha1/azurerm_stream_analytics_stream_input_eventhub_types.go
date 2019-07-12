package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStreamAnalyticsStreamInputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsStreamInputEventhubSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsStreamInputEventhubStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsStreamInputEventhubSpecSerialization struct {
	Type           string `json:"type"`
	FieldDelimiter string `json:"field_delimiter"`
	Encoding       string `json:"encoding"`
}

type AzurermStreamAnalyticsStreamInputEventhubSpec struct {
	StreamAnalyticsJobName    string                                          `json:"stream_analytics_job_name"`
	ResourceGroupName         string                                          `json:"resource_group_name"`
	EventhubName              string                                          `json:"eventhub_name"`
	SharedAccessPolicyKey     string                                          `json:"shared_access_policy_key"`
	Serialization             []AzurermStreamAnalyticsStreamInputEventhubSpec `json:"serialization"`
	Name                      string                                          `json:"name"`
	EventhubConsumerGroupName string                                          `json:"eventhub_consumer_group_name"`
	ServicebusNamespace       string                                          `json:"servicebus_namespace"`
	SharedAccessPolicyName    string                                          `json:"shared_access_policy_name"`
}

type AzurermStreamAnalyticsStreamInputEventhubStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStreamAnalyticsStreamInputEventhubList is a list of AzurermStreamAnalyticsStreamInputEventhubs
type AzurermStreamAnalyticsStreamInputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsStreamInputEventhub CRD objects
	Items []AzurermStreamAnalyticsStreamInputEventhub `json:"items,omitempty"`
}
