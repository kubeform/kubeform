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

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&BackupPlan{},
		&BackupPlanList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&S3Bucket{},
		&S3BucketList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&IamGroup{},
		&IamGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&WafIpset{},
		&WafIpsetList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&IamUser{},
		&IamUserList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&PinpointApp{},
		&PinpointAppList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&DaxCluster{},
		&DaxClusterList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&RdsCluster{},
		&RdsClusterList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IotCertificate{},
		&IotCertificateList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&Alb{},
		&AlbList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&IotThingType{},
		&IotThingTypeList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&MskCluster{},
		&MskClusterList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&Route{},
		&RouteList{},

		&SqsQueue{},
		&SqsQueueList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&WafWebACL{},
		&WafWebACLList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&NetworkACL{},
		&NetworkACLList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DxGateway{},
		&DxGatewayList{},

		&NatGateway{},
		&NatGatewayList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&LbListener{},
		&LbListenerList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&IotThing{},
		&IotThingList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&Ami{},
		&AmiList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&BackupVault{},
		&BackupVaultList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SesTemplate{},
		&SesTemplateList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&Instance{},
		&InstanceList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&AmiCopy{},
		&AmiCopyList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&KmsGrant{},
		&KmsGrantList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&KmsAlias{},
		&KmsAliasList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&KmsKey{},
		&KmsKeyList{},

		&MqBroker{},
		&MqBrokerList{},

		&RouteTable{},
		&RouteTableList{},

		&Elb{},
		&ElbList{},

		&IamPolicy{},
		&IamPolicyList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&Route53Record{},
		&Route53RecordList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&TransferServer{},
		&TransferServerList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&EcsCluster{},
		&EcsClusterList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&DbInstance{},
		&DbInstanceList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&SwfDomain{},
		&SwfDomainList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&Subnet{},
		&SubnetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SfnActivity{},
		&SfnActivityList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&EmrCluster{},
		&EmrClusterList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&EcsService{},
		&EcsServiceList{},

		&KeyPair{},
		&KeyPairList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&IotPolicy{},
		&IotPolicyList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxLag{},
		&DxLagList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&GlueJob{},
		&GlueJobList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&WafRule{},
		&WafRuleList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&IamRole{},
		&IamRoleList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&Lb{},
		&LbList{},

		&DxConnection{},
		&DxConnectionList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&FlowLog{},
		&FlowLogList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SnsTopic{},
		&SnsTopicList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&SsmParameter{},
		&SsmParameterList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&TransferUser{},
		&TransferUserList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&Eip{},
		&EipList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&EksCluster{},
		&EksClusterList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&Vpc{},
		&VpcList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
