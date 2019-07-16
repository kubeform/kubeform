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

type GameliftAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftAliasSpec   `json:"spec,omitempty"`
	Status            GameliftAliasStatus `json:"status,omitempty"`
}

type GameliftAliasSpecRoutingStrategy struct {
	// +optional
	FleetId string `json:"fleet_id,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	Type    string `json:"type"`
}

type GameliftAliasSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	RoutingStrategy []GameliftAliasSpec `json:"routing_strategy"`
}

type GameliftAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GameliftAliasList is a list of GameliftAliass
type GameliftAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameliftAlias CRD objects
	Items []GameliftAlias `json:"items,omitempty"`
}
