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

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&Project{},
		&ProjectList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&BigtableTable{},
		&BigtableTableList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ProjectService{},
		&ProjectServiceList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlUser{},
		&SqlUserList{},

		&DataflowJob{},
		&DataflowJobList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&Folder{},
		&FolderList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&StorageBucket{},
		&StorageBucketList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
