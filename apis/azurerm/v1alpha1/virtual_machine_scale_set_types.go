package v1alpha1

import (
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

type VirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetSpec   `json:"spec,omitempty"`
	Status            VirtualMachineScaleSetStatus `json:"status,omitempty"`
}

type VirtualMachineScaleSetSpecBootDiagnostics struct {
	// +optional
	Enabled    bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	StorageURI string `json:"storageURI" tf:"storage_uri"`
}

type VirtualMachineScaleSetSpecExtension struct {
	// +optional
	AutoUpgradeMinorVersion bool   `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`
	Name                    string `json:"name" tf:"name"`
	// +optional
	ProtectedSettings string `json:"-" sensitive:"true" tf:"protected_settings,omitempty"`
	// +optional
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions,omitempty"`
	Publisher                string   `json:"publisher" tf:"publisher"`
	// +optional
	Settings           string `json:"settings,omitempty" tf:"settings,omitempty"`
	Type               string `json:"type" tf:"type"`
	TypeHandlerVersion string `json:"typeHandlerVersion" tf:"type_handler_version"`
}

type VirtualMachineScaleSetSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	Type        string `json:"type" tf:"type"`
}

type VirtualMachineScaleSetSpecNetworkProfileDnsSettings struct {
	DnsServers []string `json:"dnsServers" tf:"dns_servers"`
}

type VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration struct {
	DomainNameLabel string `json:"domainNameLabel" tf:"domain_name_label"`
	IdleTimeout     int    `json:"idleTimeout" tf:"idle_timeout"`
	Name            string `json:"name" tf:"name"`
}

type VirtualMachineScaleSetSpecNetworkProfileIpConfiguration struct {
	// +optional
	ApplicationGatewayBackendAddressPoolIDS []string `json:"applicationGatewayBackendAddressPoolIDS,omitempty" tf:"application_gateway_backend_address_pool_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ApplicationSecurityGroupIDS []string `json:"applicationSecurityGroupIDS,omitempty" tf:"application_security_group_ids,omitempty"`
	// +optional
	LoadBalancerBackendAddressPoolIDS []string `json:"loadBalancerBackendAddressPoolIDS,omitempty" tf:"load_balancer_backend_address_pool_ids,omitempty"`
	// +optional
	LoadBalancerInboundNATRulesIDS []string `json:"loadBalancerInboundNATRulesIDS,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`
	Name                           string   `json:"name" tf:"name"`
	Primary                        bool     `json:"primary" tf:"primary"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PublicIPAddressConfiguration []VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration `json:"publicIPAddressConfiguration,omitempty" tf:"public_ip_address_configuration,omitempty"`
	SubnetID                     string                                                                                `json:"subnetID" tf:"subnet_id"`
}

type VirtualMachineScaleSetSpecNetworkProfile struct {
	// +optional
	AcceleratedNetworking bool `json:"acceleratedNetworking,omitempty" tf:"accelerated_networking,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DnsSettings     []VirtualMachineScaleSetSpecNetworkProfileDnsSettings     `json:"dnsSettings,omitempty" tf:"dns_settings,omitempty"`
	IpConfiguration []VirtualMachineScaleSetSpecNetworkProfileIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	// +optional
	IpForwarding bool   `json:"ipForwarding,omitempty" tf:"ip_forwarding,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id,omitempty"`
	Primary                bool   `json:"primary" tf:"primary"`
}

type VirtualMachineScaleSetSpecOsProfile struct {
	// +optional
	AdminPassword      string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	AdminUsername      string `json:"adminUsername" tf:"admin_username"`
	ComputerNamePrefix string `json:"computerNamePrefix" tf:"computer_name_prefix"`
	// +optional
	CustomData string `json:"customData,omitempty" tf:"custom_data,omitempty"`
}

type VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys struct {
	// +optional
	KeyData string `json:"keyData,omitempty" tf:"key_data,omitempty"`
	Path    string `json:"path" tf:"path"`
}

