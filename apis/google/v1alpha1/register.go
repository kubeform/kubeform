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

		&GoogleComputeSubnetworkIamPolicy{},
		&GoogleComputeSubnetworkIamPolicyList{},

		&GoogleStorageBucket{},
		&GoogleStorageBucketList{},

		&GoogleDataflowJob{},
		&GoogleDataflowJobList{},

		&GoogleOrganizationPolicy{},
		&GoogleOrganizationPolicyList{},

		&GoogleDnsManagedZone{},
		&GoogleDnsManagedZoneList{},

		&GoogleComputeNetwork{},
		&GoogleComputeNetworkList{},

		&GoogleSpannerDatabase{},
		&GoogleSpannerDatabaseList{},

		&GoogleLoggingOrganizationSink{},
		&GoogleLoggingOrganizationSinkList{},

		&GoogleCloudiotRegistry{},
		&GoogleCloudiotRegistryList{},

		&GoogleComputeHealthCheck{},
		&GoogleComputeHealthCheckList{},

		&GoogleCloudbuildTrigger{},
		&GoogleCloudbuildTriggerList{},

		&GoogleComputeImage{},
		&GoogleComputeImageList{},

		&GoogleEndpointsService{},
		&GoogleEndpointsServiceList{},

		&GoogleContainerCluster{},
		&GoogleContainerClusterList{},

		&GoogleServiceAccountKey{},
		&GoogleServiceAccountKeyList{},

		&GoogleMonitoringAlertPolicy{},
		&GoogleMonitoringAlertPolicyList{},

		&GoogleBigtableTable{},
		&GoogleBigtableTableList{},

		&GoogleBillingAccountIamMember{},
		&GoogleBillingAccountIamMemberList{},

		&GoogleSpannerInstanceIamMember{},
		&GoogleSpannerInstanceIamMemberList{},

		&GoogleSpannerDatabaseIamPolicy{},
		&GoogleSpannerDatabaseIamPolicyList{},

		&GoogleComputeSubnetworkIamMember{},
		&GoogleComputeSubnetworkIamMemberList{},

		&GoogleComputeTargetHttpProxy{},
		&GoogleComputeTargetHttpProxyList{},

		&GoogleLoggingFolderSink{},
		&GoogleLoggingFolderSinkList{},

		&GoogleProjectIamCustomRole{},
		&GoogleProjectIamCustomRoleList{},

		&GoogleBigqueryDataset{},
		&GoogleBigqueryDatasetList{},

		&GoogleStorageBucketIamPolicy{},
		&GoogleStorageBucketIamPolicyList{},

		&GoogleSqlDatabaseInstance{},
		&GoogleSqlDatabaseInstanceList{},

		&GoogleDnsRecordSet{},
		&GoogleDnsRecordSetList{},

		&GoogleComputeVpnGateway{},
		&GoogleComputeVpnGatewayList{},

		&GoogleComputeVpnTunnel{},
		&GoogleComputeVpnTunnelList{},

		&GoogleComputeHttpHealthCheck{},
		&GoogleComputeHttpHealthCheckList{},

		&GoogleServiceAccountIamPolicy{},
		&GoogleServiceAccountIamPolicyList{},

		&GooglePubsubTopicIamPolicy{},
		&GooglePubsubTopicIamPolicyList{},

		&GoogleLoggingBillingAccountExclusion{},
		&GoogleLoggingBillingAccountExclusionList{},

		&GoogleComputeTargetTcpProxy{},
		&GoogleComputeTargetTcpProxyList{},

		&GoogleSpannerInstanceIamPolicy{},
		&GoogleSpannerInstanceIamPolicyList{},

		&GoogleComputeProjectMetadataItem{},
		&GoogleComputeProjectMetadataItemList{},

		&GoogleFolder{},
		&GoogleFolderList{},

		&GoogleFolderIamBinding{},
		&GoogleFolderIamBindingList{},

		&GoogleStorageBucketAcl{},
		&GoogleStorageBucketAclList{},

		&GoogleComputeRegionAutoscaler{},
		&GoogleComputeRegionAutoscalerList{},

		&GoogleKmsCryptoKeyIamBinding{},
		&GoogleKmsCryptoKeyIamBindingList{},

		&GoogleLoggingProjectExclusion{},
		&GoogleLoggingProjectExclusionList{},

		&GoogleServiceAccount{},
		&GoogleServiceAccountList{},

		&GoogleComputeAddress{},
		&GoogleComputeAddressList{},

		&GoogleStorageObjectAccessControl{},
		&GoogleStorageObjectAccessControlList{},

		&GoogleComputeSecurityPolicy{},
		&GoogleComputeSecurityPolicyList{},

		&GooglePubsubSubscriptionIamMember{},
		&GooglePubsubSubscriptionIamMemberList{},

		&GoogleComputeUrlMap{},
		&GoogleComputeUrlMapList{},

		&GoogleContainerNodePool{},
		&GoogleContainerNodePoolList{},

		&GoogleCloudfunctionsFunction{},
		&GoogleCloudfunctionsFunctionList{},

		&GoogleBinaryAuthorizationPolicy{},
		&GoogleBinaryAuthorizationPolicyList{},

		&GoogleDataprocJob{},
		&GoogleDataprocJobList{},

		&GoogleProjectIamBinding{},
		&GoogleProjectIamBindingList{},

		&GoogleComputeRegionBackendService{},
		&GoogleComputeRegionBackendServiceList{},

		&GoogleOrganizationIamMember{},
		&GoogleOrganizationIamMemberList{},

		&GoogleComputeGlobalAddress{},
		&GoogleComputeGlobalAddressList{},

		&GoogleComputeSubnetwork{},
		&GoogleComputeSubnetworkList{},

		&GoogleComputeRoute{},
		&GoogleComputeRouteList{},

		&GoogleFilestoreInstance{},
		&GoogleFilestoreInstanceList{},

		&GoogleLoggingProjectSink{},
		&GoogleLoggingProjectSinkList{},

		&GoogleProjectUsageExportBucket{},
		&GoogleProjectUsageExportBucketList{},

		&GoogleComputeInstanceTemplate{},
		&GoogleComputeInstanceTemplateList{},

		&GoogleProjectIamPolicy{},
		&GoogleProjectIamPolicyList{},

		&GoogleComputeSslCertificate{},
		&GoogleComputeSslCertificateList{},

		&GoogleBillingAccountIamBinding{},
		&GoogleBillingAccountIamBindingList{},

		&GoogleSpannerInstanceIamBinding{},
		&GoogleSpannerInstanceIamBindingList{},

		&GoogleComputeRegionInstanceGroupManager{},
		&GoogleComputeRegionInstanceGroupManagerList{},

		&GoogleComputeDisk{},
		&GoogleComputeDiskList{},

		&GoogleComputeFirewall{},
		&GoogleComputeFirewallList{},

		&GoogleComputeNetworkPeering{},
		&GoogleComputeNetworkPeeringList{},

		&GoogleOrganizationIamBinding{},
		&GoogleOrganizationIamBindingList{},

		&GoogleBillingAccountIamPolicy{},
		&GoogleBillingAccountIamPolicyList{},

		&GoogleComputeTargetSslProxy{},
		&GoogleComputeTargetSslProxyList{},

		&GoogleSpannerInstance{},
		&GoogleSpannerInstanceList{},

		&GoogleOrganizationIamPolicy{},
		&GoogleOrganizationIamPolicyList{},

		&GoogleComputeSubnetworkIamBinding{},
		&GoogleComputeSubnetworkIamBindingList{},

		&GoogleSqlDatabase{},
		&GoogleSqlDatabaseList{},

		&GoogleComputeSnapshot{},
		&GoogleComputeSnapshotList{},

		&GoogleMonitoringNotificationChannel{},
		&GoogleMonitoringNotificationChannelList{},

		&GoogleDataprocCluster{},
		&GoogleDataprocClusterList{},

		&GoogleLoggingFolderExclusion{},
		&GoogleLoggingFolderExclusionList{},

		&GoogleStorageBucketIamMember{},
		&GoogleStorageBucketIamMemberList{},

		&GoogleRedisInstance{},
		&GoogleRedisInstanceList{},

		&GoogleFolderIamMember{},
		&GoogleFolderIamMemberList{},

		&GoogleAppEngineApplication{},
		&GoogleAppEngineApplicationList{},

		&GoogleLoggingOrganizationExclusion{},
		&GoogleLoggingOrganizationExclusionList{},

		&GoogleComputeBackendService{},
		&GoogleComputeBackendServiceList{},

		&GooglePubsubSubscriptionIamBinding{},
		&GooglePubsubSubscriptionIamBindingList{},

		&GoogleOrganizationIamCustomRole{},
		&GoogleOrganizationIamCustomRoleList{},

		&GoogleComputeInstance{},
		&GoogleComputeInstanceList{},

		&GoogleSpannerDatabaseIamMember{},
		&GoogleSpannerDatabaseIamMemberList{},

		&GoogleComputeRegionDisk{},
		&GoogleComputeRegionDiskList{},

		&GoogleKmsKeyRingIamPolicy{},
		&GoogleKmsKeyRingIamPolicyList{},

		&GoogleComputeRouterPeer{},
		&GoogleComputeRouterPeerList{},

		&GoogleFolderOrganizationPolicy{},
		&GoogleFolderOrganizationPolicyList{},

		&GoogleComposerEnvironment{},
		&GoogleComposerEnvironmentList{},

		&GoogleStorageObjectAcl{},
		&GoogleStorageObjectAclList{},

		&GoogleSqlSslCert{},
		&GoogleSqlSslCertList{},

		&GoogleComputeInstanceGroupManager{},
		&GoogleComputeInstanceGroupManagerList{},

		&GoogleBigtableInstance{},
		&GoogleBigtableInstanceList{},

		&GoogleStorageNotification{},
		&GoogleStorageNotificationList{},

		&GoogleKmsKeyRingIamBinding{},
		&GoogleKmsKeyRingIamBindingList{},

		&GoogleRuntimeconfigVariable{},
		&GoogleRuntimeconfigVariableList{},

		&GoogleComputeTargetHttpsProxy{},
		&GoogleComputeTargetHttpsProxyList{},

		&GoogleKmsCryptoKey{},
		&GoogleKmsCryptoKeyList{},

		&GoogleFolderIamPolicy{},
		&GoogleFolderIamPolicyList{},

		&GoogleComputeRouter{},
		&GoogleComputeRouterList{},

		&GoogleComputeSharedVpcServiceProject{},
		&GoogleComputeSharedVpcServiceProjectList{},

		&GoogleSourcerepoRepository{},
		&GoogleSourcerepoRepositoryList{},

		&GoogleContainerAnalysisNote{},
		&GoogleContainerAnalysisNoteList{},

		&GoogleProjectIamMember{},
		&GoogleProjectIamMemberList{},

		&GoogleComputeSharedVpcHostProject{},
		&GoogleComputeSharedVpcHostProjectList{},

		&GoogleProjectServices{},
		&GoogleProjectServicesList{},

		&GooglePubsubTopic{},
		&GooglePubsubTopicList{},

		&GoogleProjectService{},
		&GoogleProjectServiceList{},

		&GoogleComputeAttachedDisk{},
		&GoogleComputeAttachedDiskList{},

		&GoogleComputeRouterNat{},
		&GoogleComputeRouterNatList{},

		&GoogleSqlUser{},
		&GoogleSqlUserList{},

		&GoogleProject{},
		&GoogleProjectList{},

		&GoogleBigqueryTable{},
		&GoogleBigqueryTableList{},

		&GooglePubsubSubscription{},
		&GooglePubsubSubscriptionList{},

		&GoogleMonitoringUptimeCheckConfig{},
		&GoogleMonitoringUptimeCheckConfigList{},

		&GoogleStorageDefaultObjectAcl{},
		&GoogleStorageDefaultObjectAclList{},

		&GoogleComputeInstanceGroup{},
		&GoogleComputeInstanceGroupList{},

		&GoogleSpannerDatabaseIamBinding{},
		&GoogleSpannerDatabaseIamBindingList{},

		&GoogleLoggingBillingAccountSink{},
		&GoogleLoggingBillingAccountSinkList{},

		&GoogleComputeInstanceFromTemplate{},
		&GoogleComputeInstanceFromTemplateList{},

		&GoogleMonitoringGroup{},
		&GoogleMonitoringGroupList{},

		&GoogleComputeAutoscaler{},
		&GoogleComputeAutoscalerList{},

		&GoogleComputeInterconnectAttachment{},
		&GoogleComputeInterconnectAttachmentList{},

		&GoogleComputeTargetPool{},
		&GoogleComputeTargetPoolList{},

		&GoogleBinaryAuthorizationAttestor{},
		&GoogleBinaryAuthorizationAttestorList{},

		&GoogleKmsKeyRing{},
		&GoogleKmsKeyRingList{},

		&GoogleResourceManagerLien{},
		&GoogleResourceManagerLienList{},

		&GoogleStorageDefaultObjectAccessControl{},
		&GoogleStorageDefaultObjectAccessControlList{},

		&GoogleComputeProjectMetadata{},
		&GoogleComputeProjectMetadataList{},

		&GoogleServiceAccountIamBinding{},
		&GoogleServiceAccountIamBindingList{},

		&GooglePubsubSubscriptionIamPolicy{},
		&GooglePubsubSubscriptionIamPolicyList{},

		&GoogleComputeHttpsHealthCheck{},
		&GoogleComputeHttpsHealthCheckList{},

		&GoogleComputeSslPolicy{},
		&GoogleComputeSslPolicyList{},

		&GooglePubsubTopicIamMember{},
		&GooglePubsubTopicIamMemberList{},

		&GoogleStorageBucketObject{},
		&GoogleStorageBucketObjectList{},

		&GoogleStorageBucketIamBinding{},
		&GoogleStorageBucketIamBindingList{},

		&GoogleComputeForwardingRule{},
		&GoogleComputeForwardingRuleList{},

		&GoogleComputeRouterInterface{},
		&GoogleComputeRouterInterfaceList{},

		&GoogleProjectOrganizationPolicy{},
		&GoogleProjectOrganizationPolicyList{},

		&GoogleKmsKeyRingIamMember{},
		&GoogleKmsKeyRingIamMemberList{},

		&GooglePubsubTopicIamBinding{},
		&GooglePubsubTopicIamBindingList{},

		&GoogleRuntimeconfigConfig{},
		&GoogleRuntimeconfigConfigList{},

		&GoogleKmsCryptoKeyIamMember{},
		&GoogleKmsCryptoKeyIamMemberList{},

		&GoogleServiceAccountIamMember{},
		&GoogleServiceAccountIamMemberList{},

		&GoogleComputeGlobalForwardingRule{},
		&GoogleComputeGlobalForwardingRuleList{},

		&GoogleComputeBackendBucket{},
		&GoogleComputeBackendBucketList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
