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

		&ComputeInstance{},
		&ComputeInstanceList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&Project{},
		&ProjectList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&Folder{},
		&FolderList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&SqlUser{},
		&SqlUserList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&DataflowJob{},
		&DataflowJobList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
