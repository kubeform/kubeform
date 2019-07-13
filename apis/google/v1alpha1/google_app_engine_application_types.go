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

type GoogleAppEngineApplicationSpecUrlDispatchRule struct {
	Domain  string `json:"domain"`
	Path    string `json:"path"`
	Service string `json:"service"`
}

type GoogleAppEngineApplicationSpecFeatureSettings struct {
	SplitHealthChecks bool `json:"split_health_checks"`
}

type GoogleAppEngineApplicationSpec struct {
	Project         string                           `json:"project"`
	ServingStatus   string                           `json:"serving_status"`
	Name            string                           `json:"name"`
	UrlDispatchRule []GoogleAppEngineApplicationSpec `json:"url_dispatch_rule"`
	CodeBucket      string                           `json:"code_bucket"`
	AuthDomain      string                           `json:"auth_domain"`
	LocationId      string                           `json:"location_id"`
	FeatureSettings []GoogleAppEngineApplicationSpec `json:"feature_settings"`
	DefaultHostname string                           `json:"default_hostname"`
	DefaultBucket   string                           `json:"default_bucket"`
	GcrDomain       string                           `json:"gcr_domain"`
}



type GoogleAppEngineApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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