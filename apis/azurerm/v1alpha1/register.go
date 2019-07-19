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

		&Image{},
		&ImageList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&ContainerService{},
		&ContainerServiceList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&Eventhub{},
		&EventhubList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DevTestLab{},
		&DevTestLabList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&Snapshot{},
		&SnapshotList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&IotDps{},
		&IotDpsList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&RouteTable{},
		&RouteTableList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&BatchPool{},
		&BatchPoolList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&BatchAccount{},
		&BatchAccountList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&SharedImage{},
		&SharedImageList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&StorageTable{},
		&StorageTableList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NotificationHub{},
		&NotificationHubList{},

		&SqlServer{},
		&SqlServerList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&PublicIP{},
		&PublicIPList{},

		&Route{},
		&RouteList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&KeyVault{},
		&KeyVaultList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&StorageShare{},
		&StorageShareList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&ManagementLock{},
		&ManagementLockList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&SignalrService{},
		&SignalrServiceList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&RedisCache{},
		&RedisCacheList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&SearchService{},
		&SearchServiceList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MysqlServer{},
		&MysqlServerList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&AppService{},
		&AppServiceList{},

		&DnsZone{},
		&DnsZoneList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&DnsARecord{},
		&DnsARecordList{},

		&Firewall{},
		&FirewallList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&DataFactory{},
		&DataFactoryList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&Iothub{},
		&IothubList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MariadbServer{},
		&MariadbServerList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&Lb{},
		&LbList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&Subnet{},
		&SubnetList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&FunctionApp{},
		&FunctionAppList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&StorageAccount{},
		&StorageAccountList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&LbProbe{},
		&LbProbeList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&StorageBlob{},
		&StorageBlobList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
