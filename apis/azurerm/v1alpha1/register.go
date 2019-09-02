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
    
        &DnsARecord{},
        &DnsARecordList{},
    
        &RedisFirewallRule{},
        &RedisFirewallRuleList{},
    
        &Snapshot{},
        &SnapshotList{},
    
        &AutomationSchedule{},
        &AutomationScheduleList{},
    
        &AutomationVariableInt{},
        &AutomationVariableIntList{},
    
        &FirewallNetworkRuleCollection{},
        &FirewallNetworkRuleCollectionList{},
    
        &MariadbServer{},
        &MariadbServerList{},
    
        &SecurityCenterWorkspace{},
        &SecurityCenterWorkspaceList{},
    
        &SignalrService{},
        &SignalrServiceList{},
    
        &SqlVirtualNetworkRule{},
        &SqlVirtualNetworkRuleList{},
    
        &AutomationDscConfiguration{},
        &AutomationDscConfigurationList{},
    
        &DataLakeAnalyticsAccount{},
        &DataLakeAnalyticsAccountList{},
    
        &DevTestVirtualNetwork{},
        &DevTestVirtualNetworkList{},
    
        &HdinsightRserverCluster{},
        &HdinsightRserverClusterList{},
    
        &ManagedDisk{},
        &ManagedDiskList{},
    
        &MapsAccount{},
        &MapsAccountList{},
    
        &MonitorMetricAlertrule{},
        &MonitorMetricAlertruleList{},
    
        &PrivateDNSZone{},
        &PrivateDNSZoneList{},
    
        &ApiManagementAPIPolicy{},
        &ApiManagementAPIPolicyList{},
    
        &ApiManagementCertificate{},
        &ApiManagementCertificateList{},
    
        &DnsZone{},
        &DnsZoneList{},
    
        &SharedImage{},
        &SharedImageList{},
    
        &StorageShareDirectory{},
        &StorageShareDirectoryList{},
    
        &VirtualMachineDataDiskAttachment{},
        &VirtualMachineDataDiskAttachmentList{},
    
        &ResourceGroup{},
        &ResourceGroupList{},
    
        &DataFactoryLinkedServiceMysql{},
        &DataFactoryLinkedServiceMysqlList{},
    
        &ExpressRouteCircuitPeering{},
        &ExpressRouteCircuitPeeringList{},
    
        &HdinsightSparkCluster{},
        &HdinsightSparkClusterList{},
    
        &LbBackendAddressPool{},
        &LbBackendAddressPoolList{},
    
        &VirtualMachine{},
        &VirtualMachineList{},
    
        &ApiManagementProduct{},
        &ApiManagementProductList{},
    
        &CosmosdbAccount{},
        &CosmosdbAccountList{},
    
        &DataFactoryDatasetSQLServerTable{},
        &DataFactoryDatasetSQLServerTableList{},
    
        &RedisCache{},
        &RedisCacheList{},
    
        &AppService{},
        &AppServiceList{},
    
        &LogicAppTriggerCustom{},
        &LogicAppTriggerCustomList{},
    
        &NotificationHubAuthorizationRule{},
        &NotificationHubAuthorizationRuleList{},
    
        &StreamAnalyticsFunctionJavascriptUdf{},
        &StreamAnalyticsFunctionJavascriptUdfList{},
    
        &SubnetNetworkSecurityGroupAssociation{},
        &SubnetNetworkSecurityGroupAssociationList{},
    
        &VirtualNetworkGatewayConnection{},
        &VirtualNetworkGatewayConnectionList{},
    
        &DdosProtectionPlan{},
        &DdosProtectionPlanList{},
    
        &MysqlVirtualNetworkRule{},
        &MysqlVirtualNetworkRuleList{},
    
        &NetworkPacketCapture{},
        &NetworkPacketCaptureList{},
    
        &ApiManagementProductGroup{},
        &ApiManagementProductGroupList{},
    
        &AzureadApplication{},
        &AzureadApplicationList{},
    
        &CognitiveAccount{},
        &CognitiveAccountList{},
    
        &SharedImageVersion{},
        &SharedImageVersionList{},
    
        &ApiManagementGroup{},
        &ApiManagementGroupList{},
    
        &KeyVault{},
        &KeyVaultList{},
    
        &PostgresqlVirtualNetworkRule{},
        &PostgresqlVirtualNetworkRuleList{},
    
        &CdnEndpoint{},
        &CdnEndpointList{},
    
        &DataFactoryLinkedServiceSQLServer{},
        &DataFactoryLinkedServiceSQLServerList{},
    
        &HdinsightStormCluster{},
        &HdinsightStormClusterList{},
    
        &KeyVaultCertificate{},
        &KeyVaultCertificateList{},
    
        &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
        &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},
    
        &ApiManagementAPISchema{},
        &ApiManagementAPISchemaList{},
    
        &AppServiceSlot{},
        &AppServiceSlotList{},
    
        &AutomationVariableBool{},
        &AutomationVariableBoolList{},
    
        &ServicebusNamespaceAuthorizationRule{},
        &ServicebusNamespaceAuthorizationRuleList{},
    
        &PolicySetDefinition{},
        &PolicySetDefinitionList{},
    
        &StreamAnalyticsOutputBlob{},
        &StreamAnalyticsOutputBlobList{},
    
        &EventhubNamespace_{},
        &EventhubNamespace_List{},
    
        &IotDpsCertificate{},
        &IotDpsCertificateList{},
    
        &NotificationHub{},
        &NotificationHubList{},
    
        &LbNATPool{},
        &LbNATPoolList{},
    
        &StreamAnalyticsStreamInputBlob{},
        &StreamAnalyticsStreamInputBlobList{},
    
        &TemplateDeployment{},
        &TemplateDeploymentList{},
    
        &UserAssignedIdentity{},
        &UserAssignedIdentityList{},
    
        &ApplicationInsightsWebTest{},
        &ApplicationInsightsWebTestList{},
    
        &DnsNsRecord{},
        &DnsNsRecordList{},
    
        &SecurityCenterContact{},
        &SecurityCenterContactList{},
    
        &HdinsightMlServicesCluster{},
        &HdinsightMlServicesClusterList{},
    
        &StorageQueue{},
        &StorageQueueList{},
    
        &ApplicationInsightsAPIKey{},
        &ApplicationInsightsAPIKeyList{},
    
        &DataFactoryPipeline{},
        &DataFactoryPipelineList{},
    
        &HdinsightInteractiveQueryCluster{},
        &HdinsightInteractiveQueryClusterList{},
    
        &FunctionApp{},
        &FunctionAppList{},
    
        &LbNATRule{},
        &LbNATRuleList{},
    
        &PublicIP{},
        &PublicIPList{},
    
        &ApiManagementAPI{},
        &ApiManagementAPIList{},
    
        &ApiManagementGroupUser{},
        &ApiManagementGroupUserList{},
    
        &ExpressRouteCircuit{},
        &ExpressRouteCircuitList{},
    
        &ServicebusQueueAuthorizationRule{},
        &ServicebusQueueAuthorizationRuleList{},
    
        &ServicebusQueue{},
        &ServicebusQueueList{},
    
        &ApiManagementOpenidConnectProvider{},
        &ApiManagementOpenidConnectProviderList{},
    
        &DevTestLab{},
        &DevTestLabList{},
    
        &PostgresqlServer{},
        &PostgresqlServerList{},
    
        &MediaServicesAccount{},
        &MediaServicesAccountList{},
    
        &VirtualMachineExtension{},
        &VirtualMachineExtensionList{},
    
        &ConnectionMonitor{},
        &ConnectionMonitorList{},
    
        &CosmosdbMongoCollection{},
        &CosmosdbMongoCollectionList{},
    
        &DatabricksWorkspace{},
        &DatabricksWorkspaceList{},
    
        &MysqlFirewallRule{},
        &MysqlFirewallRuleList{},
    
        &StreamAnalyticsStreamInputEventhub{},
        &StreamAnalyticsStreamInputEventhubList{},
    
        &ApiManagementAPIOperation{},
        &ApiManagementAPIOperationList{},
    
        &BatchPool{},
        &BatchPoolList{},
    
        &DataLakeStoreFile{},
        &DataLakeStoreFileList{},
    
        &DevTestPolicy{},
        &DevTestPolicyList{},
    
        &RecoveryServicesProtectionPolicyVm{},
        &RecoveryServicesProtectionPolicyVmList{},
    
        &DataLakeAnalyticsFirewallRule{},
        &DataLakeAnalyticsFirewallRuleList{},
    
        &RelayNamespace{},
        &RelayNamespaceList{},
    
        &NetworkProfile{},
        &NetworkProfileList{},
    
        &PublicIPPrefix{},
        &PublicIPPrefixList{},
    
        &StorageTableEntity{},
        &StorageTableEntityList{},
    
        &VirtualMachineScaleSet{},
        &VirtualMachineScaleSetList{},
    
        &MonitorActionGroup{},
        &MonitorActionGroupList{},
    
        &NetworkInterfaceNATRuleAssociation{},
        &NetworkInterfaceNATRuleAssociationList{},
    
        &RecoveryServicesVault{},
        &RecoveryServicesVaultList{},
    
        &LogAnalyticsSolution{},
        &LogAnalyticsSolutionList{},
    
        &MonitorDiagnosticSetting{},
        &MonitorDiagnosticSettingList{},
    
        &ApiManagementSubscription{},
        &ApiManagementSubscriptionList{},
    
        &DataLakeStore{},
        &DataLakeStoreList{},
    
        &ServicebusTopic{},
        &ServicebusTopicList{},
    
        &StorageBlob{},
        &StorageBlobList{},
    
        &ApiManagementAPIVersionSet{},
        &ApiManagementAPIVersionSetList{},
    
        &BatchApplication{},
        &BatchApplicationList{},
    
        &LogicAppActionHTTP{},
        &LogicAppActionHTTPList{},
    
        &CosmosdbMongoDatabase{},
        &CosmosdbMongoDatabaseList{},
    
        &NetworkSecurityGroup{},
        &NetworkSecurityGroupList{},
    
        &ServicebusTopicAuthorizationRule{},
        &ServicebusTopicAuthorizationRuleList{},
    
        &CdnProfile{},
        &CdnProfileList{},
    
        &Iothub{},
        &IothubList{},
    
        &VirtualNetwork{},
        &VirtualNetworkList{},
    
        &EventgridEventSubscription{},
        &EventgridEventSubscriptionList{},
    
        &Firewall{},
        &FirewallList{},
    
        &SchedulerJobCollection{},
        &SchedulerJobCollectionList{},
    
        &SchedulerJob{},
        &SchedulerJobList{},
    
        &ApiManagementLogger{},
        &ApiManagementLoggerList{},
    
        &ContainerRegistry{},
        &ContainerRegistryList{},
    
        &DnsAaaaRecord{},
        &DnsAaaaRecordList{},
    
        &NetworkConnectionMonitor{},
        &NetworkConnectionMonitorList{},
    
        &DnsTxtRecord{},
        &DnsTxtRecordList{},
    
        &LogAnalyticsLinkedService{},
        &LogAnalyticsLinkedServiceList{},
    
        &MonitorActivityLogAlert{},
        &MonitorActivityLogAlertList{},
    
        &SecurityCenterSubscriptionPricing{},
        &SecurityCenterSubscriptionPricingList{},
    
        &MonitorMetricAlert{},
        &MonitorMetricAlertList{},
    
        &PolicyDefinition{},
        &PolicyDefinitionList{},
    
        &AppServiceActiveSlot{},
        &AppServiceActiveSlotList{},
    
        &AzureadServicePrincipalPassword{},
        &AzureadServicePrincipalPasswordList{},
    
        &LocalNetworkGateway{},
        &LocalNetworkGatewayList{},
    
        &ManagementLock{},
        &ManagementLockList{},
    
        &PostgresqlDatabase{},
        &PostgresqlDatabaseList{},
    
        &KubernetesCluster{},
        &KubernetesClusterList{},
    
        &MariadbFirewallRule{},
        &MariadbFirewallRuleList{},
    
        &ApiManagement{},
        &ApiManagementList{},
    
        &ContainerGroup{},
        &ContainerGroupList{},
    
        &DnsMxRecord{},
        &DnsMxRecordList{},
    
        &RoleAssignment{},
        &RoleAssignmentList{},
    
        &ServicebusNamespace{},
        &ServicebusNamespaceList{},
    
        &DataLakeStoreFirewallRule{},
        &DataLakeStoreFirewallRuleList{},
    
        &RecoveryServicesProtectedVm{},
        &RecoveryServicesProtectedVmList{},
    
        &SharedImageGallery{},
        &SharedImageGalleryList{},
    
        &EventhubNamespaceAuthorizationRule{},
        &EventhubNamespaceAuthorizationRuleList{},
    
        &Lb{},
        &LbList{},
    
        &NetworkSecurityRule{},
        &NetworkSecurityRuleList{},
    
        &StreamAnalyticsOutputServicebusQueue{},
        &StreamAnalyticsOutputServicebusQueueList{},
    
        &AnalysisServicesServer{},
        &AnalysisServicesServerList{},
    
        &AutomationCredential{},
        &AutomationCredentialList{},
    
        &EventgridDomain{},
        &EventgridDomainList{},
    
        &CosmosdbCassandraKeyspace{},
        &CosmosdbCassandraKeyspaceList{},
    
        &SqlActiveDirectoryAdministrator{},
        &SqlActiveDirectoryAdministratorList{},
    
        &StorageAccount{},
        &StorageAccountList{},
    
        &StreamAnalyticsStreamInputIothub{},
        &StreamAnalyticsStreamInputIothubList{},
    
        &ApiManagementProperty{},
        &ApiManagementPropertyList{},
    
        &AutomationRunbook{},
        &AutomationRunbookList{},
    
        &AutomationVariableString{},
        &AutomationVariableStringList{},
    
        &MssqlElasticpool{},
        &MssqlElasticpoolList{},
    
        &PrivateDNSARecord{},
        &PrivateDNSARecordList{},
    
        &ServicebusSubscription{},
        &ServicebusSubscriptionList{},
    
        &StorageTable{},
        &StorageTableList{},
    
        &AutomationVariableDatetime{},
        &AutomationVariableDatetimeList{},
    
        &DnsCnameRecord{},
        &DnsCnameRecordList{},
    
        &LogAnalyticsWorkspace{},
        &LogAnalyticsWorkspaceList{},
    
        &ApiManagementBackend{},
        &ApiManagementBackendList{},
    
        &DataFactoryDatasetMysql{},
        &DataFactoryDatasetMysqlList{},
    
        &IothubConsumerGroup{},
        &IothubConsumerGroupList{},
    
        &HdinsightKafkaCluster{},
        &HdinsightKafkaClusterList{},
    
        &KeyVaultAccessPolicy{},
        &KeyVaultAccessPolicyList{},
    
        &NetworkDdosProtectionPlan{},
        &NetworkDdosProtectionPlanList{},
    
        &Route{},
        &RouteList{},
    
        &ApiManagementUser{},
        &ApiManagementUserList{},
    
        &AutomationModule{},
        &AutomationModuleList{},
    
        &DnsPtrRecord{},
        &DnsPtrRecordList{},
    
        &Eventhub{},
        &EventhubList{},
    
        &RouteTable{},
        &RouteTableList{},
    
        &SqlDatabase{},
        &SqlDatabaseList{},
    
        &AvailabilitySet{},
        &AvailabilitySetList{},
    
        &LogicAppTriggerHTTPRequest{},
        &LogicAppTriggerHTTPRequestList{},
    
        &TrafficManagerProfile{},
        &TrafficManagerProfileList{},
    
        &AutomationAccount{},
        &AutomationAccountList{},
    
        &CosmosdbTable{},
        &CosmosdbTableList{},
    
        &LbRule{},
        &LbRuleList{},
    
        &EventhubAuthorizationRule{},
        &EventhubAuthorizationRuleList{},
    
        &IotDps{},
        &IotDpsList{},
    
        &PostgresqlFirewallRule{},
        &PostgresqlFirewallRuleList{},
    
        &VirtualNetworkGateway{},
        &VirtualNetworkGatewayList{},
    
        &DnsSrvRecord{},
        &DnsSrvRecordList{},
    
        &FirewallApplicationRuleCollection{},
        &FirewallApplicationRuleCollectionList{},
    
        &StorageShare{},
        &StorageShareList{},
    
        &EventgridTopic{},
        &EventgridTopicList{},
    
        &FirewallNATRuleCollection{},
        &FirewallNATRuleCollectionList{},
    
        &MysqlServer{},
        &MysqlServerList{},
    
        &IothubSharedAccessPolicy{},
        &IothubSharedAccessPolicyList{},
    
        &NetworkWatcher{},
        &NetworkWatcherList{},
    
        &PostgresqlConfiguration{},
        &PostgresqlConfigurationList{},
    
        &RoleDefinition{},
        &RoleDefinitionList{},
    
        &SearchService{},
        &SearchServiceList{},
    
        &BatchCertificate{},
        &BatchCertificateList{},
    
        &DevspaceController{},
        &DevspaceControllerList{},
    
        &EventhubConsumerGroup{},
        &EventhubConsumerGroupList{},
    
        &MonitorAutoscaleSetting{},
        &MonitorAutoscaleSettingList{},
    
        &BatchAccount{},
        &BatchAccountList{},
    
        &DataFactory{},
        &DataFactoryList{},
    
        &DataFactoryLinkedServicePostgresql{},
        &DataFactoryLinkedServicePostgresqlList{},
    
        &ManagementGroup{},
        &ManagementGroupList{},
    
        &NetworkInterfaceBackendAddressPoolAssociation{},
        &NetworkInterfaceBackendAddressPoolAssociationList{},
    
        &StreamAnalyticsOutputEventhub{},
        &StreamAnalyticsOutputEventhubList{},
    
        &VirtualNetworkPeering{},
        &VirtualNetworkPeeringList{},
    
        &ApiManagementAuthorizationServer{},
        &ApiManagementAuthorizationServerList{},
    
        &AppServicePlan{},
        &AppServicePlanList{},
    
        &LogicAppTriggerRecurrence{},
        &LogicAppTriggerRecurrenceList{},
    
        &DataFactoryDatasetPostgresql{},
        &DataFactoryDatasetPostgresqlList{},
    
        &PacketCapture{},
        &PacketCaptureList{},
    
        &SqlElasticpool{},
        &SqlElasticpoolList{},
    
        &SqlServer{},
        &SqlServerList{},
    
        &ContainerService{},
        &ContainerServiceList{},
    
        &DevTestWindowsVirtualMachine{},
        &DevTestWindowsVirtualMachineList{},
    
        &NetworkInterfaceApplicationSecurityGroupAssociation{},
        &NetworkInterfaceApplicationSecurityGroupAssociationList{},
    
        &MetricAlertrule{},
        &MetricAlertruleList{},
    
        &ServiceFabricCluster{},
        &ServiceFabricClusterList{},
    
        &Subnet{},
        &SubnetList{},
    
        &DataFactoryLinkedServiceDataLakeStorageGen2{},
        &DataFactoryLinkedServiceDataLakeStorageGen2List{},
    
        &KeyVaultSecret{},
        &KeyVaultSecretList{},
    
        &LogicAppActionCustom{},
        &LogicAppActionCustomList{},
    
        &LbOutboundRule{},
        &LbOutboundRuleList{},
    
        &MariadbDatabase{},
        &MariadbDatabaseList{},
    
        &MonitorLogProfile{},
        &MonitorLogProfileList{},
    
        &ApplicationInsights{},
        &ApplicationInsightsList{},
    
        &Image{},
        &ImageList{},
    
        &SubnetRouteTableAssociation{},
        &SubnetRouteTableAssociationList{},
    
        &ApiManagementAPIOperationPolicy{},
        &ApiManagementAPIOperationPolicyList{},
    
        &AutoscaleSetting{},
        &AutoscaleSettingList{},
    
        &NetworkInterface{},
        &NetworkInterfaceList{},
    
        &NotificationHubNamespace_{},
        &NotificationHubNamespace_List{},
    
        &StreamAnalyticsJob{},
        &StreamAnalyticsJobList{},
    
        &ApiManagementProductAPI{},
        &ApiManagementProductAPIList{},
    
        &DevTestLinuxVirtualMachine{},
        &DevTestLinuxVirtualMachineList{},
    
        &HdinsightHadoopCluster{},
        &HdinsightHadoopClusterList{},
    
        &MysqlConfiguration{},
        &MysqlConfigurationList{},
    
        &PolicyAssignment{},
        &PolicyAssignmentList{},
    
        &StreamAnalyticsOutputMssql{},
        &StreamAnalyticsOutputMssqlList{},
    
        &TrafficManagerEndpoint{},
        &TrafficManagerEndpointList{},
    
        &ApplicationGateway{},
        &ApplicationGatewayList{},
    
        &KeyVaultKey{},
        &KeyVaultKeyList{},
    
        &LogicAppWorkflow{},
        &LogicAppWorkflowList{},
    
        &AzureadServicePrincipal{},
        &AzureadServicePrincipalList{},
    
        &LbProbe{},
        &LbProbeList{},
    
        &LogAnalyticsWorkspaceLinkedService{},
        &LogAnalyticsWorkspaceLinkedServiceList{},
    
        &MysqlDatabase{},
        &MysqlDatabaseList{},
    
        &ServicebusSubscriptionRule{},
        &ServicebusSubscriptionRuleList{},
    
        &SqlFirewallRule{},
        &SqlFirewallRuleList{},
    
        &StorageContainer{},
        &StorageContainerList{},
    
        &ApiManagementProductPolicy{},
        &ApiManagementProductPolicyList{},
    
        &AppServiceCustomHostnameBinding{},
        &AppServiceCustomHostnameBindingList{},
    
        &CosmosdbSQLDatabase{},
        &CosmosdbSQLDatabaseList{},
    
        &ExpressRouteCircuitAuthorization{},
        &ExpressRouteCircuitAuthorizationList{},
    
        &HdinsightHbaseCluster{},
        &HdinsightHbaseClusterList{},
    
        &ApplicationSecurityGroup{},
        &ApplicationSecurityGroupList{},
    
        &AutomationDscNodeconfiguration{},
        &AutomationDscNodeconfigurationList{},
    
        &DnsCaaRecord{},
        &DnsCaaRecordList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
