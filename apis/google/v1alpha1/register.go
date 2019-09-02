package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/google"
)

var SchemeGroupVersion = schema.GroupVersion{Group: google.GroupName, Version: "v1alpha1"}

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
    
        &ComputeRegionAutoscaler{},
        &ComputeRegionAutoscalerList{},
    
        &OrganizationPolicy{},
        &OrganizationPolicyList{},
    
        &ComputeRouterNAT{},
        &ComputeRouterNATList{},
    
        &ProjectService{},
        &ProjectServiceList{},
    
        &ComputeSSLCertificate{},
        &ComputeSSLCertificateList{},
    
        &ProjectOrganizationPolicy{},
        &ProjectOrganizationPolicyList{},
    
        &ServiceAccountKey{},
        &ServiceAccountKeyList{},
    
        &ProjectServices{},
        &ProjectServicesList{},
    
        &ComputeProjectMetadata{},
        &ComputeProjectMetadataList{},
    
        &ComputeAddress{},
        &ComputeAddressList{},
    
        &ResourceManagerLien{},
        &ResourceManagerLienList{},
    
        &SqlDatabase{},
        &SqlDatabaseList{},
    
        &CloudiotRegistry{},
        &CloudiotRegistryList{},
    
        &ContainerAnalysisNote{},
        &ContainerAnalysisNoteList{},
    
        &StorageBucket{},
        &StorageBucketList{},
    
        &SpannerInstanceIamPolicy{},
        &SpannerInstanceIamPolicyList{},
    
        &BigtableInstance{},
        &BigtableInstanceList{},
    
        &SqlDatabaseInstance{},
        &SqlDatabaseInstanceList{},
    
        &ComputeInstanceFromTemplate{},
        &ComputeInstanceFromTemplateList{},
    
        &PubsubSubscription{},
        &PubsubSubscriptionList{},
    
        &BillingAccountIamPolicy{},
        &BillingAccountIamPolicyList{},
    
        &RedisInstance{},
        &RedisInstanceList{},
    
        &CloudfunctionsFunction{},
        &CloudfunctionsFunctionList{},
    
        &ComputeSubnetworkIamPolicy{},
        &ComputeSubnetworkIamPolicyList{},
    
        &ComputeRegionInstanceGroupManager{},
        &ComputeRegionInstanceGroupManagerList{},
    
        &PubsubSubscriptionIamMember{},
        &PubsubSubscriptionIamMemberList{},
    
        &ComputeBackendService{},
        &ComputeBackendServiceList{},
    
        &SpannerDatabaseIamPolicy{},
        &SpannerDatabaseIamPolicyList{},
    
        &ComputeFirewall{},
        &ComputeFirewallList{},
    
        &ComputeForwardingRule{},
        &ComputeForwardingRuleList{},
    
        &ComputeRouterInterface{},
        &ComputeRouterInterfaceList{},
    
        &ServiceAccountIamMember{},
        &ServiceAccountIamMemberList{},
    
        &ComputeSharedVpcServiceProject{},
        &ComputeSharedVpcServiceProjectList{},
    
        &ComputeTargetTcpProxy{},
        &ComputeTargetTcpProxyList{},
    
        &ComposerEnvironment{},
        &ComposerEnvironmentList{},
    
        &ComputeInstanceGroup{},
        &ComputeInstanceGroupList{},
    
        &ContainerCluster{},
        &ContainerClusterList{},
    
        &StorageBucketIamPolicy{},
        &StorageBucketIamPolicyList{},
    
        &EndpointsService{},
        &EndpointsServiceList{},
    
        &ProjectIamCustomRole{},
        &ProjectIamCustomRoleList{},
    
        &KmsKeyRingIamMember{},
        &KmsKeyRingIamMemberList{},
    
        &BigqueryDataset{},
        &BigqueryDatasetList{},
    
        &StorageBucketIamMember{},
        &StorageBucketIamMemberList{},
    
        &ComputeInstanceGroupManager{},
        &ComputeInstanceGroupManagerList{},
    
        &PubsubTopicIamBinding{},
        &PubsubTopicIamBindingList{},
    
        &ComputeHTTPHealthCheck{},
        &ComputeHTTPHealthCheckList{},
    
        &RuntimeconfigVariable{},
        &RuntimeconfigVariableList{},
    
        &ServiceAccountIamBinding{},
        &ServiceAccountIamBindingList{},
    
        &ComputeRegionBackendService{},
        &ComputeRegionBackendServiceList{},
    
        &LoggingFolderExclusion{},
        &LoggingFolderExclusionList{},
    
        &SpannerInstance{},
        &SpannerInstanceList{},
    
        &PubsubTopic{},
        &PubsubTopicList{},
    
        &LoggingFolderSink{},
        &LoggingFolderSinkList{},
    
        &StorageBucketACL{},
        &StorageBucketACLList{},
    
        &SpannerDatabaseIamBinding{},
        &SpannerDatabaseIamBindingList{},
    
        &KmsCryptoKeyIamBinding{},
        &KmsCryptoKeyIamBindingList{},
    
        &OrganizationIamMember{},
        &OrganizationIamMemberList{},
    
        &StorageBucketObject{},
        &StorageBucketObjectList{},
    
        &BillingAccountIamBinding{},
        &BillingAccountIamBindingList{},
    
        &KmsKeyRingIamBinding{},
        &KmsKeyRingIamBindingList{},
    
        &ComputeAutoscaler{},
        &ComputeAutoscalerList{},
    
        &FolderIamBinding{},
        &FolderIamBindingList{},
    
        &DnsRecordSet{},
        &DnsRecordSetList{},
    
        &LoggingProjectExclusion{},
        &LoggingProjectExclusionList{},
    
        &ComputeHealthCheck{},
        &ComputeHealthCheckList{},
    
        &StorageObjectACL{},
        &StorageObjectACLList{},
    
        &LoggingBillingAccountSink{},
        &LoggingBillingAccountSinkList{},
    
        &ServiceAccountIamPolicy{},
        &ServiceAccountIamPolicyList{},
    
        &OrganizationIamCustomRole{},
        &OrganizationIamCustomRoleList{},
    
        &ProjectIamPolicy{},
        &ProjectIamPolicyList{},
    
        &BinaryAuthorizationPolicy{},
        &BinaryAuthorizationPolicyList{},
    
        &ComputeSnapshot{},
        &ComputeSnapshotList{},
    
        &KmsCryptoKeyIamMember{},
        &KmsCryptoKeyIamMemberList{},
    
        &LoggingOrganizationSink{},
        &LoggingOrganizationSinkList{},
    
        &ProjectIamMember{},
        &ProjectIamMemberList{},
    
        &PubsubTopicIamPolicy{},
        &PubsubTopicIamPolicyList{},
    
        &KmsKeyRingIamPolicy{},
        &KmsKeyRingIamPolicyList{},
    
        &LoggingProjectSink{},
        &LoggingProjectSinkList{},
    
        &BinaryAuthorizationAttestor{},
        &BinaryAuthorizationAttestorList{},
    
        &ComputeBackendBucket{},
        &ComputeBackendBucketList{},
    
        &MonitoringAlertPolicy{},
        &MonitoringAlertPolicyList{},
    
        &ComputeDisk{},
        &ComputeDiskList{},
    
        &ComputeTargetHTTPProxy{},
        &ComputeTargetHTTPProxyList{},
    
        &ComputeRouter{},
        &ComputeRouterList{},
    
        &SpannerInstanceIamMember{},
        &SpannerInstanceIamMemberList{},
    
        &ComputeRouterPeer{},
        &ComputeRouterPeerList{},
    
        &ComputeProjectMetadataItem{},
        &ComputeProjectMetadataItemList{},
    
        &ComputeRegionDisk{},
        &ComputeRegionDiskList{},
    
        &Project{},
        &ProjectList{},
    
        &OrganizationIamPolicy{},
        &OrganizationIamPolicyList{},
    
        &ComputeURLMap{},
        &ComputeURLMapList{},
    
        &DataprocJob{},
        &DataprocJobList{},
    
        &ProjectIamBinding{},
        &ProjectIamBindingList{},
    
        &ComputeNetworkPeering{},
        &ComputeNetworkPeeringList{},
    
        &Folder{},
        &FolderList{},
    
        &KmsKeyRing{},
        &KmsKeyRingList{},
    
        &ComputeVPNTunnel{},
        &ComputeVPNTunnelList{},
    
        &MonitoringGroup{},
        &MonitoringGroupList{},
    
        &KmsCryptoKey{},
        &KmsCryptoKeyList{},
    
        &DataflowJob{},
        &DataflowJobList{},
    
        &ComputeInstance{},
        &ComputeInstanceList{},
    
        &ProjectUsageExportBucket{},
        &ProjectUsageExportBucketList{},
    
        &ComputeTargetSSLProxy{},
        &ComputeTargetSSLProxyList{},
    
        &ComputeSubnetworkIamBinding{},
        &ComputeSubnetworkIamBindingList{},
    
        &AppEngineApplication{},
        &AppEngineApplicationList{},
    
        &PubsubSubscriptionIamPolicy{},
        &PubsubSubscriptionIamPolicyList{},
    
        &SqlSSLCert{},
        &SqlSSLCertList{},
    
        &RuntimeconfigConfig{},
        &RuntimeconfigConfigList{},
    
        &StorageNotification{},
        &StorageNotificationList{},
    
        &SpannerDatabaseIamMember{},
        &SpannerDatabaseIamMemberList{},
    
        &MonitoringUptimeCheckConfig{},
        &MonitoringUptimeCheckConfigList{},
    
        &ComputeVPNGateway{},
        &ComputeVPNGatewayList{},
    
        &StorageObjectAccessControl{},
        &StorageObjectAccessControlList{},
    
        &SpannerInstanceIamBinding{},
        &SpannerInstanceIamBindingList{},
    
        &ComputeInstanceTemplate{},
        &ComputeInstanceTemplateList{},
    
        &ComputeSecurityPolicy{},
        &ComputeSecurityPolicyList{},
    
        &ComputeTargetPool{},
        &ComputeTargetPoolList{},
    
        &OrganizationIamBinding{},
        &OrganizationIamBindingList{},
    
        &ComputeSSLPolicy{},
        &ComputeSSLPolicyList{},
    
        &CloudbuildTrigger{},
        &CloudbuildTriggerList{},
    
        &ComputeSharedVpcHostProject{},
        &ComputeSharedVpcHostProjectList{},
    
        &SpannerDatabase{},
        &SpannerDatabaseList{},
    
        &FolderIamPolicy{},
        &FolderIamPolicyList{},
    
        &ComputeSubnetwork{},
        &ComputeSubnetworkList{},
    
        &SqlUser{},
        &SqlUserList{},
    
        &ContainerNodePool{},
        &ContainerNodePoolList{},
    
        &SourcerepoRepository{},
        &SourcerepoRepositoryList{},
    
        &ServiceAccount{},
        &ServiceAccountList{},
    
        &BigqueryTable{},
        &BigqueryTableList{},
    
        &ComputeTargetHTTPSProxy{},
        &ComputeTargetHTTPSProxyList{},
    
        &ComputeHTTPSHealthCheck{},
        &ComputeHTTPSHealthCheckList{},
    
        &ComputeRoute{},
        &ComputeRouteList{},
    
        &FilestoreInstance{},
        &FilestoreInstanceList{},
    
        &PubsubSubscriptionIamBinding{},
        &PubsubSubscriptionIamBindingList{},
    
        &ComputeNetwork{},
        &ComputeNetworkList{},
    
        &ComputeAttachedDisk{},
        &ComputeAttachedDiskList{},
    
        &ComputeImage{},
        &ComputeImageList{},
    
        &ComputeGlobalAddress{},
        &ComputeGlobalAddressList{},
    
        &FolderIamMember{},
        &FolderIamMemberList{},
    
        &DataprocCluster{},
        &DataprocClusterList{},
    
        &FolderOrganizationPolicy{},
        &FolderOrganizationPolicyList{},
    
        &ComputeSubnetworkIamMember{},
        &ComputeSubnetworkIamMemberList{},
    
        &DnsManagedZone{},
        &DnsManagedZoneList{},
    
        &MonitoringNotificationChannel{},
        &MonitoringNotificationChannelList{},
    
        &PubsubTopicIamMember{},
        &PubsubTopicIamMemberList{},
    
        &ComputeGlobalForwardingRule{},
        &ComputeGlobalForwardingRuleList{},
    
        &BillingAccountIamMember{},
        &BillingAccountIamMemberList{},
    
        &StorageDefaultObjectACL{},
        &StorageDefaultObjectACLList{},
    
        &StorageBucketIamBinding{},
        &StorageBucketIamBindingList{},
    
        &ComputeInterconnectAttachment{},
        &ComputeInterconnectAttachmentList{},
    
        &LoggingOrganizationExclusion{},
        &LoggingOrganizationExclusionList{},
    
        &BigtableTable{},
        &BigtableTableList{},
    
        &LoggingBillingAccountExclusion{},
        &LoggingBillingAccountExclusionList{},
    
        &StorageDefaultObjectAccessControl{},
        &StorageDefaultObjectAccessControlList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
