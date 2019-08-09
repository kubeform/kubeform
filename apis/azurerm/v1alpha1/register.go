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

		&DevTestLab{},
		&DevTestLabList{},

		&LbProbe{},
		&LbProbeList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactory{},
		&DataFactoryList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&PublicIP{},
		&PublicIPList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StorageQueue{},
		&StorageQueueList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&RouteTable{},
		&RouteTableList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&Lb{},
		&LbList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&StorageTable{},
		&StorageTableList{},

		&AppService{},
		&AppServiceList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&FunctionApp{},
		&FunctionAppList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&StorageShare{},
		&StorageShareList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&Route{},
		&RouteList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&KeyVault{},
		&KeyVaultList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&MapsAccount{},
		&MapsAccountList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&SignalrService{},
		&SignalrServiceList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DnsZone{},
		&DnsZoneList{},

		&Firewall{},
		&FirewallList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&Subnet{},
		&SubnetList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&RedisCache{},
		&RedisCacheList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&Image{},
		&ImageList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&SearchService{},
		&SearchServiceList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&IotDps{},
		&IotDpsList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&StorageContainer{},
		&StorageContainerList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&Eventhub{},
		&EventhubList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&MariadbServer{},
		&MariadbServerList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ContainerService{},
		&ContainerServiceList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&Iothub{},
		&IothubList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&StorageBlob{},
		&StorageBlobList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&BatchPool{},
		&BatchPoolList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&SharedImage{},
		&SharedImageList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ManagementLock{},
		&ManagementLockList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&BatchAccount{},
		&BatchAccountList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&LbRule{},
		&LbRuleList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SqlServer{},
		&SqlServerList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MysqlServer{},
		&MysqlServerList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DnsARecord{},
		&DnsARecordList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&Snapshot{},
		&SnapshotList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
