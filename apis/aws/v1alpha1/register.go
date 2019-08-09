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

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&BackupPlan{},
		&BackupPlanList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&Eip{},
		&EipList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&SsmActivation{},
		&SsmActivationList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&IamGroup{},
		&IamGroupList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IotPolicy{},
		&IotPolicyList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&NatGateway{},
		&NatGatewayList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&Alb{},
		&AlbList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&TransferUser{},
		&TransferUserList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&IamRole{},
		&IamRoleList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IotThingType{},
		&IotThingTypeList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&DxConnection{},
		&DxConnectionList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&KeyPair{},
		&KeyPairList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&EksCluster{},
		&EksClusterList{},

		&GlueJob{},
		&GlueJobList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&Route{},
		&RouteList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&Elb{},
		&ElbList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&AlbListener{},
		&AlbListenerList{},

		&NetworkACL{},
		&NetworkACLList{},

		&WafIpset{},
		&WafIpsetList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DxGateway{},
		&DxGatewayList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&DxLag{},
		&DxLagList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&S3Bucket{},
		&S3BucketList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&KmsKey{},
		&KmsKeyList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&EcsCluster{},
		&EcsClusterList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&EmrCluster{},
		&EmrClusterList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&Vpc{},
		&VpcList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&PinpointApp{},
		&PinpointAppList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&SnsTopic{},
		&SnsTopicList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&MqBroker{},
		&MqBrokerList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&SsmParameter{},
		&SsmParameterList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&TransferServer{},
		&TransferServerList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&Ami{},
		&AmiList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&KmsGrant{},
		&KmsGrantList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&Instance{},
		&InstanceList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&RouteTable{},
		&RouteTableList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&FlowLog{},
		&FlowLogList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&Route53Record{},
		&Route53RecordList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SwfDomain{},
		&SwfDomainList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IotCertificate{},
		&IotCertificateList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafWebACL{},
		&WafWebACLList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&EcsService{},
		&EcsServiceList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&BackupVault{},
		&BackupVaultList{},

		&IotThing{},
		&IotThingList{},

		&KmsAlias{},
		&KmsAliasList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&WafRule{},
		&WafRuleList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&LbListener{},
		&LbListenerList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&IamUser{},
		&IamUserList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DbInstance{},
		&DbInstanceList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&Subnet{},
		&SubnetList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&Lb{},
		&LbList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&SqsQueue{},
		&SqsQueueList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&DaxCluster{},
		&DaxClusterList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
