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

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec,omitempty"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type VirtualMachineSpecOsProfile struct {
	// +optional
	AdminPassword string `json:"admin_password,omitempty"`
	AdminUsername string `json:"admin_username"`
	ComputerName  string `json:"computer_name"`
}

type VirtualMachineSpecOsProfileLinuxConfigSshKeys struct {
	KeyData string `json:"key_data"`
	Path    string `json:"path"`
}

type VirtualMachineSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool `json:"disable_password_authentication"`
	// +optional
	SshKeys *[]VirtualMachineSpecOsProfileLinuxConfig `json:"ssh_keys,omitempty"`
}

type VirtualMachineSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore string `json:"certificate_store,omitempty"`
	CertificateUrl   string `json:"certificate_url"`
}

type VirtualMachineSpecOsProfileSecrets struct {
	SourceVaultId string `json:"source_vault_id"`
	// +optional
	VaultCertificates *[]VirtualMachineSpecOsProfileSecrets `json:"vault_certificates,omitempty"`
}

type VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   string `json:"component"`
	Content     string `json:"content"`
	Pass        string `json:"pass"`
	SettingName string `json:"setting_name"`
}

type VirtualMachineSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateUrl string `json:"certificate_url,omitempty"`
	Protocol       string `json:"protocol"`
}

type VirtualMachineSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig *[]VirtualMachineSpecOsProfileWindowsConfig `json:"additional_unattend_config,omitempty"`
	// +optional
	EnableAutomaticUpgrades bool `json:"enable_automatic_upgrades,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provision_vm_agent,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
	// +optional
	Winrm *[]VirtualMachineSpecOsProfileWindowsConfig `json:"winrm,omitempty"`
}

type VirtualMachineSpecPlan struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
}

type VirtualMachineSpecStorageOsDisk struct {
	CreateOption string `json:"create_option"`
	// +optional
	ImageUri string `json:"image_uri,omitempty"`
	Name     string `json:"name"`
	// +optional
	VhdUri string `json:"vhd_uri,omitempty"`
	// +optional
	WriteAcceleratorEnabled bool `json:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics *[]VirtualMachineSpec `json:"boot_diagnostics,omitempty"`
	// +optional
	DeleteDataDisksOnTermination bool `json:"delete_data_disks_on_termination,omitempty"`
	// +optional
	DeleteOsDiskOnTermination bool     `json:"delete_os_disk_on_termination,omitempty"`
	Location                  string   `json:"location"`
	Name                      string   `json:"name"`
	NetworkInterfaceIds       []string `json:"network_interface_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfile *[]VirtualMachineSpec `json:"os_profile,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileLinuxConfig *[]VirtualMachineSpec `json:"os_profile_linux_config,omitempty"`
	// +optional
	OsProfileSecrets *[]VirtualMachineSpec `json:"os_profile_secrets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileWindowsConfig *[]VirtualMachineSpec `json:"os_profile_windows_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan *[]VirtualMachineSpec `json:"plan,omitempty"`
	// +optional
	PrimaryNetworkInterfaceId string `json:"primary_network_interface_id,omitempty"`
	ResourceGroupName         string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	StorageOsDisk []VirtualMachineSpec `json:"storage_os_disk"`
	VmSize        string               `json:"vm_size"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type VirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
