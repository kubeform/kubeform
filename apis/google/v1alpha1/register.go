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
    
        &GoogleMonitoringAlertPolicy{},
        &GoogleMonitoringAlertPolicyList{},
    
        &GoogleComputeSubnetworkIamPolicy{},
        &GoogleComputeSubnetworkIamPolicyList{},
    
        &GoogleLoggingBillingAccountSink{},
        &GoogleLoggingBillingAccountSinkList{},
    
        &GooglePubsubSubscriptionIamPolicy{},
        &GooglePubsubSubscriptionIamPolicyList{},
    
        &GoogleComputeSecurityPolicy{},
        &GoogleComputeSecurityPolicyList{},
    
        &GoogleFolder{},
        &GoogleFolderList{},
    
        &GoogleComputeVpnGateway{},
        &GoogleComputeVpnGatewayList{},
    
        &GoogleSpannerDatabase{},
        &GoogleSpannerDatabaseList{},
    
        &GoogleSpannerInstanceIamMember{},
        &GoogleSpannerInstanceIamMemberList{},
    
        &GoogleSpannerInstance{},
        &GoogleSpannerInstanceList{},
    
        &GooglePubsubTopic{},
        &GooglePubsubTopicList{},
    
        &GoogleBinaryAuthorizationAttestor{},
        &GoogleBinaryAuthorizationAttestorList{},
    
        &GoogleMonitoringNotificationChannel{},
        &GoogleMonitoringNotificationChannelList{},
    
        &GoogleComputeInstanceGroupManager{},
        &GoogleComputeInstanceGroupManagerList{},
    
        &GoogleBigtableTable{},
        &GoogleBigtableTableList{},
    
        &GoogleOrganizationIamCustomRole{},
        &GoogleOrganizationIamCustomRoleList{},
    
        &GoogleComputeForwardingRule{},
        &GoogleComputeForwardingRuleList{},
    
        &GoogleCloudfunctionsFunction{},
        &GoogleCloudfunctionsFunctionList{},
    
        &GoogleComposerEnvironment{},
        &GoogleComposerEnvironmentList{},
    
        &GoogleBigqueryTable{},
        &GoogleBigqueryTableList{},
    
        &GoogleComputeNetwork{},
        &GoogleComputeNetworkList{},
    
        &GoogleComputeUrlMap{},
        &GoogleComputeUrlMapList{},
    
        &GoogleProject{},
        &GoogleProjectList{},
    
        &GoogleStorageNotification{},
        &GoogleStorageNotificationList{},
    
        &GoogleComputeAddress{},
        &GoogleComputeAddressList{},
    
        &GoogleComputeProjectMetadata{},
        &GoogleComputeProjectMetadataList{},
    
        &GoogleProjectIamMember{},
        &GoogleProjectIamMemberList{},
    
        &GoogleSqlDatabaseInstance{},
        &GoogleSqlDatabaseInstanceList{},
    
        &GoogleOrganizationIamBinding{},
        &GoogleOrganizationIamBindingList{},
    
        &GoogleStorageBucketIamPolicy{},
        &GoogleStorageBucketIamPolicyList{},
    
        &GoogleResourceManagerLien{},
        &GoogleResourceManagerLienList{},
    
        &GoogleSpannerDatabaseIamPolicy{},
        &GoogleSpannerDatabaseIamPolicyList{},
    
        &GoogleKmsKeyRingIamMember{},
        &GoogleKmsKeyRingIamMemberList{},
    
        &GoogleSpannerInstanceIamPolicy{},
        &GoogleSpannerInstanceIamPolicyList{},
    
        &GoogleComputeVpnTunnel{},
        &GoogleComputeVpnTunnelList{},
    
        &GoogleStorageBucketIamMember{},
        &GoogleStorageBucketIamMemberList{},
    
        &GoogleComputeFirewall{},
        &GoogleComputeFirewallList{},
    
        &GoogleComputeSharedVpcServiceProject{},
        &GoogleComputeSharedVpcServiceProjectList{},
    
        &GooglePubsubSubscriptionIamBinding{},
        &GooglePubsubSubscriptionIamBindingList{},
    
        &GoogleAppEngineApplication{},
        &GoogleAppEngineApplicationList{},
    
        &GoogleKmsKeyRingIamPolicy{},
        &GoogleKmsKeyRingIamPolicyList{},
    
        &GoogleSpannerDatabaseIamMember{},
        &GoogleSpannerDatabaseIamMemberList{},
    
        &GoogleComputeSubnetworkIamBinding{},
        &GoogleComputeSubnetworkIamBindingList{},
    
        &GoogleComputeRouter{},
        &GoogleComputeRouterList{},
    
        &GoogleComputeDisk{},
        &GoogleComputeDiskList{},
    
        &GoogleStorageBucketIamBinding{},
        &GoogleStorageBucketIamBindingList{},
    
        &GoogleComputeInstance{},
        &GoogleComputeInstanceList{},
    
        &GoogleComputeHttpsHealthCheck{},
        &GoogleComputeHttpsHealthCheckList{},
    
        &GoogleStorageBucketObject{},
        &GoogleStorageBucketObjectList{},
    
        &GoogleFolderIamMember{},
        &GoogleFolderIamMemberList{},
    
        &GoogleProjectIamBinding{},
        &GoogleProjectIamBindingList{},
    
        &GoogleComputeHttpHealthCheck{},
        &GoogleComputeHttpHealthCheckList{},
    
        &GoogleSourcerepoRepository{},
        &GoogleSourcerepoRepositoryList{},
    
        &GoogleContainerCluster{},
        &GoogleContainerClusterList{},
    
        &GoogleStorageBucket{},
        &GoogleStorageBucketList{},
    
        &GoogleSqlDatabase{},
        &GoogleSqlDatabaseList{},
    
        &GoogleLoggingBillingAccountExclusion{},
        &GoogleLoggingBillingAccountExclusionList{},
    
        &GooglePubsubSubscriptionIamMember{},
        &GooglePubsubSubscriptionIamMemberList{},
    
        &GoogleComputeBackendBucket{},
        &GoogleComputeBackendBucketList{},
    
        &GoogleStorageBucketAcl{},
        &GoogleStorageBucketAclList{},
    
        &GoogleProjectOrganizationPolicy{},
        &GoogleProjectOrganizationPolicyList{},
    
        &GoogleStorageDefaultObjectAcl{},
        &GoogleStorageDefaultObjectAclList{},
    
        &GoogleComputeGlobalAddress{},
        &GoogleComputeGlobalAddressList{},
    
        &GoogleContainerAnalysisNote{},
        &GoogleContainerAnalysisNoteList{},
    
        &GoogleOrganizationIamMember{},
        &GoogleOrganizationIamMemberList{},
    
        &GoogleOrganizationPolicy{},
        &GoogleOrganizationPolicyList{},
    
        &GoogleProjectService{},
        &GoogleProjectServiceList{},
    
        &GoogleLoggingFolderSink{},
        &GoogleLoggingFolderSinkList{},
    
        &GoogleLoggingOrganizationExclusion{},
        &GoogleLoggingOrganizationExclusionList{},
    
        &GoogleProjectIamCustomRole{},
        &GoogleProjectIamCustomRoleList{},
    
        &GoogleFolderOrganizationPolicy{},
        &GoogleFolderOrganizationPolicyList{},
    
        &GoogleComputeRouterNat{},
        &GoogleComputeRouterNatList{},
    
        &GoogleComputeProjectMetadataItem{},
        &GoogleComputeProjectMetadataItemList{},
    
        &GoogleComputeTargetTcpProxy{},
        &GoogleComputeTargetTcpProxyList{},
    
        &GoogleComputeRoute{},
        &GoogleComputeRouteList{},
    
        &GoogleComputeSubnetwork{},
        &GoogleComputeSubnetworkList{},
    
        &GooglePubsubTopicIamMember{},
        &GooglePubsubTopicIamMemberList{},
    
        &GoogleComputeSslCertificate{},
        &GoogleComputeSslCertificateList{},
    
        &GoogleBigqueryDataset{},
        &GoogleBigqueryDatasetList{},
    
        &GoogleServiceAccountIamMember{},
        &GoogleServiceAccountIamMemberList{},
    
        &GoogleComputeInstanceFromTemplate{},
        &GoogleComputeInstanceFromTemplateList{},
    
        &GoogleComputeTargetHttpProxy{},
        &GoogleComputeTargetHttpProxyList{},
    
        &GoogleBillingAccountIamMember{},
        &GoogleBillingAccountIamMemberList{},
    
        &GooglePubsubSubscription{},
        &GooglePubsubSubscriptionList{},
    
        &GoogleLoggingOrganizationSink{},
        &GoogleLoggingOrganizationSinkList{},
    
        &GoogleComputeBackendService{},
        &GoogleComputeBackendServiceList{},
    
        &GoogleOrganizationIamPolicy{},
        &GoogleOrganizationIamPolicyList{},
    
        &GoogleComputeRouterPeer{},
        &GoogleComputeRouterPeerList{},
    
        &GoogleSqlSslCert{},
        &GoogleSqlSslCertList{},
    
        &GoogleKmsCryptoKey{},
        &GoogleKmsCryptoKeyList{},
    
        &GoogleLoggingProjectSink{},
        &GoogleLoggingProjectSinkList{},
    
        &GoogleServiceAccountIamPolicy{},
        &GoogleServiceAccountIamPolicyList{},
    
        &GooglePubsubTopicIamPolicy{},
        &GooglePubsubTopicIamPolicyList{},
    
        &GoogleBillingAccountIamPolicy{},
        &GoogleBillingAccountIamPolicyList{},
    
        &GoogleSpannerInstanceIamBinding{},
        &GoogleSpannerInstanceIamBindingList{},
    
        &GoogleComputeHealthCheck{},
        &GoogleComputeHealthCheckList{},
    
        &GoogleComputeInterconnectAttachment{},
        &GoogleComputeInterconnectAttachmentList{},
    
        &GoogleBigtableInstance{},
        &GoogleBigtableInstanceList{},
    
        &GoogleComputeSharedVpcHostProject{},
        &GoogleComputeSharedVpcHostProjectList{},
    
        &GoogleComputeInstanceTemplate{},
        &GoogleComputeInstanceTemplateList{},
    
        &GoogleDataprocJob{},
        &GoogleDataprocJobList{},
    
        &GoogleDnsRecordSet{},
        &GoogleDnsRecordSetList{},
    
        &GoogleProjectUsageExportBucket{},
        &GoogleProjectUsageExportBucketList{},
    
        &GoogleKmsKeyRing{},
        &GoogleKmsKeyRingList{},
    
        &GoogleComputeAutoscaler{},
        &GoogleComputeAutoscalerList{},
    
        &GoogleDataprocCluster{},
        &GoogleDataprocClusterList{},
    
        &GoogleSqlUser{},
        &GoogleSqlUserList{},
    
        &GoogleProjectServices{},
        &GoogleProjectServicesList{},
    
        &GooglePubsubTopicIamBinding{},
        &GooglePubsubTopicIamBindingList{},
    
        &GoogleKmsCryptoKeyIamMember{},
        &GoogleKmsCryptoKeyIamMemberList{},
    
        &GoogleComputeSslPolicy{},
        &GoogleComputeSslPolicyList{},
    
        &GoogleComputeRegionAutoscaler{},
        &GoogleComputeRegionAutoscalerList{},
    
        &GoogleComputeAttachedDisk{},
        &GoogleComputeAttachedDiskList{},
    
        &GoogleFolderIamPolicy{},
        &GoogleFolderIamPolicyList{},
    
        &GoogleSpannerDatabaseIamBinding{},
        &GoogleSpannerDatabaseIamBindingList{},
    
        &GoogleComputeGlobalForwardingRule{},
        &GoogleComputeGlobalForwardingRuleList{},
    
        &GoogleServiceAccountIamBinding{},
        &GoogleServiceAccountIamBindingList{},
    
        &GoogleProjectIamPolicy{},
        &GoogleProjectIamPolicyList{},
    
        &GoogleComputeRegionDisk{},
        &GoogleComputeRegionDiskList{},
    
        &GoogleMonitoringGroup{},
        &GoogleMonitoringGroupList{},
    
        &GoogleComputeImage{},
        &GoogleComputeImageList{},
    
        &GoogleComputeRouterInterface{},
        &GoogleComputeRouterInterfaceList{},
    
        &GoogleComputeTargetPool{},
        &GoogleComputeTargetPoolList{},
    
        &GoogleDataflowJob{},
        &GoogleDataflowJobList{},
    
        &GoogleStorageObjectAcl{},
        &GoogleStorageObjectAclList{},
    
        &GoogleEndpointsService{},
        &GoogleEndpointsServiceList{},
    
        &GoogleCloudbuildTrigger{},
        &GoogleCloudbuildTriggerList{},
    
        &GoogleServiceAccount{},
        &GoogleServiceAccountList{},
    
        &GoogleRuntimeconfigConfig{},
        &GoogleRuntimeconfigConfigList{},
    
        &GoogleStorageObjectAccessControl{},
        &GoogleStorageObjectAccessControlList{},
    
        &GoogleMonitoringUptimeCheckConfig{},
        &GoogleMonitoringUptimeCheckConfigList{},
    
        &GoogleBillingAccountIamBinding{},
        &GoogleBillingAccountIamBindingList{},
    
        &GoogleDnsManagedZone{},
        &GoogleDnsManagedZoneList{},
    
        &GoogleComputeSnapshot{},
        &GoogleComputeSnapshotList{},
    
        &GoogleComputeRegionBackendService{},
        &GoogleComputeRegionBackendServiceList{},
    
        &GoogleComputeNetworkPeering{},
        &GoogleComputeNetworkPeeringList{},
    
        &GoogleServiceAccountKey{},
        &GoogleServiceAccountKeyList{},
    
        &GoogleKmsCryptoKeyIamBinding{},
        &GoogleKmsCryptoKeyIamBindingList{},
    
        &GoogleComputeInstanceGroup{},
        &GoogleComputeInstanceGroupList{},
    
        &GoogleKmsKeyRingIamBinding{},
        &GoogleKmsKeyRingIamBindingList{},
    
        &GoogleBinaryAuthorizationPolicy{},
        &GoogleBinaryAuthorizationPolicyList{},
    
        &GoogleComputeTargetSslProxy{},
        &GoogleComputeTargetSslProxyList{},
    
        &GoogleRedisInstance{},
        &GoogleRedisInstanceList{},
    
        &GoogleLoggingProjectExclusion{},
        &GoogleLoggingProjectExclusionList{},
    
        &GoogleStorageDefaultObjectAccessControl{},
        &GoogleStorageDefaultObjectAccessControlList{},
    
        &GoogleFolderIamBinding{},
        &GoogleFolderIamBindingList{},
    
        &GoogleComputeRegionInstanceGroupManager{},
        &GoogleComputeRegionInstanceGroupManagerList{},
    
        &GoogleRuntimeconfigVariable{},
        &GoogleRuntimeconfigVariableList{},
    
        &GoogleContainerNodePool{},
        &GoogleContainerNodePoolList{},
    
        &GoogleLoggingFolderExclusion{},
        &GoogleLoggingFolderExclusionList{},
    
        &GoogleComputeTargetHttpsProxy{},
        &GoogleComputeTargetHttpsProxyList{},
    
        &GoogleFilestoreInstance{},
        &GoogleFilestoreInstanceList{},
    
        &GoogleComputeSubnetworkIamMember{},
        &GoogleComputeSubnetworkIamMemberList{},
    
        &GoogleCloudiotRegistry{},
        &GoogleCloudiotRegistryList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
