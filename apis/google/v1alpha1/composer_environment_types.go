package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

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
	// +kubebuilder:validation:UniqueItems=true
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
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
	PypiPackages map[string]string `json:"pypiPackages,omitempty" tf:"pypi_packages,omitempty"`
}

type ComposerEnvironmentSpecConfig struct {
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
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Config []ComposerEnvironmentSpecConfig `json:"config,omitempty" tf:"config,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region      string                    `json:"region,omitempty" tf:"region,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComposerEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
