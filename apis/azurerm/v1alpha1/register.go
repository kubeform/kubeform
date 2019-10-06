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

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ApiManagement{},
		&ApiManagementList{},

		&DataFactory{},
		&DataFactoryList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&StorageQueue{},
		&StorageQueueList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&Lb{},
		&LbList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&AppService{},
		&AppServiceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&StorageBlob{},
		&StorageBlobList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&StorageContainer{},
		&StorageContainerList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&LbProbe{},
		&LbProbeList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&ManagementLock{},
		&ManagementLockList{},

		&StorageShare{},
		&StorageShareList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&Snapshot{},
		&SnapshotList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&IotDps{},
		&IotDpsList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ContainerService{},
		&ContainerServiceList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&Eventhub{},
		&EventhubList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&RouteTable{},
		&RouteTableList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&Route{},
		&RouteList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&DnsZone{},
		&DnsZoneList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MysqlServer{},
		&MysqlServerList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&Firewall{},
		&FirewallList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&PublicIP{},
		&PublicIPList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&StorageTable{},
		&StorageTableList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&BatchPool{},
		&BatchPoolList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NotificationHub{},
		&NotificationHubList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DnsARecord{},
		&DnsARecordList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&FunctionApp{},
		&FunctionAppList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&RedisCache{},
		&RedisCacheList{},

		&SharedImage{},
		&SharedImageList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&LbRule{},
		&LbRuleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&Subnet{},
		&SubnetList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&KeyVault{},
		&KeyVaultList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&SignalrService{},
		&SignalrServiceList{},

		&SqlServer{},
		&SqlServerList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&Image{},
		&ImageList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&SearchService{},
		&SearchServiceList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&MapsAccount{},
		&MapsAccountList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&Iothub{},
		&IothubList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
