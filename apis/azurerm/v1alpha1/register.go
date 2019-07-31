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

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&BatchPool{},
		&BatchPoolList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&MariadbServer{},
		&MariadbServerList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&StorageShare{},
		&StorageShareList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DevTestLab{},
		&DevTestLabList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&IotDps{},
		&IotDpsList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&SignalrService{},
		&SignalrServiceList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AppService{},
		&AppServiceList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&Route{},
		&RouteList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&Image{},
		&ImageList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&StorageBlob{},
		&StorageBlobList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&StorageAccount{},
		&StorageAccountList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&DnsARecord{},
		&DnsARecordList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SqlServer{},
		&SqlServerList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&Iothub{},
		&IothubList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&SearchService{},
		&SearchServiceList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&Firewall{},
		&FirewallList{},

		&RedisCache{},
		&RedisCacheList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DataFactory{},
		&DataFactoryList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&Snapshot{},
		&SnapshotList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&BatchAccount{},
		&BatchAccountList{},

		&StorageTable{},
		&StorageTableList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PublicIP{},
		&PublicIPList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&CdnProfile{},
		&CdnProfileList{},

		&KeyVault{},
		&KeyVaultList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&Subnet{},
		&SubnetList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&LbProbe{},
		&LbProbeList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&Eventhub{},
		&EventhubList{},

		&Lb{},
		&LbList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StorageQueue{},
		&StorageQueueList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&DnsZone{},
		&DnsZoneList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&MysqlServer{},
		&MysqlServerList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&FunctionApp{},
		&FunctionAppList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
