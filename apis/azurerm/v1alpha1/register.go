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

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MariadbServer{},
		&MariadbServerList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&BatchAccount{},
		&BatchAccountList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&FunctionApp{},
		&FunctionAppList{},

		&ManagementLock{},
		&ManagementLockList{},

		&Route{},
		&RouteList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&StorageShare{},
		&StorageShareList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&BatchPool{},
		&BatchPoolList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&KeyVault{},
		&KeyVaultList{},

		&MysqlServer{},
		&MysqlServerList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&LbRule{},
		&LbRuleList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&Image{},
		&ImageList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&RouteTable{},
		&RouteTableList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&Iothub{},
		&IothubList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&LbProbe{},
		&LbProbeList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&Snapshot{},
		&SnapshotList{},

		&StorageContainer{},
		&StorageContainerList{},

		&AppService{},
		&AppServiceList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NotificationHub{},
		&NotificationHubList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&SearchService{},
		&SearchServiceList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DataFactory{},
		&DataFactoryList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StorageQueue{},
		&StorageQueueList{},

		&Firewall{},
		&FirewallList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&ApiManagement{},
		&ApiManagementList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&StorageTable{},
		&StorageTableList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ContainerService{},
		&ContainerServiceList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&Eventhub{},
		&EventhubList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&PublicIP{},
		&PublicIPList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DnsARecord{},
		&DnsARecordList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&RedisCache{},
		&RedisCacheList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&DevTestLab{},
		&DevTestLabList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&IotDps{},
		&IotDpsList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&StorageAccount{},
		&StorageAccountList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&DnsZone{},
		&DnsZoneList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&Lb{},
		&LbList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SharedImage{},
		&SharedImageList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
