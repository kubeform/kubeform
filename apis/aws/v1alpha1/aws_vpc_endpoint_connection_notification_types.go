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

type AwsVpcEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointConnectionNotificationSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

type AwsVpcEndpointConnectionNotificationSpec struct {
	ConnectionNotificationArn string   `json:"connection_notification_arn"`
	ConnectionEvents          []string `json:"connection_events"`
	State                     string   `json:"state"`
	NotificationType          string   `json:"notification_type"`
	VpcEndpointServiceId      string   `json:"vpc_endpoint_service_id"`
	VpcEndpointId             string   `json:"vpc_endpoint_id"`
}

type AwsVpcEndpointConnectionNotificationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcEndpointConnectionNotificationList is a list of AwsVpcEndpointConnectionNotifications
type AwsVpcEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointConnectionNotification CRD objects
	Items []AwsVpcEndpointConnectionNotification `json:"items,omitempty"`
}
