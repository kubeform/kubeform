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

type AzurermAppServiceSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceSlotSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceSlotStatus `json:"status,omitempty"`
}

type AzurermAppServiceSlotSpecSiteConfigCors struct {
	AllowedOrigins     []string `json:"allowed_origins"`
	SupportCredentials bool     `json:"support_credentials"`
}

type AzurermAppServiceSlotSpecSiteConfigIpRestriction struct {
	IpAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
}

type AzurermAppServiceSlotSpecSiteConfig struct {
	Cors                   []AzurermAppServiceSlotSpecSiteConfig `json:"cors"`
	DefaultDocuments       []string                              `json:"default_documents"`
	WebsocketsEnabled      bool                                  `json:"websockets_enabled"`
	FtpsState              string                                `json:"ftps_state"`
	LinuxFxVersion         string                                `json:"linux_fx_version"`
	WindowsFxVersion       string                                `json:"windows_fx_version"`
	PhpVersion             string                                `json:"php_version"`
	ScmType                string                                `json:"scm_type"`
	VirtualNetworkName     string                                `json:"virtual_network_name"`
	AlwaysOn               bool                                  `json:"always_on"`
	Http2Enabled           bool                                  `json:"http2_enabled"`
	IpRestriction          []AzurermAppServiceSlotSpecSiteConfig `json:"ip_restriction"`
	JavaContainer          string                                `json:"java_container"`
	JavaContainerVersion   string                                `json:"java_container_version"`
	RemoteDebuggingVersion string                                `json:"remote_debugging_version"`
	AppCommandLine         string                                `json:"app_command_line"`
	DotnetFrameworkVersion string                                `json:"dotnet_framework_version"`
	JavaVersion            string                                `json:"java_version"`
	LocalMysqlEnabled      bool                                  `json:"local_mysql_enabled"`
	PythonVersion          string                                `json:"python_version"`
	ManagedPipelineMode    string                                `json:"managed_pipeline_mode"`
	RemoteDebuggingEnabled bool                                  `json:"remote_debugging_enabled"`
	Use32BitWorkerProcess  bool                                  `json:"use_32_bit_worker_process"`
	MinTlsVersion          string                                `json:"min_tls_version"`
}

type AzurermAppServiceSlotSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermAppServiceSlotSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermAppServiceSlotSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermAppServiceSlotSpec struct {
	Name                  string                      `json:"name"`
	SiteConfig            []AzurermAppServiceSlotSpec `json:"site_config"`
	AppServicePlanId      string                      `json:"app_service_plan_id"`
	AppSettings           map[string]string           `json:"app_settings"`
	ResourceGroupName     string                      `json:"resource_group_name"`
	Identity              []AzurermAppServiceSlotSpec `json:"identity"`
	ClientAffinityEnabled bool                        `json:"client_affinity_enabled"`
	HttpsOnly             bool                        `json:"https_only"`
	Tags                  map[string]string           `json:"tags"`
	Location              string                      `json:"location"`
	AppServiceName        string                      `json:"app_service_name"`
	Enabled               bool                        `json:"enabled"`
	ConnectionString      []AzurermAppServiceSlotSpec `json:"connection_string"`
	SiteCredential        []AzurermAppServiceSlotSpec `json:"site_credential"`
	DefaultSiteHostname   string                      `json:"default_site_hostname"`
}

type AzurermAppServiceSlotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAppServiceSlotList is a list of AzurermAppServiceSlots
type AzurermAppServiceSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppServiceSlot CRD objects
	Items []AzurermAppServiceSlot `json:"items,omitempty"`
}
