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

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&Eventhub{},
		&EventhubList{},

		&NotificationHub{},
		&NotificationHubList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DnsZone{},
		&DnsZoneList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&CdnProfile{},
		&CdnProfileList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&Route{},
		&RouteList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AppService{},
		&AppServiceList{},

		&BatchAccount{},
		&BatchAccountList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LbRule{},
		&LbRuleList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&RedisCache{},
		&RedisCacheList{},

		&StorageShare{},
		&StorageShareList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&IotDps{},
		&IotDpsList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&BatchPool{},
		&BatchPoolList{},

		&StorageBlob{},
		&StorageBlobList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ContainerService{},
		&ContainerServiceList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MysqlServer{},
		&MysqlServerList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&Iothub{},
		&IothubList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&Firewall{},
		&FirewallList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&RouteTable{},
		&RouteTableList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StorageTable{},
		&StorageTableList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&Subnet{},
		&SubnetList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&SignalrService{},
		&SignalrServiceList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&KeyVault{},
		&KeyVaultList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&PublicIP{},
		&PublicIPList{},

		&StorageAccount{},
		&StorageAccountList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&FunctionApp{},
		&FunctionAppList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SearchService{},
		&SearchServiceList{},

		&Snapshot{},
		&SnapshotList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&Image{},
		&ImageList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DataFactory{},
		&DataFactoryList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&LbProbe{},
		&LbProbeList{},

		&Lb{},
		&LbList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&StorageContainer{},
		&StorageContainerList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
