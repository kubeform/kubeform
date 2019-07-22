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

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&IamRole{},
		&IamRoleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&Eip{},
		&EipList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&Subnet{},
		&SubnetList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&DxGateway{},
		&DxGatewayList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&Lb{},
		&LbList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&MqBroker{},
		&MqBrokerList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&IamGroup{},
		&IamGroupList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&LbListener{},
		&LbListenerList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&SwfDomain{},
		&SwfDomainList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&IamUser{},
		&IamUserList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&NetworkACL{},
		&NetworkACLList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&EipAssociation{},
		&EipAssociationList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&Codepipeline{},
		&CodepipelineList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SsmActivation{},
		&SsmActivationList{},

		&TransferServer{},
		&TransferServerList{},

		&WafIpset{},
		&WafIpsetList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&Route{},
		&RouteList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DaxCluster{},
		&DaxClusterList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&S3Bucket{},
		&S3BucketList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&Alb{},
		&AlbList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&WafWebACL{},
		&WafWebACLList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&DxLag{},
		&DxLagList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&EcsService{},
		&EcsServiceList{},

		&RouteTable{},
		&RouteTableList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&MskCluster{},
		&MskClusterList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&GlueJob{},
		&GlueJobList{},

		&KeyPair{},
		&KeyPairList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&TransferUser{},
		&TransferUserList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&SfnActivity{},
		&SfnActivityList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&Vpc{},
		&VpcList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&SesTemplate{},
		&SesTemplateList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&DxConnection{},
		&DxConnectionList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&Ami{},
		&AmiList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SnsTopic{},
		&SnsTopicList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&Elb{},
		&ElbList{},

		&PinpointApp{},
		&PinpointAppList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&WafRule{},
		&WafRuleList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&Instance{},
		&InstanceList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&SsmParameter{},
		&SsmParameterList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&KmsGrant{},
		&KmsGrantList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EmrCluster{},
		&EmrClusterList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&IotCertificate{},
		&IotCertificateList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&IotThingType{},
		&IotThingTypeList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&FlowLog{},
		&FlowLogList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&SqsQueue{},
		&SqsQueueList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&BackupVault{},
		&BackupVaultList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IotThing{},
		&IotThingList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&EksCluster{},
		&EksClusterList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&KmsAlias{},
		&KmsAliasList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&KmsKey{},
		&KmsKeyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&AmiCopy{},
		&AmiCopyList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&DbInstance{},
		&DbInstanceList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&NatGateway{},
		&NatGatewayList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Route53Record{},
		&Route53RecordList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
