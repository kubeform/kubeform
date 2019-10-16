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

		&IotThingType{},
		&IotThingTypeList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&SfnActivity{},
		&SfnActivityList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&Instance{},
		&InstanceList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&DxConnection{},
		&DxConnectionList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&KeyPair{},
		&KeyPairList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&WafIpset{},
		&WafIpsetList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&IamRole{},
		&IamRoleList{},

		&S3Bucket{},
		&S3BucketList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&Subnet{},
		&SubnetList{},

		&TransferServer{},
		&TransferServerList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&SesTemplate{},
		&SesTemplateList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&IamUser{},
		&IamUserList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&AlbListener{},
		&AlbListenerList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&NetworkACL{},
		&NetworkACLList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&Vpc{},
		&VpcList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&RouteTable{},
		&RouteTableList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&PinpointApp{},
		&PinpointAppList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&EcsService{},
		&EcsServiceList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&AmiCopy{},
		&AmiCopyList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&Route53Record{},
		&Route53RecordList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&LbListener{},
		&LbListenerList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&KmsKey{},
		&KmsKeyList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&Eip{},
		&EipList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&GlueJob{},
		&GlueJobList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&DaxCluster{},
		&DaxClusterList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IotThing{},
		&IotThingList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&IotCertificate{},
		&IotCertificateList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&WafRule{},
		&WafRuleList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&EmrCluster{},
		&EmrClusterList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&RdsCluster{},
		&RdsClusterList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&EipAssociation{},
		&EipAssociationList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&DxGateway{},
		&DxGatewayList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SsmActivation{},
		&SsmActivationList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DxLag{},
		&DxLagList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DbInstance{},
		&DbInstanceList{},

		&EksCluster{},
		&EksClusterList{},

		&KmsAlias{},
		&KmsAliasList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IotPolicy{},
		&IotPolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&NatGateway{},
		&NatGatewayList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&IamGroup{},
		&IamGroupList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&WafWebACL{},
		&WafWebACLList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&Alb{},
		&AlbList{},

		&Lb{},
		&LbList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Route{},
		&RouteList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&MqBroker{},
		&MqBrokerList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&SsmParameter{},
		&SsmParameterList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&SqsQueue{},
		&SqsQueueList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&TransferUser{},
		&TransferUserList{},

		&BackupVault{},
		&BackupVaultList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&FlowLog{},
		&FlowLogList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&Elb{},
		&ElbList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&Ami{},
		&AmiList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
