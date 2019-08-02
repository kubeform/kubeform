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

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&StorageContainer{},
		&StorageContainerList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&Image{},
		&ImageList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SearchService{},
		&SearchServiceList{},

		&SharedImage{},
		&SharedImageList{},

		&SignalrService{},
		&SignalrServiceList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ManagementLock{},
		&ManagementLockList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LbRule{},
		&LbRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&DnsZone{},
		&DnsZoneList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&BatchPool{},
		&BatchPoolList{},

		&KeyVault{},
		&KeyVaultList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&StorageAccount{},
		&StorageAccountList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&RedisCache{},
		&RedisCacheList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&StorageShare{},
		&StorageShareList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&LbProbe{},
		&LbProbeList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&Subnet{},
		&SubnetList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StorageBlob{},
		&StorageBlobList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&PublicIP{},
		&PublicIPList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagement{},
		&ApiManagementList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&Iothub{},
		&IothubList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StorageQueue{},
		&StorageQueueList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&Lb{},
		&LbList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DataFactory{},
		&DataFactoryList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&Firewall{},
		&FirewallList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&DevTestLab{},
		&DevTestLabList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&DnsARecord{},
		&DnsARecordList{},

		&Eventhub{},
		&EventhubList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&SqlServer{},
		&SqlServerList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&FunctionApp{},
		&FunctionAppList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&AppService{},
		&AppServiceList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&ContainerService{},
		&ContainerServiceList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&Snapshot{},
		&SnapshotList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&CdnProfile{},
		&CdnProfileList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageTable{},
		&StorageTableList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NotificationHub{},
		&NotificationHubList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&IotDps{},
		&IotDpsList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&Route{},
		&RouteList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
