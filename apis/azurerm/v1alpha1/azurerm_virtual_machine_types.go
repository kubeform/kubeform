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

type AzurermVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineSpecPlan struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Product   string `json:"product"`
}

type AzurermVirtualMachineSpecStorageOsDisk struct {
	ManagedDiskType         string `json:"managed_disk_type"`
	CreateOption            string `json:"create_option"`
	DiskSizeGb              int    `json:"disk_size_gb"`
	Name                    string `json:"name"`
	VhdUri                  string `json:"vhd_uri"`
	ManagedDiskId           string `json:"managed_disk_id"`
	WriteAcceleratorEnabled bool   `json:"write_accelerator_enabled"`
	OsType                  string `json:"os_type"`
	ImageUri                string `json:"image_uri"`
	Caching                 string `json:"caching"`
}

type AzurermVirtualMachineSpecOsProfileSecretsVaultCertificates struct {
	CertificateUrl   string `json:"certificate_url"`
	CertificateStore string `json:"certificate_store"`
}

type AzurermVirtualMachineSpecOsProfileSecrets struct {
	VaultCertificates []AzurermVirtualMachineSpecOsProfileSecrets `json:"vault_certificates"`
	SourceVaultId     string                                      `json:"source_vault_id"`
}

type AzurermVirtualMachineSpecStorageDataDisk struct {
	VhdUri                  string `json:"vhd_uri"`
	ManagedDiskId           string `json:"managed_disk_id"`
	Caching                 string `json:"caching"`
	WriteAcceleratorEnabled bool   `json:"write_accelerator_enabled"`
	Name                    string `json:"name"`
	ManagedDiskType         string `json:"managed_disk_type"`
	CreateOption            string `json:"create_option"`
	DiskSizeGb              int    `json:"disk_size_gb"`
	Lun                     int    `json:"lun"`
}

type AzurermVirtualMachineSpecOsProfile struct {
	ComputerName  string `json:"computer_name"`
	AdminUsername string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
	CustomData    string `json:"custom_data"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Pass        string `json:"pass"`
	Component   string `json:"component"`
	SettingName string `json:"setting_name"`
	Content     string `json:"content"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfigWinrm struct {
	Protocol       string `json:"protocol"`
	CertificateUrl string `json:"certificate_url"`
}

type AzurermVirtualMachineSpecOsProfileWindowsConfig struct {
	AdditionalUnattendConfig []AzurermVirtualMachineSpecOsProfileWindowsConfig `json:"additional_unattend_config"`
	ProvisionVmAgent         bool                                              `json:"provision_vm_agent"`
	EnableAutomaticUpgrades  bool                                              `json:"enable_automatic_upgrades"`
	Timezone                 string                                            `json:"timezone"`
	Winrm                    []AzurermVirtualMachineSpecOsProfileWindowsConfig `json:"winrm"`
}

type AzurermVirtualMachineSpecIdentity struct {
	Type        string   `json:"type"`
	PrincipalId string   `json:"principal_id"`
	IdentityIds []string `json:"identity_ids"`
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

type AzurermVirtualMachineSpecOsProfileLinuxConfigSshKeys struct {
	Path    string `json:"path"`
	KeyData string `json:"key_data"`
}

type AzurermVirtualMachineSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool                                            `json:"disable_password_authentication"`
	SshKeys                       []AzurermVirtualMachineSpecOsProfileLinuxConfig `json:"ssh_keys"`
}

type AzurermVirtualMachineSpec struct {
	Plan                         []AzurermVirtualMachineSpec `json:"plan"`
	StorageOsDisk                []AzurermVirtualMachineSpec `json:"storage_os_disk"`
	OsProfileSecrets             []AzurermVirtualMachineSpec `json:"os_profile_secrets"`
	NetworkInterfaceIds          []string                    `json:"network_interface_ids"`
	PrimaryNetworkInterfaceId    string                      `json:"primary_network_interface_id"`
	Name                         string                      `json:"name"`
	ResourceGroupName            string                      `json:"resource_group_name"`
	Zones                        []string                    `json:"zones"`
	AvailabilitySetId            string                      `json:"availability_set_id"`
	VmSize                       string                      `json:"vm_size"`
	DeleteOsDiskOnTermination    bool                        `json:"delete_os_disk_on_termination"`
	StorageDataDisk              []AzurermVirtualMachineSpec `json:"storage_data_disk"`
	DeleteDataDisksOnTermination bool                        `json:"delete_data_disks_on_termination"`
	OsProfile                    []AzurermVirtualMachineSpec `json:"os_profile"`
	OsProfileWindowsConfig       []AzurermVirtualMachineSpec `json:"os_profile_windows_config"`
	Tags                         map[string]string           `json:"tags"`
	Location                     string                      `json:"location"`
	Identity                     []AzurermVirtualMachineSpec `json:"identity"`
	LicenseType                  string                      `json:"license_type"`
	StorageImageReference        []AzurermVirtualMachineSpec `json:"storage_image_reference"`
	BootDiagnostics              []AzurermVirtualMachineSpec `json:"boot_diagnostics"`
	OsProfileLinuxConfig         []AzurermVirtualMachineSpec `json:"os_profile_linux_config"`
}



type AzurermVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermVirtualMachineList is a list of AzurermVirtualMachines
type AzurermVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachine CRD objects
	Items []AzurermVirtualMachine `json:"items,omitempty"`
}