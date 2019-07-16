package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	// +optional
	Attributes    map[string]string `json:"attributes,omitempty"`
	FixedRate     json.Number       `json:"fixed_rate"`
	Host          string            `json:"host"`
	HttpMethod    string            `json:"http_method"`
	Priority      int               `json:"priority"`
	ReservoirSize int               `json:"reservoir_size"`
	ResourceArn   string            `json:"resource_arn"`
	// +optional
	RuleName    string `json:"rule_name,omitempty"`
	ServiceName string `json:"service_name"`
	ServiceType string `json:"service_type"`
	UrlPath     string `json:"url_path"`
	Version     int    `json:"version"`
}

type XraySamplingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
