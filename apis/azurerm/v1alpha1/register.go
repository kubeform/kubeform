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
    
        &AzurermApiManagementGroup{},
        &AzurermApiManagementGroupList{},
    
        &AzurermAppService{},
        &AzurermAppServiceList{},
    
        &AzurermEventgridEventSubscription{},
        &AzurermEventgridEventSubscriptionList{},
    
        &AzurermKubernetesCluster{},
        &AzurermKubernetesClusterList{},
    
        &AzurermNetworkProfile{},
        &AzurermNetworkProfileList{},
    
        &AzurermPostgresqlDatabase{},
        &AzurermPostgresqlDatabaseList{},
    
        &AzurermRecoveryServicesProtectionPolicyVm{},
        &AzurermRecoveryServicesProtectionPolicyVmList{},
    
        &AzurermDevTestPolicy{},
        &AzurermDevTestPolicyList{},
    
        &AzurermIotDpsCertificate{},
        &AzurermIotDpsCertificateList{},
    
        &AzurermVirtualNetworkGateway{},
        &AzurermVirtualNetworkGatewayList{},
    
        &AzurermConnectionMonitor{},
        &AzurermConnectionMonitorList{},
    
        &AzurermMysqlVirtualNetworkRule{},
        &AzurermMysqlVirtualNetworkRuleList{},
    
        &AzurermSecurityCenterSubscriptionPricing{},
        &AzurermSecurityCenterSubscriptionPricingList{},
    
        &AzurermApiManagementApiPolicy{},
        &AzurermApiManagementApiPolicyList{},
    
        &AzurermApplicationInsightsWebTest{},
        &AzurermApplicationInsightsWebTestList{},
    
        &AzurermAutoscaleSetting{},
        &AzurermAutoscaleSettingList{},
    
        &AzurermDnsPtrRecord{},
        &AzurermDnsPtrRecordList{},
    
        &AzurermVirtualMachineExtension{},
        &AzurermVirtualMachineExtensionList{},
    
        &AzurermHdinsightMlServicesCluster{},
        &AzurermHdinsightMlServicesClusterList{},
    
        &AzurermStreamAnalyticsOutputBlob{},
        &AzurermStreamAnalyticsOutputBlobList{},
    
        &AzurermApiManagementLogger{},
        &AzurermApiManagementLoggerList{},
    
        &AzurermAvailabilitySet{},
        &AzurermAvailabilitySetList{},
    
        &AzurermManagedDisk{},
        &AzurermManagedDiskList{},
    
        &AzurermManagementLock{},
        &AzurermManagementLockList{},
    
        &AzurermNetworkConnectionMonitor{},
        &AzurermNetworkConnectionMonitorList{},
    
        &AzurermAutomationModule{},
        &AzurermAutomationModuleList{},
    
        &AzurermDataFactoryLinkedServicePostgresql{},
        &AzurermDataFactoryLinkedServicePostgresqlList{},
    
        &AzurermAutomationAccount{},
        &AzurermAutomationAccountList{},
    
        &AzurermAzureadServicePrincipalPassword{},
        &AzurermAzureadServicePrincipalPasswordList{},
    
        &AzurermCosmosdbSqlDatabase{},
        &AzurermCosmosdbSqlDatabaseList{},
    
        &AzurermKeyVaultSecret{},
        &AzurermKeyVaultSecretList{},
    
        &AzurermSqlServer{},
        &AzurermSqlServerList{},
    
        &AzurermStorageAccount{},
        &AzurermStorageAccountList{},
    
        &AzurermAutomationVariableString{},
        &AzurermAutomationVariableStringList{},
    
        &AzurermEventgridDomain{},
        &AzurermEventgridDomainList{},
    
        &AzurermVirtualMachineScaleSet{},
        &AzurermVirtualMachineScaleSetList{},
    
        &AzurermDnsZone{},
        &AzurermDnsZoneList{},
    
        &AzurermPostgresqlConfiguration{},
        &AzurermPostgresqlConfigurationList{},
    
        &AzurermPublicIpPrefix{},
        &AzurermPublicIpPrefixList{},
    
        &AzurermRelayNamespace{},
        &AzurermRelayNamespaceList{},
    
        &AzurermApiManagementGroupUser{},
        &AzurermApiManagementGroupUserList{},
    
        &AzurermServiceFabricCluster{},
        &AzurermServiceFabricClusterList{},
    
        &AzurermServicebusSubscriptionRule{},
        &AzurermServicebusSubscriptionRuleList{},
    
        &AzurermContainerRegistry{},
        &AzurermContainerRegistryList{},
    
        &AzurermAutomationVariableInt{},
        &AzurermAutomationVariableIntList{},
    
        &AzurermImage{},
        &AzurermImageList{},
    
        &AzurermIothubSharedAccessPolicy{},
        &AzurermIothubSharedAccessPolicyList{},
    
        &AzurermNetworkInterface{},
        &AzurermNetworkInterfaceList{},
    
        &AzurermPublicIp{},
        &AzurermPublicIpList{},
    
        &AzurermSecurityCenterWorkspace{},
        &AzurermSecurityCenterWorkspaceList{},
    
        &AzurermAzureadApplication{},
        &AzurermAzureadApplicationList{},
    
        &AzurermDevTestLab{},
        &AzurermDevTestLabList{},
    
        &AzurermVirtualNetworkGatewayConnection{},
        &AzurermVirtualNetworkGatewayConnectionList{},
    
        &AzurermDataFactoryDatasetSqlServerTable{},
        &AzurermDataFactoryDatasetSqlServerTableList{},
    
        &AzurermDnsTxtRecord{},
        &AzurermDnsTxtRecordList{},
    
        &AzurermMonitorActionGroup{},
        &AzurermMonitorActionGroupList{},
    
        &AzurermPostgresqlServer{},
        &AzurermPostgresqlServerList{},
    
        &AzurermServicebusTopic{},
        &AzurermServicebusTopicList{},
    
        &AzurermSqlDatabase{},
        &AzurermSqlDatabaseList{},
    
        &AzurermTemplateDeployment{},
        &AzurermTemplateDeploymentList{},
    
        &AzurermKeyVaultAccessPolicy{},
        &AzurermKeyVaultAccessPolicyList{},
    
        &AzurermMonitorAutoscaleSetting{},
        &AzurermMonitorAutoscaleSettingList{},
    
        &AzurermPacketCapture{},
        &AzurermPacketCaptureList{},
    
        &AzurermApiManagementOpenidConnectProvider{},
        &AzurermApiManagementOpenidConnectProviderList{},
    
        &AzurermDataFactoryDatasetMysql{},
        &AzurermDataFactoryDatasetMysqlList{},
    
        &AzurermSecurityCenterContact{},
        &AzurermSecurityCenterContactList{},
    
        &AzurermSqlFirewallRule{},
        &AzurermSqlFirewallRuleList{},
    
        &AzurermMariadbDatabase{},
        &AzurermMariadbDatabaseList{},
    
        &AzurermRedisCache{},
        &AzurermRedisCacheList{},
    
        &AzurermBatchAccount{},
        &AzurermBatchAccountList{},
    
        &AzurermVirtualMachineDataDiskAttachment{},
        &AzurermVirtualMachineDataDiskAttachmentList{},
    
        &AzurermMysqlFirewallRule{},
        &AzurermMysqlFirewallRuleList{},
    
        &AzurermRoute{},
        &AzurermRouteList{},
    
        &AzurermSnapshot{},
        &AzurermSnapshotList{},
    
        &AzurermEventgridTopic{},
        &AzurermEventgridTopicList{},
    
        &AzurermMonitorMetricAlertrule{},
        &AzurermMonitorMetricAlertruleList{},
    
        &AzurermNetworkSecurityGroup{},
        &AzurermNetworkSecurityGroupList{},
    
        &AzurermServicebusNamespace{},
        &AzurermServicebusNamespaceList{},
    
        &AzurermApiManagementProductGroup{},
        &AzurermApiManagementProductGroupList{},
    
        &AzurermDevTestVirtualNetwork{},
        &AzurermDevTestVirtualNetworkList{},
    
        &AzurermDnsAaaaRecord{},
        &AzurermDnsAaaaRecordList{},
    
        &AzurermLogAnalyticsLinkedService{},
        &AzurermLogAnalyticsLinkedServiceList{},
    
        &AzurermSqlActiveDirectoryAdministrator{},
        &AzurermSqlActiveDirectoryAdministratorList{},
    
        &AzurermStreamAnalyticsStreamInputEventhub{},
        &AzurermStreamAnalyticsStreamInputEventhubList{},
    
        &AzurermApiManagementProductPolicy{},
        &AzurermApiManagementProductPolicyList{},
    
        &AzurermAppServiceActiveSlot{},
        &AzurermAppServiceActiveSlotList{},
    
        &AzurermCosmosdbCassandraKeyspace{},
        &AzurermCosmosdbCassandraKeyspaceList{},
    
        &AzurermDataLakeStoreFile{},
        &AzurermDataLakeStoreFileList{},
    
        &AzurermIotDps{},
        &AzurermIotDpsList{},
    
        &AzurermLogicAppActionHttp{},
        &AzurermLogicAppActionHttpList{},
    
        &AzurermMonitorDiagnosticSetting{},
        &AzurermMonitorDiagnosticSettingList{},
    
        &AzurermCosmosdbAccount{},
        &AzurermCosmosdbAccountList{},
    
        &AzurermDataFactoryLinkedServiceSqlServer{},
        &AzurermDataFactoryLinkedServiceSqlServerList{},
    
        &AzurermHdinsightInteractiveQueryCluster{},
        &AzurermHdinsightInteractiveQueryClusterList{},
    
        &AzurermLogicAppTriggerRecurrence{},
        &AzurermLogicAppTriggerRecurrenceList{},
    
        &AzurermCdnEndpoint{},
        &AzurermCdnEndpointList{},
    
        &AzurermContainerService{},
        &AzurermContainerServiceList{},
    
        &AzurermDdosProtectionPlan{},
        &AzurermDdosProtectionPlanList{},
    
        &AzurermLogAnalyticsWorkspace{},
        &AzurermLogAnalyticsWorkspaceList{},
    
        &AzurermMysqlDatabase{},
        &AzurermMysqlDatabaseList{},
    
        &AzurermSubnetNetworkSecurityGroupAssociation{},
        &AzurermSubnetNetworkSecurityGroupAssociationList{},
    
        &AzurermExpressRouteCircuit{},
        &AzurermExpressRouteCircuitList{},
    
        &AzurermHdinsightStormCluster{},
        &AzurermHdinsightStormClusterList{},
    
        &AzurermKeyVaultCertificate{},
        &AzurermKeyVaultCertificateList{},
    
        &AzurermMetricAlertrule{},
        &AzurermMetricAlertruleList{},
    
        &AzurermStreamAnalyticsOutputServicebusQueue{},
        &AzurermStreamAnalyticsOutputServicebusQueueList{},
    
        &AzurermApiManagementApi{},
        &AzurermApiManagementApiList{},
    
        &AzurermApiManagementUser{},
        &AzurermApiManagementUserList{},
    
        &AzurermCosmosdbTable{},
        &AzurermCosmosdbTableList{},
    
        &AzurermServicebusSubscription{},
        &AzurermServicebusSubscriptionList{},
    
        &AzurermDataLakeAnalyticsAccount{},
        &AzurermDataLakeAnalyticsAccountList{},
    
        &AzurermLbNatRule{},
        &AzurermLbNatRuleList{},
    
        &AzurermMonitorLogProfile{},
        &AzurermMonitorLogProfileList{},
    
        &AzurermPolicyAssignment{},
        &AzurermPolicyAssignmentList{},
    
        &AzurermSqlVirtualNetworkRule{},
        &AzurermSqlVirtualNetworkRuleList{},
    
        &AzurermMonitorActivityLogAlert{},
        &AzurermMonitorActivityLogAlertList{},
    
        &AzurermServicebusQueueAuthorizationRule{},
        &AzurermServicebusQueueAuthorizationRuleList{},
    
        &AzurermHdinsightHadoopCluster{},
        &AzurermHdinsightHadoopClusterList{},
    
        &AzurermLbOutboundRule{},
        &AzurermLbOutboundRuleList{},
    
        &AzurermMariadbFirewallRule{},
        &AzurermMariadbFirewallRuleList{},
    
        &AzurermLogicAppWorkflow{},
        &AzurermLogicAppWorkflowList{},
    
        &AzurermSchedulerJobCollection{},
        &AzurermSchedulerJobCollectionList{},
    
        &AzurermStreamAnalyticsFunctionJavascriptUdf{},
        &AzurermStreamAnalyticsFunctionJavascriptUdfList{},
    
        &AzurermDevTestLinuxVirtualMachine{},
        &AzurermDevTestLinuxVirtualMachineList{},
    
        &AzurermHdinsightHbaseCluster{},
        &AzurermHdinsightHbaseClusterList{},
    
        &AzurermStorageContainer{},
        &AzurermStorageContainerList{},
    
        &AzurermTrafficManagerProfile{},
        &AzurermTrafficManagerProfileList{},
    
        &AzurermApiManagementProduct{},
        &AzurermApiManagementProductList{},
    
        &AzurermApiManagementProductApi{},
        &AzurermApiManagementProductApiList{},
    
        &AzurermEventhubConsumerGroup{},
        &AzurermEventhubConsumerGroupList{},
    
        &AzurermManagementGroup{},
        &AzurermManagementGroupList{},
    
        &AzurermNetworkWatcher{},
        &AzurermNetworkWatcherList{},
    
        &AzurermRoleAssignment{},
        &AzurermRoleAssignmentList{},
    
        &AzurermSqlElasticpool{},
        &AzurermSqlElasticpoolList{},
    
        &AzurermDnsCnameRecord{},
        &AzurermDnsCnameRecordList{},
    
        &AzurermServicebusNamespaceAuthorizationRule{},
        &AzurermServicebusNamespaceAuthorizationRuleList{},
    
        &AzurermAutomationSchedule{},
        &AzurermAutomationScheduleList{},
    
        &AzurermCosmosdbMongoCollection{},
        &AzurermCosmosdbMongoCollectionList{},
    
        &AzurermMysqlServer{},
        &AzurermMysqlServerList{},
    
        &AzurermStreamAnalyticsStreamInputIothub{},
        &AzurermStreamAnalyticsStreamInputIothubList{},
    
        &AzurermUserAssignedIdentity{},
        &AzurermUserAssignedIdentityList{},
    
        &AzurermVirtualNetwork{},
        &AzurermVirtualNetworkList{},
    
        &AzurermDataFactoryDatasetPostgresql{},
        &AzurermDataFactoryDatasetPostgresqlList{},
    
        &AzurermFirewallApplicationRuleCollection{},
        &AzurermFirewallApplicationRuleCollectionList{},
    
        &AzurermFirewallNetworkRuleCollection{},
        &AzurermFirewallNetworkRuleCollectionList{},
    
        &AzurermPostgresqlVirtualNetworkRule{},
        &AzurermPostgresqlVirtualNetworkRuleList{},
    
        &AzurermDevTestWindowsVirtualMachine{},
        &AzurermDevTestWindowsVirtualMachineList{},
    
        &AzurermEventhubAuthorizationRule{},
        &AzurermEventhubAuthorizationRuleList{},
    
        &AzurermSharedImageGallery{},
        &AzurermSharedImageGalleryList{},
    
        &AzurermStorageShare{},
        &AzurermStorageShareList{},
    
        &AzurermAutomationRunbook{},
        &AzurermAutomationRunbookList{},
    
        &AzurermDnsCaaRecord{},
        &AzurermDnsCaaRecordList{},
    
        &AzurermMonitorMetricAlert{},
        &AzurermMonitorMetricAlertList{},
    
        &AzurermStreamAnalyticsJob{},
        &AzurermStreamAnalyticsJobList{},
    
        &AzurermCdnProfile{},
        &AzurermCdnProfileList{},
    
        &AzurermApiManagementAuthorizationServer{},
        &AzurermApiManagementAuthorizationServerList{},
    
        &AzurermApiManagementCertificate{},
        &AzurermApiManagementCertificateList{},
    
        &AzurermApiManagementProperty{},
        &AzurermApiManagementPropertyList{},
    
        &AzurermAutomationVariableBool{},
        &AzurermAutomationVariableBoolList{},
    
        &AzurermFunctionApp{},
        &AzurermFunctionAppList{},
    
        &AzurermNetworkInterfaceApplicationSecurityGroupAssociation{},
        &AzurermNetworkInterfaceApplicationSecurityGroupAssociationList{},
    
        &AzurermAzureadServicePrincipal{},
        &AzurermAzureadServicePrincipalList{},
    
        &AzurermDnsMxRecord{},
        &AzurermDnsMxRecordList{},
    
        &AzurermLbNatPool{},
        &AzurermLbNatPoolList{},
    
        &AzurermSharedImage{},
        &AzurermSharedImageList{},
    
        &AzurermVirtualMachine{},
        &AzurermVirtualMachineList{},
    
        &AzurermAutomationDscConfiguration{},
        &AzurermAutomationDscConfigurationList{},
    
        &AzurermBatchPool{},
        &AzurermBatchPoolList{},
    
        &AzurermDataFactory{},
        &AzurermDataFactoryList{},
    
        &AzurermDataLakeStore{},
        &AzurermDataLakeStoreList{},
    
        &AzurermIothub{},
        &AzurermIothubList{},
    
        &AzurermLbProbe{},
        &AzurermLbProbeList{},
    
        &AzurermNetworkSecurityRule{},
        &AzurermNetworkSecurityRuleList{},
    
        &AzurermAppServiceSlot{},
        &AzurermAppServiceSlotList{},
    
        &AzurermMariadbServer{},
        &AzurermMariadbServerList{},
    
        &AzurermNotificationHubAuthorizationRule{},
        &AzurermNotificationHubAuthorizationRuleList{},
    
        &AzurermRouteTable{},
        &AzurermRouteTableList{},
    
        &AzurermStorageQueue{},
        &AzurermStorageQueueList{},
    
        &AzurermHdinsightSparkCluster{},
        &AzurermHdinsightSparkClusterList{},
    
        &AzurermDnsARecord{},
        &AzurermDnsARecordList{},
    
        &AzurermEventhubNamespace{},
        &AzurermEventhubNamespaceList{},
    
        &AzurermRedisFirewallRule{},
        &AzurermRedisFirewallRuleList{},
    
        &AzurermSearchService{},
        &AzurermSearchServiceList{},
    
        &AzurermApplicationInsights{},
        &AzurermApplicationInsightsList{},
    
        &AzurermDataFactoryLinkedServiceDataLakeStorageGen2{},
        &AzurermDataFactoryLinkedServiceDataLakeStorageGen2List{},
    
        &AzurermLocalNetworkGateway{},
        &AzurermLocalNetworkGatewayList{},
    
        &AzurermLogicAppTriggerHttpRequest{},
        &AzurermLogicAppTriggerHttpRequestList{},
    
        &AzurermRecoveryServicesVault{},
        &AzurermRecoveryServicesVaultList{},
    
        &AzurermExpressRouteCircuitAuthorization{},
        &AzurermExpressRouteCircuitAuthorizationList{},
    
        &AzurermHdinsightKafkaCluster{},
        &AzurermHdinsightKafkaClusterList{},
    
        &AzurermKeyVaultKey{},
        &AzurermKeyVaultKeyList{},
    
        &AzurermNetworkInterfaceNatRuleAssociation{},
        &AzurermNetworkInterfaceNatRuleAssociationList{},
    
        &AzurermPolicyDefinition{},
        &AzurermPolicyDefinitionList{},
    
        &AzurermStorageBlob{},
        &AzurermStorageBlobList{},
    
        &AzurermCognitiveAccount{},
        &AzurermCognitiveAccountList{},
    
        &AzurermMediaServicesAccount{},
        &AzurermMediaServicesAccountList{},
    
        &AzurermStreamAnalyticsOutputMssql{},
        &AzurermStreamAnalyticsOutputMssqlList{},
    
        &AzurermVirtualNetworkPeering{},
        &AzurermVirtualNetworkPeeringList{},
    
        &AzurermAppServicePlan{},
        &AzurermAppServicePlanList{},
    
        &AzurermDnsSrvRecord{},
        &AzurermDnsSrvRecordList{},
    
        &AzurermFirewall{},
        &AzurermFirewallList{},
    
        &AzurermIothubConsumerGroup{},
        &AzurermIothubConsumerGroupList{},
    
        &AzurermNotificationHubNamespace{},
        &AzurermNotificationHubNamespaceList{},
    
        &AzurermNotificationHub{},
        &AzurermNotificationHubList{},
    
        &AzurermDataLakeStoreFirewallRule{},
        &AzurermDataLakeStoreFirewallRuleList{},
    
        &AzurermLogicAppActionCustom{},
        &AzurermLogicAppActionCustomList{},
    
        &AzurermSharedImageVersion{},
        &AzurermSharedImageVersionList{},
    
        &AzurermSignalrService{},
        &AzurermSignalrServiceList{},
    
        &AzurermApiManagementApiSchema{},
        &AzurermApiManagementApiSchemaList{},
    
        &AzurermDnsNsRecord{},
        &AzurermDnsNsRecordList{},
    
        &AzurermExpressRouteCircuitPeering{},
        &AzurermExpressRouteCircuitPeeringList{},
    
        &AzurermLb{},
        &AzurermLbList{},
    
        &AzurermNetworkDdosProtectionPlan{},
        &AzurermNetworkDdosProtectionPlanList{},
    
        &AzurermStreamAnalyticsOutputEventhub{},
        &AzurermStreamAnalyticsOutputEventhubList{},
    
        &AzurermApiManagementApiOperation{},
        &AzurermApiManagementApiOperationList{},
    
        &AzurermApiManagementApiVersionSet{},
        &AzurermApiManagementApiVersionSetList{},
    
        &AzurermBatchCertificate{},
        &AzurermBatchCertificateList{},
    
        &AzurermEventhubNamespaceAuthorizationRule{},
        &AzurermEventhubNamespaceAuthorizationRuleList{},
    
        &AzurermLbRule{},
        &AzurermLbRuleList{},
    
        &AzurermPostgresqlFirewallRule{},
        &AzurermPostgresqlFirewallRuleList{},
    
        &AzurermPrivateDnsZone{},
        &AzurermPrivateDnsZoneList{},
    
        &AzurermSubnetRouteTableAssociation{},
        &AzurermSubnetRouteTableAssociationList{},
    
        &AzurermTrafficManagerEndpoint{},
        &AzurermTrafficManagerEndpointList{},
    
        &AzurermApplicationGateway{},
        &AzurermApplicationGatewayList{},
    
        &AzurermApplicationInsightsApiKey{},
        &AzurermApplicationInsightsApiKeyList{},
    
        &AzurermContainerGroup{},
        &AzurermContainerGroupList{},
    
        &AzurermDataFactoryLinkedServiceMysql{},
        &AzurermDataFactoryLinkedServiceMysqlList{},
    
        &AzurermDataFactoryPipeline{},
        &AzurermDataFactoryPipelineList{},
    
        &AzurermApplicationSecurityGroup{},
        &AzurermApplicationSecurityGroupList{},
    
        &AzurermLogAnalyticsSolution{},
        &AzurermLogAnalyticsSolutionList{},
    
        &AzurermLogicAppTriggerCustom{},
        &AzurermLogicAppTriggerCustomList{},
    
        &AzurermAutomationCredential{},
        &AzurermAutomationCredentialList{},
    
        &AzurermPolicySetDefinition{},
        &AzurermPolicySetDefinitionList{},
    
        &AzurermResourceGroup{},
        &AzurermResourceGroupList{},
    
        &AzurermServicebusQueue{},
        &AzurermServicebusQueueList{},
    
        &AzurermNetworkPacketCapture{},
        &AzurermNetworkPacketCaptureList{},
    
        &AzurermAutomationDscNodeconfiguration{},
        &AzurermAutomationDscNodeconfigurationList{},
    
        &AzurermDevspaceController{},
        &AzurermDevspaceControllerList{},
    
        &AzurermKeyVault{},
        &AzurermKeyVaultList{},
    
        &AzurermLbBackendAddressPool{},
        &AzurermLbBackendAddressPoolList{},
    
        &AzurermSubnet{},
        &AzurermSubnetList{},
    
        &AzurermStorageTable{},
        &AzurermStorageTableList{},
    
        &AzurermApiManagementApiOperationPolicy{},
        &AzurermApiManagementApiOperationPolicyList{},
    
        &AzurermDatabricksWorkspace{},
        &AzurermDatabricksWorkspaceList{},
    
        &AzurermEventhub{},
        &AzurermEventhubList{},
    
        &AzurermHdinsightRserverCluster{},
        &AzurermHdinsightRserverClusterList{},
    
        &AzurermMssqlElasticpool{},
        &AzurermMssqlElasticpoolList{},
    
        &AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
        &AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},
    
        &AzurermSchedulerJob{},
        &AzurermSchedulerJobList{},
    
        &AzurermAppServiceCustomHostnameBinding{},
        &AzurermAppServiceCustomHostnameBindingList{},
    
        &AzurermAutomationVariableDatetime{},
        &AzurermAutomationVariableDatetimeList{},
    
        &AzurermDataLakeAnalyticsFirewallRule{},
        &AzurermDataLakeAnalyticsFirewallRuleList{},
    
        &AzurermFirewallNatRuleCollection{},
        &AzurermFirewallNatRuleCollectionList{},
    
        &AzurermLogAnalyticsWorkspaceLinkedService{},
        &AzurermLogAnalyticsWorkspaceLinkedServiceList{},
    
        &AzurermApiManagement{},
        &AzurermApiManagementList{},
    
        &AzurermCosmosdbMongoDatabase{},
        &AzurermCosmosdbMongoDatabaseList{},
    
        &AzurermRecoveryServicesProtectedVm{},
        &AzurermRecoveryServicesProtectedVmList{},
    
        &AzurermRoleDefinition{},
        &AzurermRoleDefinitionList{},
    
        &AzurermApiManagementSubscription{},
        &AzurermApiManagementSubscriptionList{},
    
        &AzurermMysqlConfiguration{},
        &AzurermMysqlConfigurationList{},
    
        &AzurermNetworkInterfaceBackendAddressPoolAssociation{},
        &AzurermNetworkInterfaceBackendAddressPoolAssociationList{},
    
        &AzurermServicebusTopicAuthorizationRule{},
        &AzurermServicebusTopicAuthorizationRuleList{},
    
        &AzurermStreamAnalyticsStreamInputBlob{},
        &AzurermStreamAnalyticsStreamInputBlobList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
