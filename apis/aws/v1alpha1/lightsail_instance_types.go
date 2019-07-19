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

type LightsailInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailInstanceSpec   `json:"spec,omitempty"`
	Status            LightsailInstanceStatus `json:"status,omitempty"`
}

type LightsailInstanceSpec struct {
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`
	BlueprintID      string `json:"blueprintID" tf:"blueprint_id"`
	BundleID         string `json:"bundleID" tf:"bundle_id"`
	// +optional
	KeyPairName string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData    string                    `json:"userData,omitempty" tf:"user_data,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LightsailInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailInstanceList is a list of LightsailInstances
type LightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailInstance CRD objects
	Items []LightsailInstance `json:"items,omitempty"`
}
