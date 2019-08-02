package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type XraySamplingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              XraySamplingRuleSpec   `json:"spec,omitempty"`
	Status            XraySamplingRuleStatus `json:"status,omitempty"`
}

type XraySamplingRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Attributes    map[string]string `json:"attributes,omitempty" tf:"attributes,omitempty"`
	FixedRate     json.Number       `json:"fixedRate" tf:"fixed_rate"`
	Host          string            `json:"host" tf:"host"`
	HttpMethod    string            `json:"httpMethod" tf:"http_method"`
	Priority      int               `json:"priority" tf:"priority"`
	ReservoirSize int               `json:"reservoirSize" tf:"reservoir_size"`
	ResourceArn   string            `json:"resourceArn" tf:"resource_arn"`
	// +optional
	RuleName    string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`
	ServiceName string `json:"serviceName" tf:"service_name"`
	ServiceType string `json:"serviceType" tf:"service_type"`
	UrlPath     string `json:"urlPath" tf:"url_path"`
	Version     int    `json:"version" tf:"version"`
}

type XraySamplingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *XraySamplingRuleSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// XraySamplingRuleList is a list of XraySamplingRules
type XraySamplingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of XraySamplingRule CRD objects
	Items []XraySamplingRule `json:"items,omitempty"`
}
