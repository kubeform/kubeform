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
    
        &AwsVpnConnection{},
        &AwsVpnConnectionList{},
    
        &AwsCognitoIdentityProvider{},
        &AwsCognitoIdentityProviderList{},
    
        &AwsDmsReplicationSubnetGroup{},
        &AwsDmsReplicationSubnetGroupList{},
    
        &AwsSesDomainMailFrom{},
        &AwsSesDomainMailFromList{},
    
        &AwsCloudformationStackSetInstance{},
        &AwsCloudformationStackSetInstanceList{},
    
        &AwsNetworkAcl{},
        &AwsNetworkAclList{},
    
        &AwsAlb{},
        &AwsAlbList{},
    
        &AwsCloudwatchEventRule{},
        &AwsCloudwatchEventRuleList{},
    
        &AwsCloudwatchLogMetricFilter{},
        &AwsCloudwatchLogMetricFilterList{},
    
        &AwsOpsworksUserProfile{},
        &AwsOpsworksUserProfileList{},
    
        &AwsEbsVolume{},
        &AwsEbsVolumeList{},
    
        &AwsGlueCatalogDatabase{},
        &AwsGlueCatalogDatabaseList{},
    
        &AwsOrganizationsAccount{},
        &AwsOrganizationsAccountList{},
    
        &AwsVpcEndpointServiceAllowedPrincipal{},
        &AwsVpcEndpointServiceAllowedPrincipalList{},
    
        &AwsWafWebAcl{},
        &AwsWafWebAclList{},
    
        &AwsBatchJobQueue{},
        &AwsBatchJobQueueList{},
    
        &AwsDevicefarmProject{},
        &AwsDevicefarmProjectList{},
    
        &AwsDocdbClusterInstance{},
        &AwsDocdbClusterInstanceList{},
    
        &AwsDbEventSubscription{},
        &AwsDbEventSubscriptionList{},
    
        &AwsEc2ClientVpnNetworkAssociation{},
        &AwsEc2ClientVpnNetworkAssociationList{},
    
        &AwsVpcEndpointSubnetAssociation{},
        &AwsVpcEndpointSubnetAssociationList{},
    
        &AwsWafRule{},
        &AwsWafRuleList{},
    
        &AwsConfigConfigurationRecorderStatus{},
        &AwsConfigConfigurationRecorderStatusList{},
    
        &AwsDatasyncAgent{},
        &AwsDatasyncAgentList{},
    
        &AwsOpsworksStaticWebLayer{},
        &AwsOpsworksStaticWebLayerList{},
    
        &AwsS3BucketPolicy{},
        &AwsS3BucketPolicyList{},
    
        &AwsVpcEndpoint{},
        &AwsVpcEndpointList{},
    
        &AwsElastictranscoderPreset{},
        &AwsElastictranscoderPresetList{},
    
        &AwsInstance{},
        &AwsInstanceList{},
    
        &AwsCodedeployDeploymentConfig{},
        &AwsCodedeployDeploymentConfigList{},
    
        &AwsIotPolicyAttachment{},
        &AwsIotPolicyAttachmentList{},
    
        &AwsRedshiftSubnetGroup{},
        &AwsRedshiftSubnetGroupList{},
    
        &AwsS3BucketPublicAccessBlock{},
        &AwsS3BucketPublicAccessBlockList{},
    
        &AwsApiGatewayDomainName{},
        &AwsApiGatewayDomainNameList{},
    
        &AwsAppsyncGraphqlApi{},
        &AwsAppsyncGraphqlApiList{},
    
        &AwsSsmPatchGroup{},
        &AwsSsmPatchGroupList{},
    
        &AwsStoragegatewayCachedIscsiVolume{},
        &AwsStoragegatewayCachedIscsiVolumeList{},
    
        &AwsSnsTopicSubscription{},
        &AwsSnsTopicSubscriptionList{},
    
        &AwsAppmeshMesh{},
        &AwsAppmeshMeshList{},
    
        &AwsCloudwatchLogDestinationPolicy{},
        &AwsCloudwatchLogDestinationPolicyList{},
    
        &AwsRoute53ZoneAssociation{},
        &AwsRoute53ZoneAssociationList{},
    
        &AwsSesDomainIdentityVerification{},
        &AwsSesDomainIdentityVerificationList{},
    
        &AwsGlueSecurityConfiguration{},
        &AwsGlueSecurityConfigurationList{},
    
        &AwsIamPolicy{},
        &AwsIamPolicyList{},
    
        &AwsKmsAlias{},
        &AwsKmsAliasList{},
    
        &AwsSecretsmanagerSecret{},
        &AwsSecretsmanagerSecretList{},
    
        &AwsS3BucketMetric{},
        &AwsS3BucketMetricList{},
    
        &AwsAlbTargetGroupAttachment{},
        &AwsAlbTargetGroupAttachmentList{},
    
        &AwsDocdbClusterParameterGroup{},
        &AwsDocdbClusterParameterGroupList{},
    
        &AwsIamGroupPolicy{},
        &AwsIamGroupPolicyList{},
    
        &AwsWafregionalByteMatchSet{},
        &AwsWafregionalByteMatchSetList{},
    
        &AwsGlobalacceleratorEndpointGroup{},
        &AwsGlobalacceleratorEndpointGroupList{},
    
        &AwsLambdaFunction{},
        &AwsLambdaFunctionList{},
    
        &AwsDocdbCluster{},
        &AwsDocdbClusterList{},
    
        &AwsElasticBeanstalkEnvironment{},
        &AwsElasticBeanstalkEnvironmentList{},
    
        &AwsSqsQueuePolicy{},
        &AwsSqsQueuePolicyList{},
    
        &AwsSfnStateMachine{},
        &AwsSfnStateMachineList{},
    
        &AwsVpcEndpointService{},
        &AwsVpcEndpointServiceList{},
    
        &AwsApiGatewayMethodSettings{},
        &AwsApiGatewayMethodSettingsList{},
    
        &AwsCloud9EnvironmentEc2{},
        &AwsCloud9EnvironmentEc2List{},
    
        &AwsCognitoUserPoolDomain{},
        &AwsCognitoUserPoolDomainList{},
    
        &AwsMacieS3BucketAssociation{},
        &AwsMacieS3BucketAssociationList{},
    
        &AwsIamUserSshKey{},
        &AwsIamUserSshKeyList{},
    
        &AwsIamUserLoginProfile{},
        &AwsIamUserLoginProfileList{},
    
        &AwsKinesisAnalyticsApplication{},
        &AwsKinesisAnalyticsApplicationList{},
    
        &AwsSnsTopicPolicy{},
        &AwsSnsTopicPolicyList{},
    
        &AwsCognitoResourceServer{},
        &AwsCognitoResourceServerList{},
    
        &AwsGlueConnection{},
        &AwsGlueConnectionList{},
    
        &AwsEbsEncryptionByDefault{},
        &AwsEbsEncryptionByDefaultList{},
    
        &AwsElasticBeanstalkConfigurationTemplate{},
        &AwsElasticBeanstalkConfigurationTemplateList{},
    
        &AwsVpnConnectionRoute{},
        &AwsVpnConnectionRouteList{},
    
        &AwsInternetGateway{},
        &AwsInternetGatewayList{},
    
        &AwsRoute53QueryLog{},
        &AwsRoute53QueryLogList{},
    
        &AwsS3BucketInventory{},
        &AwsS3BucketInventoryList{},
    
        &AwsAppautoscalingPolicy{},
        &AwsAppautoscalingPolicyList{},
    
        &AwsIamAccountAlias{},
        &AwsIamAccountAliasList{},
    
        &AwsCodebuildProject{},
        &AwsCodebuildProjectList{},
    
        &AwsGuarddutyThreatintelset{},
        &AwsGuarddutyThreatintelsetList{},
    
        &AwsNeptuneClusterSnapshot{},
        &AwsNeptuneClusterSnapshotList{},
    
        &AwsConfigAggregateAuthorization{},
        &AwsConfigAggregateAuthorizationList{},
    
        &AwsCloudwatchDashboard{},
        &AwsCloudwatchDashboardList{},
    
        &AwsEc2TransitGatewayRouteTablePropagation{},
        &AwsEc2TransitGatewayRouteTablePropagationList{},
    
        &AwsNeptuneSubnetGroup{},
        &AwsNeptuneSubnetGroupList{},
    
        &AwsBackupSelection{},
        &AwsBackupSelectionList{},
    
        &AwsCognitoUserPool{},
        &AwsCognitoUserPoolList{},
    
        &AwsEc2TransitGateway{},
        &AwsEc2TransitGatewayList{},
    
        &AwsWafregionalWebAcl{},
        &AwsWafregionalWebAclList{},
    
        &AwsAutoscalingPolicy{},
        &AwsAutoscalingPolicyList{},
    
        &AwsDlmLifecyclePolicy{},
        &AwsDlmLifecyclePolicyList{},
    
        &AwsSsmActivation{},
        &AwsSsmActivationList{},
    
        &AwsPinpointAdmChannel{},
        &AwsPinpointAdmChannelList{},
    
        &AwsXraySamplingRule{},
        &AwsXraySamplingRuleList{},
    
        &AwsAlbListener{},
        &AwsAlbListenerList{},
    
        &AwsEc2TransitGatewayRouteTable{},
        &AwsEc2TransitGatewayRouteTableList{},
    
        &AwsSecurityGroup{},
        &AwsSecurityGroupList{},
    
        &AwsElb{},
        &AwsElbList{},
    
        &AwsIamServerCertificate{},
        &AwsIamServerCertificateList{},
    
        &AwsSesActiveReceiptRuleSet{},
        &AwsSesActiveReceiptRuleSetList{},
    
        &AwsCodepipeline{},
        &AwsCodepipelineList{},
    
        &AwsDocdbSubnetGroup{},
        &AwsDocdbSubnetGroupList{},
    
        &AwsSesReceiptRule{},
        &AwsSesReceiptRuleList{},
    
        &AwsDefaultSubnet{},
        &AwsDefaultSubnetList{},
    
        &AwsLbListener{},
        &AwsLbListenerList{},
    
        &AwsLightsailDomain{},
        &AwsLightsailDomainList{},
    
        &AwsRamResourceShare{},
        &AwsRamResourceShareList{},
    
        &AwsAppmeshVirtualService{},
        &AwsAppmeshVirtualServiceList{},
    
        &AwsCodebuildWebhook{},
        &AwsCodebuildWebhookList{},
    
        &AwsSsmPatchBaseline{},
        &AwsSsmPatchBaselineList{},
    
        &AwsApiGatewayMethodResponse{},
        &AwsApiGatewayMethodResponseList{},
    
        &AwsAppCookieStickinessPolicy{},
        &AwsAppCookieStickinessPolicyList{},
    
        &AwsSecurityhubStandardsSubscription{},
        &AwsSecurityhubStandardsSubscriptionList{},
    
        &AwsLbListenerRule{},
        &AwsLbListenerRuleList{},
    
        &AwsRoute53ResolverRuleAssociation{},
        &AwsRoute53ResolverRuleAssociationList{},
    
        &AwsAutoscalingNotification{},
        &AwsAutoscalingNotificationList{},
    
        &AwsLoadBalancerBackendServerPolicy{},
        &AwsLoadBalancerBackendServerPolicyList{},
    
        &AwsDirectoryServiceLogSubscription{},
        &AwsDirectoryServiceLogSubscriptionList{},
    
        &AwsIamPolicyAttachment{},
        &AwsIamPolicyAttachmentList{},
    
        &AwsRdsClusterInstance{},
        &AwsRdsClusterInstanceList{},
    
        &AwsCloudwatchLogDestination{},
        &AwsCloudwatchLogDestinationList{},
    
        &AwsCloudwatchMetricAlarm{},
        &AwsCloudwatchMetricAlarmList{},
    
        &AwsGuarddutyDetector{},
        &AwsGuarddutyDetectorList{},
    
        &AwsWafRegexPatternSet{},
        &AwsWafRegexPatternSetList{},
    
        &AwsWafRuleGroup{},
        &AwsWafRuleGroupList{},
    
        &AwsDatasyncLocationEfs{},
        &AwsDatasyncLocationEfsList{},
    
        &AwsDaxSubnetGroup{},
        &AwsDaxSubnetGroupList{},
    
        &AwsIamGroupPolicyAttachment{},
        &AwsIamGroupPolicyAttachmentList{},
    
        &AwsRoute53DelegationSet{},
        &AwsRoute53DelegationSetList{},
    
        &AwsAcmpcaCertificateAuthority{},
        &AwsAcmpcaCertificateAuthorityList{},
    
        &AwsElbAttachment{},
        &AwsElbAttachmentList{},
    
        &AwsSesIdentityNotificationTopic{},
        &AwsSesIdentityNotificationTopicList{},
    
        &AwsAmi{},
        &AwsAmiList{},
    
        &AwsDbParameterGroup{},
        &AwsDbParameterGroupList{},
    
        &AwsElasticBeanstalkApplication{},
        &AwsElasticBeanstalkApplicationList{},
    
        &AwsApiGatewayApiKey{},
        &AwsApiGatewayApiKeyList{},
    
        &AwsEc2TransitGatewayRoute{},
        &AwsEc2TransitGatewayRouteList{},
    
        &AwsNeptuneEventSubscription{},
        &AwsNeptuneEventSubscriptionList{},
    
        &AwsSsmParameter{},
        &AwsSsmParameterList{},
    
        &AwsDmsEndpoint{},
        &AwsDmsEndpointList{},
    
        &AwsLoadBalancerListenerPolicy{},
        &AwsLoadBalancerListenerPolicyList{},
    
        &AwsSnsPlatformApplication{},
        &AwsSnsPlatformApplicationList{},
    
        &AwsPinpointBaiduChannel{},
        &AwsPinpointBaiduChannelList{},
    
        &AwsNetworkInterfaceAttachment{},
        &AwsNetworkInterfaceAttachmentList{},
    
        &AwsRamResourceAssociation{},
        &AwsRamResourceAssociationList{},
    
        &AwsEfsFileSystem{},
        &AwsEfsFileSystemList{},
    
        &AwsKmsKey{},
        &AwsKmsKeyList{},
    
        &AwsSpotDatafeedSubscription{},
        &AwsSpotDatafeedSubscriptionList{},
    
        &AwsEc2ClientVpnEndpoint{},
        &AwsEc2ClientVpnEndpointList{},
    
        &AwsEc2Fleet{},
        &AwsEc2FleetList{},
    
        &AwsIotRoleAlias{},
        &AwsIotRoleAliasList{},
    
        &AwsNetworkAclRule{},
        &AwsNetworkAclRuleList{},
    
        &AwsSnsSmsPreferences{},
        &AwsSnsSmsPreferencesList{},
    
        &AwsCloudwatchLogStream{},
        &AwsCloudwatchLogStreamList{},
    
        &AwsIamUserGroupMembership{},
        &AwsIamUserGroupMembershipList{},
    
        &AwsDatasyncLocationNfs{},
        &AwsDatasyncLocationNfsList{},
    
        &AwsElasticacheCluster{},
        &AwsElasticacheClusterList{},
    
        &AwsServiceDiscoveryHttpNamespace{},
        &AwsServiceDiscoveryHttpNamespaceList{},
    
        &AwsVpcPeeringConnectionOptions{},
        &AwsVpcPeeringConnectionOptionsList{},
    
        &AwsApiGatewayRestApi{},
        &AwsApiGatewayRestApiList{},
    
        &AwsAutoscalingAttachment{},
        &AwsAutoscalingAttachmentList{},
    
        &AwsSecurityhubProductSubscription{},
        &AwsSecurityhubProductSubscriptionList{},
    
        &AwsApiGatewayAccount{},
        &AwsApiGatewayAccountList{},
    
        &AwsApiGatewayDocumentationVersion{},
        &AwsApiGatewayDocumentationVersionList{},
    
        &AwsSsmMaintenanceWindow{},
        &AwsSsmMaintenanceWindowList{},
    
        &AwsAppautoscalingScheduledAction{},
        &AwsAppautoscalingScheduledActionList{},
    
        &AwsSecurityGroupRule{},
        &AwsSecurityGroupRuleList{},
    
        &AwsWafregionalSqlInjectionMatchSet{},
        &AwsWafregionalSqlInjectionMatchSetList{},
    
        &AwsEcsTaskDefinition{},
        &AwsEcsTaskDefinitionList{},
    
        &AwsGameliftBuild{},
        &AwsGameliftBuildList{},
    
        &AwsWafregionalRegexPatternSet{},
        &AwsWafregionalRegexPatternSetList{},
    
        &AwsGameliftAlias{},
        &AwsGameliftAliasList{},
    
        &AwsS3BucketNotification{},
        &AwsS3BucketNotificationList{},
    
        &AwsOpsworksApplication{},
        &AwsOpsworksApplicationList{},
    
        &AwsLbTargetGroupAttachment{},
        &AwsLbTargetGroupAttachmentList{},
    
        &AwsCloudhsmV2Hsm{},
        &AwsCloudhsmV2HsmList{},
    
        &AwsCurReportDefinition{},
        &AwsCurReportDefinitionList{},
    
        &AwsIotTopicRule{},
        &AwsIotTopicRuleList{},
    
        &AwsCloudhsmV2Cluster{},
        &AwsCloudhsmV2ClusterList{},
    
        &AwsEcrRepositoryPolicy{},
        &AwsEcrRepositoryPolicyList{},
    
        &AwsEipAssociation{},
        &AwsEipAssociationList{},
    
        &AwsGlueCrawler{},
        &AwsGlueCrawlerList{},
    
        &AwsApiGatewayGatewayResponse{},
        &AwsApiGatewayGatewayResponseList{},
    
        &AwsApiGatewayMethod{},
        &AwsApiGatewayMethodList{},
    
        &AwsRoute53ResolverEndpoint{},
        &AwsRoute53ResolverEndpointList{},
    
        &AwsSwfDomain{},
        &AwsSwfDomainList{},
    
        &AwsWafByteMatchSet{},
        &AwsWafByteMatchSetList{},
    
        &AwsWafSqlInjectionMatchSet{},
        &AwsWafSqlInjectionMatchSetList{},
    
        &AwsEc2TransitGatewayVpcAttachmentAccepter{},
        &AwsEc2TransitGatewayVpcAttachmentAccepterList{},
    
        &AwsGuarddutyInviteAccepter{},
        &AwsGuarddutyInviteAccepterList{},
    
        &AwsLoadBalancerPolicy{},
        &AwsLoadBalancerPolicyList{},
    
        &AwsOrganizationsPolicyAttachment{},
        &AwsOrganizationsPolicyAttachmentList{},
    
        &AwsVpcIpv4CidrBlockAssociation{},
        &AwsVpcIpv4CidrBlockAssociationList{},
    
        &AwsPinpointEmailChannel{},
        &AwsPinpointEmailChannelList{},
    
        &AwsElasticsearchDomainPolicy{},
        &AwsElasticsearchDomainPolicyList{},
    
        &AwsGuarddutyIpset{},
        &AwsGuarddutyIpsetList{},
    
        &AwsVpcDhcpOptionsAssociation{},
        &AwsVpcDhcpOptionsAssociationList{},
    
        &AwsWafGeoMatchSet{},
        &AwsWafGeoMatchSetList{},
    
        &AwsAcmCertificateValidation{},
        &AwsAcmCertificateValidationList{},
    
        &AwsSnsTopic{},
        &AwsSnsTopicList{},
    
        &AwsOpsworksStack{},
        &AwsOpsworksStackList{},
    
        &AwsTransferServer{},
        &AwsTransferServerList{},
    
        &AwsApiGatewayUsagePlanKey{},
        &AwsApiGatewayUsagePlanKeyList{},
    
        &AwsDirectoryServiceDirectory{},
        &AwsDirectoryServiceDirectoryList{},
    
        &AwsOpsworksCustomLayer{},
        &AwsOpsworksCustomLayerList{},
    
        &AwsVolumeAttachment{},
        &AwsVolumeAttachmentList{},
    
        &AwsVpnGatewayRoutePropagation{},
        &AwsVpnGatewayRoutePropagationList{},
    
        &AwsDxConnection{},
        &AwsDxConnectionList{},
    
        &AwsLaunchConfiguration{},
        &AwsLaunchConfigurationList{},
    
        &AwsKmsCiphertext{},
        &AwsKmsCiphertextList{},
    
        &AwsCustomerGateway{},
        &AwsCustomerGatewayList{},
    
        &AwsIamRolePolicyAttachment{},
        &AwsIamRolePolicyAttachmentList{},
    
        &AwsGlobalacceleratorAccelerator{},
        &AwsGlobalacceleratorAcceleratorList{},
    
        &AwsNeptuneParameterGroup{},
        &AwsNeptuneParameterGroupList{},
    
        &AwsRoute53Zone{},
        &AwsRoute53ZoneList{},
    
        &AwsSesEventDestination{},
        &AwsSesEventDestinationList{},
    
        &AwsWafXssMatchSet{},
        &AwsWafXssMatchSetList{},
    
        &AwsCognitoIdentityPool{},
        &AwsCognitoIdentityPoolList{},
    
        &AwsDbSnapshot{},
        &AwsDbSnapshotList{},
    
        &AwsDatasyncLocationS3{},
        &AwsDatasyncLocationS3List{},
    
        &AwsEbsSnapshot{},
        &AwsEbsSnapshotList{},
    
        &AwsEcsCluster{},
        &AwsEcsClusterList{},
    
        &AwsElasticacheReplicationGroup{},
        &AwsElasticacheReplicationGroupList{},
    
        &AwsElasticacheSecurityGroup{},
        &AwsElasticacheSecurityGroupList{},
    
        &AwsNeptuneCluster{},
        &AwsNeptuneClusterList{},
    
        &AwsAppmeshVirtualRouter{},
        &AwsAppmeshVirtualRouterList{},
    
        &AwsConfigConfigurationAggregator{},
        &AwsConfigConfigurationAggregatorList{},
    
        &AwsDmsReplicationInstance{},
        &AwsDmsReplicationInstanceList{},
    
        &AwsWorklinkWebsiteCertificateAuthorityAssociation{},
        &AwsWorklinkWebsiteCertificateAuthorityAssociationList{},
    
        &AwsPinpointApnsChannel{},
        &AwsPinpointApnsChannelList{},
    
        &AwsDbInstance{},
        &AwsDbInstanceList{},
    
        &AwsDmsCertificate{},
        &AwsDmsCertificateList{},
    
        &AwsEc2CapacityReservation{},
        &AwsEc2CapacityReservationList{},
    
        &AwsLambdaEventSourceMapping{},
        &AwsLambdaEventSourceMappingList{},
    
        &AwsApiGatewayAuthorizer{},
        &AwsApiGatewayAuthorizerList{},
    
        &AwsAutoscalingGroup{},
        &AwsAutoscalingGroupList{},
    
        &AwsSagemakerModel{},
        &AwsSagemakerModelList{},
    
        &AwsSagemakerNotebookInstanceLifecycleConfiguration{},
        &AwsSagemakerNotebookInstanceLifecycleConfigurationList{},
    
        &AwsWafregionalRateBasedRule{},
        &AwsWafregionalRateBasedRuleList{},
    
        &AwsBatchJobDefinition{},
        &AwsBatchJobDefinitionList{},
    
        &AwsIamInstanceProfile{},
        &AwsIamInstanceProfileList{},
    
        &AwsPlacementGroup{},
        &AwsPlacementGroupList{},
    
        &AwsCodecommitRepository{},
        &AwsCodecommitRepositoryList{},
    
        &AwsGuarddutyMember{},
        &AwsGuarddutyMemberList{},
    
        &AwsResourcegroupsGroup{},
        &AwsResourcegroupsGroupList{},
    
        &AwsPinpointApnsVoipChannel{},
        &AwsPinpointApnsVoipChannelList{},
    
        &AwsCloudformationStackSet{},
        &AwsCloudformationStackSetList{},
    
        &AwsCloudwatchEventTarget{},
        &AwsCloudwatchEventTargetList{},
    
        &AwsEfsMountTarget{},
        &AwsEfsMountTargetList{},
    
        &AwsOpsworksPhpAppLayer{},
        &AwsOpsworksPhpAppLayerList{},
    
        &AwsSesDomainIdentity{},
        &AwsSesDomainIdentityList{},
    
        &AwsSubnet{},
        &AwsSubnetList{},
    
        &AwsPinpointApp{},
        &AwsPinpointAppList{},
    
        &AwsAppmeshRoute{},
        &AwsAppmeshRouteList{},
    
        &AwsBackupVault{},
        &AwsBackupVaultList{},
    
        &AwsDmsReplicationTask{},
        &AwsDmsReplicationTaskList{},
    
        &AwsEc2TransitGatewayRouteTableAssociation{},
        &AwsEc2TransitGatewayRouteTableAssociationList{},
    
        &AwsKinesisFirehoseDeliveryStream{},
        &AwsKinesisFirehoseDeliveryStreamList{},
    
        &AwsDefaultNetworkAcl{},
        &AwsDefaultNetworkAclList{},
    
        &AwsRedshiftCluster{},
        &AwsRedshiftClusterList{},
    
        &AwsCloudfrontOriginAccessIdentity{},
        &AwsCloudfrontOriginAccessIdentityList{},
    
        &AwsConfigConfigurationRecorder{},
        &AwsConfigConfigurationRecorderList{},
    
        &AwsBudgetsBudget{},
        &AwsBudgetsBudgetList{},
    
        &AwsDxConnectionAssociation{},
        &AwsDxConnectionAssociationList{},
    
        &AwsMacieMemberAccountAssociation{},
        &AwsMacieMemberAccountAssociationList{},
    
        &AwsDefaultVpc{},
        &AwsDefaultVpcList{},
    
        &AwsApiGatewayStage{},
        &AwsApiGatewayStageList{},
    
        &AwsApiGatewayVpcLink{},
        &AwsApiGatewayVpcLinkList{},
    
        &AwsNeptuneClusterParameterGroup{},
        &AwsNeptuneClusterParameterGroupList{},
    
        &AwsDefaultRouteTable{},
        &AwsDefaultRouteTableList{},
    
        &AwsSesReceiptFilter{},
        &AwsSesReceiptFilterList{},
    
        &AwsVpcEndpointConnectionNotification{},
        &AwsVpcEndpointConnectionNotificationList{},
    
        &AwsVpnGatewayAttachment{},
        &AwsVpnGatewayAttachmentList{},
    
        &AwsAppautoscalingTarget{},
        &AwsAppautoscalingTargetList{},
    
        &AwsCodecommitTrigger{},
        &AwsCodecommitTriggerList{},
    
        &AwsDefaultSecurityGroup{},
        &AwsDefaultSecurityGroupList{},
    
        &AwsWafregionalXssMatchSet{},
        &AwsWafregionalXssMatchSetList{},
    
        &AwsPinpointApnsSandboxChannel{},
        &AwsPinpointApnsSandboxChannelList{},
    
        &AwsAlbTargetGroup{},
        &AwsAlbTargetGroupList{},
    
        &AwsGlueTrigger{},
        &AwsGlueTriggerList{},
    
        &AwsSagemakerNotebookInstance{},
        &AwsSagemakerNotebookInstanceList{},
    
        &AwsLambdaAlias{},
        &AwsLambdaAliasList{},
    
        &AwsOpsworksRailsAppLayer{},
        &AwsOpsworksRailsAppLayerList{},
    
        &AwsLb{},
        &AwsLbList{},
    
        &AwsCodepipelineWebhook{},
        &AwsCodepipelineWebhookList{},
    
        &AwsIamUser{},
        &AwsIamUserList{},
    
        &AwsIamAccessKey{},
        &AwsIamAccessKeyList{},
    
        &AwsAppsyncApiKey{},
        &AwsAppsyncApiKeyList{},
    
        &AwsDxHostedPublicVirtualInterface{},
        &AwsDxHostedPublicVirtualInterfaceList{},
    
        &AwsShieldProtection{},
        &AwsShieldProtectionList{},
    
        &AwsSqsQueue{},
        &AwsSqsQueueList{},
    
        &AwsElasticBeanstalkApplicationVersion{},
        &AwsElasticBeanstalkApplicationVersionList{},
    
        &AwsRoute53HealthCheck{},
        &AwsRoute53HealthCheckList{},
    
        &AwsEgressOnlyInternetGateway{},
        &AwsEgressOnlyInternetGatewayList{},
    
        &AwsLightsailInstance{},
        &AwsLightsailInstanceList{},
    
        &AwsRouteTable{},
        &AwsRouteTableList{},
    
        &AwsSagemakerEndpointConfiguration{},
        &AwsSagemakerEndpointConfigurationList{},
    
        &AwsApiGatewayClientCertificate{},
        &AwsApiGatewayClientCertificateList{},
    
        &AwsCloudwatchLogSubscriptionFilter{},
        &AwsCloudwatchLogSubscriptionFilterList{},
    
        &AwsLightsailStaticIpAttachment{},
        &AwsLightsailStaticIpAttachmentList{},
    
        &AwsSimpledbDomain{},
        &AwsSimpledbDomainList{},
    
        &AwsEc2TransitGatewayVpcAttachment{},
        &AwsEc2TransitGatewayVpcAttachmentList{},
    
        &AwsLbListenerCertificate{},
        &AwsLbListenerCertificateList{},
    
        &AwsAlbListenerRule{},
        &AwsAlbListenerRuleList{},
    
        &AwsApiGatewayUsagePlan{},
        &AwsApiGatewayUsagePlanList{},
    
        &AwsAppsyncResolver{},
        &AwsAppsyncResolverList{},
    
        &AwsGlobalacceleratorListener{},
        &AwsGlobalacceleratorListenerList{},
    
        &AwsOpsworksGangliaLayer{},
        &AwsOpsworksGangliaLayerList{},
    
        &AwsWorklinkFleet{},
        &AwsWorklinkFleetList{},
    
        &AwsApiGatewayIntegrationResponse{},
        &AwsApiGatewayIntegrationResponseList{},
    
        &AwsAutoscalingSchedule{},
        &AwsAutoscalingScheduleList{},
    
        &AwsNetworkInterface{},
        &AwsNetworkInterfaceList{},
    
        &AwsDbOptionGroup{},
        &AwsDbOptionGroupList{},
    
        &AwsEksCluster{},
        &AwsEksClusterList{},
    
        &AwsKeyPair{},
        &AwsKeyPairList{},
    
        &AwsStoragegatewayUploadBuffer{},
        &AwsStoragegatewayUploadBufferList{},
    
        &AwsSecurityhubAccount{},
        &AwsSecurityhubAccountList{},
    
        &AwsWafRateBasedRule{},
        &AwsWafRateBasedRuleList{},
    
        &AwsAthenaNamedQuery{},
        &AwsAthenaNamedQueryList{},
    
        &AwsEcrLifecyclePolicy{},
        &AwsEcrLifecyclePolicyList{},
    
        &AwsMqConfiguration{},
        &AwsMqConfigurationList{},
    
        &AwsEmrInstanceGroup{},
        &AwsEmrInstanceGroupList{},
    
        &AwsIamUserPolicy{},
        &AwsIamUserPolicyList{},
    
        &AwsIamGroup{},
        &AwsIamGroupList{},
    
        &AwsDatapipelinePipeline{},
        &AwsDatapipelinePipelineList{},
    
        &AwsDynamodbTable{},
        &AwsDynamodbTableList{},
    
        &AwsEcrRepository{},
        &AwsEcrRepositoryList{},
    
        &AwsEmrSecurityConfiguration{},
        &AwsEmrSecurityConfigurationList{},
    
        &AwsInspectorAssessmentTarget{},
        &AwsInspectorAssessmentTargetList{},
    
        &AwsServicecatalogPortfolio{},
        &AwsServicecatalogPortfolioList{},
    
        &AwsCognitoIdentityPoolRolesAttachment{},
        &AwsCognitoIdentityPoolRolesAttachmentList{},
    
        &AwsDocdbClusterSnapshot{},
        &AwsDocdbClusterSnapshotList{},
    
        &AwsRdsClusterEndpoint{},
        &AwsRdsClusterEndpointList{},
    
        &AwsAmiFromInstance{},
        &AwsAmiFromInstanceList{},
    
        &AwsMediaStoreContainerPolicy{},
        &AwsMediaStoreContainerPolicyList{},
    
        &AwsApiGatewayModel{},
        &AwsApiGatewayModelList{},
    
        &AwsPinpointGcmChannel{},
        &AwsPinpointGcmChannelList{},
    
        &AwsTransferSshKey{},
        &AwsTransferSshKeyList{},
    
        &AwsCloudformationStack{},
        &AwsCloudformationStackList{},
    
        &AwsInspectorAssessmentTemplate{},
        &AwsInspectorAssessmentTemplateList{},
    
        &AwsApiGatewayDeployment{},
        &AwsApiGatewayDeploymentList{},
    
        &AwsServiceDiscoveryPrivateDnsNamespace{},
        &AwsServiceDiscoveryPrivateDnsNamespaceList{},
    
        &AwsLambdaPermission{},
        &AwsLambdaPermissionList{},
    
        &AwsOrganizationsOrganization{},
        &AwsOrganizationsOrganizationList{},
    
        &AwsStoragegatewayGateway{},
        &AwsStoragegatewayGatewayList{},
    
        &AwsSpotFleetRequest{},
        &AwsSpotFleetRequestList{},
    
        &AwsWafRegexMatchSet{},
        &AwsWafRegexMatchSetList{},
    
        &AwsWafregionalIpset{},
        &AwsWafregionalIpsetList{},
    
        &AwsAppsyncDatasource{},
        &AwsAppsyncDatasourceList{},
    
        &AwsDxHostedPrivateVirtualInterface{},
        &AwsDxHostedPrivateVirtualInterfaceList{},
    
        &AwsWafregionalSizeConstraintSet{},
        &AwsWafregionalSizeConstraintSetList{},
    
        &AwsCognitoUserGroup{},
        &AwsCognitoUserGroupList{},
    
        &AwsProxyProtocolPolicy{},
        &AwsProxyProtocolPolicyList{},
    
        &AwsInspectorResourceGroup{},
        &AwsInspectorResourceGroupList{},
    
        &AwsWafregionalRegexMatchSet{},
        &AwsWafregionalRegexMatchSetList{},
    
        &AwsWafregionalRuleGroup{},
        &AwsWafregionalRuleGroupList{},
    
        &AwsEbsDefaultKmsKey{},
        &AwsEbsDefaultKmsKeyList{},
    
        &AwsOpsworksJavaAppLayer{},
        &AwsOpsworksJavaAppLayerList{},
    
        &AwsMediaPackageChannel{},
        &AwsMediaPackageChannelList{},
    
        &AwsNeptuneClusterInstance{},
        &AwsNeptuneClusterInstanceList{},
    
        &AwsS3Bucket{},
        &AwsS3BucketList{},
    
        &AwsVpnGateway{},
        &AwsVpnGatewayList{},
    
        &AwsDxPublicVirtualInterface{},
        &AwsDxPublicVirtualInterfaceList{},
    
        &AwsLaunchTemplate{},
        &AwsLaunchTemplateList{},
    
        &AwsAutoscalingLifecycleHook{},
        &AwsAutoscalingLifecycleHookList{},
    
        &AwsLbCookieStickinessPolicy{},
        &AwsLbCookieStickinessPolicyList{},
    
        &AwsRoute{},
        &AwsRouteList{},
    
        &AwsSecretsmanagerSecretVersion{},
        &AwsSecretsmanagerSecretVersionList{},
    
        &AwsSsmMaintenanceWindowTask{},
        &AwsSsmMaintenanceWindowTaskList{},
    
        &AwsDefaultVpcDhcpOptions{},
        &AwsDefaultVpcDhcpOptionsList{},
    
        &AwsAmiLaunchPermission{},
        &AwsAmiLaunchPermissionList{},
    
        &AwsApiGatewayResource{},
        &AwsApiGatewayResourceList{},
    
        &AwsWafregionalGeoMatchSet{},
        &AwsWafregionalGeoMatchSetList{},
    
        &AwsWafregionalWebAclAssociation{},
        &AwsWafregionalWebAclAssociationList{},
    
        &AwsApiGatewayDocumentationPart{},
        &AwsApiGatewayDocumentationPartList{},
    
        &AwsLambdaLayerVersion{},
        &AwsLambdaLayerVersionList{},
    
        &AwsCloudfrontDistribution{},
        &AwsCloudfrontDistributionList{},
    
        &AwsSsmAssociation{},
        &AwsSsmAssociationList{},
    
        &AwsSfnActivity{},
        &AwsSfnActivityList{},
    
        &AwsAcmCertificate{},
        &AwsAcmCertificateList{},
    
        &AwsAthenaWorkgroup{},
        &AwsAthenaWorkgroupList{},
    
        &AwsLightsailStaticIp{},
        &AwsLightsailStaticIpList{},
    
        &AwsIotPolicy{},
        &AwsIotPolicyList{},
    
        &AwsIotThing{},
        &AwsIotThingList{},
    
        &AwsCodedeployApp{},
        &AwsCodedeployAppList{},
    
        &AwsSesTemplate{},
        &AwsSesTemplateList{},
    
        &AwsSsmMaintenanceWindowTarget{},
        &AwsSsmMaintenanceWindowTargetList{},
    
        &AwsElasticsearchDomain{},
        &AwsElasticsearchDomainList{},
    
        &AwsGlueClassifier{},
        &AwsGlueClassifierList{},
    
        &AwsSsmResourceDataSync{},
        &AwsSsmResourceDataSyncList{},
    
        &AwsLbTargetGroup{},
        &AwsLbTargetGroupList{},
    
        &AwsEcsService{},
        &AwsEcsServiceList{},
    
        &AwsGlueCatalogTable{},
        &AwsGlueCatalogTableList{},
    
        &AwsRdsClusterParameterGroup{},
        &AwsRdsClusterParameterGroupList{},
    
        &AwsVpcPeeringConnection{},
        &AwsVpcPeeringConnectionList{},
    
        &AwsWafregionalRule{},
        &AwsWafregionalRuleList{},
    
        &AwsDxPrivateVirtualInterface{},
        &AwsDxPrivateVirtualInterfaceList{},
    
        &AwsDynamodbTableItem{},
        &AwsDynamodbTableItemList{},
    
        &AwsSesIdentityPolicy{},
        &AwsSesIdentityPolicyList{},
    
        &AwsAlbListenerCertificate{},
        &AwsAlbListenerCertificateList{},
    
        &AwsElasticacheSubnetGroup{},
        &AwsElasticacheSubnetGroupList{},
    
        &AwsRoute53Record{},
        &AwsRoute53RecordList{},
    
        &AwsDbSubnetGroup{},
        &AwsDbSubnetGroupList{},
    
        &AwsRdsCluster{},
        &AwsRdsClusterList{},
    
        &AwsAppsyncFunction{},
        &AwsAppsyncFunctionList{},
    
        &AwsCognitoUserPoolClient{},
        &AwsCognitoUserPoolClientList{},
    
        &AwsIotThingPrincipalAttachment{},
        &AwsIotThingPrincipalAttachmentList{},
    
        &AwsRedshiftParameterGroup{},
        &AwsRedshiftParameterGroupList{},
    
        &AwsDxGatewayAssociation{},
        &AwsDxGatewayAssociationList{},
    
        &AwsElasticacheParameterGroup{},
        &AwsElasticacheParameterGroupList{},
    
        &AwsSesEmailIdentity{},
        &AwsSesEmailIdentityList{},
    
        &AwsDaxCluster{},
        &AwsDaxClusterList{},
    
        &AwsDbClusterSnapshot{},
        &AwsDbClusterSnapshotList{},
    
        &AwsEmrCluster{},
        &AwsEmrClusterList{},
    
        &AwsGlacierVault{},
        &AwsGlacierVaultList{},
    
        &AwsStoragegatewayCache{},
        &AwsStoragegatewayCacheList{},
    
        &AwsDirectoryServiceConditionalForwarder{},
        &AwsDirectoryServiceConditionalForwarderList{},
    
        &AwsServiceDiscoveryService{},
        &AwsServiceDiscoveryServiceList{},
    
        &AwsIamServiceLinkedRole{},
        &AwsIamServiceLinkedRoleList{},
    
        &AwsPinpointEventStream{},
        &AwsPinpointEventStreamList{},
    
        &AwsKinesisStream{},
        &AwsKinesisStreamList{},
    
        &AwsMskCluster{},
        &AwsMskClusterList{},
    
        &AwsVpcDhcpOptions{},
        &AwsVpcDhcpOptionsList{},
    
        &AwsVpcEndpointRouteTableAssociation{},
        &AwsVpcEndpointRouteTableAssociationList{},
    
        &AwsCloudfrontPublicKey{},
        &AwsCloudfrontPublicKeyList{},
    
        &AwsIamOpenidConnectProvider{},
        &AwsIamOpenidConnectProviderList{},
    
        &AwsCodedeployDeploymentGroup{},
        &AwsCodedeployDeploymentGroupList{},
    
        &AwsWafIpset{},
        &AwsWafIpsetList{},
    
        &AwsElastictranscoderPipeline{},
        &AwsElastictranscoderPipelineList{},
    
        &AwsIamGroupMembership{},
        &AwsIamGroupMembershipList{},
    
        &AwsWafSizeConstraintSet{},
        &AwsWafSizeConstraintSetList{},
    
        &AwsCloudtrail{},
        &AwsCloudtrailList{},
    
        &AwsEip{},
        &AwsEipList{},
    
        &AwsIamRolePolicy{},
        &AwsIamRolePolicyList{},
    
        &AwsLightsailKeyPair{},
        &AwsLightsailKeyPairList{},
    
        &AwsOpsworksInstance{},
        &AwsOpsworksInstanceList{},
    
        &AwsOrganizationsPolicy{},
        &AwsOrganizationsPolicyList{},
    
        &AwsOrganizationsOrganizationalUnit{},
        &AwsOrganizationsOrganizationalUnitList{},
    
        &AwsStoragegatewayNfsFileShare{},
        &AwsStoragegatewayNfsFileShareList{},
    
        &AwsCloudwatchLogResourcePolicy{},
        &AwsCloudwatchLogResourcePolicyList{},
    
        &AwsDbInstanceRoleAssociation{},
        &AwsDbInstanceRoleAssociationList{},
    
        &AwsDynamodbGlobalTable{},
        &AwsDynamodbGlobalTableList{},
    
        &AwsIamUserPolicyAttachment{},
        &AwsIamUserPolicyAttachmentList{},
    
        &AwsDxGatewayAssociationProposal{},
        &AwsDxGatewayAssociationProposalList{},
    
        &AwsDxLag{},
        &AwsDxLagList{},
    
        &AwsIotCertificate{},
        &AwsIotCertificateList{},
    
        &AwsRdsGlobalCluster{},
        &AwsRdsGlobalClusterList{},
    
        &AwsS3AccountPublicAccessBlock{},
        &AwsS3AccountPublicAccessBlockList{},
    
        &AwsConfigConfigRule{},
        &AwsConfigConfigRuleList{},
    
        &AwsDatasyncTask{},
        &AwsDatasyncTaskList{},
    
        &AwsOpsworksMemcachedLayer{},
        &AwsOpsworksMemcachedLayerList{},
    
        &AwsSesDomainDkim{},
        &AwsSesDomainDkimList{},
    
        &AwsDaxParameterGroup{},
        &AwsDaxParameterGroupList{},
    
        &AwsOpsworksNodejsAppLayer{},
        &AwsOpsworksNodejsAppLayerList{},
    
        &AwsLicensemanagerAssociation{},
        &AwsLicensemanagerAssociationList{},
    
        &AwsMainRouteTableAssociation{},
        &AwsMainRouteTableAssociationList{},
    
        &AwsSsmDocument{},
        &AwsSsmDocumentList{},
    
        &AwsStoragegatewaySmbFileShare{},
        &AwsStoragegatewaySmbFileShareList{},
    
        &AwsGameliftGameSessionQueue{},
        &AwsGameliftGameSessionQueueList{},
    
        &AwsIamRole{},
        &AwsIamRoleList{},
    
        &AwsBackupPlan{},
        &AwsBackupPlanList{},
    
        &AwsSesConfigurationSet{},
        &AwsSesConfigurationSetList{},
    
        &AwsIamSamlProvider{},
        &AwsIamSamlProviderList{},
    
        &AwsKmsExternalKey{},
        &AwsKmsExternalKeyList{},
    
        &AwsMskConfiguration{},
        &AwsMskConfigurationList{},
    
        &AwsOpsworksRdsDbInstance{},
        &AwsOpsworksRdsDbInstanceList{},
    
        &AwsRedshiftSecurityGroup{},
        &AwsRedshiftSecurityGroupList{},
    
        &AwsS3BucketObject{},
        &AwsS3BucketObjectList{},
    
        &AwsDxGateway{},
        &AwsDxGatewayList{},
    
        &AwsEbsSnapshotCopy{},
        &AwsEbsSnapshotCopyList{},
    
        &AwsSpotInstanceRequest{},
        &AwsSpotInstanceRequestList{},
    
        &AwsPinpointApnsVoipSandboxChannel{},
        &AwsPinpointApnsVoipSandboxChannelList{},
    
        &AwsApiGatewayIntegration{},
        &AwsApiGatewayIntegrationList{},
    
        &AwsApiGatewayRequestValidator{},
        &AwsApiGatewayRequestValidatorList{},
    
        &AwsNetworkInterfaceSgAttachment{},
        &AwsNetworkInterfaceSgAttachmentList{},
    
        &AwsTransferUser{},
        &AwsTransferUserList{},
    
        &AwsAthenaDatabase{},
        &AwsAthenaDatabaseList{},
    
        &AwsDxHostedPublicVirtualInterfaceAccepter{},
        &AwsDxHostedPublicVirtualInterfaceAccepterList{},
    
        &AwsIamAccountPasswordPolicy{},
        &AwsIamAccountPasswordPolicyList{},
    
        &AwsGameliftFleet{},
        &AwsGameliftFleetList{},
    
        &AwsLicensemanagerLicenseConfiguration{},
        &AwsLicensemanagerLicenseConfigurationList{},
    
        &AwsRedshiftSnapshotCopyGrant{},
        &AwsRedshiftSnapshotCopyGrantList{},
    
        &AwsVpc{},
        &AwsVpcList{},
    
        &AwsAmiCopy{},
        &AwsAmiCopyList{},
    
        &AwsDxBgpPeer{},
        &AwsDxBgpPeerList{},
    
        &AwsRoute53ResolverRule{},
        &AwsRoute53ResolverRuleList{},
    
        &AwsFlowLog{},
        &AwsFlowLogList{},
    
        &AwsRedshiftEventSubscription{},
        &AwsRedshiftEventSubscriptionList{},
    
        &AwsMqBroker{},
        &AwsMqBrokerList{},
    
        &AwsMediaStoreContainer{},
        &AwsMediaStoreContainerList{},
    
        &AwsNatGateway{},
        &AwsNatGatewayList{},
    
        &AwsDbSecurityGroup{},
        &AwsDbSecurityGroupList{},
    
        &AwsGlueJob{},
        &AwsGlueJobList{},
    
        &AwsCloudwatchLogGroup{},
        &AwsCloudwatchLogGroupList{},
    
        &AwsSesReceiptRuleSet{},
        &AwsSesReceiptRuleSetList{},
    
        &AwsKmsGrant{},
        &AwsKmsGrantList{},
    
        &AwsOpsworksHaproxyLayer{},
        &AwsOpsworksHaproxyLayerList{},
    
        &AwsPinpointSmsChannel{},
        &AwsPinpointSmsChannelList{},
    
        &AwsCloudwatchEventPermission{},
        &AwsCloudwatchEventPermissionList{},
    
        &AwsGlacierVaultLock{},
        &AwsGlacierVaultLockList{},
    
        &AwsOpsworksMysqlLayer{},
        &AwsOpsworksMysqlLayerList{},
    
        &AwsOpsworksPermission{},
        &AwsOpsworksPermissionList{},
    
        &AwsRouteTableAssociation{},
        &AwsRouteTableAssociationList{},
    
        &AwsServiceDiscoveryPublicDnsNamespace{},
        &AwsServiceDiscoveryPublicDnsNamespaceList{},
    
        &AwsStoragegatewayWorkingStorage{},
        &AwsStoragegatewayWorkingStorageList{},
    
        &AwsApiGatewayBasePathMapping{},
        &AwsApiGatewayBasePathMappingList{},
    
        &AwsAppmeshVirtualNode{},
        &AwsAppmeshVirtualNodeList{},
    
        &AwsLbSslNegotiationPolicy{},
        &AwsLbSslNegotiationPolicyList{},
    
        &AwsVpcPeeringConnectionAccepter{},
        &AwsVpcPeeringConnectionAccepterList{},
    
        &AwsConfigDeliveryChannel{},
        &AwsConfigDeliveryChannelList{},
    
        &AwsDxHostedPrivateVirtualInterfaceAccepter{},
        &AwsDxHostedPrivateVirtualInterfaceAccepterList{},
    
        &AwsIotThingType{},
        &AwsIotThingTypeList{},
    
        &AwsRamPrincipalAssociation{},
        &AwsRamPrincipalAssociationList{},
    
        &AwsSagemakerEndpoint{},
        &AwsSagemakerEndpointList{},
    
        &AwsSnapshotCreateVolumePermission{},
        &AwsSnapshotCreateVolumePermissionList{},
    
        &AwsBatchComputeEnvironment{},
        &AwsBatchComputeEnvironmentList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
