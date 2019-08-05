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

		&ShieldProtection{},
		&ShieldProtectionList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&MskCluster{},
		&MskClusterList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&IamRole{},
		&IamRoleList{},

		&PinpointApp{},
		&PinpointAppList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&DxConnection{},
		&DxConnectionList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DbInstance{},
		&DbInstanceList{},

		&DxGateway{},
		&DxGatewayList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&SqsQueue{},
		&SqsQueueList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&Eip{},
		&EipList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&NetworkACL{},
		&NetworkACLList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&Lb{},
		&LbList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&BackupVault{},
		&BackupVaultList{},

		&EmrCluster{},
		&EmrClusterList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&SesTemplate{},
		&SesTemplateList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&TransferServer{},
		&TransferServerList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&TransferUser{},
		&TransferUserList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&EksCluster{},
		&EksClusterList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&IotThingType{},
		&IotThingTypeList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Subnet{},
		&SubnetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IotThing{},
		&IotThingList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&IotCertificate{},
		&IotCertificateList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&WafWebACL{},
		&WafWebACLList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&IamGroup{},
		&IamGroupList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SwfDomain{},
		&SwfDomainList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&NatGateway{},
		&NatGatewayList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&Ami{},
		&AmiList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&GlueJob{},
		&GlueJobList{},

		&Vpc{},
		&VpcList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&FlowLog{},
		&FlowLogList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&Route{},
		&RouteList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&WafRule{},
		&WafRuleList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SsmParameter{},
		&SsmParameterList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&EcsCluster{},
		&EcsClusterList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SfnActivity{},
		&SfnActivityList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DaxCluster{},
		&DaxClusterList{},

		&Elb{},
		&ElbList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&KmsKey{},
		&KmsKeyList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&KmsAlias{},
		&KmsAliasList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&Route53Record{},
		&Route53RecordList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&MqBroker{},
		&MqBrokerList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&Codepipeline{},
		&CodepipelineList{},

		&Instance{},
		&InstanceList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&EcsService{},
		&EcsServiceList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&RouteTable{},
		&RouteTableList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&LbListener{},
		&LbListenerList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&IamUser{},
		&IamUserList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&DxLag{},
		&DxLagList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&KeyPair{},
		&KeyPairList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&AlbListener{},
		&AlbListenerList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&WafIpset{},
		&WafIpsetList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&S3Bucket{},
		&S3BucketList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&EipAssociation{},
		&EipAssociationList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&KmsGrant{},
		&KmsGrantList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&Alb{},
		&AlbList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
