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

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&Snapshot{},
		&SnapshotList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&IotDps{},
		&IotDpsList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&MysqlServer{},
		&MysqlServerList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&Image{},
		&ImageList{},

		&DataFactory{},
		&DataFactoryList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&DnsARecord{},
		&DnsARecordList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&StorageContainer{},
		&StorageContainerList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&StorageBlob{},
		&StorageBlobList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&CdnProfile{},
		&CdnProfileList{},

		&LbRule{},
		&LbRuleList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ApiManagement{},
		&ApiManagementList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&BatchPool{},
		&BatchPoolList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LbProbe{},
		&LbProbeList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AppService{},
		&AppServiceList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ManagementLock{},
		&ManagementLockList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&PublicIP{},
		&PublicIPList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&DnsZone{},
		&DnsZoneList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&KeyVault{},
		&KeyVaultList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&Iothub{},
		&IothubList{},

		&Lb{},
		&LbList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Route{},
		&RouteList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ContainerService{},
		&ContainerServiceList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&SharedImage{},
		&SharedImageList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&NotificationHub{},
		&NotificationHubList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&DevTestLab{},
		&DevTestLabList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&Firewall{},
		&FirewallList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&SearchService{},
		&SearchServiceList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&RedisCache{},
		&RedisCacheList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&MariadbServer{},
		&MariadbServerList{},

		&Eventhub{},
		&EventhubList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&SignalrService{},
		&SignalrServiceList{},

		&StorageTable{},
		&StorageTableList{},

		&BatchAccount{},
		&BatchAccountList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SqlServer{},
		&SqlServerList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ContainerGroup{},
		&ContainerGroupList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
