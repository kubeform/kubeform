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

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&SqsQueue{},
		&SqsQueueList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&EcsCluster{},
		&EcsClusterList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&DxLag{},
		&DxLagList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&Subnet{},
		&SubnetList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&TransferUser{},
		&TransferUserList{},

		&Alb{},
		&AlbList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&Route53Record{},
		&Route53RecordList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&IamGroup{},
		&IamGroupList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&IotThing{},
		&IotThingList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&AlbListener{},
		&AlbListenerList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&KmsGrant{},
		&KmsGrantList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&IotThingType{},
		&IotThingTypeList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&NetworkACL{},
		&NetworkACLList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&GlueJob{},
		&GlueJobList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&SesTemplate{},
		&SesTemplateList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MskCluster{},
		&MskClusterList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&SfnActivity{},
		&SfnActivityList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&MqBroker{},
		&MqBrokerList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&Lb{},
		&LbList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&EmrCluster{},
		&EmrClusterList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&EipAssociation{},
		&EipAssociationList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&Vpc{},
		&VpcList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&Eip{},
		&EipList{},

		&DxGateway{},
		&DxGatewayList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&DaxCluster{},
		&DaxClusterList{},

		&KeyPair{},
		&KeyPairList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SwfDomain{},
		&SwfDomainList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&Elb{},
		&ElbList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&S3Bucket{},
		&S3BucketList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&KmsKey{},
		&KmsKeyList{},

		&LbListener{},
		&LbListenerList{},

		&IamRole{},
		&IamRoleList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&IamPolicy{},
		&IamPolicyList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&Instance{},
		&InstanceList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&WafRule{},
		&WafRuleList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&IamUser{},
		&IamUserList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&EksCluster{},
		&EksClusterList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&Codepipeline{},
		&CodepipelineList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafWebACL{},
		&WafWebACLList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SsmParameter{},
		&SsmParameterList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&EcsService{},
		&EcsServiceList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&SsmActivation{},
		&SsmActivationList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DbInstance{},
		&DbInstanceList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&KmsAlias{},
		&KmsAliasList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&PinpointApp{},
		&PinpointAppList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&TransferServer{},
		&TransferServerList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&FlowLog{},
		&FlowLogList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&IotPolicy{},
		&IotPolicyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&Ami{},
		&AmiList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&Route{},
		&RouteList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&AmiCopy{},
		&AmiCopyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&IotCertificate{},
		&IotCertificateList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
