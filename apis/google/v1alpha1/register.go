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

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&Project{},
		&ProjectList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ProjectService{},
		&ProjectServiceList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&Folder{},
		&FolderList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&BigtableTable{},
		&BigtableTableList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&DataflowJob{},
		&DataflowJobList{},

		&DataprocJob{},
		&DataprocJobList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeImage{},
		&ComputeImageList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
