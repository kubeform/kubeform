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

type GoogleComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status            GoogleComposerEnvironmentStatus `json:"status,omitempty"`
}

type GoogleComposerEnvironmentSpecConfigNodeConfig struct {
	Network        string   `json:"network"`
	Subnetwork     string   `json:"subnetwork"`
	DiskSizeGb     int      `json:"disk_size_gb"`
	OauthScopes    []string `json:"oauth_scopes"`
	ServiceAccount string   `json:"service_account"`
	Tags           []string `json:"tags"`
	Zone           string   `json:"zone"`
	MachineType    string   `json:"machine_type"`
}

type GoogleComposerEnvironmentSpecConfigSoftwareConfig struct {
	AirflowConfigOverrides map[string]string `json:"airflow_config_overrides"`
	PypiPackages           map[string]string `json:"pypi_packages"`
	EnvVariables           map[string]string `json:"env_variables"`
	ImageVersion           string            `json:"image_version"`
}

type GoogleComposerEnvironmentSpecConfig struct {
	NodeCount      int                                   `json:"node_count"`
	NodeConfig     []GoogleComposerEnvironmentSpecConfig `json:"node_config"`
	SoftwareConfig []GoogleComposerEnvironmentSpecConfig `json:"software_config"`
	AirflowUri     string                                `json:"airflow_uri"`
	DagGcsPrefix   string                                `json:"dag_gcs_prefix"`
	GkeCluster     string                                `json:"gke_cluster"`
}

type GoogleComposerEnvironmentSpec struct {
	Project string                          `json:"project"`
	Config  []GoogleComposerEnvironmentSpec `json:"config"`
	Labels  map[string]string               `json:"labels"`
	Name    string                          `json:"name"`
	Region  string                          `json:"region"`
}

type GoogleComposerEnvironmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComposerEnvironmentList is a list of GoogleComposerEnvironments
type GoogleComposerEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComposerEnvironment CRD objects
	Items []GoogleComposerEnvironment `json:"items,omitempty"`
}
