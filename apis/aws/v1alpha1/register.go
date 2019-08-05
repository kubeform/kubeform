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

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IamUser{},
		&IamUserList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&Ami{},
		&AmiList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&S3Bucket{},
		&S3BucketList{},

		&TransferServer{},
		&TransferServerList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&DxConnection{},
		&DxConnectionList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&IotCertificate{},
		&IotCertificateList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&IamGroup{},
		&IamGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&KmsAlias{},
		&KmsAliasList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&KeyPair{},
		&KeyPairList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&LbListener{},
		&LbListenerList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&Alb{},
		&AlbList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&EmrCluster{},
		&EmrClusterList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&KmsGrant{},
		&KmsGrantList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EksCluster{},
		&EksClusterList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&Codepipeline{},
		&CodepipelineList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&Eip{},
		&EipList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&Route{},
		&RouteList{},

		&RouteTable{},
		&RouteTableList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EcsService{},
		&EcsServiceList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&Elb{},
		&ElbList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&WafRule{},
		&WafRuleList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&DxGateway{},
		&DxGatewayList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&FlowLog{},
		&FlowLogList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&DxLag{},
		&DxLagList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&Subnet{},
		&SubnetList{},

		&TransferUser{},
		&TransferUserList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&GlueJob{},
		&GlueJobList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&IamRole{},
		&IamRoleList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&Vpc{},
		&VpcList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&IotThing{},
		&IotThingList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&IamPolicy{},
		&IamPolicyList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&Instance{},
		&InstanceList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&NetworkACL{},
		&NetworkACLList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&Route53Record{},
		&Route53RecordList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&IotThingType{},
		&IotThingTypeList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DbInstance{},
		&DbInstanceList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SsmParameter{},
		&SsmParameterList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&Lb{},
		&LbList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&MqBroker{},
		&MqBrokerList{},

		&MskCluster{},
		&MskClusterList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&IotPolicy{},
		&IotPolicyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&KmsKey{},
		&KmsKeyList{},

		&SqsQueue{},
		&SqsQueueList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafWebACL{},
		&WafWebACLList{},

		&BackupVault{},
		&BackupVaultList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&WafIpset{},
		&WafIpsetList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&BackupPlan{},
		&BackupPlanList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&SwfDomain{},
		&SwfDomainList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&PinpointApp{},
		&PinpointAppList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
