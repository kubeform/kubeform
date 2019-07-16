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

type FlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec,omitempty"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

type FlowLogSpec struct {
	// +optional
	EniId string `json:"eni_id,omitempty"`
	// +optional
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	// +optional
	LogDestinationType string `json:"log_destination_type,omitempty"`
	// +optional
	SubnetId    string `json:"subnet_id,omitempty"`
	TrafficType string `json:"traffic_type"`
	// +optional
	VpcId string `json:"vpc_id,omitempty"`
}

type FlowLogStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
