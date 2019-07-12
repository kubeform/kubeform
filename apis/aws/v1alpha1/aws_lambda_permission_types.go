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

type AwsLambdaPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaPermissionSpec   `json:"spec,omitempty"`
	Status            AwsLambdaPermissionStatus `json:"status,omitempty"`
}

type AwsLambdaPermissionSpec struct {
	Qualifier         string `json:"qualifier"`
	SourceAccount     string `json:"source_account"`
	StatementIdPrefix string `json:"statement_id_prefix"`
	EventSourceToken  string `json:"event_source_token"`
	FunctionName      string `json:"function_name"`
	Principal         string `json:"principal"`
	SourceArn         string `json:"source_arn"`
	StatementId       string `json:"statement_id"`
	Action            string `json:"action"`
}

type AwsLambdaPermissionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLambdaPermissionList is a list of AwsLambdaPermissions
type AwsLambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaPermission CRD objects
	Items []AwsLambdaPermission `json:"items,omitempty"`
}
