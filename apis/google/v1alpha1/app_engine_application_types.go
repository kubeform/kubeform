package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppEngineApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppEngineApplicationSpec   `json:"spec,omitempty"`
	Status            AppEngineApplicationStatus `json:"status,omitempty"`
}

type AppEngineApplicationSpecFeatureSettings struct {
	// +optional
	SplitHealthChecks bool `json:"splitHealthChecks,omitempty" tf:"split_health_checks,omitempty"`
}

type AppEngineApplicationSpecUrlDispatchRule struct {
	Domain  string `json:"domain" tf:"domain"`
	Path    string `json:"path" tf:"path"`
	Service string `json:"service" tf:"service"`
}

type AppEngineApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AuthDomain      string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`
	CodeBucket      string `json:"codeBucket" tf:"code_bucket"`
	DefaultBucket   string `json:"defaultBucket" tf:"default_bucket"`
	DefaultHostname string `json:"defaultHostname" tf:"default_hostname"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FeatureSettings []AppEngineApplicationSpecFeatureSettings `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`
	GcrDomain       string                                    `json:"gcrDomain" tf:"gcr_domain"`
	LocationID      string                                    `json:"locationID" tf:"location_id"`
	Name            string                                    `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ServingStatus   string                                    `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
	UrlDispatchRule []AppEngineApplicationSpecUrlDispatchRule `json:"urlDispatchRule" tf:"url_dispatch_rule"`
}

type AppEngineApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppEngineApplicationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppEngineApplicationList is a list of AppEngineApplications
type AppEngineApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppEngineApplication CRD objects
	Items []AppEngineApplication `json:"items,omitempty"`
}
