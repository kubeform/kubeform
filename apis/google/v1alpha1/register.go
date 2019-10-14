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

		&StorageNotification{},
		&StorageNotificationList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&Folder{},
		&FolderList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&DataprocJob{},
		&DataprocJobList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ProjectService{},
		&ProjectServiceList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&BigtableTable{},
		&BigtableTableList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&Project{},
		&ProjectList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&StorageBucket{},
		&StorageBucketList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
