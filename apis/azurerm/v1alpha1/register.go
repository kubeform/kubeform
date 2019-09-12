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

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&StorageBlob{},
		&StorageBlobList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&BatchPool{},
		&BatchPoolList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&SharedImage{},
		&SharedImageList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&DnsARecord{},
		&DnsARecordList{},

		&SignalrService{},
		&SignalrServiceList{},

		&StorageAccount{},
		&StorageAccountList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&Firewall{},
		&FirewallList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&StorageShare{},
		&StorageShareList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&BatchAccount{},
		&BatchAccountList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&Route{},
		&RouteList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&AppService{},
		&AppServiceList{},

		&Image{},
		&ImageList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&MapsAccount{},
		&MapsAccountList{},

		&SearchService{},
		&SearchServiceList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&Eventhub{},
		&EventhubList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&Snapshot{},
		&SnapshotList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&PublicIP{},
		&PublicIPList{},

		&RedisCache{},
		&RedisCacheList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&NotificationHub{},
		&NotificationHubList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&LbProbe{},
		&LbProbeList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&StorageTable{},
		&StorageTableList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&SqlServer{},
		&SqlServerList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MysqlServer{},
		&MysqlServerList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&Lb{},
		&LbList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DevTestLab{},
		&DevTestLabList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&KeyVault{},
		&KeyVaultList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MariadbServer{},
		&MariadbServerList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&Iothub{},
		&IothubList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&StorageContainer{},
		&StorageContainerList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&LbRule{},
		&LbRuleList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ContainerService{},
		&ContainerServiceList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&RouteTable{},
		&RouteTableList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DnsZone{},
		&DnsZoneList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&IotDps{},
		&IotDpsList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
