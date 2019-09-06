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

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&NetworkACL{},
		&NetworkACLList{},

		&Route{},
		&RouteList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&DxLag{},
		&DxLagList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafWebACL{},
		&WafWebACLList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SwfDomain{},
		&SwfDomainList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&GlueJob{},
		&GlueJobList{},

		&NatGateway{},
		&NatGatewayList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&WafRule{},
		&WafRuleList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&IamGroup{},
		&IamGroupList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&EksCluster{},
		&EksClusterList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&DbInstance{},
		&DbInstanceList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&Elb{},
		&ElbList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&DxGateway{},
		&DxGatewayList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&Alb{},
		&AlbList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&Codepipeline{},
		&CodepipelineList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&IamRole{},
		&IamRoleList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&KmsAlias{},
		&KmsAliasList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&LbListener{},
		&LbListenerList{},

		&Ami{},
		&AmiList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&IotThing{},
		&IotThingList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&MqBroker{},
		&MqBrokerList{},

		&RouteTable{},
		&RouteTableList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SesTemplate{},
		&SesTemplateList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&Instance{},
		&InstanceList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&KeyPair{},
		&KeyPairList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&EipAssociation{},
		&EipAssociationList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&Subnet{},
		&SubnetList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DxConnection{},
		&DxConnectionList{},

		&KmsGrant{},
		&KmsGrantList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&IotThingType{},
		&IotThingTypeList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&KmsKey{},
		&KmsKeyList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&Eip{},
		&EipList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&EcsService{},
		&EcsServiceList{},

		&Route53Record{},
		&Route53RecordList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&PinpointApp{},
		&PinpointAppList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&SnsTopic{},
		&SnsTopicList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&TransferUser{},
		&TransferUserList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&IamUser{},
		&IamUserList{},

		&TransferServer{},
		&TransferServerList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SsmParameter{},
		&SsmParameterList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&SfnActivity{},
		&SfnActivityList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IotCertificate{},
		&IotCertificateList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&FlowLog{},
		&FlowLogList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&SsmDocument{},
		&SsmDocumentList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AmiCopy{},
		&AmiCopyList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&Vpc{},
		&VpcList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&EmrCluster{},
		&EmrClusterList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&BackupPlan{},
		&BackupPlanList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&SsmActivation{},
		&SsmActivationList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&BackupVault{},
		&BackupVaultList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&Lb{},
		&LbList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&S3Bucket{},
		&S3BucketList{},

		&SqsQueue{},
		&SqsQueueList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
