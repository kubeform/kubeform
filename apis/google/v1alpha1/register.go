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

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&Folder{},
		&FolderList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ProjectService{},
		&ProjectServiceList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&StorageBucket{},
		&StorageBucketList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&Project{},
		&ProjectList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&BigtableTable{},
		&BigtableTableList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&DataflowJob{},
		&DataflowJobList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
