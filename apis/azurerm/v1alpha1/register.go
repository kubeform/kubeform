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

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogicAppTriggerHttpRequest{},
		&LogicAppTriggerHttpRequestList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&CosmosdbSqlDatabase{},
		&CosmosdbSqlDatabaseList{},

		&IotDps{},
		&IotDpsList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&LbNatPool{},
		&LbNatPoolList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&StorageContainer{},
		&StorageContainerList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DevTestLab{},
		&DevTestLabList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&PublicIp{},
		&PublicIpList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&SqlServer{},
		&SqlServerList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApiManagementApiVersionSet{},
		&ApiManagementApiVersionSetList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&KeyVault{},
		&KeyVaultList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&LbNatRule{},
		&LbNatRuleList{},

		&PublicIpPrefix{},
		&PublicIpPrefixList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LbProbe{},
		&LbProbeList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ApiManagementApiOperationPolicy{},
		&ApiManagementApiOperationPolicyList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&Snapshot{},
		&SnapshotList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&Image{},
		&ImageList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementProductApi{},
		&ApiManagementProductApiList{},

		&FirewallNatRuleCollection{},
		&FirewallNatRuleCollectionList{},

		&FunctionApp{},
		&FunctionAppList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataFactoryLinkedServiceSqlServer{},
		&DataFactoryLinkedServiceSqlServerList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&Route{},
		&RouteList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&MysqlServer{},
		&MysqlServerList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactoryDatasetSqlServerTable{},
		&DataFactoryDatasetSqlServerTableList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DnsARecord{},
		&DnsARecordList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&MariadbServer{},
		&MariadbServerList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsZone{},
		&DnsZoneList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&SearchService{},
		&SearchServiceList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppActionHttp{},
		&LogicAppActionHttpList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StorageQueue{},
		&StorageQueueList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApiManagementApiSchema{},
		&ApiManagementApiSchemaList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&Firewall{},
		&FirewallList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&PrivateDnsZone{},
		&PrivateDnsZoneList{},

		&RouteTable{},
		&RouteTableList{},

		&StorageTable{},
		&StorageTableList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&ApiManagementApiPolicy{},
		&ApiManagementApiPolicyList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AppService{},
		&AppServiceList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DataFactory{},
		&DataFactoryList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementApi{},
		&ApiManagementApiList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&SignalrService{},
		&SignalrServiceList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SharedImage{},
		&SharedImageList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&NetworkInterfaceNatRuleAssociation{},
		&NetworkInterfaceNatRuleAssociationList{},

		&RedisCache{},
		&RedisCacheList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApiManagementApiOperation{},
		&ApiManagementApiOperationList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&Iothub{},
		&IothubList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&Subnet{},
		&SubnetList{},

		&ApplicationInsightsApiKey{},
		&ApplicationInsightsApiKeyList{},

		&BatchPool{},
		&BatchPoolList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ContainerService{},
		&ContainerServiceList{},

		&Eventhub{},
		&EventhubList{},

		&Lb{},
		&LbList{},

		&StorageShare{},
		&StorageShareList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
