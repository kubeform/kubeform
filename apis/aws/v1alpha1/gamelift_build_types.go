package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Bucket  string `json:"bucket" tf:"bucket"`
	Key     string `json:"key" tf:"key"`
	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type GameliftBuildSpec struct {
	Name            string `json:"name" tf:"name"`
	OperatingSystem string `json:"operatingSystem" tf:"operating_system"`
	// +kubebuilder:validation:MaxItems=1
	StorageLocation []GameliftBuildSpecStorageLocation `json:"storageLocation" tf:"storage_location"`
	// +optional
	Version     string                    `json:"version,omitempty" tf:"version,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GameliftBuildStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
