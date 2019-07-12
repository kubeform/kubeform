package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineSpecStorageOsDisk struct {
	Name                    string `json:"name"`
	VhdUri                  string `json:"vhd_uri"`
	ManagedDiskType         string `json:"managed_disk_type"`
	Caching                 string `json:"caching"`
	WriteAcceleratorEnabled bool   `json:"write_accelerator_enabled"`
	OsType                  string `json:"os_type"`
	ManagedDiskId           string `json:"managed_disk_id"`
	ImageUri                string `json:"image_uri"`
	CreateOption            string `json:"create_option"`
	DiskSizeGb              int    `json:"disk_size_gb"`
}

type AzurermVirtualMachineSpecOsProfileSecretsVaultCertificates struct {
	CertificateUrl   string `json:"certificate_url"`
	CertificateStore string `json:"certificate_store"`
}

type AzurermVirtualMachineSpecOsProfileSecrets struct {
	SourceVaultId     string                                      `json:"source_vault_id"`
	VaultCertificates []AzurermVirtualMachineSpecOsProfileSecrets `json:"vault_certificates"`
}

type AzurermVirtualMachineSpecIdentity struct {
	Type        string   `json:"type"`
	PrincipalId string   `json:"principal_id"`
	IdentityIds []string `json:"identity_ids"`
}

type AzurermVirtualMachineSpecPlan struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Product   string `json:"product"`
}

type AzurermVirtualMachineSpecStorageImageReference struct {
	Id        string `json:"id"`
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type AzurermVirtualMachineSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfigWinrm struct {
	Protocol       string `json:"protocol"`
	CertificateUrl string `json:"certificate_url"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Pass        string `json:"pass"`
	Component   string `json:"component"`
	SettingName string `json:"setting_name"`
	Content     string `json:"content"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfig struct {
	EnableAutomaticUpgrades  bool                                              `json:"enable_automatic_upgrades"`
	Timezone                 string                                            `json:"timezone"`
	Winrm                    []AzurermVirtualMachineSpecOsProfileWindowsConfig `json:"winrm"`
	AdditionalUnattendConfig []AzurermVirtualMachineSpecOsProfileWindowsConfig `json:"additional_unattend_config"`
	ProvisionVmAgent         bool                                              `json:"provision_vm_agent"`
}

type AzurermVirtualMachineSpecStorageDataDisk struct {
	CreateOption            string `json:"create_option"`
	Caching                 string `json:"caching"`
	Lun                     int    `json:"lun"`
	WriteAcceleratorEnabled bool   `json:"write_accelerator_enabled"`
	VhdUri                  string `json:"vhd_uri"`
	ManagedDiskType         string `json:"managed_disk_type"`
	DiskSizeGb              int    `json:"disk_size_gb"`
	Name                    string `json:"name"`
	ManagedDiskId           string `json:"managed_disk_id"`
}

type AzurermVirtualMachineSpecOsProfile struct {
	ComputerName  string `json:"computer_name"`
	AdminUsername string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
	CustomData    string `json:"custom_data"`
}

type AzurermVirtualMachineSpecOsProfileLinuxConfigSshKeys struct {
	Path    string `json:"path"`
	KeyData string `json:"key_data"`
}

type AzurermVirtualMachineSpecOsProfileLinuxConfig struct {
	SshKeys                       []AzurermVirtualMachineSpecOsProfileLinuxConfig `json:"ssh_keys"`
	DisablePasswordAuthentication bool                                            `json:"disable_password_authentication"`
}

type AzurermVirtualMachineSpec struct {
	Location                     string                      `json:"location"`
	ResourceGroupName            string                      `json:"resource_group_name"`
	AvailabilitySetId            string                      `json:"availability_set_id"`
	StorageOsDisk                []AzurermVirtualMachineSpec `json:"storage_os_disk"`
	DeleteOsDiskOnTermination    bool                        `json:"delete_os_disk_on_termination"`
	DeleteDataDisksOnTermination bool                        `json:"delete_data_disks_on_termination"`
	PrimaryNetworkInterfaceId    string                      `json:"primary_network_interface_id"`
	Name                         string                      `json:"name"`
	OsProfileSecrets             []AzurermVirtualMachineSpec `json:"os_profile_secrets"`
	Tags                         map[string]string           `json:"tags"`
	Identity                     []AzurermVirtualMachineSpec `json:"identity"`
	Plan                         []AzurermVirtualMachineSpec `json:"plan"`
	StorageImageReference        []AzurermVirtualMachineSpec `json:"storage_image_reference"`
	BootDiagnostics              []AzurermVirtualMachineSpec `json:"boot_diagnostics"`
	OsProfileWindowsConfig       []AzurermVirtualMachineSpec `json:"os_profile_windows_config"`
	NetworkInterfaceIds          []string                    `json:"network_interface_ids"`
	Zones                        []string                    `json:"zones"`
	VmSize                       string                      `json:"vm_size"`
	StorageDataDisk              []AzurermVirtualMachineSpec `json:"storage_data_disk"`
	OsProfile                    []AzurermVirtualMachineSpec `json:"os_profile"`
	OsProfileLinuxConfig         []AzurermVirtualMachineSpec `json:"os_profile_linux_config"`
	LicenseType                  string                      `json:"license_type"`
}

type AzurermVirtualMachineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualMachineList is a list of AzurermVirtualMachines
type AzurermVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachine CRD objects
	Items []AzurermVirtualMachine `json:"items,omitempty"`
}
