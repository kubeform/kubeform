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

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&DataflowJob{},
		&DataflowJobList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&BigtableTable{},
		&BigtableTableList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&StorageBucket{},
		&StorageBucketList{},

		&Folder{},
		&FolderList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&SqlUser{},
		&SqlUserList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&Project{},
		&ProjectList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ComputeImage{},
		&ComputeImageList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
