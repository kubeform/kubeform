package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayCacheSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayCacheStatus `json:"status,omitempty"`
}

type AwsStoragegatewayCacheSpec struct {
	DiskId     string `json:"disk_id"`
	GatewayArn string `json:"gateway_arn"`
}

type AwsStoragegatewayCacheStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayCacheList is a list of AwsStoragegatewayCaches
type AwsStoragegatewayCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayCache CRD objects
	Items []AwsStoragegatewayCache `json:"items,omitempty"`
}
