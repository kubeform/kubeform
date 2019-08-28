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
    
        &ApiManagementGroupUser{},
        &ApiManagementGroupUserList{},
    
        &HdinsightHbaseCluster{},
        &HdinsightHbaseClusterList{},
    
        &KeyVaultAccessPolicy{},
        &KeyVaultAccessPolicyList{},
    
        &RelayNamespace{},
        &RelayNamespaceList{},
    
        &StreamAnalyticsStreamInputBlob{},
        &StreamAnalyticsStreamInputBlobList{},
    
        &Subnet{},
        &SubnetList{},
    
        &VirtualMachineScaleSet{},
        &VirtualMachineScaleSetList{},
    
        &DnsAaaaRecord{},
        &DnsAaaaRecordList{},
    
        &MonitorAutoscaleSetting{},
        &MonitorAutoscaleSettingList{},
    
        &ServicebusNamespaceAuthorizationRule{},
        &ServicebusNamespaceAuthorizationRuleList{},
    
        &UserAssignedIdentity{},
        &UserAssignedIdentityList{},
    
        &ContainerRegistry{},
        &ContainerRegistryList{},
    
        &MetricAlertrule{},
        &MetricAlertruleList{},
    
        &StreamAnalyticsOutputMssql{},
        &StreamAnalyticsOutputMssqlList{},
    
        &Snapshot{},
        &SnapshotList{},
    
        &AutomationVariableDatetime{},
        &AutomationVariableDatetimeList{},
    
        &LogAnalyticsWorkspaceLinkedService{},
        &LogAnalyticsWorkspaceLinkedServiceList{},
    
        &MysqlConfiguration{},
        &MysqlConfigurationList{},
    
        &NotificationHub{},
        &NotificationHubList{},
    
        &MediaServicesAccount{},
        &MediaServicesAccountList{},
    
        &MysqlDatabase{},
        &MysqlDatabaseList{},
    
        &Route{},
        &RouteList{},
    
        &ApiManagementProductGroup{},
        &ApiManagementProductGroupList{},
    
        &FirewallNATRuleCollection{},
        &FirewallNATRuleCollectionList{},
    
        &KeyVaultCertificate{},
        &KeyVaultCertificateList{},
    
        &MonitorLogProfile{},
        &MonitorLogProfileList{},
    
        &StorageShareDirectory{},
        &StorageShareDirectoryList{},
    
        &CosmosdbMongoCollection{},
        &CosmosdbMongoCollectionList{},
    
        &DevTestLab{},
        &DevTestLabList{},
    
        &EventhubNamespaceAuthorizationRule{},
        &EventhubNamespaceAuthorizationRuleList{},
    
        &LbProbe{},
        &LbProbeList{},
    
        &MonitorActivityLogAlert{},
        &MonitorActivityLogAlertList{},
    
        &VirtualMachine{},
        &VirtualMachineList{},
    
        &IotDpsCertificate{},
        &IotDpsCertificateList{},
    
        &LogicAppWorkflow{},
        &LogicAppWorkflowList{},
    
        &ConnectionMonitor{},
        &ConnectionMonitorList{},
    
        &DevTestWindowsVirtualMachine{},
        &DevTestWindowsVirtualMachineList{},
    
        &EventgridEventSubscription{},
        &EventgridEventSubscriptionList{},
    
        &KeyVaultSecret{},
        &KeyVaultSecretList{},
    
        &LogicAppActionHTTP{},
        &LogicAppActionHTTPList{},
    
        &PostgresqlVirtualNetworkRule{},
        &PostgresqlVirtualNetworkRuleList{},
    
        &DevTestLinuxVirtualMachine{},
        &DevTestLinuxVirtualMachineList{},
    
        &MysqlVirtualNetworkRule{},
        &MysqlVirtualNetworkRuleList{},
    
        &EventhubConsumerGroup{},
        &EventhubConsumerGroupList{},
    
        &HdinsightHadoopCluster{},
        &HdinsightHadoopClusterList{},
    
        &HdinsightKafkaCluster{},
        &HdinsightKafkaClusterList{},
    
        &PolicyDefinition{},
        &PolicyDefinitionList{},
    
        &ServicebusTopicAuthorizationRule{},
        &ServicebusTopicAuthorizationRuleList{},
    
        &SqlServer{},
        &SqlServerList{},
    
        &Iothub{},
        &IothubList{},
    
        &MapsAccount{},
        &MapsAccountList{},
    
        &AutomationRunbook{},
        &AutomationRunbookList{},
    
        &AzureadServicePrincipal{},
        &AzureadServicePrincipalList{},
    
        &BatchCertificate{},
        &BatchCertificateList{},
    
        &FirewallNetworkRuleCollection{},
        &FirewallNetworkRuleCollectionList{},
    
        &Lb{},
        &LbList{},
    
        &StreamAnalyticsOutputServicebusQueue{},
        &StreamAnalyticsOutputServicebusQueueList{},
    
        &AnalysisServicesServer{},
        &AnalysisServicesServerList{},
    
        &AutomationDscNodeconfiguration{},
        &AutomationDscNodeconfigurationList{},
    
        &DevTestPolicy{},
        &DevTestPolicyList{},
    
        &LbNATRule{},
        &LbNATRuleList{},
    
        &MysqlServer{},
        &MysqlServerList{},
    
        &PostgresqlConfiguration{},
        &PostgresqlConfigurationList{},
    
        &RoleAssignment{},
        &RoleAssignmentList{},
    
        &AzureadServicePrincipalPassword{},
        &AzureadServicePrincipalPasswordList{},
    
        &ApplicationSecurityGroup{},
        &ApplicationSecurityGroupList{},
    
        &ContainerGroup{},
        &ContainerGroupList{},
    
        &DataFactoryPipeline{},
        &DataFactoryPipelineList{},
    
        &RecoveryServicesVault{},
        &RecoveryServicesVaultList{},
    
        &RedisFirewallRule{},
        &RedisFirewallRuleList{},
    
        &ApiManagementProduct{},
        &ApiManagementProductList{},
    
        &AppServiceSlot{},
        &AppServiceSlotList{},
    
        &AutomationVariableBool{},
        &AutomationVariableBoolList{},
    
        &IothubConsumerGroup{},
        &IothubConsumerGroupList{},
    
        &MonitorDiagnosticSetting{},
        &MonitorDiagnosticSettingList{},
    
        &PostgresqlFirewallRule{},
        &PostgresqlFirewallRuleList{},
    
        &SchedulerJobCollection{},
        &SchedulerJobCollectionList{},
    
        &BatchAccount{},
        &BatchAccountList{},
    
        &CosmosdbMongoDatabase{},
        &CosmosdbMongoDatabaseList{},
    
        &MariadbDatabase{},
        &MariadbDatabaseList{},
    
        &MonitorActionGroup{},
        &MonitorActionGroupList{},
    
        &NotificationHubNamespace_{},
        &NotificationHubNamespace_List{},
    
        &VirtualNetworkGateway{},
        &VirtualNetworkGatewayList{},
    
        &DataLakeStore{},
        &DataLakeStoreList{},
    
        &ExpressRouteCircuitAuthorization{},
        &ExpressRouteCircuitAuthorizationList{},
    
        &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
        &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},
    
        &StorageBlob{},
        &StorageBlobList{},
    
        &TrafficManagerProfile{},
        &TrafficManagerProfileList{},
    
        &ApiManagement{},
        &ApiManagementList{},
    
        &ApiManagementProductAPI{},
        &ApiManagementProductAPIList{},
    
        &DnsCaaRecord{},
        &DnsCaaRecordList{},
    
        &ExpressRouteCircuitPeering{},
        &ExpressRouteCircuitPeeringList{},
    
        &SqlElasticpool{},
        &SqlElasticpoolList{},
    
        &ApiManagementAPIPolicy{},
        &ApiManagementAPIPolicyList{},
    
        &ApplicationInsights{},
        &ApplicationInsightsList{},
    
        &AutomationDscConfiguration{},
        &AutomationDscConfigurationList{},
    
        &DataFactoryDatasetMysql{},
        &DataFactoryDatasetMysqlList{},
    
        &DdosProtectionPlan{},
        &DdosProtectionPlanList{},
    
        &NotificationHubAuthorizationRule{},
        &NotificationHubAuthorizationRuleList{},
    
        &SqlDatabase{},
        &SqlDatabaseList{},
    
        &StreamAnalyticsFunctionJavascriptUdf{},
        &StreamAnalyticsFunctionJavascriptUdfList{},
    
        &EventhubAuthorizationRule{},
        &EventhubAuthorizationRuleList{},
    
        &HdinsightSparkCluster{},
        &HdinsightSparkClusterList{},
    
        &BatchApplication{},
        &BatchApplicationList{},
    
        &DnsTxtRecord{},
        &DnsTxtRecordList{},
    
        &Image{},
        &ImageList{},
    
        &ManagementGroup{},
        &ManagementGroupList{},
    
        &StreamAnalyticsStreamInputIothub{},
        &StreamAnalyticsStreamInputIothubList{},
    
        &TemplateDeployment{},
        &TemplateDeploymentList{},
    
        &NetworkProfile{},
        &NetworkProfileList{},
    
        &ResourceGroup{},
        &ResourceGroupList{},
    
        &ApiManagementAPIOperation{},
        &ApiManagementAPIOperationList{},
    
        &Eventhub{},
        &EventhubList{},
    
        &ExpressRouteCircuit{},
        &ExpressRouteCircuitList{},
    
        &NetworkConnectionMonitor{},
        &NetworkConnectionMonitorList{},
    
        &SecurityCenterWorkspace{},
        &SecurityCenterWorkspaceList{},
    
        &StorageAccount{},
        &StorageAccountList{},
    
        &AppServiceCustomHostnameBinding{},
        &AppServiceCustomHostnameBindingList{},
    
        &HdinsightMlServicesCluster{},
        &HdinsightMlServicesClusterList{},
    
        &KeyVault{},
        &KeyVaultList{},
    
        &AutomationModule{},
        &AutomationModuleList{},
    
        &DataFactoryLinkedServiceMysql{},
        &DataFactoryLinkedServiceMysqlList{},
    
        &EventgridTopic{},
        &EventgridTopicList{},
    
        &SqlActiveDirectoryAdministrator{},
        &SqlActiveDirectoryAdministratorList{},
    
        &StorageShare{},
        &StorageShareList{},
    
        &DataFactory{},
        &DataFactoryList{},
    
        &DnsZone{},
        &DnsZoneList{},
    
        &EventgridDomain{},
        &EventgridDomainList{},
    
        &NetworkSecurityGroup{},
        &NetworkSecurityGroupList{},
    
        &SqlVirtualNetworkRule{},
        &SqlVirtualNetworkRuleList{},
    
        &StorageContainer{},
        &StorageContainerList{},
    
        &ApiManagementSubscription{},
        &ApiManagementSubscriptionList{},
    
        &CognitiveAccount{},
        &CognitiveAccountList{},
    
        &ServicebusSubscriptionRule{},
        &ServicebusSubscriptionRuleList{},
    
        &StreamAnalyticsOutputBlob{},
        &StreamAnalyticsOutputBlobList{},
    
        &ApiManagementAPIVersionSet{},
        &ApiManagementAPIVersionSetList{},
    
        &ApiManagementAuthorizationServer{},
        &ApiManagementAuthorizationServerList{},
    
        &AutomationVariableString{},
        &AutomationVariableStringList{},
    
        &LbRule{},
        &LbRuleList{},
    
        &MariadbFirewallRule{},
        &MariadbFirewallRuleList{},
    
        &PacketCapture{},
        &PacketCaptureList{},
    
        &SecurityCenterSubscriptionPricing{},
        &SecurityCenterSubscriptionPricingList{},
    
        &DataFactoryLinkedServiceDataLakeStorageGen2{},
        &DataFactoryLinkedServiceDataLakeStorageGen2List{},
    
        &DataLakeStoreFirewallRule{},
        &DataLakeStoreFirewallRuleList{},
    
        &LbBackendAddressPool{},
        &LbBackendAddressPoolList{},
    
        &LogicAppTriggerHTTPRequest{},
        &LogicAppTriggerHTTPRequestList{},
    
        &LogicAppTriggerRecurrence{},
        &LogicAppTriggerRecurrenceList{},
    
        &ServicebusTopic{},
        &ServicebusTopicList{},
    
        &DataFactoryDatasetPostgresql{},
        &DataFactoryDatasetPostgresqlList{},
    
        &DnsCnameRecord{},
        &DnsCnameRecordList{},
    
        &Firewall{},
        &FirewallList{},
    
        &LogAnalyticsWorkspace{},
        &LogAnalyticsWorkspaceList{},
    
        &DataFactoryDatasetSQLServerTable{},
        &DataFactoryDatasetSQLServerTableList{},
    
        &LogicAppTriggerCustom{},
        &LogicAppTriggerCustomList{},
    
        &PrivateDNSARecord{},
        &PrivateDNSARecordList{},
    
        &SecurityCenterContact{},
        &SecurityCenterContactList{},
    
        &StreamAnalyticsOutputEventhub{},
        &StreamAnalyticsOutputEventhubList{},
    
        &IothubSharedAccessPolicy{},
        &IothubSharedAccessPolicyList{},
    
        &LocalNetworkGateway{},
        &LocalNetworkGatewayList{},
    
        &NetworkWatcher{},
        &NetworkWatcherList{},
    
        &PostgresqlDatabase{},
        &PostgresqlDatabaseList{},
    
        &PrivateDNSZone{},
        &PrivateDNSZoneList{},
    
        &RecoveryServicesProtectionPolicyVm{},
        &RecoveryServicesProtectionPolicyVmList{},
    
        &SharedImage{},
        &SharedImageList{},
    
        &AppServiceActiveSlot{},
        &AppServiceActiveSlotList{},
    
        &ApiManagementBackend{},
        &ApiManagementBackendList{},
    
        &CosmosdbTable{},
        &CosmosdbTableList{},
    
        &SharedImageGallery{},
        &SharedImageGalleryList{},
    
        &BatchPool{},
        &BatchPoolList{},
    
        &CdnProfile{},
        &CdnProfileList{},
    
        &LogicAppActionCustom{},
        &LogicAppActionCustomList{},
    
        &VirtualNetworkPeering{},
        &VirtualNetworkPeeringList{},
    
        &ApplicationGateway{},
        &ApplicationGatewayList{},
    
        &CosmosdbCassandraKeyspace{},
        &CosmosdbCassandraKeyspaceList{},
    
        &DataLakeAnalyticsFirewallRule{},
        &DataLakeAnalyticsFirewallRuleList{},
    
        &DnsNsRecord{},
        &DnsNsRecordList{},
    
        &DnsPtrRecord{},
        &DnsPtrRecordList{},
    
        &IotDps{},
        &IotDpsList{},
    
        &ServicebusQueueAuthorizationRule{},
        &ServicebusQueueAuthorizationRuleList{},
    
        &ApiManagementProductPolicy{},
        &ApiManagementProductPolicyList{},
    
        &ApplicationInsightsWebTest{},
        &ApplicationInsightsWebTestList{},
    
        &DnsARecord{},
        &DnsARecordList{},
    
        &StorageQueue{},
        &StorageQueueList{},
    
        &DevspaceController{},
        &DevspaceControllerList{},
    
        &SchedulerJob{},
        &SchedulerJobList{},
    
        &AzureadApplication{},
        &AzureadApplicationList{},
    
        &DatabricksWorkspace{},
        &DatabricksWorkspaceList{},
    
        &ManagementLock{},
        &ManagementLockList{},
    
        &AvailabilitySet{},
        &AvailabilitySetList{},
    
        &CdnEndpoint{},
        &CdnEndpointList{},
    
        &PolicyAssignment{},
        &PolicyAssignmentList{},
    
        &ServicebusQueue{},
        &ServicebusQueueList{},
    
        &DataLakeAnalyticsAccount{},
        &DataLakeAnalyticsAccountList{},
    
        &ServicebusSubscription{},
        &ServicebusSubscriptionList{},
    
        &ApiManagementAPI{},
        &ApiManagementAPIList{},
    
        &ApiManagementAPISchema{},
        &ApiManagementAPISchemaList{},
    
        &AutomationCredential{},
        &AutomationCredentialList{},
    
        &ManagedDisk{},
        &ManagedDiskList{},
    
        &NetworkSecurityRule{},
        &NetworkSecurityRuleList{},
    
        &SearchService{},
        &SearchServiceList{},
    
        &StreamAnalyticsStreamInputEventhub{},
        &StreamAnalyticsStreamInputEventhubList{},
    
        &ApiManagementProperty{},
        &ApiManagementPropertyList{},
    
        &AutoscaleSetting{},
        &AutoscaleSettingList{},
    
        &ContainerService{},
        &ContainerServiceList{},
    
        &RedisCache{},
        &RedisCacheList{},
    
        &StreamAnalyticsJob{},
        &StreamAnalyticsJobList{},
    
        &VirtualNetwork{},
        &VirtualNetworkList{},
    
        &FunctionApp{},
        &FunctionAppList{},
    
        &NetworkInterfaceNATRuleAssociation{},
        &NetworkInterfaceNATRuleAssociationList{},
    
        &PolicySetDefinition{},
        &PolicySetDefinitionList{},
    
        &PostgresqlServer{},
        &PostgresqlServerList{},
    
        &DataFactoryLinkedServiceSQLServer{},
        &DataFactoryLinkedServiceSQLServerList{},
    
        &NetworkPacketCapture{},
        &NetworkPacketCaptureList{},
    
        &RecoveryServicesProtectedVm{},
        &RecoveryServicesProtectedVmList{},
    
        &AppService{},
        &AppServiceList{},
    
        &AutomationSchedule{},
        &AutomationScheduleList{},
    
        &HdinsightStormCluster{},
        &HdinsightStormClusterList{},
    
        &KubernetesCluster{},
        &KubernetesClusterList{},
    
        &ServicebusNamespace{},
        &ServicebusNamespaceList{},
    
        &StorageTable{},
        &StorageTableList{},
    
        &HdinsightInteractiveQueryCluster{},
        &HdinsightInteractiveQueryClusterList{},
    
        &NetworkInterfaceApplicationSecurityGroupAssociation{},
        &NetworkInterfaceApplicationSecurityGroupAssociationList{},
    
        &PublicIP{},
        &PublicIPList{},
    
        &SignalrService{},
        &SignalrServiceList{},
    
        &ApiManagementCertificate{},
        &ApiManagementCertificateList{},
    
        &AutomationVariableInt{},
        &AutomationVariableIntList{},
    
        &CosmosdbAccount{},
        &CosmosdbAccountList{},
    
        &DataLakeStoreFile{},
        &DataLakeStoreFileList{},
    
        &MysqlFirewallRule{},
        &MysqlFirewallRuleList{},
    
        &DnsMxRecord{},
        &DnsMxRecordList{},
    
        &FirewallApplicationRuleCollection{},
        &FirewallApplicationRuleCollectionList{},
    
        &TrafficManagerEndpoint{},
        &TrafficManagerEndpointList{},
    
        &VirtualMachineDataDiskAttachment{},
        &VirtualMachineDataDiskAttachmentList{},
    
        &SharedImageVersion{},
        &SharedImageVersionList{},
    
        &SubnetNetworkSecurityGroupAssociation{},
        &SubnetNetworkSecurityGroupAssociationList{},
    
        &ApiManagementLogger{},
        &ApiManagementLoggerList{},
    
        &CosmosdbSQLDatabase{},
        &CosmosdbSQLDatabaseList{},
    
        &DnsSrvRecord{},
        &DnsSrvRecordList{},
    
        &HdinsightRserverCluster{},
        &HdinsightRserverClusterList{},
    
        &MonitorMetricAlertrule{},
        &MonitorMetricAlertruleList{},
    
        &ServiceFabricCluster{},
        &ServiceFabricClusterList{},
    
        &NetworkInterfaceBackendAddressPoolAssociation{},
        &NetworkInterfaceBackendAddressPoolAssociationList{},
    
        &RouteTable{},
        &RouteTableList{},
    
        &AppServicePlan{},
        &AppServicePlanList{},
    
        &LogAnalyticsLinkedService{},
        &LogAnalyticsLinkedServiceList{},
    
        &NetworkInterface{},
        &NetworkInterfaceList{},
    
        &SqlFirewallRule{},
        &SqlFirewallRuleList{},
    
        &DataFactoryLinkedServicePostgresql{},
        &DataFactoryLinkedServicePostgresqlList{},
    
        &KeyVaultKey{},
        &KeyVaultKeyList{},
    
        &LbOutboundRule{},
        &LbOutboundRuleList{},
    
        &LogAnalyticsSolution{},
        &LogAnalyticsSolutionList{},
    
        &VirtualMachineExtension{},
        &VirtualMachineExtensionList{},
    
        &ApplicationInsightsAPIKey{},
        &ApplicationInsightsAPIKeyList{},
    
        &ApiManagementUser{},
        &ApiManagementUserList{},
    
        &AutomationAccount{},
        &AutomationAccountList{},
    
        &DevTestVirtualNetwork{},
        &DevTestVirtualNetworkList{},
    
        &LbNATPool{},
        &LbNATPoolList{},
    
        &PublicIPPrefix{},
        &PublicIPPrefixList{},
    
        &ApiManagementOpenidConnectProvider{},
        &ApiManagementOpenidConnectProviderList{},
    
        &NetworkDdosProtectionPlan{},
        &NetworkDdosProtectionPlanList{},
    
        &VirtualNetworkGatewayConnection{},
        &VirtualNetworkGatewayConnectionList{},
    
        &ApiManagementAPIOperationPolicy{},
        &ApiManagementAPIOperationPolicyList{},
    
        &MariadbServer{},
        &MariadbServerList{},
    
        &MssqlElasticpool{},
        &MssqlElasticpoolList{},
    
        &RoleDefinition{},
        &RoleDefinitionList{},
    
        &ApiManagementGroup{},
        &ApiManagementGroupList{},
    
        &EventhubNamespace_{},
        &EventhubNamespace_List{},
    
        &MonitorMetricAlert{},
        &MonitorMetricAlertList{},
    
        &StorageTableEntity{},
        &StorageTableEntityList{},
    
        &SubnetRouteTableAssociation{},
        &SubnetRouteTableAssociationList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
