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
	GcrDomain       string                       `json:"gcr_domain"`
	LocationId      string                       `json:"location_id"`
	FeatureSettings []GoogleProjectSpecAppEngine `json:"feature_settings"`
	Name            string                       `json:"name"`
	CodeBucket      string                       `json:"code_bucket"`
	DefaultHostname string                       `json:"default_hostname"`
	DefaultBucket   string                       `json:"default_bucket"`
	AuthDomain      string                       `json:"auth_domain"`
	ServingStatus   string                       `json:"serving_status"`
	UrlDispatchRule []GoogleProjectSpecAppEngine `json:"url_dispatch_rule"`
}

type GoogleProjectSpec struct {
	Labels            map[string]string   `json:"labels"`
	AppEngine         []GoogleProjectSpec `json:"app_engine"`
	OrgId             string              `json:"org_id"`
	PolicyEtag        string              `json:"policy_etag"`
	AutoCreateNetwork bool                `json:"auto_create_network"`
	Name              string              `json:"name"`
	FolderId          string              `json:"folder_id"`
	PolicyData        string              `json:"policy_data"`
	Number            string              `json:"number"`
	BillingAccount    string              `json:"billing_account"`
	ProjectId         string              `json:"project_id"`
	SkipDelete        bool                `json:"skip_delete"`
}

type GoogleProjectStatus struct {
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
