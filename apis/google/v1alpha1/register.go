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

		&GoogleFolder{},
		&GoogleFolderList{},

		&GoogleSqlUser{},
		&GoogleSqlUserList{},

		&GoogleLoggingProjectSink{},
		&GoogleLoggingProjectSinkList{},

		&GoogleComputeRouterInterface{},
		&GoogleComputeRouterInterfaceList{},

		&GoogleComputeInstance{},
		&GoogleComputeInstanceList{},

		&GoogleComputeDisk{},
		&GoogleComputeDiskList{},

		&GoogleComputeTargetHttpsProxy{},
		&GoogleComputeTargetHttpsProxyList{},

		&GoogleLoggingBillingAccountSink{},
		&GoogleLoggingBillingAccountSinkList{},

		&GoogleDataprocCluster{},
		&GoogleDataprocClusterList{},

		&GoogleComputeTargetTcpProxy{},
		&GoogleComputeTargetTcpProxyList{},

		&GoogleComputeRegionAutoscaler{},
		&GoogleComputeRegionAutoscalerList{},

		&GoogleComputeBackendService{},
		&GoogleComputeBackendServiceList{},

		&GoogleMonitoringNotificationChannel{},
		&GoogleMonitoringNotificationChannelList{},

		&GoogleOrganizationIamMember{},
		&GoogleOrganizationIamMemberList{},

		&GoogleStorageBucketIamMember{},
		&GoogleStorageBucketIamMemberList{},

		&GoogleOrganizationIamPolicy{},
		&GoogleOrganizationIamPolicyList{},

		&GoogleKmsCryptoKeyIamMember{},
		&GoogleKmsCryptoKeyIamMemberList{},

		&GoogleSpannerInstanceIamBinding{},
		&GoogleSpannerInstanceIamBindingList{},

		&GoogleComputeSubnetworkIamPolicy{},
		&GoogleComputeSubnetworkIamPolicyList{},

		&GoogleKmsCryptoKey{},
		&GoogleKmsCryptoKeyList{},

		&GoogleComputeSubnetworkIamMember{},
		&GoogleComputeSubnetworkIamMemberList{},

		&GoogleContainerNodePool{},
		&GoogleContainerNodePoolList{},

		&GoogleSourcerepoRepository{},
		&GoogleSourcerepoRepositoryList{},

		&GoogleResourceManagerLien{},
		&GoogleResourceManagerLienList{},

		&GoogleSpannerDatabaseIamBinding{},
		&GoogleSpannerDatabaseIamBindingList{},

		&GoogleComputeAttachedDisk{},
		&GoogleComputeAttachedDiskList{},

		&GoogleStorageBucketIamPolicy{},
		&GoogleStorageBucketIamPolicyList{},

		&GoogleBillingAccountIamBinding{},
		&GoogleBillingAccountIamBindingList{},

		&GooglePubsubSubscription{},
		&GooglePubsubSubscriptionList{},

		&GoogleStorageBucketIamBinding{},
		&GoogleStorageBucketIamBindingList{},

		&GoogleRuntimeconfigVariable{},
		&GoogleRuntimeconfigVariableList{},

		&GoogleCloudbuildTrigger{},
		&GoogleCloudbuildTriggerList{},

		&GoogleComputeAddress{},
		&GoogleComputeAddressList{},

		&GoogleRedisInstance{},
		&GoogleRedisInstanceList{},

		&GoogleMonitoringUptimeCheckConfig{},
		&GoogleMonitoringUptimeCheckConfigList{},

		&GoogleOrganizationPolicy{},
		&GoogleOrganizationPolicyList{},

		&GoogleCloudfunctionsFunction{},
		&GoogleCloudfunctionsFunctionList{},

		&GoogleServiceAccountIamPolicy{},
		&GoogleServiceAccountIamPolicyList{},

		&GoogleComputeHealthCheck{},
		&GoogleComputeHealthCheckList{},

		&GoogleComputeTargetSslProxy{},
		&GoogleComputeTargetSslProxyList{},

		&GoogleComputeSharedVpcServiceProject{},
		&GoogleComputeSharedVpcServiceProjectList{},

		&GoogleProjectOrganizationPolicy{},
		&GoogleProjectOrganizationPolicyList{},

		&GoogleComputeProjectMetadataItem{},
		&GoogleComputeProjectMetadataItemList{},

		&GoogleBinaryAuthorizationPolicy{},
		&GoogleBinaryAuthorizationPolicyList{},

		&GoogleComputeTargetHttpProxy{},
		&GoogleComputeTargetHttpProxyList{},

		&GoogleStorageBucket{},
		&GoogleStorageBucketList{},

		&GoogleProjectUsageExportBucket{},
		&GoogleProjectUsageExportBucketList{},

		&GoogleSqlDatabase{},
		&GoogleSqlDatabaseList{},

		&GoogleComputeRouter{},
		&GoogleComputeRouterList{},

		&GoogleComputeSecurityPolicy{},
		&GoogleComputeSecurityPolicyList{},

		&GoogleComputeGlobalForwardingRule{},
		&GoogleComputeGlobalForwardingRuleList{},

		&GoogleBigqueryDataset{},
		&GoogleBigqueryDatasetList{},

		&GoogleComputeInstanceFromTemplate{},
		&GoogleComputeInstanceFromTemplateList{},

		&GoogleFolderIamMember{},
		&GoogleFolderIamMemberList{},

		&GoogleComputeAutoscaler{},
		&GoogleComputeAutoscalerList{},

		&GoogleFilestoreInstance{},
		&GoogleFilestoreInstanceList{},

		&GoogleServiceAccountKey{},
		&GoogleServiceAccountKeyList{},

		&GoogleComputeForwardingRule{},
		&GoogleComputeForwardingRuleList{},

		&GoogleFolderIamBinding{},
		&GoogleFolderIamBindingList{},

		&GoogleLoggingFolderExclusion{},
		&GoogleLoggingFolderExclusionList{},

		&GoogleComputeNetwork{},
		&GoogleComputeNetworkList{},

		&GoogleBillingAccountIamPolicy{},
		&GoogleBillingAccountIamPolicyList{},

		&GoogleProjectIamCustomRole{},
		&GoogleProjectIamCustomRoleList{},

		&GoogleSqlDatabaseInstance{},
		&GoogleSqlDatabaseInstanceList{},

		&GoogleKmsKeyRingIamPolicy{},
		&GoogleKmsKeyRingIamPolicyList{},

		&GoogleComputeGlobalAddress{},
		&GoogleComputeGlobalAddressList{},

		&GoogleDnsManagedZone{},
		&GoogleDnsManagedZoneList{},

		&GoogleComputeRouterNat{},
		&GoogleComputeRouterNatList{},

		&GoogleComputeRoute{},
		&GoogleComputeRouteList{},

		&GoogleComputeSubnetwork{},
		&GoogleComputeSubnetworkList{},

		&GoogleComputeHttpHealthCheck{},
		&GoogleComputeHttpHealthCheckList{},

		&GoogleComputeInstanceGroupManager{},
		&GoogleComputeInstanceGroupManagerList{},

		&GoogleFolderIamPolicy{},
		&GoogleFolderIamPolicyList{},

		&GoogleDnsRecordSet{},
		&GoogleDnsRecordSetList{},

		&GoogleOrganizationIamBinding{},
		&GoogleOrganizationIamBindingList{},

		&GoogleStorageObjectAccessControl{},
		&GoogleStorageObjectAccessControlList{},

		&GoogleKmsCryptoKeyIamBinding{},
		&GoogleKmsCryptoKeyIamBindingList{},

		&GoogleAppEngineApplication{},
		&GoogleAppEngineApplicationList{},

		&GoogleServiceAccount{},
		&GoogleServiceAccountList{},

		&GoogleBigtableTable{},
		&GoogleBigtableTableList{},

		&GoogleComputeRegionDisk{},
		&GoogleComputeRegionDiskList{},

		&GoogleComputeRegionInstanceGroupManager{},
		&GoogleComputeRegionInstanceGroupManagerList{},

		&GoogleSpannerDatabaseIamMember{},
		&GoogleSpannerDatabaseIamMemberList{},

		&GoogleMonitoringGroup{},
		&GoogleMonitoringGroupList{},

		&GoogleCloudiotRegistry{},
		&GoogleCloudiotRegistryList{},

		&GoogleProjectIamBinding{},
		&GoogleProjectIamBindingList{},

		&GoogleStorageDefaultObjectAcl{},
		&GoogleStorageDefaultObjectAclList{},

		&GoogleComputeInterconnectAttachment{},
		&GoogleComputeInterconnectAttachmentList{},

		&GoogleComputeVpnTunnel{},
		&GoogleComputeVpnTunnelList{},

		&GoogleContainerAnalysisNote{},
		&GoogleContainerAnalysisNoteList{},

		&GooglePubsubTopicIamPolicy{},
		&GooglePubsubTopicIamPolicyList{},

		&GoogleBigtableInstance{},
		&GoogleBigtableInstanceList{},

		&GoogleLoggingProjectExclusion{},
		&GoogleLoggingProjectExclusionList{},

		&GoogleBinaryAuthorizationAttestor{},
		&GoogleBinaryAuthorizationAttestorList{},

		&GoogleLoggingBillingAccountExclusion{},
		&GoogleLoggingBillingAccountExclusionList{},

		&GoogleStorageObjectAcl{},
		&GoogleStorageObjectAclList{},

		&GoogleComputeSubnetworkIamBinding{},
		&GoogleComputeSubnetworkIamBindingList{},

		&GooglePubsubTopic{},
		&GooglePubsubTopicList{},

		&GoogleContainerCluster{},
		&GoogleContainerClusterList{},

		&GoogleComposerEnvironment{},
		&GoogleComposerEnvironmentList{},

		&GoogleServiceAccountIamBinding{},
		&GoogleServiceAccountIamBindingList{},

		&GoogleComputeSnapshot{},
		&GoogleComputeSnapshotList{},

		&GoogleStorageDefaultObjectAccessControl{},
		&GoogleStorageDefaultObjectAccessControlList{},

		&GoogleProjectIamPolicy{},
		&GoogleProjectIamPolicyList{},

		&GoogleSpannerInstanceIamPolicy{},
		&GoogleSpannerInstanceIamPolicyList{},

		&GoogleProjectServices{},
		&GoogleProjectServicesList{},

		&GoogleSpannerInstanceIamMember{},
		&GoogleSpannerInstanceIamMemberList{},

		&GoogleSpannerDatabaseIamPolicy{},
		&GoogleSpannerDatabaseIamPolicyList{},

		&GoogleProjectIamMember{},
		&GoogleProjectIamMemberList{},

		&GoogleMonitoringAlertPolicy{},
		&GoogleMonitoringAlertPolicyList{},

		&GoogleStorageNotification{},
		&GoogleStorageNotificationList{},

		&GoogleComputeInstanceTemplate{},
		&GoogleComputeInstanceTemplateList{},

		&GoogleComputeNetworkPeering{},
		&GoogleComputeNetworkPeeringList{},

		&GoogleDataprocJob{},
		&GoogleDataprocJobList{},

		&GoogleKmsKeyRingIamBinding{},
		&GoogleKmsKeyRingIamBindingList{},

		&GoogleLoggingFolderSink{},
		&GoogleLoggingFolderSinkList{},

		&GooglePubsubSubscriptionIamMember{},
		&GooglePubsubSubscriptionIamMemberList{},

		&GoogleSqlSslCert{},
		&GoogleSqlSslCertList{},

		&GoogleBillingAccountIamMember{},
		&GoogleBillingAccountIamMemberList{},

		&GoogleProject{},
		&GoogleProjectList{},

		&GoogleComputeVpnGateway{},
		&GoogleComputeVpnGatewayList{},

		&GoogleOrganizationIamCustomRole{},
		&GoogleOrganizationIamCustomRoleList{},

		&GoogleComputeImage{},
		&GoogleComputeImageList{},

		&GoogleComputeRouterPeer{},
		&GoogleComputeRouterPeerList{},

		&GoogleComputeHttpsHealthCheck{},
		&GoogleComputeHttpsHealthCheckList{},

		&GoogleComputeTargetPool{},
		&GoogleComputeTargetPoolList{},

		&GooglePubsubTopicIamMember{},
		&GooglePubsubTopicIamMemberList{},

		&GoogleFolderOrganizationPolicy{},
		&GoogleFolderOrganizationPolicyList{},

		&GooglePubsubSubscriptionIamBinding{},
		&GooglePubsubSubscriptionIamBindingList{},

		&GoogleKmsKeyRing{},
		&GoogleKmsKeyRingList{},

		&GoogleComputeSslCertificate{},
		&GoogleComputeSslCertificateList{},

		&GoogleStorageBucketAcl{},
		&GoogleStorageBucketAclList{},

		&GoogleComputeRegionBackendService{},
		&GoogleComputeRegionBackendServiceList{},

		&GoogleStorageBucketObject{},
		&GoogleStorageBucketObjectList{},

		&GoogleEndpointsService{},
		&GoogleEndpointsServiceList{},

		&GooglePubsubSubscriptionIamPolicy{},
		&GooglePubsubSubscriptionIamPolicyList{},

		&GoogleComputeSharedVpcHostProject{},
		&GoogleComputeSharedVpcHostProjectList{},

		&GoogleBigqueryTable{},
		&GoogleBigqueryTableList{},

		&GoogleRuntimeconfigConfig{},
		&GoogleRuntimeconfigConfigList{},

		&GoogleComputeInstanceGroup{},
		&GoogleComputeInstanceGroupList{},

		&GoogleComputeFirewall{},
		&GoogleComputeFirewallList{},

		&GoogleSpannerInstance{},
		&GoogleSpannerInstanceList{},

		&GoogleServiceAccountIamMember{},
		&GoogleServiceAccountIamMemberList{},

		&GoogleComputeBackendBucket{},
		&GoogleComputeBackendBucketList{},

		&GoogleLoggingOrganizationSink{},
		&GoogleLoggingOrganizationSinkList{},

		&GoogleLoggingOrganizationExclusion{},
		&GoogleLoggingOrganizationExclusionList{},

		&GoogleDataflowJob{},
		&GoogleDataflowJobList{},

		&GoogleComputeSslPolicy{},
		&GoogleComputeSslPolicyList{},

		&GoogleProjectService{},
		&GoogleProjectServiceList{},

		&GoogleKmsKeyRingIamMember{},
		&GoogleKmsKeyRingIamMemberList{},

		&GoogleSpannerDatabase{},
		&GoogleSpannerDatabaseList{},

		&GooglePubsubTopicIamBinding{},
		&GooglePubsubTopicIamBindingList{},

		&GoogleComputeUrlMap{},
		&GoogleComputeUrlMapList{},

		&GoogleComputeProjectMetadata{},
		&GoogleComputeProjectMetadataList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
