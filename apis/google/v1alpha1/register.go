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

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&Folder{},
		&FolderList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&SqlUser{},
		&SqlUserList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ProjectService{},
		&ProjectServiceList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeImage{},
		&ComputeImageList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&BigtableTable{},
		&BigtableTableList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&Project{},
		&ProjectList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
