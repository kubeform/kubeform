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

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&Folder{},
		&FolderList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&BigtableTable{},
		&BigtableTableList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&Project{},
		&ProjectList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&StorageBucket{},
		&StorageBucketList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
