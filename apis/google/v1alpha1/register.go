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

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeImage{},
		&ComputeImageList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&DataflowJob{},
		&DataflowJobList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&Folder{},
		&FolderList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&Project{},
		&ProjectList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&BigtableTable{},
		&BigtableTableList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
