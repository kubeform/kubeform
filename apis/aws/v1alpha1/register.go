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

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DxConnection{},
		&DxConnectionList{},

		&Alb{},
		&AlbList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&FlowLog{},
		&FlowLogList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&GlueJob{},
		&GlueJobList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&KmsGrant{},
		&KmsGrantList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&NatGateway{},
		&NatGatewayList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&Subnet{},
		&SubnetList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&TransferUser{},
		&TransferUserList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&RouteTable{},
		&RouteTableList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&IotCertificate{},
		&IotCertificateList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&RdsCluster{},
		&RdsClusterList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&EcsService{},
		&EcsServiceList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&SqsQueue{},
		&SqsQueueList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&PinpointApp{},
		&PinpointAppList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&EksCluster{},
		&EksClusterList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Instance{},
		&InstanceList{},

		&IotThing{},
		&IotThingList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&KeyPair{},
		&KeyPairList{},

		&WafRule{},
		&WafRuleList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&Route53Record{},
		&Route53RecordList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&AmiCopy{},
		&AmiCopyList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&NetworkACL{},
		&NetworkACLList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&EmrCluster{},
		&EmrClusterList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&Eip{},
		&EipList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&KmsAlias{},
		&KmsAliasList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AlbListener{},
		&AlbListenerList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&IamRole{},
		&IamRoleList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&SsmActivation{},
		&SsmActivationList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IotThingType{},
		&IotThingTypeList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&Route{},
		&RouteList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&BackupVault{},
		&BackupVaultList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafWebACL{},
		&WafWebACLList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&Elb{},
		&ElbList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&S3Bucket{},
		&S3BucketList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&DbInstance{},
		&DbInstanceList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&SsmParameter{},
		&SsmParameterList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&KmsKey{},
		&KmsKeyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SnsTopic{},
		&SnsTopicList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&Vpc{},
		&VpcList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&MqBroker{},
		&MqBrokerList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&EipAssociation{},
		&EipAssociationList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&EcsCluster{},
		&EcsClusterList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&TransferServer{},
		&TransferServerList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&IotPolicy{},
		&IotPolicyList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&DxLag{},
		&DxLagList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&LbListener{},
		&LbListenerList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&IamUser{},
		&IamUserList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&Lb{},
		&LbList{},

		&Ami{},
		&AmiList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamGroup{},
		&IamGroupList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&MskCluster{},
		&MskClusterList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DxGateway{},
		&DxGatewayList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
