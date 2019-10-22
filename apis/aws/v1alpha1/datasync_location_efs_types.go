package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DatasyncLocationEfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncLocationEfsSpec   `json:"spec,omitempty"`
	Status            DatasyncLocationEfsStatus `json:"status,omitempty"`
}

type DatasyncLocationEfsSpecEc2Config struct {
	SecurityGroupArns []string `json:"securityGroupArns" tf:"security_group_arns"`
	SubnetArn         string   `json:"subnetArn" tf:"subnet_arn"`
}

type DatasyncLocationEfsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Ec2Config        []DatasyncLocationEfsSpecEc2Config `json:"ec2Config" tf:"ec2_config"`
	EfsFileSystemArn string                             `json:"efsFileSystemArn" tf:"efs_file_system_arn"`
	// +optional
	Subdirectory string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Uri string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type DatasyncLocationEfsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DatasyncLocationEfsSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
