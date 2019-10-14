package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Service string `json:"service,omitempty" tf:"service,omitempty"`
}

type AppEngineApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AuthDomain string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`
	// +optional
	CodeBucket string `json:"codeBucket,omitempty" tf:"code_bucket,omitempty"`
	// +optional
	DefaultBucket string `json:"defaultBucket,omitempty" tf:"default_bucket,omitempty"`
	// +optional
	DefaultHostname string `json:"defaultHostname,omitempty" tf:"default_hostname,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FeatureSettings []AppEngineApplicationSpecFeatureSettings `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`
	// +optional
	GcrDomain  string `json:"gcrDomain,omitempty" tf:"gcr_domain,omitempty"`
	LocationID string `json:"locationID" tf:"location_id"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ServingStatus string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
	// +optional
	UrlDispatchRule []AppEngineApplicationSpecUrlDispatchRule `json:"urlDispatchRule,omitempty" tf:"url_dispatch_rule,omitempty"`
}

type AppEngineApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppEngineApplicationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
