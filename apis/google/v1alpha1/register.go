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

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&BigtableTable{},
		&BigtableTableList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&Project{},
		&ProjectList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeImage{},
		&ComputeImageList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&SqlUser{},
		&SqlUserList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&Folder{},
		&FolderList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ProjectService{},
		&ProjectServiceList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
