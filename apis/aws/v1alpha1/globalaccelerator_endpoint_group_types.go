package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GlobalacceleratorEndpointGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorEndpointGroupSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorEndpointGroupStatus `json:"status,omitempty"`
}

type GlobalacceleratorEndpointGroupSpecEndpointConfiguration struct {
	// +optional
	EndpointId string `json:"endpoint_id,omitempty"`
	// +optional
	Weight int `json:"weight,omitempty"`
}

type GlobalacceleratorEndpointGroupSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	EndpointConfiguration *[]GlobalacceleratorEndpointGroupSpec `json:"endpoint_configuration,omitempty"`
	// +optional
	HealthCheckIntervalSeconds int `json:"health_check_interval_seconds,omitempty"`
	// +optional
	HealthCheckPath string `json:"health_check_path,omitempty"`
	// +optional
	HealthCheckPort int `json:"health_check_port,omitempty"`
	// +optional
	HealthCheckProtocol string `json:"health_check_protocol,omitempty"`
	ListenerArn         string `json:"listener_arn"`
	// +optional
	ThresholdCount int `json:"threshold_count,omitempty"`
	// +optional
	TrafficDialPercentage json.Number `json:"traffic_dial_percentage,omitempty"`
}

type GlobalacceleratorEndpointGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalacceleratorEndpointGroupList is a list of GlobalacceleratorEndpointGroups
type GlobalacceleratorEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalacceleratorEndpointGroup CRD objects
	Items []GlobalacceleratorEndpointGroup `json:"items,omitempty"`
}
