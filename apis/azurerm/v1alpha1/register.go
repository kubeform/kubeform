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

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DevTestLab{},
		&DevTestLabList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&SqlServer{},
		&SqlServerList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&Route{},
		&RouteList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DnsARecord{},
		&DnsARecordList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&SignalrService{},
		&SignalrServiceList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&PublicIP{},
		&PublicIPList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DnsZone{},
		&DnsZoneList{},

		&RedisCache{},
		&RedisCacheList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ContainerService{},
		&ContainerServiceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&AppService{},
		&AppServiceList{},

		&FunctionApp{},
		&FunctionAppList{},

		&Lb{},
		&LbList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&Snapshot{},
		&SnapshotList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&NotificationHub{},
		&NotificationHubList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&Firewall{},
		&FirewallList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DataFactory{},
		&DataFactoryList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagement{},
		&ApiManagementList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&LbProbe{},
		&LbProbeList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&Image{},
		&ImageList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&Eventhub{},
		&EventhubList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&KeyVault{},
		&KeyVaultList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&SearchService{},
		&SearchServiceList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MapsAccount{},
		&MapsAccountList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&StorageBlob{},
		&StorageBlobList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&Iothub{},
		&IothubList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&BatchAccount{},
		&BatchAccountList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&CdnProfile{},
		&CdnProfileList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&IotDps{},
		&IotDpsList{},

		&MariadbServer{},
		&MariadbServerList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&Subnet{},
		&SubnetList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&MysqlServer{},
		&MysqlServerList{},

		&StorageContainer{},
		&StorageContainerList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&BatchPool{},
		&BatchPoolList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&LbRule{},
		&LbRuleList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&StorageTable{},
		&StorageTableList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&ManagementLock{},
		&ManagementLockList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&StorageQueue{},
		&StorageQueueList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
