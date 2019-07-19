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

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&EcsService{},
		&EcsServiceList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&DxGateway{},
		&DxGatewayList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&S3Bucket{},
		&S3BucketList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&WafWebACL{},
		&WafWebACLList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IotThingType{},
		&IotThingTypeList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&TransferServer{},
		&TransferServerList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&KeyPair{},
		&KeyPairList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&BackupPlan{},
		&BackupPlanList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&KmsGrant{},
		&KmsGrantList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&PinpointApp{},
		&PinpointAppList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&IamUser{},
		&IamUserList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&Alb{},
		&AlbList{},

		&DaxCluster{},
		&DaxClusterList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SqsQueue{},
		&SqsQueueList{},

		&AmiCopy{},
		&AmiCopyList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&DbInstance{},
		&DbInstanceList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&Instance{},
		&InstanceList{},

		&MqBroker{},
		&MqBrokerList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&IamPolicy{},
		&IamPolicyList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&GlueJob{},
		&GlueJobList{},

		&Route{},
		&RouteList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&EipAssociation{},
		&EipAssociationList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&WafRule{},
		&WafRuleList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&AlbListener{},
		&AlbListenerList{},

		&SwfDomain{},
		&SwfDomainList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&Elb{},
		&ElbList{},

		&SsmParameter{},
		&SsmParameterList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&NatGateway{},
		&NatGatewayList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&Lb{},
		&LbList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&FlowLog{},
		&FlowLogList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&IotThing{},
		&IotThingList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&KmsKey{},
		&KmsKeyList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&Subnet{},
		&SubnetList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&Codepipeline{},
		&CodepipelineList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&Eip{},
		&EipList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&SnsTopic{},
		&SnsTopicList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&IotCertificate{},
		&IotCertificateList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WafIpset{},
		&WafIpsetList{},

		&Ami{},
		&AmiList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SesTemplate{},
		&SesTemplateList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&EksCluster{},
		&EksClusterList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&EmrCluster{},
		&EmrClusterList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&MskCluster{},
		&MskClusterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&SsmActivation{},
		&SsmActivationList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&IamRole{},
		&IamRoleList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&IamGroup{},
		&IamGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&TransferUser{},
		&TransferUserList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&RouteTable{},
		&RouteTableList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DxLag{},
		&DxLagList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&EcsCluster{},
		&EcsClusterList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&BackupVault{},
		&BackupVaultList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&SfnActivity{},
		&SfnActivityList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&NetworkACL{},
		&NetworkACLList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&KmsAlias{},
		&KmsAliasList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&Route53Record{},
		&Route53RecordList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&Vpc{},
		&VpcList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&LbListener{},
		&LbListenerList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
