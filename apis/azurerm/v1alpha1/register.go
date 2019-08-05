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

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&Route{},
		&RouteList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DataFactory{},
		&DataFactoryList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&KeyVault{},
		&KeyVaultList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&AppService{},
		&AppServiceList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&RouteTable{},
		&RouteTableList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&BatchPool{},
		&BatchPoolList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&RedisCache{},
		&RedisCacheList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&Iothub{},
		&IothubList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&DnsZone{},
		&DnsZoneList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&SharedImage{},
		&SharedImageList{},

		&SignalrService{},
		&SignalrServiceList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ManagementLock{},
		&ManagementLockList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&PublicIP{},
		&PublicIPList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LbProbe{},
		&LbProbeList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&StorageAccount{},
		&StorageAccountList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagement{},
		&ApiManagementList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageTable{},
		&StorageTableList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NotificationHub{},
		&NotificationHubList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&Lb{},
		&LbList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&StorageShare{},
		&StorageShareList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&DnsARecord{},
		&DnsARecordList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&StorageContainer{},
		&StorageContainerList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&BatchAccount{},
		&BatchAccountList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&Firewall{},
		&FirewallList{},

		&SearchService{},
		&SearchServiceList{},

		&Snapshot{},
		&SnapshotList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&Image{},
		&ImageList{},

		&IotDps{},
		&IotDpsList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&LbRule{},
		&LbRuleList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MariadbServer{},
		&MariadbServerList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&Eventhub{},
		&EventhubList{},

		&FunctionApp{},
		&FunctionAppList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
