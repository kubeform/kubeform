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
    
        &AzurermSubnetNetworkSecurityGroupAssociation{},
        &AzurermSubnetNetworkSecurityGroupAssociationList{},
    
        &AzurermApiManagementAuthorizationServer{},
        &AzurermApiManagementAuthorizationServerList{},
    
        &AzurermCdnProfile{},
        &AzurermCdnProfileList{},
    
        &AzurermEventhubNamespace{},
        &AzurermEventhubNamespaceList{},
    
        &AzurermNotificationHubAuthorizationRule{},
        &AzurermNotificationHubAuthorizationRuleList{},
    
        &AzurermSqlActiveDirectoryAdministrator{},
        &AzurermSqlActiveDirectoryAdministratorList{},
    
        &AzurermDnsSrvRecord{},
        &AzurermDnsSrvRecordList{},
    
        &AzurermKeyVault{},
        &AzurermKeyVaultList{},
    
        &AzurermNetworkInterface{},
        &AzurermNetworkInterfaceList{},
    
        &AzurermServiceFabricCluster{},
        &AzurermServiceFabricClusterList{},
    
        &AzurermServicebusTopicAuthorizationRule{},
        &AzurermServicebusTopicAuthorizationRuleList{},
    
        &AzurermBatchPool{},
        &AzurermBatchPoolList{},
    
        &AzurermExpressRouteCircuit{},
        &AzurermExpressRouteCircuitList{},
    
        &AzurermSqlElasticpool{},
        &AzurermSqlElasticpoolList{},
    
        &AzurermApiManagement{},
        &AzurermApiManagementList{},
    
        &AzurermConnectionMonitor{},
        &AzurermConnectionMonitorList{},
    
        &AzurermPublicIpPrefix{},
        &AzurermPublicIpPrefixList{},
    
        &AzurermFirewallNatRuleCollection{},
        &AzurermFirewallNatRuleCollectionList{},
    
        &AzurermHdinsightKafkaCluster{},
        &AzurermHdinsightKafkaClusterList{},
    
        &AzurermSignalrService{},
        &AzurermSignalrServiceList{},
    
        &AzurermVirtualNetworkPeering{},
        &AzurermVirtualNetworkPeeringList{},
    
        &AzurermCosmosdbTable{},
        &AzurermCosmosdbTableList{},
    
        &AzurermPostgresqlFirewallRule{},
        &AzurermPostgresqlFirewallRuleList{},
    
        &AzurermSecurityCenterWorkspace{},
        &AzurermSecurityCenterWorkspaceList{},
    
        &AzurermTrafficManagerEndpoint{},
        &AzurermTrafficManagerEndpointList{},
    
        &AzurermUserAssignedIdentity{},
        &AzurermUserAssignedIdentityList{},
    
        &AzurermSharedImageVersion{},
        &AzurermSharedImageVersionList{},
    
        &AzurermStorageShare{},
        &AzurermStorageShareList{},
    
        &AzurermAzureadApplication{},
        &AzurermAzureadApplicationList{},
    
        &AzurermCosmosdbMongoDatabase{},
        &AzurermCosmosdbMongoDatabaseList{},
    
        &AzurermDataLakeStore{},
        &AzurermDataLakeStoreList{},
    
        &AzurermLogicAppActionHttp{},
        &AzurermLogicAppActionHttpList{},
    
        &AzurermPostgresqlVirtualNetworkRule{},
        &AzurermPostgresqlVirtualNetworkRuleList{},
    
        &AzurermApiManagementOpenidConnectProvider{},
        &AzurermApiManagementOpenidConnectProviderList{},
    
        &AzurermImage{},
        &AzurermImageList{},
    
        &AzurermHdinsightInteractiveQueryCluster{},
        &AzurermHdinsightInteractiveQueryClusterList{},
    
        &AzurermMonitorLogProfile{},
        &AzurermMonitorLogProfileList{},
    
        &AzurermNetworkInterfaceBackendAddressPoolAssociation{},
        &AzurermNetworkInterfaceBackendAddressPoolAssociationList{},
    
        &AzurermServicebusSubscriptionRule{},
        &AzurermServicebusSubscriptionRuleList{},
    
        &AzurermApiManagementUser{},
        &AzurermApiManagementUserList{},
    
        &AzurermApplicationInsights{},
        &AzurermApplicationInsightsList{},
    
        &AzurermLogAnalyticsWorkspace{},
        &AzurermLogAnalyticsWorkspaceList{},
    
        &AzurermMysqlServer{},
        &AzurermMysqlServerList{},
    
        &AzurermNetworkDdosProtectionPlan{},
        &AzurermNetworkDdosProtectionPlanList{},
    
        &AzurermNetworkPacketCapture{},
        &AzurermNetworkPacketCaptureList{},
    
        &AzurermStreamAnalyticsOutputEventhub{},
        &AzurermStreamAnalyticsOutputEventhubList{},
    
        &AzurermTrafficManagerProfile{},
        &AzurermTrafficManagerProfileList{},
    
        &AzurermApiManagementCertificate{},
        &AzurermApiManagementCertificateList{},
    
        &AzurermAppServiceActiveSlot{},
        &AzurermAppServiceActiveSlotList{},
    
        &AzurermDnsPtrRecord{},
        &AzurermDnsPtrRecordList{},
    
        &AzurermKeyVaultKey{},
        &AzurermKeyVaultKeyList{},
    
        &AzurermMetricAlertrule{},
        &AzurermMetricAlertruleList{},
    
        &AzurermApiManagementProductPolicy{},
        &AzurermApiManagementProductPolicyList{},
    
        &AzurermDnsMxRecord{},
        &AzurermDnsMxRecordList{},
    
        &AzurermNetworkConnectionMonitor{},
        &AzurermNetworkConnectionMonitorList{},
    
        &AzurermRedisFirewallRule{},
        &AzurermRedisFirewallRuleList{},
    
        &AzurermAutomationModule{},
        &AzurermAutomationModuleList{},
    
        &AzurermCosmosdbMongoCollection{},
        &AzurermCosmosdbMongoCollectionList{},
    
        &AzurermVirtualNetworkGatewayConnection{},
        &AzurermVirtualNetworkGatewayConnectionList{},
    
        &AzurermApiManagementGroupUser{},
        &AzurermApiManagementGroupUserList{},
    
        &AzurermDevTestVirtualNetwork{},
        &AzurermDevTestVirtualNetworkList{},
    
        &AzurermEventhub{},
        &AzurermEventhubList{},
    
        &AzurermPolicyAssignment{},
        &AzurermPolicyAssignmentList{},
    
        &AzurermDataFactoryPipeline{},
        &AzurermDataFactoryPipelineList{},
    
        &AzurermSubnet{},
        &AzurermSubnetList{},
    
        &AzurermBatchCertificate{},
        &AzurermBatchCertificateList{},
    
        &AzurermMariadbServer{},
        &AzurermMariadbServerList{},
    
        &AzurermMonitorDiagnosticSetting{},
        &AzurermMonitorDiagnosticSettingList{},
    
        &AzurermContainerGroup{},
        &AzurermContainerGroupList{},
    
        &AzurermLbOutboundRule{},
        &AzurermLbOutboundRuleList{},
    
        &AzurermApplicationInsightsWebTest{},
        &AzurermApplicationInsightsWebTestList{},
    
        &AzurermDataFactoryLinkedServicePostgresql{},
        &AzurermDataFactoryLinkedServicePostgresqlList{},
    
        &AzurermPostgresqlServer{},
        &AzurermPostgresqlServerList{},
    
        &AzurermSharedImage{},
        &AzurermSharedImageList{},
    
        &AzurermMonitorMetricAlertrule{},
        &AzurermMonitorMetricAlertruleList{},
    
        &AzurermNetworkSecurityGroup{},
        &AzurermNetworkSecurityGroupList{},
    
        &AzurermRecoveryServicesProtectedVm{},
        &AzurermRecoveryServicesProtectedVmList{},
    
        &AzurermDdosProtectionPlan{},
        &AzurermDdosProtectionPlanList{},
    
        &AzurermDnsAaaaRecord{},
        &AzurermDnsAaaaRecordList{},
    
        &AzurermDnsNsRecord{},
        &AzurermDnsNsRecordList{},
    
        &AzurermLbNatRule{},
        &AzurermLbNatRuleList{},
    
        &AzurermMonitorMetricAlert{},
        &AzurermMonitorMetricAlertList{},
    
        &AzurermStreamAnalyticsFunctionJavascriptUdf{},
        &AzurermStreamAnalyticsFunctionJavascriptUdfList{},
    
        &AzurermStorageQueue{},
        &AzurermStorageQueueList{},
    
        &AzurermExpressRouteCircuitPeering{},
        &AzurermExpressRouteCircuitPeeringList{},
    
        &AzurermLogAnalyticsSolution{},
        &AzurermLogAnalyticsSolutionList{},
    
        &AzurermPrivateDnsZone{},
        &AzurermPrivateDnsZoneList{},
    
        &AzurermRelayNamespace{},
        &AzurermRelayNamespaceList{},
    
        &AzurermStreamAnalyticsStreamInputBlob{},
        &AzurermStreamAnalyticsStreamInputBlobList{},
    
        &AzurermPostgresqlDatabase{},
        &AzurermPostgresqlDatabaseList{},
    
        &AzurermAutomationDscConfiguration{},
        &AzurermAutomationDscConfigurationList{},
    
        &AzurermAutomationVariableDatetime{},
        &AzurermAutomationVariableDatetimeList{},
    
        &AzurermAvailabilitySet{},
        &AzurermAvailabilitySetList{},
    
        &AzurermCognitiveAccount{},
        &AzurermCognitiveAccountList{},
    
        &AzurermLogicAppTriggerCustom{},
        &AzurermLogicAppTriggerCustomList{},
    
        &AzurermEventgridDomain{},
        &AzurermEventgridDomainList{},
    
        &AzurermHdinsightStormCluster{},
        &AzurermHdinsightStormClusterList{},
    
        &AzurermVirtualNetwork{},
        &AzurermVirtualNetworkList{},
    
        &AzurermApiManagementProductGroup{},
        &AzurermApiManagementProductGroupList{},
    
        &AzurermLogAnalyticsLinkedService{},
        &AzurermLogAnalyticsLinkedServiceList{},
    
        &AzurermSchedulerJobCollection{},
        &AzurermSchedulerJobCollectionList{},
    
        &AzurermMediaServicesAccount{},
        &AzurermMediaServicesAccountList{},
    
        &AzurermMssqlElasticpool{},
        &AzurermMssqlElasticpoolList{},
    
        &AzurermSqlServer{},
        &AzurermSqlServerList{},
    
        &AzurermBatchAccount{},
        &AzurermBatchAccountList{},
    
        &AzurermHdinsightSparkCluster{},
        &AzurermHdinsightSparkClusterList{},
    
        &AzurermLbBackendAddressPool{},
        &AzurermLbBackendAddressPoolList{},
    
        &AzurermLogicAppWorkflow{},
        &AzurermLogicAppWorkflowList{},
    
        &AzurermMariadbDatabase{},
        &AzurermMariadbDatabaseList{},
    
        &AzurermVirtualMachineDataDiskAttachment{},
        &AzurermVirtualMachineDataDiskAttachmentList{},
    
        &AzurermLogicAppActionCustom{},
        &AzurermLogicAppActionCustomList{},
    
        &AzurermLogicAppTriggerHttpRequest{},
        &AzurermLogicAppTriggerHttpRequestList{},
    
        &AzurermRecoveryServicesVault{},
        &AzurermRecoveryServicesVaultList{},
    
        &AzurermVirtualNetworkGateway{},
        &AzurermVirtualNetworkGatewayList{},
    
        &AzurermDnsCaaRecord{},
        &AzurermDnsCaaRecordList{},
    
        &AzurermMysqlFirewallRule{},
        &AzurermMysqlFirewallRuleList{},
    
        &AzurermNotificationHub{},
        &AzurermNotificationHubList{},
    
        &AzurermRouteTable{},
        &AzurermRouteTableList{},
    
        &AzurermSqlFirewallRule{},
        &AzurermSqlFirewallRuleList{},
    
        &AzurermFirewallApplicationRuleCollection{},
        &AzurermFirewallApplicationRuleCollectionList{},
    
        &AzurermLogicAppTriggerRecurrence{},
        &AzurermLogicAppTriggerRecurrenceList{},
    
        &AzurermStreamAnalyticsJob{},
        &AzurermStreamAnalyticsJobList{},
    
        &AzurermDnsARecord{},
        &AzurermDnsARecordList{},
    
        &AzurermNetworkInterfaceApplicationSecurityGroupAssociation{},
        &AzurermNetworkInterfaceApplicationSecurityGroupAssociationList{},
    
        &AzurermApplicationSecurityGroup{},
        &AzurermApplicationSecurityGroupList{},
    
        &AzurermAutomationCredential{},
        &AzurermAutomationCredentialList{},
    
        &AzurermMysqlConfiguration{},
        &AzurermMysqlConfigurationList{},
    
        &AzurermServicebusTopic{},
        &AzurermServicebusTopicList{},
    
        &AzurermStreamAnalyticsStreamInputIothub{},
        &AzurermStreamAnalyticsStreamInputIothubList{},
    
        &AzurermApiManagementApiOperationPolicy{},
        &AzurermApiManagementApiOperationPolicyList{},
    
        &AzurermApplicationGateway{},
        &AzurermApplicationGatewayList{},
    
        &AzurermIothubConsumerGroup{},
        &AzurermIothubConsumerGroupList{},
    
        &AzurermRoute{},
        &AzurermRouteList{},
    
        &AzurermSnapshot{},
        &AzurermSnapshotList{},
    
        &AzurermAutomationAccount{},
        &AzurermAutomationAccountList{},
    
        &AzurermFirewallNetworkRuleCollection{},
        &AzurermFirewallNetworkRuleCollectionList{},
    
        &AzurermHdinsightHadoopCluster{},
        &AzurermHdinsightHadoopClusterList{},
    
        &AzurermAutomationSchedule{},
        &AzurermAutomationScheduleList{},
    
        &AzurermEventgridEventSubscription{},
        &AzurermEventgridEventSubscriptionList{},
    
        &AzurermMonitorAutoscaleSetting{},
        &AzurermMonitorAutoscaleSettingList{},
    
        &AzurermStorageTable{},
        &AzurermStorageTableList{},
    
        &AzurermAppServiceCustomHostnameBinding{},
        &AzurermAppServiceCustomHostnameBindingList{},
    
        &AzurermHdinsightRserverCluster{},
        &AzurermHdinsightRserverClusterList{},
    
        &AzurermKeyVaultCertificate{},
        &AzurermKeyVaultCertificateList{},
    
        &AzurermKeyVaultSecret{},
        &AzurermKeyVaultSecretList{},
    
        &AzurermPolicyDefinition{},
        &AzurermPolicyDefinitionList{},
    
        &AzurermRoleDefinition{},
        &AzurermRoleDefinitionList{},
    
        &AzurermStreamAnalyticsStreamInputEventhub{},
        &AzurermStreamAnalyticsStreamInputEventhubList{},
    
        &AzurermApiManagementApiVersionSet{},
        &AzurermApiManagementApiVersionSetList{},
    
        &AzurermCosmosdbAccount{},
        &AzurermCosmosdbAccountList{},
    
        &AzurermDevTestLab{},
        &AzurermDevTestLabList{},
    
        &AzurermEventgridTopic{},
        &AzurermEventgridTopicList{},
    
        &AzurermIotDps{},
        &AzurermIotDpsList{},
    
        &AzurermApiManagementLogger{},
        &AzurermApiManagementLoggerList{},
    
        &AzurermContainerService{},
        &AzurermContainerServiceList{},
    
        &AzurermCosmosdbSqlDatabase{},
        &AzurermCosmosdbSqlDatabaseList{},
    
        &AzurermEventhubConsumerGroup{},
        &AzurermEventhubConsumerGroupList{},
    
        &AzurermAutomationVariableBool{},
        &AzurermAutomationVariableBoolList{},
    
        &AzurermAutomationVariableString{},
        &AzurermAutomationVariableStringList{},
    
        &AzurermRecoveryServicesProtectionPolicyVm{},
        &AzurermRecoveryServicesProtectionPolicyVmList{},
    
        &AzurermSharedImageGallery{},
        &AzurermSharedImageGalleryList{},
    
        &AzurermStreamAnalyticsOutputBlob{},
        &AzurermStreamAnalyticsOutputBlobList{},
    
        &AzurermVirtualMachineScaleSet{},
        &AzurermVirtualMachineScaleSetList{},
    
        &AzurermDataLakeStoreFile{},
        &AzurermDataLakeStoreFileList{},
    
        &AzurermDnsTxtRecord{},
        &AzurermDnsTxtRecordList{},
    
        &AzurermPacketCapture{},
        &AzurermPacketCaptureList{},
    
        &AzurermStreamAnalyticsOutputMssql{},
        &AzurermStreamAnalyticsOutputMssqlList{},
    
        &AzurermStreamAnalyticsOutputServicebusQueue{},
        &AzurermStreamAnalyticsOutputServicebusQueueList{},
    
        &AzurermApiManagementProperty{},
        &AzurermApiManagementPropertyList{},
    
        &AzurermDatabricksWorkspace{},
        &AzurermDatabricksWorkspaceList{},
    
        &AzurermLbRule{},
        &AzurermLbRuleList{},
    
        &AzurermIothubSharedAccessPolicy{},
        &AzurermIothubSharedAccessPolicyList{},
    
        &AzurermKubernetesCluster{},
        &AzurermKubernetesClusterList{},
    
        &AzurermLb{},
        &AzurermLbList{},
    
        &AzurermApiManagementApiSchema{},
        &AzurermApiManagementApiSchemaList{},
    
        &AzurermApiManagementProduct{},
        &AzurermApiManagementProductList{},
    
        &AzurermAppService{},
        &AzurermAppServiceList{},
    
        &AzurermCosmosdbCassandraKeyspace{},
        &AzurermCosmosdbCassandraKeyspaceList{},
    
        &AzurermDevspaceController{},
        &AzurermDevspaceControllerList{},
    
        &AzurermSearchService{},
        &AzurermSearchServiceList{},
    
        &AzurermServicebusNamespace{},
        &AzurermServicebusNamespaceList{},
    
        &AzurermAppServiceSlot{},
        &AzurermAppServiceSlotList{},
    
        &AzurermStorageAccount{},
        &AzurermStorageAccountList{},
    
        &AzurermMysqlDatabase{},
        &AzurermMysqlDatabaseList{},
    
        &AzurermContainerRegistry{},
        &AzurermContainerRegistryList{},
    
        &AzurermDataLakeAnalyticsAccount{},
        &AzurermDataLakeAnalyticsAccountList{},
    
        &AzurermServicebusQueueAuthorizationRule{},
        &AzurermServicebusQueueAuthorizationRuleList{},
    
        &AzurermSqlDatabase{},
        &AzurermSqlDatabaseList{},
    
        &AzurermTemplateDeployment{},
        &AzurermTemplateDeploymentList{},
    
        &AzurermDevTestWindowsVirtualMachine{},
        &AzurermDevTestWindowsVirtualMachineList{},
    
        &AzurermEventhubNamespaceAuthorizationRule{},
        &AzurermEventhubNamespaceAuthorizationRuleList{},
    
        &AzurermLocalNetworkGateway{},
        &AzurermLocalNetworkGatewayList{},
    
        &AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
        &AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},
    
        &AzurermSecurityCenterContact{},
        &AzurermSecurityCenterContactList{},
    
        &AzurermAutomationVariableInt{},
        &AzurermAutomationVariableIntList{},
    
        &AzurermDnsCnameRecord{},
        &AzurermDnsCnameRecordList{},
    
        &AzurermMonitorActivityLogAlert{},
        &AzurermMonitorActivityLogAlertList{},
    
        &AzurermSchedulerJob{},
        &AzurermSchedulerJobList{},
    
        &AzurermServicebusNamespaceAuthorizationRule{},
        &AzurermServicebusNamespaceAuthorizationRuleList{},
    
        &AzurermVirtualMachineExtension{},
        &AzurermVirtualMachineExtensionList{},
    
        &AzurermApiManagementSubscription{},
        &AzurermApiManagementSubscriptionList{},
    
        &AzurermAppServicePlan{},
        &AzurermAppServicePlanList{},
    
        &AzurermAzureadServicePrincipalPassword{},
        &AzurermAzureadServicePrincipalPasswordList{},
    
        &AzurermIotDpsCertificate{},
        &AzurermIotDpsCertificateList{},
    
        &AzurermMysqlVirtualNetworkRule{},
        &AzurermMysqlVirtualNetworkRuleList{},
    
        &AzurermApplicationInsightsApiKey{},
        &AzurermApplicationInsightsApiKeyList{},
    
        &AzurermManagedDisk{},
        &AzurermManagedDiskList{},
    
        &AzurermNetworkWatcher{},
        &AzurermNetworkWatcherList{},
    
        &AzurermRoleAssignment{},
        &AzurermRoleAssignmentList{},
    
        &AzurermServicebusSubscription{},
        &AzurermServicebusSubscriptionList{},
    
        &AzurermDataFactoryDatasetMysql{},
        &AzurermDataFactoryDatasetMysqlList{},
    
        &AzurermMariadbFirewallRule{},
        &AzurermMariadbFirewallRuleList{},
    
        &AzurermPolicySetDefinition{},
        &AzurermPolicySetDefinitionList{},
    
        &AzurermApiManagementApiPolicy{},
        &AzurermApiManagementApiPolicyList{},
    
        &AzurermHdinsightHbaseCluster{},
        &AzurermHdinsightHbaseClusterList{},
    
        &AzurermNetworkInterfaceNatRuleAssociation{},
        &AzurermNetworkInterfaceNatRuleAssociationList{},
    
        &AzurermKeyVaultAccessPolicy{},
        &AzurermKeyVaultAccessPolicyList{},
    
        &AzurermApiManagementApi{},
        &AzurermApiManagementApiList{},
    
        &AzurermAutomationDscNodeconfiguration{},
        &AzurermAutomationDscNodeconfigurationList{},
    
        &AzurermDataFactoryDatasetSqlServerTable{},
        &AzurermDataFactoryDatasetSqlServerTableList{},
    
        &AzurermDataFactoryLinkedServiceSqlServer{},
        &AzurermDataFactoryLinkedServiceSqlServerList{},
    
        &AzurermDevTestPolicy{},
        &AzurermDevTestPolicyList{},
    
        &AzurermDataFactoryLinkedServiceDataLakeStorageGen2{},
        &AzurermDataFactoryLinkedServiceDataLakeStorageGen2List{},
    
        &AzurermEventhubAuthorizationRule{},
        &AzurermEventhubAuthorizationRuleList{},
    
        &AzurermPostgresqlConfiguration{},
        &AzurermPostgresqlConfigurationList{},
    
        &AzurermCdnEndpoint{},
        &AzurermCdnEndpointList{},
    
        &AzurermDataFactoryDatasetPostgresql{},
        &AzurermDataFactoryDatasetPostgresqlList{},
    
        &AzurermFirewall{},
        &AzurermFirewallList{},
    
        &AzurermHdinsightMlServicesCluster{},
        &AzurermHdinsightMlServicesClusterList{},
    
        &AzurermApiManagementGroup{},
        &AzurermApiManagementGroupList{},
    
        &AzurermAutomationRunbook{},
        &AzurermAutomationRunbookList{},
    
        &AzurermNetworkProfile{},
        &AzurermNetworkProfileList{},
    
        &AzurermRedisCache{},
        &AzurermRedisCacheList{},
    
        &AzurermAutoscaleSetting{},
        &AzurermAutoscaleSettingList{},
    
        &AzurermDataFactory{},
        &AzurermDataFactoryList{},
    
        &AzurermNetworkSecurityRule{},
        &AzurermNetworkSecurityRuleList{},
    
        &AzurermDataFactoryLinkedServiceMysql{},
        &AzurermDataFactoryLinkedServiceMysqlList{},
    
        &AzurermNotificationHubNamespace{},
        &AzurermNotificationHubNamespaceList{},
    
        &AzurermPublicIp{},
        &AzurermPublicIpList{},
    
        &AzurermResourceGroup{},
        &AzurermResourceGroupList{},
    
        &AzurermServicebusQueue{},
        &AzurermServicebusQueueList{},
    
        &AzurermApiManagementApiOperation{},
        &AzurermApiManagementApiOperationList{},
    
        &AzurermAzureadServicePrincipal{},
        &AzurermAzureadServicePrincipalList{},
    
        &AzurermLogAnalyticsWorkspaceLinkedService{},
        &AzurermLogAnalyticsWorkspaceLinkedServiceList{},
    
        &AzurermMonitorActionGroup{},
        &AzurermMonitorActionGroupList{},
    
        &AzurermStorageBlob{},
        &AzurermStorageBlobList{},
    
        &AzurermFunctionApp{},
        &AzurermFunctionAppList{},
    
        &AzurermIothub{},
        &AzurermIothubList{},
    
        &AzurermDnsZone{},
        &AzurermDnsZoneList{},
    
        &AzurermLbNatPool{},
        &AzurermLbNatPoolList{},
    
        &AzurermSubnetRouteTableAssociation{},
        &AzurermSubnetRouteTableAssociationList{},
    
        &AzurermSecurityCenterSubscriptionPricing{},
        &AzurermSecurityCenterSubscriptionPricingList{},
    
        &AzurermSqlVirtualNetworkRule{},
        &AzurermSqlVirtualNetworkRuleList{},
    
        &AzurermStorageContainer{},
        &AzurermStorageContainerList{},
    
        &AzurermApiManagementProductApi{},
        &AzurermApiManagementProductApiList{},
    
        &AzurermDataLakeStoreFirewallRule{},
        &AzurermDataLakeStoreFirewallRuleList{},
    
        &AzurermDevTestLinuxVirtualMachine{},
        &AzurermDevTestLinuxVirtualMachineList{},
    
        &AzurermExpressRouteCircuitAuthorization{},
        &AzurermExpressRouteCircuitAuthorizationList{},
    
        &AzurermLbProbe{},
        &AzurermLbProbeList{},
    
        &AzurermVirtualMachine{},
        &AzurermVirtualMachineList{},
    
        &AzurermDataLakeAnalyticsFirewallRule{},
        &AzurermDataLakeAnalyticsFirewallRuleList{},
    
        &AzurermManagementGroup{},
        &AzurermManagementGroupList{},
    
        &AzurermManagementLock{},
        &AzurermManagementLockList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
