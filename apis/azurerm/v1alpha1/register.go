/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	"kubeform.dev/kubeform/apis/azurerm"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppService{},
		&AppServiceList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&BatchAccount{},
		&BatchAccountList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&BatchPool{},
		&BatchPoolList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&ContainerService{},
		&ContainerServiceList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataFactory{},
		&DataFactoryList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&DnsZone{},
		&DnsZoneList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&Eventhub{},
		&EventhubList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&Firewall{},
		&FirewallList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&FunctionApp{},
		&FunctionAppList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&Image{},
		&ImageList{},

		&IotDps{},
		&IotDpsList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&Iothub{},
		&IothubList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&KeyVault{},
		&KeyVaultList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&Lb{},
		&LbList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LbProbe{},
		&LbProbeList{},

		&LbRule{},
		&LbRuleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MapsAccount{},
		&MapsAccountList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&NotificationHub{},
		&NotificationHubList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&PacketCapture{},
		&PacketCaptureList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&PublicIP{},
		&PublicIPList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RedisCache{},
		&RedisCacheList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&Route{},
		&RouteList{},

		&RouteTable{},
		&RouteTableList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SearchService{},
		&SearchServiceList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&SharedImage{},
		&SharedImageList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&Snapshot{},
		&SnapshotList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageContainer{},
		&StorageContainerList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StorageShare{},
		&StorageShareList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&StorageTable{},
		&StorageTableList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&Subnet{},
		&SubnetList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
