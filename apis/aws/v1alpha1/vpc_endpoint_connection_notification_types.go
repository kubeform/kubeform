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

type VpcEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointConnectionNotificationSpec   `json:"spec,omitempty"`
	Status            VpcEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

type VpcEndpointConnectionNotificationSpec struct {
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	ConnectionEvents          []string `json:"connection_events"`
	ConnectionNotificationArn string   `json:"connection_notification_arn"`
	// +optional
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// +optional
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id,omitempty"`
}

type VpcEndpointConnectionNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointConnectionNotificationList is a list of VpcEndpointConnectionNotifications
type VpcEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointConnectionNotification CRD objects
	Items []VpcEndpointConnectionNotification `json:"items,omitempty"`
}
