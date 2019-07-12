/*
Copyright 2019 The KubeDB Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	"kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

type GoogleV1alpha1Interface interface {
	RESTClient() rest.Interface
	GoogleAppEngineApplicationsGetter
	GoogleBigqueryDatasetsGetter
	GoogleBigqueryTablesGetter
	GoogleBigtableInstancesGetter
	GoogleBigtableTablesGetter
	GoogleBillingAccountIamBindingsGetter
	GoogleBillingAccountIamMembersGetter
	GoogleBillingAccountIamPoliciesGetter
	GoogleBinaryAuthorizationAttestorsGetter
	GoogleBinaryAuthorizationPoliciesGetter
	GoogleCloudbuildTriggersGetter
	GoogleCloudfunctionsFunctionsGetter
	GoogleCloudiotRegistriesGetter
	GoogleComposerEnvironmentsGetter
	GoogleComputeAddressesGetter
	GoogleComputeAttachedDisksGetter
	GoogleComputeAutoscalersGetter
	GoogleComputeBackendBucketsGetter
	GoogleComputeBackendServicesGetter
	GoogleComputeDisksGetter
	GoogleComputeFirewallsGetter
	GoogleComputeForwardingRulesGetter
	GoogleComputeGlobalAddressesGetter
	GoogleComputeGlobalForwardingRulesGetter
	GoogleComputeHealthChecksGetter
	GoogleComputeHttpHealthChecksGetter
	GoogleComputeHttpsHealthChecksGetter
	GoogleComputeImagesGetter
	GoogleComputeInstancesGetter
	GoogleComputeInstanceFromTemplatesGetter
	GoogleComputeInstanceGroupsGetter
	GoogleComputeInstanceGroupManagersGetter
	GoogleComputeInstanceTemplatesGetter
	GoogleComputeInterconnectAttachmentsGetter
	GoogleComputeNetworksGetter
	GoogleComputeNetworkPeeringsGetter
	GoogleComputeProjectMetadatasGetter
	GoogleComputeProjectMetadataItemsGetter
	GoogleComputeRegionAutoscalersGetter
	GoogleComputeRegionBackendServicesGetter
	GoogleComputeRegionDisksGetter
	GoogleComputeRegionInstanceGroupManagersGetter
	GoogleComputeRoutesGetter
	GoogleComputeRoutersGetter
	GoogleComputeRouterInterfacesGetter
	GoogleComputeRouterNatsGetter
	GoogleComputeRouterPeersGetter
	GoogleComputeSecurityPoliciesGetter
	GoogleComputeSharedVpcHostProjectsGetter
	GoogleComputeSharedVpcServiceProjectsGetter
	GoogleComputeSnapshotsGetter
	GoogleComputeSslCertificatesGetter
	GoogleComputeSslPoliciesGetter
	GoogleComputeSubnetworksGetter
	GoogleComputeSubnetworkIamBindingsGetter
	GoogleComputeSubnetworkIamMembersGetter
	GoogleComputeSubnetworkIamPoliciesGetter
	GoogleComputeTargetHttpProxiesGetter
	GoogleComputeTargetHttpsProxiesGetter
	GoogleComputeTargetPoolsGetter
	GoogleComputeTargetSslProxiesGetter
	GoogleComputeTargetTcpProxiesGetter
	GoogleComputeUrlMapsGetter
	GoogleComputeVpnGatewaysGetter
	GoogleComputeVpnTunnelsGetter
	GoogleContainerAnalysisNotesGetter
	GoogleContainerClustersGetter
	GoogleContainerNodePoolsGetter
	GoogleDataflowJobsGetter
	GoogleDataprocClustersGetter
	GoogleDataprocJobsGetter
	GoogleDnsManagedZonesGetter
	GoogleDnsRecordSetsGetter
	GoogleEndpointsServicesGetter
	GoogleFilestoreInstancesGetter
	GoogleFoldersGetter
	GoogleFolderIamBindingsGetter
	GoogleFolderIamMembersGetter
	GoogleFolderIamPoliciesGetter
	GoogleFolderOrganizationPoliciesGetter
	GoogleKmsCryptoKeysGetter
	GoogleKmsCryptoKeyIamBindingsGetter
	GoogleKmsCryptoKeyIamMembersGetter
	GoogleKmsKeyRingsGetter
	GoogleKmsKeyRingIamBindingsGetter
	GoogleKmsKeyRingIamMembersGetter
	GoogleKmsKeyRingIamPoliciesGetter
	GoogleLoggingBillingAccountExclusionsGetter
	GoogleLoggingBillingAccountSinksGetter
	GoogleLoggingFolderExclusionsGetter
	GoogleLoggingFolderSinksGetter
	GoogleLoggingOrganizationExclusionsGetter
	GoogleLoggingOrganizationSinksGetter
	GoogleLoggingProjectExclusionsGetter
	GoogleLoggingProjectSinksGetter
	GoogleMonitoringAlertPoliciesGetter
	GoogleMonitoringGroupsGetter
	GoogleMonitoringNotificationChannelsGetter
	GoogleMonitoringUptimeCheckConfigsGetter
	GoogleOrganizationIamBindingsGetter
	GoogleOrganizationIamCustomRolesGetter
	GoogleOrganizationIamMembersGetter
	GoogleOrganizationIamPoliciesGetter
	GoogleOrganizationPoliciesGetter
	GoogleProjectsGetter
	GoogleProjectIamBindingsGetter
	GoogleProjectIamCustomRolesGetter
	GoogleProjectIamMembersGetter
	GoogleProjectIamPoliciesGetter
	GoogleProjectOrganizationPoliciesGetter
	GoogleProjectServicesGetter
	GoogleProjectServicesesGetter
	GoogleProjectUsageExportBucketsGetter
	GooglePubsubSubscriptionsGetter
	GooglePubsubSubscriptionIamBindingsGetter
	GooglePubsubSubscriptionIamMembersGetter
	GooglePubsubSubscriptionIamPoliciesGetter
	GooglePubsubTopicsGetter
	GooglePubsubTopicIamBindingsGetter
	GooglePubsubTopicIamMembersGetter
	GooglePubsubTopicIamPoliciesGetter
	GoogleRedisInstancesGetter
	GoogleResourceManagerLiensGetter
	GoogleRuntimeconfigConfigsGetter
	GoogleRuntimeconfigVariablesGetter
	GoogleServiceAccountsGetter
	GoogleServiceAccountIamBindingsGetter
	GoogleServiceAccountIamMembersGetter
	GoogleServiceAccountIamPoliciesGetter
	GoogleServiceAccountKeysGetter
	GoogleSourcerepoRepositoriesGetter
	GoogleSpannerDatabasesGetter
	GoogleSpannerDatabaseIamBindingsGetter
	GoogleSpannerDatabaseIamMembersGetter
	GoogleSpannerDatabaseIamPoliciesGetter
	GoogleSpannerInstancesGetter
	GoogleSpannerInstanceIamBindingsGetter
	GoogleSpannerInstanceIamMembersGetter
	GoogleSpannerInstanceIamPoliciesGetter
	GoogleSqlDatabasesGetter
	GoogleSqlDatabaseInstancesGetter
	GoogleSqlSslCertsGetter
	GoogleSqlUsersGetter
	GoogleStorageBucketsGetter
	GoogleStorageBucketAclsGetter
	GoogleStorageBucketIamBindingsGetter
	GoogleStorageBucketIamMembersGetter
	GoogleStorageBucketIamPoliciesGetter
	GoogleStorageBucketObjectsGetter
	GoogleStorageDefaultObjectAccessControlsGetter
	GoogleStorageDefaultObjectAclsGetter
	GoogleStorageNotificationsGetter
	GoogleStorageObjectAccessControlsGetter
	GoogleStorageObjectAclsGetter
}

// GoogleV1alpha1Client is used to interact with features provided by the google.kubeform.com group.
type GoogleV1alpha1Client struct {
	restClient rest.Interface
}

func (c *GoogleV1alpha1Client) GoogleAppEngineApplications() GoogleAppEngineApplicationInterface {
	return newGoogleAppEngineApplications(c)
}

func (c *GoogleV1alpha1Client) GoogleBigqueryDatasets() GoogleBigqueryDatasetInterface {
	return newGoogleBigqueryDatasets(c)
}

func (c *GoogleV1alpha1Client) GoogleBigqueryTables() GoogleBigqueryTableInterface {
	return newGoogleBigqueryTables(c)
}

func (c *GoogleV1alpha1Client) GoogleBigtableInstances() GoogleBigtableInstanceInterface {
	return newGoogleBigtableInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleBigtableTables() GoogleBigtableTableInterface {
	return newGoogleBigtableTables(c)
}

func (c *GoogleV1alpha1Client) GoogleBillingAccountIamBindings() GoogleBillingAccountIamBindingInterface {
	return newGoogleBillingAccountIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleBillingAccountIamMembers() GoogleBillingAccountIamMemberInterface {
	return newGoogleBillingAccountIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleBillingAccountIamPolicies() GoogleBillingAccountIamPolicyInterface {
	return newGoogleBillingAccountIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleBinaryAuthorizationAttestors() GoogleBinaryAuthorizationAttestorInterface {
	return newGoogleBinaryAuthorizationAttestors(c)
}

func (c *GoogleV1alpha1Client) GoogleBinaryAuthorizationPolicies() GoogleBinaryAuthorizationPolicyInterface {
	return newGoogleBinaryAuthorizationPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleCloudbuildTriggers() GoogleCloudbuildTriggerInterface {
	return newGoogleCloudbuildTriggers(c)
}

func (c *GoogleV1alpha1Client) GoogleCloudfunctionsFunctions() GoogleCloudfunctionsFunctionInterface {
	return newGoogleCloudfunctionsFunctions(c)
}

func (c *GoogleV1alpha1Client) GoogleCloudiotRegistries() GoogleCloudiotRegistryInterface {
	return newGoogleCloudiotRegistries(c)
}

func (c *GoogleV1alpha1Client) GoogleComposerEnvironments() GoogleComposerEnvironmentInterface {
	return newGoogleComposerEnvironments(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeAddresses() GoogleComputeAddressInterface {
	return newGoogleComputeAddresses(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeAttachedDisks() GoogleComputeAttachedDiskInterface {
	return newGoogleComputeAttachedDisks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeAutoscalers() GoogleComputeAutoscalerInterface {
	return newGoogleComputeAutoscalers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeBackendBuckets() GoogleComputeBackendBucketInterface {
	return newGoogleComputeBackendBuckets(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeBackendServices() GoogleComputeBackendServiceInterface {
	return newGoogleComputeBackendServices(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeDisks() GoogleComputeDiskInterface {
	return newGoogleComputeDisks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeFirewalls() GoogleComputeFirewallInterface {
	return newGoogleComputeFirewalls(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeForwardingRules() GoogleComputeForwardingRuleInterface {
	return newGoogleComputeForwardingRules(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeGlobalAddresses() GoogleComputeGlobalAddressInterface {
	return newGoogleComputeGlobalAddresses(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeGlobalForwardingRules() GoogleComputeGlobalForwardingRuleInterface {
	return newGoogleComputeGlobalForwardingRules(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeHealthChecks() GoogleComputeHealthCheckInterface {
	return newGoogleComputeHealthChecks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeHttpHealthChecks() GoogleComputeHttpHealthCheckInterface {
	return newGoogleComputeHttpHealthChecks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeHttpsHealthChecks() GoogleComputeHttpsHealthCheckInterface {
	return newGoogleComputeHttpsHealthChecks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeImages() GoogleComputeImageInterface {
	return newGoogleComputeImages(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInstances() GoogleComputeInstanceInterface {
	return newGoogleComputeInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInstanceFromTemplates() GoogleComputeInstanceFromTemplateInterface {
	return newGoogleComputeInstanceFromTemplates(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInstanceGroups() GoogleComputeInstanceGroupInterface {
	return newGoogleComputeInstanceGroups(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInstanceGroupManagers() GoogleComputeInstanceGroupManagerInterface {
	return newGoogleComputeInstanceGroupManagers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInstanceTemplates() GoogleComputeInstanceTemplateInterface {
	return newGoogleComputeInstanceTemplates(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeInterconnectAttachments() GoogleComputeInterconnectAttachmentInterface {
	return newGoogleComputeInterconnectAttachments(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeNetworks() GoogleComputeNetworkInterface {
	return newGoogleComputeNetworks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeNetworkPeerings() GoogleComputeNetworkPeeringInterface {
	return newGoogleComputeNetworkPeerings(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeProjectMetadatas() GoogleComputeProjectMetadataInterface {
	return newGoogleComputeProjectMetadatas(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeProjectMetadataItems() GoogleComputeProjectMetadataItemInterface {
	return newGoogleComputeProjectMetadataItems(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRegionAutoscalers() GoogleComputeRegionAutoscalerInterface {
	return newGoogleComputeRegionAutoscalers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRegionBackendServices() GoogleComputeRegionBackendServiceInterface {
	return newGoogleComputeRegionBackendServices(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRegionDisks() GoogleComputeRegionDiskInterface {
	return newGoogleComputeRegionDisks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRegionInstanceGroupManagers() GoogleComputeRegionInstanceGroupManagerInterface {
	return newGoogleComputeRegionInstanceGroupManagers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRoutes() GoogleComputeRouteInterface {
	return newGoogleComputeRoutes(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRouters() GoogleComputeRouterInterface {
	return newGoogleComputeRouters(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRouterInterfaces() GoogleComputeRouterInterfaceInterface {
	return newGoogleComputeRouterInterfaces(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRouterNats() GoogleComputeRouterNatInterface {
	return newGoogleComputeRouterNats(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeRouterPeers() GoogleComputeRouterPeerInterface {
	return newGoogleComputeRouterPeers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSecurityPolicies() GoogleComputeSecurityPolicyInterface {
	return newGoogleComputeSecurityPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSharedVpcHostProjects() GoogleComputeSharedVpcHostProjectInterface {
	return newGoogleComputeSharedVpcHostProjects(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSharedVpcServiceProjects() GoogleComputeSharedVpcServiceProjectInterface {
	return newGoogleComputeSharedVpcServiceProjects(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSnapshots() GoogleComputeSnapshotInterface {
	return newGoogleComputeSnapshots(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSslCertificates() GoogleComputeSslCertificateInterface {
	return newGoogleComputeSslCertificates(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSslPolicies() GoogleComputeSslPolicyInterface {
	return newGoogleComputeSslPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSubnetworks() GoogleComputeSubnetworkInterface {
	return newGoogleComputeSubnetworks(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSubnetworkIamBindings() GoogleComputeSubnetworkIamBindingInterface {
	return newGoogleComputeSubnetworkIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSubnetworkIamMembers() GoogleComputeSubnetworkIamMemberInterface {
	return newGoogleComputeSubnetworkIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeSubnetworkIamPolicies() GoogleComputeSubnetworkIamPolicyInterface {
	return newGoogleComputeSubnetworkIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeTargetHttpProxies() GoogleComputeTargetHttpProxyInterface {
	return newGoogleComputeTargetHttpProxies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeTargetHttpsProxies() GoogleComputeTargetHttpsProxyInterface {
	return newGoogleComputeTargetHttpsProxies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeTargetPools() GoogleComputeTargetPoolInterface {
	return newGoogleComputeTargetPools(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeTargetSslProxies() GoogleComputeTargetSslProxyInterface {
	return newGoogleComputeTargetSslProxies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeTargetTcpProxies() GoogleComputeTargetTcpProxyInterface {
	return newGoogleComputeTargetTcpProxies(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeUrlMaps() GoogleComputeUrlMapInterface {
	return newGoogleComputeUrlMaps(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeVpnGateways() GoogleComputeVpnGatewayInterface {
	return newGoogleComputeVpnGateways(c)
}

func (c *GoogleV1alpha1Client) GoogleComputeVpnTunnels() GoogleComputeVpnTunnelInterface {
	return newGoogleComputeVpnTunnels(c)
}

func (c *GoogleV1alpha1Client) GoogleContainerAnalysisNotes() GoogleContainerAnalysisNoteInterface {
	return newGoogleContainerAnalysisNotes(c)
}

func (c *GoogleV1alpha1Client) GoogleContainerClusters() GoogleContainerClusterInterface {
	return newGoogleContainerClusters(c)
}

func (c *GoogleV1alpha1Client) GoogleContainerNodePools() GoogleContainerNodePoolInterface {
	return newGoogleContainerNodePools(c)
}

func (c *GoogleV1alpha1Client) GoogleDataflowJobs() GoogleDataflowJobInterface {
	return newGoogleDataflowJobs(c)
}

func (c *GoogleV1alpha1Client) GoogleDataprocClusters() GoogleDataprocClusterInterface {
	return newGoogleDataprocClusters(c)
}

func (c *GoogleV1alpha1Client) GoogleDataprocJobs() GoogleDataprocJobInterface {
	return newGoogleDataprocJobs(c)
}

func (c *GoogleV1alpha1Client) GoogleDnsManagedZones() GoogleDnsManagedZoneInterface {
	return newGoogleDnsManagedZones(c)
}

func (c *GoogleV1alpha1Client) GoogleDnsRecordSets() GoogleDnsRecordSetInterface {
	return newGoogleDnsRecordSets(c)
}

func (c *GoogleV1alpha1Client) GoogleEndpointsServices() GoogleEndpointsServiceInterface {
	return newGoogleEndpointsServices(c)
}

func (c *GoogleV1alpha1Client) GoogleFilestoreInstances() GoogleFilestoreInstanceInterface {
	return newGoogleFilestoreInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleFolders() GoogleFolderInterface {
	return newGoogleFolders(c)
}

func (c *GoogleV1alpha1Client) GoogleFolderIamBindings() GoogleFolderIamBindingInterface {
	return newGoogleFolderIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleFolderIamMembers() GoogleFolderIamMemberInterface {
	return newGoogleFolderIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleFolderIamPolicies() GoogleFolderIamPolicyInterface {
	return newGoogleFolderIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleFolderOrganizationPolicies() GoogleFolderOrganizationPolicyInterface {
	return newGoogleFolderOrganizationPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsCryptoKeys() GoogleKmsCryptoKeyInterface {
	return newGoogleKmsCryptoKeys(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsCryptoKeyIamBindings() GoogleKmsCryptoKeyIamBindingInterface {
	return newGoogleKmsCryptoKeyIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsCryptoKeyIamMembers() GoogleKmsCryptoKeyIamMemberInterface {
	return newGoogleKmsCryptoKeyIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsKeyRings() GoogleKmsKeyRingInterface {
	return newGoogleKmsKeyRings(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsKeyRingIamBindings() GoogleKmsKeyRingIamBindingInterface {
	return newGoogleKmsKeyRingIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsKeyRingIamMembers() GoogleKmsKeyRingIamMemberInterface {
	return newGoogleKmsKeyRingIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleKmsKeyRingIamPolicies() GoogleKmsKeyRingIamPolicyInterface {
	return newGoogleKmsKeyRingIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingBillingAccountExclusions() GoogleLoggingBillingAccountExclusionInterface {
	return newGoogleLoggingBillingAccountExclusions(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingBillingAccountSinks() GoogleLoggingBillingAccountSinkInterface {
	return newGoogleLoggingBillingAccountSinks(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingFolderExclusions() GoogleLoggingFolderExclusionInterface {
	return newGoogleLoggingFolderExclusions(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingFolderSinks() GoogleLoggingFolderSinkInterface {
	return newGoogleLoggingFolderSinks(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingOrganizationExclusions() GoogleLoggingOrganizationExclusionInterface {
	return newGoogleLoggingOrganizationExclusions(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingOrganizationSinks() GoogleLoggingOrganizationSinkInterface {
	return newGoogleLoggingOrganizationSinks(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingProjectExclusions() GoogleLoggingProjectExclusionInterface {
	return newGoogleLoggingProjectExclusions(c)
}

func (c *GoogleV1alpha1Client) GoogleLoggingProjectSinks() GoogleLoggingProjectSinkInterface {
	return newGoogleLoggingProjectSinks(c)
}

func (c *GoogleV1alpha1Client) GoogleMonitoringAlertPolicies() GoogleMonitoringAlertPolicyInterface {
	return newGoogleMonitoringAlertPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleMonitoringGroups() GoogleMonitoringGroupInterface {
	return newGoogleMonitoringGroups(c)
}

func (c *GoogleV1alpha1Client) GoogleMonitoringNotificationChannels() GoogleMonitoringNotificationChannelInterface {
	return newGoogleMonitoringNotificationChannels(c)
}

func (c *GoogleV1alpha1Client) GoogleMonitoringUptimeCheckConfigs() GoogleMonitoringUptimeCheckConfigInterface {
	return newGoogleMonitoringUptimeCheckConfigs(c)
}

func (c *GoogleV1alpha1Client) GoogleOrganizationIamBindings() GoogleOrganizationIamBindingInterface {
	return newGoogleOrganizationIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleOrganizationIamCustomRoles() GoogleOrganizationIamCustomRoleInterface {
	return newGoogleOrganizationIamCustomRoles(c)
}

func (c *GoogleV1alpha1Client) GoogleOrganizationIamMembers() GoogleOrganizationIamMemberInterface {
	return newGoogleOrganizationIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleOrganizationIamPolicies() GoogleOrganizationIamPolicyInterface {
	return newGoogleOrganizationIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleOrganizationPolicies() GoogleOrganizationPolicyInterface {
	return newGoogleOrganizationPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleProjects() GoogleProjectInterface {
	return newGoogleProjects(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectIamBindings() GoogleProjectIamBindingInterface {
	return newGoogleProjectIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectIamCustomRoles() GoogleProjectIamCustomRoleInterface {
	return newGoogleProjectIamCustomRoles(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectIamMembers() GoogleProjectIamMemberInterface {
	return newGoogleProjectIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectIamPolicies() GoogleProjectIamPolicyInterface {
	return newGoogleProjectIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectOrganizationPolicies() GoogleProjectOrganizationPolicyInterface {
	return newGoogleProjectOrganizationPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectServices() GoogleProjectServiceInterface {
	return newGoogleProjectServices(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectServiceses() GoogleProjectServicesInterface {
	return newGoogleProjectServiceses(c)
}

func (c *GoogleV1alpha1Client) GoogleProjectUsageExportBuckets() GoogleProjectUsageExportBucketInterface {
	return newGoogleProjectUsageExportBuckets(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubSubscriptions() GooglePubsubSubscriptionInterface {
	return newGooglePubsubSubscriptions(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubSubscriptionIamBindings() GooglePubsubSubscriptionIamBindingInterface {
	return newGooglePubsubSubscriptionIamBindings(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubSubscriptionIamMembers() GooglePubsubSubscriptionIamMemberInterface {
	return newGooglePubsubSubscriptionIamMembers(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubSubscriptionIamPolicies() GooglePubsubSubscriptionIamPolicyInterface {
	return newGooglePubsubSubscriptionIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubTopics() GooglePubsubTopicInterface {
	return newGooglePubsubTopics(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubTopicIamBindings() GooglePubsubTopicIamBindingInterface {
	return newGooglePubsubTopicIamBindings(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubTopicIamMembers() GooglePubsubTopicIamMemberInterface {
	return newGooglePubsubTopicIamMembers(c)
}

func (c *GoogleV1alpha1Client) GooglePubsubTopicIamPolicies() GooglePubsubTopicIamPolicyInterface {
	return newGooglePubsubTopicIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleRedisInstances() GoogleRedisInstanceInterface {
	return newGoogleRedisInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleResourceManagerLiens() GoogleResourceManagerLienInterface {
	return newGoogleResourceManagerLiens(c)
}

func (c *GoogleV1alpha1Client) GoogleRuntimeconfigConfigs() GoogleRuntimeconfigConfigInterface {
	return newGoogleRuntimeconfigConfigs(c)
}

func (c *GoogleV1alpha1Client) GoogleRuntimeconfigVariables() GoogleRuntimeconfigVariableInterface {
	return newGoogleRuntimeconfigVariables(c)
}

func (c *GoogleV1alpha1Client) GoogleServiceAccounts() GoogleServiceAccountInterface {
	return newGoogleServiceAccounts(c)
}

func (c *GoogleV1alpha1Client) GoogleServiceAccountIamBindings() GoogleServiceAccountIamBindingInterface {
	return newGoogleServiceAccountIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleServiceAccountIamMembers() GoogleServiceAccountIamMemberInterface {
	return newGoogleServiceAccountIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleServiceAccountIamPolicies() GoogleServiceAccountIamPolicyInterface {
	return newGoogleServiceAccountIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleServiceAccountKeys() GoogleServiceAccountKeyInterface {
	return newGoogleServiceAccountKeys(c)
}

func (c *GoogleV1alpha1Client) GoogleSourcerepoRepositories() GoogleSourcerepoRepositoryInterface {
	return newGoogleSourcerepoRepositories(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerDatabases() GoogleSpannerDatabaseInterface {
	return newGoogleSpannerDatabases(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerDatabaseIamBindings() GoogleSpannerDatabaseIamBindingInterface {
	return newGoogleSpannerDatabaseIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerDatabaseIamMembers() GoogleSpannerDatabaseIamMemberInterface {
	return newGoogleSpannerDatabaseIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerDatabaseIamPolicies() GoogleSpannerDatabaseIamPolicyInterface {
	return newGoogleSpannerDatabaseIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerInstances() GoogleSpannerInstanceInterface {
	return newGoogleSpannerInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerInstanceIamBindings() GoogleSpannerInstanceIamBindingInterface {
	return newGoogleSpannerInstanceIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerInstanceIamMembers() GoogleSpannerInstanceIamMemberInterface {
	return newGoogleSpannerInstanceIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleSpannerInstanceIamPolicies() GoogleSpannerInstanceIamPolicyInterface {
	return newGoogleSpannerInstanceIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleSqlDatabases() GoogleSqlDatabaseInterface {
	return newGoogleSqlDatabases(c)
}

func (c *GoogleV1alpha1Client) GoogleSqlDatabaseInstances() GoogleSqlDatabaseInstanceInterface {
	return newGoogleSqlDatabaseInstances(c)
}

func (c *GoogleV1alpha1Client) GoogleSqlSslCerts() GoogleSqlSslCertInterface {
	return newGoogleSqlSslCerts(c)
}

func (c *GoogleV1alpha1Client) GoogleSqlUsers() GoogleSqlUserInterface {
	return newGoogleSqlUsers(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBuckets() GoogleStorageBucketInterface {
	return newGoogleStorageBuckets(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBucketAcls() GoogleStorageBucketAclInterface {
	return newGoogleStorageBucketAcls(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBucketIamBindings() GoogleStorageBucketIamBindingInterface {
	return newGoogleStorageBucketIamBindings(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBucketIamMembers() GoogleStorageBucketIamMemberInterface {
	return newGoogleStorageBucketIamMembers(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBucketIamPolicies() GoogleStorageBucketIamPolicyInterface {
	return newGoogleStorageBucketIamPolicies(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageBucketObjects() GoogleStorageBucketObjectInterface {
	return newGoogleStorageBucketObjects(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageDefaultObjectAccessControls() GoogleStorageDefaultObjectAccessControlInterface {
	return newGoogleStorageDefaultObjectAccessControls(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageDefaultObjectAcls() GoogleStorageDefaultObjectAclInterface {
	return newGoogleStorageDefaultObjectAcls(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageNotifications() GoogleStorageNotificationInterface {
	return newGoogleStorageNotifications(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageObjectAccessControls() GoogleStorageObjectAccessControlInterface {
	return newGoogleStorageObjectAccessControls(c)
}

func (c *GoogleV1alpha1Client) GoogleStorageObjectAcls() GoogleStorageObjectAclInterface {
	return newGoogleStorageObjectAcls(c)
}

// NewForConfig creates a new GoogleV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*GoogleV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GoogleV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new GoogleV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GoogleV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GoogleV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *GoogleV1alpha1Client {
	return &GoogleV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GoogleV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
