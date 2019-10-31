/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	"kubeform.dev/kubeform/apis/google"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&BigtableTable{},
		&BigtableTableList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&DataflowJob{},
		&DataflowJobList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&DataprocJob{},
		&DataprocJobList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&Folder{},
		&FolderList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&Project{},
		&ProjectList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ProjectServiceBatch{},
		&ProjectServiceBatchList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&SqlUser{},
		&SqlUserList{},

		&StorageBucket{},
		&StorageBucketList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
