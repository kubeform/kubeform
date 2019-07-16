package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LambdaAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaAliasSpec   `json:"spec,omitempty"`
	Status            LambdaAliasStatus `json:"status,omitempty"`
}

type LambdaAliasSpecRoutingConfig struct {
	// +optional
	AdditionalVersionWeights map[string]json.Number `json:"additional_version_weights,omitempty"`
}

type LambdaAliasSpec struct {
	// +optional
	Description     string `json:"description,omitempty"`
	FunctionName    string `json:"function_name"`
	FunctionVersion string `json:"function_version"`
	Name            string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RoutingConfig *[]LambdaAliasSpec `json:"routing_config,omitempty"`
}

type LambdaAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaAliasList is a list of LambdaAliass
type LambdaAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaAlias CRD objects
	Items []LambdaAlias `json:"items,omitempty"`
}
