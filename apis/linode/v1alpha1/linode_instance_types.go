package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeInstanceSpec   `json:"spec,omitempty"`
	Status            LinodeInstanceStatus `json:"status,omitempty"`
}

type LinodeInstanceSpecAlerts struct {
	TransferQuota int `json:"transfer_quota"`
	Io            int `json:"io"`
	Cpu           int `json:"cpu"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
}

type LinodeInstanceSpecBackupsSchedule struct {
	Window string `json:"window"`
	Day    string `json:"day"`
}

type LinodeInstanceSpecBackups struct {
	Enabled  bool                        `json:"enabled"`
	Schedule []LinodeInstanceSpecBackups `json:"schedule"`
}

type LinodeInstanceSpecDisk struct {
	Label           string            `json:"label"`
	Filesystem      string            `json:"filesystem"`
	AuthorizedUsers []string          `json:"authorized_users"`
	StackscriptData map[string]string `json:"stackscript_data"`
	StackscriptId   int               `json:"stackscript_id"`
	RootPass        string            `json:"root_pass"`
	Size            int               `json:"size"`
	Id              int               `json:"id"`
	ReadOnly        bool              `json:"read_only"`
	Image           string            `json:"image"`
	AuthorizedKeys  []string          `json:"authorized_keys"`
}

type LinodeInstanceSpecSpecs struct {
	Vcpus    int `json:"vcpus"`
	Transfer int `json:"transfer"`
	Disk     int `json:"disk"`
	Memory   int `json:"memory"`
}

type LinodeInstanceSpecConfigHelpers struct {
	Distro            bool `json:"distro"`
	ModulesDep        bool `json:"modules_dep"`
	Network           bool `json:"network"`
	DevtmpfsAutomount bool `json:"devtmpfs_automount"`
	UpdatedbDisabled  bool `json:"updatedb_disabled"`
}

type LinodeInstanceSpecConfigDevicesSda struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdb struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdc struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdd struct {
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
}

type LinodeInstanceSpecConfigDevicesSde struct {
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
}

type LinodeInstanceSpecConfigDevicesSdf struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdg struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdh struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevices struct {
	Sda []LinodeInstanceSpecConfigDevices `json:"sda"`
	Sdb []LinodeInstanceSpecConfigDevices `json:"sdb"`
	Sdc []LinodeInstanceSpecConfigDevices `json:"sdc"`
	Sdd []LinodeInstanceSpecConfigDevices `json:"sdd"`
	Sde []LinodeInstanceSpecConfigDevices `json:"sde"`
	Sdf []LinodeInstanceSpecConfigDevices `json:"sdf"`
	Sdg []LinodeInstanceSpecConfigDevices `json:"sdg"`
	Sdh []LinodeInstanceSpecConfigDevices `json:"sdh"`
}

type LinodeInstanceSpecConfig struct {
	Helpers     []LinodeInstanceSpecConfig `json:"helpers"`
	Devices     []LinodeInstanceSpecConfig `json:"devices"`
	Kernel      string                     `json:"kernel"`
	RunLevel    string                     `json:"run_level"`
	VirtMode    string                     `json:"virt_mode"`
	RootDevice  string                     `json:"root_device"`
	MemoryLimit int                        `json:"memory_limit"`
	Label       string                     `json:"label"`
	Comments    string                     `json:"comments"`
}

type LinodeInstanceSpec struct {
	BackupsEnabled   bool                 `json:"backups_enabled"`
	Alerts           []LinodeInstanceSpec `json:"alerts"`
	Tags             []string             `json:"tags"`
	IpAddress        string               `json:"ip_address"`
	AuthorizedKeys   []string             `json:"authorized_keys"`
	SwapSize         int                  `json:"swap_size"`
	Backups          []LinodeInstanceSpec `json:"backups"`
	Disk             []LinodeInstanceSpec `json:"disk"`
	BackupId         int                  `json:"backup_id"`
	StackscriptId    int                  `json:"stackscript_id"`
	Region           string               `json:"region"`
	Specs            []LinodeInstanceSpec `json:"specs"`
	StackscriptData  map[string]string    `json:"stackscript_data"`
	AuthorizedUsers  []string             `json:"authorized_users"`
	BootConfigLabel  string               `json:"boot_config_label"`
	PrivateIp        bool                 `json:"private_ip"`
	RootPass         string               `json:"root_pass"`
	Image            string               `json:"image"`
	Type             string               `json:"type"`
	Status           string               `json:"status"`
	Ipv4             []string             `json:"ipv4"`
	Label            string               `json:"label"`
	Group            string               `json:"group"`
	PrivateIpAddress string               `json:"private_ip_address"`
	WatchdogEnabled  bool                 `json:"watchdog_enabled"`
	Ipv6             string               `json:"ipv6"`
	Config           []LinodeInstanceSpec `json:"config"`
}

type LinodeInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeInstanceList is a list of LinodeInstances
type LinodeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeInstance CRD objects
	Items []LinodeInstance `json:"items,omitempty"`
}
