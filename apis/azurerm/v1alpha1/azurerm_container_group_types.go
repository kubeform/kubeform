package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermContainerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermContainerGroupSpec   `json:"spec,omitempty"`
	Status            AzurermContainerGroupStatus `json:"status,omitempty"`
}

type AzurermContainerGroupSpecContainerGpu struct {
	Count int    `json:"count"`
	Sku   string `json:"sku"`
}

type AzurermContainerGroupSpecContainerVolume struct {
	Name               string `json:"name"`
	MountPath          string `json:"mount_path"`
	ReadOnly           bool   `json:"read_only"`
	ShareName          string `json:"share_name"`
	StorageAccountName string `json:"storage_account_name"`
	StorageAccountKey  string `json:"storage_account_key"`
}

type AzurermContainerGroupSpecContainerLivenessProbeHttpGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type AzurermContainerGroupSpecContainerLivenessProbe struct {
	HttpGet             []AzurermContainerGroupSpecContainerLivenessProbe `json:"http_get"`
	InitialDelaySeconds int                                               `json:"initial_delay_seconds"`
	PeriodSeconds       int                                               `json:"period_seconds"`
	FailureThreshold    int                                               `json:"failure_threshold"`
	SuccessThreshold    int                                               `json:"success_threshold"`
	TimeoutSeconds      int                                               `json:"timeout_seconds"`
	Exec                []string                                          `json:"exec"`
}

type AzurermContainerGroupSpecContainerPorts struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type AzurermContainerGroupSpecContainerReadinessProbeHttpGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type AzurermContainerGroupSpecContainerReadinessProbe struct {
	Exec                []string                                           `json:"exec"`
	HttpGet             []AzurermContainerGroupSpecContainerReadinessProbe `json:"http_get"`
	InitialDelaySeconds int                                                `json:"initial_delay_seconds"`
	PeriodSeconds       int                                                `json:"period_seconds"`
	FailureThreshold    int                                                `json:"failure_threshold"`
	SuccessThreshold    int                                                `json:"success_threshold"`
	TimeoutSeconds      int                                                `json:"timeout_seconds"`
}

type AzurermContainerGroupSpecContainer struct {
	Name                       string                               `json:"name"`
	Memory                     float64                              `json:"memory"`
	Gpu                        []AzurermContainerGroupSpecContainer `json:"gpu"`
	Port                       int                                  `json:"port"`
	Command                    string                               `json:"command"`
	EnvironmentVariables       map[string]string                    `json:"environment_variables"`
	Volume                     []AzurermContainerGroupSpecContainer `json:"volume"`
	Protocol                   string                               `json:"protocol"`
	Commands                   []string                             `json:"commands"`
	LivenessProbe              []AzurermContainerGroupSpecContainer `json:"liveness_probe"`
	Image                      string                               `json:"image"`
	Cpu                        float64                              `json:"cpu"`
	Ports                      []AzurermContainerGroupSpecContainer `json:"ports"`
	SecureEnvironmentVariables map[string]string                    `json:"secure_environment_variables"`
	ReadinessProbe             []AzurermContainerGroupSpecContainer `json:"readiness_probe"`
}

type AzurermContainerGroupSpecDiagnosticsLogAnalytics struct {
	WorkspaceId  string            `json:"workspace_id"`
	WorkspaceKey string            `json:"workspace_key"`
	LogType      string            `json:"log_type"`
	Metadata     map[string]string `json:"metadata"`
}

type AzurermContainerGroupSpecDiagnostics struct {
	LogAnalytics []AzurermContainerGroupSpecDiagnostics `json:"log_analytics"`
}

type AzurermContainerGroupSpecImageRegistryCredential struct {
	Password string `json:"password"`
	Server   string `json:"server"`
	Username string `json:"username"`
}

type AzurermContainerGroupSpecIdentity struct {
	Type        string   `json:"type"`
	PrincipalId string   `json:"principal_id"`
	IdentityIds []string `json:"identity_ids"`
}

type AzurermContainerGroupSpec struct {
	Name                    string                      `json:"name"`
	ResourceGroupName       string                      `json:"resource_group_name"`
	IpAddressType           string                      `json:"ip_address_type"`
	Tags                    map[string]string           `json:"tags"`
	Container               []AzurermContainerGroupSpec `json:"container"`
	Diagnostics             []AzurermContainerGroupSpec `json:"diagnostics"`
	Location                string                      `json:"location"`
	RestartPolicy           string                      `json:"restart_policy"`
	DnsNameLabel            string                      `json:"dns_name_label"`
	Fqdn                    string                      `json:"fqdn"`
	OsType                  string                      `json:"os_type"`
	ImageRegistryCredential []AzurermContainerGroupSpec `json:"image_registry_credential"`
	Identity                []AzurermContainerGroupSpec `json:"identity"`
	IpAddress               string                      `json:"ip_address"`
}

type AzurermContainerGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermContainerGroupList is a list of AzurermContainerGroups
type AzurermContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerGroup CRD objects
	Items []AzurermContainerGroup `json:"items,omitempty"`
}
