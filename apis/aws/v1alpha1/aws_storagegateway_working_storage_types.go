package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayWorkingStorage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayWorkingStorageSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayWorkingStorageStatus `json:"status,omitempty"`
}

type AwsStoragegatewayWorkingStorageSpec struct {
	DiskId     string `json:"disk_id"`
	GatewayArn string `json:"gateway_arn"`
}

type AwsStoragegatewayWorkingStorageStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayWorkingStorageList is a list of AwsStoragegatewayWorkingStorages
type AwsStoragegatewayWorkingStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayWorkingStorage CRD objects
	Items []AwsStoragegatewayWorkingStorage `json:"items,omitempty"`
}
