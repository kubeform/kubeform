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

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DbInstance{},
		&DbInstanceList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&Ami{},
		&AmiList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&TransferServer{},
		&TransferServerList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&Alb{},
		&AlbList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&Elb{},
		&ElbList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&RdsCluster{},
		&RdsClusterList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&Vpc{},
		&VpcList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&KeyPair{},
		&KeyPairList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&IamUser{},
		&IamUserList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&WafRule{},
		&WafRuleList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IamRole{},
		&IamRoleList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&EcsService{},
		&EcsServiceList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SfnActivity{},
		&SfnActivityList{},

		&WafWebACL{},
		&WafWebACLList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DxConnection{},
		&DxConnectionList{},

		&IotThing{},
		&IotThingList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&DxLag{},
		&DxLagList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GlueJob{},
		&GlueJobList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&IamGroup{},
		&IamGroupList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&SqsQueue{},
		&SqsQueueList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&NetworkACL{},
		&NetworkACLList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&IotThingType{},
		&IotThingTypeList{},

		&KmsGrant{},
		&KmsGrantList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&SsmParameter{},
		&SsmParameterList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&NatGateway{},
		&NatGatewayList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&WafIpset{},
		&WafIpsetList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&MskCluster{},
		&MskClusterList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&EmrCluster{},
		&EmrClusterList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&SwfDomain{},
		&SwfDomainList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&LbListener{},
		&LbListenerList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&BackupVault{},
		&BackupVaultList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&MqBroker{},
		&MqBrokerList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&Lb{},
		&LbList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&IotCertificate{},
		&IotCertificateList{},

		&KmsAlias{},
		&KmsAliasList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&FlowLog{},
		&FlowLogList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&Route53Record{},
		&Route53RecordList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&S3Bucket{},
		&S3BucketList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&DxGateway{},
		&DxGatewayList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&TransferUser{},
		&TransferUserList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&IamPolicy{},
		&IamPolicyList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&PinpointApp{},
		&PinpointAppList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&Subnet{},
		&SubnetList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&Codepipeline{},
		&CodepipelineList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&Eip{},
		&EipList{},

		&EksCluster{},
		&EksClusterList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&RouteTable{},
		&RouteTableList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&Instance{},
		&InstanceList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&KmsKey{},
		&KmsKeyList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AlbListener{},
		&AlbListenerList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&EcsCluster{},
		&EcsClusterList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&Route{},
		&RouteList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
