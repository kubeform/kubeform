package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ContainerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerGroupSpec   `json:"spec,omitempty"`
	Status            ContainerGroupStatus `json:"status,omitempty"`
}

type ContainerGroupSpecContainerGpu struct {
	// +optional
	Count int `json:"count,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty"`
}

type ContainerGroupSpecContainerLivenessProbeHttpGet struct {
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty"`
}

type ContainerGroupSpecContainerLivenessProbe struct {
	// +optional
	Exec []string `json:"exec,omitempty"`
	// +optional
	FailureThreshold int `json:"failure_threshold,omitempty"`
	// +optional
	HttpGet *[]ContainerGroupSpecContainerLivenessProbe `json:"http_get,omitempty"`
	// +optional
	InitialDelaySeconds int `json:"initial_delay_seconds,omitempty"`
	// +optional
	PeriodSeconds int `json:"period_seconds,omitempty"`
	// +optional
	SuccessThreshold int `json:"success_threshold,omitempty"`
	// +optional
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type ContainerGroupSpecContainerReadinessProbeHttpGet struct {
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty"`
}

type ContainerGroupSpecContainerReadinessProbe struct {
	// +optional
	Exec []string `json:"exec,omitempty"`
	// +optional
	FailureThreshold int `json:"failure_threshold,omitempty"`
	// +optional
	HttpGet *[]ContainerGroupSpecContainerReadinessProbe `json:"http_get,omitempty"`
	// +optional
	InitialDelaySeconds int `json:"initial_delay_seconds,omitempty"`
	// +optional
	PeriodSeconds int `json:"period_seconds,omitempty"`
	// +optional
	SuccessThreshold int `json:"success_threshold,omitempty"`
	// +optional
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type ContainerGroupSpecContainerVolume struct {
	MountPath string `json:"mount_path"`
	Name      string `json:"name"`
	// +optional
	ReadOnly           bool   `json:"read_only,omitempty"`
	ShareName          string `json:"share_name"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageAccountName string `json:"storage_account_name"`
}

type ContainerGroupSpecContainer struct {
	Cpu json.Number `json:"cpu"`
	// +optional
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Gpu   *[]ContainerGroupSpecContainer `json:"gpu,omitempty"`
	Image string                         `json:"image"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LivenessProbe *[]ContainerGroupSpecContainer `json:"liveness_probe,omitempty"`
	Memory        json.Number                    `json:"memory"`
	Name          string                         `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReadinessProbe *[]ContainerGroupSpecContainer `json:"readiness_probe,omitempty"`
	// +optional
	SecureEnvironmentVariables map[string]string `json:"secure_environment_variables,omitempty"`
	// +optional
	Volume *[]ContainerGroupSpecContainer `json:"volume,omitempty"`
}

type ContainerGroupSpecDiagnosticsLogAnalytics struct {
	LogType string `json:"log_type"`
	// +optional
	Metadata     map[string]string `json:"metadata,omitempty"`
	WorkspaceId  string            `json:"workspace_id"`
	WorkspaceKey string            `json:"workspace_key"`
}

type ContainerGroupSpecDiagnostics struct {
	// +kubebuilder:validation:MaxItems=1
	LogAnalytics []ContainerGroupSpecDiagnostics `json:"log_analytics"`
}

type ContainerGroupSpecImageRegistryCredential struct {
	Password string `json:"password"`
	Server   string `json:"server"`
	Username string `json:"username"`
}

type ContainerGroupSpec struct {
	Container []ContainerGroupSpec `json:"container"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Diagnostics *[]ContainerGroupSpec `json:"diagnostics,omitempty"`
	// +optional
	DnsNameLabel string `json:"dns_name_label,omitempty"`
	// +optional
	ImageRegistryCredential *[]ContainerGroupSpec `json:"image_registry_credential,omitempty"`
	// +optional
	IpAddressType     string `json:"ip_address_type,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	OsType            string `json:"os_type"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	RestartPolicy string `json:"restart_policy,omitempty"`
}

type ContainerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerGroupList is a list of ContainerGroups
type ContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerGroup CRD objects
	Items []ContainerGroup `json:"items,omitempty"`
}
