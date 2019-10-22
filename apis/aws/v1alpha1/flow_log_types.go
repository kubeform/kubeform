package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec,omitempty"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

type FlowLogSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EniID string `json:"eniID,omitempty" tf:"eni_id,omitempty"`
	// +optional
	IamRoleArn string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`
	// +optional
	LogDestination string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`
	// +optional
	LogDestinationType string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`
	// +optional
	// Deprecated
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	SubnetID    string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TrafficType string `json:"trafficType" tf:"traffic_type"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type FlowLogStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FlowLogSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FlowLogList is a list of FlowLogs
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FlowLog CRD objects
	Items []FlowLog `json:"items,omitempty"`
}
