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

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DaxCluster{},
		&DaxClusterList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&Elb{},
		&ElbList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&NetworkACL{},
		&NetworkACLList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&AmiCopy{},
		&AmiCopyList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&WafRule{},
		&WafRuleList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&SsmParameter{},
		&SsmParameterList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DxLag{},
		&DxLagList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&Route{},
		&RouteList{},

		&SwfDomain{},
		&SwfDomainList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SsmActivation{},
		&SsmActivationList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&GlueJob{},
		&GlueJobList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&IamGroup{},
		&IamGroupList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&RouteTable{},
		&RouteTableList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&IotThingType{},
		&IotThingTypeList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&IotCertificate{},
		&IotCertificateList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&Ami{},
		&AmiList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SnsTopic{},
		&SnsTopicList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&WafWebACL{},
		&WafWebACLList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&KeyPair{},
		&KeyPairList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&KmsAlias{},
		&KmsAliasList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DbInstance{},
		&DbInstanceList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Route53Record{},
		&Route53RecordList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IotThing{},
		&IotThingList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DxConnection{},
		&DxConnectionList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&Lb{},
		&LbList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&FlowLog{},
		&FlowLogList{},

		&SfnActivity{},
		&SfnActivityList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&MqBroker{},
		&MqBrokerList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&Subnet{},
		&SubnetList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&EmrCluster{},
		&EmrClusterList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&Codepipeline{},
		&CodepipelineList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&NatGateway{},
		&NatGatewayList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&WafIpset{},
		&WafIpsetList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&TransferUser{},
		&TransferUserList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&S3Bucket{},
		&S3BucketList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&Vpc{},
		&VpcList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EksCluster{},
		&EksClusterList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&PinpointApp{},
		&PinpointAppList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&KmsKey{},
		&KmsKeyList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&EcsService{},
		&EcsServiceList{},

		&EipAssociation{},
		&EipAssociationList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&LbListener{},
		&LbListenerList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&TransferServer{},
		&TransferServerList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&Alb{},
		&AlbList{},

		&Eip{},
		&EipList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SqsQueue{},
		&SqsQueueList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&BackupPlan{},
		&BackupPlanList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&Instance{},
		&InstanceList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&IamRole{},
		&IamRoleList{},

		&IotPolicy{},
		&IotPolicyList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&IamUser{},
		&IamUserList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&DxGateway{},
		&DxGatewayList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
