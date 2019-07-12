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

type AzurermStreamAnalyticsOutputServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsOutputServicebusQueueSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsOutputServicebusQueueStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsOutputServicebusQueueSpecSerialization struct {
	Type           string `json:"type"`
	FieldDelimiter string `json:"field_delimiter"`
	Encoding       string `json:"encoding"`
	Format         string `json:"format"`
}

type AzurermStreamAnalyticsOutputServicebusQueueSpec struct {
	ResourceGroupName      string                                            `json:"resource_group_name"`
	QueueName              string                                            `json:"queue_name"`
	ServicebusNamespace    string                                            `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                            `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                            `json:"shared_access_policy_name"`
	Serialization          []AzurermStreamAnalyticsOutputServicebusQueueSpec `json:"serialization"`
	Name                   string                                            `json:"name"`
	StreamAnalyticsJobName string                                            `json:"stream_analytics_job_name"`
}

type AzurermStreamAnalyticsOutputServicebusQueueStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStreamAnalyticsOutputServicebusQueueList is a list of AzurermStreamAnalyticsOutputServicebusQueues
type AzurermStreamAnalyticsOutputServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsOutputServicebusQueue CRD objects
	Items []AzurermStreamAnalyticsOutputServicebusQueue `json:"items,omitempty"`
}
