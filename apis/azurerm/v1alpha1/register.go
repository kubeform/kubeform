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

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&BatchPool{},
		&BatchPoolList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&Iothub{},
		&IothubList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&StorageQueue{},
		&StorageQueueList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&StorageBlob{},
		&StorageBlobList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&Image{},
		&ImageList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&MysqlServer{},
		&MysqlServerList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataFactory{},
		&DataFactoryList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppService{},
		&AppServiceList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ManagementLock{},
		&ManagementLockList{},

		&PublicIP{},
		&PublicIPList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&Snapshot{},
		&SnapshotList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&FunctionApp{},
		&FunctionAppList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&StorageContainer{},
		&StorageContainerList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&Route{},
		&RouteList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RedisCache{},
		&RedisCacheList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&KeyVault{},
		&KeyVaultList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&Subnet{},
		&SubnetList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SearchService{},
		&SearchServiceList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LbProbe{},
		&LbProbeList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&Eventhub{},
		&EventhubList{},

		&BatchAccount{},
		&BatchAccountList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DevTestLab{},
		&DevTestLabList{},

		&LbRule{},
		&LbRuleList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DnsARecord{},
		&DnsARecordList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&Lb{},
		&LbList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&CdnProfile{},
		&CdnProfileList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&RouteTable{},
		&RouteTableList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&NotificationHub{},
		&NotificationHubList{},

		&SharedImage{},
		&SharedImageList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&IotDps{},
		&IotDpsList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DnsZone{},
		&DnsZoneList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&Firewall{},
		&FirewallList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageTable{},
		&StorageTableList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
