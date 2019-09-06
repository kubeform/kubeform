package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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
	// The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.
	// +optional
	Cpu int `json:"cpu,omitempty" tf:"cpu,omitempty"`
	// The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.
	// +optional
	Io int `json:"io,omitempty" tf:"io,omitempty"`
	// The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
	// +optional
	NetworkIn int `json:"networkIn,omitempty" tf:"network_in,omitempty"`
	// The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
	// +optional
	NetworkOut int `json:"networkOut,omitempty" tf:"network_out,omitempty"`
	// The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.
	// +optional
	TransferQuota int `json:"transferQuota,omitempty" tf:"transfer_quota,omitempty"`
}

type InstanceSpecBackupsSchedule struct {
	// The day ('Sunday'-'Saturday') of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as 'Scheduling' until the day is automatically selected.
	// +optional
	Day string `json:"day,omitempty" tf:"day,omitempty"`
	// The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically.  If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected.
	// +optional
	Window string `json:"window,omitempty" tf:"window,omitempty"`
}

type InstanceSpecBackups struct {
	// If this Linode has the Backup service enabled.
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Schedule []InstanceSpecBackupsSchedule `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type InstanceSpecConfigDevicesSda struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID int `json:"diskID,omitempty" tf:"disk_id,omitempty"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel string `json:"diskLabel,omitempty" tf:"disk_label,omitempty"`
	// The Block Storage volume ID to map to this disk slot
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
	// Populates the /dev directory early during boot without udev. Defaults to false.
	// +optional
	DevtmpfsAutomount bool `json:"devtmpfsAutomount,omitempty" tf:"devtmpfs_automount,omitempty"`
	// Controls the behavior of the Linode Config's Distribution Helper setting.
	// +optional
	Distro bool `json:"distro,omitempty" tf:"distro,omitempty"`
	// Creates a modules dependency file for the Kernel you run.
	// +optional
	ModulesDep bool `json:"modulesDep,omitempty" tf:"modules_dep,omitempty"`
	// Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.
	// +optional
	Network bool `json:"network,omitempty" tf:"network,omitempty"`
	// Disables updatedb cron job to avoid disk thrashing.
	// +optional
	UpdatedbDisabled bool `json:"updatedbDisabled,omitempty" tf:"updatedb_disabled,omitempty"`
}

type InstanceSpecConfig struct {
	// Optional field for arbitrary User comments on this Config.
	// +optional
	Comments string `json:"comments,omitempty" tf:"comments,omitempty"`
	// Device sda-sdh can be either a Disk or Volume identified by disk_label or volume_id. Only one type per slot allowed.
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Devices []InstanceSpecConfigDevices `json:"devices,omitempty" tf:"devices,omitempty"`
	// Helpers enabled when booting to this Linode Config.
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Helpers []InstanceSpecConfigHelpers `json:"helpers,omitempty" tf:"helpers,omitempty"`
	// A Kernel ID to boot a Linode with. Default is based on image choice. (examples: linode/latest-64bit, linode/grub2, linode/direct-disk)
	// +optional
	Kernel string `json:"kernel,omitempty" tf:"kernel,omitempty"`
	// The Config's label for display purposes.  Also used by `boot_config_label`.
	Label string `json:"label" tf:"label"`
	// Defaults to the total RAM of the Linode
	// +optional
	MemoryLimit int `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`
	// The root device to boot. The corresponding disk must be attached.
	// +optional
	RootDevice string `json:"rootDevice,omitempty" tf:"root_device,omitempty"`
	// Defines the state of your Linode after booting. Defaults to default.
	// +optional
	RunLevel string `json:"runLevel,omitempty" tf:"run_level,omitempty"`
	// Controls the virtualization mode. Defaults to paravirt.
	// +optional
	VirtMode string `json:"virtMode,omitempty" tf:"virt_mode,omitempty"`
}

type InstanceSpecDisk struct {
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	// The Disk filesystem can be one of: raw, swap, ext3, ext4, initrd (max 32mb)
	// +optional
	Filesystem string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`
	// The ID of the Disk (for use in Linode Image resources and Linode Instance Config Devices)
	// +optional
	ID int `json:"ID,omitempty" tf:"id,omitempty"`
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/.
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// The disks label, which acts as an identifier in Terraform.
	Label string `json:"label" tf:"label"`
	// If true, this Disk is read-only.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
	// The password that will be initialially assigned to the 'root' user account.
	// +optional
	RootPass string `json:"-" sensitive:"true" tf:"root_pass,omitempty"`
	// The size of the Disk in MB.
	Size int `json:"size" tf:"size"`
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
	// +optional
	StackscriptData map[string]string `json:"-" sensitive:"true" tf:"stackscript_data,omitempty"`
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.
	// +optional
	StackscriptID int `json:"stackscriptID,omitempty" tf:"stackscript_id,omitempty"`
}

type InstanceSpecSpecs struct {
	// The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image without specifying disks.
	// +optional
	Disk int `json:"disk,omitempty" tf:"disk,omitempty"`
	// The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.
	// +optional
	Memory int `json:"memory,omitempty" tf:"memory,omitempty"`
	// The amount of network transfer this Linode is allotted each month.
	// +optional
	Transfer int `json:"transfer,omitempty" tf:"transfer,omitempty"`
	// The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.
	// +optional
	Vcpus int `json:"vcpus,omitempty" tf:"vcpus,omitempty"`
}

type InstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Alerts []InstanceSpecAlerts `json:"alerts,omitempty" tf:"alerts,omitempty"`
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	// A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive.
	// +optional
	BackupID int `json:"backupID,omitempty" tf:"backup_id,omitempty"`
	// Information about this Linode's backups status.
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Backups []InstanceSpecBackups `json:"backups,omitempty" tf:"backups,omitempty"`
	// If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
	// +optional
	BackupsEnabled bool `json:"backupsEnabled,omitempty" tf:"backups_enabled,omitempty"`
	// The Label of the Instance Config that should be used to boot the Linode instance.
	// +optional
	BootConfigLabel string `json:"bootConfigLabel,omitempty" tf:"boot_config_label,omitempty"`
	// Configuration profiles define the VM settings and boot behavior of the Linode Instance.
	// +optional
	Config []InstanceSpecConfig `json:"config,omitempty" tf:"config,omitempty"`
	// +optional
	Disk []InstanceSpecDisk `json:"disk,omitempty" tf:"disk,omitempty"`
	// The display group of the Linode instance.
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use.
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address will be used for this field.
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ipv4 []string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`
	// This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
	// +optional
	Ipv6 string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
	// The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.
	// +optional
	PrivateIP bool `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// This Linode's Private IPv4 Address.  The regional private IP address range is 192.168.128/17 address shared by all Linode Instances in a region.
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.
	Region string `json:"region" tf:"region"`
	// The password that will be initialially assigned to the 'root' user account.
	// +optional
	RootPass string `json:"-" sensitive:"true" tf:"root_pass,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Specs []InstanceSpecSpecs `json:"specs,omitempty" tf:"specs,omitempty"`
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
	// +optional
	StackscriptData map[string]string `json:"-" sensitive:"true" tf:"stackscript_data,omitempty"`
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.
	// +optional
	StackscriptID int `json:"stackscriptID,omitempty" tf:"stackscript_id,omitempty"`
	// The status of the instance, indicating the current readiness state.
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.
	// +optional
	SwapSize int `json:"swapSize,omitempty" tf:"swap_size,omitempty"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// The type of instance to be deployed, determining the price and size.
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
	// +optional
	WatchdogEnabled bool `json:"watchdogEnabled,omitempty" tf:"watchdog_enabled,omitempty"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *InstanceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
