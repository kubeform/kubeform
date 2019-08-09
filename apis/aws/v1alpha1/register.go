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

		&FlowLog{},
		&FlowLogList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&EcsCluster{},
		&EcsClusterList{},

		&S3Bucket{},
		&S3BucketList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&KmsAlias{},
		&KmsAliasList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&BackupVault{},
		&BackupVaultList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DaxCluster{},
		&DaxClusterList{},

		&IotPolicy{},
		&IotPolicyList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&EmrCluster{},
		&EmrClusterList{},

		&IamRole{},
		&IamRoleList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&Lb{},
		&LbList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&WafWebACL{},
		&WafWebACLList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&Route53Record{},
		&Route53RecordList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&Instance{},
		&InstanceList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SsmParameter{},
		&SsmParameterList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&LbListener{},
		&LbListenerList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&IamPolicy{},
		&IamPolicyList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DxLag{},
		&DxLagList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&AlbListener{},
		&AlbListenerList{},

		&Eip{},
		&EipList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&DxGateway{},
		&DxGatewayList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&Alb{},
		&AlbList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&WafRule{},
		&WafRuleList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SwfDomain{},
		&SwfDomainList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&MqBroker{},
		&MqBrokerList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&EksCluster{},
		&EksClusterList{},

		&Route{},
		&RouteList{},

		&IotThing{},
		&IotThingList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DbInstance{},
		&DbInstanceList{},

		&Subnet{},
		&SubnetList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&KmsKey{},
		&KmsKeyList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&Ami{},
		&AmiList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&RdsCluster{},
		&RdsClusterList{},

		&Vpc{},
		&VpcList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&IamGroup{},
		&IamGroupList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&GlueJob{},
		&GlueJobList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&WafIpset{},
		&WafIpsetList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&TransferServer{},
		&TransferServerList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&IamUser{},
		&IamUserList{},

		&SesTemplate{},
		&SesTemplateList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&EipAssociation{},
		&EipAssociationList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&DxConnection{},
		&DxConnectionList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&KmsGrant{},
		&KmsGrantList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&IotCertificate{},
		&IotCertificateList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SqsQueue{},
		&SqsQueueList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&PinpointApp{},
		&PinpointAppList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&EcsService{},
		&EcsServiceList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&SsmActivation{},
		&SsmActivationList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&TransferUser{},
		&TransferUserList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&IotThingType{},
		&IotThingTypeList{},

		&NetworkACL{},
		&NetworkACLList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&Elb{},
		&ElbList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&KeyPair{},
		&KeyPairList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&Codepipeline{},
		&CodepipelineList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
