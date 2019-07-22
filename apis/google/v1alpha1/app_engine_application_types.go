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

type AppEngineApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppEngineApplicationSpec   `json:"spec,omitempty"`
	Status            AppEngineApplicationStatus `json:"status,omitempty"`
}

type AppEngineApplicationSpecFeatureSettings struct {
	// +optional
	SplitHealthChecks bool `json:"splitHealthChecks,omitempty" tf:"split_health_checks,omitempty"`
}

type AppEngineApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	AuthDomain string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FeatureSettings []AppEngineApplicationSpecFeatureSettings `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`
	LocationID      string                                    `json:"locationID" tf:"location_id"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ServingStatus string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
}

type AppEngineApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppEngineApplicationList is a list of AppEngineApplications
type AppEngineApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppEngineApplication CRD objects
	Items []AppEngineApplication `json:"items,omitempty"`
}
