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

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&BackupVault{},
		&BackupVaultList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&S3Bucket{},
		&S3BucketList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&AmiCopy{},
		&AmiCopyList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&Instance{},
		&InstanceList{},

		&IotCertificate{},
		&IotCertificateList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&Route53Record{},
		&Route53RecordList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&IamPolicy{},
		&IamPolicyList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&Vpc{},
		&VpcList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GlueJob{},
		&GlueJobList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&EipAssociation{},
		&EipAssociationList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&SsmActivation{},
		&SsmActivationList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&LbListener{},
		&LbListenerList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IotThingType{},
		&IotThingTypeList{},

		&WafIpset{},
		&WafIpsetList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&KmsKey{},
		&KmsKeyList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&WafWebACL{},
		&WafWebACLList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&NetworkACL{},
		&NetworkACLList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&Route{},
		&RouteList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&PinpointApp{},
		&PinpointAppList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&Subnet{},
		&SubnetList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&EksCluster{},
		&EksClusterList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&Eip{},
		&EipList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&EcsService{},
		&EcsServiceList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DxConnection{},
		&DxConnectionList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&FlowLog{},
		&FlowLogList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&DbInstance{},
		&DbInstanceList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DxGateway{},
		&DxGatewayList{},

		&AlbListener{},
		&AlbListenerList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&MskCluster{},
		&MskClusterList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&IamUser{},
		&IamUserList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&WafRule{},
		&WafRuleList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&Alb{},
		&AlbList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&TransferServer{},
		&TransferServerList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&SfnActivity{},
		&SfnActivityList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&Elb{},
		&ElbList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&KmsGrant{},
		&KmsGrantList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&Ami{},
		&AmiList{},

		&EmrCluster{},
		&EmrClusterList{},

		&IotThing{},
		&IotThingList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&IamGroup{},
		&IamGroupList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SsmParameter{},
		&SsmParameterList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&DxLag{},
		&DxLagList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&RdsCluster{},
		&RdsClusterList{},

		&BackupPlan{},
		&BackupPlanList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&RouteTable{},
		&RouteTableList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&Lb{},
		&LbList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&MqBroker{},
		&MqBrokerList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&TransferUser{},
		&TransferUserList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SesTemplate{},
		&SesTemplateList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Codepipeline{},
		&CodepipelineList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SqsQueue{},
		&SqsQueueList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&KmsAlias{},
		&KmsAliasList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&IamRole{},
		&IamRoleList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&KeyPair{},
		&KeyPairList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
