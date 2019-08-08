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

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&DnsARecord{},
		&DnsARecordList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&SqlServer{},
		&SqlServerList{},

		&MariadbServer{},
		&MariadbServerList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&StorageShare{},
		&StorageShareList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&RedisCache{},
		&RedisCacheList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&NotificationHub{},
		&NotificationHubList{},

		&SearchService{},
		&SearchServiceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&BatchPool{},
		&BatchPoolList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&PublicIP{},
		&PublicIPList{},

		&RouteTable{},
		&RouteTableList{},

		&StorageBlob{},
		&StorageBlobList{},

		&KeyVault{},
		&KeyVaultList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&Subnet{},
		&SubnetList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&LbRule{},
		&LbRuleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&BatchAccount{},
		&BatchAccountList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&Iothub{},
		&IothubList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&ApiManagement{},
		&ApiManagementList{},

		&AppService{},
		&AppServiceList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&Snapshot{},
		&SnapshotList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&LbProbe{},
		&LbProbeList{},

		&Lb{},
		&LbList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&Eventhub{},
		&EventhubList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&ManagementLock{},
		&ManagementLockList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&Image{},
		&ImageList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&MysqlServer{},
		&MysqlServerList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&Firewall{},
		&FirewallList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&StorageAccount{},
		&StorageAccountList{},

		&FunctionApp{},
		&FunctionAppList{},

		&IotDps{},
		&IotDpsList{},

		&StorageTable{},
		&StorageTableList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&SharedImage{},
		&SharedImageList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&DnsZone{},
		&DnsZoneList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&Route{},
		&RouteList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
