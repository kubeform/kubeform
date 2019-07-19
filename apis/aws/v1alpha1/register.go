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

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&WafWebACL{},
		&WafWebACLList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&Subnet{},
		&SubnetList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IotThingType{},
		&IotThingTypeList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&AmiCopy{},
		&AmiCopyList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&EipAssociation{},
		&EipAssociationList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&EcsService{},
		&EcsServiceList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&KmsAlias{},
		&KmsAliasList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&RouteTable{},
		&RouteTableList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&IotCertificate{},
		&IotCertificateList{},

		&MskCluster{},
		&MskClusterList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&KmsKey{},
		&KmsKeyList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&Elb{},
		&ElbList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&IamUser{},
		&IamUserList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Vpc{},
		&VpcList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&TransferServer{},
		&TransferServerList{},

		&WafIpset{},
		&WafIpsetList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&EmrCluster{},
		&EmrClusterList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&Lb{},
		&LbList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&EksCluster{},
		&EksClusterList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&KeyPair{},
		&KeyPairList{},

		&Route{},
		&RouteList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&FlowLog{},
		&FlowLogList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SqsQueue{},
		&SqsQueueList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&AlbListener{},
		&AlbListenerList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&GlueJob{},
		&GlueJobList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&KmsGrant{},
		&KmsGrantList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&WafRule{},
		&WafRuleList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&Instance{},
		&InstanceList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&RdsCluster{},
		&RdsClusterList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&Eip{},
		&EipList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&Alb{},
		&AlbList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&IotPolicy{},
		&IotPolicyList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&EcsCluster{},
		&EcsClusterList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&S3Bucket{},
		&S3BucketList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&NetworkACL{},
		&NetworkACLList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&SsmParameter{},
		&SsmParameterList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&IamGroup{},
		&IamGroupList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&Ami{},
		&AmiList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&TransferUser{},
		&TransferUserList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&LbListener{},
		&LbListenerList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&NatGateway{},
		&NatGatewayList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&IamRole{},
		&IamRoleList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IotThing{},
		&IotThingList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&Codepipeline{},
		&CodepipelineList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&DxConnection{},
		&DxConnectionList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&BackupPlan{},
		&BackupPlanList{},

		&SfnActivity{},
		&SfnActivityList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DxGateway{},
		&DxGatewayList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DxLag{},
		&DxLagList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&MqBroker{},
		&MqBrokerList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DbInstance{},
		&DbInstanceList{},

		&Route53Record{},
		&Route53RecordList{},

		&BackupVault{},
		&BackupVaultList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DaxCluster{},
		&DaxClusterList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SsmActivation{},
		&SsmActivationList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&PinpointApp{},
		&PinpointAppList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
