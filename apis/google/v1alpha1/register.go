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

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&DataprocJob{},
		&DataprocJobList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&StorageBucket{},
		&StorageBucketList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&Folder{},
		&FolderList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&Project{},
		&ProjectList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&DataflowJob{},
		&DataflowJobList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&SqlUser{},
		&SqlUserList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
