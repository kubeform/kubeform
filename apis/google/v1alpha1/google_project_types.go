package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Service string `json:"service"`
	Domain  string `json:"domain"`
	Path    string `json:"path"`
}

type GoogleProjectSpecAppEngine struct {
	FeatureSettings []GoogleProjectSpecAppEngine `json:"feature_settings"`
	Name            string                       `json:"name"`
	UrlDispatchRule []GoogleProjectSpecAppEngine `json:"url_dispatch_rule"`
	CodeBucket      string                       `json:"code_bucket"`
	DefaultHostname string                       `json:"default_hostname"`
	DefaultBucket   string                       `json:"default_bucket"`
	LocationId      string                       `json:"location_id"`
	ServingStatus   string                       `json:"serving_status"`
	GcrDomain       string                       `json:"gcr_domain"`
	AuthDomain      string                       `json:"auth_domain"`
}

type GoogleProjectSpec struct {
	AutoCreateNetwork bool                `json:"auto_create_network"`
	OrgId             string              `json:"org_id"`
	FolderId          string              `json:"folder_id"`
	PolicyEtag        string              `json:"policy_etag"`
	Number            string              `json:"number"`
	ProjectId         string              `json:"project_id"`
	SkipDelete        bool                `json:"skip_delete"`
	Name              string              `json:"name"`
	PolicyData        string              `json:"policy_data"`
	BillingAccount    string              `json:"billing_account"`
	Labels            map[string]string   `json:"labels"`
	AppEngine         []GoogleProjectSpec `json:"app_engine"`
}



type GoogleProjectStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleProjectList is a list of GoogleProjects
type GoogleProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProject CRD objects
	Items []GoogleProject `json:"items,omitempty"`
}