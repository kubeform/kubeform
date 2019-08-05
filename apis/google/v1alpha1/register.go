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

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&SqlUser{},
		&SqlUserList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&DataflowJob{},
		&DataflowJobList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&Folder{},
		&FolderList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&Project{},
		&ProjectList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&DataprocJob{},
		&DataprocJobList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
