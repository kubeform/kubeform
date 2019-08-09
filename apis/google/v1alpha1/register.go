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

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&Project{},
		&ProjectList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&Folder{},
		&FolderList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
