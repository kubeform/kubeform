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

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&ApiGatewayRestApi{},
		&ApiGatewayRestApiList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&SfnActivity{},
		&SfnActivityList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&MskCluster{},
		&MskClusterList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&EksCluster{},
		&EksClusterList{},

		&FlowLog{},
		&FlowLogList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafRule{},
		&WafRuleList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ServiceDiscoveryPrivateDnsNamespace{},
		&ServiceDiscoveryPrivateDnsNamespaceList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&ServiceDiscoveryHttpNamespace{},
		&ServiceDiscoveryHttpNamespaceList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&NetworkAcl{},
		&NetworkAclList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&Alb{},
		&AlbList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppsyncApiKey{},
		&AppsyncApiKeyList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&Instance{},
		&InstanceList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&SnsTopic{},
		&SnsTopicList{},

		&TransferUser{},
		&TransferUserList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&BackupPlan{},
		&BackupPlanList{},

		&KmsAlias{},
		&KmsAliasList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SqsQueue{},
		&SqsQueueList{},

		&TransferSshKey{},
		&TransferSshKeyList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&IamRole{},
		&IamRoleList{},

		&DxBgpPeer{},
		&DxBgpPeerList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SwfDomain{},
		&SwfDomainList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&VpcIpv4CidrBlockAssociation{},
		&VpcIpv4CidrBlockAssociationList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&Ec2ClientVpnNetworkAssociation{},
		&Ec2ClientVpnNetworkAssociationList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&ServiceDiscoveryPublicDnsNamespace{},
		&ServiceDiscoveryPublicDnsNamespaceList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&Subnet{},
		&SubnetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MqBroker{},
		&MqBrokerList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&EmrCluster{},
		&EmrClusterList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&KeyPair{},
		&KeyPairList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&IamUser{},
		&IamUserList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&WafWebAcl{},
		&WafWebAclList{},

		&WafregionalWebAcl{},
		&WafregionalWebAclList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&LightsailStaticIpAttachment{},
		&LightsailStaticIpAttachmentList{},

		&DefaultVpcDhcpOptions{},
		&DefaultVpcDhcpOptionsList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&IotCertificate{},
		&IotCertificateList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&NetworkAclRule{},
		&NetworkAclRuleList{},

		&VpcDhcpOptions{},
		&VpcDhcpOptionsList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&Ami{},
		&AmiList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&DefaultNetworkAcl{},
		&DefaultNetworkAclList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&VpcDhcpOptionsAssociation{},
		&VpcDhcpOptionsAssociationList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&Ec2ClientVpnEndpoint{},
		&Ec2ClientVpnEndpointList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&PinpointApp{},
		&PinpointAppList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&AmiCopy{},
		&AmiCopyList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&NatGateway{},
		&NatGatewayList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&IamPolicy{},
		&IamPolicyList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&S3Bucket{},
		&S3BucketList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&RouteTable{},
		&RouteTableList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&LightsailStaticIp{},
		&LightsailStaticIpList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DaxCluster{},
		&DaxClusterList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&Elb{},
		&ElbList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&EcsCluster{},
		&EcsClusterList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&Eip{},
		&EipList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&KmsKey{},
		&KmsKeyList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&IamUserSshKey{},
		&IamUserSshKeyList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&Route53Record{},
		&Route53RecordList{},

		&EcsService{},
		&EcsServiceList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&DbInstance{},
		&DbInstanceList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DxConnection{},
		&DxConnectionList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&Lb{},
		&LbList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&RdsCluster{},
		&RdsClusterList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&IotThingType{},
		&IotThingTypeList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&IotThing{},
		&IotThingList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&ApiGatewayApiKey{},
		&ApiGatewayApiKeyList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&Vpc{},
		&VpcList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&AlbListener{},
		&AlbListenerList{},

		&IamGroup{},
		&IamGroupList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&LbListener{},
		&LbListenerList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&WafIpset{},
		&WafIpsetList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&TransferServer{},
		&TransferServerList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SsmParameter{},
		&SsmParameterList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&Codepipeline{},
		&CodepipelineList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&AppsyncGraphqlApi{},
		&AppsyncGraphqlApiList{},

		&GlueJob{},
		&GlueJobList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DxLag{},
		&DxLagList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&KmsGrant{},
		&KmsGrantList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IotPolicy{},
		&IotPolicyList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&WafSqlInjectionMatchSet{},
		&WafSqlInjectionMatchSetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&Route{},
		&RouteList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&DxGateway{},
		&DxGatewayList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&WafregionalWebAclAssociation{},
		&WafregionalWebAclAssociationList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&WafregionalSqlInjectionMatchSet{},
		&WafregionalSqlInjectionMatchSetList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
