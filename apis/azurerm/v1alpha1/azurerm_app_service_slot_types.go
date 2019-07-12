package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAppServiceSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceSlotSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceSlotStatus `json:"status,omitempty"`
}

type AzurermAppServiceSlotSpecIdentity struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
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
	DotnetFrameworkVersion string                                `json:"dotnet_framework_version"`
	ScmType                string                                `json:"scm_type"`
	Use32BitWorkerProcess  bool                                  `json:"use_32_bit_worker_process"`
	WebsocketsEnabled      bool                                  `json:"websockets_enabled"`
	VirtualNetworkName     string                                `json:"virtual_network_name"`
	AppCommandLine         string                                `json:"app_command_line"`
	JavaContainer          string                                `json:"java_container"`
	ManagedPipelineMode    string                                `json:"managed_pipeline_mode"`
	PythonVersion          string                                `json:"python_version"`
	RemoteDebuggingVersion string                                `json:"remote_debugging_version"`
	LinuxFxVersion         string                                `json:"linux_fx_version"`
	WindowsFxVersion       string                                `json:"windows_fx_version"`
	AlwaysOn               bool                                  `json:"always_on"`
	DefaultDocuments       []string                              `json:"default_documents"`
	JavaContainerVersion   string                                `json:"java_container_version"`
	LocalMysqlEnabled      bool                                  `json:"local_mysql_enabled"`
	PhpVersion             string                                `json:"php_version"`
	MinTlsVersion          string                                `json:"min_tls_version"`
	Cors                   []AzurermAppServiceSlotSpecSiteConfig `json:"cors"`
	Http2Enabled           bool                                  `json:"http2_enabled"`
	IpRestriction          []AzurermAppServiceSlotSpecSiteConfig `json:"ip_restriction"`
	JavaVersion            string                                `json:"java_version"`
	RemoteDebuggingEnabled bool                                  `json:"remote_debugging_enabled"`
	FtpsState              string                                `json:"ftps_state"`
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
	Location              string                      `json:"location"`
	AppServiceName        string                      `json:"app_service_name"`
	AppSettings           map[string]string           `json:"app_settings"`
	DefaultSiteHostname   string                      `json:"default_site_hostname"`
	ResourceGroupName     string                      `json:"resource_group_name"`
	AppServicePlanId      string                      `json:"app_service_plan_id"`
	HttpsOnly             bool                        `json:"https_only"`
	Name                  string                      `json:"name"`
	Identity              []AzurermAppServiceSlotSpec `json:"identity"`
	Enabled               bool                        `json:"enabled"`
	SiteConfig            []AzurermAppServiceSlotSpec `json:"site_config"`
	ClientAffinityEnabled bool                        `json:"client_affinity_enabled"`
	ConnectionString      []AzurermAppServiceSlotSpec `json:"connection_string"`
	Tags                  map[string]string           `json:"tags"`
	SiteCredential        []AzurermAppServiceSlotSpec `json:"site_credential"`
}



type AzurermAppServiceSlotStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAppServiceSlotList is a list of AzurermAppServiceSlots
type AzurermAppServiceSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppServiceSlot CRD objects
	Items []AzurermAppServiceSlot `json:"items,omitempty"`
}