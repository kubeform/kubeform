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

		&DbSnapshot{},
		&DbSnapshotList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&Ami{},
		&AmiList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&MqBroker{},
		&MqBrokerList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&Eip{},
		&EipList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&AlbListener{},
		&AlbListenerList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&BackupPlan{},
		&BackupPlanList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&Route{},
		&RouteList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&WafWebACL{},
		&WafWebACLList{},

		&FlowLog{},
		&FlowLogList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&EksCluster{},
		&EksClusterList{},

		&NatGateway{},
		&NatGatewayList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DxLag{},
		&DxLagList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DaxCluster{},
		&DaxClusterList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&LbListener{},
		&LbListenerList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&TransferServer{},
		&TransferServerList{},

		&RouteTable{},
		&RouteTableList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&IamRole{},
		&IamRoleList{},

		&DxGateway{},
		&DxGatewayList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&KmsGrant{},
		&KmsGrantList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IotThing{},
		&IotThingList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&Vpc{},
		&VpcList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&SsmParameter{},
		&SsmParameterList{},

		&Subnet{},
		&SubnetList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&Alb{},
		&AlbList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&IamGroup{},
		&IamGroupList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&Instance{},
		&InstanceList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafRule{},
		&WafRuleList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&IotCertificate{},
		&IotCertificateList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&Route53Record{},
		&Route53RecordList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&SfnActivity{},
		&SfnActivityList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&SwfDomain{},
		&SwfDomainList{},

		&DxConnection{},
		&DxConnectionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&IotThingType{},
		&IotThingTypeList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&Lb{},
		&LbList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&KmsAlias{},
		&KmsAliasList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&AmiCopy{},
		&AmiCopyList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&EipAssociation{},
		&EipAssociationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&EmrCluster{},
		&EmrClusterList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SsmActivation{},
		&SsmActivationList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&IamUser{},
		&IamUserList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&BackupVault{},
		&BackupVaultList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&WafIpset{},
		&WafIpsetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SqsQueue{},
		&SqsQueueList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&Elb{},
		&ElbList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&Codepipeline{},
		&CodepipelineList{},

		&KmsKey{},
		&KmsKeyList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&EcsService{},
		&EcsServiceList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IamPolicy{},
		&IamPolicyList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&NetworkACL{},
		&NetworkACLList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&TransferUser{},
		&TransferUserList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&KeyPair{},
		&KeyPairList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&GlueJob{},
		&GlueJobList{},

		&S3Bucket{},
		&S3BucketList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DbInstance{},
		&DbInstanceList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&RdsCluster{},
		&RdsClusterList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&PinpointApp{},
		&PinpointAppList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
