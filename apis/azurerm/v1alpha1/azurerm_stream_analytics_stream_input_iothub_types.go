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

type AzurermStreamAnalyticsStreamInputIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsStreamInputIothubSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsStreamInputIothubStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsStreamInputIothubSpecSerialization struct {
	Type           string `json:"type"`
	FieldDelimiter string `json:"field_delimiter"`
	Encoding       string `json:"encoding"`
}

type AzurermStreamAnalyticsStreamInputIothubSpec struct {
	IothubNamespace           string                                        `json:"iothub_namespace"`
	Serialization             []AzurermStreamAnalyticsStreamInputIothubSpec `json:"serialization"`
	StreamAnalyticsJobName    string                                        `json:"stream_analytics_job_name"`
	ResourceGroupName         string                                        `json:"resource_group_name"`
	Endpoint                  string                                        `json:"endpoint"`
	EventhubConsumerGroupName string                                        `json:"eventhub_consumer_group_name"`
	SharedAccessPolicyKey     string                                        `json:"shared_access_policy_key"`
	SharedAccessPolicyName    string                                        `json:"shared_access_policy_name"`
	Name                      string                                        `json:"name"`
}



type AzurermStreamAnalyticsStreamInputIothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStreamAnalyticsStreamInputIothubList is a list of AzurermStreamAnalyticsStreamInputIothubs
type AzurermStreamAnalyticsStreamInputIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsStreamInputIothub CRD objects
	Items []AzurermStreamAnalyticsStreamInputIothub `json:"items,omitempty"`
}