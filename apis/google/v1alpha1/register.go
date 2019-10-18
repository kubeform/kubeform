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

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeImage{},
		&ComputeImageList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&DataflowJob{},
		&DataflowJobList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&Folder{},
		&FolderList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&Project{},
		&ProjectList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&StorageBucket{},
		&StorageBucketList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&DataprocJob{},
		&DataprocJobList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
