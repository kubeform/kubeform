package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDatasyncTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatasyncTaskSpec   `json:"spec,omitempty"`
	Status            AwsDatasyncTaskStatus `json:"status,omitempty"`
}

type AwsDatasyncTaskSpecOptions struct {
	PosixPermissions     string `json:"posix_permissions"`
	PreserveDeletedFiles string `json:"preserve_deleted_files"`
	PreserveDevices      string `json:"preserve_devices"`
	Uid                  string `json:"uid"`
	BytesPerSecond       int    `json:"bytes_per_second"`
	Gid                  string `json:"gid"`
	Mtime                string `json:"mtime"`
	Atime                string `json:"atime"`
	VerifyMode           string `json:"verify_mode"`
}

type AwsDatasyncTaskSpec struct {
	Arn                    string                `json:"arn"`
	CloudwatchLogGroupArn  string                `json:"cloudwatch_log_group_arn"`
	DestinationLocationArn string                `json:"destination_location_arn"`
	Name                   string                `json:"name"`
	Options                []AwsDatasyncTaskSpec `json:"options"`
	SourceLocationArn      string                `json:"source_location_arn"`
	Tags                   map[string]string     `json:"tags"`
}



type AwsDatasyncTaskStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDatasyncTaskList is a list of AwsDatasyncTasks
type AwsDatasyncTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatasyncTask CRD objects
	Items []AwsDatasyncTask `json:"items,omitempty"`
}