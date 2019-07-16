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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecConfig struct {
	// +optional
	Comments string `json:"comments,omitempty"`
	// +optional
	Kernel string `json:"kernel,omitempty"`
	Label  string `json:"label"`
	// +optional
	MemoryLimit int `json:"memory_limit,omitempty"`
	// +optional
	RunLevel string `json:"run_level,omitempty"`
	// +optional
	VirtMode string `json:"virt_mode,omitempty"`
}

type InstanceSpecDisk struct {
	// +optional
	AuthorizedKeys []string `json:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorized_users,omitempty"`
	Label           string   `json:"label"`
	// +optional
	RootPass string `json:"root_pass,omitempty"`
	Size     int    `json:"size"`
}

type InstanceSpec struct {
	// +optional
	AuthorizedKeys []string `json:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorized_users,omitempty"`
	// +optional
	BackupId int `json:"backup_id,omitempty"`
	// +optional
	Config *[]InstanceSpec `json:"config,omitempty"`
	// +optional
	Disk *[]InstanceSpec `json:"disk,omitempty"`
	// +optional
	Group string `json:"group,omitempty"`
	// +optional
	Image string `json:"image,omitempty"`
	// +optional
	PrivateIp bool   `json:"private_ip,omitempty"`
	Region    string `json:"region"`
	// +optional
	RootPass string `json:"root_pass,omitempty"`
	// +optional
	StackscriptData map[string]string `json:"stackscript_data,omitempty"`
	// +optional
	StackscriptId int `json:"stackscript_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	WatchdogEnabled bool `json:"watchdog_enabled,omitempty"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
