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

type AwsGlobalacceleratorAccelerator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlobalacceleratorAcceleratorSpec   `json:"spec,omitempty"`
	Status            AwsGlobalacceleratorAcceleratorStatus `json:"status,omitempty"`
}

type AwsGlobalacceleratorAcceleratorSpecIpSets struct {
	IpFamily    string   `json:"ip_family"`
	IpAddresses []string `json:"ip_addresses"`
}

type AwsGlobalacceleratorAcceleratorSpecAttributes struct {
	FlowLogsEnabled  bool   `json:"flow_logs_enabled"`
	FlowLogsS3Bucket string `json:"flow_logs_s3_bucket"`
	FlowLogsS3Prefix string `json:"flow_logs_s3_prefix"`
}

type AwsGlobalacceleratorAcceleratorSpec struct {
	Name          string                                `json:"name"`
	IpAddressType string                                `json:"ip_address_type"`
	Enabled       bool                                  `json:"enabled"`
	IpSets        []AwsGlobalacceleratorAcceleratorSpec `json:"ip_sets"`
	Attributes    []AwsGlobalacceleratorAcceleratorSpec `json:"attributes"`
}

type AwsGlobalacceleratorAcceleratorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlobalacceleratorAcceleratorList is a list of AwsGlobalacceleratorAccelerators
type AwsGlobalacceleratorAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlobalacceleratorAccelerator CRD objects
	Items []AwsGlobalacceleratorAccelerator `json:"items,omitempty"`
}
