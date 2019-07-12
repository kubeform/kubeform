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
	WorkspaceId  string            `json:"workspace_id"`
	WorkspaceKey string            `json:"workspace_key"`
	LogType      string            `json:"log_type"`
	Metadata     map[string]string `json:"metadata"`
}

type AzurermContainerGroupSpecDiagnostics struct {
	LogAnalytics []AzurermContainerGroupSpecDiagnostics `json:"log_analytics"`
}

type AzurermContainerGroupSpecIdentity struct {
	Type        string   `json:"type"`
	PrincipalId string   `json:"principal_id"`
	IdentityIds []string `json:"identity_ids"`
}

type AzurermContainerGroupSpecContainerVolume struct {
	ShareName          string `json:"share_name"`
	StorageAccountName string `json:"storage_account_name"`
	StorageAccountKey  string `json:"storage_account_key"`
	Name               string `json:"name"`
	MountPath          string `json:"mount_path"`
	ReadOnly           bool   `json:"read_only"`
}

type AzurermContainerGroupSpecContainerReadinessProbeHttpGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type AzurermContainerGroupSpecContainerReadinessProbe struct {
	FailureThreshold    int                                                `json:"failure_threshold"`
	SuccessThreshold    int                                                `json:"success_threshold"`
	TimeoutSeconds      int                                                `json:"timeout_seconds"`
	Exec                []string                                           `json:"exec"`
	HttpGet             []AzurermContainerGroupSpecContainerReadinessProbe `json:"http_get"`
	InitialDelaySeconds int                                                `json:"initial_delay_seconds"`
	PeriodSeconds       int                                                `json:"period_seconds"`
}

type AzurermContainerGroupSpecContainerLivenessProbeHttpGet struct {
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
	Path   string `json:"path"`
}

type AzurermContainerGroupSpecContainerLivenessProbe struct {
	TimeoutSeconds      int                                               `json:"timeout_seconds"`
	Exec                []string                                          `json:"exec"`
	HttpGet             []AzurermContainerGroupSpecContainerLivenessProbe `json:"http_get"`
	InitialDelaySeconds int                                               `json:"initial_delay_seconds"`
	PeriodSeconds       int                                               `json:"period_seconds"`
	FailureThreshold    int                                               `json:"failure_threshold"`
	SuccessThreshold    int                                               `json:"success_threshold"`
}

type AzurermContainerGroupSpecContainerGpu struct {
	Count int    `json:"count"`
	Sku   string `json:"sku"`
}

type AzurermContainerGroupSpecContainerPorts struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AzurermContainerGroupSpecContainer struct {
	Name                       string                               `json:"name"`
	Image                      string                               `json:"image"`
	Cpu                        float64                              `json:"cpu"`
	SecureEnvironmentVariables map[string]string                    `json:"secure_environment_variables"`
	Volume                     []AzurermContainerGroupSpecContainer `json:"volume"`
	ReadinessProbe             []AzurermContainerGroupSpecContainer `json:"readiness_probe"`
	Port                       int                                  `json:"port"`
	Protocol                   string                               `json:"protocol"`
	Commands                   []string                             `json:"commands"`
	LivenessProbe              []AzurermContainerGroupSpecContainer `json:"liveness_probe"`
	Memory                     float64                              `json:"memory"`
	Gpu                        []AzurermContainerGroupSpecContainer `json:"gpu"`
	Ports                      []AzurermContainerGroupSpecContainer `json:"ports"`
	EnvironmentVariables       map[string]string                    `json:"environment_variables"`
	Command                    string                               `json:"command"`
}

type AzurermContainerGroupSpecImageRegistryCredential struct {
	Password string `json:"password"`
	Server   string `json:"server"`
	Username string `json:"username"`
}

type AzurermContainerGroupSpec struct {
	IpAddress               string                      `json:"ip_address"`
	Name                    string                      `json:"name"`
	Diagnostics             []AzurermContainerGroupSpec `json:"diagnostics"`
	Fqdn                    string                      `json:"fqdn"`
	Location                string                      `json:"location"`
	ResourceGroupName       string                      `json:"resource_group_name"`
	IpAddressType           string                      `json:"ip_address_type"`
	OsType                  string                      `json:"os_type"`
	Identity                []AzurermContainerGroupSpec `json:"identity"`
	DnsNameLabel            string                      `json:"dns_name_label"`
	Container               []AzurermContainerGroupSpec `json:"container"`
	ImageRegistryCredential []AzurermContainerGroupSpec `json:"image_registry_credential"`
	Tags                    map[string]string           `json:"tags"`
	RestartPolicy           string                      `json:"restart_policy"`
}

type AzurermContainerGroupStatus struct {
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
