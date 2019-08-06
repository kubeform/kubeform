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

		&ComputeAddress{},
		&ComputeAddressList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&SqlUser{},
		&SqlUserList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&Project{},
		&ProjectList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&DataprocJob{},
		&DataprocJobList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&DataflowJob{},
		&DataflowJobList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&Folder{},
		&FolderList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&StorageBucket{},
		&StorageBucketList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
