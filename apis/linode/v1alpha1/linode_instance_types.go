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

type DiskSpec struct {
	Label           string            `json:"label"`
	Size            int               `json:"size"`
	Id              int               `json:"id"`
	ReadOnly        bool              `json:"read_only"`
	AuthorizedKeys  []string          `json:"authorized_keys"`
	AuthorizedUsers []string          `json:"authorized_users"`
	RootPass        string            `json:"root_pass"`
	Filesystem      string            `json:"filesystem"`
	Image           string            `json:"image"`
	StackscriptId   int               `json:"stackscript_id"`
	StackscriptData map[string]string `json:"stackscript_data"`
}

type AlertsSpec struct {
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
	Io            int `json:"io"`
	Cpu           int `json:"cpu"`
	NetworkIn     int `json:"network_in"`
}

type ScheduleSpec struct {
	Day    string `json:"day"`
	Window string `json:"window"`
}

type BackupsSpec struct {
	Enabled  bool           `json:"enabled"`
	Schedule []ScheduleSpec `json:"schedule"`
}

type HelpersSpec struct {
	UpdatedbDisabled  bool `json:"updatedb_disabled"`
	Distro            bool `json:"distro"`
	ModulesDep        bool `json:"modules_dep"`
	Network           bool `json:"network"`
	DevtmpfsAutomount bool `json:"devtmpfs_automount"`
}

type SdeSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SdfSpec struct {
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
}

type SdgSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SdhSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SdaSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SdbSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SdcSpec struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type SddSpec struct {
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
}

type DevicesSpec struct {
	Sde []SdeSpec `json:"sde"`
	Sdf []SdfSpec `json:"sdf"`
	Sdg []SdgSpec `json:"sdg"`
	Sdh []SdhSpec `json:"sdh"`
	Sda []SdaSpec `json:"sda"`
	Sdb []SdbSpec `json:"sdb"`
	Sdc []SdcSpec `json:"sdc"`
	Sdd []SddSpec `json:"sdd"`
}

type ConfigSpec struct {
	Helpers     []HelpersSpec `json:"helpers"`
	Devices     []DevicesSpec `json:"devices"`
	RunLevel    string        `json:"run_level"`
	Label       string        `json:"label"`
	Kernel      string        `json:"kernel"`
	VirtMode    string        `json:"virt_mode"`
	RootDevice  string        `json:"root_device"`
	Comments    string        `json:"comments"`
	MemoryLimit int           `json:"memory_limit"`
}

type SpecsSpec struct {
	Disk     int `json:"disk"`
	Memory   int `json:"memory"`
	Vcpus    int `json:"vcpus"`
	Transfer int `json:"transfer"`
}

type LinodeInstanceSpec struct {
	Type             string            `json:"type"`
	PrivateIp        bool              `json:"private_ip"`
	Region           string            `json:"region"`
	IpAddress        string            `json:"ip_address"`
	AuthorizedUsers  []string          `json:"authorized_users"`
	Disk             []DiskSpec        `json:"disk"`
	Ipv4             []string          `json:"ipv4"`
	PrivateIpAddress string            `json:"private_ip_address"`
	RootPass         string            `json:"root_pass"`
	Group            string            `json:"group"`
	Status           string            `json:"status"`
	Alerts           []AlertsSpec      `json:"alerts"`
	StackscriptData  map[string]string `json:"stackscript_data"`
	Label            string            `json:"label"`
	Backups          []BackupsSpec     `json:"backups"`
	Ipv6             string            `json:"ipv6"`
	AuthorizedKeys   []string          `json:"authorized_keys"`
	SwapSize         int               `json:"swap_size"`
	WatchdogEnabled  bool              `json:"watchdog_enabled"`
	Tags             []string          `json:"tags"`
	Config           []ConfigSpec      `json:"config"`
	Image            string            `json:"image"`
	BackupId         int               `json:"backup_id"`
	StackscriptId    int               `json:"stackscript_id"`
	BootConfigLabel  string            `json:"boot_config_label"`
	BackupsEnabled   bool              `json:"backups_enabled"`
	Specs            []SpecsSpec       `json:"specs"`
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