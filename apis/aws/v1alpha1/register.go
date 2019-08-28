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
    
        &Alb{},
        &AlbList{},
    
        &AppmeshVirtualService{},
        &AppmeshVirtualServiceList{},
    
        &CodecommitRepository{},
        &CodecommitRepositoryList{},
    
        &DocdbClusterInstance{},
        &DocdbClusterInstanceList{},
    
        &EcrLifecyclePolicy{},
        &EcrLifecyclePolicyList{},
    
        &LbSSLNegotiationPolicy{},
        &LbSSLNegotiationPolicyList{},
    
        &SsmResourceDataSync{},
        &SsmResourceDataSyncList{},
    
        &AcmpcaCertificateAuthority{},
        &AcmpcaCertificateAuthorityList{},
    
        &ApiGatewayClientCertificate{},
        &ApiGatewayClientCertificateList{},
    
        &WafSQLInjectionMatchSet{},
        &WafSQLInjectionMatchSetList{},
    
        &IamPolicyAttachment{},
        &IamPolicyAttachmentList{},
    
        &Route53ZoneAssociation{},
        &Route53ZoneAssociationList{},
    
        &SecretsmanagerSecretVersion{},
        &SecretsmanagerSecretVersionList{},
    
        &ApiGatewayIntegration{},
        &ApiGatewayIntegrationList{},
    
        &VpcDHCPOptions{},
        &VpcDHCPOptionsList{},
    
        &PinpointGcmChannel{},
        &PinpointGcmChannelList{},
    
        &AlbListenerCertificate{},
        &AlbListenerCertificateList{},
    
        &AppautoscalingScheduledAction{},
        &AppautoscalingScheduledActionList{},
    
        &AutoscalingPolicy{},
        &AutoscalingPolicyList{},
    
        &Vpc{},
        &VpcList{},
    
        &WafregionalRegexPatternSet{},
        &WafregionalRegexPatternSetList{},
    
        &DbClusterSnapshot{},
        &DbClusterSnapshotList{},
    
        &DbSubnetGroup{},
        &DbSubnetGroupList{},
    
        &ElasticBeanstalkApplicationVersion{},
        &ElasticBeanstalkApplicationVersionList{},
    
        &Elb{},
        &ElbList{},
    
        &RdsGlobalCluster{},
        &RdsGlobalClusterList{},
    
        &CloudwatchEventTarget{},
        &CloudwatchEventTargetList{},
    
        &Instance{},
        &InstanceList{},
    
        &NetworkACLRule{},
        &NetworkACLRuleList{},
    
        &RedshiftCluster{},
        &RedshiftClusterList{},
    
        &ApiGatewayModel{},
        &ApiGatewayModelList{},
    
        &DbEventSubscription{},
        &DbEventSubscriptionList{},
    
        &WafregionalByteMatchSet{},
        &WafregionalByteMatchSetList{},
    
        &GlobalacceleratorListener{},
        &GlobalacceleratorListenerList{},
    
        &DocdbSubnetGroup{},
        &DocdbSubnetGroupList{},
    
        &IamAccountAlias{},
        &IamAccountAliasList{},
    
        &MacieS3BucketAssociation{},
        &MacieS3BucketAssociationList{},
    
        &OpsworksRdsDbInstance{},
        &OpsworksRdsDbInstanceList{},
    
        &CloudwatchLogStream{},
        &CloudwatchLogStreamList{},
    
        &PinpointApnsVoipChannel{},
        &PinpointApnsVoipChannelList{},
    
        &AcmCertificate{},
        &AcmCertificateList{},
    
        &IotThingType{},
        &IotThingTypeList{},
    
        &LambdaPermission{},
        &LambdaPermissionList{},
    
        &ApiGatewayDomainName{},
        &ApiGatewayDomainNameList{},
    
        &CognitoUserPool{},
        &CognitoUserPoolList{},
    
        &EbsSnapshot{},
        &EbsSnapshotList{},
    
        &PinpointAdmChannel{},
        &PinpointAdmChannelList{},
    
        &LambdaEventSourceMapping{},
        &LambdaEventSourceMappingList{},
    
        &SecurityGroup{},
        &SecurityGroupList{},
    
        &ApiGatewayRestAPI{},
        &ApiGatewayRestAPIList{},
    
        &CodepipelineWebhook{},
        &CodepipelineWebhookList{},
    
        &GlueCatalogTable{},
        &GlueCatalogTableList{},
    
        &GlueJob{},
        &GlueJobList{},
    
        &IamAccessKey{},
        &IamAccessKeyList{},
    
        &IotTopicRule{},
        &IotTopicRuleList{},
    
        &WafByteMatchSet{},
        &WafByteMatchSetList{},
    
        &RamResourceAssociation{},
        &RamResourceAssociationList{},
    
        &S3BucketInventory{},
        &S3BucketInventoryList{},
    
        &StoragegatewaySmbFileShare{},
        &StoragegatewaySmbFileShareList{},
    
        &ConfigAggregateAuthorization{},
        &ConfigAggregateAuthorizationList{},
    
        &CognitoUserPoolDomain{},
        &CognitoUserPoolDomainList{},
    
        &DocdbClusterSnapshot{},
        &DocdbClusterSnapshotList{},
    
        &GlueConnection{},
        &GlueConnectionList{},
    
        &OpsworksJavaAppLayer{},
        &OpsworksJavaAppLayerList{},
    
        &S3BucketObject{},
        &S3BucketObjectList{},
    
        &SesReceiptRuleSet{},
        &SesReceiptRuleSetList{},
    
        &AppsyncResolver{},
        &AppsyncResolverList{},
    
        &CloudfrontPublicKey{},
        &CloudfrontPublicKeyList{},
    
        &CodedeployDeploymentConfig{},
        &CodedeployDeploymentConfigList{},
    
        &EfsFileSystem{},
        &EfsFileSystemList{},
    
        &ElasticacheParameterGroup{},
        &ElasticacheParameterGroupList{},
    
        &SagemakerNotebookInstanceLifecycleConfiguration{},
        &SagemakerNotebookInstanceLifecycleConfigurationList{},
    
        &ConfigDeliveryChannel{},
        &ConfigDeliveryChannelList{},
    
        &LicensemanagerAssociation{},
        &LicensemanagerAssociationList{},
    
        &XraySamplingRule{},
        &XraySamplingRuleList{},
    
        &EmrInstanceGroup{},
        &EmrInstanceGroupList{},
    
        &VolumeAttachment{},
        &VolumeAttachmentList{},
    
        &Cloudtrail{},
        &CloudtrailList{},
    
        &ElastictranscoderPreset{},
        &ElastictranscoderPresetList{},
    
        &KinesisFirehoseDeliveryStream{},
        &KinesisFirehoseDeliveryStreamList{},
    
        &RedshiftSubnetGroup{},
        &RedshiftSubnetGroupList{},
    
        &SnsTopic{},
        &SnsTopicList{},
    
        &ApiGatewayAuthorizer{},
        &ApiGatewayAuthorizerList{},
    
        &DocdbClusterParameterGroup{},
        &DocdbClusterParameterGroupList{},
    
        &IamUserLoginProfile{},
        &IamUserLoginProfileList{},
    
        &CloudwatchLogGroup{},
        &CloudwatchLogGroupList{},
    
        &KmsAlias{},
        &KmsAliasList{},
    
        &OpsworksUserProfile{},
        &OpsworksUserProfileList{},
    
        &SesTemplate{},
        &SesTemplateList{},
    
        &WafregionalRegexMatchSet{},
        &WafregionalRegexMatchSetList{},
    
        &MainRouteTableAssociation{},
        &MainRouteTableAssociationList{},
    
        &NeptuneCluster{},
        &NeptuneClusterList{},
    
        &SagemakerModel{},
        &SagemakerModelList{},
    
        &SagemakerEndpointConfiguration{},
        &SagemakerEndpointConfigurationList{},
    
        &TransferUser{},
        &TransferUserList{},
    
        &WafRegexPatternSet{},
        &WafRegexPatternSetList{},
    
        &ApiGatewayDeployment{},
        &ApiGatewayDeploymentList{},
    
        &AppmeshVirtualRouter{},
        &AppmeshVirtualRouterList{},
    
        &ElasticBeanstalkEnvironment{},
        &ElasticBeanstalkEnvironmentList{},
    
        &NetworkACL{},
        &NetworkACLList{},
    
        &WafregionalRule{},
        &WafregionalRuleList{},
    
        &DatasyncAgent{},
        &DatasyncAgentList{},
    
        &LightsailKeyPair{},
        &LightsailKeyPairList{},
    
        &OrganizationsPolicy{},
        &OrganizationsPolicyList{},
    
        &CodedeployDeploymentGroup{},
        &CodedeployDeploymentGroupList{},
    
        &DaxCluster{},
        &DaxClusterList{},
    
        &DxConnection{},
        &DxConnectionList{},
    
        &NatGateway{},
        &NatGatewayList{},
    
        &RamResourceShare{},
        &RamResourceShareList{},
    
        &SwfDomain{},
        &SwfDomainList{},
    
        &VpcPeeringConnectionOptions{},
        &VpcPeeringConnectionOptionsList{},
    
        &DbOptionGroup{},
        &DbOptionGroupList{},
    
        &ElasticBeanstalkConfigurationTemplate{},
        &ElasticBeanstalkConfigurationTemplateList{},
    
        &RdsClusterParameterGroup{},
        &RdsClusterParameterGroupList{},
    
        &PlacementGroup{},
        &PlacementGroupList{},
    
        &SnsTopicPolicy{},
        &SnsTopicPolicyList{},
    
        &InspectorAssessmentTarget{},
        &InspectorAssessmentTargetList{},
    
        &StoragegatewayCache{},
        &StoragegatewayCacheList{},
    
        &CognitoIdentityPool{},
        &CognitoIdentityPoolList{},
    
        &ApiGatewayDocumentationPart{},
        &ApiGatewayDocumentationPartList{},
    
        &DatasyncLocationS3{},
        &DatasyncLocationS3List{},
    
        &GameliftAlias{},
        &GameliftAliasList{},
    
        &OrganizationsPolicyAttachment{},
        &OrganizationsPolicyAttachmentList{},
    
        &ServiceDiscoveryPrivateDNSNamespace{},
        &ServiceDiscoveryPrivateDNSNamespaceList{},
    
        &DbInstanceRoleAssociation{},
        &DbInstanceRoleAssociationList{},
    
        &IamGroupPolicyAttachment{},
        &IamGroupPolicyAttachmentList{},
    
        &RedshiftParameterGroup{},
        &RedshiftParameterGroupList{},
    
        &DefaultRouteTable{},
        &DefaultRouteTableList{},
    
        &SecurityGroupRule{},
        &SecurityGroupRuleList{},
    
        &Codepipeline{},
        &CodepipelineList{},
    
        &LbListener{},
        &LbListenerList{},
    
        &IamAccountPasswordPolicy{},
        &IamAccountPasswordPolicyList{},
    
        &DefaultNetworkACL{},
        &DefaultNetworkACLList{},
    
        &WafregionalXssMatchSet{},
        &WafregionalXssMatchSetList{},
    
        &ApiGatewayBasePathMapping{},
        &ApiGatewayBasePathMappingList{},
    
        &DaxSubnetGroup{},
        &DaxSubnetGroupList{},
    
        &DxGateway{},
        &DxGatewayList{},
    
        &RdsClusterEndpoint{},
        &RdsClusterEndpointList{},
    
        &WafregionalRateBasedRule{},
        &WafregionalRateBasedRuleList{},
    
        &AlbTargetGroup{},
        &AlbTargetGroupList{},
    
        &CloudwatchLogMetricFilter{},
        &CloudwatchLogMetricFilterList{},
    
        &SesDomainDkim{},
        &SesDomainDkimList{},
    
        &VpcEndpoint{},
        &VpcEndpointList{},
    
        &VpcEndpointServiceAllowedPrincipal{},
        &VpcEndpointServiceAllowedPrincipalList{},
    
        &WafIpset{},
        &WafIpsetList{},
    
        &BackupPlan{},
        &BackupPlanList{},
    
        &NeptuneClusterInstance{},
        &NeptuneClusterInstanceList{},
    
        &EgressOnlyInternetGateway{},
        &EgressOnlyInternetGatewayList{},
    
        &StoragegatewayGateway{},
        &StoragegatewayGatewayList{},
    
        &BatchJobDefinition{},
        &BatchJobDefinitionList{},
    
        &GameliftGameSessionQueue{},
        &GameliftGameSessionQueueList{},
    
        &OpsworksPhpAppLayer{},
        &OpsworksPhpAppLayerList{},
    
        &AppsyncAPIKey{},
        &AppsyncAPIKeyList{},
    
        &CloudhsmV2Hsm{},
        &CloudhsmV2HsmList{},
    
        &KmsExternalKey{},
        &KmsExternalKeyList{},
    
        &SagemakerEndpoint{},
        &SagemakerEndpointList{},
    
        &ServiceDiscoveryPublicDNSNamespace{},
        &ServiceDiscoveryPublicDNSNamespaceList{},
    
        &ApiGatewayAPIKey{},
        &ApiGatewayAPIKeyList{},
    
        &CloudwatchEventRule{},
        &CloudwatchEventRuleList{},
    
        &CodedeployApp{},
        &CodedeployAppList{},
    
        &GlacierVault{},
        &GlacierVaultList{},
    
        &GuarddutyThreatintelset{},
        &GuarddutyThreatintelsetList{},
    
        &RamPrincipalAssociation{},
        &RamPrincipalAssociationList{},
    
        &StoragegatewayNfsFileShare{},
        &StoragegatewayNfsFileShareList{},
    
        &CloudwatchDashboard{},
        &CloudwatchDashboardList{},
    
        &MqBroker{},
        &MqBrokerList{},
    
        &Route{},
        &RouteList{},
    
        &IotThingPrincipalAttachment{},
        &IotThingPrincipalAttachmentList{},
    
        &LambdaLayerVersion{},
        &LambdaLayerVersionList{},
    
        &CloudwatchLogDestinationPolicy{},
        &CloudwatchLogDestinationPolicyList{},
    
        &EbsVolume{},
        &EbsVolumeList{},
    
        &EfsMountTarget{},
        &EfsMountTargetList{},
    
        &IamOpenidConnectProvider{},
        &IamOpenidConnectProviderList{},
    
        &IamUserGroupMembership{},
        &IamUserGroupMembershipList{},
    
        &InspectorAssessmentTemplate{},
        &InspectorAssessmentTemplateList{},
    
        &DmsReplicationTask{},
        &DmsReplicationTaskList{},
    
        &ElbAttachment{},
        &ElbAttachmentList{},
    
        &NetworkInterfaceAttachment{},
        &NetworkInterfaceAttachmentList{},
    
        &AppmeshMesh{},
        &AppmeshMeshList{},
    
        &CloudwatchLogSubscriptionFilter{},
        &CloudwatchLogSubscriptionFilterList{},
    
        &GlueTrigger{},
        &GlueTriggerList{},
    
        &KinesisStream{},
        &KinesisStreamList{},
    
        &S3Bucket{},
        &S3BucketList{},
    
        &WafregionalWebACL{},
        &WafregionalWebACLList{},
    
        &CloudfrontDistribution{},
        &CloudfrontDistributionList{},
    
        &DlmLifecyclePolicy{},
        &DlmLifecyclePolicyList{},
    
        &PinpointApnsChannel{},
        &PinpointApnsChannelList{},
    
        &AlbListenerRule{},
        &AlbListenerRuleList{},
    
        &DocdbCluster{},
        &DocdbClusterList{},
    
        &DxConnectionAssociation{},
        &DxConnectionAssociationList{},
    
        &Ec2TransitGatewayRoute{},
        &Ec2TransitGatewayRouteList{},
    
        &CloudformationStackSetInstance{},
        &CloudformationStackSetInstanceList{},
    
        &GameliftBuild{},
        &GameliftBuildList{},
    
        &LoadBalancerPolicy{},
        &LoadBalancerPolicyList{},
    
        &MediaStoreContainer{},
        &MediaStoreContainerList{},
    
        &CloudwatchLogResourcePolicy{},
        &CloudwatchLogResourcePolicyList{},
    
        &IamRolePolicy{},
        &IamRolePolicyList{},
    
        &OpsworksMemcachedLayer{},
        &OpsworksMemcachedLayerList{},
    
        &SqsQueuePolicy{},
        &SqsQueuePolicyList{},
    
        &GlueClassifier{},
        &GlueClassifierList{},
    
        &OrganizationsOrganizationalUnit{},
        &OrganizationsOrganizationalUnitList{},
    
        &SesEventDestination{},
        &SesEventDestinationList{},
    
        &S3BucketPolicy{},
        &S3BucketPolicyList{},
    
        &SfnActivity{},
        &SfnActivityList{},
    
        &CloudwatchEventPermission{},
        &CloudwatchEventPermissionList{},
    
        &WafregionalWebACLAssociation{},
        &WafregionalWebACLAssociationList{},
    
        &Cloud9EnvironmentEc2{},
        &Cloud9EnvironmentEc2List{},
    
        &CloudformationStack{},
        &CloudformationStackList{},
    
        &Ec2TransitGatewayRouteTableAssociation{},
        &Ec2TransitGatewayRouteTableAssociationList{},
    
        &LambdaFunction{},
        &LambdaFunctionList{},
    
        &LbListenerRule{},
        &LbListenerRuleList{},
    
        &CloudformationStackSet{},
        &CloudformationStackSetList{},
    
        &DxHostedPrivateVirtualInterface{},
        &DxHostedPrivateVirtualInterfaceList{},
    
        &Ec2TransitGatewayRouteTablePropagation{},
        &Ec2TransitGatewayRouteTablePropagationList{},
    
        &EcsCluster{},
        &EcsClusterList{},
    
        &Route53ResolverRuleAssociation{},
        &Route53ResolverRuleAssociationList{},
    
        &SpotDatafeedSubscription{},
        &SpotDatafeedSubscriptionList{},
    
        &CodecommitTrigger{},
        &CodecommitTriggerList{},
    
        &Ec2ClientVPNNetworkAssociation{},
        &Ec2ClientVPNNetworkAssociationList{},
    
        &MediaStoreContainerPolicy{},
        &MediaStoreContainerPolicyList{},
    
        &Eip{},
        &EipList{},
    
        &ElasticacheSecurityGroup{},
        &ElasticacheSecurityGroupList{},
    
        &RedshiftEventSubscription{},
        &RedshiftEventSubscriptionList{},
    
        &Route53ResolverEndpoint{},
        &Route53ResolverEndpointList{},
    
        &SsmAssociation{},
        &SsmAssociationList{},
    
        &DynamodbTable{},
        &DynamodbTableList{},
    
        &OpsworksGangliaLayer{},
        &OpsworksGangliaLayerList{},
    
        &SnsSmsPreferences{},
        &SnsSmsPreferencesList{},
    
        &DxPublicVirtualInterface{},
        &DxPublicVirtualInterfaceList{},
    
        &DynamodbGlobalTable{},
        &DynamodbGlobalTableList{},
    
        &IamUserPolicyAttachment{},
        &IamUserPolicyAttachmentList{},
    
        &LaunchConfiguration{},
        &LaunchConfigurationList{},
    
        &NeptuneEventSubscription{},
        &NeptuneEventSubscriptionList{},
    
        &SsmMaintenanceWindowTask{},
        &SsmMaintenanceWindowTaskList{},
    
        &CognitoUserPoolClient{},
        &CognitoUserPoolClientList{},
    
        &ElasticacheReplicationGroup{},
        &ElasticacheReplicationGroupList{},
    
        &SesDomainIdentityVerification{},
        &SesDomainIdentityVerificationList{},
    
        &Lb{},
        &LbList{},
    
        &GlacierVaultLock{},
        &GlacierVaultLockList{},
    
        &IamUserSSHKey{},
        &IamUserSSHKeyList{},
    
        &IamUser{},
        &IamUserList{},
    
        &ResourcegroupsGroup{},
        &ResourcegroupsGroupList{},
    
        &AcmCertificateValidation{},
        &AcmCertificateValidationList{},
    
        &DmsEndpoint{},
        &DmsEndpointList{},
    
        &LbCookieStickinessPolicy{},
        &LbCookieStickinessPolicyList{},
    
        &NetworkInterface{},
        &NetworkInterfaceList{},
    
        &SesDomainMailFrom{},
        &SesDomainMailFromList{},
    
        &IamServiceLinkedRole{},
        &IamServiceLinkedRoleList{},
    
        &VpcEndpointService{},
        &VpcEndpointServiceList{},
    
        &BatchJobQueue{},
        &BatchJobQueueList{},
    
        &CloudhsmV2Cluster{},
        &CloudhsmV2ClusterList{},
    
        &DbSnapshot{},
        &DbSnapshotList{},
    
        &LightsailInstance{},
        &LightsailInstanceList{},
    
        &NeptuneParameterGroup{},
        &NeptuneParameterGroupList{},
    
        &S3BucketPublicAccessBlock{},
        &S3BucketPublicAccessBlockList{},
    
        &WafRuleGroup{},
        &WafRuleGroupList{},
    
        &DatasyncTask{},
        &DatasyncTaskList{},
    
        &SsmMaintenanceWindowTarget{},
        &SsmMaintenanceWindowTargetList{},
    
        &SnapshotCreateVolumePermission{},
        &SnapshotCreateVolumePermissionList{},
    
        &VpcEndpointSubnetAssociation{},
        &VpcEndpointSubnetAssociationList{},
    
        &CognitoUserGroup{},
        &CognitoUserGroupList{},
    
        &DbParameterGroup{},
        &DbParameterGroupList{},
    
        &DxHostedPrivateVirtualInterfaceAccepter{},
        &DxHostedPrivateVirtualInterfaceAccepterList{},
    
        &EbsSnapshotCopy{},
        &EbsSnapshotCopyList{},
    
        &LbTargetGroup{},
        &LbTargetGroupList{},
    
        &AlbTargetGroupAttachment{},
        &AlbTargetGroupAttachmentList{},
    
        &AutoscalingNotification{},
        &AutoscalingNotificationList{},
    
        &IamGroupMembership{},
        &IamGroupMembershipList{},
    
        &MacieMemberAccountAssociation{},
        &MacieMemberAccountAssociationList{},
    
        &MqConfiguration{},
        &MqConfigurationList{},
    
        &Route53DelegationSet{},
        &Route53DelegationSetList{},
    
        &SesConfigurationSet{},
        &SesConfigurationSetList{},
    
        &WafregionalSizeConstraintSet{},
        &WafregionalSizeConstraintSetList{},
    
        &PinpointSmsChannel{},
        &PinpointSmsChannelList{},
    
        &ApiGatewayStage{},
        &ApiGatewayStageList{},
    
        &DxHostedPublicVirtualInterfaceAccepter{},
        &DxHostedPublicVirtualInterfaceAccepterList{},
    
        &LicensemanagerLicenseConfiguration{},
        &LicensemanagerLicenseConfigurationList{},
    
        &OpsworksHaproxyLayer{},
        &OpsworksHaproxyLayerList{},
    
        &TransferServer{},
        &TransferServerList{},
    
        &WafregionalIpset{},
        &WafregionalIpsetList{},
    
        &DmsReplicationSubnetGroup{},
        &DmsReplicationSubnetGroupList{},
    
        &Ec2ClientVPNEndpoint{},
        &Ec2ClientVPNEndpointList{},
    
        &IamUserPolicy{},
        &IamUserPolicyList{},
    
        &Ec2CapacityReservation{},
        &Ec2CapacityReservationList{},
    
        &GameliftFleet{},
        &GameliftFleetList{},
    
        &LbTargetGroupAttachment{},
        &LbTargetGroupAttachmentList{},
    
        &OpsworksMysqlLayer{},
        &OpsworksMysqlLayerList{},
    
        &SesActiveReceiptRuleSet{},
        &SesActiveReceiptRuleSetList{},
    
        &WafregionalSQLInjectionMatchSet{},
        &WafregionalSQLInjectionMatchSetList{},
    
        &VpcEndpointRouteTableAssociation{},
        &VpcEndpointRouteTableAssociationList{},
    
        &ApiGatewayAccount{},
        &ApiGatewayAccountList{},
    
        &ElasticacheCluster{},
        &ElasticacheClusterList{},
    
        &IamRole{},
        &IamRoleList{},
    
        &KmsCiphertext{},
        &KmsCiphertextList{},
    
        &LightsailStaticIP{},
        &LightsailStaticIPList{},
    
        &VpcPeeringConnectionAccepter{},
        &VpcPeeringConnectionAccepterList{},
    
        &CognitoResourceServer{},
        &CognitoResourceServerList{},
    
        &IotRoleAlias{},
        &IotRoleAliasList{},
    
        &LaunchTemplate{},
        &LaunchTemplateList{},
    
        &Subnet{},
        &SubnetList{},
    
        &ConfigConfigurationRecorderStatus_{},
        &ConfigConfigurationRecorderStatus_List{},
    
        &Ec2TransitGatewayVpcAttachment{},
        &Ec2TransitGatewayVpcAttachmentList{},
    
        &SesReceiptFilter{},
        &SesReceiptFilterList{},
    
        &DefaultSecurityGroup{},
        &DefaultSecurityGroupList{},
    
        &AppsyncDatasource{},
        &AppsyncDatasourceList{},
    
        &AutoscalingLifecycleHook{},
        &AutoscalingLifecycleHookList{},
    
        &DatasyncLocationNfs{},
        &DatasyncLocationNfsList{},
    
        &OpsworksStaticWebLayer{},
        &OpsworksStaticWebLayerList{},
    
        &ApiGatewayUsagePlanKey{},
        &ApiGatewayUsagePlanKeyList{},
    
        &EcrRepositoryPolicy{},
        &EcrRepositoryPolicyList{},
    
        &IamGroupPolicy{},
        &IamGroupPolicyList{},
    
        &InternetGateway{},
        &InternetGatewayList{},
    
        &LoadBalancerListenerPolicy{},
        &LoadBalancerListenerPolicyList{},
    
        &SnsTopicSubscription{},
        &SnsTopicSubscriptionList{},
    
        &SagemakerNotebookInstance{},
        &SagemakerNotebookInstanceList{},
    
        &DbInstance{},
        &DbInstanceList{},
    
        &DirectoryServiceConditionalForwarder{},
        &DirectoryServiceConditionalForwarderList{},
    
        &Ec2TransitGateway{},
        &Ec2TransitGatewayList{},
    
        &BudgetsBudget{},
        &BudgetsBudgetList{},
    
        &CloudfrontOriginAccessIdentity{},
        &CloudfrontOriginAccessIdentityList{},
    
        &ApiGatewayMethodSettings{},
        &ApiGatewayMethodSettingsList{},
    
        &AppautoscalingTarget{},
        &AppautoscalingTargetList{},
    
        &CognitoIdentityProvider{},
        &CognitoIdentityProviderList{},
    
        &GuarddutyInviteAccepter{},
        &GuarddutyInviteAccepterList{},
    
        &SesIdentityNotificationTopic{},
        &SesIdentityNotificationTopicList{},
    
        &SecurityhubProductSubscription{},
        &SecurityhubProductSubscriptionList{},
    
        &AppCookieStickinessPolicy{},
        &AppCookieStickinessPolicyList{},
    
        &GuarddutyMember{},
        &GuarddutyMemberList{},
    
        &Route53ResolverRule{},
        &Route53ResolverRuleList{},
    
        &MediaPackageChannel{},
        &MediaPackageChannelList{},
    
        &DevicefarmProject{},
        &DevicefarmProjectList{},
    
        &EksCluster{},
        &EksClusterList{},
    
        &ElasticsearchDomainPolicy{},
        &ElasticsearchDomainPolicyList{},
    
        &Route53Record{},
        &Route53RecordList{},
    
        &SsmPatchBaseline{},
        &SsmPatchBaselineList{},
    
        &WafWebACL{},
        &WafWebACLList{},
    
        &BackupVault{},
        &BackupVaultList{},
    
        &VpcEndpointConnectionNotification{},
        &VpcEndpointConnectionNotificationList{},
    
        &SesDomainIdentity{},
        &SesDomainIdentityList{},
    
        &StoragegatewayWorkingStorage{},
        &StoragegatewayWorkingStorageList{},
    
        &SpotFleetRequest{},
        &SpotFleetRequestList{},
    
        &PinpointBaiduChannel{},
        &PinpointBaiduChannelList{},
    
        &Ami{},
        &AmiList{},
    
        &GlueCatalogDatabase{},
        &GlueCatalogDatabaseList{},
    
        &IamSamlProvider{},
        &IamSamlProviderList{},
    
        &SecretsmanagerSecret{},
        &SecretsmanagerSecretList{},
    
        &SecurityhubAccount{},
        &SecurityhubAccountList{},
    
        &SsmActivation{},
        &SsmActivationList{},
    
        &IamRolePolicyAttachment{},
        &IamRolePolicyAttachmentList{},
    
        &WafGeoMatchSet{},
        &WafGeoMatchSetList{},
    
        &DmsReplicationInstance{},
        &DmsReplicationInstanceList{},
    
        &EcsService{},
        &EcsServiceList{},
    
        &OrganizationsAccount{},
        &OrganizationsAccountList{},
    
        &VpnGatewayRoutePropagation{},
        &VpnGatewayRoutePropagationList{},
    
        &ApiGatewayResource{},
        &ApiGatewayResourceList{},
    
        &SesReceiptRule{},
        &SesReceiptRuleList{},
    
        &StoragegatewayUploadBuffer{},
        &StoragegatewayUploadBufferList{},
    
        &AmiLaunchPermission{},
        &AmiLaunchPermissionList{},
    
        &EipAssociation{},
        &EipAssociationList{},
    
        &DefaultSubnet{},
        &DefaultSubnetList{},
    
        &TransferSSHKey{},
        &TransferSSHKeyList{},
    
        &CodebuildProject{},
        &CodebuildProjectList{},
    
        &DxGatewayAssociation{},
        &DxGatewayAssociationList{},
    
        &ElastictranscoderPipeline{},
        &ElastictranscoderPipelineList{},
    
        &IotPolicyAttachment{},
        &IotPolicyAttachmentList{},
    
        &RedshiftSnapshotCopyGrant{},
        &RedshiftSnapshotCopyGrantList{},
    
        &Route53QueryLog{},
        &Route53QueryLogList{},
    
        &ApiGatewayMethodResponse{},
        &ApiGatewayMethodResponseList{},
    
        &ApiGatewayVpcLink{},
        &ApiGatewayVpcLinkList{},
    
        &AutoscalingAttachment{},
        &AutoscalingAttachmentList{},
    
        &CognitoIdentityPoolRolesAttachment{},
        &CognitoIdentityPoolRolesAttachmentList{},
    
        &ProxyProtocolPolicy{},
        &ProxyProtocolPolicyList{},
    
        &PinpointApp{},
        &PinpointAppList{},
    
        &AthenaNamedQuery{},
        &AthenaNamedQueryList{},
    
        &AutoscalingGroup{},
        &AutoscalingGroupList{},
    
        &SimpledbDomain{},
        &SimpledbDomainList{},
    
        &PinpointApnsVoipSandboxChannel{},
        &PinpointApnsVoipSandboxChannelList{},
    
        &WafregionalRuleGroup{},
        &WafregionalRuleGroupList{},
    
        &GuarddutyDetector{},
        &GuarddutyDetectorList{},
    
        &KmsGrant{},
        &KmsGrantList{},
    
        &RouteTable{},
        &RouteTableList{},
    
        &StoragegatewayCachedIscsiVolume{},
        &StoragegatewayCachedIscsiVolumeList{},
    
        &DefaultVpcDHCPOptions{},
        &DefaultVpcDHCPOptionsList{},
    
        &WafSizeConstraintSet{},
        &WafSizeConstraintSetList{},
    
        &DxHostedPublicVirtualInterface{},
        &DxHostedPublicVirtualInterfaceList{},
    
        &FlowLog{},
        &FlowLogList{},
    
        &VpcDHCPOptionsAssociation{},
        &VpcDHCPOptionsAssociationList{},
    
        &LightsailDomain{},
        &LightsailDomainList{},
    
        &Route53Zone{},
        &Route53ZoneList{},
    
        &SecurityhubStandardsSubscription{},
        &SecurityhubStandardsSubscriptionList{},
    
        &SsmParameter{},
        &SsmParameterList{},
    
        &AppsyncGraphqlAPI{},
        &AppsyncGraphqlAPIList{},
    
        &BackupSelection{},
        &BackupSelectionList{},
    
        &IotCertificate{},
        &IotCertificateList{},
    
        &NeptuneSubnetGroup{},
        &NeptuneSubnetGroupList{},
    
        &RouteTableAssociation{},
        &RouteTableAssociationList{},
    
        &S3BucketNotification{},
        &S3BucketNotificationList{},
    
        &AlbListener{},
        &AlbListenerList{},
    
        &LbListenerCertificate{},
        &LbListenerCertificateList{},
    
        &ApiGatewayIntegrationResponse{},
        &ApiGatewayIntegrationResponseList{},
    
        &AppautoscalingPolicy{},
        &AppautoscalingPolicyList{},
    
        &NeptuneClusterSnapshot{},
        &NeptuneClusterSnapshotList{},
    
        &SsmMaintenanceWindow{},
        &SsmMaintenanceWindowList{},
    
        &WafRateBasedRule{},
        &WafRateBasedRuleList{},
    
        &WorklinkWebsiteCertificateAuthorityAssociation{},
        &WorklinkWebsiteCertificateAuthorityAssociationList{},
    
        &WafRule{},
        &WafRuleList{},
    
        &WorklinkFleet{},
        &WorklinkFleetList{},
    
        &AthenaDatabase{},
        &AthenaDatabaseList{},
    
        &CurReportDefinition{},
        &CurReportDefinitionList{},
    
        &DirectoryServiceDirectory{},
        &DirectoryServiceDirectoryList{},
    
        &IamPolicy{},
        &IamPolicyList{},
    
        &RedshiftSecurityGroup{},
        &RedshiftSecurityGroupList{},
    
        &VpcPeeringConnection{},
        &VpcPeeringConnectionList{},
    
        &PinpointEmailChannel{},
        &PinpointEmailChannelList{},
    
        &ElasticsearchDomain{},
        &ElasticsearchDomainList{},
    
        &GlobalacceleratorAccelerator{},
        &GlobalacceleratorAcceleratorList{},
    
        &IamGroup{},
        &IamGroupList{},
    
        &CodebuildWebhook{},
        &CodebuildWebhookList{},
    
        &LambdaAlias{},
        &LambdaAliasList{},
    
        &OpsworksNodejsAppLayer{},
        &OpsworksNodejsAppLayerList{},
    
        &ServiceDiscoveryService{},
        &ServiceDiscoveryServiceList{},
    
        &VpcIpv4CIDRBlockAssociation{},
        &VpcIpv4CIDRBlockAssociationList{},
    
        &WafRegexMatchSet{},
        &WafRegexMatchSetList{},
    
        &AutoscalingSchedule{},
        &AutoscalingScheduleList{},
    
        &InspectorResourceGroup{},
        &InspectorResourceGroupList{},
    
        &LoadBalancerBackendServerPolicy{},
        &LoadBalancerBackendServerPolicyList{},
    
        &SqsQueue{},
        &SqsQueueList{},
    
        &VpnGateway{},
        &VpnGatewayList{},
    
        &DatasyncLocationEfs{},
        &DatasyncLocationEfsList{},
    
        &DxLag{},
        &DxLagList{},
    
        &AppmeshRoute{},
        &AppmeshRouteList{},
    
        &GlueCrawler{},
        &GlueCrawlerList{},
    
        &Route53HealthCheck{},
        &Route53HealthCheckList{},
    
        &ConfigConfigurationAggregator{},
        &ConfigConfigurationAggregatorList{},
    
        &SpotInstanceRequest{},
        &SpotInstanceRequestList{},
    
        &SsmDocument{},
        &SsmDocumentList{},
    
        &DynamodbTableItem{},
        &DynamodbTableItemList{},
    
        &ApiGatewayMethod{},
        &ApiGatewayMethodList{},
    
        &KeyPair{},
        &KeyPairList{},
    
        &OpsworksPermission{},
        &OpsworksPermissionList{},
    
        &NetworkInterfaceSgAttachment{},
        &NetworkInterfaceSgAttachmentList{},
    
        &ServicecatalogPortfolio{},
        &ServicecatalogPortfolioList{},
    
        &PinpointApnsSandboxChannel{},
        &PinpointApnsSandboxChannelList{},
    
        &AmiCopy{},
        &AmiCopyList{},
    
        &AppmeshVirtualNode{},
        &AppmeshVirtualNodeList{},
    
        &DmsCertificate{},
        &DmsCertificateList{},
    
        &DxBGPPeer{},
        &DxBGPPeerList{},
    
        &GlueSecurityConfiguration{},
        &GlueSecurityConfigurationList{},
    
        &DefaultVpc{},
        &DefaultVpcList{},
    
        &ConfigConfigRule{},
        &ConfigConfigRuleList{},
    
        &IamServerCertificate{},
        &IamServerCertificateList{},
    
        &SsmPatchGroup{},
        &SsmPatchGroupList{},
    
        &VpnGatewayAttachment{},
        &VpnGatewayAttachmentList{},
    
        &WafregionalGeoMatchSet{},
        &WafregionalGeoMatchSetList{},
    
        &ElasticacheSubnetGroup{},
        &ElasticacheSubnetGroupList{},
    
        &IamInstanceProfile{},
        &IamInstanceProfileList{},
    
        &OpsworksStack{},
        &OpsworksStackList{},
    
        &CloudwatchLogDestination{},
        &CloudwatchLogDestinationList{},
    
        &CustomerGateway{},
        &CustomerGatewayList{},
    
        &DbSecurityGroup{},
        &DbSecurityGroupList{},
    
        &ServiceDiscoveryHTTPNamespace{},
        &ServiceDiscoveryHTTPNamespaceList{},
    
        &WafXssMatchSet{},
        &WafXssMatchSetList{},
    
        &EmrSecurityConfiguration{},
        &EmrSecurityConfigurationList{},
    
        &AmiFromInstance{},
        &AmiFromInstanceList{},
    
        &S3BucketMetric{},
        &S3BucketMetricList{},
    
        &CloudwatchMetricAlarm{},
        &CloudwatchMetricAlarmList{},
    
        &DxGatewayAssociationProposal{},
        &DxGatewayAssociationProposalList{},
    
        &EcrRepository{},
        &EcrRepositoryList{},
    
        &GuarddutyIpset{},
        &GuarddutyIpsetList{},
    
        &OpsworksInstance{},
        &OpsworksInstanceList{},
    
        &PinpointEventStream{},
        &PinpointEventStreamList{},
    
        &EcsTaskDefinition{},
        &EcsTaskDefinitionList{},
    
        &NeptuneClusterParameterGroup{},
        &NeptuneClusterParameterGroupList{},
    
        &OpsworksApplication{},
        &OpsworksApplicationList{},
    
        &S3AccountPublicAccessBlock{},
        &S3AccountPublicAccessBlockList{},
    
        &RdsCluster{},
        &RdsClusterList{},
    
        &SfnStateMachine{},
        &SfnStateMachineList{},
    
        &KinesisAnalyticsApplication{},
        &KinesisAnalyticsApplicationList{},
    
        &LightsailStaticIPAttachment{},
        &LightsailStaticIPAttachmentList{},
    
        &ApiGatewayGatewayResponse{},
        &ApiGatewayGatewayResponseList{},
    
        &ConfigConfigurationRecorder{},
        &ConfigConfigurationRecorderList{},
    
        &DaxParameterGroup{},
        &DaxParameterGroupList{},
    
        &ElasticBeanstalkApplication{},
        &ElasticBeanstalkApplicationList{},
    
        &EmrCluster{},
        &EmrClusterList{},
    
        &IotPolicy{},
        &IotPolicyList{},
    
        &OrganizationsOrganization{},
        &OrganizationsOrganizationList{},
    
        &ApiGatewayDocumentationVersion{},
        &ApiGatewayDocumentationVersionList{},
    
        &Ec2Fleet{},
        &Ec2FleetList{},
    
        &KmsKey{},
        &KmsKeyList{},
    
        &OpsworksCustomLayer{},
        &OpsworksCustomLayerList{},
    
        &VpnConnection{},
        &VpnConnectionList{},
    
        &RdsClusterInstance{},
        &RdsClusterInstanceList{},
    
        &SnsPlatformApplication{},
        &SnsPlatformApplicationList{},
    
        &BatchComputeEnvironment{},
        &BatchComputeEnvironmentList{},
    
        &ApiGatewayRequestValidator{},
        &ApiGatewayRequestValidatorList{},
    
        &DxPrivateVirtualInterface{},
        &DxPrivateVirtualInterfaceList{},
    
        &Ec2TransitGatewayRouteTable{},
        &Ec2TransitGatewayRouteTableList{},
    
        &VpnConnectionRoute{},
        &VpnConnectionRouteList{},
    
        &ApiGatewayUsagePlan{},
        &ApiGatewayUsagePlanList{},
    
        &IotThing{},
        &IotThingList{},
    
        &OpsworksRailsAppLayer{},
        &OpsworksRailsAppLayerList{},
    
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
