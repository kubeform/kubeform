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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecAlerts struct {
	// +optional
	Cpu int `json:"cpu,omitempty" tf:"cpu,omitempty"`
	// +optional
	Io int `json:"io,omitempty" tf:"io,omitempty"`
	// +optional
	NetworkIn int `json:"networkIn,omitempty" tf:"network_in,omitempty"`
	// +optional
	NetworkOut int `json:"networkOut,omitempty" tf:"network_out,omitempty"`
	// +optional
	TransferQuota int `json:"transferQuota,omitempty" tf:"transfer_quota,omitempty"`
}

type InstanceSpecConfigDevicesSda struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdb struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdc struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdd struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSde struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdf struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdg struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevicesSdh struct {
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// +optional
	VolumeID int `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
}

type InstanceSpecConfigDevices struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sda []InstanceSpecConfigDevicesSda `json:"sda,omitempty" tf:"sda,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdb []InstanceSpecConfigDevicesSdb `json:"sdb,omitempty" tf:"sdb,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdc []InstanceSpecConfigDevicesSdc `json:"sdc,omitempty" tf:"sdc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdd []InstanceSpecConfigDevicesSdd `json:"sdd,omitempty" tf:"sdd,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sde []InstanceSpecConfigDevicesSde `json:"sde,omitempty" tf:"sde,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdf []InstanceSpecConfigDevicesSdf `json:"sdf,omitempty" tf:"sdf,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdg []InstanceSpecConfigDevicesSdg `json:"sdg,omitempty" tf:"sdg,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Sdh []InstanceSpecConfigDevicesSdh `json:"sdh,omitempty" tf:"sdh,omitempty"`
}

type InstanceSpecConfigHelpers struct {
	// +optional
	DevtmpfsAutomount bool `json:"devtmpfsAutomount,omitempty" tf:"devtmpfs_automount,omitempty"`
	// +optional
	Distro bool `json:"distro,omitempty" tf:"distro,omitempty"`
	// +optional
	ModulesDep bool `json:"modulesDep,omitempty" tf:"modules_dep,omitempty"`
	// +optional
	Network bool `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	UpdatedbDisabled bool `json:"updatedbDisabled,omitempty" tf:"updatedb_disabled,omitempty"`
}

type InstanceSpecConfig struct {
	// +optional
	Comments string `json:"comments,omitempty" tf:"comments,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Devices []InstanceSpecConfigDevices `json:"devices,omitempty" tf:"devices,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Helpers []InstanceSpecConfigHelpers `json:"helpers,omitempty" tf:"helpers,omitempty"`
	// +optional
	Kernel string `json:"kernel,omitempty" tf:"kernel,omitempty"`
	Label  string `json:"label" tf:"label"`
	// +optional
	MemoryLimit int `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`
	// +optional
	RootDevice string `json:"rootDevice,omitempty" tf:"root_device,omitempty"`
	// +optional
	RunLevel string `json:"runLevel,omitempty" tf:"run_level,omitempty"`
	// +optional
	VirtMode string `json:"virtMode,omitempty" tf:"virt_mode,omitempty"`
}

type InstanceSpecDisk struct {
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	// +optional
	Filesystem string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	Label string `json:"label" tf:"label"`
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
	// +optional
	Size int `json:"size" tf:"size"`
	// +optional
	// +optional
	StackscriptID int `json:"stackscriptID,omitempty" tf:"stackscript_id,omitempty"`
}

type InstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Alerts []InstanceSpecAlerts `json:"alerts,omitempty" tf:"alerts,omitempty"`
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	// +optional
	BackupID int `json:"backupID,omitempty" tf:"backup_id,omitempty"`
	// +optional
	BackupsEnabled bool `json:"backupsEnabled,omitempty" tf:"backups_enabled,omitempty"`
	// +optional
	BootConfigLabel string `json:"bootConfigLabel,omitempty" tf:"boot_config_label,omitempty"`
	// +optional
	Config []InstanceSpecConfig `json:"config,omitempty" tf:"config,omitempty"`
	// +optional
	Disk []InstanceSpecDisk `json:"disk,omitempty" tf:"disk,omitempty"`
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// +optional
	PrivateIP bool   `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	Region    string `json:"region" tf:"region"`
	// +optional
	// +optional
	// +optional
	StackscriptID int `json:"stackscriptID,omitempty" tf:"stackscript_id,omitempty"`
	// +optional
	SwapSize int `json:"swapSize,omitempty" tf:"swap_size,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	WatchdogEnabled bool `json:"watchdogEnabled,omitempty" tf:"watchdog_enabled,omitempty"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
