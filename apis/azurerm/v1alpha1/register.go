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

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RedisCache{},
		&RedisCacheList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&Eventhub{},
		&EventhubList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&FunctionApp{},
		&FunctionAppList{},

		&KeyVault{},
		&KeyVaultList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&StorageQueue{},
		&StorageQueueList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&ContainerService{},
		&ContainerServiceList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&BatchPool{},
		&BatchPoolList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&StorageShare{},
		&StorageShareList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&SearchService{},
		&SearchServiceList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&Image{},
		&ImageList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&MysqlServer{},
		&MysqlServerList{},

		&StorageAccount{},
		&StorageAccountList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&Route{},
		&RouteList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&StorageContainer{},
		&StorageContainerList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&Firewall{},
		&FirewallList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&IotDps{},
		&IotDpsList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&ManagementLock{},
		&ManagementLockList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&Iothub{},
		&IothubList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsZone{},
		&DnsZoneList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&SqlServer{},
		&SqlServerList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&SharedImage{},
		&SharedImageList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&Snapshot{},
		&SnapshotList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&StorageTable{},
		&StorageTableList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&MariadbServer{},
		&MariadbServerList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&Subnet{},
		&SubnetList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&StorageBlob{},
		&StorageBlobList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&PublicIP{},
		&PublicIPList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&LbProbe{},
		&LbProbeList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DnsARecord{},
		&DnsARecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppService{},
		&AppServiceList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&Lb{},
		&LbList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&LbRule{},
		&LbRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
