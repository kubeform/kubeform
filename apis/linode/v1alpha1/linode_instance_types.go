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

type LinodeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeInstanceSpec   `json:"spec,omitempty"`
	Status            LinodeInstanceStatus `json:"status,omitempty"`
}

type LinodeInstanceSpecSpecs struct {
	Disk     int `json:"disk"`
	Memory   int `json:"memory"`
	Vcpus    int `json:"vcpus"`
	Transfer int `json:"transfer"`
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
	Size            int               `json:"size"`
	Filesystem      string            `json:"filesystem"`
	Image           string            `json:"image"`
	StackscriptData map[string]string `json:"stackscript_data"`
	Label           string            `json:"label"`
	ReadOnly        bool              `json:"read_only"`
	AuthorizedKeys  []string          `json:"authorized_keys"`
	AuthorizedUsers []string          `json:"authorized_users"`
	StackscriptId   int               `json:"stackscript_id"`
	RootPass        string            `json:"root_pass"`
	Id              int               `json:"id"`
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
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
}

type LinodeInstanceSpecConfigDevicesSde struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSdf struct {
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
}

type LinodeInstanceSpecConfigDevicesSdg struct {
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
}

type LinodeInstanceSpecConfigDevicesSdh struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSda struct {
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
}

type LinodeInstanceSpecConfigDevices struct {
	Sdb []LinodeInstanceSpecConfigDevices `json:"sdb"`
	Sdc []LinodeInstanceSpecConfigDevices `json:"sdc"`
	Sdd []LinodeInstanceSpecConfigDevices `json:"sdd"`
	Sde []LinodeInstanceSpecConfigDevices `json:"sde"`
	Sdf []LinodeInstanceSpecConfigDevices `json:"sdf"`
	Sdg []LinodeInstanceSpecConfigDevices `json:"sdg"`
	Sdh []LinodeInstanceSpecConfigDevices `json:"sdh"`
	Sda []LinodeInstanceSpecConfigDevices `json:"sda"`
}

type LinodeInstanceSpecConfigHelpers struct {
	Distro            bool `json:"distro"`
	ModulesDep        bool `json:"modules_dep"`
	Network           bool `json:"network"`
	DevtmpfsAutomount bool `json:"devtmpfs_automount"`
	UpdatedbDisabled  bool `json:"updatedb_disabled"`
}

type LinodeInstanceSpecConfig struct {
	Devices     []LinodeInstanceSpecConfig `json:"devices"`
	Kernel      string                     `json:"kernel"`
	VirtMode    string                     `json:"virt_mode"`
	RootDevice  string                     `json:"root_device"`
	Label       string                     `json:"label"`
	Helpers     []LinodeInstanceSpecConfig `json:"helpers"`
	RunLevel    string                     `json:"run_level"`
	Comments    string                     `json:"comments"`
	MemoryLimit int                        `json:"memory_limit"`
}

type LinodeInstanceSpecAlerts struct {
	Io            int `json:"io"`
	Cpu           int `json:"cpu"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
}

type LinodeInstanceSpec struct {
	StackscriptData  map[string]string    `json:"stackscript_data"`
	Type             string               `json:"type"`
	Ipv4             []string             `json:"ipv4"`
	Specs            []LinodeInstanceSpec `json:"specs"`
	Backups          []LinodeInstanceSpec `json:"backups"`
	StackscriptId    int                  `json:"stackscript_id"`
	BootConfigLabel  string               `json:"boot_config_label"`
	AuthorizedKeys   []string             `json:"authorized_keys"`
	WatchdogEnabled  bool                 `json:"watchdog_enabled"`
	Disk             []LinodeInstanceSpec `json:"disk"`
	PrivateIp        bool                 `json:"private_ip"`
	SwapSize         int                  `json:"swap_size"`
	Tags             []string             `json:"tags"`
	Region           string               `json:"region"`
	PrivateIpAddress string               `json:"private_ip_address"`
	Image            string               `json:"image"`
	RootPass         string               `json:"root_pass"`
	Config           []LinodeInstanceSpec `json:"config"`
	Label            string               `json:"label"`
	Group            string               `json:"group"`
	IpAddress        string               `json:"ip_address"`
	Status           string               `json:"status"`
	Ipv6             string               `json:"ipv6"`
	AuthorizedUsers  []string             `json:"authorized_users"`
	BackupId         int                  `json:"backup_id"`
	BackupsEnabled   bool                 `json:"backups_enabled"`
	Alerts           []LinodeInstanceSpec `json:"alerts"`
}

type LinodeInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeInstanceList is a list of LinodeInstances
type LinodeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeInstance CRD objects
	Items []LinodeInstance `json:"items,omitempty"`
}
