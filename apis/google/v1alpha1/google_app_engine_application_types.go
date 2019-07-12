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
	Service string `json:"service"`
	Domain  string `json:"domain"`
	Path    string `json:"path"`
}

type GoogleAppEngineApplicationSpec struct {
	AuthDomain      string                           `json:"auth_domain"`
	DefaultHostname string                           `json:"default_hostname"`
	GcrDomain       string                           `json:"gcr_domain"`
	CodeBucket      string                           `json:"code_bucket"`
	DefaultBucket   string                           `json:"default_bucket"`
	Project         string                           `json:"project"`
	LocationId      string                           `json:"location_id"`
	ServingStatus   string                           `json:"serving_status"`
	FeatureSettings []GoogleAppEngineApplicationSpec `json:"feature_settings"`
	Name            string                           `json:"name"`
	UrlDispatchRule []GoogleAppEngineApplicationSpec `json:"url_dispatch_rule"`
}

type GoogleAppEngineApplicationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleAppEngineApplicationList is a list of GoogleAppEngineApplications
type GoogleAppEngineApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleAppEngineApplication CRD objects
	Items []GoogleAppEngineApplication `json:"items,omitempty"`
}
