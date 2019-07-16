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

type DatasyncTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncTaskSpec   `json:"spec,omitempty"`
	Status            DatasyncTaskStatus `json:"status,omitempty"`
}

type DatasyncTaskSpecOptions struct {
	// +optional
	Atime string `json:"atime,omitempty"`
	// +optional
	BytesPerSecond int `json:"bytes_per_second,omitempty"`
	// +optional
	Gid string `json:"gid,omitempty"`
	// +optional
	Mtime string `json:"mtime,omitempty"`
	// +optional
	PosixPermissions string `json:"posix_permissions,omitempty"`
	// +optional
	PreserveDeletedFiles string `json:"preserve_deleted_files,omitempty"`
	// +optional
	PreserveDevices string `json:"preserve_devices,omitempty"`
	// +optional
	Uid string `json:"uid,omitempty"`
	// +optional
	VerifyMode string `json:"verify_mode,omitempty"`
}

type DatasyncTaskSpec struct {
	// +optional
	CloudwatchLogGroupArn  string `json:"cloudwatch_log_group_arn,omitempty"`
	DestinationLocationArn string `json:"destination_location_arn"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Options           *[]DatasyncTaskSpec `json:"options,omitempty"`
	SourceLocationArn string              `json:"source_location_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DatasyncTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatasyncTaskList is a list of DatasyncTasks
type DatasyncTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatasyncTask CRD objects
	Items []DatasyncTask `json:"items,omitempty"`
}
