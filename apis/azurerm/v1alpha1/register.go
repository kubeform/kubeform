package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/azurerm"
)

var SchemeGroupVersion = schema.GroupVersion{Group: azurerm.GroupName, Version: "v1alpha1"}

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

		&AzurermLbOutboundRule{},
		&AzurermLbOutboundRuleList{},

		&AzurermApiManagementProductGroup{},
		&AzurermApiManagementProductGroupList{},

		&AzurermAppServiceSlot{},
		&AzurermAppServiceSlotList{},

		&AzurermAutomationVariableInt{},
		&AzurermAutomationVariableIntList{},

		&AzurermCdnEndpoint{},
		&AzurermCdnEndpointList{},

		&AzurermDataFactory{},
		&AzurermDataFactoryList{},

		&AzurermApplicationInsightsWebTest{},
		&AzurermApplicationInsightsWebTestList{},

		&AzurermDnsZone{},
		&AzurermDnsZoneList{},

		&AzurermServicebusTopic{},
		&AzurermServicebusTopicList{},

		&AzurermSqlDatabase{},
		&AzurermSqlDatabaseList{},

		&AzurermSubnetRouteTableAssociation{},
		&AzurermSubnetRouteTableAssociationList{},

		&AzurermApiManagementApiVersionSet{},
		&AzurermApiManagementApiVersionSetList{},

		&AzurermContainerRegistry{},
		&AzurermContainerRegistryList{},

		&AzurermCosmosdbMongoDatabase{},
		&AzurermCosmosdbMongoDatabaseList{},

		&AzurermStorageTable{},
		&AzurermStorageTableList{},

		&AzurermEventhubAuthorizationRule{},
		&AzurermEventhubAuthorizationRuleList{},

		&AzurermManagementGroup{},
		&AzurermManagementGroupList{},

		&AzurermSecurityCenterSubscriptionPricing{},
		&AzurermSecurityCenterSubscriptionPricingList{},

		&AzurermNetworkSecurityGroup{},
		&AzurermNetworkSecurityGroupList{},

		&AzurermPolicySetDefinition{},
		&AzurermPolicySetDefinitionList{},

		&AzurermSchedulerJob{},
		&AzurermSchedulerJobList{},

		&AzurermLogAnalyticsSolution{},
		&AzurermLogAnalyticsSolutionList{},

		&AzurermSqlServer{},
		&AzurermSqlServerList{},

		&AzurermVirtualMachineScaleSet{},
		&AzurermVirtualMachineScaleSetList{},

		&AzurermStreamAnalyticsStreamInputEventhub{},
		&AzurermStreamAnalyticsStreamInputEventhubList{},

		&AzurermCognitiveAccount{},
		&AzurermCognitiveAccountList{},

		&AzurermDatabricksWorkspace{},
		&AzurermDatabricksWorkspaceList{},

		&AzurermDdosProtectionPlan{},
		&AzurermDdosProtectionPlanList{},

		&AzurermDnsCnameRecord{},
		&AzurermDnsCnameRecordList{},

		&AzurermRedisFirewallRule{},
		&AzurermRedisFirewallRuleList{},

		&AzurermCosmosdbCassandraKeyspace{},
		&AzurermCosmosdbCassandraKeyspaceList{},

		&AzurermLogAnalyticsWorkspace{},
		&AzurermLogAnalyticsWorkspaceList{},

		&AzurermMysqlDatabase{},
		&AzurermMysqlDatabaseList{},

		&AzurermVirtualNetworkGateway{},
		&AzurermVirtualNetworkGatewayList{},

		&AzurermMonitorAutoscaleSetting{},
		&AzurermMonitorAutoscaleSettingList{},

		&AzurermSecurityCenterWorkspace{},
		&AzurermSecurityCenterWorkspaceList{},

		&AzurermUserAssignedIdentity{},
		&AzurermUserAssignedIdentityList{},

		&AzurermAutomationDscNodeconfiguration{},
		&AzurermAutomationDscNodeconfigurationList{},

		&AzurermDevTestWindowsVirtualMachine{},
		&AzurermDevTestWindowsVirtualMachineList{},

		&AzurermEventgridDomain{},
		&AzurermEventgridDomainList{},

		&AzurermKeyVaultAccessPolicy{},
		&AzurermKeyVaultAccessPolicyList{},

		&AzurermKeyVaultKey{},
		&AzurermKeyVaultKeyList{},

		&AzurermAppServiceActiveSlot{},
		&AzurermAppServiceActiveSlotList{},

		&AzurermApplicationGateway{},
		&AzurermApplicationGatewayList{},

		&AzurermFirewallApplicationRuleCollection{},
		&AzurermFirewallApplicationRuleCollectionList{},

		&AzurermMysqlVirtualNetworkRule{},
		&AzurermMysqlVirtualNetworkRuleList{},

		&AzurermRecoveryServicesProtectionPolicyVm{},
		&AzurermRecoveryServicesProtectionPolicyVmList{},

		&AzurermTrafficManagerProfile{},
		&AzurermTrafficManagerProfileList{},

		&AzurermApiManagementCertificate{},
		&AzurermApiManagementCertificateList{},

		&AzurermCdnProfile{},
		&AzurermCdnProfileList{},

		&AzurermDevTestPolicy{},
		&AzurermDevTestPolicyList{},

		&AzurermFunctionApp{},
		&AzurermFunctionAppList{},

		&AzurermKeyVaultSecret{},
		&AzurermKeyVaultSecretList{},

		&AzurermVirtualMachineExtension{},
		&AzurermVirtualMachineExtensionList{},

		&AzurermBatchAccount{},
		&AzurermBatchAccountList{},

		&AzurermHdinsightRserverCluster{},
		&AzurermHdinsightRserverClusterList{},

		&AzurermMonitorActivityLogAlert{},
		&AzurermMonitorActivityLogAlertList{},

		&AzurermNetworkInterfaceBackendAddressPoolAssociation{},
		&AzurermNetworkInterfaceBackendAddressPoolAssociationList{},

		&AzurermPublicIpPrefix{},
		&AzurermPublicIpPrefixList{},

		&AzurermApiManagementLogger{},
		&AzurermApiManagementLoggerList{},

		&AzurermApiManagementProduct{},
		&AzurermApiManagementProductList{},

		&AzurermAppServicePlan{},
		&AzurermAppServicePlanList{},

		&AzurermLogicAppActionCustom{},
		&AzurermLogicAppActionCustomList{},

		&AzurermServicebusTopicAuthorizationRule{},
		&AzurermServicebusTopicAuthorizationRuleList{},

		&AzurermApiManagementApiOperationPolicy{},
		&AzurermApiManagementApiOperationPolicyList{},

		&AzurermDataLakeStoreFile{},
		&AzurermDataLakeStoreFileList{},

		&AzurermNetworkInterfaceApplicationSecurityGroupAssociation{},
		&AzurermNetworkInterfaceApplicationSecurityGroupAssociationList{},

		&AzurermPolicyDefinition{},
		&AzurermPolicyDefinitionList{},

		&AzurermDataFactoryLinkedServiceSqlServer{},
		&AzurermDataFactoryLinkedServiceSqlServerList{},

		&AzurermDnsTxtRecord{},
		&AzurermDnsTxtRecordList{},

		&AzurermNetworkConnectionMonitor{},
		&AzurermNetworkConnectionMonitorList{},

		&AzurermDevspaceController{},
		&AzurermDevspaceControllerList{},

		&AzurermNotificationHub{},
		&AzurermNotificationHubList{},

		&AzurermPostgresqlDatabase{},
		&AzurermPostgresqlDatabaseList{},

		&AzurermSecurityCenterContact{},
		&AzurermSecurityCenterContactList{},

		&AzurermSqlActiveDirectoryAdministrator{},
		&AzurermSqlActiveDirectoryAdministratorList{},

		&AzurermCosmosdbSqlDatabase{},
		&AzurermCosmosdbSqlDatabaseList{},

		&AzurermDnsCaaRecord{},
		&AzurermDnsCaaRecordList{},

		&AzurermMonitorActionGroup{},
		&AzurermMonitorActionGroupList{},

		&AzurermRecoveryServicesVault{},
		&AzurermRecoveryServicesVaultList{},

		&AzurermStreamAnalyticsJob{},
		&AzurermStreamAnalyticsJobList{},

		&AzurermVirtualMachineDataDiskAttachment{},
		&AzurermVirtualMachineDataDiskAttachmentList{},

		&AzurermAutoscaleSetting{},
		&AzurermAutoscaleSettingList{},

		&AzurermDataFactoryLinkedServiceDataLakeStorageGen2{},
		&AzurermDataFactoryLinkedServiceDataLakeStorageGen2List{},

		&AzurermHdinsightHbaseCluster{},
		&AzurermHdinsightHbaseClusterList{},

		&AzurermLbRule{},
		&AzurermLbRuleList{},

		&AzurermNotificationHubNamespace{},
		&AzurermNotificationHubNamespaceList{},

		&AzurermApiManagementProductApi{},
		&AzurermApiManagementProductApiList{},

		&AzurermBatchCertificate{},
		&AzurermBatchCertificateList{},

		&AzurermPacketCapture{},
		&AzurermPacketCaptureList{},

		&AzurermServicebusSubscriptionRule{},
		&AzurermServicebusSubscriptionRuleList{},

		&AzurermDataLakeAnalyticsFirewallRule{},
		&AzurermDataLakeAnalyticsFirewallRuleList{},

		&AzurermStreamAnalyticsOutputEventhub{},
		&AzurermStreamAnalyticsOutputEventhubList{},

		&AzurermHdinsightInteractiveQueryCluster{},
		&AzurermHdinsightInteractiveQueryClusterList{},

		&AzurermHdinsightMlServicesCluster{},
		&AzurermHdinsightMlServicesClusterList{},

		&AzurermLogAnalyticsWorkspaceLinkedService{},
		&AzurermLogAnalyticsWorkspaceLinkedServiceList{},

		&AzurermRecoveryServicesProtectedVm{},
		&AzurermRecoveryServicesProtectedVmList{},

		&AzurermStreamAnalyticsOutputMssql{},
		&AzurermStreamAnalyticsOutputMssqlList{},

		&AzurermApiManagementAuthorizationServer{},
		&AzurermApiManagementAuthorizationServerList{},

		&AzurermLogicAppTriggerCustom{},
		&AzurermLogicAppTriggerCustomList{},

		&AzurermVirtualMachine{},
		&AzurermVirtualMachineList{},

		&AzurermIothubConsumerGroup{},
		&AzurermIothubConsumerGroupList{},

		&AzurermLbNatPool{},
		&AzurermLbNatPoolList{},

		&AzurermPostgresqlVirtualNetworkRule{},
		&AzurermPostgresqlVirtualNetworkRuleList{},

		&AzurermRelayNamespace{},
		&AzurermRelayNamespaceList{},

		&AzurermServicebusNamespaceAuthorizationRule{},
		&AzurermServicebusNamespaceAuthorizationRuleList{},

		&AzurermAutomationSchedule{},
		&AzurermAutomationScheduleList{},

		&AzurermAzureadServicePrincipal{},
		&AzurermAzureadServicePrincipalList{},

		&AzurermDnsMxRecord{},
		&AzurermDnsMxRecordList{},

		&AzurermEventhubNamespace{},
		&AzurermEventhubNamespaceList{},

		&AzurermNetworkPacketCapture{},
		&AzurermNetworkPacketCaptureList{},

		&AzurermBatchPool{},
		&AzurermBatchPoolList{},

		&AzurermDataFactoryLinkedServiceMysql{},
		&AzurermDataFactoryLinkedServiceMysqlList{},

		&AzurermDnsPtrRecord{},
		&AzurermDnsPtrRecordList{},

		&AzurermLogicAppTriggerRecurrence{},
		&AzurermLogicAppTriggerRecurrenceList{},

		&AzurermMysqlServer{},
		&AzurermMysqlServerList{},

		&AzurermRedisCache{},
		&AzurermRedisCacheList{},

		&AzurermStorageQueue{},
		&AzurermStorageQueueList{},

		&AzurermContainerService{},
		&AzurermContainerServiceList{},

		&AzurermDevTestLab{},
		&AzurermDevTestLabList{},

		&AzurermDnsSrvRecord{},
		&AzurermDnsSrvRecordList{},

		&AzurermFirewall{},
		&AzurermFirewallList{},

		&AzurermIothubSharedAccessPolicy{},
		&AzurermIothubSharedAccessPolicyList{},

		&AzurermNetworkProfile{},
		&AzurermNetworkProfileList{},

		&AzurermDnsARecord{},
		&AzurermDnsARecordList{},

		&AzurermStorageShare{},
		&AzurermStorageShareList{},

		&AzurermStreamAnalyticsFunctionJavascriptUdf{},
		&AzurermStreamAnalyticsFunctionJavascriptUdfList{},

		&AzurermTrafficManagerEndpoint{},
		&AzurermTrafficManagerEndpointList{},

		&AzurermLogicAppWorkflow{},
		&AzurermLogicAppWorkflowList{},

		&AzurermMariadbServer{},
		&AzurermMariadbServerList{},

		&AzurermNotificationHubAuthorizationRule{},
		&AzurermNotificationHubAuthorizationRuleList{},

		&AzurermLb{},
		&AzurermLbList{},

		&AzurermLogAnalyticsLinkedService{},
		&AzurermLogAnalyticsLinkedServiceList{},

		&AzurermStreamAnalyticsStreamInputIothub{},
		&AzurermStreamAnalyticsStreamInputIothubList{},

		&AzurermLogicAppTriggerHttpRequest{},
		&AzurermLogicAppTriggerHttpRequestList{},

		&AzurermSubnet{},
		&AzurermSubnetList{},

		&AzurermApplicationInsightsApiKey{},
		&AzurermApplicationInsightsApiKeyList{},

		&AzurermAzureadApplication{},
		&AzurermAzureadApplicationList{},

		&AzurermApiManagementGroup{},
		&AzurermApiManagementGroupList{},

		&AzurermRoleDefinition{},
		&AzurermRoleDefinitionList{},

		&AzurermRoute{},
		&AzurermRouteList{},

		&AzurermSchedulerJobCollection{},
		&AzurermSchedulerJobCollectionList{},

		&AzurermNetworkDdosProtectionPlan{},
		&AzurermNetworkDdosProtectionPlanList{},

		&AzurermSqlFirewallRule{},
		&AzurermSqlFirewallRuleList{},

		&AzurermHdinsightHadoopCluster{},
		&AzurermHdinsightHadoopClusterList{},

		&AzurermMariadbDatabase{},
		&AzurermMariadbDatabaseList{},

		&AzurermServicebusNamespace{},
		&AzurermServicebusNamespaceList{},

		&AzurermApplicationSecurityGroup{},
		&AzurermApplicationSecurityGroupList{},

		&AzurermAutomationAccount{},
		&AzurermAutomationAccountList{},

		&AzurermAzureadServicePrincipalPassword{},
		&AzurermAzureadServicePrincipalPasswordList{},

		&AzurermCosmosdbAccount{},
		&AzurermCosmosdbAccountList{},

		&AzurermDataFactoryDatasetMysql{},
		&AzurermDataFactoryDatasetMysqlList{},

		&AzurermConnectionMonitor{},
		&AzurermConnectionMonitorList{},

		&AzurermMediaServicesAccount{},
		&AzurermMediaServicesAccountList{},

		&AzurermSignalrService{},
		&AzurermSignalrServiceList{},

		&AzurermNetworkInterface{},
		&AzurermNetworkInterfaceList{},

		&AzurermNetworkSecurityRule{},
		&AzurermNetworkSecurityRuleList{},

		&AzurermApiManagementSubscription{},
		&AzurermApiManagementSubscriptionList{},

		&AzurermFirewallNatRuleCollection{},
		&AzurermFirewallNatRuleCollectionList{},

		&AzurermIotDpsCertificate{},
		&AzurermIotDpsCertificateList{},

		&AzurermMariadbFirewallRule{},
		&AzurermMariadbFirewallRuleList{},

		&AzurermMysqlFirewallRule{},
		&AzurermMysqlFirewallRuleList{},

		&AzurermDnsNsRecord{},
		&AzurermDnsNsRecordList{},

		&AzurermKeyVault{},
		&AzurermKeyVaultList{},

		&AzurermLbProbe{},
		&AzurermLbProbeList{},

		&AzurermPostgresqlFirewallRule{},
		&AzurermPostgresqlFirewallRuleList{},

		&AzurermEventgridTopic{},
		&AzurermEventgridTopicList{},

		&AzurermManagedDisk{},
		&AzurermManagedDiskList{},

		&AzurermManagementLock{},
		&AzurermManagementLockList{},

		&AzurermMysqlConfiguration{},
		&AzurermMysqlConfigurationList{},

		&AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&AzurermMonitorMetricAlertrule{},
		&AzurermMonitorMetricAlertruleList{},

		&AzurermCosmosdbMongoCollection{},
		&AzurermCosmosdbMongoCollectionList{},

		&AzurermMssqlElasticpool{},
		&AzurermMssqlElasticpoolList{},

		&AzurermServicebusQueueAuthorizationRule{},
		&AzurermServicebusQueueAuthorizationRuleList{},

		&AzurermApiManagementGroupUser{},
		&AzurermApiManagementGroupUserList{},

		&AzurermAppService{},
		&AzurermAppServiceList{},

		&AzurermAutomationVariableString{},
		&AzurermAutomationVariableStringList{},

		&AzurermIotDps{},
		&AzurermIotDpsList{},

		&AzurermIothub{},
		&AzurermIothubList{},

		&AzurermSqlVirtualNetworkRule{},
		&AzurermSqlVirtualNetworkRuleList{},

		&AzurermVirtualNetwork{},
		&AzurermVirtualNetworkList{},

		&AzurermHdinsightKafkaCluster{},
		&AzurermHdinsightKafkaClusterList{},

		&AzurermNetworkWatcher{},
		&AzurermNetworkWatcherList{},

		&AzurermSubnetNetworkSecurityGroupAssociation{},
		&AzurermSubnetNetworkSecurityGroupAssociationList{},

		&AzurermDnsAaaaRecord{},
		&AzurermDnsAaaaRecordList{},

		&AzurermSharedImageGallery{},
		&AzurermSharedImageGalleryList{},

		&AzurermSqlElasticpool{},
		&AzurermSqlElasticpoolList{},

		&AzurermDevTestVirtualNetwork{},
		&AzurermDevTestVirtualNetworkList{},

		&AzurermImage{},
		&AzurermImageList{},

		&AzurermServicebusSubscription{},
		&AzurermServicebusSubscriptionList{},

		&AzurermStorageBlob{},
		&AzurermStorageBlobList{},

		&AzurermApiManagement{},
		&AzurermApiManagementList{},

		&AzurermApiManagementApiOperation{},
		&AzurermApiManagementApiOperationList{},

		&AzurermDataFactoryLinkedServicePostgresql{},
		&AzurermDataFactoryLinkedServicePostgresqlList{},

		&AzurermExpressRouteCircuitPeering{},
		&AzurermExpressRouteCircuitPeeringList{},

		&AzurermHdinsightSparkCluster{},
		&AzurermHdinsightSparkClusterList{},

		&AzurermAppServiceCustomHostnameBinding{},
		&AzurermAppServiceCustomHostnameBindingList{},

		&AzurermLbNatRule{},
		&AzurermLbNatRuleList{},

		&AzurermApiManagementProductPolicy{},
		&AzurermApiManagementProductPolicyList{},

		&AzurermMonitorDiagnosticSetting{},
		&AzurermMonitorDiagnosticSettingList{},

		&AzurermServiceFabricCluster{},
		&AzurermServiceFabricClusterList{},

		&AzurermServicebusQueue{},
		&AzurermServicebusQueueList{},

		&AzurermSharedImageVersion{},
		&AzurermSharedImageVersionList{},

		&AzurermAutomationRunbook{},
		&AzurermAutomationRunbookList{},

		&AzurermDataLakeAnalyticsAccount{},
		&AzurermDataLakeAnalyticsAccountList{},

		&AzurermEventgridEventSubscription{},
		&AzurermEventgridEventSubscriptionList{},

		&AzurermExpressRouteCircuitAuthorization{},
		&AzurermExpressRouteCircuitAuthorizationList{},

		&AzurermLogicAppActionHttp{},
		&AzurermLogicAppActionHttpList{},

		&AzurermSnapshot{},
		&AzurermSnapshotList{},

		&AzurermAutomationVariableDatetime{},
		&AzurermAutomationVariableDatetimeList{},

		&AzurermCosmosdbTable{},
		&AzurermCosmosdbTableList{},

		&AzurermEventhub{},
		&AzurermEventhubList{},

		&AzurermKubernetesCluster{},
		&AzurermKubernetesClusterList{},

		&AzurermVirtualNetworkPeering{},
		&AzurermVirtualNetworkPeeringList{},

		&AzurermNetworkInterfaceNatRuleAssociation{},
		&AzurermNetworkInterfaceNatRuleAssociationList{},

		&AzurermPolicyAssignment{},
		&AzurermPolicyAssignmentList{},

		&AzurermPublicIp{},
		&AzurermPublicIpList{},

		&AzurermRoleAssignment{},
		&AzurermRoleAssignmentList{},

		&AzurermStreamAnalyticsOutputServicebusQueue{},
		&AzurermStreamAnalyticsOutputServicebusQueueList{},

		&AzurermPostgresqlServer{},
		&AzurermPostgresqlServerList{},

		&AzurermDataFactoryDatasetPostgresql{},
		&AzurermDataFactoryDatasetPostgresqlList{},

		&AzurermDataLakeStoreFirewallRule{},
		&AzurermDataLakeStoreFirewallRuleList{},

		&AzurermFirewallNetworkRuleCollection{},
		&AzurermFirewallNetworkRuleCollectionList{},

		&AzurermHdinsightStormCluster{},
		&AzurermHdinsightStormClusterList{},

		&AzurermLbBackendAddressPool{},
		&AzurermLbBackendAddressPoolList{},

		&AzurermApiManagementApi{},
		&AzurermApiManagementApiList{},

		&AzurermApiManagementProperty{},
		&AzurermApiManagementPropertyList{},

		&AzurermAutomationCredential{},
		&AzurermAutomationCredentialList{},

		&AzurermMonitorLogProfile{},
		&AzurermMonitorLogProfileList{},

		&AzurermPostgresqlConfiguration{},
		&AzurermPostgresqlConfigurationList{},

		&AzurermApiManagementApiSchema{},
		&AzurermApiManagementApiSchemaList{},

		&AzurermApiManagementUser{},
		&AzurermApiManagementUserList{},

		&AzurermDevTestLinuxVirtualMachine{},
		&AzurermDevTestLinuxVirtualMachineList{},

		&AzurermKeyVaultCertificate{},
		&AzurermKeyVaultCertificateList{},

		&AzurermMetricAlertrule{},
		&AzurermMetricAlertruleList{},

		&AzurermDataFactoryPipeline{},
		&AzurermDataFactoryPipelineList{},

		&AzurermEventhubNamespaceAuthorizationRule{},
		&AzurermEventhubNamespaceAuthorizationRuleList{},

		&AzurermMonitorMetricAlert{},
		&AzurermMonitorMetricAlertList{},

		&AzurermStreamAnalyticsOutputBlob{},
		&AzurermStreamAnalyticsOutputBlobList{},

		&AzurermStreamAnalyticsStreamInputBlob{},
		&AzurermStreamAnalyticsStreamInputBlobList{},

		&AzurermStorageContainer{},
		&AzurermStorageContainerList{},

		&AzurermVirtualNetworkGatewayConnection{},
		&AzurermVirtualNetworkGatewayConnectionList{},

		&AzurermApplicationInsights{},
		&AzurermApplicationInsightsList{},

		&AzurermAvailabilitySet{},
		&AzurermAvailabilitySetList{},

		&AzurermDataLakeStore{},
		&AzurermDataLakeStoreList{},

		&AzurermLocalNetworkGateway{},
		&AzurermLocalNetworkGatewayList{},

		&AzurermPrivateDnsZone{},
		&AzurermPrivateDnsZoneList{},

		&AzurermResourceGroup{},
		&AzurermResourceGroupList{},

		&AzurermRouteTable{},
		&AzurermRouteTableList{},

		&AzurermSearchService{},
		&AzurermSearchServiceList{},

		&AzurermApiManagementOpenidConnectProvider{},
		&AzurermApiManagementOpenidConnectProviderList{},

		&AzurermAutomationModule{},
		&AzurermAutomationModuleList{},

		&AzurermAutomationVariableBool{},
		&AzurermAutomationVariableBoolList{},

		&AzurermDataFactoryDatasetSqlServerTable{},
		&AzurermDataFactoryDatasetSqlServerTableList{},

		&AzurermExpressRouteCircuit{},
		&AzurermExpressRouteCircuitList{},

		&AzurermStorageAccount{},
		&AzurermStorageAccountList{},

		&AzurermTemplateDeployment{},
		&AzurermTemplateDeploymentList{},

		&AzurermApiManagementApiPolicy{},
		&AzurermApiManagementApiPolicyList{},

		&AzurermAutomationDscConfiguration{},
		&AzurermAutomationDscConfigurationList{},

		&AzurermContainerGroup{},
		&AzurermContainerGroupList{},

		&AzurermEventhubConsumerGroup{},
		&AzurermEventhubConsumerGroupList{},

		&AzurermSharedImage{},
		&AzurermSharedImageList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
