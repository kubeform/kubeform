package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	EvaluateTargetHealth bool   `json:"evaluateTargetHealth" tf:"evaluate_target_health"`
	Name                 string `json:"name" tf:"name"`
	ZoneID               string `json:"zoneID" tf:"zone_id"`
}

type Route53RecordSpecFailoverRoutingPolicy struct {
	Type string `json:"type" tf:"type"`
}

type Route53RecordSpecGeolocationRoutingPolicy struct {
	// +optional
	Continent string `json:"continent,omitempty" tf:"continent,omitempty"`
	// +optional
	Country string `json:"country,omitempty" tf:"country,omitempty"`
	// +optional
	Subdivision string `json:"subdivision,omitempty" tf:"subdivision,omitempty"`
}

type Route53RecordSpecLatencyRoutingPolicy struct {
	Region string `json:"region" tf:"region"`
}

type Route53RecordSpecWeightedRoutingPolicy struct {
	Weight int `json:"weight" tf:"weight"`
}

type Route53RecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Alias []Route53RecordSpecAlias `json:"alias,omitempty" tf:"alias,omitempty"`
	// +optional
	AllowOverwrite bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`
	// +optional
	FailoverRoutingPolicy []Route53RecordSpecFailoverRoutingPolicy `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy,omitempty"`
	// +optional
	GeolocationRoutingPolicy []Route53RecordSpecGeolocationRoutingPolicy `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy,omitempty"`
	// +optional
	HealthCheckID string `json:"healthCheckID,omitempty" tf:"health_check_id,omitempty"`
	// +optional
	LatencyRoutingPolicy []Route53RecordSpecLatencyRoutingPolicy `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy,omitempty"`
	// +optional
	MultivalueAnswerRoutingPolicy bool   `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy,omitempty"`
	Name                          string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Records []string `json:"records,omitempty" tf:"records,omitempty"`
	// +optional
	SetIdentifier string `json:"setIdentifier,omitempty" tf:"set_identifier,omitempty"`
	// +optional
	Ttl  int    `json:"ttl,omitempty" tf:"ttl,omitempty"`
	Type string `json:"type" tf:"type"`
	// +optional
	WeightedRoutingPolicy []Route53RecordSpecWeightedRoutingPolicy `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy,omitempty"`
	ZoneID                string                                   `json:"zoneID" tf:"zone_id"`
}

type Route53RecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *Route53RecordSpec `json:"state,omitempty"`
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
