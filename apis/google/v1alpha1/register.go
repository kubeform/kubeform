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

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ProjectService{},
		&ProjectServiceList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&Folder{},
		&FolderList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeImage{},
		&ComputeImageList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&Project{},
		&ProjectList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&DataflowJob{},
		&DataflowJobList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&BigtableTable{},
		&BigtableTableList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&SqlUser{},
		&SqlUserList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
