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

type GoogleProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectSpec   `json:"spec,omitempty"`
	Status            GoogleProjectStatus `json:"status,omitempty"`
}

type GoogleProjectSpecAppEngineFeatureSettings struct {
	SplitHealthChecks bool `json:"split_health_checks"`
}

type GoogleProjectSpecAppEngineUrlDispatchRule struct {
	Domain  string `json:"domain"`
	Path    string `json:"path"`
	Service string `json:"service"`
}

type GoogleProjectSpecAppEngine struct {
	DefaultHostname string                       `json:"default_hostname"`
	AuthDomain      string                       `json:"auth_domain"`
	ServingStatus   string                       `json:"serving_status"`
	FeatureSettings []GoogleProjectSpecAppEngine `json:"feature_settings"`
	Name            string                       `json:"name"`
	CodeBucket      string                       `json:"code_bucket"`
	LocationId      string                       `json:"location_id"`
	UrlDispatchRule []GoogleProjectSpecAppEngine `json:"url_dispatch_rule"`
	DefaultBucket   string                       `json:"default_bucket"`
	GcrDomain       string                       `json:"gcr_domain"`
}

type GoogleProjectSpec struct {
	ProjectId         string              `json:"project_id"`
	SkipDelete        bool                `json:"skip_delete"`
	Name              string              `json:"name"`
	OrgId             string              `json:"org_id"`
	FolderId          string              `json:"folder_id"`
	PolicyEtag        string              `json:"policy_etag"`
	Labels            map[string]string   `json:"labels"`
	AutoCreateNetwork bool                `json:"auto_create_network"`
	PolicyData        string              `json:"policy_data"`
	Number            string              `json:"number"`
	BillingAccount    string              `json:"billing_account"`
	AppEngine         []GoogleProjectSpec `json:"app_engine"`
}



type GoogleProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectList is a list of GoogleProjects
type GoogleProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProject CRD objects
	Items []GoogleProject `json:"items,omitempty"`
}