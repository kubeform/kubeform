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

type AwsGameliftBuild struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGameliftBuildSpec   `json:"spec,omitempty"`
	Status            AwsGameliftBuildStatus `json:"status,omitempty"`
}

type AwsGameliftBuildSpecStorageLocation struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	RoleArn string `json:"role_arn"`
}

type AwsGameliftBuildSpec struct {
	OperatingSystem string                 `json:"operating_system"`
	StorageLocation []AwsGameliftBuildSpec `json:"storage_location"`
	Version         string                 `json:"version"`
	Name            string                 `json:"name"`
}



type AwsGameliftBuildStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGameliftBuildList is a list of AwsGameliftBuilds
type AwsGameliftBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGameliftBuild CRD objects
	Items []AwsGameliftBuild `json:"items,omitempty"`
}