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

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&AppService{},
		&AppServiceList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&StorageShare{},
		&StorageShareList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&Eventhub{},
		&EventhubList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&CdnProfile{},
		&CdnProfileList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&BatchPool{},
		&BatchPoolList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&ManagementLock{},
		&ManagementLockList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&LbProbe{},
		&LbProbeList{},

		&MapsAccount{},
		&MapsAccountList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ContainerService{},
		&ContainerServiceList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&SearchService{},
		&SearchServiceList{},

		&StorageAccount{},
		&StorageAccountList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&StorageTable{},
		&StorageTableList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ApiManagement{},
		&ApiManagementList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&IotDps{},
		&IotDpsList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&Snapshot{},
		&SnapshotList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&DataFactory{},
		&DataFactoryList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&Image{},
		&ImageList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&LbRule{},
		&LbRuleList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&DnsZone{},
		&DnsZoneList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&DevTestLab{},
		&DevTestLabList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&RedisCache{},
		&RedisCacheList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&DnsARecord{},
		&DnsARecordList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&FunctionApp{},
		&FunctionAppList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&Subnet{},
		&SubnetList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&Firewall{},
		&FirewallList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&Iothub{},
		&IothubList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MariadbServer{},
		&MariadbServerList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&KeyVault{},
		&KeyVaultList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&RouteTable{},
		&RouteTableList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&Route{},
		&RouteList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&NotificationHub{},
		&NotificationHubList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&StorageBlob{},
		&StorageBlobList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&PublicIP{},
		&PublicIPList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&StorageContainer{},
		&StorageContainerList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&Lb{},
		&LbList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
