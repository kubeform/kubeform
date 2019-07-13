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

type AzurermContainerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermContainerGroupSpec   `json:"spec,omitempty"`
	Status            AzurermContainerGroupStatus `json:"status,omitempty"`
}

type AzurermContainerGroupSpecDiagnosticsLogAnalytics struct {
	LogType      string            `json:"log_type"`
	Metadata     map[string]string `json:"metadata"`
	WorkspaceId  string            `json:"workspace_id"`
	WorkspaceKey string            `json:"workspace_key"`
}

type AzurermContainerGroupSpecDiagnostics struct {
	LogAnalytics []AzurermContainerGroupSpecDiagnostics `json:"log_analytics"`
}

type AzurermContainerGroupSpecImageRegistryCredential struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermContainerGroupSpecIdentity struct {
	Type        string   `json:"type"`
	PrincipalId string   `json:"principal_id"`
	IdentityIds []string `json:"identity_ids"`
}

type AzurermContainerGroupSpecContainerPorts struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AzurermContainerGroupSpecContainerReadinessProbeHttpGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type AzurermContainerGroupSpecContainerReadinessProbe struct {
	TimeoutSeconds      int                                                `json:"timeout_seconds"`
	Exec                []string                                           `json:"exec"`
	HttpGet             []AzurermContainerGroupSpecContainerReadinessProbe `json:"http_get"`
	InitialDelaySeconds int                                                `json:"initial_delay_seconds"`
	PeriodSeconds       int                                                `json:"period_seconds"`
	FailureThreshold    int                                                `json:"failure_threshold"`
	SuccessThreshold    int                                                `json:"success_threshold"`
}

type AzurermContainerGroupSpecContainerVolume struct {
	StorageAccountKey  string `json:"storage_account_key"`
	Name               string `json:"name"`
	MountPath          string `json:"mount_path"`
	ReadOnly           bool   `json:"read_only"`
	ShareName          string `json:"share_name"`
	StorageAccountName string `json:"storage_account_name"`
}

type AzurermContainerGroupSpecContainerLivenessProbeHttpGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type AzurermContainerGroupSpecContainerLivenessProbe struct {
	InitialDelaySeconds int                                               `json:"initial_delay_seconds"`
	PeriodSeconds       int                                               `json:"period_seconds"`
	FailureThreshold    int                                               `json:"failure_threshold"`
	SuccessThreshold    int                                               `json:"success_threshold"`
	TimeoutSeconds      int                                               `json:"timeout_seconds"`
	Exec                []string                                          `json:"exec"`
	HttpGet             []AzurermContainerGroupSpecContainerLivenessProbe `json:"http_get"`
}

type AzurermContainerGroupSpecContainerGpu struct {
	Count int    `json:"count"`
	Sku   string `json:"sku"`
}

type AzurermContainerGroupSpecContainer struct {
	Cpu                        float64                              `json:"cpu"`
	Ports                      []AzurermContainerGroupSpecContainer `json:"ports"`
	Memory                     float64                              `json:"memory"`
	Protocol                   string                               `json:"protocol"`
	EnvironmentVariables       map[string]string                    `json:"environment_variables"`
	Command                    string                               `json:"command"`
	ReadinessProbe             []AzurermContainerGroupSpecContainer `json:"readiness_probe"`
	SecureEnvironmentVariables map[string]string                    `json:"secure_environment_variables"`
	Commands                   []string                             `json:"commands"`
	Volume                     []AzurermContainerGroupSpecContainer `json:"volume"`
	LivenessProbe              []AzurermContainerGroupSpecContainer `json:"liveness_probe"`
	Name                       string                               `json:"name"`
	Image                      string                               `json:"image"`
	Gpu                        []AzurermContainerGroupSpecContainer `json:"gpu"`
	Port                       int                                  `json:"port"`
}

type AzurermContainerGroupSpec struct {
	IpAddressType           string                      `json:"ip_address_type"`
	Diagnostics             []AzurermContainerGroupSpec `json:"diagnostics"`
	Fqdn                    string                      `json:"fqdn"`
	DnsNameLabel            string                      `json:"dns_name_label"`
	OsType                  string                      `json:"os_type"`
	Tags                    map[string]string           `json:"tags"`
	RestartPolicy           string                      `json:"restart_policy"`
	Name                    string                      `json:"name"`
	ResourceGroupName       string                      `json:"resource_group_name"`
	ImageRegistryCredential []AzurermContainerGroupSpec `json:"image_registry_credential"`
	IpAddress               string                      `json:"ip_address"`
	Location                string                      `json:"location"`
	Identity                []AzurermContainerGroupSpec `json:"identity"`
	Container               []AzurermContainerGroupSpec `json:"container"`
}



type AzurermContainerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermContainerGroupList is a list of AzurermContainerGroups
type AzurermContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerGroup CRD objects
	Items []AzurermContainerGroup `json:"items,omitempty"`
}