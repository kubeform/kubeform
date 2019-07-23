package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	EndpointID string `json:"endpointID,omitempty" tf:"endpoint_id,omitempty"`
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}

type GlobalacceleratorEndpointGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	EndpointConfiguration []GlobalacceleratorEndpointGroupSpecEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`
	// +optional
	EndpointGroupRegion string `json:"endpointGroupRegion,omitempty" tf:"endpoint_group_region,omitempty"`
	// +optional
	HealthCheckIntervalSeconds int `json:"healthCheckIntervalSeconds,omitempty" tf:"health_check_interval_seconds,omitempty"`
	// +optional
	HealthCheckPath string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`
	// +optional
	HealthCheckPort int `json:"healthCheckPort,omitempty" tf:"health_check_port,omitempty"`
	// +optional
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`
	ListenerArn         string `json:"listenerArn" tf:"listener_arn"`
	// +optional
	ThresholdCount int `json:"thresholdCount,omitempty" tf:"threshold_count,omitempty"`
	// +optional
	TrafficDialPercentage json.Number `json:"trafficDialPercentage,omitempty" tf:"traffic_dial_percentage,omitempty"`
}

type GlobalacceleratorEndpointGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
