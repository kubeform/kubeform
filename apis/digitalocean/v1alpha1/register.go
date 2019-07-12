package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/digitalocean"
)

var SchemeGroupVersion = schema.GroupVersion{Group: digitalocean.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
    
        &DigitaloceanDropletSnapshot{},
        &DigitaloceanDropletSnapshotList{},
    
        &DigitaloceanLoadbalancer{},
        &DigitaloceanLoadbalancerList{},
    
        &DigitaloceanProject{},
        &DigitaloceanProjectList{},
    
        &DigitaloceanDatabaseCluster{},
        &DigitaloceanDatabaseClusterList{},
    
        &DigitaloceanDroplet{},
        &DigitaloceanDropletList{},
    
        &DigitaloceanKubernetesCluster{},
        &DigitaloceanKubernetesClusterList{},
    
        &DigitaloceanRecord{},
        &DigitaloceanRecordList{},
    
        &DigitaloceanSpacesBucket{},
        &DigitaloceanSpacesBucketList{},
    
        &DigitaloceanTag{},
        &DigitaloceanTagList{},
    
        &DigitaloceanVolumeSnapshot{},
        &DigitaloceanVolumeSnapshotList{},
    
        &DigitaloceanCertificate{},
        &DigitaloceanCertificateList{},
    
        &DigitaloceanCdn{},
        &DigitaloceanCdnList{},
    
        &DigitaloceanFloatingIp{},
        &DigitaloceanFloatingIpList{},
    
        &DigitaloceanKubernetesNodePool{},
        &DigitaloceanKubernetesNodePoolList{},
    
        &DigitaloceanSshKey{},
        &DigitaloceanSshKeyList{},
    
        &DigitaloceanVolume{},
        &DigitaloceanVolumeList{},
    
        &DigitaloceanVolumeAttachment{},
        &DigitaloceanVolumeAttachmentList{},
    
        &DigitaloceanDomain{},
        &DigitaloceanDomainList{},
    
        &DigitaloceanFirewall{},
        &DigitaloceanFirewallList{},
    
        &DigitaloceanFloatingIpAssignment{},
        &DigitaloceanFloatingIpAssignmentList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
