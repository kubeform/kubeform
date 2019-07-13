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

type AzurermAppServiceSlotSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermAppServiceSlotSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermAppServiceSlotSpecSiteConfigIpRestriction struct {
	IpAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
}

type AzurermAppServiceSlotSpecSiteConfigCors struct {
	AllowedOrigins     []string `json:"allowed_origins"`
	SupportCredentials bool     `json:"support_credentials"`
}

type AzurermAppServiceSlotSpecSiteConfig struct {
	LocalMysqlEnabled      bool                                  `json:"local_mysql_enabled"`
	WebsocketsEnabled      bool                                  `json:"websockets_enabled"`
	LinuxFxVersion         string                                `json:"linux_fx_version"`
	IpRestriction          []AzurermAppServiceSlotSpecSiteConfig `json:"ip_restriction"`
	JavaContainer          string                                `json:"java_container"`
	MinTlsVersion          string                                `json:"min_tls_version"`
	AlwaysOn               bool                                  `json:"always_on"`
	AppCommandLine         string                                `json:"app_command_line"`
	DefaultDocuments       []string                              `json:"default_documents"`
	PhpVersion             string                                `json:"php_version"`
	PythonVersion          string                                `json:"python_version"`
	RemoteDebuggingEnabled bool                                  `json:"remote_debugging_enabled"`
	RemoteDebuggingVersion string                                `json:"remote_debugging_version"`
	FtpsState              string                                `json:"ftps_state"`
	DotnetFrameworkVersion string                                `json:"dotnet_framework_version"`
	Http2Enabled           bool                                  `json:"http2_enabled"`
	JavaVersion            string                                `json:"java_version"`
	WindowsFxVersion       string                                `json:"windows_fx_version"`
	Cors                   []AzurermAppServiceSlotSpecSiteConfig `json:"cors"`
	Use32BitWorkerProcess  bool                                  `json:"use_32_bit_worker_process"`
	VirtualNetworkName     string                                `json:"virtual_network_name"`
	JavaContainerVersion   string                                `json:"java_container_version"`
	ManagedPipelineMode    string                                `json:"managed_pipeline_mode"`
	ScmType                string                                `json:"scm_type"`
}

type AzurermAppServiceSlotSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermAppServiceSlotSpec struct {
	AppSettings           map[string]string           `json:"app_settings"`
	Name                  string                      `json:"name"`
	Identity              []AzurermAppServiceSlotSpec `json:"identity"`
	AppServicePlanId      string                      `json:"app_service_plan_id"`
	Enabled               bool                        `json:"enabled"`
	Location              string                      `json:"location"`
	AppServiceName        string                      `json:"app_service_name"`
	HttpsOnly             bool                        `json:"https_only"`
	DefaultSiteHostname   string                      `json:"default_site_hostname"`
	Tags                  map[string]string           `json:"tags"`
	SiteCredential        []AzurermAppServiceSlotSpec `json:"site_credential"`
	ResourceGroupName     string                      `json:"resource_group_name"`
	SiteConfig            []AzurermAppServiceSlotSpec `json:"site_config"`
	ClientAffinityEnabled bool                        `json:"client_affinity_enabled"`
	ConnectionString      []AzurermAppServiceSlotSpec `json:"connection_string"`
}



type AzurermAppServiceSlotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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