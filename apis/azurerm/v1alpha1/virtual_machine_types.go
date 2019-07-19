package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec,omitempty"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled" tf:"enabled"`
	StorageURI string `json:"storageURI" tf:"storage_uri"`
}

type VirtualMachineSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	Type        string   `json:"type" tf:"type"`
}

type VirtualMachineSpecOsProfile struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	AdminPassword *core.LocalObjectReference `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`
	AdminUsername string                     `json:"adminUsername" tf:"admin_username"`
	ComputerName  string                     `json:"computerName" tf:"computer_name"`
	// +optional
	CustomData string `json:"customData,omitempty" tf:"custom_data,omitempty"`
}

type VirtualMachineSpecOsProfileLinuxConfigSshKeys struct {
	KeyData string `json:"keyData" tf:"key_data"`
	Path    string `json:"path" tf:"path"`
}

type VirtualMachineSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication" tf:"disable_password_authentication"`
	// +optional
	SshKeys []VirtualMachineSpecOsProfileLinuxConfigSshKeys `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
}

type VirtualMachineSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore string `json:"certificateStore,omitempty" tf:"certificate_store,omitempty"`
	CertificateURL   string `json:"certificateURL" tf:"certificate_url"`
}

type VirtualMachineSpecOsProfileSecrets struct {
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
	// +optional
	VaultCertificates []VirtualMachineSpecOsProfileSecretsVaultCertificates `json:"vaultCertificates,omitempty" tf:"vault_certificates,omitempty"`
}

type VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component string `json:"component" tf:"component"`
	// Sensitive Data. Provide secret name which contains one value only
	Content     *core.LocalObjectReference `json:"content" tf:"content"`
	Pass        string                     `json:"pass" tf:"pass"`
	SettingName string                     `json:"settingName" tf:"setting_name"`
}

type VirtualMachineSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateURL string `json:"certificateURL,omitempty" tf:"certificate_url,omitempty"`
	Protocol       string `json:"protocol" tf:"protocol"`
}

type VirtualMachineSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig []VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config,omitempty"`
	// +optional
	EnableAutomaticUpgrades bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
	// +optional
	Winrm []VirtualMachineSpecOsProfileWindowsConfigWinrm `json:"winrm,omitempty" tf:"winrm,omitempty"`
}

type VirtualMachineSpecPlan struct {
	Name      string `json:"name" tf:"name"`
	Product   string `json:"product" tf:"product"`
	Publisher string `json:"publisher" tf:"publisher"`
}

type VirtualMachineSpecStorageDataDisk struct {
	// +optional
	Caching      string `json:"caching,omitempty" tf:"caching,omitempty"`
	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	Lun        int `json:"lun" tf:"lun"`
	// +optional
	ManagedDiskID string `json:"managedDiskID,omitempty" tf:"managed_disk_id,omitempty"`
	// +optional
	ManagedDiskType string `json:"managedDiskType,omitempty" tf:"managed_disk_type,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	VhdURI string `json:"vhdURI,omitempty" tf:"vhd_uri,omitempty"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineSpecStorageImageReference struct {
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

type VirtualMachineSpecStorageOsDisk struct {
	// +optional
	Caching      string `json:"caching,omitempty" tf:"caching,omitempty"`
	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	ImageURI string `json:"imageURI,omitempty" tf:"image_uri,omitempty"`
	// +optional
	ManagedDiskID string `json:"managedDiskID,omitempty" tf:"managed_disk_id,omitempty"`
	// +optional
	ManagedDiskType string `json:"managedDiskType,omitempty" tf:"managed_disk_type,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	VhdURI string `json:"vhdURI,omitempty" tf:"vhd_uri,omitempty"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineSpec struct {
	// +optional
	AvailabilitySetID string `json:"availabilitySetID,omitempty" tf:"availability_set_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics []VirtualMachineSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`
	// +optional
	DeleteDataDisksOnTermination bool `json:"deleteDataDisksOnTermination,omitempty" tf:"delete_data_disks_on_termination,omitempty"`
	// +optional
	DeleteOsDiskOnTermination bool `json:"deleteOsDiskOnTermination,omitempty" tf:"delete_os_disk_on_termination,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []VirtualMachineSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	LicenseType         string   `json:"licenseType,omitempty" tf:"license_type,omitempty"`
	Location            string   `json:"location" tf:"location"`
	Name                string   `json:"name" tf:"name"`
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfile []VirtualMachineSpecOsProfile `json:"osProfile,omitempty" tf:"os_profile,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileLinuxConfig []VirtualMachineSpecOsProfileLinuxConfig `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config,omitempty"`
	// +optional
	OsProfileSecrets []VirtualMachineSpecOsProfileSecrets `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileWindowsConfig []VirtualMachineSpecOsProfileWindowsConfig `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan []VirtualMachineSpecPlan `json:"plan,omitempty" tf:"plan,omitempty"`
	// +optional
	PrimaryNetworkInterfaceID string `json:"primaryNetworkInterfaceID,omitempty" tf:"primary_network_interface_id,omitempty"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StorageDataDisk []VirtualMachineSpecStorageDataDisk `json:"storageDataDisk,omitempty" tf:"storage_data_disk,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	StorageImageReference []VirtualMachineSpecStorageImageReference `json:"storageImageReference,omitempty" tf:"storage_image_reference,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	StorageOsDisk []VirtualMachineSpecStorageOsDisk `json:"storageOsDisk" tf:"storage_os_disk"`
	// +optional
	Tags   map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VmSize string            `json:"vmSize" tf:"vm_size"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones       []string                  `json:"zones,omitempty" tf:"zones,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineList is a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachine CRD objects
	Items []VirtualMachine `json:"items,omitempty"`
}
