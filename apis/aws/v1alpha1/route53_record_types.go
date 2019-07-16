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

type Route53Record struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53RecordSpec   `json:"spec,omitempty"`
	Status            Route53RecordStatus `json:"status,omitempty"`
}

type Route53RecordSpecAlias struct {
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
	Name                 string `json:"name"`
	ZoneId               string `json:"zone_id"`
}

type Route53RecordSpecFailoverRoutingPolicy struct {
	Type string `json:"type"`
}

type Route53RecordSpecGeolocationRoutingPolicy struct {
	// +optional
	Continent string `json:"continent,omitempty"`
	// +optional
	Country string `json:"country,omitempty"`
	// +optional
	Subdivision string `json:"subdivision,omitempty"`
}

type Route53RecordSpecLatencyRoutingPolicy struct {
	Region string `json:"region"`
}

type Route53RecordSpecWeightedRoutingPolicy struct {
	Weight int `json:"weight"`
}

type Route53RecordSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Alias *[]Route53RecordSpec `json:"alias,omitempty"`
	// +optional
	FailoverRoutingPolicy *[]Route53RecordSpec `json:"failover_routing_policy,omitempty"`
	// +optional
	GeolocationRoutingPolicy *[]Route53RecordSpec `json:"geolocation_routing_policy,omitempty"`
	// +optional
	HealthCheckId string `json:"health_check_id,omitempty"`
	// +optional
	LatencyRoutingPolicy *[]Route53RecordSpec `json:"latency_routing_policy,omitempty"`
	// +optional
	MultivalueAnswerRoutingPolicy bool   `json:"multivalue_answer_routing_policy,omitempty"`
	Name                          string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Records []string `json:"records,omitempty"`
	// +optional
	SetIdentifier string `json:"set_identifier,omitempty"`
	// +optional
	Ttl  int    `json:"ttl,omitempty"`
	Type string `json:"type"`
	// +optional
	WeightedRoutingPolicy *[]Route53RecordSpec `json:"weighted_routing_policy,omitempty"`
	ZoneId                string               `json:"zone_id"`
}

type Route53RecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53RecordList is a list of Route53Records
type Route53RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53Record CRD objects
	Items []Route53Record `json:"items,omitempty"`
}
