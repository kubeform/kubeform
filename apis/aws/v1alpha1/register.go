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

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&KmsGrant{},
		&KmsGrantList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&EksCluster{},
		&EksClusterList{},

		&IotThing{},
		&IotThingList{},

		&NetworkACL{},
		&NetworkACLList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&NatGateway{},
		&NatGatewayList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&AmiCopy{},
		&AmiCopyList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&WafRule{},
		&WafRuleList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&S3Bucket{},
		&S3BucketList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&EipAssociation{},
		&EipAssociationList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&Elb{},
		&ElbList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&Vpc{},
		&VpcList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&RouteTable{},
		&RouteTableList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&EcsService{},
		&EcsServiceList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&Route53Record{},
		&Route53RecordList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&KeyPair{},
		&KeyPairList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DxConnection{},
		&DxConnectionList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&MqBroker{},
		&MqBrokerList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&Instance{},
		&InstanceList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&GlueJob{},
		&GlueJobList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&EmrCluster{},
		&EmrClusterList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&IamGroup{},
		&IamGroupList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&WafIpset{},
		&WafIpsetList{},

		&WafWebACL{},
		&WafWebACLList{},

		&PinpointApp{},
		&PinpointAppList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&KmsAlias{},
		&KmsAliasList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&FlowLog{},
		&FlowLogList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&IamPolicy{},
		&IamPolicyList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IotCertificate{},
		&IotCertificateList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SfnActivity{},
		&SfnActivityList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&KmsKey{},
		&KmsKeyList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&DxGateway{},
		&DxGatewayList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&Route{},
		&RouteList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&IotThingType{},
		&IotThingTypeList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&LbListener{},
		&LbListenerList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SesTemplate{},
		&SesTemplateList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&Alb{},
		&AlbList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&IotPolicy{},
		&IotPolicyList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&TransferUser{},
		&TransferUserList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&Codepipeline{},
		&CodepipelineList{},

		&AlbListener{},
		&AlbListenerList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&Lb{},
		&LbList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&DaxCluster{},
		&DaxClusterList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SqsQueue{},
		&SqsQueueList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&EcsCluster{},
		&EcsClusterList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&MskCluster{},
		&MskClusterList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&BackupVault{},
		&BackupVaultList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SwfDomain{},
		&SwfDomainList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&DxLag{},
		&DxLagList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&DbInstance{},
		&DbInstanceList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IamUser{},
		&IamUserList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&Ami{},
		&AmiList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&Eip{},
		&EipList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&Subnet{},
		&SubnetList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&BackupPlan{},
		&BackupPlanList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&SsmActivation{},
		&SsmActivationList{},

		&TransferServer{},
		&TransferServerList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IamRole{},
		&IamRoleList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&SsmParameter{},
		&SsmParameterList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
