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

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&Codepipeline{},
		&CodepipelineList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&NetworkACL{},
		&NetworkACLList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&IotThing{},
		&IotThingList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DaxCluster{},
		&DaxClusterList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&KmsGrant{},
		&KmsGrantList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&KeyPair{},
		&KeyPairList{},

		&KmsAlias{},
		&KmsAliasList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&Alb{},
		&AlbList{},

		&DbInstance{},
		&DbInstanceList{},

		&EcsService{},
		&EcsServiceList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&AmiCopy{},
		&AmiCopyList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&WafIpset{},
		&WafIpsetList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&IamGroup{},
		&IamGroupList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&TransferServer{},
		&TransferServerList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&SqsQueue{},
		&SqsQueueList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&RdsCluster{},
		&RdsClusterList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DxLag{},
		&DxLagList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&TransferUser{},
		&TransferUserList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&Subnet{},
		&SubnetList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&Ami{},
		&AmiList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&Route{},
		&RouteList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&Lb{},
		&LbList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SesTemplate{},
		&SesTemplateList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&IotThingType{},
		&IotThingTypeList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&RouteTable{},
		&RouteTableList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&EmrCluster{},
		&EmrClusterList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&Vpc{},
		&VpcList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxGateway{},
		&DxGatewayList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&Elb{},
		&ElbList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&SwfDomain{},
		&SwfDomainList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&IamUser{},
		&IamUserList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&Instance{},
		&InstanceList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&KmsKey{},
		&KmsKeyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&IamRole{},
		&IamRoleList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SfnActivity{},
		&SfnActivityList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&FlowLog{},
		&FlowLogList{},

		&Route53Record{},
		&Route53RecordList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafRule{},
		&WafRuleList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&LbListener{},
		&LbListenerList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&IotCertificate{},
		&IotCertificateList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&AlbListener{},
		&AlbListenerList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&GlueJob{},
		&GlueJobList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&EipAssociation{},
		&EipAssociationList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&S3Bucket{},
		&S3BucketList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SsmParameter{},
		&SsmParameterList{},

		&EksCluster{},
		&EksClusterList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&PinpointApp{},
		&PinpointAppList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&WafWebACL{},
		&WafWebACLList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&SsmActivation{},
		&SsmActivationList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&Eip{},
		&EipList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
