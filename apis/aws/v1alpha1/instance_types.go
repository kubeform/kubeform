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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpu_credits,omitempty"`
}

type InstanceSpec struct {
	Ami string `json:"ami"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification *[]InstanceSpec `json:"credit_specification,omitempty"`
	// +optional
	DisableApiTermination bool `json:"disable_api_termination,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebs_optimized,omitempty"`
	// +optional
	GetPasswordData bool `json:"get_password_data,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iam_instance_profile,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instance_initiated_shutdown_behavior,omitempty"`
	InstanceType                      string `json:"instance_type"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty"`
	// +optional
	SourceDestCheck bool `json:"source_dest_check,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	UserData string `json:"user_data,omitempty"`
	// +optional
	UserDataBase64 string `json:"user_data_base64,omitempty"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
