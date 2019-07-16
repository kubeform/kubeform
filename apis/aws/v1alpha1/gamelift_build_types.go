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

type GameliftBuild struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftBuildSpec   `json:"spec,omitempty"`
	Status            GameliftBuildStatus `json:"status,omitempty"`
}

type GameliftBuildSpecStorageLocation struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	RoleArn string `json:"role_arn"`
}

type GameliftBuildSpec struct {
	Name            string `json:"name"`
	OperatingSystem string `json:"operating_system"`
	// +kubebuilder:validation:MaxItems=1
	StorageLocation []GameliftBuildSpec `json:"storage_location"`
	// +optional
	Version string `json:"version,omitempty"`
}

type GameliftBuildStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GameliftBuildList is a list of GameliftBuilds
type GameliftBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameliftBuild CRD objects
	Items []GameliftBuild `json:"items,omitempty"`
}
