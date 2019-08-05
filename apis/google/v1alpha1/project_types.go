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

type Project struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec,omitempty"`
	Status            ProjectStatus `json:"status,omitempty"`
}

type ProjectSpecAppEngineFeatureSettings struct {
	// +optional
	SplitHealthChecks bool `json:"splitHealthChecks,omitempty" tf:"split_health_checks,omitempty"`
}

type ProjectSpecAppEngineUrlDispatchRule struct {
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Service string `json:"service,omitempty" tf:"service,omitempty"`
}

type ProjectSpecAppEngine struct {
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
	FeatureSettings []ProjectSpecAppEngineFeatureSettings `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`
	// +optional
	GcrDomain string `json:"gcrDomain,omitempty" tf:"gcr_domain,omitempty"`
	// +optional
	LocationID string `json:"locationID,omitempty" tf:"location_id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	ServingStatus string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
	// +optional
	UrlDispatchRule []ProjectSpecAppEngineUrlDispatchRule `json:"urlDispatchRule,omitempty" tf:"url_dispatch_rule,omitempty"`
}

type ProjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	AppEngine []ProjectSpecAppEngine `json:"appEngine,omitempty" tf:"app_engine,omitempty"`
	// +optional
	AutoCreateNetwork bool `json:"autoCreateNetwork,omitempty" tf:"auto_create_network,omitempty"`
	// +optional
	BillingAccount string `json:"billingAccount,omitempty" tf:"billing_account,omitempty"`
	// +optional
	FolderID string `json:"folderID,omitempty" tf:"folder_id,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Number string `json:"number,omitempty" tf:"number,omitempty"`
	// +optional
	OrgID     string `json:"orgID,omitempty" tf:"org_id,omitempty"`
	ProjectID string `json:"projectID" tf:"project_id"`
	// +optional
	SkipDelete bool `json:"skipDelete,omitempty" tf:"skip_delete,omitempty"`
}

type ProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ProjectSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectList is a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Project CRD objects
	Items []Project `json:"items,omitempty"`
}
