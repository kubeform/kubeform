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

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DataprocJob{},
		&DataprocJobList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&DataflowJob{},
		&DataflowJobList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&Folder{},
		&FolderList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&BigtableTable{},
		&BigtableTableList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&SqlUser{},
		&SqlUserList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&StorageBucket{},
		&StorageBucketList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&Project{},
		&ProjectList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&RedisInstance{},
		&RedisInstanceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
