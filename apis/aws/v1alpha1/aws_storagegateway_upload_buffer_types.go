package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayUploadBuffer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayUploadBufferSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayUploadBufferStatus `json:"status,omitempty"`
}

type AwsStoragegatewayUploadBufferSpec struct {
	DiskId     string `json:"disk_id"`
	GatewayArn string `json:"gateway_arn"`
}



type AwsStoragegatewayUploadBufferStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayUploadBufferList is a list of AwsStoragegatewayUploadBuffers
type AwsStoragegatewayUploadBufferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayUploadBuffer CRD objects
	Items []AwsStoragegatewayUploadBuffer `json:"items,omitempty"`
}