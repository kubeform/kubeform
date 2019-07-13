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

type AzurermVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineScaleSetSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm struct {
	Protocol       string `json:"protocol"`
	CertificateUrl string `json:"certificate_url"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Pass        string `json:"pass"`
	Component   string `json:"component"`
	SettingName string `json:"setting_name"`
	Content     string `json:"content"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfig struct {
	EnableAutomaticUpgrades  bool                                                      `json:"enable_automatic_upgrades"`
	Winrm                    []AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"winrm"`
	AdditionalUnattendConfig []AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"additional_unattend_config"`
	ProvisionVmAgent         bool                                                      `json:"provision_vm_agent"`
}

type AzurermVirtualMachineScaleSetSpecPlan struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Product   string `json:"product"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileDnsSettings struct {
	DnsServers []string `json:"dns_servers"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIpAddressConfiguration struct {
	Name            string `json:"name"`
	IdleTimeout     int    `json:"idle_timeout"`
	DomainNameLabel string `json:"domain_name_label"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfiguration struct {
	LoadBalancerBackendAddressPoolIds       []string                                                         `json:"load_balancer_backend_address_pool_ids"`
	LoadBalancerInboundNatRulesIds          []string                                                         `json:"load_balancer_inbound_nat_rules_ids"`
	Primary                                 bool                                                             `json:"primary"`
	PublicIpAddressConfiguration            []AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfiguration `json:"public_ip_address_configuration"`
	Name                                    string                                                           `json:"name"`
	SubnetId                                string                                                           `json:"subnet_id"`
	ApplicationGatewayBackendAddressPoolIds []string                                                         `json:"application_gateway_backend_address_pool_ids"`
	ApplicationSecurityGroupIds             []string                                                         `json:"application_security_group_ids"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfile struct {
	NetworkSecurityGroupId string                                            `json:"network_security_group_id"`
	DnsSettings            []AzurermVirtualMachineScaleSetSpecNetworkProfile `json:"dns_settings"`
	IpConfiguration        []AzurermVirtualMachineScaleSetSpecNetworkProfile `json:"ip_configuration"`
	Name                   string                                            `json:"name"`
	Primary                bool                                              `json:"primary"`
	AcceleratedNetworking  bool                                              `json:"accelerated_networking"`
	IpForwarding           bool                                              `json:"ip_forwarding"`
}

type AzurermVirtualMachineScaleSetSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type AzurermVirtualMachineScaleSetSpecExtension struct {
	ProtectedSettings        string   `json:"protected_settings"`
	Name                     string   `json:"name"`
	Publisher                string   `json:"publisher"`
	Type                     string   `json:"type"`
	TypeHandlerVersion       string   `json:"type_handler_version"`
	AutoUpgradeMinorVersion  bool     `json:"auto_upgrade_minor_version"`
	ProvisionAfterExtensions []string `json:"provision_after_extensions"`
	Settings                 string   `json:"settings"`
}

type AzurermVirtualMachineScaleSetSpecIdentity struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	Type        string   `json:"type"`
}

type AzurermVirtualMachineScaleSetSpecSku struct {
	Name     string `json:"name"`
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
}

type AzurermVirtualMachineScaleSetSpecRollingUpgradePolicy struct {
	MaxBatchInstancePercent             int    `json:"max_batch_instance_percent"`
	MaxUnhealthyInstancePercent         int    `json:"max_unhealthy_instance_percent"`
	MaxUnhealthyUpgradedInstancePercent int    `json:"max_unhealthy_upgraded_instance_percent"`
	PauseTimeBetweenBatches             string `json:"pause_time_between_batches"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys struct {
	Path    string `json:"path"`
	KeyData string `json:"key_data"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool                                                    `json:"disable_password_authentication"`
	SshKeys                       []AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfig `json:"ssh_keys"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileOsDisk struct {
	Name            string   `json:"name"`
	Image           string   `json:"image"`
	VhdContainers   []string `json:"vhd_containers"`
	ManagedDiskType string   `json:"managed_disk_type"`
	Caching         string   `json:"caching"`
	OsType          string   `json:"os_type"`
	CreateOption    string   `json:"create_option"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileImageReference struct {
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
	Id        string `json:"id"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileDataDisk struct {
	Lun             int    `json:"lun"`
	CreateOption    string `json:"create_option"`
	Caching         string `json:"caching"`
	DiskSizeGb      int    `json:"disk_size_gb"`
	ManagedDiskType string `json:"managed_disk_type"`
}

type AzurermVirtualMachineScaleSetSpecOsProfile struct {
	AdminPassword      string `json:"admin_password"`
	CustomData         string `json:"custom_data"`
	ComputerNamePrefix string `json:"computer_name_prefix"`
	AdminUsername      string `json:"admin_username"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates struct {
	CertificateUrl   string `json:"certificate_url"`
	CertificateStore string `json:"certificate_store"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileSecrets struct {
	SourceVaultId     string                                              `json:"source_vault_id"`
	VaultCertificates []AzurermVirtualMachineScaleSetSpecOsProfileSecrets `json:"vault_certificates"`
}

type AzurermVirtualMachineScaleSetSpec struct {
	Tags                         map[string]string                   `json:"tags"`
	UpgradePolicyMode            string                              `json:"upgrade_policy_mode"`
	AutomaticOsUpgrade           bool                                `json:"automatic_os_upgrade"`
	Overprovision                bool                                `json:"overprovision"`
	Priority                     string                              `json:"priority"`
	EvictionPolicy               string                              `json:"eviction_policy"`
	OsProfileWindowsConfig       []AzurermVirtualMachineScaleSetSpec `json:"os_profile_windows_config"`
	Plan                         []AzurermVirtualMachineScaleSetSpec `json:"plan"`
	LicenseType                  string                              `json:"license_type"`
	Location                     string                              `json:"location"`
	SinglePlacementGroup         bool                                `json:"single_placement_group"`
	NetworkProfile               []AzurermVirtualMachineScaleSetSpec `json:"network_profile"`
	BootDiagnostics              []AzurermVirtualMachineScaleSetSpec `json:"boot_diagnostics"`
	Extension                    []AzurermVirtualMachineScaleSetSpec `json:"extension"`
	ResourceGroupName            string                              `json:"resource_group_name"`
	Zones                        []string                            `json:"zones"`
	Identity                     []AzurermVirtualMachineScaleSetSpec `json:"identity"`
	HealthProbeId                string                              `json:"health_probe_id"`
	Name                         string                              `json:"name"`
	Sku                          []AzurermVirtualMachineScaleSetSpec `json:"sku"`
	RollingUpgradePolicy         []AzurermVirtualMachineScaleSetSpec `json:"rolling_upgrade_policy"`
	OsProfileLinuxConfig         []AzurermVirtualMachineScaleSetSpec `json:"os_profile_linux_config"`
	StorageProfileOsDisk         []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_os_disk"`
	StorageProfileImageReference []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_image_reference"`
	StorageProfileDataDisk       []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_data_disk"`
	OsProfile                    []AzurermVirtualMachineScaleSetSpec `json:"os_profile"`
	OsProfileSecrets             []AzurermVirtualMachineScaleSetSpec `json:"os_profile_secrets"`
}



type AzurermVirtualMachineScaleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermVirtualMachineScaleSetList is a list of AzurermVirtualMachineScaleSets
type AzurermVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachineScaleSet CRD objects
	Items []AzurermVirtualMachineScaleSet `json:"items,omitempty"`
}