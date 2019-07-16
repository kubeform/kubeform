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

type LambdaPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaPermissionSpec   `json:"spec,omitempty"`
	Status            LambdaPermissionStatus `json:"status,omitempty"`
}

type LambdaPermissionSpec struct {
	Action string `json:"action"`
	// +optional
	EventSourceToken string `json:"event_source_token,omitempty"`
	FunctionName     string `json:"function_name"`
	Principal        string `json:"principal"`
	// +optional
	Qualifier string `json:"qualifier,omitempty"`
	// +optional
	SourceAccount string `json:"source_account,omitempty"`
	// +optional
	SourceArn string `json:"source_arn,omitempty"`
	// +optional
	StatementIdPrefix string `json:"statement_id_prefix,omitempty"`
}

type LambdaPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaPermissionList is a list of LambdaPermissions
type LambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaPermission CRD objects
	Items []LambdaPermission `json:"items,omitempty"`
}
