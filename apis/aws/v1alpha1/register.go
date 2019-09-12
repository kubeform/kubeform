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

		&OpsworksStack{},
		&OpsworksStackList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DxConnection{},
		&DxConnectionList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&RouteTable{},
		&RouteTableList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&S3Bucket{},
		&S3BucketList{},

		&LbListener{},
		&LbListenerList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&KeyPair{},
		&KeyPairList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&Route53Record{},
		&Route53RecordList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&DaxCluster{},
		&DaxClusterList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&IotThingType{},
		&IotThingTypeList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SsmParameter{},
		&SsmParameterList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&DxLag{},
		&DxLagList{},

		&IotCertificate{},
		&IotCertificateList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&Elb{},
		&ElbList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&DbInstance{},
		&DbInstanceList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IotPolicy{},
		&IotPolicyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&EmrCluster{},
		&EmrClusterList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&FlowLog{},
		&FlowLogList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&BackupVault{},
		&BackupVaultList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EcsService{},
		&EcsServiceList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&AlbListener{},
		&AlbListenerList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&TransferServer{},
		&TransferServerList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&KmsKey{},
		&KmsKeyList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&EcsCluster{},
		&EcsClusterList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&IamGroup{},
		&IamGroupList{},

		&IotThing{},
		&IotThingList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&Eip{},
		&EipList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&EipAssociation{},
		&EipAssociationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SwfDomain{},
		&SwfDomainList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&Ami{},
		&AmiList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&Subnet{},
		&SubnetList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&Vpc{},
		&VpcList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IamUser{},
		&IamUserList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&NatGateway{},
		&NatGatewayList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&TransferUser{},
		&TransferUserList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&Route{},
		&RouteList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SsmActivation{},
		&SsmActivationList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&NetworkACL{},
		&NetworkACLList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&EksCluster{},
		&EksClusterList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&SnsTopic{},
		&SnsTopicList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&KmsGrant{},
		&KmsGrantList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&Codepipeline{},
		&CodepipelineList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&Alb{},
		&AlbList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AmiCopy{},
		&AmiCopyList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&WafRule{},
		&WafRuleList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&IamRole{},
		&IamRoleList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&WafWebACL{},
		&WafWebACLList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&DxGateway{},
		&DxGatewayList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&SqsQueue{},
		&SqsQueueList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&PinpointApp{},
		&PinpointAppList{},

		&IamPolicy{},
		&IamPolicyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&Lb{},
		&LbList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&Instance{},
		&InstanceList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&GlueJob{},
		&GlueJobList{},

		&KmsAlias{},
		&KmsAliasList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&WafIpset{},
		&WafIpsetList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&SecurityGroup{},
		&SecurityGroupList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
