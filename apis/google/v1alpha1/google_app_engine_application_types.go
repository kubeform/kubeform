package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleAppEngineApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleAppEngineApplicationSpec   `json:"spec,omitempty"`
	Status            GoogleAppEngineApplicationStatus `json:"status,omitempty"`
}

type GoogleAppEngineApplicationSpecFeatureSettings struct {
	SplitHealthChecks bool `json:"split_health_checks"`
}

type GoogleAppEngineApplicationSpecUrlDispatchRule struct {
	Domain  string `json:"domain"`
	Path    string `json:"path"`
	Service string `json:"service"`
}

type GoogleAppEngineApplicationSpec struct {
	GcrDomain       string                           `json:"gcr_domain"`
	LocationId      string                           `json:"location_id"`
	FeatureSettings []GoogleAppEngineApplicationSpec `json:"feature_settings"`
	Name            string                           `json:"name"`
	DefaultBucket   string                           `json:"default_bucket"`
	CodeBucket      string                           `json:"code_bucket"`
	DefaultHostname string                           `json:"default_hostname"`
	Project         string                           `json:"project"`
	AuthDomain      string                           `json:"auth_domain"`
	ServingStatus   string                           `json:"serving_status"`
	UrlDispatchRule []GoogleAppEngineApplicationSpec `json:"url_dispatch_rule"`
}

type GoogleAppEngineApplicationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleAppEngineApplicationList is a list of GoogleAppEngineApplications
type GoogleAppEngineApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleAppEngineApplication CRD objects
	Items []GoogleAppEngineApplication `json:"items,omitempty"`
}
