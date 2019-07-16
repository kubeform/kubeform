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

type DatasyncLocationEfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncLocationEfsSpec   `json:"spec,omitempty"`
	Status            DatasyncLocationEfsStatus `json:"status,omitempty"`
}

type DatasyncLocationEfsSpecEc2Config struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupArns []string `json:"security_group_arns"`
	SubnetArn         string   `json:"subnet_arn"`
}

type DatasyncLocationEfsSpec struct {
	// +kubebuilder:validation:MaxItems=1
	Ec2Config        []DatasyncLocationEfsSpec `json:"ec2_config"`
	EfsFileSystemArn string                    `json:"efs_file_system_arn"`
	// +optional
	Subdirectory string `json:"subdirectory,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DatasyncLocationEfsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatasyncLocationEfsList is a list of DatasyncLocationEfss
type DatasyncLocationEfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatasyncLocationEfs CRD objects
	Items []DatasyncLocationEfs `json:"items,omitempty"`
}
