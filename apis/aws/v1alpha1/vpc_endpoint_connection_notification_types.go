package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	ConnectionEvents          []string `json:"connectionEvents" tf:"connection_events"`
	ConnectionNotificationArn string   `json:"connectionNotificationArn" tf:"connection_notification_arn"`
	NotificationType          string   `json:"notificationType" tf:"notification_type"`
	State                     string   `json:"state" tf:"state"`
	// +optional
	VpcEndpointID string `json:"vpcEndpointID,omitempty" tf:"vpc_endpoint_id,omitempty"`
	// +optional
	VpcEndpointServiceID string `json:"vpcEndpointServiceID,omitempty" tf:"vpc_endpoint_service_id,omitempty"`
}

type VpcEndpointConnectionNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcEndpointConnectionNotificationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
