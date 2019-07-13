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

type AwsCognitoResourceServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoResourceServerSpec   `json:"spec,omitempty"`
	Status            AwsCognitoResourceServerStatus `json:"status,omitempty"`
}

type AwsCognitoResourceServerSpecScope struct {
	ScopeName        string `json:"scope_name"`
	ScopeDescription string `json:"scope_description"`
}

type AwsCognitoResourceServerSpec struct {
	UserPoolId       string                         `json:"user_pool_id"`
	ScopeIdentifiers []string                       `json:"scope_identifiers"`
	Identifier       string                         `json:"identifier"`
	Name             string                         `json:"name"`
	Scope            []AwsCognitoResourceServerSpec `json:"scope"`
}



type AwsCognitoResourceServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCognitoResourceServerList is a list of AwsCognitoResourceServers
type AwsCognitoResourceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoResourceServer CRD objects
	Items []AwsCognitoResourceServer `json:"items,omitempty"`
}