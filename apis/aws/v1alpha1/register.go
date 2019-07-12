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

		&AwsDatasyncLocationS3{},
		&AwsDatasyncLocationS3List{},

		&AwsIamGroupPolicyAttachment{},
		&AwsIamGroupPolicyAttachmentList{},

		&AwsSesDomainIdentity{},
		&AwsSesDomainIdentityList{},

		&AwsWafregionalRateBasedRule{},
		&AwsWafregionalRateBasedRuleList{},

		&AwsCloudwatchLogDestination{},
		&AwsCloudwatchLogDestinationList{},

		&AwsConfigDeliveryChannel{},
		&AwsConfigDeliveryChannelList{},

		&AwsSagemakerNotebookInstanceLifecycleConfiguration{},
		&AwsSagemakerNotebookInstanceLifecycleConfigurationList{},

		&AwsSesDomainMailFrom{},
		&AwsSesDomainMailFromList{},

		&AwsApiGatewayStage{},
		&AwsApiGatewayStageList{},

		&AwsDbOptionGroup{},
		&AwsDbOptionGroupList{},

		&AwsDirectoryServiceDirectory{},
		&AwsDirectoryServiceDirectoryList{},

		&AwsApiGatewayAccount{},
		&AwsApiGatewayAccountList{},

		&AwsApiGatewayDocumentationVersion{},
		&AwsApiGatewayDocumentationVersionList{},

		&AwsConfigConfigRule{},
		&AwsConfigConfigRuleList{},

		&AwsEcsService{},
		&AwsEcsServiceList{},

		&AwsS3AccountPublicAccessBlock{},
		&AwsS3AccountPublicAccessBlockList{},

		&AwsDmsReplicationInstance{},
		&AwsDmsReplicationInstanceList{},

		&AwsIamGroup{},
		&AwsIamGroupList{},

		&AwsEc2TransitGateway{},
		&AwsEc2TransitGatewayList{},

		&AwsOpsworksMemcachedLayer{},
		&AwsOpsworksMemcachedLayerList{},

		&AwsAutoscalingAttachment{},
		&AwsAutoscalingAttachmentList{},

		&AwsCloudformationStack{},
		&AwsCloudformationStackList{},

		&AwsLambdaFunction{},
		&AwsLambdaFunctionList{},

		&AwsSesReceiptFilter{},
		&AwsSesReceiptFilterList{},

		&AwsStoragegatewayWorkingStorage{},
		&AwsStoragegatewayWorkingStorageList{},

		&AwsWafIpset{},
		&AwsWafIpsetList{},

		&AwsWafRule{},
		&AwsWafRuleList{},

		&AwsDatapipelinePipeline{},
		&AwsDatapipelinePipelineList{},

		&AwsElasticacheParameterGroup{},
		&AwsElasticacheParameterGroupList{},

		&AwsDbSnapshot{},
		&AwsDbSnapshotList{},

		&AwsDmsReplicationSubnetGroup{},
		&AwsDmsReplicationSubnetGroupList{},

		&AwsEcrLifecyclePolicy{},
		&AwsEcrLifecyclePolicyList{},

		&AwsOpsworksRdsDbInstance{},
		&AwsOpsworksRdsDbInstanceList{},

		&AwsCognitoUserPool{},
		&AwsCognitoUserPoolList{},

		&AwsEbsSnapshot{},
		&AwsEbsSnapshotList{},

		&AwsEc2CapacityReservation{},
		&AwsEc2CapacityReservationList{},

		&AwsStoragegatewayUploadBuffer{},
		&AwsStoragegatewayUploadBufferList{},

		&AwsSqsQueue{},
		&AwsSqsQueueList{},

		&AwsPinpointApnsChannel{},
		&AwsPinpointApnsChannelList{},

		&AwsCodedeployApp{},
		&AwsCodedeployAppList{},

		&AwsDynamodbTable{},
		&AwsDynamodbTableList{},

		&AwsGameliftBuild{},
		&AwsGameliftBuildList{},

		&AwsIotThing{},
		&AwsIotThingList{},

		&AwsWafregionalSizeConstraintSet{},
		&AwsWafregionalSizeConstraintSetList{},

		&AwsCloudwatchLogMetricFilter{},
		&AwsCloudwatchLogMetricFilterList{},

		&AwsDbInstance{},
		&AwsDbInstanceList{},

		&AwsDefaultSecurityGroup{},
		&AwsDefaultSecurityGroupList{},

		&AwsCloudwatchLogResourcePolicy{},
		&AwsCloudwatchLogResourcePolicyList{},

		&AwsEc2TransitGatewayRouteTable{},
		&AwsEc2TransitGatewayRouteTableList{},

		&AwsAutoscalingGroup{},
		&AwsAutoscalingGroupList{},

		&AwsGlueCatalogDatabase{},
		&AwsGlueCatalogDatabaseList{},

		&AwsSesEmailIdentity{},
		&AwsSesEmailIdentityList{},

		&AwsS3BucketNotification{},
		&AwsS3BucketNotificationList{},

		&AwsDefaultVpcDhcpOptions{},
		&AwsDefaultVpcDhcpOptionsList{},

		&AwsPinpointAdmChannel{},
		&AwsPinpointAdmChannelList{},

		&AwsAcmCertificateValidation{},
		&AwsAcmCertificateValidationList{},

		&AwsAppsyncApiKey{},
		&AwsAppsyncApiKeyList{},

		&AwsGameliftFleet{},
		&AwsGameliftFleetList{},

		&AwsGuarddutyInviteAccepter{},
		&AwsGuarddutyInviteAccepterList{},

		&AwsLicensemanagerLicenseConfiguration{},
		&AwsLicensemanagerLicenseConfigurationList{},

		&AwsSagemakerEndpoint{},
		&AwsSagemakerEndpointList{},

		&AwsSnapshotCreateVolumePermission{},
		&AwsSnapshotCreateVolumePermissionList{},

		&AwsSfnStateMachine{},
		&AwsSfnStateMachineList{},

		&AwsDynamodbGlobalTable{},
		&AwsDynamodbGlobalTableList{},

		&AwsEc2TransitGatewayVpcAttachmentAccepter{},
		&AwsEc2TransitGatewayVpcAttachmentAccepterList{},

		&AwsSnsPlatformApplication{},
		&AwsSnsPlatformApplicationList{},

		&AwsVpc{},
		&AwsVpcList{},

		&AwsWafregionalRuleGroup{},
		&AwsWafregionalRuleGroupList{},

		&AwsCloudwatchLogDestinationPolicy{},
		&AwsCloudwatchLogDestinationPolicyList{},

		&AwsElasticacheReplicationGroup{},
		&AwsElasticacheReplicationGroupList{},

		&AwsRdsClusterParameterGroup{},
		&AwsRdsClusterParameterGroupList{},

		&AwsRoute53ResolverRuleAssociation{},
		&AwsRoute53ResolverRuleAssociationList{},

		&AwsRouteTable{},
		&AwsRouteTableList{},

		&AwsSesIdentityPolicy{},
		&AwsSesIdentityPolicyList{},

		&AwsServicecatalogPortfolio{},
		&AwsServicecatalogPortfolioList{},

		&AwsVpcPeeringConnection{},
		&AwsVpcPeeringConnectionList{},

		&AwsCloudwatchMetricAlarm{},
		&AwsCloudwatchMetricAlarmList{},

		&AwsDocdbCluster{},
		&AwsDocdbClusterList{},

		&AwsWafWebAcl{},
		&AwsWafWebAclList{},

		&AwsDbSecurityGroup{},
		&AwsDbSecurityGroupList{},

		&AwsGlacierVault{},
		&AwsGlacierVaultList{},

		&AwsS3Bucket{},
		&AwsS3BucketList{},

		&AwsCognitoResourceServer{},
		&AwsCognitoResourceServerList{},

		&AwsCustomerGateway{},
		&AwsCustomerGatewayList{},

		&AwsLightsailKeyPair{},
		&AwsLightsailKeyPairList{},

		&AwsWafSqlInjectionMatchSet{},
		&AwsWafSqlInjectionMatchSetList{},

		&AwsPinpointEmailChannel{},
		&AwsPinpointEmailChannelList{},

		&AwsIamAccessKey{},
		&AwsIamAccessKeyList{},

		&AwsIamAccountPasswordPolicy{},
		&AwsIamAccountPasswordPolicyList{},

		&AwsSesReceiptRuleSet{},
		&AwsSesReceiptRuleSetList{},

		&AwsS3BucketPolicy{},
		&AwsS3BucketPolicyList{},

		&AwsCognitoUserGroup{},
		&AwsCognitoUserGroupList{},

		&AwsEgressOnlyInternetGateway{},
		&AwsEgressOnlyInternetGatewayList{},

		&AwsDbSubnetGroup{},
		&AwsDbSubnetGroupList{},

		&AwsOpsworksInstance{},
		&AwsOpsworksInstanceList{},

		&AwsApiGatewayMethod{},
		&AwsApiGatewayMethodList{},

		&AwsCodebuildProject{},
		&AwsCodebuildProjectList{},

		&AwsGuarddutyDetector{},
		&AwsGuarddutyDetectorList{},

		&AwsOrganizationsAccount{},
		&AwsOrganizationsAccountList{},

		&AwsSecretsmanagerSecretVersion{},
		&AwsSecretsmanagerSecretVersionList{},

		&AwsCodebuildWebhook{},
		&AwsCodebuildWebhookList{},

		&AwsDirectoryServiceConditionalForwarder{},
		&AwsDirectoryServiceConditionalForwarderList{},

		&AwsVpcEndpointSubnetAssociation{},
		&AwsVpcEndpointSubnetAssociationList{},

		&AwsVpnGateway{},
		&AwsVpnGatewayList{},

		&AwsWafregionalByteMatchSet{},
		&AwsWafregionalByteMatchSetList{},

		&AwsEc2ClientVpnNetworkAssociation{},
		&AwsEc2ClientVpnNetworkAssociationList{},

		&AwsSnsTopicSubscription{},
		&AwsSnsTopicSubscriptionList{},

		&AwsVpcEndpointConnectionNotification{},
		&AwsVpcEndpointConnectionNotificationList{},

		&AwsVpcEndpointServiceAllowedPrincipal{},
		&AwsVpcEndpointServiceAllowedPrincipalList{},

		&AwsWafSizeConstraintSet{},
		&AwsWafSizeConstraintSetList{},

		&AwsWafregionalWebAclAssociation{},
		&AwsWafregionalWebAclAssociationList{},

		&AwsIotPolicy{},
		&AwsIotPolicyList{},

		&AwsServiceDiscoveryService{},
		&AwsServiceDiscoveryServiceList{},

		&AwsLaunchTemplate{},
		&AwsLaunchTemplateList{},

		&AwsRedshiftSnapshotCopyGrant{},
		&AwsRedshiftSnapshotCopyGrantList{},

		&AwsWafregionalIpset{},
		&AwsWafregionalIpsetList{},

		&AwsGlacierVaultLock{},
		&AwsGlacierVaultLockList{},

		&AwsIamUserSshKey{},
		&AwsIamUserSshKeyList{},

		&AwsAmiFromInstance{},
		&AwsAmiFromInstanceList{},

		&AwsNeptuneClusterInstance{},
		&AwsNeptuneClusterInstanceList{},

		&AwsIamUser{},
		&AwsIamUserList{},

		&AwsKinesisFirehoseDeliveryStream{},
		&AwsKinesisFirehoseDeliveryStreamList{},

		&AwsAmi{},
		&AwsAmiList{},

		&AwsElastictranscoderPreset{},
		&AwsElastictranscoderPresetList{},

		&AwsRdsGlobalCluster{},
		&AwsRdsGlobalClusterList{},

		&AwsVpnGatewayRoutePropagation{},
		&AwsVpnGatewayRoutePropagationList{},

		&AwsKmsKey{},
		&AwsKmsKeyList{},

		&AwsNeptuneCluster{},
		&AwsNeptuneClusterList{},

		&AwsRoute53QueryLog{},
		&AwsRoute53QueryLogList{},

		&AwsElasticacheSecurityGroup{},
		&AwsElasticacheSecurityGroupList{},

		&AwsRamResourceShare{},
		&AwsRamResourceShareList{},

		&AwsKinesisAnalyticsApplication{},
		&AwsKinesisAnalyticsApplicationList{},

		&AwsResourcegroupsGroup{},
		&AwsResourcegroupsGroupList{},

		&AwsRoute53DelegationSet{},
		&AwsRoute53DelegationSetList{},

		&AwsAlbListenerCertificate{},
		&AwsAlbListenerCertificateList{},

		&AwsAppmeshVirtualNode{},
		&AwsAppmeshVirtualNodeList{},

		&AwsGlueSecurityConfiguration{},
		&AwsGlueSecurityConfigurationList{},

		&AwsApiGatewayGatewayResponse{},
		&AwsApiGatewayGatewayResponseList{},

		&AwsIotThingType{},
		&AwsIotThingTypeList{},

		&AwsAppsyncFunction{},
		&AwsAppsyncFunctionList{},

		&AwsIamServiceLinkedRole{},
		&AwsIamServiceLinkedRoleList{},

		&AwsDatasyncTask{},
		&AwsDatasyncTaskList{},

		&AwsMacieS3BucketAssociation{},
		&AwsMacieS3BucketAssociationList{},

		&AwsSpotFleetRequest{},
		&AwsSpotFleetRequestList{},

		&AwsWafregionalXssMatchSet{},
		&AwsWafregionalXssMatchSetList{},

		&AwsBackupPlan{},
		&AwsBackupPlanList{},

		&AwsConfigConfigurationRecorder{},
		&AwsConfigConfigurationRecorderList{},

		&AwsNetworkAcl{},
		&AwsNetworkAclList{},

		&AwsSesReceiptRule{},
		&AwsSesReceiptRuleList{},

		&AwsStoragegatewaySmbFileShare{},
		&AwsStoragegatewaySmbFileShareList{},

		&AwsBatchJobDefinition{},
		&AwsBatchJobDefinitionList{},

		&AwsCurReportDefinition{},
		&AwsCurReportDefinitionList{},

		&AwsLightsailInstance{},
		&AwsLightsailInstanceList{},

		&AwsMediaPackageChannel{},
		&AwsMediaPackageChannelList{},

		&AwsOpsworksApplication{},
		&AwsOpsworksApplicationList{},

		&AwsCognitoUserPoolDomain{},
		&AwsCognitoUserPoolDomainList{},

		&AwsIamRolePolicy{},
		&AwsIamRolePolicyList{},

		&AwsWafregionalSqlInjectionMatchSet{},
		&AwsWafregionalSqlInjectionMatchSetList{},

		&AwsOrganizationsPolicy{},
		&AwsOrganizationsPolicyList{},

		&AwsSesConfigurationSet{},
		&AwsSesConfigurationSetList{},

		&AwsDaxParameterGroup{},
		&AwsDaxParameterGroupList{},

		&AwsGlobalacceleratorAccelerator{},
		&AwsGlobalacceleratorAcceleratorList{},

		&AwsRdsCluster{},
		&AwsRdsClusterList{},

		&AwsVpcPeeringConnectionAccepter{},
		&AwsVpcPeeringConnectionAccepterList{},

		&AwsCloudwatchEventPermission{},
		&AwsCloudwatchEventPermissionList{},

		&AwsDatasyncLocationEfs{},
		&AwsDatasyncLocationEfsList{},

		&AwsEbsSnapshotCopy{},
		&AwsEbsSnapshotCopyList{},

		&AwsEc2TransitGatewayRouteTableAssociation{},
		&AwsEc2TransitGatewayRouteTableAssociationList{},

		&AwsIamUserPolicyAttachment{},
		&AwsIamUserPolicyAttachmentList{},

		&AwsApiGatewayUsagePlan{},
		&AwsApiGatewayUsagePlanList{},

		&AwsDxHostedPublicVirtualInterfaceAccepter{},
		&AwsDxHostedPublicVirtualInterfaceAccepterList{},

		&AwsGlueTrigger{},
		&AwsGlueTriggerList{},

		&AwsPinpointApnsVoipSandboxChannel{},
		&AwsPinpointApnsVoipSandboxChannelList{},

		&AwsCloudformationStackSetInstance{},
		&AwsCloudformationStackSetInstanceList{},

		&AwsLaunchConfiguration{},
		&AwsLaunchConfigurationList{},

		&AwsAppautoscalingPolicy{},
		&AwsAppautoscalingPolicyList{},

		&AwsBackupVault{},
		&AwsBackupVaultList{},

		&AwsFlowLog{},
		&AwsFlowLogList{},

		&AwsOpsworksHaproxyLayer{},
		&AwsOpsworksHaproxyLayerList{},

		&AwsOpsworksCustomLayer{},
		&AwsOpsworksCustomLayerList{},

		&AwsSfnActivity{},
		&AwsSfnActivityList{},

		&AwsAppmeshVirtualRouter{},
		&AwsAppmeshVirtualRouterList{},

		&AwsEc2Fleet{},
		&AwsEc2FleetList{},

		&AwsMqBroker{},
		&AwsMqBrokerList{},

		&AwsDaxCluster{},
		&AwsDaxClusterList{},

		&AwsDaxSubnetGroup{},
		&AwsDaxSubnetGroupList{},

		&AwsKmsExternalKey{},
		&AwsKmsExternalKeyList{},

		&AwsOpsworksUserProfile{},
		&AwsOpsworksUserProfileList{},

		&AwsRoute53ResolverRule{},
		&AwsRoute53ResolverRuleList{},

		&AwsIamOpenidConnectProvider{},
		&AwsIamOpenidConnectProviderList{},

		&AwsIotThingPrincipalAttachment{},
		&AwsIotThingPrincipalAttachmentList{},

		&AwsRdsClusterEndpoint{},
		&AwsRdsClusterEndpointList{},

		&AwsVpcEndpoint{},
		&AwsVpcEndpointList{},

		&AwsDbClusterSnapshot{},
		&AwsDbClusterSnapshotList{},

		&AwsIamSamlProvider{},
		&AwsIamSamlProviderList{},

		&AwsSesIdentityNotificationTopic{},
		&AwsSesIdentityNotificationTopicList{},

		&AwsDlmLifecyclePolicy{},
		&AwsDlmLifecyclePolicyList{},

		&AwsDxGatewayAssociationProposal{},
		&AwsDxGatewayAssociationProposalList{},

		&AwsEc2TransitGatewayVpcAttachment{},
		&AwsEc2TransitGatewayVpcAttachmentList{},

		&AwsIamGroupMembership{},
		&AwsIamGroupMembershipList{},

		&AwsIamRolePolicyAttachment{},
		&AwsIamRolePolicyAttachmentList{},

		&AwsInternetGateway{},
		&AwsInternetGatewayList{},

		&AwsXraySamplingRule{},
		&AwsXraySamplingRuleList{},

		&AwsAppCookieStickinessPolicy{},
		&AwsAppCookieStickinessPolicyList{},

		&AwsCodepipelineWebhook{},
		&AwsCodepipelineWebhookList{},

		&AwsElasticBeanstalkEnvironment{},
		&AwsElasticBeanstalkEnvironmentList{},

		&AwsElastictranscoderPipeline{},
		&AwsElastictranscoderPipelineList{},

		&AwsGameliftAlias{},
		&AwsGameliftAliasList{},

		&AwsMediaStoreContainer{},
		&AwsMediaStoreContainerList{},

		&AwsSagemakerNotebookInstance{},
		&AwsSagemakerNotebookInstanceList{},

		&AwsDevicefarmProject{},
		&AwsDevicefarmProjectList{},

		&AwsDxGatewayAssociation{},
		&AwsDxGatewayAssociationList{},

		&AwsAmiCopy{},
		&AwsAmiCopyList{},

		&AwsWorklinkWebsiteCertificateAuthorityAssociation{},
		&AwsWorklinkWebsiteCertificateAuthorityAssociationList{},

		&AwsAutoscalingSchedule{},
		&AwsAutoscalingScheduleList{},

		&AwsWafByteMatchSet{},
		&AwsWafByteMatchSetList{},

		&AwsDxConnection{},
		&AwsDxConnectionList{},

		&AwsCloudfrontDistribution{},
		&AwsCloudfrontDistributionList{},

		&AwsDocdbSubnetGroup{},
		&AwsDocdbSubnetGroupList{},

		&AwsApiGatewayVpcLink{},
		&AwsApiGatewayVpcLinkList{},

		&AwsS3BucketObject{},
		&AwsS3BucketObjectList{},

		&AwsApiGatewayAuthorizer{},
		&AwsApiGatewayAuthorizerList{},

		&AwsApiGatewayMethodResponse{},
		&AwsApiGatewayMethodResponseList{},

		&AwsRoute53Zone{},
		&AwsRoute53ZoneList{},

		&AwsInstance{},
		&AwsInstanceList{},

		&AwsIotRoleAlias{},
		&AwsIotRoleAliasList{},

		&AwsGlueConnection{},
		&AwsGlueConnectionList{},

		&AwsWafregionalRegexMatchSet{},
		&AwsWafregionalRegexMatchSetList{},

		&AwsKmsCiphertext{},
		&AwsKmsCiphertextList{},

		&AwsSecurityhubStandardsSubscription{},
		&AwsSecurityhubStandardsSubscriptionList{},

		&AwsApiGatewayIntegrationResponse{},
		&AwsApiGatewayIntegrationResponseList{},

		&AwsDxHostedPrivateVirtualInterface{},
		&AwsDxHostedPrivateVirtualInterfaceList{},

		&AwsGlueCatalogTable{},
		&AwsGlueCatalogTableList{},

		&AwsOrganizationsOrganizationalUnit{},
		&AwsOrganizationsOrganizationalUnitList{},

		&AwsRedshiftEventSubscription{},
		&AwsRedshiftEventSubscriptionList{},

		&AwsSsmPatchGroup{},
		&AwsSsmPatchGroupList{},

		&AwsDxPrivateVirtualInterface{},
		&AwsDxPrivateVirtualInterfaceList{},

		&AwsEfsMountTarget{},
		&AwsEfsMountTargetList{},

		&AwsMskConfiguration{},
		&AwsMskConfigurationList{},

		&AwsOpsworksStack{},
		&AwsOpsworksStackList{},

		&AwsSsmParameter{},
		&AwsSsmParameterList{},

		&AwsPinpointGcmChannel{},
		&AwsPinpointGcmChannelList{},

		&AwsCognitoIdentityPool{},
		&AwsCognitoIdentityPoolList{},

		&AwsKmsGrant{},
		&AwsKmsGrantList{},

		&AwsElasticacheSubnetGroup{},
		&AwsElasticacheSubnetGroupList{},

		&AwsIamUserGroupMembership{},
		&AwsIamUserGroupMembershipList{},

		&AwsSsmDocument{},
		&AwsSsmDocumentList{},

		&AwsAmiLaunchPermission{},
		&AwsAmiLaunchPermissionList{},

		&AwsCodedeployDeploymentGroup{},
		&AwsCodedeployDeploymentGroupList{},

		&AwsStoragegatewayNfsFileShare{},
		&AwsStoragegatewayNfsFileShareList{},

		&AwsAppsyncDatasource{},
		&AwsAppsyncDatasourceList{},

		&AwsCodepipeline{},
		&AwsCodepipelineList{},

		&AwsDxConnectionAssociation{},
		&AwsDxConnectionAssociationList{},

		&AwsInspectorResourceGroup{},
		&AwsInspectorResourceGroupList{},

		&AwsSimpledbDomain{},
		&AwsSimpledbDomainList{},

		&AwsApiGatewayDeployment{},
		&AwsApiGatewayDeploymentList{},

		&AwsCognitoUserPoolClient{},
		&AwsCognitoUserPoolClientList{},

		&AwsCodedeployDeploymentConfig{},
		&AwsCodedeployDeploymentConfigList{},

		&AwsRamPrincipalAssociation{},
		&AwsRamPrincipalAssociationList{},

		&AwsRamResourceAssociation{},
		&AwsRamResourceAssociationList{},

		&AwsAthenaDatabase{},
		&AwsAthenaDatabaseList{},

		&AwsDynamodbTableItem{},
		&AwsDynamodbTableItemList{},

		&AwsMediaStoreContainerPolicy{},
		&AwsMediaStoreContainerPolicyList{},

		&AwsRoute53HealthCheck{},
		&AwsRoute53HealthCheckList{},

		&AwsSecurityGroupRule{},
		&AwsSecurityGroupRuleList{},

		&AwsGuarddutyThreatintelset{},
		&AwsGuarddutyThreatintelsetList{},

		&AwsIamInstanceProfile{},
		&AwsIamInstanceProfileList{},

		&AwsDbEventSubscription{},
		&AwsDbEventSubscriptionList{},

		&AwsSsmPatchBaseline{},
		&AwsSsmPatchBaselineList{},

		&AwsApiGatewayResource{},
		&AwsApiGatewayResourceList{},

		&AwsCloud9EnvironmentEc2{},
		&AwsCloud9EnvironmentEc2List{},

		&AwsIotCertificate{},
		&AwsIotCertificateList{},

		&AwsCloudhsmV2Hsm{},
		&AwsCloudhsmV2HsmList{},

		&AwsDxHostedPublicVirtualInterface{},
		&AwsDxHostedPublicVirtualInterfaceList{},

		&AwsNetworkAclRule{},
		&AwsNetworkAclRuleList{},

		&AwsNetworkInterfaceSgAttachment{},
		&AwsNetworkInterfaceSgAttachmentList{},

		&AwsSsmMaintenanceWindowTarget{},
		&AwsSsmMaintenanceWindowTargetList{},

		&AwsBatchComputeEnvironment{},
		&AwsBatchComputeEnvironmentList{},

		&AwsLbTargetGroup{},
		&AwsLbTargetGroupList{},

		&AwsLoadBalancerPolicy{},
		&AwsLoadBalancerPolicyList{},

		&AwsLoadBalancerBackendServerPolicy{},
		&AwsLoadBalancerBackendServerPolicyList{},

		&AwsDxBgpPeer{},
		&AwsDxBgpPeerList{},

		&AwsAlbListenerRule{},
		&AwsAlbListenerRuleList{},

		&AwsOpsworksNodejsAppLayer{},
		&AwsOpsworksNodejsAppLayerList{},

		&AwsSecurityhubProductSubscription{},
		&AwsSecurityhubProductSubscriptionList{},

		&AwsLightsailDomain{},
		&AwsLightsailDomainList{},

		&AwsNeptuneClusterParameterGroup{},
		&AwsNeptuneClusterParameterGroupList{},

		&AwsNeptuneClusterSnapshot{},
		&AwsNeptuneClusterSnapshotList{},

		&AwsDefaultSubnet{},
		&AwsDefaultSubnetList{},

		&AwsAcmCertificate{},
		&AwsAcmCertificateList{},

		&AwsDocdbClusterSnapshot{},
		&AwsDocdbClusterSnapshotList{},

		&AwsMskCluster{},
		&AwsMskClusterList{},

		&AwsOpsworksPhpAppLayer{},
		&AwsOpsworksPhpAppLayerList{},

		&AwsOpsworksRailsAppLayer{},
		&AwsOpsworksRailsAppLayerList{},

		&AwsDxGateway{},
		&AwsDxGatewayList{},

		&AwsElasticBeanstalkApplicationVersion{},
		&AwsElasticBeanstalkApplicationVersionList{},

		&AwsRdsClusterInstance{},
		&AwsRdsClusterInstanceList{},

		&AwsServiceDiscoveryPrivateDnsNamespace{},
		&AwsServiceDiscoveryPrivateDnsNamespaceList{},

		&AwsCloudwatchDashboard{},
		&AwsCloudwatchDashboardList{},

		&AwsDefaultNetworkAcl{},
		&AwsDefaultNetworkAclList{},

		&AwsAthenaWorkgroup{},
		&AwsAthenaWorkgroupList{},

		&AwsLb{},
		&AwsLbList{},

		&AwsRouteTableAssociation{},
		&AwsRouteTableAssociationList{},

		&AwsWafRuleGroup{},
		&AwsWafRuleGroupList{},

		&AwsEbsDefaultKmsKey{},
		&AwsEbsDefaultKmsKeyList{},

		&AwsGlueCrawler{},
		&AwsGlueCrawlerList{},

		&AwsEmrInstanceGroup{},
		&AwsEmrInstanceGroupList{},

		&AwsOpsworksMysqlLayer{},
		&AwsOpsworksMysqlLayerList{},

		&AwsS3BucketPublicAccessBlock{},
		&AwsS3BucketPublicAccessBlockList{},

		&AwsSnsTopicPolicy{},
		&AwsSnsTopicPolicyList{},

		&AwsConfigAggregateAuthorization{},
		&AwsConfigAggregateAuthorizationList{},

		&AwsDmsReplicationTask{},
		&AwsDmsReplicationTaskList{},

		&AwsConfigConfigurationRecorderStatus{},
		&AwsConfigConfigurationRecorderStatusList{},

		&AwsDmsEndpoint{},
		&AwsDmsEndpointList{},

		&AwsBudgetsBudget{},
		&AwsBudgetsBudgetList{},

		&AwsCloudwatchLogStream{},
		&AwsCloudwatchLogStreamList{},

		&AwsSecretsmanagerSecret{},
		&AwsSecretsmanagerSecretList{},

		&AwsVpnConnectionRoute{},
		&AwsVpnConnectionRouteList{},

		&AwsDbParameterGroup{},
		&AwsDbParameterGroupList{},

		&AwsOpsworksJavaAppLayer{},
		&AwsOpsworksJavaAppLayerList{},

		&AwsOrganizationsOrganization{},
		&AwsOrganizationsOrganizationList{},

		&AwsSesTemplate{},
		&AwsSesTemplateList{},

		&AwsVpcPeeringConnectionOptions{},
		&AwsVpcPeeringConnectionOptionsList{},

		&AwsCloudwatchEventRule{},
		&AwsCloudwatchEventRuleList{},

		&AwsIamUserPolicy{},
		&AwsIamUserPolicyList{},

		&AwsAppmeshRoute{},
		&AwsAppmeshRouteList{},

		&AwsTransferSshKey{},
		&AwsTransferSshKeyList{},

		&AwsWorklinkFleet{},
		&AwsWorklinkFleetList{},

		&AwsPinpointEventStream{},
		&AwsPinpointEventStreamList{},

		&AwsPlacementGroup{},
		&AwsPlacementGroupList{},

		&AwsSqsQueuePolicy{},
		&AwsSqsQueuePolicyList{},

		&AwsTransferServer{},
		&AwsTransferServerList{},

		&AwsDatasyncAgent{},
		&AwsDatasyncAgentList{},

		&AwsEmrCluster{},
		&AwsEmrClusterList{},

		&AwsRoute53ResolverEndpoint{},
		&AwsRoute53ResolverEndpointList{},

		&AwsDbInstanceRoleAssociation{},
		&AwsDbInstanceRoleAssociationList{},

		&AwsIamUserLoginProfile{},
		&AwsIamUserLoginProfileList{},

		&AwsOrganizationsPolicyAttachment{},
		&AwsOrganizationsPolicyAttachmentList{},

		&AwsShieldProtection{},
		&AwsShieldProtectionList{},

		&AwsEbsEncryptionByDefault{},
		&AwsEbsEncryptionByDefaultList{},

		&AwsElasticacheCluster{},
		&AwsElasticacheClusterList{},

		&AwsBackupSelection{},
		&AwsBackupSelectionList{},

		&AwsEcrRepository{},
		&AwsEcrRepositoryList{},

		&AwsEksCluster{},
		&AwsEksClusterList{},

		&AwsGlueJob{},
		&AwsGlueJobList{},

		&AwsMqConfiguration{},
		&AwsMqConfigurationList{},

		&AwsNeptuneParameterGroup{},
		&AwsNeptuneParameterGroupList{},

		&AwsAcmpcaCertificateAuthority{},
		&AwsAcmpcaCertificateAuthorityList{},

		&AwsApiGatewayRequestValidator{},
		&AwsApiGatewayRequestValidatorList{},

		&AwsDefaultRouteTable{},
		&AwsDefaultRouteTableList{},

		&AwsPinpointApnsVoipChannel{},
		&AwsPinpointApnsVoipChannelList{},

		&AwsLbListenerCertificate{},
		&AwsLbListenerCertificateList{},

		&AwsDocdbClusterParameterGroup{},
		&AwsDocdbClusterParameterGroupList{},

		&AwsDxLag{},
		&AwsDxLagList{},

		&AwsSsmAssociation{},
		&AwsSsmAssociationList{},

		&AwsVolumeAttachment{},
		&AwsVolumeAttachmentList{},

		&AwsAlbTargetGroup{},
		&AwsAlbTargetGroupList{},

		&AwsInspectorAssessmentTarget{},
		&AwsInspectorAssessmentTargetList{},

		&AwsLightsailStaticIp{},
		&AwsLightsailStaticIpList{},

		&AwsRoute53ZoneAssociation{},
		&AwsRoute53ZoneAssociationList{},

		&AwsRoute{},
		&AwsRouteList{},

		&AwsCodecommitRepository{},
		&AwsCodecommitRepositoryList{},

		&AwsLicensemanagerAssociation{},
		&AwsLicensemanagerAssociationList{},

		&AwsElbAttachment{},
		&AwsElbAttachmentList{},

		&AwsNatGateway{},
		&AwsNatGatewayList{},

		&AwsRoute53Record{},
		&AwsRoute53RecordList{},

		&AwsSagemakerModel{},
		&AwsSagemakerModelList{},

		&AwsAppsyncGraphqlApi{},
		&AwsAppsyncGraphqlApiList{},

		&AwsCloudwatchLogGroup{},
		&AwsCloudwatchLogGroupList{},

		&AwsOpsworksStaticWebLayer{},
		&AwsOpsworksStaticWebLayerList{},

		&AwsSubnet{},
		&AwsSubnetList{},

		&AwsEc2TransitGatewayRouteTablePropagation{},
		&AwsEc2TransitGatewayRouteTablePropagationList{},

		&AwsGameliftGameSessionQueue{},
		&AwsGameliftGameSessionQueueList{},

		&AwsCodecommitTrigger{},
		&AwsCodecommitTriggerList{},

		&AwsEcrRepositoryPolicy{},
		&AwsEcrRepositoryPolicyList{},

		&AwsLbSslNegotiationPolicy{},
		&AwsLbSslNegotiationPolicyList{},

		&AwsApiGatewayClientCertificate{},
		&AwsApiGatewayClientCertificateList{},

		&AwsCognitoIdentityProvider{},
		&AwsCognitoIdentityProviderList{},

		&AwsSpotInstanceRequest{},
		&AwsSpotInstanceRequestList{},

		&AwsPinpointSmsChannel{},
		&AwsPinpointSmsChannelList{},

		&AwsSpotDatafeedSubscription{},
		&AwsSpotDatafeedSubscriptionList{},

		&AwsPinpointBaiduChannel{},
		&AwsPinpointBaiduChannelList{},

		&AwsPinpointApnsSandboxChannel{},
		&AwsPinpointApnsSandboxChannelList{},

		&AwsVpcEndpointService{},
		&AwsVpcEndpointServiceList{},

		&AwsVpnGatewayAttachment{},
		&AwsVpnGatewayAttachmentList{},

		&AwsRedshiftSubnetGroup{},
		&AwsRedshiftSubnetGroupList{},

		&AwsEipAssociation{},
		&AwsEipAssociationList{},

		&AwsRedshiftSecurityGroup{},
		&AwsRedshiftSecurityGroupList{},

		&AwsGlobalacceleratorListener{},
		&AwsGlobalacceleratorListenerList{},

		&AwsVpcDhcpOptions{},
		&AwsVpcDhcpOptionsList{},

		&AwsLambdaEventSourceMapping{},
		&AwsLambdaEventSourceMappingList{},

		&AwsStoragegatewayCachedIscsiVolume{},
		&AwsStoragegatewayCachedIscsiVolumeList{},

		&AwsWafregionalRule{},
		&AwsWafregionalRuleList{},

		&AwsCloudwatchEventTarget{},
		&AwsCloudwatchEventTargetList{},

		&AwsCognitoIdentityPoolRolesAttachment{},
		&AwsCognitoIdentityPoolRolesAttachmentList{},

		&AwsLbCookieStickinessPolicy{},
		&AwsLbCookieStickinessPolicyList{},

		&AwsNeptuneSubnetGroup{},
		&AwsNeptuneSubnetGroupList{},

		&AwsRedshiftCluster{},
		&AwsRedshiftClusterList{},

		&AwsPinpointApp{},
		&AwsPinpointAppList{},

		&AwsAlbListener{},
		&AwsAlbListenerList{},

		&AwsLbListener{},
		&AwsLbListenerList{},

		&AwsAppmeshMesh{},
		&AwsAppmeshMeshList{},

		&AwsIamAccountAlias{},
		&AwsIamAccountAliasList{},

		&AwsSecurityhubAccount{},
		&AwsSecurityhubAccountList{},

		&AwsDefaultVpc{},
		&AwsDefaultVpcList{},

		&AwsWafregionalWebAcl{},
		&AwsWafregionalWebAclList{},

		&AwsNetworkInterfaceAttachment{},
		&AwsNetworkInterfaceAttachmentList{},

		&AwsRedshiftParameterGroup{},
		&AwsRedshiftParameterGroupList{},

		&AwsCloudtrail{},
		&AwsCloudtrailList{},

		&AwsEcsTaskDefinition{},
		&AwsEcsTaskDefinitionList{},

		&AwsIotTopicRule{},
		&AwsIotTopicRuleList{},

		&AwsNeptuneEventSubscription{},
		&AwsNeptuneEventSubscriptionList{},

		&AwsAutoscalingLifecycleHook{},
		&AwsAutoscalingLifecycleHookList{},

		&AwsAutoscalingNotification{},
		&AwsAutoscalingNotificationList{},

		&AwsElasticBeanstalkApplication{},
		&AwsElasticBeanstalkApplicationList{},

		&AwsOpsworksPermission{},
		&AwsOpsworksPermissionList{},

		&AwsApiGatewayMethodSettings{},
		&AwsApiGatewayMethodSettingsList{},

		&AwsEcsCluster{},
		&AwsEcsClusterList{},

		&AwsIamGroupPolicy{},
		&AwsIamGroupPolicyList{},

		&AwsIamPolicy{},
		&AwsIamPolicyList{},

		&AwsWafRegexPatternSet{},
		&AwsWafRegexPatternSetList{},

		&AwsBatchJobQueue{},
		&AwsBatchJobQueueList{},

		&AwsApiGatewayRestApi{},
		&AwsApiGatewayRestApiList{},

		&AwsDatasyncLocationNfs{},
		&AwsDatasyncLocationNfsList{},

		&AwsSesEventDestination{},
		&AwsSesEventDestinationList{},

		&AwsVpcDhcpOptionsAssociation{},
		&AwsVpcDhcpOptionsAssociationList{},

		&AwsElasticBeanstalkConfigurationTemplate{},
		&AwsElasticBeanstalkConfigurationTemplateList{},

		&AwsMainRouteTableAssociation{},
		&AwsMainRouteTableAssociationList{},

		&AwsEfsFileSystem{},
		&AwsEfsFileSystemList{},

		&AwsS3BucketMetric{},
		&AwsS3BucketMetricList{},

		&AwsVpcIpv4CidrBlockAssociation{},
		&AwsVpcIpv4CidrBlockAssociationList{},

		&AwsLbListenerRule{},
		&AwsLbListenerRuleList{},

		&AwsAlbTargetGroupAttachment{},
		&AwsAlbTargetGroupAttachmentList{},

		&AwsElasticsearchDomain{},
		&AwsElasticsearchDomainList{},

		&AwsLambdaPermission{},
		&AwsLambdaPermissionList{},

		&AwsLambdaAlias{},
		&AwsLambdaAliasList{},

		&AwsSesDomainIdentityVerification{},
		&AwsSesDomainIdentityVerificationList{},

		&AwsSesDomainDkim{},
		&AwsSesDomainDkimList{},

		&AwsWafregionalRegexPatternSet{},
		&AwsWafregionalRegexPatternSetList{},

		&AwsApiGatewayApiKey{},
		&AwsApiGatewayApiKeyList{},

		&AwsAppautoscalingScheduledAction{},
		&AwsAppautoscalingScheduledActionList{},

		&AwsIamRole{},
		&AwsIamRoleList{},

		&AwsLoadBalancerListenerPolicy{},
		&AwsLoadBalancerListenerPolicyList{},

		&AwsStoragegatewayGateway{},
		&AwsStoragegatewayGatewayList{},

		&AwsLbTargetGroupAttachment{},
		&AwsLbTargetGroupAttachmentList{},

		&AwsCloudfrontOriginAccessIdentity{},
		&AwsCloudfrontOriginAccessIdentityList{},

		&AwsElasticsearchDomainPolicy{},
		&AwsElasticsearchDomainPolicyList{},

		&AwsDxPublicVirtualInterface{},
		&AwsDxPublicVirtualInterfaceList{},

		&AwsElb{},
		&AwsElbList{},

		&AwsS3BucketInventory{},
		&AwsS3BucketInventoryList{},

		&AwsWafRegexMatchSet{},
		&AwsWafRegexMatchSetList{},

		&AwsAppsyncResolver{},
		&AwsAppsyncResolverList{},

		&AwsCloudformationStackSet{},
		&AwsCloudformationStackSetList{},

		&AwsInspectorAssessmentTemplate{},
		&AwsInspectorAssessmentTemplateList{},

		&AwsProxyProtocolPolicy{},
		&AwsProxyProtocolPolicyList{},

		&AwsApiGatewayDomainName{},
		&AwsApiGatewayDomainNameList{},

		&AwsEc2TransitGatewayRoute{},
		&AwsEc2TransitGatewayRouteList{},

		&AwsEbsVolume{},
		&AwsEbsVolumeList{},

		&AwsStoragegatewayCache{},
		&AwsStoragegatewayCacheList{},

		&AwsWafGeoMatchSet{},
		&AwsWafGeoMatchSetList{},

		&AwsApiGatewayUsagePlanKey{},
		&AwsApiGatewayUsagePlanKeyList{},

		&AwsDxHostedPrivateVirtualInterfaceAccepter{},
		&AwsDxHostedPrivateVirtualInterfaceAccepterList{},

		&AwsServiceDiscoveryHttpNamespace{},
		&AwsServiceDiscoveryHttpNamespaceList{},

		&AwsSnsTopic{},
		&AwsSnsTopicList{},

		&AwsWafregionalGeoMatchSet{},
		&AwsWafregionalGeoMatchSetList{},

		&AwsEc2ClientVpnEndpoint{},
		&AwsEc2ClientVpnEndpointList{},

		&AwsSsmResourceDataSync{},
		&AwsSsmResourceDataSyncList{},

		&AwsCloudwatchLogSubscriptionFilter{},
		&AwsCloudwatchLogSubscriptionFilterList{},

		&AwsDmsCertificate{},
		&AwsDmsCertificateList{},

		&AwsKeyPair{},
		&AwsKeyPairList{},

		&AwsServiceDiscoveryPublicDnsNamespace{},
		&AwsServiceDiscoveryPublicDnsNamespaceList{},

		&AwsSsmActivation{},
		&AwsSsmActivationList{},

		&AwsWafXssMatchSet{},
		&AwsWafXssMatchSetList{},

		&AwsApiGatewayDocumentationPart{},
		&AwsApiGatewayDocumentationPartList{},

		&AwsApiGatewayModel{},
		&AwsApiGatewayModelList{},

		&AwsEmrSecurityConfiguration{},
		&AwsEmrSecurityConfigurationList{},

		&AwsSecurityGroup{},
		&AwsSecurityGroupList{},

		&AwsMacieMemberAccountAssociation{},
		&AwsMacieMemberAccountAssociationList{},

		&AwsSagemakerEndpointConfiguration{},
		&AwsSagemakerEndpointConfigurationList{},

		&AwsVpcEndpointRouteTableAssociation{},
		&AwsVpcEndpointRouteTableAssociationList{},

		&AwsAthenaNamedQuery{},
		&AwsAthenaNamedQueryList{},

		&AwsConfigConfigurationAggregator{},
		&AwsConfigConfigurationAggregatorList{},

		&AwsDocdbClusterInstance{},
		&AwsDocdbClusterInstanceList{},

		&AwsSsmMaintenanceWindowTask{},
		&AwsSsmMaintenanceWindowTaskList{},

		&AwsGuarddutyIpset{},
		&AwsGuarddutyIpsetList{},

		&AwsKmsAlias{},
		&AwsKmsAliasList{},

		&AwsVpnConnection{},
		&AwsVpnConnectionList{},

		&AwsWafRateBasedRule{},
		&AwsWafRateBasedRuleList{},

		&AwsGlobalacceleratorEndpointGroup{},
		&AwsGlobalacceleratorEndpointGroupList{},

		&AwsIotPolicyAttachment{},
		&AwsIotPolicyAttachmentList{},

		&AwsAppautoscalingTarget{},
		&AwsAppautoscalingTargetList{},

		&AwsNetworkInterface{},
		&AwsNetworkInterfaceList{},

		&AwsSesActiveReceiptRuleSet{},
		&AwsSesActiveReceiptRuleSetList{},

		&AwsGlueClassifier{},
		&AwsGlueClassifierList{},

		&AwsOpsworksGangliaLayer{},
		&AwsOpsworksGangliaLayerList{},

		&AwsSwfDomain{},
		&AwsSwfDomainList{},

		&AwsTransferUser{},
		&AwsTransferUserList{},

		&AwsAlb{},
		&AwsAlbList{},

		&AwsKinesisStream{},
		&AwsKinesisStreamList{},

		&AwsLambdaLayerVersion{},
		&AwsLambdaLayerVersionList{},

		&AwsSsmMaintenanceWindow{},
		&AwsSsmMaintenanceWindowList{},

		&AwsCloudhsmV2Cluster{},
		&AwsCloudhsmV2ClusterList{},

		&AwsLightsailStaticIpAttachment{},
		&AwsLightsailStaticIpAttachmentList{},

		&AwsIamPolicyAttachment{},
		&AwsIamPolicyAttachmentList{},

		&AwsIamServerCertificate{},
		&AwsIamServerCertificateList{},

		&AwsApiGatewayBasePathMapping{},
		&AwsApiGatewayBasePathMappingList{},

		&AwsEip{},
		&AwsEipList{},

		&AwsDirectoryServiceLogSubscription{},
		&AwsDirectoryServiceLogSubscriptionList{},

		&AwsGuarddutyMember{},
		&AwsGuarddutyMemberList{},

		&AwsApiGatewayIntegration{},
		&AwsApiGatewayIntegrationList{},

		&AwsCloudfrontPublicKey{},
		&AwsCloudfrontPublicKeyList{},

		&AwsAppmeshVirtualService{},
		&AwsAppmeshVirtualServiceList{},

		&AwsSnsSmsPreferences{},
		&AwsSnsSmsPreferencesList{},

		&AwsAutoscalingPolicy{},
		&AwsAutoscalingPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
