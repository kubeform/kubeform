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

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&DataflowJob{},
		&DataflowJobList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&Project{},
		&ProjectList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&DataprocJob{},
		&DataprocJobList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeImage{},
		&ComputeImageList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&SqlUser{},
		&SqlUserList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&StorageBucket{},
		&StorageBucketList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&Folder{},
		&FolderList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ProjectService{},
		&ProjectServiceList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
