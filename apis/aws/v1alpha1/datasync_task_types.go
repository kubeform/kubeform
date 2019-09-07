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

type DatasyncTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncTaskSpec   `json:"spec,omitempty"`
	Status            DatasyncTaskStatus `json:"status,omitempty"`
}

type DatasyncTaskSpecOptions struct {
	// +optional
	Atime string `json:"atime,omitempty" tf:"atime,omitempty"`
	// +optional
	BytesPerSecond int `json:"bytesPerSecond,omitempty" tf:"bytes_per_second,omitempty"`
	// +optional
	Gid string `json:"gid,omitempty" tf:"gid,omitempty"`
	// +optional
	Mtime string `json:"mtime,omitempty" tf:"mtime,omitempty"`
	// +optional
	PosixPermissions string `json:"posixPermissions,omitempty" tf:"posix_permissions,omitempty"`
	// +optional
	PreserveDeletedFiles string `json:"preserveDeletedFiles,omitempty" tf:"preserve_deleted_files,omitempty"`
	// +optional
	PreserveDevices string `json:"preserveDevices,omitempty" tf:"preserve_devices,omitempty"`
	// +optional
	Uid string `json:"uid,omitempty" tf:"uid,omitempty"`
	// +optional
	VerifyMode string `json:"verifyMode,omitempty" tf:"verify_mode,omitempty"`
}

type DatasyncTaskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CloudwatchLogGroupArn  string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn,omitempty"`
	DestinationLocationArn string `json:"destinationLocationArn" tf:"destination_location_arn"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Options           []DatasyncTaskSpecOptions `json:"options,omitempty" tf:"options,omitempty"`
	SourceLocationArn string                    `json:"sourceLocationArn" tf:"source_location_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DatasyncTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DatasyncTaskSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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