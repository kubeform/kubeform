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

type AwsGlobalacceleratorEndpointGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlobalacceleratorEndpointGroupSpec   `json:"spec,omitempty"`
	Status            AwsGlobalacceleratorEndpointGroupStatus `json:"status,omitempty"`
}

type AwsGlobalacceleratorEndpointGroupSpecEndpointConfiguration struct {
	EndpointId string `json:"endpoint_id"`
	Weight     int    `json:"weight"`
}

type AwsGlobalacceleratorEndpointGroupSpec struct {
	HealthCheckProtocol        string                                  `json:"health_check_protocol"`
	ThresholdCount             int                                     `json:"threshold_count"`
	EndpointConfiguration      []AwsGlobalacceleratorEndpointGroupSpec `json:"endpoint_configuration"`
	EndpointGroupRegion        string                                  `json:"endpoint_group_region"`
	HealthCheckPort            int                                     `json:"health_check_port"`
	HealthCheckPath            string                                  `json:"health_check_path"`
	TrafficDialPercentage      float64                                 `json:"traffic_dial_percentage"`
	ListenerArn                string                                  `json:"listener_arn"`
	HealthCheckIntervalSeconds int                                     `json:"health_check_interval_seconds"`
}

type AwsGlobalacceleratorEndpointGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlobalacceleratorEndpointGroupList is a list of AwsGlobalacceleratorEndpointGroups
type AwsGlobalacceleratorEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlobalacceleratorEndpointGroup CRD objects
	Items []AwsGlobalacceleratorEndpointGroup `json:"items,omitempty"`
}
