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
    
        &ComputeSSLCertificate{},
        &ComputeSSLCertificateList{},
    
        &ComputeTargetHTTPProxy{},
        &ComputeTargetHTTPProxyList{},
    
        &ComputeDisk{},
        &ComputeDiskList{},
    
        &ComputeRouterNAT{},
        &ComputeRouterNATList{},
    
        &ComputeGlobalAddress{},
        &ComputeGlobalAddressList{},
    
        &DataprocCluster{},
        &DataprocClusterList{},
    
        &LoggingOrganizationExclusion{},
        &LoggingOrganizationExclusionList{},
    
        &SpannerDatabaseIamBinding{},
        &SpannerDatabaseIamBindingList{},
    
        &KmsKeyRingIamBinding{},
        &KmsKeyRingIamBindingList{},
    
        &ComputeSubnetwork{},
        &ComputeSubnetworkList{},
    
        &StorageDefaultObjectACL{},
        &StorageDefaultObjectACLList{},
    
        &ServiceAccountIamPolicy{},
        &ServiceAccountIamPolicyList{},
    
        &ComputeAttachedDisk{},
        &ComputeAttachedDiskList{},
    
        &ComputeSecurityPolicy{},
        &ComputeSecurityPolicyList{},
    
        &StorageNotification{},
        &StorageNotificationList{},
    
        &ComputeBackendBucket{},
        &ComputeBackendBucketList{},
    
        &ComputeTargetSSLProxy{},
        &ComputeTargetSSLProxyList{},
    
        &MonitoringUptimeCheckConfig{},
        &MonitoringUptimeCheckConfigList{},
    
        &SqlSSLCert{},
        &SqlSSLCertList{},
    
        &KmsKeyRingIamPolicy{},
        &KmsKeyRingIamPolicyList{},
    
        &BillingAccountIamPolicy{},
        &BillingAccountIamPolicyList{},
    
        &ComputeSSLPolicy{},
        &ComputeSSLPolicyList{},
    
        &PubsubSubscriptionIamMember{},
        &PubsubSubscriptionIamMemberList{},
    
        &PubsubSubscription{},
        &PubsubSubscriptionList{},
    
        &SqlDatabaseInstance{},
        &SqlDatabaseInstanceList{},
    
        &ComputeInstanceFromTemplate{},
        &ComputeInstanceFromTemplateList{},
    
        &Folder{},
        &FolderList{},
    
        &LoggingProjectExclusion{},
        &LoggingProjectExclusionList{},
    
        &KmsKeyRing{},
        &KmsKeyRingList{},
    
        &ComputeInterconnectAttachment{},
        &ComputeInterconnectAttachmentList{},
    
        &ComputeRegionAutoscaler{},
        &ComputeRegionAutoscalerList{},
    
        &BillingAccountIamMember{},
        &BillingAccountIamMemberList{},
    
        &OrganizationIamMember{},
        &OrganizationIamMemberList{},
    
        &ComputeVPNGateway{},
        &ComputeVPNGatewayList{},
    
        &ComputeFirewall{},
        &ComputeFirewallList{},
    
        &Project{},
        &ProjectList{},
    
        &KmsCryptoKeyIamBinding{},
        &KmsCryptoKeyIamBindingList{},
    
        &BinaryAuthorizationPolicy{},
        &BinaryAuthorizationPolicyList{},
    
        &ComputeTargetTcpProxy{},
        &ComputeTargetTcpProxyList{},
    
        &ComputeAutoscaler{},
        &ComputeAutoscalerList{},
    
        &StorageBucketACL{},
        &StorageBucketACLList{},
    
        &LoggingOrganizationSink{},
        &LoggingOrganizationSinkList{},
    
        &ComputeBackendService{},
        &ComputeBackendServiceList{},
    
        &ContainerNodePool{},
        &ContainerNodePoolList{},
    
        &SqlDatabase{},
        &SqlDatabaseList{},
    
        &ComputeSubnetworkIamBinding{},
        &ComputeSubnetworkIamBindingList{},
    
        &ServiceAccountIamBinding{},
        &ServiceAccountIamBindingList{},
    
        &LoggingFolderSink{},
        &LoggingFolderSinkList{},
    
        &SpannerInstanceIamPolicy{},
        &SpannerInstanceIamPolicyList{},
    
        &PubsubSubscriptionIamBinding{},
        &PubsubSubscriptionIamBindingList{},
    
        &ContainerAnalysisNote{},
        &ContainerAnalysisNoteList{},
    
        &ContainerCluster{},
        &ContainerClusterList{},
    
        &SqlUser{},
        &SqlUserList{},
    
        &ComputeTargetHTTPSProxy{},
        &ComputeTargetHTTPSProxyList{},
    
        &RedisInstance{},
        &RedisInstanceList{},
    
        &PubsubSubscriptionIamPolicy{},
        &PubsubSubscriptionIamPolicyList{},
    
        &ComputeNetwork{},
        &ComputeNetworkList{},
    
        &SpannerDatabase{},
        &SpannerDatabaseList{},
    
        &ServiceAccountKey{},
        &ServiceAccountKeyList{},
    
        &BigqueryDataset{},
        &BigqueryDatasetList{},
    
        &ComputeGlobalForwardingRule{},
        &ComputeGlobalForwardingRuleList{},
    
        &SpannerDatabaseIamPolicy{},
        &SpannerDatabaseIamPolicyList{},
    
        &ComputeRoute{},
        &ComputeRouteList{},
    
        &FilestoreInstance{},
        &FilestoreInstanceList{},
    
        &SourcerepoRepository{},
        &SourcerepoRepositoryList{},
    
        &LoggingFolderExclusion{},
        &LoggingFolderExclusionList{},
    
        &DataprocJob{},
        &DataprocJobList{},
    
        &ComputeRegionDisk{},
        &ComputeRegionDiskList{},
    
        &SpannerInstanceIamMember{},
        &SpannerInstanceIamMemberList{},
    
        &ProjectUsageExportBucket{},
        &ProjectUsageExportBucketList{},
    
        &ProjectIamBinding{},
        &ProjectIamBindingList{},
    
        &StorageObjectACL{},
        &StorageObjectACLList{},
    
        &ProjectIamMember{},
        &ProjectIamMemberList{},
    
        &ComputeSubnetworkIamPolicy{},
        &ComputeSubnetworkIamPolicyList{},
    
        &AppEngineApplication{},
        &AppEngineApplicationList{},
    
        &ComposerEnvironment{},
        &ComposerEnvironmentList{},
    
        &MonitoringNotificationChannel{},
        &MonitoringNotificationChannelList{},
    
        &ProjectIamCustomRole{},
        &ProjectIamCustomRoleList{},
    
        &LoggingBillingAccountSink{},
        &LoggingBillingAccountSinkList{},
    
        &EndpointsService{},
        &EndpointsServiceList{},
    
        &CloudfunctionsFunction{},
        &CloudfunctionsFunctionList{},
    
        &BigqueryTable{},
        &BigqueryTableList{},
    
        &ResourceManagerLien{},
        &ResourceManagerLienList{},
    
        &KmsKeyRingIamMember{},
        &KmsKeyRingIamMemberList{},
    
        &ComputeInstanceTemplate{},
        &ComputeInstanceTemplateList{},
    
        &ComputeSharedVpcServiceProject{},
        &ComputeSharedVpcServiceProjectList{},
    
        &StorageBucket{},
        &StorageBucketList{},
    
        &ComputeInstanceGroup{},
        &ComputeInstanceGroupList{},
    
        &MonitoringAlertPolicy{},
        &MonitoringAlertPolicyList{},
    
        &ComputeTargetPool{},
        &ComputeTargetPoolList{},
    
        &FolderIamMember{},
        &FolderIamMemberList{},
    
        &BillingAccountIamBinding{},
        &BillingAccountIamBindingList{},
    
        &DataflowJob{},
        &DataflowJobList{},
    
        &BinaryAuthorizationAttestor{},
        &BinaryAuthorizationAttestorList{},
    
        &ComputeRouter{},
        &ComputeRouterList{},
    
        &ComputeVPNTunnel{},
        &ComputeVPNTunnelList{},
    
        &ComputeSnapshot{},
        &ComputeSnapshotList{},
    
        &ServiceAccountIamMember{},
        &ServiceAccountIamMemberList{},
    
        &FolderIamBinding{},
        &FolderIamBindingList{},
    
        &CloudiotRegistry{},
        &CloudiotRegistryList{},
    
        &ComputeInstance{},
        &ComputeInstanceList{},
    
        &MonitoringGroup{},
        &MonitoringGroupList{},
    
        &ServiceAccount{},
        &ServiceAccountList{},
    
        &ComputeRegionBackendService{},
        &ComputeRegionBackendServiceList{},
    
        &ComputeInstanceGroupManager{},
        &ComputeInstanceGroupManagerList{},
    
        &RuntimeconfigConfig{},
        &RuntimeconfigConfigList{},
    
        &BigtableTable{},
        &BigtableTableList{},
    
        &OrganizationIamBinding{},
        &OrganizationIamBindingList{},
    
        &PubsubTopic{},
        &PubsubTopicList{},
    
        &BigtableInstance{},
        &BigtableInstanceList{},
    
        &ComputeImage{},
        &ComputeImageList{},
    
        &ComputeHTTPSHealthCheck{},
        &ComputeHTTPSHealthCheckList{},
    
        &OrganizationIamCustomRole{},
        &OrganizationIamCustomRoleList{},
    
        &PubsubTopicIamPolicy{},
        &PubsubTopicIamPolicyList{},
    
        &ProjectOrganizationPolicy{},
        &ProjectOrganizationPolicyList{},
    
        &DnsManagedZone{},
        &DnsManagedZoneList{},
    
        &ComputeProjectMetadataItem{},
        &ComputeProjectMetadataItemList{},
    
        &StorageBucketIamBinding{},
        &StorageBucketIamBindingList{},
    
        &KmsCryptoKey{},
        &KmsCryptoKeyList{},
    
        &ComputeRegionInstanceGroupManager{},
        &ComputeRegionInstanceGroupManagerList{},
    
        &LoggingProjectSink{},
        &LoggingProjectSinkList{},
    
        &ComputeProjectMetadata{},
        &ComputeProjectMetadataList{},
    
        &OrganizationIamPolicy{},
        &OrganizationIamPolicyList{},
    
        &SpannerInstance{},
        &SpannerInstanceList{},
    
        &StorageDefaultObjectAccessControl{},
        &StorageDefaultObjectAccessControlList{},
    
        &FolderIamPolicy{},
        &FolderIamPolicyList{},
    
        &SpannerInstanceIamBinding{},
        &SpannerInstanceIamBindingList{},
    
        &PubsubTopicIamBinding{},
        &PubsubTopicIamBindingList{},
    
        &ComputeAddress{},
        &ComputeAddressList{},
    
        &ProjectIamPolicy{},
        &ProjectIamPolicyList{},
    
        &ProjectService{},
        &ProjectServiceList{},
    
        &ComputeSharedVpcHostProject{},
        &ComputeSharedVpcHostProjectList{},
    
        &ComputeForwardingRule{},
        &ComputeForwardingRuleList{},
    
        &DnsRecordSet{},
        &DnsRecordSetList{},
    
        &OrganizationPolicy{},
        &OrganizationPolicyList{},
    
        &ProjectServices{},
        &ProjectServicesList{},
    
        &StorageBucketIamPolicy{},
        &StorageBucketIamPolicyList{},
    
        &ComputeHealthCheck{},
        &ComputeHealthCheckList{},
    
        &ComputeRouterPeer{},
        &ComputeRouterPeerList{},
    
        &PubsubTopicIamMember{},
        &PubsubTopicIamMemberList{},
    
        &ComputeNetworkPeering{},
        &ComputeNetworkPeeringList{},
    
        &KmsCryptoKeyIamMember{},
        &KmsCryptoKeyIamMemberList{},
    
        &StorageBucketIamMember{},
        &StorageBucketIamMemberList{},
    
        &LoggingBillingAccountExclusion{},
        &LoggingBillingAccountExclusionList{},
    
        &RuntimeconfigVariable{},
        &RuntimeconfigVariableList{},
    
        &ComputeHTTPHealthCheck{},
        &ComputeHTTPHealthCheckList{},
    
        &StorageObjectAccessControl{},
        &StorageObjectAccessControlList{},
    
        &StorageBucketObject{},
        &StorageBucketObjectList{},
    
        &CloudbuildTrigger{},
        &CloudbuildTriggerList{},
    
        &FolderOrganizationPolicy{},
        &FolderOrganizationPolicyList{},
    
        &ComputeSubnetworkIamMember{},
        &ComputeSubnetworkIamMemberList{},
    
        &ComputeRouterInterface{},
        &ComputeRouterInterfaceList{},
    
        &ComputeURLMap{},
        &ComputeURLMapList{},
    
        &SpannerDatabaseIamMember{},
        &SpannerDatabaseIamMemberList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
