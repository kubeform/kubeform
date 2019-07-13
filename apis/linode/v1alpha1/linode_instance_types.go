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

type LinodeInstanceSpecAlerts struct {
	Cpu           int `json:"cpu"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
	Io            int `json:"io"`
}

type LinodeInstanceSpecConfigDevicesSdh struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
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
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
}

type LinodeInstanceSpecConfigDevicesSdd struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type LinodeInstanceSpecConfigDevicesSde struct {
	VolumeId  int    `json:"volume_id"`
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
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

type LinodeInstanceSpecConfigDevices struct {
	Sdh []LinodeInstanceSpecConfigDevices `json:"sdh"`
	Sda []LinodeInstanceSpecConfigDevices `json:"sda"`
	Sdb []LinodeInstanceSpecConfigDevices `json:"sdb"`
	Sdc []LinodeInstanceSpecConfigDevices `json:"sdc"`
	Sdd []LinodeInstanceSpecConfigDevices `json:"sdd"`
	Sde []LinodeInstanceSpecConfigDevices `json:"sde"`
	Sdf []LinodeInstanceSpecConfigDevices `json:"sdf"`
	Sdg []LinodeInstanceSpecConfigDevices `json:"sdg"`
}

type LinodeInstanceSpecConfigHelpers struct {
	UpdatedbDisabled  bool `json:"updatedb_disabled"`
	Distro            bool `json:"distro"`
	ModulesDep        bool `json:"modules_dep"`
	Network           bool `json:"network"`
	DevtmpfsAutomount bool `json:"devtmpfs_automount"`
}

type LinodeInstanceSpecConfig struct {
	VirtMode    string                     `json:"virt_mode"`
	RootDevice  string                     `json:"root_device"`
	Comments    string                     `json:"comments"`
	Label       string                     `json:"label"`
	Devices     []LinodeInstanceSpecConfig `json:"devices"`
	Kernel      string                     `json:"kernel"`
	Helpers     []LinodeInstanceSpecConfig `json:"helpers"`
	RunLevel    string                     `json:"run_level"`
	MemoryLimit int                        `json:"memory_limit"`
}

type LinodeInstanceSpecDisk struct {
	Label           string            `json:"label"`
	Id              int               `json:"id"`
	Image           string            `json:"image"`
	AuthorizedUsers []string          `json:"authorized_users"`
	Size            int               `json:"size"`
	Filesystem      string            `json:"filesystem"`
	ReadOnly        bool              `json:"read_only"`
	AuthorizedKeys  []string          `json:"authorized_keys"`
	StackscriptId   int               `json:"stackscript_id"`
	StackscriptData map[string]string `json:"stackscript_data"`
	RootPass        string            `json:"root_pass"`
}

type LinodeInstanceSpec struct {
	Image            string               `json:"image"`
	Tags             []string             `json:"tags"`
	Region           string               `json:"region"`
	PrivateIpAddress string               `json:"private_ip_address"`
	WatchdogEnabled  bool                 `json:"watchdog_enabled"`
	StackscriptData  map[string]string    `json:"stackscript_data"`
	Type             string               `json:"type"`
	Specs            []LinodeInstanceSpec `json:"specs"`
	Backups          []LinodeInstanceSpec `json:"backups"`
	Group            string               `json:"group"`
	Ipv6             string               `json:"ipv6"`
	Ipv4             []string             `json:"ipv4"`
	SwapSize         int                  `json:"swap_size"`
	BackupsEnabled   bool                 `json:"backups_enabled"`
	BackupId         int                  `json:"backup_id"`
	BootConfigLabel  string               `json:"boot_config_label"`
	AuthorizedUsers  []string             `json:"authorized_users"`
	RootPass         string               `json:"root_pass"`
	Alerts           []LinodeInstanceSpec `json:"alerts"`
	Status           string               `json:"status"`
	Label            string               `json:"label"`
	PrivateIp        bool                 `json:"private_ip"`
	AuthorizedKeys   []string             `json:"authorized_keys"`
	StackscriptId    int                  `json:"stackscript_id"`
	IpAddress        string               `json:"ip_address"`
	Config           []LinodeInstanceSpec `json:"config"`
	Disk             []LinodeInstanceSpec `json:"disk"`
}



type LinodeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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