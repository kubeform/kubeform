## CognitoUserPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoUserPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoUserPoolSpec](#CognitoUserPoolSpec)***||
| `status` | ***[CognitoUserPoolStatus](#CognitoUserPoolStatus)***||
## CognitoUserPoolSpec
##### (Appears on:[CognitoUserPool](#CognitoUserPool), [CognitoUserPoolStatus](#CognitoUserPoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adminCreateUserConfig` | ***[[]CognitoUserPoolSpecAdminCreateUserConfig](#CognitoUserPoolSpecAdminCreateUserConfig)***| ***(Optional)*** |
| `aliasAttributes` | ***[]string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoVerifiedAttributes` | ***[]string***| ***(Optional)*** |
| `creationDate` | ***string***| ***(Optional)*** |
| `deviceConfiguration` | ***[[]CognitoUserPoolSpecDeviceConfiguration](#CognitoUserPoolSpecDeviceConfiguration)***| ***(Optional)*** |
| `emailConfiguration` | ***[[]CognitoUserPoolSpecEmailConfiguration](#CognitoUserPoolSpecEmailConfiguration)***| ***(Optional)*** |
| `emailVerificationMessage` | ***string***| ***(Optional)*** |
| `emailVerificationSubject` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `lambdaConfig` | ***[[]CognitoUserPoolSpecLambdaConfig](#CognitoUserPoolSpecLambdaConfig)***| ***(Optional)*** |
| `lastModifiedDate` | ***string***| ***(Optional)*** |
| `mfaConfiguration` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `passwordPolicy` | ***[[]CognitoUserPoolSpecPasswordPolicy](#CognitoUserPoolSpecPasswordPolicy)***| ***(Optional)*** |
| `schema` | ***[[]CognitoUserPoolSpecSchema](#CognitoUserPoolSpecSchema)***| ***(Optional)*** |
| `smsAuthenticationMessage` | ***string***| ***(Optional)*** |
| `smsConfiguration` | ***[[]CognitoUserPoolSpecSmsConfiguration](#CognitoUserPoolSpecSmsConfiguration)***| ***(Optional)*** |
| `smsVerificationMessage` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `userPoolAddOns` | ***[[]CognitoUserPoolSpecUserPoolAddOns](#CognitoUserPoolSpecUserPoolAddOns)***| ***(Optional)*** |
| `usernameAttributes` | ***[]string***| ***(Optional)*** |
| `verificationMessageTemplate` | ***[[]CognitoUserPoolSpecVerificationMessageTemplate](#CognitoUserPoolSpecVerificationMessageTemplate)***| ***(Optional)*** |
## CognitoUserPoolSpecAdminCreateUserConfig
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowAdminCreateUserOnly` | ***bool***| ***(Optional)*** |
| `inviteMessageTemplate` | ***[[]CognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate](#CognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate)***| ***(Optional)*** |
| `unusedAccountValidityDays` | ***int***| ***(Optional)*** |
## CognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate
##### (Appears on:[CognitoUserPoolSpecAdminCreateUserConfig](#CognitoUserPoolSpecAdminCreateUserConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `emailMessage` | ***string***| ***(Optional)*** |
| `emailSubject` | ***string***| ***(Optional)*** |
| `smsMessage` | ***string***| ***(Optional)*** |
## CognitoUserPoolSpecDeviceConfiguration
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `challengeRequiredOnNewDevice` | ***bool***| ***(Optional)*** |
| `deviceOnlyRememberedOnUserPrompt` | ***bool***| ***(Optional)*** |
## CognitoUserPoolSpecEmailConfiguration
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `replyToEmailAddress` | ***string***| ***(Optional)*** |
| `sourceArn` | ***string***| ***(Optional)*** |
## CognitoUserPoolSpecLambdaConfig
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `createAuthChallenge` | ***string***| ***(Optional)*** |
| `customMessage` | ***string***| ***(Optional)*** |
| `defineAuthChallenge` | ***string***| ***(Optional)*** |
| `postAuthentication` | ***string***| ***(Optional)*** |
| `postConfirmation` | ***string***| ***(Optional)*** |
| `preAuthentication` | ***string***| ***(Optional)*** |
| `preSignUp` | ***string***| ***(Optional)*** |
| `preTokenGeneration` | ***string***| ***(Optional)*** |
| `userMigration` | ***string***| ***(Optional)*** |
| `verifyAuthChallengeResponse` | ***string***| ***(Optional)*** |
## CognitoUserPoolSpecPasswordPolicy
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minimumLength` | ***int***| ***(Optional)*** |
| `requireLowercase` | ***bool***| ***(Optional)*** |
| `requireNumbers` | ***bool***| ***(Optional)*** |
| `requireSymbols` | ***bool***| ***(Optional)*** |
| `requireUppercase` | ***bool***| ***(Optional)*** |
## CognitoUserPoolSpecSchema
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `attributeDataType` | ***string***||
| `developerOnlyAttribute` | ***bool***| ***(Optional)*** |
| `mutable` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `numberAttributeConstraints` | ***[[]CognitoUserPoolSpecSchemaNumberAttributeConstraints](#CognitoUserPoolSpecSchemaNumberAttributeConstraints)***| ***(Optional)*** |
| `required` | ***bool***| ***(Optional)*** |
| `stringAttributeConstraints` | ***[[]CognitoUserPoolSpecSchemaStringAttributeConstraints](#CognitoUserPoolSpecSchemaStringAttributeConstraints)***| ***(Optional)*** |
## CognitoUserPoolSpecSchemaNumberAttributeConstraints
##### (Appears on:[CognitoUserPoolSpecSchema](#CognitoUserPoolSpecSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxValue` | ***string***| ***(Optional)*** |
| `minValue` | ***string***| ***(Optional)*** |
## CognitoUserPoolSpecSchemaStringAttributeConstraints
##### (Appears on:[CognitoUserPoolSpecSchema](#CognitoUserPoolSpecSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxLength` | ***string***| ***(Optional)*** |
| `minLength` | ***string***| ***(Optional)*** |
## CognitoUserPoolSpecSmsConfiguration
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `externalID` | ***string***||
| `snsCallerArn` | ***string***||
## CognitoUserPoolSpecUserPoolAddOns
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `advancedSecurityMode` | ***string***||
## CognitoUserPoolSpecVerificationMessageTemplate
##### (Appears on:[CognitoUserPoolSpec](#CognitoUserPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultEmailOption` | ***string***| ***(Optional)*** |
| `emailMessage` | ***string***| ***(Optional)*** |
| `emailMessageByLink` | ***string***| ***(Optional)*** |
| `emailSubject` | ***string***| ***(Optional)*** |
| `emailSubjectByLink` | ***string***| ***(Optional)*** |
| `smsMessage` | ***string***| ***(Optional)*** |
## CognitoUserPoolStatus
##### (Appears on:[CognitoUserPool](#CognitoUserPool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoUserPoolSpec](#CognitoUserPoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
