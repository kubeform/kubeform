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

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&SharedImage{},
		&SharedImageList{},

		&AppService{},
		&AppServiceList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StorageTable{},
		&StorageTableList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&Route{},
		&RouteList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&BatchAccount{},
		&BatchAccountList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&KeyVault{},
		&KeyVaultList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NotificationHub{},
		&NotificationHubList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&Eventhub{},
		&EventhubList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LbProbe{},
		&LbProbeList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ManagementLock{},
		&ManagementLockList{},

		&StorageAccount{},
		&StorageAccountList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&FunctionApp{},
		&FunctionAppList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&BatchPool{},
		&BatchPoolList{},

		&DevTestLab{},
		&DevTestLabList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&RouteTable{},
		&RouteTableList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&RedisCache{},
		&RedisCacheList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&CdnProfile{},
		&CdnProfileList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MariadbServer{},
		&MariadbServerList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&Snapshot{},
		&SnapshotList{},

		&ContainerService{},
		&ContainerServiceList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageShare{},
		&StorageShareList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&IotDps{},
		&IotDpsList{},

		&SignalrService{},
		&SignalrServiceList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DnsZone{},
		&DnsZoneList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DataFactory{},
		&DataFactoryList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&Subnet{},
		&SubnetList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&Firewall{},
		&FirewallList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&Iothub{},
		&IothubList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&PublicIP{},
		&PublicIPList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MapsAccount{},
		&MapsAccountList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&DnsARecord{},
		&DnsARecordList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&Lb{},
		&LbList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&MysqlServer{},
		&MysqlServerList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&Image{},
		&ImageList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&StorageQueue{},
		&StorageQueueList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&SearchService{},
		&SearchServiceList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&SqlServer{},
		&SqlServerList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
