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

		&ResourceGroup{},
		&ResourceGroupList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&Eventhub{},
		&EventhubList{},

		&FunctionApp{},
		&FunctionAppList{},

		&SignalrService{},
		&SignalrServiceList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&DataFactory{},
		&DataFactoryList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&Route{},
		&RouteList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LbRule{},
		&LbRuleList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&Lb{},
		&LbList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&DnsZone{},
		&DnsZoneList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&Firewall{},
		&FirewallList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&StorageTable{},
		&StorageTableList{},

		&Iothub{},
		&IothubList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&NotificationHub{},
		&NotificationHubList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&Snapshot{},
		&SnapshotList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&ContainerService{},
		&ContainerServiceList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&Subnet{},
		&SubnetList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RedisCache{},
		&RedisCacheList{},

		&StorageQueue{},
		&StorageQueueList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&KeyVault{},
		&KeyVaultList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&SearchService{},
		&SearchServiceList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&DevTestLab{},
		&DevTestLabList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&PublicIP{},
		&PublicIPList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&IotDps{},
		&IotDpsList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&BatchPool{},
		&BatchPoolList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppService{},
		&AppServiceList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&LbProbe{},
		&LbProbeList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&BatchAccount{},
		&BatchAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&MysqlServer{},
		&MysqlServerList{},

		&StorageShare{},
		&StorageShareList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&Image{},
		&ImageList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ManagementLock{},
		&ManagementLockList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ApiManagement{},
		&ApiManagementList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&StorageContainer{},
		&StorageContainerList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
