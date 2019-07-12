package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Record struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53RecordSpec   `json:"spec,omitempty"`
	Status            AwsRoute53RecordStatus `json:"status,omitempty"`
}

type AwsRoute53RecordSpecAlias struct {
	ZoneId               string `json:"zone_id"`
	Name                 string `json:"name"`
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
}

type AwsRoute53RecordSpecLatencyRoutingPolicy struct {
	Region string `json:"region"`
}

type AwsRoute53RecordSpecGeolocationRoutingPolicy struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

type AwsRoute53RecordSpecWeightedRoutingPolicy struct {
	Weight int `json:"weight"`
}

type AwsRoute53RecordSpecFailoverRoutingPolicy struct {
	Type string `json:"type"`
}

type AwsRoute53RecordSpec struct {
	Alias                         []AwsRoute53RecordSpec `json:"alias"`
	LatencyRoutingPolicy          []AwsRoute53RecordSpec `json:"latency_routing_policy"`
	GeolocationRoutingPolicy      []AwsRoute53RecordSpec `json:"geolocation_routing_policy"`
	Records                       []string               `json:"records"`
	SetIdentifier                 string                 `json:"set_identifier"`
	WeightedRoutingPolicy         []AwsRoute53RecordSpec `json:"weighted_routing_policy"`
	Name                          string                 `json:"name"`
	MultivalueAnswerRoutingPolicy bool                   `json:"multivalue_answer_routing_policy"`
	HealthCheckId                 string                 `json:"health_check_id"`
	ZoneId                        string                 `json:"zone_id"`
	Type                          string                 `json:"type"`
	Ttl                           int                    `json:"ttl"`
	FailoverRoutingPolicy         []AwsRoute53RecordSpec `json:"failover_routing_policy"`
	AllowOverwrite                bool                   `json:"allow_overwrite"`
	Fqdn                          string                 `json:"fqdn"`
}

type AwsRoute53RecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Records
type AwsRoute53RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53Record CRD objects
	Items []AwsRoute53Record `json:"items,omitempty"`
}
