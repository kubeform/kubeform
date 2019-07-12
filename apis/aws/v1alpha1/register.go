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

		&AwsCloudwatchEventRule{},
		&AwsCloudwatchEventRuleList{},

		&AwsWafregionalRegexMatchSet{},
		&AwsWafregionalRegexMatchSetList{},

		&AwsDbSecurityGroup{},
		&AwsDbSecurityGroupList{},

		&AwsBackupSelection{},
		&AwsBackupSelectionList{},

		&AwsNetworkInterfaceSgAttachment{},
		&AwsNetworkInterfaceSgAttachmentList{},

		&AwsConfigConfigurationAggregator{},
		&AwsConfigConfigurationAggregatorList{},

		&AwsOpsworksGangliaLayer{},
		&AwsOpsworksGangliaLayerList{},

		&AwsS3BucketMetric{},
		&AwsS3BucketMetricList{},

		&AwsIamGroup{},
		&AwsIamGroupList{},

		&AwsDefaultNetworkAcl{},
		&AwsDefaultNetworkAclList{},

		&AwsEc2TransitGatewayRouteTablePropagation{},
		&AwsEc2TransitGatewayRouteTablePropagationList{},

		&AwsCognitoUserPoolClient{},
		&AwsCognitoUserPoolClientList{},

		&AwsEbsSnapshot{},
		&AwsEbsSnapshotList{},

		&AwsKeyPair{},
		&AwsKeyPairList{},

		&AwsVpcEndpointRouteTableAssociation{},
		&AwsVpcEndpointRouteTableAssociationList{},

		&AwsWafregionalWebAclAssociation{},
		&AwsWafregionalWebAclAssociationList{},

		&AwsBatchJobDefinition{},
		&AwsBatchJobDefinitionList{},

		&AwsDocdbCluster{},
		&AwsDocdbClusterList{},

		&AwsDocdbClusterInstance{},
		&AwsDocdbClusterInstanceList{},

		&AwsElastictranscoderPreset{},
		&AwsElastictranscoderPresetList{},

		&AwsRdsCluster{},
		&AwsRdsClusterList{},

		&AwsStoragegatewayGateway{},
		&AwsStoragegatewayGatewayList{},

		&AwsNeptuneEventSubscription{},
		&AwsNeptuneEventSubscriptionList{},

		&AwsAcmpcaCertificateAuthority{},
		&AwsAcmpcaCertificateAuthorityList{},

		&AwsGlueJob{},
		&AwsGlueJobList{},

		&AwsLbCookieStickinessPolicy{},
		&AwsLbCookieStickinessPolicyList{},

		&AwsVpcEndpointServiceAllowedPrincipal{},
		&AwsVpcEndpointServiceAllowedPrincipalList{},

		&AwsPinpointApnsChannel{},
		&AwsPinpointApnsChannelList{},

		&AwsCodedeployDeploymentConfig{},
		&AwsCodedeployDeploymentConfigList{},

		&AwsDirectoryServiceConditionalForwarder{},
		&AwsDirectoryServiceConditionalForwarderList{},

		&AwsElb{},
		&AwsElbList{},

		&AwsRouteTableAssociation{},
		&AwsRouteTableAssociationList{},

		&AwsDefaultVpcDhcpOptions{},
		&AwsDefaultVpcDhcpOptionsList{},

		&AwsVpcPeeringConnection{},
		&AwsVpcPeeringConnectionList{},

		&AwsDbParameterGroup{},
		&AwsDbParameterGroupList{},

		&AwsSesActiveReceiptRuleSet{},
		&AwsSesActiveReceiptRuleSetList{},

		&AwsDefaultSubnet{},
		&AwsDefaultSubnetList{},

		&AwsCloudwatchLogResourcePolicy{},
		&AwsCloudwatchLogResourcePolicyList{},

		&AwsIotThingType{},
		&AwsIotThingTypeList{},

		&AwsS3BucketInventory{},
		&AwsS3BucketInventoryList{},

		&AwsWafregionalGeoMatchSet{},
		&AwsWafregionalGeoMatchSetList{},

		&AwsCodecommitTrigger{},
		&AwsCodecommitTriggerList{},

		&AwsCurReportDefinition{},
		&AwsCurReportDefinitionList{},

		&AwsNatGateway{},
		&AwsNatGatewayList{},

		&AwsSnsTopicSubscription{},
		&AwsSnsTopicSubscriptionList{},

		&AwsVpnConnectionRoute{},
		&AwsVpnConnectionRouteList{},

		&AwsApiGatewayRestApi{},
		&AwsApiGatewayRestApiList{},

		&AwsDynamodbTable{},
		&AwsDynamodbTableList{},

		&AwsNeptuneSubnetGroup{},
		&AwsNeptuneSubnetGroupList{},

		&AwsSsmPatchBaseline{},
		&AwsSsmPatchBaselineList{},

		&AwsSnsSmsPreferences{},
		&AwsSnsSmsPreferencesList{},

		&AwsWafWebAcl{},
		&AwsWafWebAclList{},

		&AwsElasticsearchDomain{},
		&AwsElasticsearchDomainList{},

		&AwsWafregionalRateBasedRule{},
		&AwsWafregionalRateBasedRuleList{},

		&AwsAutoscalingSchedule{},
		&AwsAutoscalingScheduleList{},

		&AwsElastictranscoderPipeline{},
		&AwsElastictranscoderPipelineList{},

		&AwsEmrSecurityConfiguration{},
		&AwsEmrSecurityConfigurationList{},

		&AwsLoadBalancerBackendServerPolicy{},
		&AwsLoadBalancerBackendServerPolicyList{},

		&AwsCognitoResourceServer{},
		&AwsCognitoResourceServerList{},

		&AwsDbInstance{},
		&AwsDbInstanceList{},

		&AwsSagemakerModel{},
		&AwsSagemakerModelList{},

		&AwsIamGroupPolicyAttachment{},
		&AwsIamGroupPolicyAttachmentList{},

		&AwsIotCertificate{},
		&AwsIotCertificateList{},

		&AwsRedshiftCluster{},
		&AwsRedshiftClusterList{},

		&AwsWafregionalIpset{},
		&AwsWafregionalIpsetList{},

		&AwsApiGatewayClientCertificate{},
		&AwsApiGatewayClientCertificateList{},

		&AwsConfigDeliveryChannel{},
		&AwsConfigDeliveryChannelList{},

		&AwsEc2ClientVpnEndpoint{},
		&AwsEc2ClientVpnEndpointList{},

		&AwsGuarddutyIpset{},
		&AwsGuarddutyIpsetList{},

		&AwsRoute53ResolverRule{},
		&AwsRoute53ResolverRuleList{},

		&AwsVpcDhcpOptionsAssociation{},
		&AwsVpcDhcpOptionsAssociationList{},

		&AwsCloudwatchEventTarget{},
		&AwsCloudwatchEventTargetList{},

		&AwsElbAttachment{},
		&AwsElbAttachmentList{},

		&AwsLicensemanagerAssociation{},
		&AwsLicensemanagerAssociationList{},

		&AwsRoute53QueryLog{},
		&AwsRoute53QueryLogList{},

		&AwsSqsQueue{},
		&AwsSqsQueueList{},

		&AwsPinpointApp{},
		&AwsPinpointAppList{},

		&AwsDocdbClusterParameterGroup{},
		&AwsDocdbClusterParameterGroupList{},

		&AwsElasticBeanstalkApplication{},
		&AwsElasticBeanstalkApplicationList{},

		&AwsLoadBalancerPolicy{},
		&AwsLoadBalancerPolicyList{},

		&AwsMqBroker{},
		&AwsMqBrokerList{},

		&AwsSesTemplate{},
		&AwsSesTemplateList{},

		&AwsVpnGateway{},
		&AwsVpnGatewayList{},

		&AwsRdsClusterParameterGroup{},
		&AwsRdsClusterParameterGroupList{},

		&AwsS3Bucket{},
		&AwsS3BucketList{},

		&AwsConfigConfigRule{},
		&AwsConfigConfigRuleList{},

		&AwsCloudhsmV2Hsm{},
		&AwsCloudhsmV2HsmList{},

		&AwsCodepipeline{},
		&AwsCodepipelineList{},

		&AwsCodepipelineWebhook{},
		&AwsCodepipelineWebhookList{},

		&AwsEmrInstanceGroup{},
		&AwsEmrInstanceGroupList{},

		&AwsNetworkAcl{},
		&AwsNetworkAclList{},

		&AwsWafSqlInjectionMatchSet{},
		&AwsWafSqlInjectionMatchSetList{},

		&AwsEcsService{},
		&AwsEcsServiceList{},

		&AwsInternetGateway{},
		&AwsInternetGatewayList{},

		&AwsWorklinkFleet{},
		&AwsWorklinkFleetList{},

		&AwsBudgetsBudget{},
		&AwsBudgetsBudgetList{},

		&AwsDxHostedPublicVirtualInterfaceAccepter{},
		&AwsDxHostedPublicVirtualInterfaceAccepterList{},

		&AwsEcrRepositoryPolicy{},
		&AwsEcrRepositoryPolicyList{},

		&AwsGameliftGameSessionQueue{},
		&AwsGameliftGameSessionQueueList{},

		&AwsDxGatewayAssociationProposal{},
		&AwsDxGatewayAssociationProposalList{},

		&AwsOrganizationsOrganizationalUnit{},
		&AwsOrganizationsOrganizationalUnitList{},

		&AwsRedshiftParameterGroup{},
		&AwsRedshiftParameterGroupList{},

		&AwsServiceDiscoveryPublicDnsNamespace{},
		&AwsServiceDiscoveryPublicDnsNamespaceList{},

		&AwsWafregionalSizeConstraintSet{},
		&AwsWafregionalSizeConstraintSetList{},

		&AwsEksCluster{},
		&AwsEksClusterList{},

		&AwsElasticsearchDomainPolicy{},
		&AwsElasticsearchDomainPolicyList{},

		&AwsSesReceiptRule{},
		&AwsSesReceiptRuleList{},

		&AwsSwfDomain{},
		&AwsSwfDomainList{},

		&AwsPinpointApnsVoipChannel{},
		&AwsPinpointApnsVoipChannelList{},

		&AwsCodedeployApp{},
		&AwsCodedeployAppList{},

		&AwsGlacierVault{},
		&AwsGlacierVaultList{},

		&AwsGlueSecurityConfiguration{},
		&AwsGlueSecurityConfigurationList{},

		&AwsKmsKey{},
		&AwsKmsKeyList{},

		&AwsOpsworksStack{},
		&AwsOpsworksStackList{},

		&AwsOpsworksHaproxyLayer{},
		&AwsOpsworksHaproxyLayerList{},

		&AwsOpsworksRailsAppLayer{},
		&AwsOpsworksRailsAppLayerList{},

		&AwsRamResourceAssociation{},
		&AwsRamResourceAssociationList{},

		&AwsSpotInstanceRequest{},
		&AwsSpotInstanceRequestList{},

		&AwsVpcEndpoint{},
		&AwsVpcEndpointList{},

		&AwsAppmeshVirtualNode{},
		&AwsAppmeshVirtualNodeList{},

		&AwsElasticBeanstalkApplicationVersion{},
		&AwsElasticBeanstalkApplicationVersionList{},

		&AwsIamUserPolicy{},
		&AwsIamUserPolicyList{},

		&AwsLambdaEventSourceMapping{},
		&AwsLambdaEventSourceMappingList{},

		&AwsSesIdentityNotificationTopic{},
		&AwsSesIdentityNotificationTopicList{},

		&AwsS3BucketPublicAccessBlock{},
		&AwsS3BucketPublicAccessBlockList{},

		&AwsAppsyncGraphqlApi{},
		&AwsAppsyncGraphqlApiList{},

		&AwsDbClusterSnapshot{},
		&AwsDbClusterSnapshotList{},

		&AwsEc2TransitGatewayVpcAttachment{},
		&AwsEc2TransitGatewayVpcAttachmentList{},

		&AwsSnsTopicPolicy{},
		&AwsSnsTopicPolicyList{},

		&AwsPinpointSmsChannel{},
		&AwsPinpointSmsChannelList{},

		&AwsIamServerCertificate{},
		&AwsIamServerCertificateList{},

		&AwsSecurityhubProductSubscription{},
		&AwsSecurityhubProductSubscriptionList{},

		&AwsTransferUser{},
		&AwsTransferUserList{},

		&AwsOpsworksMysqlLayer{},
		&AwsOpsworksMysqlLayerList{},

		&AwsPinpointApnsSandboxChannel{},
		&AwsPinpointApnsSandboxChannelList{},

		&AwsAcmCertificateValidation{},
		&AwsAcmCertificateValidationList{},

		&AwsDxConnectionAssociation{},
		&AwsDxConnectionAssociationList{},

		&AwsEc2Fleet{},
		&AwsEc2FleetList{},

		&AwsElasticacheParameterGroup{},
		&AwsElasticacheParameterGroupList{},

		&AwsGameliftBuild{},
		&AwsGameliftBuildList{},

		&AwsIotThing{},
		&AwsIotThingList{},

		&AwsAlbTargetGroupAttachment{},
		&AwsAlbTargetGroupAttachmentList{},

		&AwsApiGatewayRequestValidator{},
		&AwsApiGatewayRequestValidatorList{},

		&AwsDxBgpPeer{},
		&AwsDxBgpPeerList{},

		&AwsVpcPeeringConnectionAccepter{},
		&AwsVpcPeeringConnectionAccepterList{},

		&AwsCloudformationStackSetInstance{},
		&AwsCloudformationStackSetInstanceList{},

		&AwsLicensemanagerLicenseConfiguration{},
		&AwsLicensemanagerLicenseConfigurationList{},

		&AwsSsmMaintenanceWindowTarget{},
		&AwsSsmMaintenanceWindowTargetList{},

		&AwsAutoscalingLifecycleHook{},
		&AwsAutoscalingLifecycleHookList{},

		&AwsCognitoIdentityPoolRolesAttachment{},
		&AwsCognitoIdentityPoolRolesAttachmentList{},

		&AwsPinpointBaiduChannel{},
		&AwsPinpointBaiduChannelList{},

		&AwsApiGatewayMethodSettings{},
		&AwsApiGatewayMethodSettingsList{},

		&AwsCloudwatchLogGroup{},
		&AwsCloudwatchLogGroupList{},

		&AwsDbSnapshot{},
		&AwsDbSnapshotList{},

		&AwsRouteTable{},
		&AwsRouteTableList{},

		&AwsSnsPlatformApplication{},
		&AwsSnsPlatformApplicationList{},

		&AwsWafSizeConstraintSet{},
		&AwsWafSizeConstraintSetList{},

		&AwsDmsCertificate{},
		&AwsDmsCertificateList{},

		&AwsKmsExternalKey{},
		&AwsKmsExternalKeyList{},

		&AwsS3BucketNotification{},
		&AwsS3BucketNotificationList{},

		&AwsWafRateBasedRule{},
		&AwsWafRateBasedRuleList{},

		&AwsCloudtrail{},
		&AwsCloudtrailList{},

		&AwsConfigAggregateAuthorization{},
		&AwsConfigAggregateAuthorizationList{},

		&AwsLoadBalancerListenerPolicy{},
		&AwsLoadBalancerListenerPolicyList{},

		&AwsNetworkAclRule{},
		&AwsNetworkAclRuleList{},

		&AwsRamPrincipalAssociation{},
		&AwsRamPrincipalAssociationList{},

		&AwsSsmMaintenanceWindowTask{},
		&AwsSsmMaintenanceWindowTaskList{},

		&AwsDmsEndpoint{},
		&AwsDmsEndpointList{},

		&AwsRedshiftSecurityGroup{},
		&AwsRedshiftSecurityGroupList{},

		&AwsSesReceiptFilter{},
		&AwsSesReceiptFilterList{},

		&AwsWafIpset{},
		&AwsWafIpsetList{},

		&AwsBackupVault{},
		&AwsBackupVaultList{},

		&AwsDmsReplicationSubnetGroup{},
		&AwsDmsReplicationSubnetGroupList{},

		&AwsWafGeoMatchSet{},
		&AwsWafGeoMatchSetList{},

		&AwsApiGatewayDocumentationVersion{},
		&AwsApiGatewayDocumentationVersionList{},

		&AwsDxPrivateVirtualInterface{},
		&AwsDxPrivateVirtualInterfaceList{},

		&AwsIamUserGroupMembership{},
		&AwsIamUserGroupMembershipList{},

		&AwsLightsailStaticIp{},
		&AwsLightsailStaticIpList{},

		&AwsBatchComputeEnvironment{},
		&AwsBatchComputeEnvironmentList{},

		&AwsDatasyncLocationNfs{},
		&AwsDatasyncLocationNfsList{},

		&AwsIamServiceLinkedRole{},
		&AwsIamServiceLinkedRoleList{},

		&AwsIamUserPolicyAttachment{},
		&AwsIamUserPolicyAttachmentList{},

		&AwsLambdaPermission{},
		&AwsLambdaPermissionList{},

		&AwsMainRouteTableAssociation{},
		&AwsMainRouteTableAssociationList{},

		&AwsApiGatewayUsagePlanKey{},
		&AwsApiGatewayUsagePlanKeyList{},

		&AwsEmrCluster{},
		&AwsEmrClusterList{},

		&AwsAmiCopy{},
		&AwsAmiCopyList{},

		&AwsCloudformationStack{},
		&AwsCloudformationStackList{},

		&AwsElasticBeanstalkEnvironment{},
		&AwsElasticBeanstalkEnvironmentList{},

		&AwsOrganizationsOrganization{},
		&AwsOrganizationsOrganizationList{},

		&AwsS3BucketPolicy{},
		&AwsS3BucketPolicyList{},

		&AwsKmsCiphertext{},
		&AwsKmsCiphertextList{},

		&AwsLambdaFunction{},
		&AwsLambdaFunctionList{},

		&AwsLbListener{},
		&AwsLbListenerList{},

		&AwsEcsCluster{},
		&AwsEcsClusterList{},

		&AwsIamPolicy{},
		&AwsIamPolicyList{},

		&AwsOpsworksUserProfile{},
		&AwsOpsworksUserProfileList{},

		&AwsApiGatewayUsagePlan{},
		&AwsApiGatewayUsagePlanList{},

		&AwsAppmeshVirtualService{},
		&AwsAppmeshVirtualServiceList{},

		&AwsCloudwatchDashboard{},
		&AwsCloudwatchDashboardList{},

		&AwsIamGroupMembership{},
		&AwsIamGroupMembershipList{},

		&AwsIotTopicRule{},
		&AwsIotTopicRuleList{},

		&AwsCognitoIdentityProvider{},
		&AwsCognitoIdentityProviderList{},

		&AwsGlueCatalogDatabase{},
		&AwsGlueCatalogDatabaseList{},

		&AwsProxyProtocolPolicy{},
		&AwsProxyProtocolPolicyList{},

		&AwsRoute53Zone{},
		&AwsRoute53ZoneList{},

		&AwsDefaultVpc{},
		&AwsDefaultVpcList{},

		&AwsPinpointApnsVoipSandboxChannel{},
		&AwsPinpointApnsVoipSandboxChannelList{},

		&AwsWafRegexMatchSet{},
		&AwsWafRegexMatchSetList{},

		&AwsWafregionalSqlInjectionMatchSet{},
		&AwsWafregionalSqlInjectionMatchSetList{},

		&AwsApiGatewayApiKey{},
		&AwsApiGatewayApiKeyList{},

		&AwsApiGatewayGatewayResponse{},
		&AwsApiGatewayGatewayResponseList{},

		&AwsIotThingPrincipalAttachment{},
		&AwsIotThingPrincipalAttachmentList{},

		&AwsRdsClusterInstance{},
		&AwsRdsClusterInstanceList{},

		&AwsResourcegroupsGroup{},
		&AwsResourcegroupsGroupList{},

		&AwsRoute53DelegationSet{},
		&AwsRoute53DelegationSetList{},

		&AwsEbsSnapshotCopy{},
		&AwsEbsSnapshotCopyList{},

		&AwsNetworkInterface{},
		&AwsNetworkInterfaceList{},

		&AwsSubnet{},
		&AwsSubnetList{},

		&AwsWafregionalWebAcl{},
		&AwsWafregionalWebAclList{},

		&AwsAutoscalingPolicy{},
		&AwsAutoscalingPolicyList{},

		&AwsIamUser{},
		&AwsIamUserList{},

		&AwsDefaultRouteTable{},
		&AwsDefaultRouteTableList{},

		&AwsLbListenerCertificate{},
		&AwsLbListenerCertificateList{},

		&AwsAppmeshVirtualRouter{},
		&AwsAppmeshVirtualRouterList{},

		&AwsAppsyncDatasource{},
		&AwsAppsyncDatasourceList{},

		&AwsOpsworksMemcachedLayer{},
		&AwsOpsworksMemcachedLayerList{},

		&AwsSpotFleetRequest{},
		&AwsSpotFleetRequestList{},

		&AwsVpc{},
		&AwsVpcList{},

		&AwsVpnGatewayAttachment{},
		&AwsVpnGatewayAttachmentList{},

		&AwsCodebuildWebhook{},
		&AwsCodebuildWebhookList{},

		&AwsInspectorAssessmentTemplate{},
		&AwsInspectorAssessmentTemplateList{},

		&AwsIotPolicyAttachment{},
		&AwsIotPolicyAttachmentList{},

		&AwsWafByteMatchSet{},
		&AwsWafByteMatchSetList{},

		&AwsAmiFromInstance{},
		&AwsAmiFromInstanceList{},

		&AwsGlueCrawler{},
		&AwsGlueCrawlerList{},

		&AwsSecretsmanagerSecretVersion{},
		&AwsSecretsmanagerSecretVersionList{},

		&AwsVpcDhcpOptions{},
		&AwsVpcDhcpOptionsList{},

		&AwsS3AccountPublicAccessBlock{},
		&AwsS3AccountPublicAccessBlockList{},

		&AwsDefaultSecurityGroup{},
		&AwsDefaultSecurityGroupList{},

		&AwsApiGatewayResource{},
		&AwsApiGatewayResourceList{},

		&AwsApiGatewayStage{},
		&AwsApiGatewayStageList{},

		&AwsDbSubnetGroup{},
		&AwsDbSubnetGroupList{},

		&AwsIamInstanceProfile{},
		&AwsIamInstanceProfileList{},

		&AwsCloudwatchLogSubscriptionFilter{},
		&AwsCloudwatchLogSubscriptionFilterList{},

		&AwsIamAccessKey{},
		&AwsIamAccessKeyList{},

		&AwsOpsworksRdsDbInstance{},
		&AwsOpsworksRdsDbInstanceList{},

		&AwsRoute53HealthCheck{},
		&AwsRoute53HealthCheckList{},

		&AwsSagemakerEndpointConfiguration{},
		&AwsSagemakerEndpointConfigurationList{},

		&AwsDatasyncLocationS3{},
		&AwsDatasyncLocationS3List{},

		&AwsElasticacheSecurityGroup{},
		&AwsElasticacheSecurityGroupList{},

		&AwsDmsReplicationTask{},
		&AwsDmsReplicationTaskList{},

		&AwsDynamodbTableItem{},
		&AwsDynamodbTableItemList{},

		&AwsSecretsmanagerSecret{},
		&AwsSecretsmanagerSecretList{},

		&AwsEcrLifecyclePolicy{},
		&AwsEcrLifecyclePolicyList{},

		&AwsKmsAlias{},
		&AwsKmsAliasList{},

		&AwsRoute53ZoneAssociation{},
		&AwsRoute53ZoneAssociationList{},

		&AwsSesConfigurationSet{},
		&AwsSesConfigurationSetList{},

		&AwsLightsailKeyPair{},
		&AwsLightsailKeyPairList{},

		&AwsStoragegatewayWorkingStorage{},
		&AwsStoragegatewayWorkingStorageList{},

		&AwsPinpointGcmChannel{},
		&AwsPinpointGcmChannelList{},

		&AwsIamAccountAlias{},
		&AwsIamAccountAliasList{},

		&AwsSesDomainIdentityVerification{},
		&AwsSesDomainIdentityVerificationList{},

		&AwsAlbTargetGroup{},
		&AwsAlbTargetGroupList{},

		&AwsEgressOnlyInternetGateway{},
		&AwsEgressOnlyInternetGatewayList{},

		&AwsIotPolicy{},
		&AwsIotPolicyList{},

		&AwsWafRegexPatternSet{},
		&AwsWafRegexPatternSetList{},

		&AwsDirectoryServiceDirectory{},
		&AwsDirectoryServiceDirectoryList{},

		&AwsEip{},
		&AwsEipList{},

		&AwsBackupPlan{},
		&AwsBackupPlanList{},

		&AwsCloudwatchLogDestinationPolicy{},
		&AwsCloudwatchLogDestinationPolicyList{},

		&AwsOpsworksApplication{},
		&AwsOpsworksApplicationList{},

		&AwsRedshiftSnapshotCopyGrant{},
		&AwsRedshiftSnapshotCopyGrantList{},

		&AwsSagemakerNotebookInstanceLifecycleConfiguration{},
		&AwsSagemakerNotebookInstanceLifecycleConfigurationList{},

		&AwsPinpointEmailChannel{},
		&AwsPinpointEmailChannelList{},

		&AwsDxGatewayAssociation{},
		&AwsDxGatewayAssociationList{},

		&AwsEipAssociation{},
		&AwsEipAssociationList{},

		&AwsServicecatalogPortfolio{},
		&AwsServicecatalogPortfolioList{},

		&AwsDxConnection{},
		&AwsDxConnectionList{},

		&AwsIotRoleAlias{},
		&AwsIotRoleAliasList{},

		&AwsS3BucketObject{},
		&AwsS3BucketObjectList{},

		&AwsSecurityhubStandardsSubscription{},
		&AwsSecurityhubStandardsSubscriptionList{},

		&AwsStoragegatewayCache{},
		&AwsStoragegatewayCacheList{},

		&AwsVpnConnection{},
		&AwsVpnConnectionList{},

		&AwsVpcPeeringConnectionOptions{},
		&AwsVpcPeeringConnectionOptionsList{},

		&AwsAcmCertificate{},
		&AwsAcmCertificateList{},

		&AwsAppmeshMesh{},
		&AwsAppmeshMeshList{},

		&AwsCodebuildProject{},
		&AwsCodebuildProjectList{},

		&AwsEc2TransitGateway{},
		&AwsEc2TransitGatewayList{},

		&AwsLaunchTemplate{},
		&AwsLaunchTemplateList{},

		&AwsRdsClusterEndpoint{},
		&AwsRdsClusterEndpointList{},

		&AwsApiGatewayVpcLink{},
		&AwsApiGatewayVpcLinkList{},

		&AwsLbSslNegotiationPolicy{},
		&AwsLbSslNegotiationPolicyList{},

		&AwsOpsworksNodejsAppLayer{},
		&AwsOpsworksNodejsAppLayerList{},

		&AwsSnsTopic{},
		&AwsSnsTopicList{},

		&AwsApiGatewayDomainName{},
		&AwsApiGatewayDomainNameList{},

		&AwsMacieMemberAccountAssociation{},
		&AwsMacieMemberAccountAssociationList{},

		&AwsMacieS3BucketAssociation{},
		&AwsMacieS3BucketAssociationList{},

		&AwsSsmAssociation{},
		&AwsSsmAssociationList{},

		&AwsWafregionalXssMatchSet{},
		&AwsWafregionalXssMatchSetList{},

		&AwsDbEventSubscription{},
		&AwsDbEventSubscriptionList{},

		&AwsSesReceiptRuleSet{},
		&AwsSesReceiptRuleSetList{},

		&AwsSsmDocument{},
		&AwsSsmDocumentList{},

		&AwsSsmMaintenanceWindow{},
		&AwsSsmMaintenanceWindowList{},

		&AwsCloudwatchEventPermission{},
		&AwsCloudwatchEventPermissionList{},

		&AwsGlueConnection{},
		&AwsGlueConnectionList{},

		&AwsGuarddutyDetector{},
		&AwsGuarddutyDetectorList{},

		&AwsInstance{},
		&AwsInstanceList{},

		&AwsNeptuneParameterGroup{},
		&AwsNeptuneParameterGroupList{},

		&AwsOpsworksPermission{},
		&AwsOpsworksPermissionList{},

		&AwsAutoscalingAttachment{},
		&AwsAutoscalingAttachmentList{},

		&AwsEc2TransitGatewayRouteTable{},
		&AwsEc2TransitGatewayRouteTableList{},

		&AwsFlowLog{},
		&AwsFlowLogList{},

		&AwsIamAccountPasswordPolicy{},
		&AwsIamAccountPasswordPolicyList{},

		&AwsIamRole{},
		&AwsIamRoleList{},

		&AwsLightsailDomain{},
		&AwsLightsailDomainList{},

		&AwsAppsyncApiKey{},
		&AwsAppsyncApiKeyList{},

		&AwsAthenaDatabase{},
		&AwsAthenaDatabaseList{},

		&AwsSsmParameter{},
		&AwsSsmParameterList{},

		&AwsStoragegatewayCachedIscsiVolume{},
		&AwsStoragegatewayCachedIscsiVolumeList{},

		&AwsVolumeAttachment{},
		&AwsVolumeAttachmentList{},

		&AwsVpcEndpointSubnetAssociation{},
		&AwsVpcEndpointSubnetAssociationList{},

		&AwsRoute53Record{},
		&AwsRoute53RecordList{},

		&AwsAppCookieStickinessPolicy{},
		&AwsAppCookieStickinessPolicyList{},

		&AwsCloudwatchMetricAlarm{},
		&AwsCloudwatchMetricAlarmList{},

		&AwsCloudwatchLogStream{},
		&AwsCloudwatchLogStreamList{},

		&AwsDlmLifecyclePolicy{},
		&AwsDlmLifecyclePolicyList{},

		&AwsLaunchConfiguration{},
		&AwsLaunchConfigurationList{},

		&AwsRdsGlobalCluster{},
		&AwsRdsGlobalClusterList{},

		&AwsWafregionalRegexPatternSet{},
		&AwsWafregionalRegexPatternSetList{},

		&AwsLbTargetGroupAttachment{},
		&AwsLbTargetGroupAttachmentList{},

		&AwsAppautoscalingScheduledAction{},
		&AwsAppautoscalingScheduledActionList{},

		&AwsDaxParameterGroup{},
		&AwsDaxParameterGroupList{},

		&AwsMqConfiguration{},
		&AwsMqConfigurationList{},

		&AwsAlbListener{},
		&AwsAlbListenerList{},

		&AwsAmiLaunchPermission{},
		&AwsAmiLaunchPermissionList{},

		&AwsApiGatewayMethodResponse{},
		&AwsApiGatewayMethodResponseList{},

		&AwsDxPublicVirtualInterface{},
		&AwsDxPublicVirtualInterfaceList{},

		&AwsInspectorAssessmentTarget{},
		&AwsInspectorAssessmentTargetList{},

		&AwsStoragegatewayNfsFileShare{},
		&AwsStoragegatewayNfsFileShareList{},

		&AwsVpnGatewayRoutePropagation{},
		&AwsVpnGatewayRoutePropagationList{},

		&AwsAlb{},
		&AwsAlbList{},

		&AwsAppautoscalingPolicy{},
		&AwsAppautoscalingPolicyList{},

		&AwsCloudfrontOriginAccessIdentity{},
		&AwsCloudfrontOriginAccessIdentityList{},

		&AwsEcrRepository{},
		&AwsEcrRepositoryList{},

		&AwsLambdaAlias{},
		&AwsLambdaAliasList{},

		&AwsOrganizationsPolicyAttachment{},
		&AwsOrganizationsPolicyAttachmentList{},

		&AwsServiceDiscoveryHttpNamespace{},
		&AwsServiceDiscoveryHttpNamespaceList{},

		&AwsApiGatewayMethod{},
		&AwsApiGatewayMethodList{},

		&AwsNeptuneClusterInstance{},
		&AwsNeptuneClusterInstanceList{},

		&AwsDatasyncLocationEfs{},
		&AwsDatasyncLocationEfsList{},

		&AwsIamRolePolicy{},
		&AwsIamRolePolicyList{},

		&AwsOpsworksPhpAppLayer{},
		&AwsOpsworksPhpAppLayerList{},

		&AwsIamRolePolicyAttachment{},
		&AwsIamRolePolicyAttachmentList{},

		&AwsLb{},
		&AwsLbList{},

		&AwsDmsReplicationInstance{},
		&AwsDmsReplicationInstanceList{},

		&AwsGuarddutyInviteAccepter{},
		&AwsGuarddutyInviteAccepterList{},

		&AwsSesDomainIdentity{},
		&AwsSesDomainIdentityList{},

		&AwsSpotDatafeedSubscription{},
		&AwsSpotDatafeedSubscriptionList{},

		&AwsWafXssMatchSet{},
		&AwsWafXssMatchSetList{},

		&AwsSecurityGroupRule{},
		&AwsSecurityGroupRuleList{},

		&AwsSecurityhubAccount{},
		&AwsSecurityhubAccountList{},

		&AwsCloudfrontDistribution{},
		&AwsCloudfrontDistributionList{},

		&AwsDxGateway{},
		&AwsDxGatewayList{},

		&AwsDxHostedPublicVirtualInterface{},
		&AwsDxHostedPublicVirtualInterfaceList{},

		&AwsGameliftAlias{},
		&AwsGameliftAliasList{},

		&AwsGameliftFleet{},
		&AwsGameliftFleetList{},

		&AwsSecurityGroup{},
		&AwsSecurityGroupList{},

		&AwsDaxSubnetGroup{},
		&AwsDaxSubnetGroupList{},

		&AwsDynamodbGlobalTable{},
		&AwsDynamodbGlobalTableList{},

		&AwsEc2CapacityReservation{},
		&AwsEc2CapacityReservationList{},

		&AwsOpsworksStaticWebLayer{},
		&AwsOpsworksStaticWebLayerList{},

		&AwsRedshiftSubnetGroup{},
		&AwsRedshiftSubnetGroupList{},

		&AwsDaxCluster{},
		&AwsDaxClusterList{},

		&AwsLambdaLayerVersion{},
		&AwsLambdaLayerVersionList{},

		&AwsOpsworksCustomLayer{},
		&AwsOpsworksCustomLayerList{},

		&AwsSsmPatchGroup{},
		&AwsSsmPatchGroupList{},

		&AwsSfnActivity{},
		&AwsSfnActivityList{},

		&AwsEfsMountTarget{},
		&AwsEfsMountTargetList{},

		&AwsElasticacheReplicationGroup{},
		&AwsElasticacheReplicationGroupList{},

		&AwsGuarddutyThreatintelset{},
		&AwsGuarddutyThreatintelsetList{},

		&AwsMediaPackageChannel{},
		&AwsMediaPackageChannelList{},

		&AwsTransferSshKey{},
		&AwsTransferSshKeyList{},

		&AwsApiGatewayBasePathMapping{},
		&AwsApiGatewayBasePathMappingList{},

		&AwsEfsFileSystem{},
		&AwsEfsFileSystemList{},

		&AwsKinesisFirehoseDeliveryStream{},
		&AwsKinesisFirehoseDeliveryStreamList{},

		&AwsRoute{},
		&AwsRouteList{},

		&AwsVpcEndpointConnectionNotification{},
		&AwsVpcEndpointConnectionNotificationList{},

		&AwsAutoscalingGroup{},
		&AwsAutoscalingGroupList{},

		&AwsCognitoIdentityPool{},
		&AwsCognitoIdentityPoolList{},

		&AwsPlacementGroup{},
		&AwsPlacementGroupList{},

		&AwsWorklinkWebsiteCertificateAuthorityAssociation{},
		&AwsWorklinkWebsiteCertificateAuthorityAssociationList{},

		&AwsVpcEndpointService{},
		&AwsVpcEndpointServiceList{},

		&AwsPinpointEventStream{},
		&AwsPinpointEventStreamList{},

		&AwsApiGatewayAuthorizer{},
		&AwsApiGatewayAuthorizerList{},

		&AwsDatasyncAgent{},
		&AwsDatasyncAgentList{},

		&AwsEc2ClientVpnNetworkAssociation{},
		&AwsEc2ClientVpnNetworkAssociationList{},

		&AwsElasticacheCluster{},
		&AwsElasticacheClusterList{},

		&AwsElasticBeanstalkConfigurationTemplate{},
		&AwsElasticBeanstalkConfigurationTemplateList{},

		&AwsGlueClassifier{},
		&AwsGlueClassifierList{},

		&AwsKinesisStream{},
		&AwsKinesisStreamList{},

		&AwsNeptuneCluster{},
		&AwsNeptuneClusterList{},

		&AwsApiGatewayDocumentationPart{},
		&AwsApiGatewayDocumentationPartList{},

		&AwsAppmeshRoute{},
		&AwsAppmeshRouteList{},

		&AwsCodedeployDeploymentGroup{},
		&AwsCodedeployDeploymentGroupList{},

		&AwsVpcIpv4CidrBlockAssociation{},
		&AwsVpcIpv4CidrBlockAssociationList{},

		&AwsCognitoUserPool{},
		&AwsCognitoUserPoolList{},

		&AwsIamGroupPolicy{},
		&AwsIamGroupPolicyList{},

		&AwsCloudfrontPublicKey{},
		&AwsCloudfrontPublicKeyList{},

		&AwsServiceDiscoveryPrivateDnsNamespace{},
		&AwsServiceDiscoveryPrivateDnsNamespaceList{},

		&AwsAlbListenerCertificate{},
		&AwsAlbListenerCertificateList{},

		&AwsApiGatewayIntegration{},
		&AwsApiGatewayIntegrationList{},

		&AwsRedshiftEventSubscription{},
		&AwsRedshiftEventSubscriptionList{},

		&AwsDatasyncTask{},
		&AwsDatasyncTaskList{},

		&AwsDocdbSubnetGroup{},
		&AwsDocdbSubnetGroupList{},

		&AwsApiGatewayModel{},
		&AwsApiGatewayModelList{},

		&AwsGlobalacceleratorAccelerator{},
		&AwsGlobalacceleratorAcceleratorList{},

		&AwsRoute53ResolverEndpoint{},
		&AwsRoute53ResolverEndpointList{},

		&AwsStoragegatewaySmbFileShare{},
		&AwsStoragegatewaySmbFileShareList{},

		&AwsStoragegatewayUploadBuffer{},
		&AwsStoragegatewayUploadBufferList{},

		&AwsCloudwatchLogMetricFilter{},
		&AwsCloudwatchLogMetricFilterList{},

		&AwsDocdbClusterSnapshot{},
		&AwsDocdbClusterSnapshotList{},

		&AwsApiGatewayDeployment{},
		&AwsApiGatewayDeploymentList{},

		&AwsIamOpenidConnectProvider{},
		&AwsIamOpenidConnectProviderList{},

		&AwsLightsailStaticIpAttachment{},
		&AwsLightsailStaticIpAttachmentList{},

		&AwsMediaStoreContainer{},
		&AwsMediaStoreContainerList{},

		&AwsOrganizationsPolicy{},
		&AwsOrganizationsPolicyList{},

		&AwsSnapshotCreateVolumePermission{},
		&AwsSnapshotCreateVolumePermissionList{},

		&AwsAutoscalingNotification{},
		&AwsAutoscalingNotificationList{},

		&AwsIamUserLoginProfile{},
		&AwsIamUserLoginProfileList{},

		&AwsWafregionalRuleGroup{},
		&AwsWafregionalRuleGroupList{},

		&AwsOpsworksJavaAppLayer{},
		&AwsOpsworksJavaAppLayerList{},

		&AwsAppautoscalingTarget{},
		&AwsAppautoscalingTargetList{},

		&AwsIamPolicyAttachment{},
		&AwsIamPolicyAttachmentList{},

		&AwsServiceDiscoveryService{},
		&AwsServiceDiscoveryServiceList{},

		&AwsDbInstanceRoleAssociation{},
		&AwsDbInstanceRoleAssociationList{},

		&AwsRoute53ResolverRuleAssociation{},
		&AwsRoute53ResolverRuleAssociationList{},

		&AwsSimpledbDomain{},
		&AwsSimpledbDomainList{},

		&AwsWafregionalByteMatchSet{},
		&AwsWafregionalByteMatchSetList{},

		&AwsSagemakerNotebookInstance{},
		&AwsSagemakerNotebookInstanceList{},

		&AwsWafRule{},
		&AwsWafRuleList{},

		&AwsDbOptionGroup{},
		&AwsDbOptionGroupList{},

		&AwsDxLag{},
		&AwsDxLagList{},

		&AwsGuarddutyMember{},
		&AwsGuarddutyMemberList{},

		&AwsKinesisAnalyticsApplication{},
		&AwsKinesisAnalyticsApplicationList{},

		&AwsNeptuneClusterSnapshot{},
		&AwsNeptuneClusterSnapshotList{},

		&AwsSagemakerEndpoint{},
		&AwsSagemakerEndpointList{},

		&AwsAlbListenerRule{},
		&AwsAlbListenerRuleList{},

		&AwsCloudhsmV2Cluster{},
		&AwsCloudhsmV2ClusterList{},

		&AwsDxHostedPrivateVirtualInterface{},
		&AwsDxHostedPrivateVirtualInterfaceList{},

		&AwsEc2TransitGatewayRoute{},
		&AwsEc2TransitGatewayRouteList{},

		&AwsGlobalacceleratorListener{},
		&AwsGlobalacceleratorListenerList{},

		&AwsCognitoUserGroup{},
		&AwsCognitoUserGroupList{},

		&AwsCognitoUserPoolDomain{},
		&AwsCognitoUserPoolDomainList{},

		&AwsGlacierVaultLock{},
		&AwsGlacierVaultLockList{},

		&AwsCodecommitRepository{},
		&AwsCodecommitRepositoryList{},

		&AwsCustomerGateway{},
		&AwsCustomerGatewayList{},

		&AwsEbsVolume{},
		&AwsEbsVolumeList{},

		&AwsIamUserSshKey{},
		&AwsIamUserSshKeyList{},

		&AwsOrganizationsAccount{},
		&AwsOrganizationsAccountList{},

		&AwsAthenaNamedQuery{},
		&AwsAthenaNamedQueryList{},

		&AwsRamResourceShare{},
		&AwsRamResourceShareList{},

		&AwsSfnStateMachine{},
		&AwsSfnStateMachineList{},

		&AwsCloudformationStackSet{},
		&AwsCloudformationStackSetList{},

		&AwsConfigConfigurationRecorder{},
		&AwsConfigConfigurationRecorderList{},

		&AwsDevicefarmProject{},
		&AwsDevicefarmProjectList{},

		&AwsElasticacheSubnetGroup{},
		&AwsElasticacheSubnetGroupList{},

		&AwsInspectorResourceGroup{},
		&AwsInspectorResourceGroupList{},

		&AwsSsmActivation{},
		&AwsSsmActivationList{},

		&AwsLbListenerRule{},
		&AwsLbListenerRuleList{},

		&AwsAppsyncResolver{},
		&AwsAppsyncResolverList{},

		&AwsCloud9EnvironmentEc2{},
		&AwsCloud9EnvironmentEc2List{},

		&AwsApiGatewayAccount{},
		&AwsApiGatewayAccountList{},

		&AwsOpsworksInstance{},
		&AwsOpsworksInstanceList{},

		&AwsTransferServer{},
		&AwsTransferServerList{},

		&AwsWafregionalRule{},
		&AwsWafregionalRuleList{},

		&AwsGlueTrigger{},
		&AwsGlueTriggerList{},

		&AwsKmsGrant{},
		&AwsKmsGrantList{},

		&AwsDxHostedPrivateVirtualInterfaceAccepter{},
		&AwsDxHostedPrivateVirtualInterfaceAccepterList{},

		&AwsLightsailInstance{},
		&AwsLightsailInstanceList{},

		&AwsNetworkInterfaceAttachment{},
		&AwsNetworkInterfaceAttachmentList{},

		&AwsSesDomainDkim{},
		&AwsSesDomainDkimList{},

		&AwsSesDomainMailFrom{},
		&AwsSesDomainMailFromList{},

		&AwsBatchJobQueue{},
		&AwsBatchJobQueueList{},

		&AwsXraySamplingRule{},
		&AwsXraySamplingRuleList{},

		&AwsApiGatewayIntegrationResponse{},
		&AwsApiGatewayIntegrationResponseList{},

		&AwsEcsTaskDefinition{},
		&AwsEcsTaskDefinitionList{},

		&AwsNeptuneClusterParameterGroup{},
		&AwsNeptuneClusterParameterGroupList{},

		&AwsSqsQueuePolicy{},
		&AwsSqsQueuePolicyList{},

		&AwsWafRuleGroup{},
		&AwsWafRuleGroupList{},

		&AwsMediaStoreContainerPolicy{},
		&AwsMediaStoreContainerPolicyList{},

		&AwsSsmResourceDataSync{},
		&AwsSsmResourceDataSyncList{},

		&AwsCloudwatchLogDestination{},
		&AwsCloudwatchLogDestinationList{},

		&AwsIamSamlProvider{},
		&AwsIamSamlProviderList{},

		&AwsSesEventDestination{},
		&AwsSesEventDestinationList{},

		&AwsPinpointAdmChannel{},
		&AwsPinpointAdmChannelList{},

		&AwsAmi{},
		&AwsAmiList{},

		&AwsEc2TransitGatewayRouteTableAssociation{},
		&AwsEc2TransitGatewayRouteTableAssociationList{},

		&AwsGlueCatalogTable{},
		&AwsGlueCatalogTableList{},

		&AwsLbTargetGroup{},
		&AwsLbTargetGroupList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
