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

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&DataflowJob{},
		&DataflowJobList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&Project{},
		&ProjectList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&Folder{},
		&FolderList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&StorageBucket{},
		&StorageBucketList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
