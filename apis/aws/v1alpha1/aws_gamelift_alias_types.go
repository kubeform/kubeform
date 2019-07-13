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

type AwsGameliftAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGameliftAliasSpec   `json:"spec,omitempty"`
	Status            AwsGameliftAliasStatus `json:"status,omitempty"`
}

type AwsGameliftAliasSpecRoutingStrategy struct {
	FleetId string `json:"fleet_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type AwsGameliftAliasSpec struct {
	Description     string                 `json:"description"`
	RoutingStrategy []AwsGameliftAliasSpec `json:"routing_strategy"`
	Arn             string                 `json:"arn"`
	Name            string                 `json:"name"`
}



type AwsGameliftAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGameliftAliasList is a list of AwsGameliftAliass
type AwsGameliftAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGameliftAlias CRD objects
	Items []AwsGameliftAlias `json:"items,omitempty"`
}