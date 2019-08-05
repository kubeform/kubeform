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

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&IotThing{},
		&IotThingList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&Eip{},
		&EipList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&LbListener{},
		&LbListenerList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SnsTopic{},
		&SnsTopicList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&TransferServer{},
		&TransferServerList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&KeyPair{},
		&KeyPairList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&NetworkACL{},
		&NetworkACLList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DbInstance{},
		&DbInstanceList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&Ami{},
		&AmiList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&FlowLog{},
		&FlowLogList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SfnActivity{},
		&SfnActivityList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&WafWebACL{},
		&WafWebACLList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&KmsGrant{},
		&KmsGrantList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&Route53Record{},
		&Route53RecordList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&Elb{},
		&ElbList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&Alb{},
		&AlbList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&Vpc{},
		&VpcList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&Codepipeline{},
		&CodepipelineList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&IotThingType{},
		&IotThingTypeList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&Subnet{},
		&SubnetList{},

		&EipAssociation{},
		&EipAssociationList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&AmiCopy{},
		&AmiCopyList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&Instance{},
		&InstanceList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&RouteTable{},
		&RouteTableList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&EmrCluster{},
		&EmrClusterList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&S3Bucket{},
		&S3BucketList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&WafRule{},
		&WafRuleList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IotCertificate{},
		&IotCertificateList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DxLag{},
		&DxLagList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&NatGateway{},
		&NatGatewayList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&EksCluster{},
		&EksClusterList{},

		&TransferUser{},
		&TransferUserList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&AlbListener{},
		&AlbListenerList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&IamUser{},
		&IamUserList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&MqBroker{},
		&MqBrokerList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&Lb{},
		&LbList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&IotPolicy{},
		&IotPolicyList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DxGateway{},
		&DxGatewayList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&KmsAlias{},
		&KmsAliasList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&Route{},
		&RouteList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&IamGroup{},
		&IamGroupList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DaxCluster{},
		&DaxClusterList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&KmsKey{},
		&KmsKeyList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SqsQueue{},
		&SqsQueueList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&EcsService{},
		&EcsServiceList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GlueJob{},
		&GlueJobList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&SsmParameter{},
		&SsmParameterList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&BackupVault{},
		&BackupVaultList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&MskCluster{},
		&MskClusterList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&PinpointApp{},
		&PinpointAppList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&IamRole{},
		&IamRoleList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&EcsCluster{},
		&EcsClusterList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&BackupPlan{},
		&BackupPlanList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DxConnection{},
		&DxConnectionList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&IamPolicy{},
		&IamPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
