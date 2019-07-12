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

type AzurermVirtualMachineScaleSetSpecOsProfile struct {
	ComputerNamePrefix string `json:"computer_name_prefix"`
	AdminUsername      string `json:"admin_username"`
	AdminPassword      string `json:"admin_password"`
	CustomData         string `json:"custom_data"`
}

type AzurermVirtualMachineScaleSetSpecPlan struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Product   string `json:"product"`
}

type AzurermVirtualMachineScaleSetSpecIdentity struct {
	Type        string   `json:"type"`
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
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
	ProvisionVmAgent         bool                                                      `json:"provision_vm_agent"`
	EnableAutomaticUpgrades  bool                                                      `json:"enable_automatic_upgrades"`
	Winrm                    []AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"winrm"`
	AdditionalUnattendConfig []AzurermVirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"additional_unattend_config"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys struct {
	Path    string `json:"path"`
	KeyData string `json:"key_data"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool                                                    `json:"disable_password_authentication"`
	SshKeys                       []AzurermVirtualMachineScaleSetSpecOsProfileLinuxConfig `json:"ssh_keys"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIpAddressConfiguration struct {
	Name            string `json:"name"`
	IdleTimeout     int    `json:"idle_timeout"`
	DomainNameLabel string `json:"domain_name_label"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfiguration struct {
	Primary                                 bool                                                             `json:"primary"`
	PublicIpAddressConfiguration            []AzurermVirtualMachineScaleSetSpecNetworkProfileIpConfiguration `json:"public_ip_address_configuration"`
	Name                                    string                                                           `json:"name"`
	SubnetId                                string                                                           `json:"subnet_id"`
	ApplicationGatewayBackendAddressPoolIds []string                                                         `json:"application_gateway_backend_address_pool_ids"`
	ApplicationSecurityGroupIds             []string                                                         `json:"application_security_group_ids"`
	LoadBalancerBackendAddressPoolIds       []string                                                         `json:"load_balancer_backend_address_pool_ids"`
	LoadBalancerInboundNatRulesIds          []string                                                         `json:"load_balancer_inbound_nat_rules_ids"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfileDnsSettings struct {
	DnsServers []string `json:"dns_servers"`
}

type AzurermVirtualMachineScaleSetSpecNetworkProfile struct {
	IpConfiguration        []AzurermVirtualMachineScaleSetSpecNetworkProfile `json:"ip_configuration"`
	Name                   string                                            `json:"name"`
	Primary                bool                                              `json:"primary"`
	AcceleratedNetworking  bool                                              `json:"accelerated_networking"`
	IpForwarding           bool                                              `json:"ip_forwarding"`
	NetworkSecurityGroupId string                                            `json:"network_security_group_id"`
	DnsSettings            []AzurermVirtualMachineScaleSetSpecNetworkProfile `json:"dns_settings"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileDataDisk struct {
	DiskSizeGb      int    `json:"disk_size_gb"`
	ManagedDiskType string `json:"managed_disk_type"`
	Lun             int    `json:"lun"`
	CreateOption    string `json:"create_option"`
	Caching         string `json:"caching"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileImageReference struct {
	Sku       string `json:"sku"`
	Version   string `json:"version"`
	Id        string `json:"id"`
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
}

type AzurermVirtualMachineScaleSetSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type AzurermVirtualMachineScaleSetSpecStorageProfileOsDisk struct {
	Caching         string   `json:"caching"`
	OsType          string   `json:"os_type"`
	CreateOption    string   `json:"create_option"`
	Name            string   `json:"name"`
	Image           string   `json:"image"`
	VhdContainers   []string `json:"vhd_containers"`
	ManagedDiskType string   `json:"managed_disk_type"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates struct {
	CertificateUrl   string `json:"certificate_url"`
	CertificateStore string `json:"certificate_store"`
}

type AzurermVirtualMachineScaleSetSpecOsProfileSecrets struct {
	SourceVaultId     string                                              `json:"source_vault_id"`
	VaultCertificates []AzurermVirtualMachineScaleSetSpecOsProfileSecrets `json:"vault_certificates"`
}

type AzurermVirtualMachineScaleSetSpecExtension struct {
	ProvisionAfterExtensions []string `json:"provision_after_extensions"`
	Settings                 string   `json:"settings"`
	ProtectedSettings        string   `json:"protected_settings"`
	Name                     string   `json:"name"`
	Publisher                string   `json:"publisher"`
	Type                     string   `json:"type"`
	TypeHandlerVersion       string   `json:"type_handler_version"`
	AutoUpgradeMinorVersion  bool     `json:"auto_upgrade_minor_version"`
}

type AzurermVirtualMachineScaleSetSpec struct {
	LicenseType                  string                              `json:"license_type"`
	AutomaticOsUpgrade           bool                                `json:"automatic_os_upgrade"`
	HealthProbeId                string                              `json:"health_probe_id"`
	SinglePlacementGroup         bool                                `json:"single_placement_group"`
	OsProfile                    []AzurermVirtualMachineScaleSetSpec `json:"os_profile"`
	Plan                         []AzurermVirtualMachineScaleSetSpec `json:"plan"`
	Name                         string                              `json:"name"`
	Identity                     []AzurermVirtualMachineScaleSetSpec `json:"identity"`
	Sku                          []AzurermVirtualMachineScaleSetSpec `json:"sku"`
	UpgradePolicyMode            string                              `json:"upgrade_policy_mode"`
	Tags                         map[string]string                   `json:"tags"`
	Zones                        []string                            `json:"zones"`
	EvictionPolicy               string                              `json:"eviction_policy"`
	RollingUpgradePolicy         []AzurermVirtualMachineScaleSetSpec `json:"rolling_upgrade_policy"`
	OsProfileWindowsConfig       []AzurermVirtualMachineScaleSetSpec `json:"os_profile_windows_config"`
	OsProfileLinuxConfig         []AzurermVirtualMachineScaleSetSpec `json:"os_profile_linux_config"`
	ResourceGroupName            string                              `json:"resource_group_name"`
	Priority                     string                              `json:"priority"`
	NetworkProfile               []AzurermVirtualMachineScaleSetSpec `json:"network_profile"`
	Location                     string                              `json:"location"`
	Overprovision                bool                                `json:"overprovision"`
	StorageProfileDataDisk       []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_data_disk"`
	StorageProfileImageReference []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_image_reference"`
	BootDiagnostics              []AzurermVirtualMachineScaleSetSpec `json:"boot_diagnostics"`
	StorageProfileOsDisk         []AzurermVirtualMachineScaleSetSpec `json:"storage_profile_os_disk"`
	OsProfileSecrets             []AzurermVirtualMachineScaleSetSpec `json:"os_profile_secrets"`
	Extension                    []AzurermVirtualMachineScaleSetSpec `json:"extension"`
}

type AzurermVirtualMachineScaleSetStatus struct {
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