type VirtualMachineScaleSetSpecOsProfileLinuxConfig struct {
	// +optional
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`
	// +optional
	SshKeys []VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
}

type VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore string `json:"certificateStore,omitempty" tf:"certificate_store,omitempty"`
	CertificateURL   string `json:"certificateURL" tf:"certificate_url"`
}

type VirtualMachineScaleSetSpecOsProfileSecrets struct {
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
	// +optional
	VaultCertificates []VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates `json:"vaultCertificates,omitempty" tf:"vault_certificates,omitempty"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   string `json:"component" tf:"component"`
	Content     string `json:"-" sensitive:"true" tf:"content"`
	Pass        string `json:"pass" tf:"pass"`
	SettingName string `json:"settingName" tf:"setting_name"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateURL string `json:"certificateURL,omitempty" tf:"certificate_url,omitempty"`
	Protocol       string `json:"protocol" tf:"protocol"`
}

type VirtualMachineScaleSetSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig []VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config,omitempty"`
	// +optional
	EnableAutomaticUpgrades bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`
	// +optional
	Winrm []VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm `json:"winrm,omitempty" tf:"winrm,omitempty"`
}

type VirtualMachineScaleSetSpecPlan struct {
	Name      string `json:"name" tf:"name"`
	Product   string `json:"product" tf:"product"`
	Publisher string `json:"publisher" tf:"publisher"`
}

type VirtualMachineScaleSetSpecRollingUpgradePolicy struct {
	// +optional
	MaxBatchInstancePercent int `json:"maxBatchInstancePercent,omitempty" tf:"max_batch_instance_percent,omitempty"`
	// +optional
	MaxUnhealthyInstancePercent int `json:"maxUnhealthyInstancePercent,omitempty" tf:"max_unhealthy_instance_percent,omitempty"`
	// +optional
	MaxUnhealthyUpgradedInstancePercent int `json:"maxUnhealthyUpgradedInstancePercent,omitempty" tf:"max_unhealthy_upgraded_instance_percent,omitempty"`
	// +optional
	PauseTimeBetweenBatches string `json:"pauseTimeBetweenBatches,omitempty" tf:"pause_time_between_batches,omitempty"`
}

type VirtualMachineScaleSetSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type VirtualMachineScaleSetSpecStorageProfileDataDisk struct {
	// +optional
	Caching      string `json:"caching,omitempty" tf:"caching,omitempty"`
	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	Lun        int `json:"lun" tf:"lun"`
	// +optional
	ManagedDiskType string `json:"managedDiskType,omitempty" tf:"managed_disk_type,omitempty"`
}

type VirtualMachineScaleSetSpecStorageProfileImageReference struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Offer string `json:"offer,omitempty" tf:"offer,omitempty"`
	// +optional
	Publisher string `json:"publisher,omitempty" tf:"publisher,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type VirtualMachineScaleSetSpecStorageProfileOsDisk struct {
	// +optional
	Caching      string `json:"caching,omitempty" tf:"caching,omitempty"`
	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	ManagedDiskType string `json:"managedDiskType,omitempty" tf:"managed_disk_type,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	VhdContainers []string `json:"vhdContainers,omitempty" tf:"vhd_containers,omitempty"`
}

type VirtualMachineScaleSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutomaticOsUpgrade bool `json:"automaticOsUpgrade,omitempty" tf:"automatic_os_upgrade,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics []VirtualMachineScaleSetSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`
	// +optional
	EvictionPolicy string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`
	// +optional
	Extension []VirtualMachineScaleSetSpecExtension `json:"extension,omitempty" tf:"extension,omitempty"`
	// +optional
	HealthProbeID string `json:"healthProbeID,omitempty" tf:"health_probe_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []VirtualMachineScaleSetSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	LicenseType    string                                     `json:"licenseType,omitempty" tf:"license_type,omitempty"`
	Location       string                                     `json:"location" tf:"location"`
	Name           string                                     `json:"name" tf:"name"`
	NetworkProfile []VirtualMachineScaleSetSpecNetworkProfile `json:"networkProfile" tf:"network_profile"`
	// +kubebuilder:validation:MaxItems=1
	OsProfile []VirtualMachineScaleSetSpecOsProfile `json:"osProfile" tf:"os_profile"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OsProfileLinuxConfig []VirtualMachineScaleSetSpecOsProfileLinuxConfig `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config,omitempty"`
	// +optional
	OsProfileSecrets []VirtualMachineScaleSetSpecOsProfileSecrets `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OsProfileWindowsConfig []VirtualMachineScaleSetSpecOsProfileWindowsConfig `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config,omitempty"`
	// +optional
	Overprovision bool `json:"overprovision,omitempty" tf:"overprovision,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan []VirtualMachineScaleSetSpecPlan `json:"plan,omitempty" tf:"plan,omitempty"`
	// +optional
	Priority          string `json:"priority,omitempty" tf:"priority,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RollingUpgradePolicy []VirtualMachineScaleSetSpecRollingUpgradePolicy `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy,omitempty"`
	// +optional
	SinglePlacementGroup bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []VirtualMachineScaleSetSpecSku `json:"sku" tf:"sku"`
	// +optional
	StorageProfileDataDisk []VirtualMachineScaleSetSpecStorageProfileDataDisk `json:"storageProfileDataDisk,omitempty" tf:"storage_profile_data_disk,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageProfileImageReference []VirtualMachineScaleSetSpecStorageProfileImageReference `json:"storageProfileImageReference,omitempty" tf:"storage_profile_image_reference,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfileOsDisk []VirtualMachineScaleSetSpecStorageProfileOsDisk `json:"storageProfileOsDisk" tf:"storage_profile_os_disk"`
	// +optional
	Tags              map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	UpgradePolicyMode string            `json:"upgradePolicyMode" tf:"upgrade_policy_mode"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type VirtualMachineScaleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualMachineScaleSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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