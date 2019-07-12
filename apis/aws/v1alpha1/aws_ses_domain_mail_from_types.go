package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainMailFrom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesDomainMailFromSpec   `json:"spec,omitempty"`
	Status            AwsSesDomainMailFromStatus `json:"status,omitempty"`
}

type AwsSesDomainMailFromSpec struct {
	Domain              string `json:"domain"`
	MailFromDomain      string `json:"mail_from_domain"`
	BehaviorOnMxFailure string `json:"behavior_on_mx_failure"`
}

type AwsSesDomainMailFromStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainMailFromList is a list of AwsSesDomainMailFroms
type AwsSesDomainMailFromList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesDomainMailFrom CRD objects
	Items []AwsSesDomainMailFrom `json:"items,omitempty"`
}
