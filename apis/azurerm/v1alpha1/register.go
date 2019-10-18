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

		&VirtualMachine{},
		&VirtualMachineList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&LbProbe{},
		&LbProbeList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&SharedImage{},
		&SharedImageList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&KeyVault{},
		&KeyVaultList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&RedisCache{},
		&RedisCacheList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&Firewall{},
		&FirewallList{},

		&Iothub{},
		&IothubList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&MariadbServer{},
		&MariadbServerList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&AppService{},
		&AppServiceList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&LbRule{},
		&LbRuleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DnsARecord{},
		&DnsARecordList{},

		&SqlServer{},
		&SqlServerList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&SignalrService{},
		&SignalrServiceList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&BatchPool{},
		&BatchPoolList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&Route{},
		&RouteList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&Eventhub{},
		&EventhubList{},

		&DataFactory{},
		&DataFactoryList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&StorageQueue{},
		&StorageQueueList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&StorageAccount{},
		&StorageAccountList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageContainer{},
		&StorageContainerList{},

		&BatchAccount{},
		&BatchAccountList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&Image{},
		&ImageList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&RouteTable{},
		&RouteTableList{},

		&SearchService{},
		&SearchServiceList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&FunctionApp{},
		&FunctionAppList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&CdnProfile{},
		&CdnProfileList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&DnsZone{},
		&DnsZoneList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&MapsAccount{},
		&MapsAccountList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&StorageTable{},
		&StorageTableList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ContainerService{},
		&ContainerServiceList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&Lb{},
		&LbList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagement{},
		&ApiManagementList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DevTestLab{},
		&DevTestLabList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&IotDps{},
		&IotDpsList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&PublicIP{},
		&PublicIPList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&Subnet{},
		&SubnetList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Snapshot{},
		&SnapshotList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&NotificationHub{},
		&NotificationHubList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
