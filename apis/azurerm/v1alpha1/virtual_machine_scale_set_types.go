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

type VirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetSpec   `json:"spec,omitempty"`
	Status            VirtualMachineScaleSetStatus `json:"status,omitempty"`
}

type VirtualMachineScaleSetSpecBootDiagnostics struct {
	// +optional
	Enabled    bool   `json:"enabled,omitempty"`
	StorageUri string `json:"storage_uri"`
}

type VirtualMachineScaleSetSpecExtension struct {
	// +optional
	AutoUpgradeMinorVersion bool   `json:"auto_upgrade_minor_version,omitempty"`
	Name                    string `json:"name"`
	// +optional
	ProtectedSettings string `json:"protected_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ProvisionAfterExtensions []string `json:"provision_after_extensions,omitempty"`
	Publisher                string   `json:"publisher"`
	// +optional
	Settings           string `json:"settings,omitempty"`
	Type               string `json:"type"`
	TypeHandlerVersion string `json:"type_handler_version"`
}

type VirtualMachineScaleSetSpecNetworkProfileDnsSettings struct {
	DnsServers []string `json:"dns_servers"`
}

type VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIpAddressConfiguration struct {
	DomainNameLabel string `json:"domain_name_label"`
	IdleTimeout     int    `json:"idle_timeout"`
	Name            string `json:"name"`
}

type VirtualMachineScaleSetSpecNetworkProfileIpConfiguration struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApplicationGatewayBackendAddressPoolIds []string `json:"application_gateway_backend_address_pool_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:UniqueItems=true
	ApplicationSecurityGroupIds []string `json:"application_security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LoadBalancerBackendAddressPoolIds []string `json:"load_balancer_backend_address_pool_ids,omitempty"`
	Name                              string   `json:"name"`
	Primary                           bool     `json:"primary"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PublicIpAddressConfiguration *[]VirtualMachineScaleSetSpecNetworkProfileIpConfiguration `json:"public_ip_address_configuration,omitempty"`
	SubnetId                     string                                                     `json:"subnet_id"`
}

type VirtualMachineScaleSetSpecNetworkProfile struct {
	// +optional
	AcceleratedNetworking bool `json:"accelerated_networking,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DnsSettings     *[]VirtualMachineScaleSetSpecNetworkProfile `json:"dns_settings,omitempty"`
	IpConfiguration []VirtualMachineScaleSetSpecNetworkProfile  `json:"ip_configuration"`
	// +optional
	IpForwarding bool   `json:"ip_forwarding,omitempty"`
	Name         string `json:"name"`
	// +optional
	NetworkSecurityGroupId string `json:"network_security_group_id,omitempty"`
	Primary                bool   `json:"primary"`
}

type VirtualMachineScaleSetSpecOsProfile struct {
	// +optional
	AdminPassword      string `json:"admin_password,omitempty"`
	AdminUsername      string `json:"admin_username"`
	ComputerNamePrefix string `json:"computer_name_prefix"`
	// +optional
	CustomData string `json:"custom_data,omitempty"`
}

type VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore string `json:"certificate_store,omitempty"`
	CertificateUrl   string `json:"certificate_url"`
}

type VirtualMachineScaleSetSpecOsProfileSecrets struct {
	SourceVaultId string `json:"source_vault_id"`
	// +optional
	VaultCertificates *[]VirtualMachineScaleSetSpecOsProfileSecrets `json:"vault_certificates,omitempty"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   string `json:"component"`
	Content     string `json:"content"`
	Pass        string `json:"pass"`
	SettingName string `json:"setting_name"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateUrl string `json:"certificate_url,omitempty"`
	Protocol       string `json:"protocol"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig *[]VirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"additional_unattend_config,omitempty"`
	// +optional
	EnableAutomaticUpgrades bool `json:"enable_automatic_upgrades,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provision_vm_agent,omitempty"`
	// +optional
	Winrm *[]VirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"winrm,omitempty"`
}

type VirtualMachineScaleSetSpecPlan struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
}

type VirtualMachineScaleSetSpecRollingUpgradePolicy struct {
	// +optional
	MaxBatchInstancePercent int `json:"max_batch_instance_percent,omitempty"`
	// +optional
	MaxUnhealthyInstancePercent int `json:"max_unhealthy_instance_percent,omitempty"`
	// +optional
	MaxUnhealthyUpgradedInstancePercent int `json:"max_unhealthy_upgraded_instance_percent,omitempty"`
	// +optional
	PauseTimeBetweenBatches string `json:"pause_time_between_batches,omitempty"`
}

type VirtualMachineScaleSetSpecSku struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

type VirtualMachineScaleSetSpecStorageProfileDataDisk struct {
	CreateOption string `json:"create_option"`
	Lun          int    `json:"lun"`
}

type VirtualMachineScaleSetSpecStorageProfileOsDisk struct {
	CreateOption string `json:"create_option"`
	// +optional
	Image string `json:"image,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	OsType string `json:"os_type,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VhdContainers []string `json:"vhd_containers,omitempty"`
}

type VirtualMachineScaleSetSpec struct {
	// +optional
	AutomaticOsUpgrade bool `json:"automatic_os_upgrade,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics *[]VirtualMachineScaleSetSpec `json:"boot_diagnostics,omitempty"`
	// +optional
	EvictionPolicy string `json:"eviction_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Extension *[]VirtualMachineScaleSetSpec `json:"extension,omitempty"`
	// +optional
	HealthProbeId string `json:"health_probe_id,omitempty"`
	Location      string `json:"location"`
	Name          string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	NetworkProfile []VirtualMachineScaleSetSpec `json:"network_profile"`
	// +kubebuilder:validation:MaxItems=1
	OsProfile []VirtualMachineScaleSetSpec `json:"os_profile"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OsProfileSecrets *[]VirtualMachineScaleSetSpec `json:"os_profile_secrets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileWindowsConfig *[]VirtualMachineScaleSetSpec `json:"os_profile_windows_config,omitempty"`
	// +optional
	Overprovision bool `json:"overprovision,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Plan *[]VirtualMachineScaleSetSpec `json:"plan,omitempty"`
	// +optional
	Priority          string `json:"priority,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RollingUpgradePolicy *[]VirtualMachineScaleSetSpec `json:"rolling_upgrade_policy,omitempty"`
	// +optional
	SinglePlacementGroup bool `json:"single_placement_group,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []VirtualMachineScaleSetSpec `json:"sku"`
	// +optional
	StorageProfileDataDisk *[]VirtualMachineScaleSetSpec `json:"storage_profile_data_disk,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	StorageProfileOsDisk []VirtualMachineScaleSetSpec `json:"storage_profile_os_disk"`
	UpgradePolicyMode    string                       `json:"upgrade_policy_mode"`
	// +optional
	Zones []string `json:"zones,omitempty"`
}

type VirtualMachineScaleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineScaleSetList is a list of VirtualMachineScaleSets
type VirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachineScaleSet CRD objects
	Items []VirtualMachineScaleSet `json:"items,omitempty"`
}
