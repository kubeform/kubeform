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

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&WafIpset{},
		&WafIpsetList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&Eip{},
		&EipList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&KeyPair{},
		&KeyPairList{},

		&Alb{},
		&AlbList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&LbListener{},
		&LbListenerList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&GlueJob{},
		&GlueJobList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&TransferUser{},
		&TransferUserList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&NatGateway{},
		&NatGatewayList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&SnsTopic{},
		&SnsTopicList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&Lb{},
		&LbList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&MqBroker{},
		&MqBrokerList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&IamGroup{},
		&IamGroupList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&NetworkACL{},
		&NetworkACLList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&IamRole{},
		&IamRoleList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&AmiCopy{},
		&AmiCopyList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&Elb{},
		&ElbList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&Route53Record{},
		&Route53RecordList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SqsQueue{},
		&SqsQueueList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AlbListener{},
		&AlbListenerList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&BackupPlan{},
		&BackupPlanList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&MskCluster{},
		&MskClusterList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&EmrCluster{},
		&EmrClusterList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&Ami{},
		&AmiList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&Instance{},
		&InstanceList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&EksCluster{},
		&EksClusterList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&IotCertificate{},
		&IotCertificateList{},

		&KmsAlias{},
		&KmsAliasList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DxGateway{},
		&DxGatewayList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&TransferServer{},
		&TransferServerList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&BackupVault{},
		&BackupVaultList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&SsmParameter{},
		&SsmParameterList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&Subnet{},
		&SubnetList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&IotThingType{},
		&IotThingTypeList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&PinpointApp{},
		&PinpointAppList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&SfnActivity{},
		&SfnActivityList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&WafRule{},
		&WafRuleList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&DxLag{},
		&DxLagList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&IamPolicy{},
		&IamPolicyList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DxConnection{},
		&DxConnectionList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&KmsKey{},
		&KmsKeyList{},

		&S3Bucket{},
		&S3BucketList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&Vpc{},
		&VpcList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&Codepipeline{},
		&CodepipelineList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&EcsService{},
		&EcsServiceList{},

		&FlowLog{},
		&FlowLogList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&WafWebACL{},
		&WafWebACLList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&SesTemplate{},
		&SesTemplateList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&DbInstance{},
		&DbInstanceList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&RdsCluster{},
		&RdsClusterList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&EcsCluster{},
		&EcsClusterList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IamUser{},
		&IamUserList{},

		&IotThing{},
		&IotThingList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&Route{},
		&RouteList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
