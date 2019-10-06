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

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Vpc{},
		&VpcList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SsmParameter{},
		&SsmParameterList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&KmsAlias{},
		&KmsAliasList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&Route{},
		&RouteList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&Instance{},
		&InstanceList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&Route53Record{},
		&Route53RecordList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&NetworkACL{},
		&NetworkACLList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IamUser{},
		&IamUserList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&LbListener{},
		&LbListenerList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&NatGateway{},
		&NatGatewayList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&EmrCluster{},
		&EmrClusterList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&Alb{},
		&AlbList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&Subnet{},
		&SubnetList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&DxLag{},
		&DxLagList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamPolicy{},
		&IamPolicyList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&WafRule{},
		&WafRuleList{},

		&WafWebACL{},
		&WafWebACLList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&EcsCluster{},
		&EcsClusterList{},

		&FlowLog{},
		&FlowLogList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&Eip{},
		&EipList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&TransferServer{},
		&TransferServerList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&DaxCluster{},
		&DaxClusterList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&IotPolicy{},
		&IotPolicyList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&IotThingType{},
		&IotThingTypeList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&Elb{},
		&ElbList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&PinpointApp{},
		&PinpointAppList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&AlbListener{},
		&AlbListenerList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&IamGroup{},
		&IamGroupList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&KmsKey{},
		&KmsKeyList{},

		&MqBroker{},
		&MqBrokerList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&Lb{},
		&LbList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SsmActivation{},
		&SsmActivationList{},

		&WafIpset{},
		&WafIpsetList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&RouteTable{},
		&RouteTableList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&KeyPair{},
		&KeyPairList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&IotThing{},
		&IotThingList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SnsTopic{},
		&SnsTopicList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&KmsGrant{},
		&KmsGrantList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&SwfDomain{},
		&SwfDomainList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&RdsCluster{},
		&RdsClusterList{},

		&EcsService{},
		&EcsServiceList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&TransferUser{},
		&TransferUserList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&GlueJob{},
		&GlueJobList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SqsQueue{},
		&SqsQueueList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&S3Bucket{},
		&S3BucketList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&EipAssociation{},
		&EipAssociationList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&BackupVault{},
		&BackupVaultList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&SfnActivity{},
		&SfnActivityList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&IotCertificate{},
		&IotCertificateList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DxGateway{},
		&DxGatewayList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&Ami{},
		&AmiList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&EksCluster{},
		&EksClusterList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&IamRole{},
		&IamRoleList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&DbInstance{},
		&DbInstanceList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
