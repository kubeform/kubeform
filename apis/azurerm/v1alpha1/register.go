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

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&Image{},
		&ImageList{},

		&DnsARecord{},
		&DnsARecordList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&Subnet{},
		&SubnetList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&LbProbe{},
		&LbProbeList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&DevTestLab{},
		&DevTestLabList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&ContainerService{},
		&ContainerServiceList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PublicIP{},
		&PublicIPList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&ManagementLock{},
		&ManagementLockList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&Iothub{},
		&IothubList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&StorageContainer{},
		&StorageContainerList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&SqlServer{},
		&SqlServerList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&Firewall{},
		&FirewallList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&SearchService{},
		&SearchServiceList{},

		&StorageQueue{},
		&StorageQueueList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&CdnProfile{},
		&CdnProfileList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&Eventhub{},
		&EventhubList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&DataFactory{},
		&DataFactoryList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&StorageTable{},
		&StorageTableList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&LbRule{},
		&LbRuleList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&Snapshot{},
		&SnapshotList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MysqlServer{},
		&MysqlServerList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DnsZone{},
		&DnsZoneList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&BatchPool{},
		&BatchPoolList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&FunctionApp{},
		&FunctionAppList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&BatchAccount{},
		&BatchAccountList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&IotDps{},
		&IotDpsList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&KeyVault{},
		&KeyVaultList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&RouteTable{},
		&RouteTableList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&AppService{},
		&AppServiceList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RedisCache{},
		&RedisCacheList{},

		&StorageShare{},
		&StorageShareList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&Route{},
		&RouteList{},

		&StorageAccount{},
		&StorageAccountList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&Lb{},
		&LbList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
