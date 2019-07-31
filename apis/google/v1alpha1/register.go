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

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&SqlUser{},
		&SqlUserList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&Folder{},
		&FolderList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeImage{},
		&ComputeImageList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&Project{},
		&ProjectList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
