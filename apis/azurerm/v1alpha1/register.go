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

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&BatchPool{},
		&BatchPoolList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsZone{},
		&DnsZoneList{},

		&PublicIP{},
		&PublicIPList{},

		&DataFactory{},
		&DataFactoryList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&LbNATRule{},
		&LbNATRuleList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&RouteTable{},
		&RouteTableList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&SignalrService{},
		&SignalrServiceList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&Iothub{},
		&IothubList{},

		&KeyVault{},
		&KeyVaultList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StorageShare{},
		&StorageShareList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&StorageContainer{},
		&StorageContainerList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&Eventhub{},
		&EventhubList{},

		&Firewall{},
		&FirewallList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&FunctionApp{},
		&FunctionAppList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&Subnet{},
		&SubnetList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NotificationHub{},
		&NotificationHubList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&BatchAccount{},
		&BatchAccountList{},

		&MysqlServer{},
		&MysqlServerList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&ManagementLock{},
		&ManagementLockList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&SearchService{},
		&SearchServiceList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DnsARecord{},
		&DnsARecordList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AppService{},
		&AppServiceList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&IotDps{},
		&IotDpsList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&LbProbe{},
		&LbProbeList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&Lb{},
		&LbList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SharedImage{},
		&SharedImageList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&Image{},
		&ImageList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageTable{},
		&StorageTableList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DevTestLab{},
		&DevTestLabList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&RedisCache{},
		&RedisCacheList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&Snapshot{},
		&SnapshotList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&Route{},
		&RouteList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
