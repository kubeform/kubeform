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

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&BigtableTable{},
		&BigtableTableList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ProjectService{},
		&ProjectServiceList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&Folder{},
		&FolderList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&Project{},
		&ProjectList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
