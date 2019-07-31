package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/aws"
)

var SchemeGroupVersion = schema.GroupVersion{Group: aws.GroupName, Version: "v1alpha1"}

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

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&GlueJob{},
		&GlueJobList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&IamUser{},
		&IamUserList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&Ami{},
		&AmiList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SqsQueue{},
		&SqsQueueList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&VpnGateway{},
		&VpnGatewayList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&LbListener{},
		&LbListenerList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&Route53Record{},
		&Route53RecordList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&Subnet{},
		&SubnetList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&IotPolicy{},
		&IotPolicyList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&EipAssociation{},
		&EipAssociationList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&IamPolicy{},
		&IamPolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&Eip{},
		&EipList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RdsCluster{},
		&RdsClusterList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Elb{},
		&ElbList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&DxConnection{},
		&DxConnectionList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&DaxCluster{},
		&DaxClusterList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&MskCluster{},
		&MskClusterList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafRule{},
		&WafRuleList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&KmsAlias{},
		&KmsAliasList{},

		&TransferUser{},
		&TransferUserList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&EmrCluster{},
		&EmrClusterList{},

		&SfnActivity{},
		&SfnActivityList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&S3Bucket{},
		&S3BucketList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&WafWebACL{},
		&WafWebACLList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&SsmActivation{},
		&SsmActivationList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&PinpointApp{},
		&PinpointAppList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&Lb{},
		&LbList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IamGroup{},
		&IamGroupList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&Vpc{},
		&VpcList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&Route{},
		&RouteList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&Alb{},
		&AlbList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&KmsKey{},
		&KmsKeyList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&BackupVault{},
		&BackupVaultList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DbInstance{},
		&DbInstanceList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&SesTemplate{},
		&SesTemplateList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&FlowLog{},
		&FlowLogList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&IotThing{},
		&IotThingList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SsmParameter{},
		&SsmParameterList{},

		&SnsTopic{},
		&SnsTopicList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IotCertificate{},
		&IotCertificateList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&NatGateway{},
		&NatGatewayList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&TransferServer{},
		&TransferServerList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&Codepipeline{},
		&CodepipelineList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&EksCluster{},
		&EksClusterList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&NetworkACL{},
		&NetworkACLList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&AlbListener{},
		&AlbListenerList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&IamRole{},
		&IamRoleList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&KeyPair{},
		&KeyPairList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DxGateway{},
		&DxGatewayList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&RouteTable{},
		&RouteTableList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&Instance{},
		&InstanceList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&EcsService{},
		&EcsServiceList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&DxLag{},
		&DxLagList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&IotThingType{},
		&IotThingTypeList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
