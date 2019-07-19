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

		&ProjectService{},
		&ProjectServiceList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&Project{},
		&ProjectList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&DataflowJob{},
		&DataflowJobList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&BigtableTable{},
		&BigtableTableList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&SqlUser{},
		&SqlUserList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&Folder{},
		&FolderList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
