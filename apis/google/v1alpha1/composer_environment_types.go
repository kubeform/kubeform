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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status            ComposerEnvironmentStatus `json:"status,omitempty"`
}

type ComposerEnvironmentSpecConfigNodeConfig struct {
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComposerEnvironmentSpecConfigSoftwareConfig struct {
	// +optional
	AirflowConfigOverrides map[string]string `json:"airflowConfigOverrides,omitempty" tf:"airflow_config_overrides,omitempty"`
	// +optional
	EnvVariables map[string]string `json:"envVariables,omitempty" tf:"env_variables,omitempty"`
	// +optional
	ImageVersion string `json:"imageVersion,omitempty" tf:"image_version,omitempty"`
	// +optional
	PypiPackages map[string]string `json:"pypiPackages,omitempty" tf:"pypi_packages,omitempty"`
}

type ComposerEnvironmentSpecConfig struct {
	// +optional
	AirflowURI string `json:"airflowURI,omitempty" tf:"airflow_uri,omitempty"`
	// +optional
	DagGcsPrefix string `json:"dagGcsPrefix,omitempty" tf:"dag_gcs_prefix,omitempty"`
	// +optional
	GkeCluster string `json:"gkeCluster,omitempty" tf:"gke_cluster,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NodeConfig []ComposerEnvironmentSpecConfigNodeConfig `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`
	// +optional
	NodeCount int `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SoftwareConfig []ComposerEnvironmentSpecConfigSoftwareConfig `json:"softwareConfig,omitempty" tf:"software_config,omitempty"`
}

type ComposerEnvironmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Config []ComposerEnvironmentSpecConfig `json:"config,omitempty" tf:"config,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
}

type ComposerEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComposerEnvironmentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComposerEnvironmentList is a list of ComposerEnvironments
type ComposerEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComposerEnvironment CRD objects
	Items []ComposerEnvironment `json:"items,omitempty"`
}
