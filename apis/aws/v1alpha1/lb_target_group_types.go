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

type LbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbTargetGroupSpec   `json:"spec,omitempty"`
	Status            LbTargetGroupStatus `json:"status,omitempty"`
}

type LbTargetGroupSpec struct {
	// +optional
	DeregistrationDelay int `json:"deregistration_delay,omitempty"`
	// +optional
	LambdaMultiValueHeadersEnabled bool `json:"lambda_multi_value_headers_enabled,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// +optional
	ProxyProtocolV2 bool `json:"proxy_protocol_v2,omitempty"`
	// +optional
	SlowStart int `json:"slow_start,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TargetType string `json:"target_type,omitempty"`
	// +optional
	VpcId string `json:"vpc_id,omitempty"`
}

type LbTargetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbTargetGroupList is a list of LbTargetGroups
type LbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbTargetGroup CRD objects
	Items []LbTargetGroup `json:"items,omitempty"`
}
