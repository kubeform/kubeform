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

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&Firewall{},
		&FirewallList{},

		&StorageAccount{},
		&StorageAccountList{},

		&PublicIP{},
		&PublicIPList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DnsARecord{},
		&DnsARecordList{},

		&Eventhub{},
		&EventhubList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&Route{},
		&RouteList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MysqlServer{},
		&MysqlServerList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&StorageBlob{},
		&StorageBlobList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactory{},
		&DataFactoryList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&FunctionApp{},
		&FunctionAppList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&RouteTable{},
		&RouteTableList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&Snapshot{},
		&SnapshotList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&StorageTable{},
		&StorageTableList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&StorageShare{},
		&StorageShareList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppService{},
		&AppServiceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&Iothub{},
		&IothubList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&IotDps{},
		&IotDpsList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&Lb{},
		&LbList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&SharedImage{},
		&SharedImageList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&ManagementLock{},
		&ManagementLockList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RedisCache{},
		&RedisCacheList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LbRule{},
		&LbRuleList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsZone{},
		&DnsZoneList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&SearchService{},
		&SearchServiceList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&Subnet{},
		&SubnetList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&ApiManagement{},
		&ApiManagementList{},

		&Image{},
		&ImageList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&BatchPool{},
		&BatchPoolList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StorageContainer{},
		&StorageContainerList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&SqlServer{},
		&SqlServerList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&LbProbe{},
		&LbProbeList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&KeyVault{},
		&KeyVaultList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
