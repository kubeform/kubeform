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

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&Folder{},
		&FolderList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeVpnGateway{},
		&ComputeVpnGatewayList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&Project{},
		&ProjectList{},

		&ComputeUrlMap{},
		&ComputeUrlMapList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&StorageObjectAcl{},
		&StorageObjectAclList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeHttpsHealthCheck{},
		&ComputeHttpsHealthCheckList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&StorageBucket{},
		&StorageBucketList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ComputeTargetHttpsProxy{},
		&ComputeTargetHttpsProxyList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeHttpHealthCheck{},
		&ComputeHttpHealthCheckList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeRouterNat{},
		&ComputeRouterNatList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeVpnTunnel{},
		&ComputeVpnTunnelList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&SqlUser{},
		&SqlUserList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeTargetHttpProxy{},
		&ComputeTargetHttpProxyList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageDefaultObjectAcl{},
		&StorageDefaultObjectAclList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&DataflowJob{},
		&DataflowJobList{},

		&StorageBucketAcl{},
		&StorageBucketAclList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
