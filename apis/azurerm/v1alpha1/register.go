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

		&BatchAccount{},
		&BatchAccountList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&IotDps{},
		&IotDpsList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&Image{},
		&ImageList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&Route{},
		&RouteList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&RouteTable{},
		&RouteTableList{},

		&SearchService{},
		&SearchServiceList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&SqlServer{},
		&SqlServerList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&AppService{},
		&AppServiceList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MapsAccount{},
		&MapsAccountList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DnsZone{},
		&DnsZoneList{},

		&KeyVault{},
		&KeyVaultList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&LbRule{},
		&LbRuleList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&ContainerService{},
		&ContainerServiceList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&CdnProfile{},
		&CdnProfileList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&Subnet{},
		&SubnetList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&StorageTable{},
		&StorageTableList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&StorageBlob{},
		&StorageBlobList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&Iothub{},
		&IothubList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&Lb{},
		&LbList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&Firewall{},
		&FirewallList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&Snapshot{},
		&SnapshotList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MysqlServer{},
		&MysqlServerList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&ApiManagement{},
		&ApiManagementList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&BatchPool{},
		&BatchPoolList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&RedisCache{},
		&RedisCacheList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&LbProbe{},
		&LbProbeList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&MariadbServer{},
		&MariadbServerList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DataFactory{},
		&DataFactoryList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&SharedImage{},
		&SharedImageList{},

		&SignalrService{},
		&SignalrServiceList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&Eventhub{},
		&EventhubList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PublicIP{},
		&PublicIPList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
