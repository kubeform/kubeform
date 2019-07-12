package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailStaticIpSpec   `json:"spec,omitempty"`
	Status            AwsLightsailStaticIpStatus `json:"status,omitempty"`
}

type AwsLightsailStaticIpSpec struct {
	IpAddress   string `json:"ip_address"`
	Arn         string `json:"arn"`
	SupportCode string `json:"support_code"`
	Name        string `json:"name"`
}

type AwsLightsailStaticIpStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailStaticIpList is a list of AwsLightsailStaticIps
type AwsLightsailStaticIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailStaticIp CRD objects
	Items []AwsLightsailStaticIp `json:"items,omitempty"`
}
