package v1alpha1

import (
	"encoding/json"

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

type ContainerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerGroupSpec   `json:"spec,omitempty"`
	Status            ContainerGroupStatus `json:"status,omitempty"`
}

type ContainerGroupSpecContainerGpu struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
}

type ContainerGroupSpecContainerLivenessProbeHttpGet struct {
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type ContainerGroupSpecContainerLivenessProbe struct {
	// +optional
	Exec []string `json:"exec,omitempty" tf:"exec,omitempty"`
	// +optional
	FailureThreshold int `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
	// +optional
	HttpGet []ContainerGroupSpecContainerLivenessProbeHttpGet `json:"httpGet,omitempty" tf:"http_get,omitempty"`
	// +optional
	InitialDelaySeconds int `json:"initialDelaySeconds,omitempty" tf:"initial_delay_seconds,omitempty"`
	// +optional
	PeriodSeconds int `json:"periodSeconds,omitempty" tf:"period_seconds,omitempty"`
	// +optional
	SuccessThreshold int `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`
	// +optional
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type ContainerGroupSpecContainerPorts struct {
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type ContainerGroupSpecContainerReadinessProbeHttpGet struct {
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type ContainerGroupSpecContainerReadinessProbe struct {
	// +optional
	Exec []string `json:"exec,omitempty" tf:"exec,omitempty"`
	// +optional
	FailureThreshold int `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
	// +optional
	HttpGet []ContainerGroupSpecContainerReadinessProbeHttpGet `json:"httpGet,omitempty" tf:"http_get,omitempty"`
	// +optional
	InitialDelaySeconds int `json:"initialDelaySeconds,omitempty" tf:"initial_delay_seconds,omitempty"`
	// +optional
	PeriodSeconds int `json:"periodSeconds,omitempty" tf:"period_seconds,omitempty"`
	// +optional
	SuccessThreshold int `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`
	// +optional
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type ContainerGroupSpecContainerVolume struct {
	MountPath string `json:"mountPath" tf:"mount_path"`
	Name      string `json:"name" tf:"name"`
	// +optional
	ReadOnly           bool   `json:"readOnly,omitempty" tf:"read_only,omitempty"`
	ShareName          string `json:"shareName" tf:"share_name"`
	StorageAccountKey  string `json:"storageAccountKey" tf:"storage_account_key"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
}

type ContainerGroupSpecContainer struct {
	// +optional
	// Deprecated
	Command string `json:"command,omitempty" tf:"command,omitempty"`
	// +optional
	Commands []string    `json:"commands,omitempty" tf:"commands,omitempty"`
	Cpu      json.Number `json:"cpu" tf:"cpu"`
	// +optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Gpu   []ContainerGroupSpecContainerGpu `json:"gpu,omitempty" tf:"gpu,omitempty"`
	Image string                           `json:"image" tf:"image"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LivenessProbe []ContainerGroupSpecContainerLivenessProbe `json:"livenessProbe,omitempty" tf:"liveness_probe,omitempty"`
	Memory        json.Number                                `json:"memory" tf:"memory"`
	Name          string                                     `json:"name" tf:"name"`
	// +optional
	// Deprecated
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Ports []ContainerGroupSpecContainerPorts `json:"ports,omitempty" tf:"ports,omitempty"`
	// +optional
	// Deprecated
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReadinessProbe []ContainerGroupSpecContainerReadinessProbe `json:"readinessProbe,omitempty" tf:"readiness_probe,omitempty"`
	// +optional
	SecureEnvironmentVariables map[string]string `json:"-" sensitive:"true" tf:"secure_environment_variables,omitempty"`
	// +optional
	Volume []ContainerGroupSpecContainerVolume `json:"volume,omitempty" tf:"volume,omitempty"`
}

type ContainerGroupSpecDiagnosticsLogAnalytics struct {
	LogType string `json:"logType" tf:"log_type"`
	// +optional
	Metadata     map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	WorkspaceID  string            `json:"workspaceID" tf:"workspace_id"`
	WorkspaceKey string            `json:"-" sensitive:"true" tf:"workspace_key"`
}

type ContainerGroupSpecDiagnostics struct {
	// +kubebuilder:validation:MaxItems=1
	LogAnalytics []ContainerGroupSpecDiagnosticsLogAnalytics `json:"logAnalytics" tf:"log_analytics"`
}

type ContainerGroupSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	Type        string `json:"type" tf:"type"`
}

type ContainerGroupSpecImageRegistryCredential struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Server   string `json:"server" tf:"server"`
	Username string `json:"username" tf:"username"`
}

type ContainerGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Container []ContainerGroupSpecContainer `json:"container" tf:"container"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Diagnostics []ContainerGroupSpecDiagnostics `json:"diagnostics,omitempty" tf:"diagnostics,omitempty"`
	// +optional
	DnsNameLabel string `json:"dnsNameLabel,omitempty" tf:"dns_name_label,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []ContainerGroupSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	ImageRegistryCredential []ContainerGroupSpecImageRegistryCredential `json:"imageRegistryCredential,omitempty" tf:"image_registry_credential,omitempty"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	// +optional
	IpAddressType     string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	OsType            string `json:"osType" tf:"os_type"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RestartPolicy string `json:"restartPolicy,omitempty" tf:"restart_policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ContainerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ContainerGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
