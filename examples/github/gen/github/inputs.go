package github

import "time"

type AuditLogOrderField string

const (
	AuditLogOrderField_CREATED_AT AuditLogOrderField = "CREATED_AT"
)

type CheckAnnotationLevel string

const (
	CheckAnnotationLevel_FAILURE CheckAnnotationLevel = "FAILURE"
	CheckAnnotationLevel_NOTICE  CheckAnnotationLevel = "NOTICE"
	CheckAnnotationLevel_WARNING CheckAnnotationLevel = "WARNING"
)

type CheckConclusionState string

const (
	CheckConclusionState_ACTION_REQUIRED CheckConclusionState = "ACTION_REQUIRED"
	CheckConclusionState_CANCELLED       CheckConclusionState = "CANCELLED"
	CheckConclusionState_FAILURE         CheckConclusionState = "FAILURE"
	CheckConclusionState_NEUTRAL         CheckConclusionState = "NEUTRAL"
	CheckConclusionState_SKIPPED         CheckConclusionState = "SKIPPED"
	CheckConclusionState_STALE           CheckConclusionState = "STALE"
	CheckConclusionState_STARTUP_FAILURE CheckConclusionState = "STARTUP_FAILURE"
	CheckConclusionState_SUCCESS         CheckConclusionState = "SUCCESS"
	CheckConclusionState_TIMED_OUT       CheckConclusionState = "TIMED_OUT"
)

type CheckRunType string

const (
	CheckRunType_ALL    CheckRunType = "ALL"
	CheckRunType_LATEST CheckRunType = "LATEST"
)

type CheckStatusState string

const (
	CheckStatusState_COMPLETED   CheckStatusState = "COMPLETED"
	CheckStatusState_IN_PROGRESS CheckStatusState = "IN_PROGRESS"
	CheckStatusState_PENDING     CheckStatusState = "PENDING"
	CheckStatusState_QUEUED      CheckStatusState = "QUEUED"
	CheckStatusState_REQUESTED   CheckStatusState = "REQUESTED"
	CheckStatusState_WAITING     CheckStatusState = "WAITING"
)

type CollaboratorAffiliation string

const (
	CollaboratorAffiliation_ALL     CollaboratorAffiliation = "ALL"
	CollaboratorAffiliation_DIRECT  CollaboratorAffiliation = "DIRECT"
	CollaboratorAffiliation_OUTSIDE CollaboratorAffiliation = "OUTSIDE"
)

type CommentAuthorAssociation string

const (
	CommentAuthorAssociation_COLLABORATOR           CommentAuthorAssociation = "COLLABORATOR"
	CommentAuthorAssociation_CONTRIBUTOR            CommentAuthorAssociation = "CONTRIBUTOR"
	CommentAuthorAssociation_FIRST_TIMER            CommentAuthorAssociation = "FIRST_TIMER"
	CommentAuthorAssociation_FIRST_TIME_CONTRIBUTOR CommentAuthorAssociation = "FIRST_TIME_CONTRIBUTOR"
	CommentAuthorAssociation_MANNEQUIN              CommentAuthorAssociation = "MANNEQUIN"
	CommentAuthorAssociation_MEMBER                 CommentAuthorAssociation = "MEMBER"
	CommentAuthorAssociation_NONE                   CommentAuthorAssociation = "NONE"
	CommentAuthorAssociation_OWNER                  CommentAuthorAssociation = "OWNER"
)

type CommentCannotUpdateReason string

const (
	CommentCannotUpdateReason_ARCHIVED                CommentCannotUpdateReason = "ARCHIVED"
	CommentCannotUpdateReason_DENIED                  CommentCannotUpdateReason = "DENIED"
	CommentCannotUpdateReason_INSUFFICIENT_ACCESS     CommentCannotUpdateReason = "INSUFFICIENT_ACCESS"
	CommentCannotUpdateReason_LOCKED                  CommentCannotUpdateReason = "LOCKED"
	CommentCannotUpdateReason_LOGIN_REQUIRED          CommentCannotUpdateReason = "LOGIN_REQUIRED"
	CommentCannotUpdateReason_MAINTENANCE             CommentCannotUpdateReason = "MAINTENANCE"
	CommentCannotUpdateReason_VERIFIED_EMAIL_REQUIRED CommentCannotUpdateReason = "VERIFIED_EMAIL_REQUIRED"
)

type CommitContributionOrderField string

const (
	CommitContributionOrderField_COMMIT_COUNT CommitContributionOrderField = "COMMIT_COUNT"
	CommitContributionOrderField_OCCURRED_AT  CommitContributionOrderField = "OCCURRED_AT"
)

type ContributionLevel string

const (
	ContributionLevel_FIRST_QUARTILE  ContributionLevel = "FIRST_QUARTILE"
	ContributionLevel_FOURTH_QUARTILE ContributionLevel = "FOURTH_QUARTILE"
	ContributionLevel_NONE            ContributionLevel = "NONE"
	ContributionLevel_SECOND_QUARTILE ContributionLevel = "SECOND_QUARTILE"
	ContributionLevel_THIRD_QUARTILE  ContributionLevel = "THIRD_QUARTILE"
)

type DefaultRepositoryPermissionField string

const (
	DefaultRepositoryPermissionField_ADMIN DefaultRepositoryPermissionField = "ADMIN"
	DefaultRepositoryPermissionField_NONE  DefaultRepositoryPermissionField = "NONE"
	DefaultRepositoryPermissionField_READ  DefaultRepositoryPermissionField = "READ"
	DefaultRepositoryPermissionField_WRITE DefaultRepositoryPermissionField = "WRITE"
)

type DeploymentOrderField string

const (
	DeploymentOrderField_CREATED_AT DeploymentOrderField = "CREATED_AT"
)

type DeploymentProtectionRuleType string

const (
	DeploymentProtectionRuleType_REQUIRED_REVIEWERS DeploymentProtectionRuleType = "REQUIRED_REVIEWERS"
	DeploymentProtectionRuleType_WAIT_TIMER         DeploymentProtectionRuleType = "WAIT_TIMER"
)

type DeploymentReviewState string

const (
	DeploymentReviewState_APPROVED DeploymentReviewState = "APPROVED"
	DeploymentReviewState_REJECTED DeploymentReviewState = "REJECTED"
)

type DeploymentState string

const (
	DeploymentState_ABANDONED   DeploymentState = "ABANDONED"
	DeploymentState_ACTIVE      DeploymentState = "ACTIVE"
	DeploymentState_DESTROYED   DeploymentState = "DESTROYED"
	DeploymentState_ERROR       DeploymentState = "ERROR"
	DeploymentState_FAILURE     DeploymentState = "FAILURE"
	DeploymentState_INACTIVE    DeploymentState = "INACTIVE"
	DeploymentState_IN_PROGRESS DeploymentState = "IN_PROGRESS"
	DeploymentState_PENDING     DeploymentState = "PENDING"
	DeploymentState_QUEUED      DeploymentState = "QUEUED"
	DeploymentState_WAITING     DeploymentState = "WAITING"
)

type DeploymentStatusState string

const (
	DeploymentStatusState_ERROR       DeploymentStatusState = "ERROR"
	DeploymentStatusState_FAILURE     DeploymentStatusState = "FAILURE"
	DeploymentStatusState_INACTIVE    DeploymentStatusState = "INACTIVE"
	DeploymentStatusState_IN_PROGRESS DeploymentStatusState = "IN_PROGRESS"
	DeploymentStatusState_PENDING     DeploymentStatusState = "PENDING"
	DeploymentStatusState_QUEUED      DeploymentStatusState = "QUEUED"
	DeploymentStatusState_SUCCESS     DeploymentStatusState = "SUCCESS"
	DeploymentStatusState_WAITING     DeploymentStatusState = "WAITING"
)

type DiffSide string

const (
	DiffSide_LEFT  DiffSide = "LEFT"
	DiffSide_RIGHT DiffSide = "RIGHT"
)

type DiscussionOrderField string

const (
	DiscussionOrderField_CREATED_AT DiscussionOrderField = "CREATED_AT"
	DiscussionOrderField_UPDATED_AT DiscussionOrderField = "UPDATED_AT"
)

type EnterpriseAdministratorInvitationOrderField string

const (
	EnterpriseAdministratorInvitationOrderField_CREATED_AT EnterpriseAdministratorInvitationOrderField = "CREATED_AT"
)

type EnterpriseAdministratorRole string

const (
	EnterpriseAdministratorRole_BILLING_MANAGER EnterpriseAdministratorRole = "BILLING_MANAGER"
	EnterpriseAdministratorRole_OWNER           EnterpriseAdministratorRole = "OWNER"
)

type EnterpriseDefaultRepositoryPermissionSettingValue string

const (
	EnterpriseDefaultRepositoryPermissionSettingValue_ADMIN     EnterpriseDefaultRepositoryPermissionSettingValue = "ADMIN"
	EnterpriseDefaultRepositoryPermissionSettingValue_NONE      EnterpriseDefaultRepositoryPermissionSettingValue = "NONE"
	EnterpriseDefaultRepositoryPermissionSettingValue_NO_POLICY EnterpriseDefaultRepositoryPermissionSettingValue = "NO_POLICY"
	EnterpriseDefaultRepositoryPermissionSettingValue_READ      EnterpriseDefaultRepositoryPermissionSettingValue = "READ"
	EnterpriseDefaultRepositoryPermissionSettingValue_WRITE     EnterpriseDefaultRepositoryPermissionSettingValue = "WRITE"
)

type EnterpriseEnabledDisabledSettingValue string

const (
	EnterpriseEnabledDisabledSettingValue_DISABLED  EnterpriseEnabledDisabledSettingValue = "DISABLED"
	EnterpriseEnabledDisabledSettingValue_ENABLED   EnterpriseEnabledDisabledSettingValue = "ENABLED"
	EnterpriseEnabledDisabledSettingValue_NO_POLICY EnterpriseEnabledDisabledSettingValue = "NO_POLICY"
)

type EnterpriseEnabledSettingValue string

const (
	EnterpriseEnabledSettingValue_ENABLED   EnterpriseEnabledSettingValue = "ENABLED"
	EnterpriseEnabledSettingValue_NO_POLICY EnterpriseEnabledSettingValue = "NO_POLICY"
)

type EnterpriseMemberOrderField string

const (
	EnterpriseMemberOrderField_CREATED_AT EnterpriseMemberOrderField = "CREATED_AT"
	EnterpriseMemberOrderField_LOGIN      EnterpriseMemberOrderField = "LOGIN"
)

type EnterpriseMembersCanCreateRepositoriesSettingValue string

const (
	EnterpriseMembersCanCreateRepositoriesSettingValue_ALL       EnterpriseMembersCanCreateRepositoriesSettingValue = "ALL"
	EnterpriseMembersCanCreateRepositoriesSettingValue_DISABLED  EnterpriseMembersCanCreateRepositoriesSettingValue = "DISABLED"
	EnterpriseMembersCanCreateRepositoriesSettingValue_NO_POLICY EnterpriseMembersCanCreateRepositoriesSettingValue = "NO_POLICY"
	EnterpriseMembersCanCreateRepositoriesSettingValue_PRIVATE   EnterpriseMembersCanCreateRepositoriesSettingValue = "PRIVATE"
	EnterpriseMembersCanCreateRepositoriesSettingValue_PUBLIC    EnterpriseMembersCanCreateRepositoriesSettingValue = "PUBLIC"
)

type EnterpriseMembersCanMakePurchasesSettingValue string

const (
	EnterpriseMembersCanMakePurchasesSettingValue_DISABLED EnterpriseMembersCanMakePurchasesSettingValue = "DISABLED"
	EnterpriseMembersCanMakePurchasesSettingValue_ENABLED  EnterpriseMembersCanMakePurchasesSettingValue = "ENABLED"
)

type EnterpriseServerInstallationOrderField string

const (
	EnterpriseServerInstallationOrderField_CREATED_AT    EnterpriseServerInstallationOrderField = "CREATED_AT"
	EnterpriseServerInstallationOrderField_CUSTOMER_NAME EnterpriseServerInstallationOrderField = "CUSTOMER_NAME"
	EnterpriseServerInstallationOrderField_HOST_NAME     EnterpriseServerInstallationOrderField = "HOST_NAME"
)

type EnterpriseServerUserAccountEmailOrderField string

const (
	EnterpriseServerUserAccountEmailOrderField_EMAIL EnterpriseServerUserAccountEmailOrderField = "EMAIL"
)

type EnterpriseServerUserAccountOrderField string

const (
	EnterpriseServerUserAccountOrderField_LOGIN             EnterpriseServerUserAccountOrderField = "LOGIN"
	EnterpriseServerUserAccountOrderField_REMOTE_CREATED_AT EnterpriseServerUserAccountOrderField = "REMOTE_CREATED_AT"
)

type EnterpriseServerUserAccountsUploadOrderField string

const (
	EnterpriseServerUserAccountsUploadOrderField_CREATED_AT EnterpriseServerUserAccountsUploadOrderField = "CREATED_AT"
)

type EnterpriseServerUserAccountsUploadSyncState string

const (
	EnterpriseServerUserAccountsUploadSyncState_FAILURE EnterpriseServerUserAccountsUploadSyncState = "FAILURE"
	EnterpriseServerUserAccountsUploadSyncState_PENDING EnterpriseServerUserAccountsUploadSyncState = "PENDING"
	EnterpriseServerUserAccountsUploadSyncState_SUCCESS EnterpriseServerUserAccountsUploadSyncState = "SUCCESS"
)

type EnterpriseUserAccountMembershipRole string

const (
	EnterpriseUserAccountMembershipRole_MEMBER EnterpriseUserAccountMembershipRole = "MEMBER"
	EnterpriseUserAccountMembershipRole_OWNER  EnterpriseUserAccountMembershipRole = "OWNER"
)

type EnterpriseUserDeployment string

const (
	EnterpriseUserDeployment_CLOUD  EnterpriseUserDeployment = "CLOUD"
	EnterpriseUserDeployment_SERVER EnterpriseUserDeployment = "SERVER"
)

type FileViewedState string

const (
	FileViewedState_DISMISSED FileViewedState = "DISMISSED"
	FileViewedState_UNVIEWED  FileViewedState = "UNVIEWED"
	FileViewedState_VIEWED    FileViewedState = "VIEWED"
)

type FundingPlatform string

const (
	FundingPlatform_COMMUNITY_BRIDGE FundingPlatform = "COMMUNITY_BRIDGE"
	FundingPlatform_CUSTOM           FundingPlatform = "CUSTOM"
	FundingPlatform_GITHUB           FundingPlatform = "GITHUB"
	FundingPlatform_ISSUEHUNT        FundingPlatform = "ISSUEHUNT"
	FundingPlatform_KO_FI            FundingPlatform = "KO_FI"
	FundingPlatform_LIBERAPAY        FundingPlatform = "LIBERAPAY"
	FundingPlatform_OPEN_COLLECTIVE  FundingPlatform = "OPEN_COLLECTIVE"
	FundingPlatform_OTECHIE          FundingPlatform = "OTECHIE"
	FundingPlatform_PATREON          FundingPlatform = "PATREON"
	FundingPlatform_TIDELIFT         FundingPlatform = "TIDELIFT"
)

type GistOrderField string

const (
	GistOrderField_CREATED_AT GistOrderField = "CREATED_AT"
	GistOrderField_PUSHED_AT  GistOrderField = "PUSHED_AT"
	GistOrderField_UPDATED_AT GistOrderField = "UPDATED_AT"
)

type GistPrivacy string

const (
	GistPrivacy_ALL    GistPrivacy = "ALL"
	GistPrivacy_PUBLIC GistPrivacy = "PUBLIC"
	GistPrivacy_SECRET GistPrivacy = "SECRET"
)

type GitSignatureState string

const (
	GitSignatureState_BAD_CERT              GitSignatureState = "BAD_CERT"
	GitSignatureState_BAD_EMAIL             GitSignatureState = "BAD_EMAIL"
	GitSignatureState_EXPIRED_KEY           GitSignatureState = "EXPIRED_KEY"
	GitSignatureState_GPGVERIFY_ERROR       GitSignatureState = "GPGVERIFY_ERROR"
	GitSignatureState_GPGVERIFY_UNAVAILABLE GitSignatureState = "GPGVERIFY_UNAVAILABLE"
	GitSignatureState_INVALID               GitSignatureState = "INVALID"
	GitSignatureState_MALFORMED_SIG         GitSignatureState = "MALFORMED_SIG"
	GitSignatureState_NOT_SIGNING_KEY       GitSignatureState = "NOT_SIGNING_KEY"
	GitSignatureState_NO_USER               GitSignatureState = "NO_USER"
	GitSignatureState_OCSP_ERROR            GitSignatureState = "OCSP_ERROR"
	GitSignatureState_OCSP_PENDING          GitSignatureState = "OCSP_PENDING"
	GitSignatureState_OCSP_REVOKED          GitSignatureState = "OCSP_REVOKED"
	GitSignatureState_UNKNOWN_KEY           GitSignatureState = "UNKNOWN_KEY"
	GitSignatureState_UNKNOWN_SIG_TYPE      GitSignatureState = "UNKNOWN_SIG_TYPE"
	GitSignatureState_UNSIGNED              GitSignatureState = "UNSIGNED"
	GitSignatureState_UNVERIFIED_EMAIL      GitSignatureState = "UNVERIFIED_EMAIL"
	GitSignatureState_VALID                 GitSignatureState = "VALID"
)

type IdentityProviderConfigurationState string

const (
	IdentityProviderConfigurationState_CONFIGURED   IdentityProviderConfigurationState = "CONFIGURED"
	IdentityProviderConfigurationState_ENFORCED     IdentityProviderConfigurationState = "ENFORCED"
	IdentityProviderConfigurationState_UNCONFIGURED IdentityProviderConfigurationState = "UNCONFIGURED"
)

type IpAllowListEnabledSettingValue string

const (
	IpAllowListEnabledSettingValue_DISABLED IpAllowListEnabledSettingValue = "DISABLED"
	IpAllowListEnabledSettingValue_ENABLED  IpAllowListEnabledSettingValue = "ENABLED"
)

type IpAllowListEntryOrderField string

const (
	IpAllowListEntryOrderField_ALLOW_LIST_VALUE IpAllowListEntryOrderField = "ALLOW_LIST_VALUE"
	IpAllowListEntryOrderField_CREATED_AT       IpAllowListEntryOrderField = "CREATED_AT"
)

type IpAllowListForInstalledAppsEnabledSettingValue string

const (
	IpAllowListForInstalledAppsEnabledSettingValue_DISABLED IpAllowListForInstalledAppsEnabledSettingValue = "DISABLED"
	IpAllowListForInstalledAppsEnabledSettingValue_ENABLED  IpAllowListForInstalledAppsEnabledSettingValue = "ENABLED"
)

type IssueCommentOrderField string

const (
	IssueCommentOrderField_UPDATED_AT IssueCommentOrderField = "UPDATED_AT"
)

type IssueOrderField string

const (
	IssueOrderField_COMMENTS   IssueOrderField = "COMMENTS"
	IssueOrderField_CREATED_AT IssueOrderField = "CREATED_AT"
	IssueOrderField_UPDATED_AT IssueOrderField = "UPDATED_AT"
)

type IssueState string

const (
	IssueState_CLOSED IssueState = "CLOSED"
	IssueState_OPEN   IssueState = "OPEN"
)

type IssueTimelineItemsItemType string

const (
	IssueTimelineItemsItemType_ADDED_TO_PROJECT_EVENT         IssueTimelineItemsItemType = "ADDED_TO_PROJECT_EVENT"
	IssueTimelineItemsItemType_ASSIGNED_EVENT                 IssueTimelineItemsItemType = "ASSIGNED_EVENT"
	IssueTimelineItemsItemType_CLOSED_EVENT                   IssueTimelineItemsItemType = "CLOSED_EVENT"
	IssueTimelineItemsItemType_COMMENT_DELETED_EVENT          IssueTimelineItemsItemType = "COMMENT_DELETED_EVENT"
	IssueTimelineItemsItemType_CONNECTED_EVENT                IssueTimelineItemsItemType = "CONNECTED_EVENT"
	IssueTimelineItemsItemType_CONVERTED_NOTE_TO_ISSUE_EVENT  IssueTimelineItemsItemType = "CONVERTED_NOTE_TO_ISSUE_EVENT"
	IssueTimelineItemsItemType_CROSS_REFERENCED_EVENT         IssueTimelineItemsItemType = "CROSS_REFERENCED_EVENT"
	IssueTimelineItemsItemType_DEMILESTONED_EVENT             IssueTimelineItemsItemType = "DEMILESTONED_EVENT"
	IssueTimelineItemsItemType_DISCONNECTED_EVENT             IssueTimelineItemsItemType = "DISCONNECTED_EVENT"
	IssueTimelineItemsItemType_ISSUE_COMMENT                  IssueTimelineItemsItemType = "ISSUE_COMMENT"
	IssueTimelineItemsItemType_LABELED_EVENT                  IssueTimelineItemsItemType = "LABELED_EVENT"
	IssueTimelineItemsItemType_LOCKED_EVENT                   IssueTimelineItemsItemType = "LOCKED_EVENT"
	IssueTimelineItemsItemType_MARKED_AS_DUPLICATE_EVENT      IssueTimelineItemsItemType = "MARKED_AS_DUPLICATE_EVENT"
	IssueTimelineItemsItemType_MENTIONED_EVENT                IssueTimelineItemsItemType = "MENTIONED_EVENT"
	IssueTimelineItemsItemType_MILESTONED_EVENT               IssueTimelineItemsItemType = "MILESTONED_EVENT"
	IssueTimelineItemsItemType_MOVED_COLUMNS_IN_PROJECT_EVENT IssueTimelineItemsItemType = "MOVED_COLUMNS_IN_PROJECT_EVENT"
	IssueTimelineItemsItemType_PINNED_EVENT                   IssueTimelineItemsItemType = "PINNED_EVENT"
	IssueTimelineItemsItemType_REFERENCED_EVENT               IssueTimelineItemsItemType = "REFERENCED_EVENT"
	IssueTimelineItemsItemType_REMOVED_FROM_PROJECT_EVENT     IssueTimelineItemsItemType = "REMOVED_FROM_PROJECT_EVENT"
	IssueTimelineItemsItemType_RENAMED_TITLE_EVENT            IssueTimelineItemsItemType = "RENAMED_TITLE_EVENT"
	IssueTimelineItemsItemType_REOPENED_EVENT                 IssueTimelineItemsItemType = "REOPENED_EVENT"
	IssueTimelineItemsItemType_SUBSCRIBED_EVENT               IssueTimelineItemsItemType = "SUBSCRIBED_EVENT"
	IssueTimelineItemsItemType_TRANSFERRED_EVENT              IssueTimelineItemsItemType = "TRANSFERRED_EVENT"
	IssueTimelineItemsItemType_UNASSIGNED_EVENT               IssueTimelineItemsItemType = "UNASSIGNED_EVENT"
	IssueTimelineItemsItemType_UNLABELED_EVENT                IssueTimelineItemsItemType = "UNLABELED_EVENT"
	IssueTimelineItemsItemType_UNLOCKED_EVENT                 IssueTimelineItemsItemType = "UNLOCKED_EVENT"
	IssueTimelineItemsItemType_UNMARKED_AS_DUPLICATE_EVENT    IssueTimelineItemsItemType = "UNMARKED_AS_DUPLICATE_EVENT"
	IssueTimelineItemsItemType_UNPINNED_EVENT                 IssueTimelineItemsItemType = "UNPINNED_EVENT"
	IssueTimelineItemsItemType_UNSUBSCRIBED_EVENT             IssueTimelineItemsItemType = "UNSUBSCRIBED_EVENT"
	IssueTimelineItemsItemType_USER_BLOCKED_EVENT             IssueTimelineItemsItemType = "USER_BLOCKED_EVENT"
)

type LabelOrderField string

const (
	LabelOrderField_CREATED_AT LabelOrderField = "CREATED_AT"
	LabelOrderField_NAME       LabelOrderField = "NAME"
)

type LanguageOrderField string

const (
	LanguageOrderField_SIZE LanguageOrderField = "SIZE"
)

type LockReason string

const (
	LockReason_OFF_TOPIC  LockReason = "OFF_TOPIC"
	LockReason_RESOLVED   LockReason = "RESOLVED"
	LockReason_SPAM       LockReason = "SPAM"
	LockReason_TOO_HEATED LockReason = "TOO_HEATED"
)

type MergeStateStatus string

const (
	MergeStateStatus_BEHIND    MergeStateStatus = "BEHIND"
	MergeStateStatus_BLOCKED   MergeStateStatus = "BLOCKED"
	MergeStateStatus_CLEAN     MergeStateStatus = "CLEAN"
	MergeStateStatus_DIRTY     MergeStateStatus = "DIRTY"
	MergeStateStatus_DRAFT     MergeStateStatus = "DRAFT"
	MergeStateStatus_HAS_HOOKS MergeStateStatus = "HAS_HOOKS"
	MergeStateStatus_UNKNOWN   MergeStateStatus = "UNKNOWN"
	MergeStateStatus_UNSTABLE  MergeStateStatus = "UNSTABLE"
)

type MergeableState string

const (
	MergeableState_CONFLICTING MergeableState = "CONFLICTING"
	MergeableState_MERGEABLE   MergeableState = "MERGEABLE"
	MergeableState_UNKNOWN     MergeableState = "UNKNOWN"
)

type MilestoneOrderField string

const (
	MilestoneOrderField_CREATED_AT MilestoneOrderField = "CREATED_AT"
	MilestoneOrderField_DUE_DATE   MilestoneOrderField = "DUE_DATE"
	MilestoneOrderField_NUMBER     MilestoneOrderField = "NUMBER"
	MilestoneOrderField_UPDATED_AT MilestoneOrderField = "UPDATED_AT"
)

type MilestoneState string

const (
	MilestoneState_CLOSED MilestoneState = "CLOSED"
	MilestoneState_OPEN   MilestoneState = "OPEN"
)

type NotificationRestrictionSettingValue string

const (
	NotificationRestrictionSettingValue_DISABLED NotificationRestrictionSettingValue = "DISABLED"
	NotificationRestrictionSettingValue_ENABLED  NotificationRestrictionSettingValue = "ENABLED"
)

type OauthApplicationCreateAuditEntryState string

const (
	OauthApplicationCreateAuditEntryState_ACTIVE           OauthApplicationCreateAuditEntryState = "ACTIVE"
	OauthApplicationCreateAuditEntryState_PENDING_DELETION OauthApplicationCreateAuditEntryState = "PENDING_DELETION"
	OauthApplicationCreateAuditEntryState_SUSPENDED        OauthApplicationCreateAuditEntryState = "SUSPENDED"
)

type OperationType string

const (
	OperationType_ACCESS         OperationType = "ACCESS"
	OperationType_AUTHENTICATION OperationType = "AUTHENTICATION"
	OperationType_CREATE         OperationType = "CREATE"
	OperationType_MODIFY         OperationType = "MODIFY"
	OperationType_REMOVE         OperationType = "REMOVE"
	OperationType_RESTORE        OperationType = "RESTORE"
	OperationType_TRANSFER       OperationType = "TRANSFER"
)

type OrderDirection string

const (
	OrderDirection_ASC  OrderDirection = "ASC"
	OrderDirection_DESC OrderDirection = "DESC"
)

type OrgAddMemberAuditEntryPermission string

const (
	OrgAddMemberAuditEntryPermission_ADMIN OrgAddMemberAuditEntryPermission = "ADMIN"
	OrgAddMemberAuditEntryPermission_READ  OrgAddMemberAuditEntryPermission = "READ"
)

type OrgCreateAuditEntryBillingPlan string

const (
	OrgCreateAuditEntryBillingPlan_BUSINESS        OrgCreateAuditEntryBillingPlan = "BUSINESS"
	OrgCreateAuditEntryBillingPlan_BUSINESS_PLUS   OrgCreateAuditEntryBillingPlan = "BUSINESS_PLUS"
	OrgCreateAuditEntryBillingPlan_FREE            OrgCreateAuditEntryBillingPlan = "FREE"
	OrgCreateAuditEntryBillingPlan_TIERED_PER_SEAT OrgCreateAuditEntryBillingPlan = "TIERED_PER_SEAT"
	OrgCreateAuditEntryBillingPlan_UNLIMITED       OrgCreateAuditEntryBillingPlan = "UNLIMITED"
)

type OrgRemoveBillingManagerAuditEntryReason string

const (
	OrgRemoveBillingManagerAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING                  OrgRemoveBillingManagerAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
	OrgRemoveBillingManagerAuditEntryReason_SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY OrgRemoveBillingManagerAuditEntryReason = "SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY"
	OrgRemoveBillingManagerAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE           OrgRemoveBillingManagerAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
)

type OrgRemoveMemberAuditEntryMembershipType string

const (
	OrgRemoveMemberAuditEntryMembershipType_ADMIN                OrgRemoveMemberAuditEntryMembershipType = "ADMIN"
	OrgRemoveMemberAuditEntryMembershipType_BILLING_MANAGER      OrgRemoveMemberAuditEntryMembershipType = "BILLING_MANAGER"
	OrgRemoveMemberAuditEntryMembershipType_DIRECT_MEMBER        OrgRemoveMemberAuditEntryMembershipType = "DIRECT_MEMBER"
	OrgRemoveMemberAuditEntryMembershipType_OUTSIDE_COLLABORATOR OrgRemoveMemberAuditEntryMembershipType = "OUTSIDE_COLLABORATOR"
	OrgRemoveMemberAuditEntryMembershipType_UNAFFILIATED         OrgRemoveMemberAuditEntryMembershipType = "UNAFFILIATED"
)

type OrgRemoveMemberAuditEntryReason string

const (
	OrgRemoveMemberAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING                  OrgRemoveMemberAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
	OrgRemoveMemberAuditEntryReason_SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY OrgRemoveMemberAuditEntryReason = "SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY"
	OrgRemoveMemberAuditEntryReason_TWO_FACTOR_ACCOUNT_RECOVERY                     OrgRemoveMemberAuditEntryReason = "TWO_FACTOR_ACCOUNT_RECOVERY"
	OrgRemoveMemberAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE           OrgRemoveMemberAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
	OrgRemoveMemberAuditEntryReason_USER_ACCOUNT_DELETED                            OrgRemoveMemberAuditEntryReason = "USER_ACCOUNT_DELETED"
)

type OrgRemoveOutsideCollaboratorAuditEntryMembershipType string

const (
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_BILLING_MANAGER      OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "BILLING_MANAGER"
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_OUTSIDE_COLLABORATOR OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "OUTSIDE_COLLABORATOR"
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_UNAFFILIATED         OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "UNAFFILIATED"
)

type OrgRemoveOutsideCollaboratorAuditEntryReason string

const (
	OrgRemoveOutsideCollaboratorAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING        OrgRemoveOutsideCollaboratorAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
	OrgRemoveOutsideCollaboratorAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE OrgRemoveOutsideCollaboratorAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
)

type OrgUpdateDefaultRepositoryPermissionAuditEntryPermission string

const (
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_ADMIN OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "ADMIN"
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_NONE  OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "NONE"
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_READ  OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "READ"
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_WRITE OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "WRITE"
)

type OrgUpdateMemberAuditEntryPermission string

const (
	OrgUpdateMemberAuditEntryPermission_ADMIN OrgUpdateMemberAuditEntryPermission = "ADMIN"
	OrgUpdateMemberAuditEntryPermission_READ  OrgUpdateMemberAuditEntryPermission = "READ"
)

type OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility string

const (
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_ALL              OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "ALL"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_INTERNAL         OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "INTERNAL"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_NONE             OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "NONE"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PRIVATE          OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PRIVATE"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PRIVATE_INTERNAL OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PRIVATE_INTERNAL"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC           OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC_INTERNAL  OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC_INTERNAL"
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC_PRIVATE   OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC_PRIVATE"
)

type OrganizationInvitationRole string

const (
	OrganizationInvitationRole_ADMIN           OrganizationInvitationRole = "ADMIN"
	OrganizationInvitationRole_BILLING_MANAGER OrganizationInvitationRole = "BILLING_MANAGER"
	OrganizationInvitationRole_DIRECT_MEMBER   OrganizationInvitationRole = "DIRECT_MEMBER"
	OrganizationInvitationRole_REINSTATE       OrganizationInvitationRole = "REINSTATE"
)

type OrganizationInvitationType string

const (
	OrganizationInvitationType_EMAIL OrganizationInvitationType = "EMAIL"
	OrganizationInvitationType_USER  OrganizationInvitationType = "USER"
)

type OrganizationMemberRole string

const (
	OrganizationMemberRole_ADMIN  OrganizationMemberRole = "ADMIN"
	OrganizationMemberRole_MEMBER OrganizationMemberRole = "MEMBER"
)

type OrganizationMembersCanCreateRepositoriesSettingValue string

const (
	OrganizationMembersCanCreateRepositoriesSettingValue_ALL      OrganizationMembersCanCreateRepositoriesSettingValue = "ALL"
	OrganizationMembersCanCreateRepositoriesSettingValue_DISABLED OrganizationMembersCanCreateRepositoriesSettingValue = "DISABLED"
	OrganizationMembersCanCreateRepositoriesSettingValue_PRIVATE  OrganizationMembersCanCreateRepositoriesSettingValue = "PRIVATE"
)

type OrganizationOrderField string

const (
	OrganizationOrderField_CREATED_AT OrganizationOrderField = "CREATED_AT"
	OrganizationOrderField_LOGIN      OrganizationOrderField = "LOGIN"
)

type PackageFileOrderField string

const (
	PackageFileOrderField_CREATED_AT PackageFileOrderField = "CREATED_AT"
)

type PackageOrderField string

const (
	PackageOrderField_CREATED_AT PackageOrderField = "CREATED_AT"
)

type PackageType string

const (
	PackageType_DEBIAN   PackageType = "DEBIAN"
	PackageType_DOCKER   PackageType = "DOCKER"
	PackageType_MAVEN    PackageType = "MAVEN"
	PackageType_NPM      PackageType = "NPM"
	PackageType_NUGET    PackageType = "NUGET"
	PackageType_PYPI     PackageType = "PYPI"
	PackageType_RUBYGEMS PackageType = "RUBYGEMS"
)

type PackageVersionOrderField string

const (
	PackageVersionOrderField_CREATED_AT PackageVersionOrderField = "CREATED_AT"
)

type PinnableItemType string

const (
	PinnableItemType_GIST         PinnableItemType = "GIST"
	PinnableItemType_ISSUE        PinnableItemType = "ISSUE"
	PinnableItemType_ORGANIZATION PinnableItemType = "ORGANIZATION"
	PinnableItemType_PROJECT      PinnableItemType = "PROJECT"
	PinnableItemType_PULL_REQUEST PinnableItemType = "PULL_REQUEST"
	PinnableItemType_REPOSITORY   PinnableItemType = "REPOSITORY"
	PinnableItemType_TEAM         PinnableItemType = "TEAM"
	PinnableItemType_USER         PinnableItemType = "USER"
)

type PinnedDiscussionGradient string

const (
	PinnedDiscussionGradient_BLUE_MINT    PinnedDiscussionGradient = "BLUE_MINT"
	PinnedDiscussionGradient_BLUE_PURPLE  PinnedDiscussionGradient = "BLUE_PURPLE"
	PinnedDiscussionGradient_PINK_BLUE    PinnedDiscussionGradient = "PINK_BLUE"
	PinnedDiscussionGradient_PURPLE_CORAL PinnedDiscussionGradient = "PURPLE_CORAL"
	PinnedDiscussionGradient_RED_ORANGE   PinnedDiscussionGradient = "RED_ORANGE"
)

type PinnedDiscussionPattern string

const (
	PinnedDiscussionPattern_CHEVRON_UP PinnedDiscussionPattern = "CHEVRON_UP"
	PinnedDiscussionPattern_DOT        PinnedDiscussionPattern = "DOT"
	PinnedDiscussionPattern_DOT_FILL   PinnedDiscussionPattern = "DOT_FILL"
	PinnedDiscussionPattern_HEART_FILL PinnedDiscussionPattern = "HEART_FILL"
	PinnedDiscussionPattern_PLUS       PinnedDiscussionPattern = "PLUS"
	PinnedDiscussionPattern_ZAP        PinnedDiscussionPattern = "ZAP"
)

type ProjectCardArchivedState string

const (
	ProjectCardArchivedState_ARCHIVED     ProjectCardArchivedState = "ARCHIVED"
	ProjectCardArchivedState_NOT_ARCHIVED ProjectCardArchivedState = "NOT_ARCHIVED"
)

type ProjectCardState string

const (
	ProjectCardState_CONTENT_ONLY ProjectCardState = "CONTENT_ONLY"
	ProjectCardState_NOTE_ONLY    ProjectCardState = "NOTE_ONLY"
	ProjectCardState_REDACTED     ProjectCardState = "REDACTED"
)

type ProjectColumnPurpose string

const (
	ProjectColumnPurpose_DONE        ProjectColumnPurpose = "DONE"
	ProjectColumnPurpose_IN_PROGRESS ProjectColumnPurpose = "IN_PROGRESS"
	ProjectColumnPurpose_TODO        ProjectColumnPurpose = "TODO"
)

type ProjectOrderField string

const (
	ProjectOrderField_CREATED_AT ProjectOrderField = "CREATED_AT"
	ProjectOrderField_NAME       ProjectOrderField = "NAME"
	ProjectOrderField_UPDATED_AT ProjectOrderField = "UPDATED_AT"
)

type ProjectState string

const (
	ProjectState_CLOSED ProjectState = "CLOSED"
	ProjectState_OPEN   ProjectState = "OPEN"
)

type ProjectTemplate string

const (
	ProjectTemplate_AUTOMATED_KANBAN_V2      ProjectTemplate = "AUTOMATED_KANBAN_V2"
	ProjectTemplate_AUTOMATED_REVIEWS_KANBAN ProjectTemplate = "AUTOMATED_REVIEWS_KANBAN"
	ProjectTemplate_BASIC_KANBAN             ProjectTemplate = "BASIC_KANBAN"
	ProjectTemplate_BUG_TRIAGE               ProjectTemplate = "BUG_TRIAGE"
)

type PullRequestMergeMethod string

const (
	PullRequestMergeMethod_MERGE  PullRequestMergeMethod = "MERGE"
	PullRequestMergeMethod_REBASE PullRequestMergeMethod = "REBASE"
	PullRequestMergeMethod_SQUASH PullRequestMergeMethod = "SQUASH"
)

type PullRequestOrderField string

const (
	PullRequestOrderField_CREATED_AT PullRequestOrderField = "CREATED_AT"
	PullRequestOrderField_UPDATED_AT PullRequestOrderField = "UPDATED_AT"
)

type PullRequestReviewCommentState string

const (
	PullRequestReviewCommentState_PENDING   PullRequestReviewCommentState = "PENDING"
	PullRequestReviewCommentState_SUBMITTED PullRequestReviewCommentState = "SUBMITTED"
)

type PullRequestReviewDecision string

const (
	PullRequestReviewDecision_APPROVED          PullRequestReviewDecision = "APPROVED"
	PullRequestReviewDecision_CHANGES_REQUESTED PullRequestReviewDecision = "CHANGES_REQUESTED"
	PullRequestReviewDecision_REVIEW_REQUIRED   PullRequestReviewDecision = "REVIEW_REQUIRED"
)

type PullRequestReviewEvent string

const (
	PullRequestReviewEvent_APPROVE         PullRequestReviewEvent = "APPROVE"
	PullRequestReviewEvent_COMMENT         PullRequestReviewEvent = "COMMENT"
	PullRequestReviewEvent_DISMISS         PullRequestReviewEvent = "DISMISS"
	PullRequestReviewEvent_REQUEST_CHANGES PullRequestReviewEvent = "REQUEST_CHANGES"
)

type PullRequestReviewState string

const (
	PullRequestReviewState_APPROVED          PullRequestReviewState = "APPROVED"
	PullRequestReviewState_CHANGES_REQUESTED PullRequestReviewState = "CHANGES_REQUESTED"
	PullRequestReviewState_COMMENTED         PullRequestReviewState = "COMMENTED"
	PullRequestReviewState_DISMISSED         PullRequestReviewState = "DISMISSED"
	PullRequestReviewState_PENDING           PullRequestReviewState = "PENDING"
)

type PullRequestState string

const (
	PullRequestState_CLOSED PullRequestState = "CLOSED"
	PullRequestState_MERGED PullRequestState = "MERGED"
	PullRequestState_OPEN   PullRequestState = "OPEN"
)

type PullRequestTimelineItemsItemType string

const (
	PullRequestTimelineItemsItemType_ADDED_TO_PROJECT_EVENT                PullRequestTimelineItemsItemType = "ADDED_TO_PROJECT_EVENT"
	PullRequestTimelineItemsItemType_ASSIGNED_EVENT                        PullRequestTimelineItemsItemType = "ASSIGNED_EVENT"
	PullRequestTimelineItemsItemType_AUTOMATIC_BASE_CHANGE_FAILED_EVENT    PullRequestTimelineItemsItemType = "AUTOMATIC_BASE_CHANGE_FAILED_EVENT"
	PullRequestTimelineItemsItemType_AUTOMATIC_BASE_CHANGE_SUCCEEDED_EVENT PullRequestTimelineItemsItemType = "AUTOMATIC_BASE_CHANGE_SUCCEEDED_EVENT"
	PullRequestTimelineItemsItemType_AUTO_MERGE_DISABLED_EVENT             PullRequestTimelineItemsItemType = "AUTO_MERGE_DISABLED_EVENT"
	PullRequestTimelineItemsItemType_AUTO_MERGE_ENABLED_EVENT              PullRequestTimelineItemsItemType = "AUTO_MERGE_ENABLED_EVENT"
	PullRequestTimelineItemsItemType_AUTO_REBASE_ENABLED_EVENT             PullRequestTimelineItemsItemType = "AUTO_REBASE_ENABLED_EVENT"
	PullRequestTimelineItemsItemType_AUTO_SQUASH_ENABLED_EVENT             PullRequestTimelineItemsItemType = "AUTO_SQUASH_ENABLED_EVENT"
	PullRequestTimelineItemsItemType_BASE_REF_CHANGED_EVENT                PullRequestTimelineItemsItemType = "BASE_REF_CHANGED_EVENT"
	PullRequestTimelineItemsItemType_BASE_REF_DELETED_EVENT                PullRequestTimelineItemsItemType = "BASE_REF_DELETED_EVENT"
	PullRequestTimelineItemsItemType_BASE_REF_FORCE_PUSHED_EVENT           PullRequestTimelineItemsItemType = "BASE_REF_FORCE_PUSHED_EVENT"
	PullRequestTimelineItemsItemType_CLOSED_EVENT                          PullRequestTimelineItemsItemType = "CLOSED_EVENT"
	PullRequestTimelineItemsItemType_COMMENT_DELETED_EVENT                 PullRequestTimelineItemsItemType = "COMMENT_DELETED_EVENT"
	PullRequestTimelineItemsItemType_CONNECTED_EVENT                       PullRequestTimelineItemsItemType = "CONNECTED_EVENT"
	PullRequestTimelineItemsItemType_CONVERTED_NOTE_TO_ISSUE_EVENT         PullRequestTimelineItemsItemType = "CONVERTED_NOTE_TO_ISSUE_EVENT"
	PullRequestTimelineItemsItemType_CONVERT_TO_DRAFT_EVENT                PullRequestTimelineItemsItemType = "CONVERT_TO_DRAFT_EVENT"
	PullRequestTimelineItemsItemType_CROSS_REFERENCED_EVENT                PullRequestTimelineItemsItemType = "CROSS_REFERENCED_EVENT"
	PullRequestTimelineItemsItemType_DEMILESTONED_EVENT                    PullRequestTimelineItemsItemType = "DEMILESTONED_EVENT"
	PullRequestTimelineItemsItemType_DEPLOYED_EVENT                        PullRequestTimelineItemsItemType = "DEPLOYED_EVENT"
	PullRequestTimelineItemsItemType_DEPLOYMENT_ENVIRONMENT_CHANGED_EVENT  PullRequestTimelineItemsItemType = "DEPLOYMENT_ENVIRONMENT_CHANGED_EVENT"
	PullRequestTimelineItemsItemType_DISCONNECTED_EVENT                    PullRequestTimelineItemsItemType = "DISCONNECTED_EVENT"
	PullRequestTimelineItemsItemType_HEAD_REF_DELETED_EVENT                PullRequestTimelineItemsItemType = "HEAD_REF_DELETED_EVENT"
	PullRequestTimelineItemsItemType_HEAD_REF_FORCE_PUSHED_EVENT           PullRequestTimelineItemsItemType = "HEAD_REF_FORCE_PUSHED_EVENT"
	PullRequestTimelineItemsItemType_HEAD_REF_RESTORED_EVENT               PullRequestTimelineItemsItemType = "HEAD_REF_RESTORED_EVENT"
	PullRequestTimelineItemsItemType_ISSUE_COMMENT                         PullRequestTimelineItemsItemType = "ISSUE_COMMENT"
	PullRequestTimelineItemsItemType_LABELED_EVENT                         PullRequestTimelineItemsItemType = "LABELED_EVENT"
	PullRequestTimelineItemsItemType_LOCKED_EVENT                          PullRequestTimelineItemsItemType = "LOCKED_EVENT"
	PullRequestTimelineItemsItemType_MARKED_AS_DUPLICATE_EVENT             PullRequestTimelineItemsItemType = "MARKED_AS_DUPLICATE_EVENT"
	PullRequestTimelineItemsItemType_MENTIONED_EVENT                       PullRequestTimelineItemsItemType = "MENTIONED_EVENT"
	PullRequestTimelineItemsItemType_MERGED_EVENT                          PullRequestTimelineItemsItemType = "MERGED_EVENT"
	PullRequestTimelineItemsItemType_MILESTONED_EVENT                      PullRequestTimelineItemsItemType = "MILESTONED_EVENT"
	PullRequestTimelineItemsItemType_MOVED_COLUMNS_IN_PROJECT_EVENT        PullRequestTimelineItemsItemType = "MOVED_COLUMNS_IN_PROJECT_EVENT"
	PullRequestTimelineItemsItemType_PINNED_EVENT                          PullRequestTimelineItemsItemType = "PINNED_EVENT"
	PullRequestTimelineItemsItemType_PULL_REQUEST_COMMIT                   PullRequestTimelineItemsItemType = "PULL_REQUEST_COMMIT"
	PullRequestTimelineItemsItemType_PULL_REQUEST_COMMIT_COMMENT_THREAD    PullRequestTimelineItemsItemType = "PULL_REQUEST_COMMIT_COMMENT_THREAD"
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVIEW                   PullRequestTimelineItemsItemType = "PULL_REQUEST_REVIEW"
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVIEW_THREAD            PullRequestTimelineItemsItemType = "PULL_REQUEST_REVIEW_THREAD"
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVISION_MARKER          PullRequestTimelineItemsItemType = "PULL_REQUEST_REVISION_MARKER"
	PullRequestTimelineItemsItemType_READY_FOR_REVIEW_EVENT                PullRequestTimelineItemsItemType = "READY_FOR_REVIEW_EVENT"
	PullRequestTimelineItemsItemType_REFERENCED_EVENT                      PullRequestTimelineItemsItemType = "REFERENCED_EVENT"
	PullRequestTimelineItemsItemType_REMOVED_FROM_PROJECT_EVENT            PullRequestTimelineItemsItemType = "REMOVED_FROM_PROJECT_EVENT"
	PullRequestTimelineItemsItemType_RENAMED_TITLE_EVENT                   PullRequestTimelineItemsItemType = "RENAMED_TITLE_EVENT"
	PullRequestTimelineItemsItemType_REOPENED_EVENT                        PullRequestTimelineItemsItemType = "REOPENED_EVENT"
	PullRequestTimelineItemsItemType_REVIEW_DISMISSED_EVENT                PullRequestTimelineItemsItemType = "REVIEW_DISMISSED_EVENT"
	PullRequestTimelineItemsItemType_REVIEW_REQUESTED_EVENT                PullRequestTimelineItemsItemType = "REVIEW_REQUESTED_EVENT"
	PullRequestTimelineItemsItemType_REVIEW_REQUEST_REMOVED_EVENT          PullRequestTimelineItemsItemType = "REVIEW_REQUEST_REMOVED_EVENT"
	PullRequestTimelineItemsItemType_SUBSCRIBED_EVENT                      PullRequestTimelineItemsItemType = "SUBSCRIBED_EVENT"
	PullRequestTimelineItemsItemType_TRANSFERRED_EVENT                     PullRequestTimelineItemsItemType = "TRANSFERRED_EVENT"
	PullRequestTimelineItemsItemType_UNASSIGNED_EVENT                      PullRequestTimelineItemsItemType = "UNASSIGNED_EVENT"
	PullRequestTimelineItemsItemType_UNLABELED_EVENT                       PullRequestTimelineItemsItemType = "UNLABELED_EVENT"
	PullRequestTimelineItemsItemType_UNLOCKED_EVENT                        PullRequestTimelineItemsItemType = "UNLOCKED_EVENT"
	PullRequestTimelineItemsItemType_UNMARKED_AS_DUPLICATE_EVENT           PullRequestTimelineItemsItemType = "UNMARKED_AS_DUPLICATE_EVENT"
	PullRequestTimelineItemsItemType_UNPINNED_EVENT                        PullRequestTimelineItemsItemType = "UNPINNED_EVENT"
	PullRequestTimelineItemsItemType_UNSUBSCRIBED_EVENT                    PullRequestTimelineItemsItemType = "UNSUBSCRIBED_EVENT"
	PullRequestTimelineItemsItemType_USER_BLOCKED_EVENT                    PullRequestTimelineItemsItemType = "USER_BLOCKED_EVENT"
)

type PullRequestUpdateState string

const (
	PullRequestUpdateState_CLOSED PullRequestUpdateState = "CLOSED"
	PullRequestUpdateState_OPEN   PullRequestUpdateState = "OPEN"
)

type ReactionContent string

const (
	ReactionContent_CONFUSED    ReactionContent = "CONFUSED"
	ReactionContent_EYES        ReactionContent = "EYES"
	ReactionContent_HEART       ReactionContent = "HEART"
	ReactionContent_HOORAY      ReactionContent = "HOORAY"
	ReactionContent_LAUGH       ReactionContent = "LAUGH"
	ReactionContent_ROCKET      ReactionContent = "ROCKET"
	ReactionContent_THUMBS_DOWN ReactionContent = "THUMBS_DOWN"
	ReactionContent_THUMBS_UP   ReactionContent = "THUMBS_UP"
)

type ReactionOrderField string

const (
	ReactionOrderField_CREATED_AT ReactionOrderField = "CREATED_AT"
)

type RefOrderField string

const (
	RefOrderField_ALPHABETICAL    RefOrderField = "ALPHABETICAL"
	RefOrderField_TAG_COMMIT_DATE RefOrderField = "TAG_COMMIT_DATE"
)

type ReleaseOrderField string

const (
	ReleaseOrderField_CREATED_AT ReleaseOrderField = "CREATED_AT"
	ReleaseOrderField_NAME       ReleaseOrderField = "NAME"
)

type RepoAccessAuditEntryVisibility string

const (
	RepoAccessAuditEntryVisibility_INTERNAL RepoAccessAuditEntryVisibility = "INTERNAL"
	RepoAccessAuditEntryVisibility_PRIVATE  RepoAccessAuditEntryVisibility = "PRIVATE"
	RepoAccessAuditEntryVisibility_PUBLIC   RepoAccessAuditEntryVisibility = "PUBLIC"
)

type RepoAddMemberAuditEntryVisibility string

const (
	RepoAddMemberAuditEntryVisibility_INTERNAL RepoAddMemberAuditEntryVisibility = "INTERNAL"
	RepoAddMemberAuditEntryVisibility_PRIVATE  RepoAddMemberAuditEntryVisibility = "PRIVATE"
	RepoAddMemberAuditEntryVisibility_PUBLIC   RepoAddMemberAuditEntryVisibility = "PUBLIC"
)

type RepoArchivedAuditEntryVisibility string

const (
	RepoArchivedAuditEntryVisibility_INTERNAL RepoArchivedAuditEntryVisibility = "INTERNAL"
	RepoArchivedAuditEntryVisibility_PRIVATE  RepoArchivedAuditEntryVisibility = "PRIVATE"
	RepoArchivedAuditEntryVisibility_PUBLIC   RepoArchivedAuditEntryVisibility = "PUBLIC"
)

type RepoChangeMergeSettingAuditEntryMergeType string

const (
	RepoChangeMergeSettingAuditEntryMergeType_MERGE  RepoChangeMergeSettingAuditEntryMergeType = "MERGE"
	RepoChangeMergeSettingAuditEntryMergeType_REBASE RepoChangeMergeSettingAuditEntryMergeType = "REBASE"
	RepoChangeMergeSettingAuditEntryMergeType_SQUASH RepoChangeMergeSettingAuditEntryMergeType = "SQUASH"
)

type RepoCreateAuditEntryVisibility string

const (
	RepoCreateAuditEntryVisibility_INTERNAL RepoCreateAuditEntryVisibility = "INTERNAL"
	RepoCreateAuditEntryVisibility_PRIVATE  RepoCreateAuditEntryVisibility = "PRIVATE"
	RepoCreateAuditEntryVisibility_PUBLIC   RepoCreateAuditEntryVisibility = "PUBLIC"
)

type RepoDestroyAuditEntryVisibility string

const (
	RepoDestroyAuditEntryVisibility_INTERNAL RepoDestroyAuditEntryVisibility = "INTERNAL"
	RepoDestroyAuditEntryVisibility_PRIVATE  RepoDestroyAuditEntryVisibility = "PRIVATE"
	RepoDestroyAuditEntryVisibility_PUBLIC   RepoDestroyAuditEntryVisibility = "PUBLIC"
)

type RepoRemoveMemberAuditEntryVisibility string

const (
	RepoRemoveMemberAuditEntryVisibility_INTERNAL RepoRemoveMemberAuditEntryVisibility = "INTERNAL"
	RepoRemoveMemberAuditEntryVisibility_PRIVATE  RepoRemoveMemberAuditEntryVisibility = "PRIVATE"
	RepoRemoveMemberAuditEntryVisibility_PUBLIC   RepoRemoveMemberAuditEntryVisibility = "PUBLIC"
)

type ReportedContentClassifiers string

const (
	ReportedContentClassifiers_ABUSE     ReportedContentClassifiers = "ABUSE"
	ReportedContentClassifiers_DUPLICATE ReportedContentClassifiers = "DUPLICATE"
	ReportedContentClassifiers_OFF_TOPIC ReportedContentClassifiers = "OFF_TOPIC"
	ReportedContentClassifiers_OUTDATED  ReportedContentClassifiers = "OUTDATED"
	ReportedContentClassifiers_RESOLVED  ReportedContentClassifiers = "RESOLVED"
	ReportedContentClassifiers_SPAM      ReportedContentClassifiers = "SPAM"
)

type RepositoryAffiliation string

const (
	RepositoryAffiliation_COLLABORATOR        RepositoryAffiliation = "COLLABORATOR"
	RepositoryAffiliation_ORGANIZATION_MEMBER RepositoryAffiliation = "ORGANIZATION_MEMBER"
	RepositoryAffiliation_OWNER               RepositoryAffiliation = "OWNER"
)

type RepositoryContributionType string

const (
	RepositoryContributionType_COMMIT              RepositoryContributionType = "COMMIT"
	RepositoryContributionType_ISSUE               RepositoryContributionType = "ISSUE"
	RepositoryContributionType_PULL_REQUEST        RepositoryContributionType = "PULL_REQUEST"
	RepositoryContributionType_PULL_REQUEST_REVIEW RepositoryContributionType = "PULL_REQUEST_REVIEW"
	RepositoryContributionType_REPOSITORY          RepositoryContributionType = "REPOSITORY"
)

type RepositoryInteractionLimit string

const (
	RepositoryInteractionLimit_COLLABORATORS_ONLY RepositoryInteractionLimit = "COLLABORATORS_ONLY"
	RepositoryInteractionLimit_CONTRIBUTORS_ONLY  RepositoryInteractionLimit = "CONTRIBUTORS_ONLY"
	RepositoryInteractionLimit_EXISTING_USERS     RepositoryInteractionLimit = "EXISTING_USERS"
	RepositoryInteractionLimit_NO_LIMIT           RepositoryInteractionLimit = "NO_LIMIT"
)

type RepositoryInteractionLimitExpiry string

const (
	RepositoryInteractionLimitExpiry_ONE_DAY    RepositoryInteractionLimitExpiry = "ONE_DAY"
	RepositoryInteractionLimitExpiry_ONE_MONTH  RepositoryInteractionLimitExpiry = "ONE_MONTH"
	RepositoryInteractionLimitExpiry_ONE_WEEK   RepositoryInteractionLimitExpiry = "ONE_WEEK"
	RepositoryInteractionLimitExpiry_SIX_MONTHS RepositoryInteractionLimitExpiry = "SIX_MONTHS"
	RepositoryInteractionLimitExpiry_THREE_DAYS RepositoryInteractionLimitExpiry = "THREE_DAYS"
)

type RepositoryInteractionLimitOrigin string

const (
	RepositoryInteractionLimitOrigin_ORGANIZATION RepositoryInteractionLimitOrigin = "ORGANIZATION"
	RepositoryInteractionLimitOrigin_REPOSITORY   RepositoryInteractionLimitOrigin = "REPOSITORY"
	RepositoryInteractionLimitOrigin_USER         RepositoryInteractionLimitOrigin = "USER"
)

type RepositoryInvitationOrderField string

const (
	RepositoryInvitationOrderField_CREATED_AT    RepositoryInvitationOrderField = "CREATED_AT"
	RepositoryInvitationOrderField_INVITEE_LOGIN RepositoryInvitationOrderField = "INVITEE_LOGIN"
)

type RepositoryLockReason string

const (
	RepositoryLockReason_BILLING   RepositoryLockReason = "BILLING"
	RepositoryLockReason_MIGRATING RepositoryLockReason = "MIGRATING"
	RepositoryLockReason_MOVING    RepositoryLockReason = "MOVING"
	RepositoryLockReason_RENAME    RepositoryLockReason = "RENAME"
)

type RepositoryOrderField string

const (
	RepositoryOrderField_CREATED_AT RepositoryOrderField = "CREATED_AT"
	RepositoryOrderField_NAME       RepositoryOrderField = "NAME"
	RepositoryOrderField_PUSHED_AT  RepositoryOrderField = "PUSHED_AT"
	RepositoryOrderField_STARGAZERS RepositoryOrderField = "STARGAZERS"
	RepositoryOrderField_UPDATED_AT RepositoryOrderField = "UPDATED_AT"
)

type RepositoryPermission string

const (
	RepositoryPermission_ADMIN    RepositoryPermission = "ADMIN"
	RepositoryPermission_MAINTAIN RepositoryPermission = "MAINTAIN"
	RepositoryPermission_READ     RepositoryPermission = "READ"
	RepositoryPermission_TRIAGE   RepositoryPermission = "TRIAGE"
	RepositoryPermission_WRITE    RepositoryPermission = "WRITE"
)

type RepositoryPrivacy string

const (
	RepositoryPrivacy_PRIVATE RepositoryPrivacy = "PRIVATE"
	RepositoryPrivacy_PUBLIC  RepositoryPrivacy = "PUBLIC"
)

type RepositoryVisibility string

const (
	RepositoryVisibility_INTERNAL RepositoryVisibility = "INTERNAL"
	RepositoryVisibility_PRIVATE  RepositoryVisibility = "PRIVATE"
	RepositoryVisibility_PUBLIC   RepositoryVisibility = "PUBLIC"
)

type RequestableCheckStatusState string

const (
	RequestableCheckStatusState_COMPLETED   RequestableCheckStatusState = "COMPLETED"
	RequestableCheckStatusState_IN_PROGRESS RequestableCheckStatusState = "IN_PROGRESS"
	RequestableCheckStatusState_PENDING     RequestableCheckStatusState = "PENDING"
	RequestableCheckStatusState_QUEUED      RequestableCheckStatusState = "QUEUED"
	RequestableCheckStatusState_WAITING     RequestableCheckStatusState = "WAITING"
)

type SamlDigestAlgorithm string

const (
	SamlDigestAlgorithm_SHA1   SamlDigestAlgorithm = "SHA1"
	SamlDigestAlgorithm_SHA256 SamlDigestAlgorithm = "SHA256"
	SamlDigestAlgorithm_SHA384 SamlDigestAlgorithm = "SHA384"
	SamlDigestAlgorithm_SHA512 SamlDigestAlgorithm = "SHA512"
)

type SamlSignatureAlgorithm string

const (
	SamlSignatureAlgorithm_RSA_SHA1   SamlSignatureAlgorithm = "RSA_SHA1"
	SamlSignatureAlgorithm_RSA_SHA256 SamlSignatureAlgorithm = "RSA_SHA256"
	SamlSignatureAlgorithm_RSA_SHA384 SamlSignatureAlgorithm = "RSA_SHA384"
	SamlSignatureAlgorithm_RSA_SHA512 SamlSignatureAlgorithm = "RSA_SHA512"
)

type SavedReplyOrderField string

const (
	SavedReplyOrderField_UPDATED_AT SavedReplyOrderField = "UPDATED_AT"
)

type SearchType string

const (
	SearchType_DISCUSSION SearchType = "DISCUSSION"
	SearchType_ISSUE      SearchType = "ISSUE"
	SearchType_REPOSITORY SearchType = "REPOSITORY"
	SearchType_USER       SearchType = "USER"
)

type SecurityAdvisoryEcosystem string

const (
	SecurityAdvisoryEcosystem_COMPOSER SecurityAdvisoryEcosystem = "COMPOSER"
	SecurityAdvisoryEcosystem_GO       SecurityAdvisoryEcosystem = "GO"
	SecurityAdvisoryEcosystem_MAVEN    SecurityAdvisoryEcosystem = "MAVEN"
	SecurityAdvisoryEcosystem_NPM      SecurityAdvisoryEcosystem = "NPM"
	SecurityAdvisoryEcosystem_NUGET    SecurityAdvisoryEcosystem = "NUGET"
	SecurityAdvisoryEcosystem_OTHER    SecurityAdvisoryEcosystem = "OTHER"
	SecurityAdvisoryEcosystem_PIP      SecurityAdvisoryEcosystem = "PIP"
	SecurityAdvisoryEcosystem_RUBYGEMS SecurityAdvisoryEcosystem = "RUBYGEMS"
)

type SecurityAdvisoryIdentifierType string

const (
	SecurityAdvisoryIdentifierType_CVE  SecurityAdvisoryIdentifierType = "CVE"
	SecurityAdvisoryIdentifierType_GHSA SecurityAdvisoryIdentifierType = "GHSA"
)

type SecurityAdvisoryOrderField string

const (
	SecurityAdvisoryOrderField_PUBLISHED_AT SecurityAdvisoryOrderField = "PUBLISHED_AT"
	SecurityAdvisoryOrderField_UPDATED_AT   SecurityAdvisoryOrderField = "UPDATED_AT"
)

type SecurityAdvisorySeverity string

const (
	SecurityAdvisorySeverity_CRITICAL SecurityAdvisorySeverity = "CRITICAL"
	SecurityAdvisorySeverity_HIGH     SecurityAdvisorySeverity = "HIGH"
	SecurityAdvisorySeverity_LOW      SecurityAdvisorySeverity = "LOW"
	SecurityAdvisorySeverity_MODERATE SecurityAdvisorySeverity = "MODERATE"
)

type SecurityVulnerabilityOrderField string

const (
	SecurityVulnerabilityOrderField_UPDATED_AT SecurityVulnerabilityOrderField = "UPDATED_AT"
)

type SponsorableOrderField string

const (
	SponsorableOrderField_LOGIN SponsorableOrderField = "LOGIN"
)

type SponsorsActivityAction string

const (
	SponsorsActivityAction_CANCELLED_SPONSORSHIP  SponsorsActivityAction = "CANCELLED_SPONSORSHIP"
	SponsorsActivityAction_NEW_SPONSORSHIP        SponsorsActivityAction = "NEW_SPONSORSHIP"
	SponsorsActivityAction_PENDING_CHANGE         SponsorsActivityAction = "PENDING_CHANGE"
	SponsorsActivityAction_REFUND                 SponsorsActivityAction = "REFUND"
	SponsorsActivityAction_SPONSOR_MATCH_DISABLED SponsorsActivityAction = "SPONSOR_MATCH_DISABLED"
	SponsorsActivityAction_TIER_CHANGE            SponsorsActivityAction = "TIER_CHANGE"
)

type SponsorsActivityOrderField string

const (
	SponsorsActivityOrderField_TIMESTAMP SponsorsActivityOrderField = "TIMESTAMP"
)

type SponsorsActivityPeriod string

const (
	SponsorsActivityPeriod_ALL   SponsorsActivityPeriod = "ALL"
	SponsorsActivityPeriod_DAY   SponsorsActivityPeriod = "DAY"
	SponsorsActivityPeriod_MONTH SponsorsActivityPeriod = "MONTH"
	SponsorsActivityPeriod_WEEK  SponsorsActivityPeriod = "WEEK"
)

type SponsorsGoalKind string

const (
	SponsorsGoalKind_MONTHLY_SPONSORSHIP_AMOUNT SponsorsGoalKind = "MONTHLY_SPONSORSHIP_AMOUNT"
	SponsorsGoalKind_TOTAL_SPONSORS_COUNT       SponsorsGoalKind = "TOTAL_SPONSORS_COUNT"
)

type SponsorsTierOrderField string

const (
	SponsorsTierOrderField_CREATED_AT             SponsorsTierOrderField = "CREATED_AT"
	SponsorsTierOrderField_MONTHLY_PRICE_IN_CENTS SponsorsTierOrderField = "MONTHLY_PRICE_IN_CENTS"
)

type SponsorshipOrderField string

const (
	SponsorshipOrderField_CREATED_AT SponsorshipOrderField = "CREATED_AT"
)

type SponsorshipPrivacy string

const (
	SponsorshipPrivacy_PRIVATE SponsorshipPrivacy = "PRIVATE"
	SponsorshipPrivacy_PUBLIC  SponsorshipPrivacy = "PUBLIC"
)

type StarOrderField string

const (
	StarOrderField_STARRED_AT StarOrderField = "STARRED_AT"
)

type StatusState string

const (
	StatusState_ERROR    StatusState = "ERROR"
	StatusState_EXPECTED StatusState = "EXPECTED"
	StatusState_FAILURE  StatusState = "FAILURE"
	StatusState_PENDING  StatusState = "PENDING"
	StatusState_SUCCESS  StatusState = "SUCCESS"
)

type SubscriptionState string

const (
	SubscriptionState_IGNORED      SubscriptionState = "IGNORED"
	SubscriptionState_SUBSCRIBED   SubscriptionState = "SUBSCRIBED"
	SubscriptionState_UNSUBSCRIBED SubscriptionState = "UNSUBSCRIBED"
)

type TeamDiscussionCommentOrderField string

const (
	TeamDiscussionCommentOrderField_NUMBER TeamDiscussionCommentOrderField = "NUMBER"
)

type TeamDiscussionOrderField string

const (
	TeamDiscussionOrderField_CREATED_AT TeamDiscussionOrderField = "CREATED_AT"
)

type TeamMemberOrderField string

const (
	TeamMemberOrderField_CREATED_AT TeamMemberOrderField = "CREATED_AT"
	TeamMemberOrderField_LOGIN      TeamMemberOrderField = "LOGIN"
)

type TeamMemberRole string

const (
	TeamMemberRole_MAINTAINER TeamMemberRole = "MAINTAINER"
	TeamMemberRole_MEMBER     TeamMemberRole = "MEMBER"
)

type TeamMembershipType string

const (
	TeamMembershipType_ALL        TeamMembershipType = "ALL"
	TeamMembershipType_CHILD_TEAM TeamMembershipType = "CHILD_TEAM"
	TeamMembershipType_IMMEDIATE  TeamMembershipType = "IMMEDIATE"
)

type TeamOrderField string

const (
	TeamOrderField_NAME TeamOrderField = "NAME"
)

type TeamPrivacy string

const (
	TeamPrivacy_SECRET  TeamPrivacy = "SECRET"
	TeamPrivacy_VISIBLE TeamPrivacy = "VISIBLE"
)

type TeamRepositoryOrderField string

const (
	TeamRepositoryOrderField_CREATED_AT TeamRepositoryOrderField = "CREATED_AT"
	TeamRepositoryOrderField_NAME       TeamRepositoryOrderField = "NAME"
	TeamRepositoryOrderField_PERMISSION TeamRepositoryOrderField = "PERMISSION"
	TeamRepositoryOrderField_PUSHED_AT  TeamRepositoryOrderField = "PUSHED_AT"
	TeamRepositoryOrderField_STARGAZERS TeamRepositoryOrderField = "STARGAZERS"
	TeamRepositoryOrderField_UPDATED_AT TeamRepositoryOrderField = "UPDATED_AT"
)

type TeamReviewAssignmentAlgorithm string

const (
	TeamReviewAssignmentAlgorithm_LOAD_BALANCE TeamReviewAssignmentAlgorithm = "LOAD_BALANCE"
	TeamReviewAssignmentAlgorithm_ROUND_ROBIN  TeamReviewAssignmentAlgorithm = "ROUND_ROBIN"
)

type TeamRole string

const (
	TeamRole_ADMIN  TeamRole = "ADMIN"
	TeamRole_MEMBER TeamRole = "MEMBER"
)

type TopicSuggestionDeclineReason string

const (
	TopicSuggestionDeclineReason_NOT_RELEVANT        TopicSuggestionDeclineReason = "NOT_RELEVANT"
	TopicSuggestionDeclineReason_PERSONAL_PREFERENCE TopicSuggestionDeclineReason = "PERSONAL_PREFERENCE"
	TopicSuggestionDeclineReason_TOO_GENERAL         TopicSuggestionDeclineReason = "TOO_GENERAL"
	TopicSuggestionDeclineReason_TOO_SPECIFIC        TopicSuggestionDeclineReason = "TOO_SPECIFIC"
)

type UserBlockDuration string

const (
	UserBlockDuration_ONE_DAY    UserBlockDuration = "ONE_DAY"
	UserBlockDuration_ONE_MONTH  UserBlockDuration = "ONE_MONTH"
	UserBlockDuration_ONE_WEEK   UserBlockDuration = "ONE_WEEK"
	UserBlockDuration_PERMANENT  UserBlockDuration = "PERMANENT"
	UserBlockDuration_THREE_DAYS UserBlockDuration = "THREE_DAYS"
)

type UserStatusOrderField string

const (
	UserStatusOrderField_UPDATED_AT UserStatusOrderField = "UPDATED_AT"
)

type VerifiableDomainOrderField string

const (
	VerifiableDomainOrderField_CREATED_AT VerifiableDomainOrderField = "CREATED_AT"
	VerifiableDomainOrderField_DOMAIN     VerifiableDomainOrderField = "DOMAIN"
)

type AcceptEnterpriseAdministratorInvitationInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	InvitationId     string  `json:"invitationId"`
}

type AcceptTopicSuggestionInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	RepositoryId     string  `json:"repositoryId"`
}

type AddAssigneesToAssignableInput struct {
	AssignableId     string   `json:"assignableId"`
	AssigneeIds      []string `json:"assigneeIds"`
	ClientMutationId *string  `json:"clientMutationId"`
}

type AddCommentInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	SubjectId        string  `json:"subjectId"`
}

type AddDiscussionCommentInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	DiscussionId     string  `json:"discussionId"`
	ReplyToId        *string `json:"replyToId"`
}

type AddEnterpriseSupportEntitlementInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
	Login            string  `json:"login"`
}

type AddLabelsToLabelableInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	LabelIds         []string `json:"labelIds"`
	LabelableId      string   `json:"labelableId"`
}

type AddProjectCardInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ContentId        *string `json:"contentId"`
	Note             *string `json:"note"`
	ProjectColumnId  string  `json:"projectColumnId"`
}

type AddProjectColumnInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	ProjectId        string  `json:"projectId"`
}

type AddPullRequestReviewCommentInput struct {
	Body                string  `json:"body"`
	ClientMutationId    *string `json:"clientMutationId"`
	CommitOID           *string `json:"commitOID"`
	InReplyTo           *string `json:"inReplyTo"`
	Path                *string `json:"path"`
	Position            *int    `json:"position"`
	PullRequestId       *string `json:"pullRequestId"`
	PullRequestReviewId *string `json:"pullRequestReviewId"`
}

type AddPullRequestReviewInput struct {
	Body             *string                          `json:"body"`
	ClientMutationId *string                          `json:"clientMutationId"`
	Comments         []*DraftPullRequestReviewComment `json:"comments"`
	CommitOID        *string                          `json:"commitOID"`
	Event            *PullRequestReviewEvent          `json:"event"`
	PullRequestId    string                           `json:"pullRequestId"`
	Threads          []*DraftPullRequestReviewThread  `json:"threads"`
}

type AddPullRequestReviewThreadInput struct {
	Body                string    `json:"body"`
	ClientMutationId    *string   `json:"clientMutationId"`
	Line                int       `json:"line"`
	Path                string    `json:"path"`
	PullRequestId       *string   `json:"pullRequestId"`
	PullRequestReviewId *string   `json:"pullRequestReviewId"`
	Side                *DiffSide `json:"side"`
	StartLine           *int      `json:"startLine"`
	StartSide           *DiffSide `json:"startSide"`
}

type AddReactionInput struct {
	ClientMutationId *string         `json:"clientMutationId"`
	Content          ReactionContent `json:"content"`
	SubjectId        string          `json:"subjectId"`
}

type AddStarInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	StarrableId      string  `json:"starrableId"`
}

type AddUpvoteInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	SubjectId        string  `json:"subjectId"`
}

type AddVerifiableDomainInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Domain           string  `json:"domain"`
	OwnerId          string  `json:"ownerId"`
}

type ApproveDeploymentsInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	Comment          *string  `json:"comment"`
	EnvironmentIds   []string `json:"environmentIds"`
	WorkflowRunId    string   `json:"workflowRunId"`
}

type ApproveVerifiableDomainInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type ArchiveRepositoryInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	RepositoryId     string  `json:"repositoryId"`
}

type AuditLogOrder struct {
	Direction *OrderDirection     `json:"direction"`
	Field     *AuditLogOrderField `json:"field"`
}

type CancelEnterpriseAdminInvitationInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	InvitationId     string  `json:"invitationId"`
}

type ChangeUserStatusInput struct {
	ClientMutationId    *string    `json:"clientMutationId"`
	Emoji               *string    `json:"emoji"`
	ExpiresAt           *time.Time `json:"expiresAt"`
	LimitedAvailability *bool      `json:"limitedAvailability"`
	Message             *string    `json:"message"`
	OrganizationId      *string    `json:"organizationId"`
}

type CheckAnnotationData struct {
	AnnotationLevel CheckAnnotationLevel `json:"annotationLevel"`
	Location        CheckAnnotationRange `json:"location"`
	Message         string               `json:"message"`
	Path            string               `json:"path"`
	RawDetails      *string              `json:"rawDetails"`
	Title           *string              `json:"title"`
}

type CheckAnnotationRange struct {
	EndColumn   *int `json:"endColumn"`
	EndLine     int  `json:"endLine"`
	StartColumn *int `json:"startColumn"`
	StartLine   int  `json:"startLine"`
}

type CheckRunAction struct {
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
	Label       string `json:"label"`
}

type CheckRunFilter struct {
	AppId     *int              `json:"appId"`
	CheckName *string           `json:"checkName"`
	CheckType *CheckRunType     `json:"checkType"`
	Status    *CheckStatusState `json:"status"`
}

type CheckRunOutput struct {
	Annotations []CheckAnnotationData `json:"annotations"`
	Images      []CheckRunOutputImage `json:"images"`
	Summary     string                `json:"summary"`
	Text        *string               `json:"text"`
	Title       string                `json:"title"`
}

type CheckRunOutputImage struct {
	Alt      string  `json:"alt"`
	Caption  *string `json:"caption"`
	ImageUrl string  `json:"imageUrl"`
}

type CheckSuiteAutoTriggerPreference struct {
	AppId   string `json:"appId"`
	Setting bool   `json:"setting"`
}

type CheckSuiteFilter struct {
	AppId     *int    `json:"appId"`
	CheckName *string `json:"checkName"`
}

type ClearLabelsFromLabelableInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	LabelableId      string  `json:"labelableId"`
}

type CloneProjectInput struct {
	Body             *string `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	IncludeWorkflows bool    `json:"includeWorkflows"`
	Name             string  `json:"name"`
	Public           *bool   `json:"public"`
	SourceId         string  `json:"sourceId"`
	TargetOwnerId    string  `json:"targetOwnerId"`
}

type CloneTemplateRepositoryInput struct {
	ClientMutationId   *string              `json:"clientMutationId"`
	Description        *string              `json:"description"`
	IncludeAllBranches *bool                `json:"includeAllBranches"`
	Name               string               `json:"name"`
	OwnerId            string               `json:"ownerId"`
	RepositoryId       string               `json:"repositoryId"`
	Visibility         RepositoryVisibility `json:"visibility"`
}

type CloseIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
}

type ClosePullRequestInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PullRequestId    string  `json:"pullRequestId"`
}

type CommitAuthor struct {
	Emails []string `json:"emails"`
	Id     *string  `json:"id"`
}

type CommitContributionOrder struct {
	Direction OrderDirection               `json:"direction"`
	Field     CommitContributionOrderField `json:"field"`
}

type ContributionOrder struct {
	Direction OrderDirection `json:"direction"`
}

type ConvertProjectCardNoteToIssueInput struct {
	Body             *string `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	ProjectCardId    string  `json:"projectCardId"`
	RepositoryId     string  `json:"repositoryId"`
	Title            *string `json:"title"`
}

type ConvertPullRequestToDraftInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PullRequestId    string  `json:"pullRequestId"`
}

type CreateBranchProtectionRuleInput struct {
	AllowsDeletions                *bool    `json:"allowsDeletions"`
	AllowsForcePushes              *bool    `json:"allowsForcePushes"`
	ClientMutationId               *string  `json:"clientMutationId"`
	DismissesStaleReviews          *bool    `json:"dismissesStaleReviews"`
	IsAdminEnforced                *bool    `json:"isAdminEnforced"`
	Pattern                        string   `json:"pattern"`
	PushActorIds                   []string `json:"pushActorIds"`
	RepositoryId                   string   `json:"repositoryId"`
	RequiredApprovingReviewCount   *int     `json:"requiredApprovingReviewCount"`
	RequiredStatusCheckContexts    []string `json:"requiredStatusCheckContexts"`
	RequiresApprovingReviews       *bool    `json:"requiresApprovingReviews"`
	RequiresCodeOwnerReviews       *bool    `json:"requiresCodeOwnerReviews"`
	RequiresCommitSignatures       *bool    `json:"requiresCommitSignatures"`
	RequiresConversationResolution *bool    `json:"requiresConversationResolution"`
	RequiresLinearHistory          *bool    `json:"requiresLinearHistory"`
	RequiresStatusChecks           *bool    `json:"requiresStatusChecks"`
	RequiresStrictStatusChecks     *bool    `json:"requiresStrictStatusChecks"`
	RestrictsPushes                *bool    `json:"restrictsPushes"`
	RestrictsReviewDismissals      *bool    `json:"restrictsReviewDismissals"`
	ReviewDismissalActorIds        []string `json:"reviewDismissalActorIds"`
}

type CreateCheckRunInput struct {
	Actions          []CheckRunAction             `json:"actions"`
	ClientMutationId *string                      `json:"clientMutationId"`
	CompletedAt      *time.Time                   `json:"completedAt"`
	Conclusion       *CheckConclusionState        `json:"conclusion"`
	DetailsUrl       *string                      `json:"detailsUrl"`
	ExternalId       *string                      `json:"externalId"`
	HeadSha          string                       `json:"headSha"`
	Name             string                       `json:"name"`
	Output           *CheckRunOutput              `json:"output"`
	RepositoryId     string                       `json:"repositoryId"`
	StartedAt        *time.Time                   `json:"startedAt"`
	Status           *RequestableCheckStatusState `json:"status"`
}

type CreateCheckSuiteInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	HeadSha          string  `json:"headSha"`
	RepositoryId     string  `json:"repositoryId"`
}

type CreateContentAttachmentInput struct {
	Body               string  `json:"body"`
	ClientMutationId   *string `json:"clientMutationId"`
	ContentReferenceId string  `json:"contentReferenceId"`
	Title              string  `json:"title"`
}

type CreateDeploymentInput struct {
	AutoMerge        *bool    `json:"autoMerge"`
	ClientMutationId *string  `json:"clientMutationId"`
	Description      *string  `json:"description"`
	Environment      *string  `json:"environment"`
	Payload          *string  `json:"payload"`
	RefId            string   `json:"refId"`
	RepositoryId     string   `json:"repositoryId"`
	RequiredContexts []string `json:"requiredContexts"`
	Task             *string  `json:"task"`
}

type CreateDeploymentStatusInput struct {
	AutoInactive     *bool                 `json:"autoInactive"`
	ClientMutationId *string               `json:"clientMutationId"`
	DeploymentId     string                `json:"deploymentId"`
	Description      *string               `json:"description"`
	Environment      *string               `json:"environment"`
	EnvironmentUrl   *string               `json:"environmentUrl"`
	LogUrl           *string               `json:"logUrl"`
	State            DeploymentStatusState `json:"state"`
}

type CreateDiscussionInput struct {
	Body             string  `json:"body"`
	CategoryId       string  `json:"categoryId"`
	ClientMutationId *string `json:"clientMutationId"`
	RepositoryId     string  `json:"repositoryId"`
	Title            string  `json:"title"`
}

type CreateEnterpriseOrganizationInput struct {
	AdminLogins      []string `json:"adminLogins"`
	BillingEmail     string   `json:"billingEmail"`
	ClientMutationId *string  `json:"clientMutationId"`
	EnterpriseId     string   `json:"enterpriseId"`
	Login            string   `json:"login"`
	ProfileName      string   `json:"profileName"`
}

type CreateEnvironmentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	RepositoryId     string  `json:"repositoryId"`
}

type CreateIpAllowListEntryInput struct {
	AllowListValue   string  `json:"allowListValue"`
	ClientMutationId *string `json:"clientMutationId"`
	IsActive         bool    `json:"isActive"`
	Name             *string `json:"name"`
	OwnerId          string  `json:"ownerId"`
}

type CreateIssueInput struct {
	AssigneeIds      []string `json:"assigneeIds"`
	Body             *string  `json:"body"`
	ClientMutationId *string  `json:"clientMutationId"`
	IssueTemplate    *string  `json:"issueTemplate"`
	LabelIds         []string `json:"labelIds"`
	MilestoneId      *string  `json:"milestoneId"`
	ProjectIds       []string `json:"projectIds"`
	RepositoryId     string   `json:"repositoryId"`
	Title            string   `json:"title"`
}

type CreateLabelInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Color            string  `json:"color"`
	Description      *string `json:"description"`
	Name             string  `json:"name"`
	RepositoryId     string  `json:"repositoryId"`
}

type CreateProjectInput struct {
	Body             *string          `json:"body"`
	ClientMutationId *string          `json:"clientMutationId"`
	Name             string           `json:"name"`
	OwnerId          string           `json:"ownerId"`
	RepositoryIds    []string         `json:"repositoryIds"`
	Template         *ProjectTemplate `json:"template"`
}

type CreatePullRequestInput struct {
	BaseRefName         string  `json:"baseRefName"`
	Body                *string `json:"body"`
	ClientMutationId    *string `json:"clientMutationId"`
	Draft               *bool   `json:"draft"`
	HeadRefName         string  `json:"headRefName"`
	MaintainerCanModify *bool   `json:"maintainerCanModify"`
	RepositoryId        string  `json:"repositoryId"`
	Title               string  `json:"title"`
}

type CreateRefInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	Oid              string  `json:"oid"`
	RepositoryId     string  `json:"repositoryId"`
}

type CreateRepositoryInput struct {
	ClientMutationId *string              `json:"clientMutationId"`
	Description      *string              `json:"description"`
	HasIssuesEnabled *bool                `json:"hasIssuesEnabled"`
	HasWikiEnabled   *bool                `json:"hasWikiEnabled"`
	HomepageUrl      *string              `json:"homepageUrl"`
	Name             string               `json:"name"`
	OwnerId          *string              `json:"ownerId"`
	TeamId           *string              `json:"teamId"`
	Template         *bool                `json:"template"`
	Visibility       RepositoryVisibility `json:"visibility"`
}

type CreateTeamDiscussionCommentInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	DiscussionId     string  `json:"discussionId"`
}

type CreateTeamDiscussionInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	Private          *bool   `json:"private"`
	TeamId           string  `json:"teamId"`
	Title            string  `json:"title"`
}

type DeclineTopicSuggestionInput struct {
	ClientMutationId *string                      `json:"clientMutationId"`
	Name             string                       `json:"name"`
	Reason           TopicSuggestionDeclineReason `json:"reason"`
	RepositoryId     string                       `json:"repositoryId"`
}

type DeleteBranchProtectionRuleInput struct {
	BranchProtectionRuleId string  `json:"branchProtectionRuleId"`
	ClientMutationId       *string `json:"clientMutationId"`
}

type DeleteDeploymentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteDiscussionCommentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteDiscussionInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteEnvironmentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteIpAllowListEntryInput struct {
	ClientMutationId   *string `json:"clientMutationId"`
	IpAllowListEntryId string  `json:"ipAllowListEntryId"`
}

type DeleteIssueCommentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
}

type DeleteLabelInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeletePackageVersionInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PackageVersionId string  `json:"packageVersionId"`
}

type DeleteProjectCardInput struct {
	CardId           string  `json:"cardId"`
	ClientMutationId *string `json:"clientMutationId"`
}

type DeleteProjectColumnInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ColumnId         string  `json:"columnId"`
}

type DeleteProjectInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ProjectId        string  `json:"projectId"`
}

type DeletePullRequestReviewCommentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeletePullRequestReviewInput struct {
	ClientMutationId    *string `json:"clientMutationId"`
	PullRequestReviewId string  `json:"pullRequestReviewId"`
}

type DeleteRefInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	RefId            string  `json:"refId"`
}

type DeleteTeamDiscussionCommentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteTeamDiscussionInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeleteVerifiableDomainInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type DeploymentOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     DeploymentOrderField `json:"field"`
}

type DisablePullRequestAutoMergeInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PullRequestId    string  `json:"pullRequestId"`
}

type DiscussionOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     DiscussionOrderField `json:"field"`
}

type DismissPullRequestReviewInput struct {
	ClientMutationId    *string `json:"clientMutationId"`
	Message             string  `json:"message"`
	PullRequestReviewId string  `json:"pullRequestReviewId"`
}

type DraftPullRequestReviewComment struct {
	Body     string `json:"body"`
	Path     string `json:"path"`
	Position int    `json:"position"`
}

type DraftPullRequestReviewThread struct {
	Body      string    `json:"body"`
	Line      int       `json:"line"`
	Path      string    `json:"path"`
	Side      *DiffSide `json:"side"`
	StartLine *int      `json:"startLine"`
	StartSide *DiffSide `json:"startSide"`
}

type EnablePullRequestAutoMergeInput struct {
	AuthorEmail      *string                 `json:"authorEmail"`
	ClientMutationId *string                 `json:"clientMutationId"`
	CommitBody       *string                 `json:"commitBody"`
	CommitHeadline   *string                 `json:"commitHeadline"`
	MergeMethod      *PullRequestMergeMethod `json:"mergeMethod"`
	PullRequestId    string                  `json:"pullRequestId"`
}

type EnterpriseAdministratorInvitationOrder struct {
	Direction OrderDirection                              `json:"direction"`
	Field     EnterpriseAdministratorInvitationOrderField `json:"field"`
}

type EnterpriseMemberOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     EnterpriseMemberOrderField `json:"field"`
}

type EnterpriseServerInstallationOrder struct {
	Direction OrderDirection                         `json:"direction"`
	Field     EnterpriseServerInstallationOrderField `json:"field"`
}

type EnterpriseServerUserAccountEmailOrder struct {
	Direction OrderDirection                             `json:"direction"`
	Field     EnterpriseServerUserAccountEmailOrderField `json:"field"`
}

type EnterpriseServerUserAccountOrder struct {
	Direction OrderDirection                        `json:"direction"`
	Field     EnterpriseServerUserAccountOrderField `json:"field"`
}

type EnterpriseServerUserAccountsUploadOrder struct {
	Direction OrderDirection                               `json:"direction"`
	Field     EnterpriseServerUserAccountsUploadOrderField `json:"field"`
}

type FollowUserInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	UserId           string  `json:"userId"`
}

type GistOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     GistOrderField `json:"field"`
}

type ImportProjectInput struct {
	Body             *string               `json:"body"`
	ClientMutationId *string               `json:"clientMutationId"`
	ColumnImports    []ProjectColumnImport `json:"columnImports"`
	Name             string                `json:"name"`
	OwnerName        string                `json:"ownerName"`
	Public           *bool                 `json:"public"`
}

type InviteEnterpriseAdminInput struct {
	ClientMutationId *string                      `json:"clientMutationId"`
	Email            *string                      `json:"email"`
	EnterpriseId     string                       `json:"enterpriseId"`
	Invitee          *string                      `json:"invitee"`
	Role             *EnterpriseAdministratorRole `json:"role"`
}

type IpAllowListEntryOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     IpAllowListEntryOrderField `json:"field"`
}

type IssueCommentOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     IssueCommentOrderField `json:"field"`
}

type IssueFilters struct {
	Assignee         *string      `json:"assignee"`
	CreatedBy        *string      `json:"createdBy"`
	Labels           []string     `json:"labels"`
	Mentioned        *string      `json:"mentioned"`
	Milestone        *string      `json:"milestone"`
	Since            *time.Time   `json:"since"`
	States           []IssueState `json:"states"`
	ViewerSubscribed *bool        `json:"viewerSubscribed"`
}

type IssueOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     IssueOrderField `json:"field"`
}

type LabelOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     LabelOrderField `json:"field"`
}

type LanguageOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     LanguageOrderField `json:"field"`
}

type LinkRepositoryToProjectInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ProjectId        string  `json:"projectId"`
	RepositoryId     string  `json:"repositoryId"`
}

type LockLockableInput struct {
	ClientMutationId *string     `json:"clientMutationId"`
	LockReason       *LockReason `json:"lockReason"`
	LockableId       string      `json:"lockableId"`
}

type MarkDiscussionCommentAsAnswerInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type MarkFileAsViewedInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Path             string  `json:"path"`
	PullRequestId    string  `json:"pullRequestId"`
}

type MarkPullRequestReadyForReviewInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PullRequestId    string  `json:"pullRequestId"`
}

type MergeBranchInput struct {
	AuthorEmail      *string `json:"authorEmail"`
	Base             string  `json:"base"`
	ClientMutationId *string `json:"clientMutationId"`
	CommitMessage    *string `json:"commitMessage"`
	Head             string  `json:"head"`
	RepositoryId     string  `json:"repositoryId"`
}

type MergePullRequestInput struct {
	AuthorEmail      *string                 `json:"authorEmail"`
	ClientMutationId *string                 `json:"clientMutationId"`
	CommitBody       *string                 `json:"commitBody"`
	CommitHeadline   *string                 `json:"commitHeadline"`
	ExpectedHeadOid  *string                 `json:"expectedHeadOid"`
	MergeMethod      *PullRequestMergeMethod `json:"mergeMethod"`
	PullRequestId    string                  `json:"pullRequestId"`
}

type MilestoneOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     MilestoneOrderField `json:"field"`
}

type MinimizeCommentInput struct {
	Classifier       ReportedContentClassifiers `json:"classifier"`
	ClientMutationId *string                    `json:"clientMutationId"`
	SubjectId        string                     `json:"subjectId"`
}

type MoveProjectCardInput struct {
	AfterCardId      *string `json:"afterCardId"`
	CardId           string  `json:"cardId"`
	ClientMutationId *string `json:"clientMutationId"`
	ColumnId         string  `json:"columnId"`
}

type MoveProjectColumnInput struct {
	AfterColumnId    *string `json:"afterColumnId"`
	ClientMutationId *string `json:"clientMutationId"`
	ColumnId         string  `json:"columnId"`
}

type OrganizationOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     OrganizationOrderField `json:"field"`
}

type PackageFileOrder struct {
	Direction *OrderDirection        `json:"direction"`
	Field     *PackageFileOrderField `json:"field"`
}

type PackageOrder struct {
	Direction *OrderDirection    `json:"direction"`
	Field     *PackageOrderField `json:"field"`
}

type PackageVersionOrder struct {
	Direction *OrderDirection           `json:"direction"`
	Field     *PackageVersionOrderField `json:"field"`
}

type PinIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
}

type ProjectCardImport struct {
	Number     int    `json:"number"`
	Repository string `json:"repository"`
}

type ProjectColumnImport struct {
	ColumnName string              `json:"columnName"`
	Issues     []ProjectCardImport `json:"issues"`
	Position   int                 `json:"position"`
}

type ProjectOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     ProjectOrderField `json:"field"`
}

type PullRequestOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     PullRequestOrderField `json:"field"`
}

type ReactionOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     ReactionOrderField `json:"field"`
}

type RefOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     RefOrderField  `json:"field"`
}

type RefUpdate struct {
	AfterOid  string  `json:"afterOid"`
	BeforeOid *string `json:"beforeOid"`
	Force     *bool   `json:"force"`
	Name      string  `json:"name"`
}

type RegenerateEnterpriseIdentityProviderRecoveryCodesInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
}

type RegenerateVerifiableDomainTokenInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type RejectDeploymentsInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	Comment          *string  `json:"comment"`
	EnvironmentIds   []string `json:"environmentIds"`
	WorkflowRunId    string   `json:"workflowRunId"`
}

type ReleaseOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     ReleaseOrderField `json:"field"`
}

type RemoveAssigneesFromAssignableInput struct {
	AssignableId     string   `json:"assignableId"`
	AssigneeIds      []string `json:"assigneeIds"`
	ClientMutationId *string  `json:"clientMutationId"`
}

type RemoveEnterpriseAdminInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
	Login            string  `json:"login"`
}

type RemoveEnterpriseIdentityProviderInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
}

type RemoveEnterpriseOrganizationInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
	OrganizationId   string  `json:"organizationId"`
}

type RemoveEnterpriseSupportEntitlementInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	EnterpriseId     string  `json:"enterpriseId"`
	Login            string  `json:"login"`
}

type RemoveLabelsFromLabelableInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	LabelIds         []string `json:"labelIds"`
	LabelableId      string   `json:"labelableId"`
}

type RemoveOutsideCollaboratorInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	OrganizationId   string  `json:"organizationId"`
	UserId           string  `json:"userId"`
}

type RemoveReactionInput struct {
	ClientMutationId *string         `json:"clientMutationId"`
	Content          ReactionContent `json:"content"`
	SubjectId        string          `json:"subjectId"`
}

type RemoveStarInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	StarrableId      string  `json:"starrableId"`
}

type RemoveUpvoteInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	SubjectId        string  `json:"subjectId"`
}

type ReopenIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
}

type ReopenPullRequestInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	PullRequestId    string  `json:"pullRequestId"`
}

type RepositoryInvitationOrder struct {
	Direction OrderDirection                 `json:"direction"`
	Field     RepositoryInvitationOrderField `json:"field"`
}

type RepositoryOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     RepositoryOrderField `json:"field"`
}

type RequestReviewsInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	PullRequestId    string   `json:"pullRequestId"`
	TeamIds          []string `json:"teamIds"`
	Union            *bool    `json:"union"`
	UserIds          []string `json:"userIds"`
}

type RerequestCheckSuiteInput struct {
	CheckSuiteId     string  `json:"checkSuiteId"`
	ClientMutationId *string `json:"clientMutationId"`
	RepositoryId     string  `json:"repositoryId"`
}

type ResolveReviewThreadInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ThreadId         string  `json:"threadId"`
}

type SavedReplyOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     SavedReplyOrderField `json:"field"`
}

type SecurityAdvisoryIdentifierFilter struct {
	Type  SecurityAdvisoryIdentifierType `json:"type"`
	Value string                         `json:"value"`
}

type SecurityAdvisoryOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     SecurityAdvisoryOrderField `json:"field"`
}

type SecurityVulnerabilityOrder struct {
	Direction OrderDirection                  `json:"direction"`
	Field     SecurityVulnerabilityOrderField `json:"field"`
}

type SetEnterpriseIdentityProviderInput struct {
	ClientMutationId *string                `json:"clientMutationId"`
	DigestMethod     SamlDigestAlgorithm    `json:"digestMethod"`
	EnterpriseId     string                 `json:"enterpriseId"`
	IdpCertificate   string                 `json:"idpCertificate"`
	Issuer           *string                `json:"issuer"`
	SignatureMethod  SamlSignatureAlgorithm `json:"signatureMethod"`
	SsoUrl           string                 `json:"ssoUrl"`
}

type SetOrganizationInteractionLimitInput struct {
	ClientMutationId *string                           `json:"clientMutationId"`
	Expiry           *RepositoryInteractionLimitExpiry `json:"expiry"`
	Limit            RepositoryInteractionLimit        `json:"limit"`
	OrganizationId   string                            `json:"organizationId"`
}

type SetRepositoryInteractionLimitInput struct {
	ClientMutationId *string                           `json:"clientMutationId"`
	Expiry           *RepositoryInteractionLimitExpiry `json:"expiry"`
	Limit            RepositoryInteractionLimit        `json:"limit"`
	RepositoryId     string                            `json:"repositoryId"`
}

type SetUserInteractionLimitInput struct {
	ClientMutationId *string                           `json:"clientMutationId"`
	Expiry           *RepositoryInteractionLimitExpiry `json:"expiry"`
	Limit            RepositoryInteractionLimit        `json:"limit"`
	UserId           string                            `json:"userId"`
}

type SponsorableOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     SponsorableOrderField `json:"field"`
}

type SponsorsActivityOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     SponsorsActivityOrderField `json:"field"`
}

type SponsorsTierOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     SponsorsTierOrderField `json:"field"`
}

type SponsorshipOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     SponsorshipOrderField `json:"field"`
}

type StarOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     StarOrderField `json:"field"`
}

type SubmitPullRequestReviewInput struct {
	Body                *string                `json:"body"`
	ClientMutationId    *string                `json:"clientMutationId"`
	Event               PullRequestReviewEvent `json:"event"`
	PullRequestId       *string                `json:"pullRequestId"`
	PullRequestReviewId *string                `json:"pullRequestReviewId"`
}

type TeamDiscussionCommentOrder struct {
	Direction OrderDirection                  `json:"direction"`
	Field     TeamDiscussionCommentOrderField `json:"field"`
}

type TeamDiscussionOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     TeamDiscussionOrderField `json:"field"`
}

type TeamMemberOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     TeamMemberOrderField `json:"field"`
}

type TeamOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     TeamOrderField `json:"field"`
}

type TeamRepositoryOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     TeamRepositoryOrderField `json:"field"`
}

type TransferIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
	RepositoryId     string  `json:"repositoryId"`
}

type UnarchiveRepositoryInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	RepositoryId     string  `json:"repositoryId"`
}

type UnfollowUserInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	UserId           string  `json:"userId"`
}

type UnlinkRepositoryFromProjectInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ProjectId        string  `json:"projectId"`
	RepositoryId     string  `json:"repositoryId"`
}

type UnlockLockableInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	LockableId       string  `json:"lockableId"`
}

type UnmarkDiscussionCommentAsAnswerInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type UnmarkFileAsViewedInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Path             string  `json:"path"`
	PullRequestId    string  `json:"pullRequestId"`
}

type UnmarkIssueAsDuplicateInput struct {
	CanonicalId      string  `json:"canonicalId"`
	ClientMutationId *string `json:"clientMutationId"`
	DuplicateId      string  `json:"duplicateId"`
}

type UnminimizeCommentInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	SubjectId        string  `json:"subjectId"`
}

type UnpinIssueInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IssueId          string  `json:"issueId"`
}

type UnresolveReviewThreadInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	ThreadId         string  `json:"threadId"`
}

type UpdateBranchProtectionRuleInput struct {
	AllowsDeletions                *bool    `json:"allowsDeletions"`
	AllowsForcePushes              *bool    `json:"allowsForcePushes"`
	BranchProtectionRuleId         string   `json:"branchProtectionRuleId"`
	ClientMutationId               *string  `json:"clientMutationId"`
	DismissesStaleReviews          *bool    `json:"dismissesStaleReviews"`
	IsAdminEnforced                *bool    `json:"isAdminEnforced"`
	Pattern                        *string  `json:"pattern"`
	PushActorIds                   []string `json:"pushActorIds"`
	RequiredApprovingReviewCount   *int     `json:"requiredApprovingReviewCount"`
	RequiredStatusCheckContexts    []string `json:"requiredStatusCheckContexts"`
	RequiresApprovingReviews       *bool    `json:"requiresApprovingReviews"`
	RequiresCodeOwnerReviews       *bool    `json:"requiresCodeOwnerReviews"`
	RequiresCommitSignatures       *bool    `json:"requiresCommitSignatures"`
	RequiresConversationResolution *bool    `json:"requiresConversationResolution"`
	RequiresLinearHistory          *bool    `json:"requiresLinearHistory"`
	RequiresStatusChecks           *bool    `json:"requiresStatusChecks"`
	RequiresStrictStatusChecks     *bool    `json:"requiresStrictStatusChecks"`
	RestrictsPushes                *bool    `json:"restrictsPushes"`
	RestrictsReviewDismissals      *bool    `json:"restrictsReviewDismissals"`
	ReviewDismissalActorIds        []string `json:"reviewDismissalActorIds"`
}

type UpdateCheckRunInput struct {
	Actions          []CheckRunAction             `json:"actions"`
	CheckRunId       string                       `json:"checkRunId"`
	ClientMutationId *string                      `json:"clientMutationId"`
	CompletedAt      *time.Time                   `json:"completedAt"`
	Conclusion       *CheckConclusionState        `json:"conclusion"`
	DetailsUrl       *string                      `json:"detailsUrl"`
	ExternalId       *string                      `json:"externalId"`
	Name             *string                      `json:"name"`
	Output           *CheckRunOutput              `json:"output"`
	RepositoryId     string                       `json:"repositoryId"`
	StartedAt        *time.Time                   `json:"startedAt"`
	Status           *RequestableCheckStatusState `json:"status"`
}

type UpdateCheckSuitePreferencesInput struct {
	AutoTriggerPreferences []CheckSuiteAutoTriggerPreference `json:"autoTriggerPreferences"`
	ClientMutationId       *string                           `json:"clientMutationId"`
	RepositoryId           string                            `json:"repositoryId"`
}

type UpdateDiscussionCommentInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	CommentId        string  `json:"commentId"`
}

type UpdateDiscussionInput struct {
	Body             *string `json:"body"`
	CategoryId       *string `json:"categoryId"`
	ClientMutationId *string `json:"clientMutationId"`
	DiscussionId     string  `json:"discussionId"`
	Title            *string `json:"title"`
}

type UpdateEnterpriseAdministratorRoleInput struct {
	ClientMutationId *string                     `json:"clientMutationId"`
	EnterpriseId     string                      `json:"enterpriseId"`
	Login            string                      `json:"login"`
	Role             EnterpriseAdministratorRole `json:"role"`
}

type UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseDefaultRepositoryPermissionSettingInput struct {
	ClientMutationId *string                                           `json:"clientMutationId"`
	EnterpriseId     string                                            `json:"enterpriseId"`
	SettingValue     EnterpriseDefaultRepositoryPermissionSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanCreateRepositoriesSettingInput struct {
	ClientMutationId                          *string                                             `json:"clientMutationId"`
	EnterpriseId                              string                                              `json:"enterpriseId"`
	MembersCanCreateInternalRepositories      *bool                                               `json:"membersCanCreateInternalRepositories"`
	MembersCanCreatePrivateRepositories       *bool                                               `json:"membersCanCreatePrivateRepositories"`
	MembersCanCreatePublicRepositories        *bool                                               `json:"membersCanCreatePublicRepositories"`
	MembersCanCreateRepositoriesPolicyEnabled *bool                                               `json:"membersCanCreateRepositoriesPolicyEnabled"`
	SettingValue                              *EnterpriseMembersCanCreateRepositoriesSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanDeleteIssuesSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanDeleteRepositoriesSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanInviteCollaboratorsSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanMakePurchasesSettingInput struct {
	ClientMutationId *string                                       `json:"clientMutationId"`
	EnterpriseId     string                                        `json:"enterpriseId"`
	SettingValue     EnterpriseMembersCanMakePurchasesSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseMembersCanViewDependencyInsightsSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseOrganizationProjectsSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseProfileInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Description      *string `json:"description"`
	EnterpriseId     string  `json:"enterpriseId"`
	Location         *string `json:"location"`
	Name             *string `json:"name"`
	WebsiteUrl       *string `json:"websiteUrl"`
}

type UpdateEnterpriseRepositoryProjectsSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseTeamDiscussionsSettingInput struct {
	ClientMutationId *string                               `json:"clientMutationId"`
	EnterpriseId     string                                `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
}

type UpdateEnterpriseTwoFactorAuthenticationRequiredSettingInput struct {
	ClientMutationId *string                       `json:"clientMutationId"`
	EnterpriseId     string                        `json:"enterpriseId"`
	SettingValue     EnterpriseEnabledSettingValue `json:"settingValue"`
}

type UpdateEnvironmentInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	EnvironmentId    string   `json:"environmentId"`
	Reviewers        []string `json:"reviewers"`
	WaitTimer        *int     `json:"waitTimer"`
}

type UpdateIpAllowListEnabledSettingInput struct {
	ClientMutationId *string                        `json:"clientMutationId"`
	OwnerId          string                         `json:"ownerId"`
	SettingValue     IpAllowListEnabledSettingValue `json:"settingValue"`
}

type UpdateIpAllowListEntryInput struct {
	AllowListValue     string  `json:"allowListValue"`
	ClientMutationId   *string `json:"clientMutationId"`
	IpAllowListEntryId string  `json:"ipAllowListEntryId"`
	IsActive           bool    `json:"isActive"`
	Name               *string `json:"name"`
}

type UpdateIpAllowListForInstalledAppsEnabledSettingInput struct {
	ClientMutationId *string                                        `json:"clientMutationId"`
	OwnerId          string                                         `json:"ownerId"`
	SettingValue     IpAllowListForInstalledAppsEnabledSettingValue `json:"settingValue"`
}

type UpdateIssueCommentInput struct {
	Body             string  `json:"body"`
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type UpdateIssueInput struct {
	AssigneeIds      []string    `json:"assigneeIds"`
	Body             *string     `json:"body"`
	ClientMutationId *string     `json:"clientMutationId"`
	Id               string      `json:"id"`
	LabelIds         []string    `json:"labelIds"`
	MilestoneId      *string     `json:"milestoneId"`
	ProjectIds       []string    `json:"projectIds"`
	State            *IssueState `json:"state"`
	Title            *string     `json:"title"`
}

type UpdateLabelInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Color            *string `json:"color"`
	Description      *string `json:"description"`
	Id               string  `json:"id"`
	Name             *string `json:"name"`
}

type UpdateNotificationRestrictionSettingInput struct {
	ClientMutationId *string                             `json:"clientMutationId"`
	OwnerId          string                              `json:"ownerId"`
	SettingValue     NotificationRestrictionSettingValue `json:"settingValue"`
}

type UpdateProjectCardInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	IsArchived       *bool   `json:"isArchived"`
	Note             *string `json:"note"`
	ProjectCardId    string  `json:"projectCardId"`
}

type UpdateProjectColumnInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	ProjectColumnId  string  `json:"projectColumnId"`
}

type UpdateProjectInput struct {
	Body             *string       `json:"body"`
	ClientMutationId *string       `json:"clientMutationId"`
	Name             *string       `json:"name"`
	ProjectId        string        `json:"projectId"`
	Public           *bool         `json:"public"`
	State            *ProjectState `json:"state"`
}

type UpdatePullRequestInput struct {
	AssigneeIds         []string                `json:"assigneeIds"`
	BaseRefName         *string                 `json:"baseRefName"`
	Body                *string                 `json:"body"`
	ClientMutationId    *string                 `json:"clientMutationId"`
	LabelIds            []string                `json:"labelIds"`
	MaintainerCanModify *bool                   `json:"maintainerCanModify"`
	MilestoneId         *string                 `json:"milestoneId"`
	ProjectIds          []string                `json:"projectIds"`
	PullRequestId       string                  `json:"pullRequestId"`
	State               *PullRequestUpdateState `json:"state"`
	Title               *string                 `json:"title"`
}

type UpdatePullRequestReviewCommentInput struct {
	Body                       string  `json:"body"`
	ClientMutationId           *string `json:"clientMutationId"`
	PullRequestReviewCommentId string  `json:"pullRequestReviewCommentId"`
}

type UpdatePullRequestReviewInput struct {
	Body                string  `json:"body"`
	ClientMutationId    *string `json:"clientMutationId"`
	PullRequestReviewId string  `json:"pullRequestReviewId"`
}

type UpdateRefInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Force            *bool   `json:"force"`
	Oid              string  `json:"oid"`
	RefId            string  `json:"refId"`
}

type UpdateRefsInput struct {
	ClientMutationId *string     `json:"clientMutationId"`
	RefUpdates       []RefUpdate `json:"refUpdates"`
	RepositoryId     string      `json:"repositoryId"`
}

type UpdateRepositoryInput struct {
	ClientMutationId   *string `json:"clientMutationId"`
	Description        *string `json:"description"`
	HasIssuesEnabled   *bool   `json:"hasIssuesEnabled"`
	HasProjectsEnabled *bool   `json:"hasProjectsEnabled"`
	HasWikiEnabled     *bool   `json:"hasWikiEnabled"`
	HomepageUrl        *string `json:"homepageUrl"`
	Name               *string `json:"name"`
	RepositoryId       string  `json:"repositoryId"`
	Template           *bool   `json:"template"`
}

type UpdateSubscriptionInput struct {
	ClientMutationId *string           `json:"clientMutationId"`
	State            SubscriptionState `json:"state"`
	SubscribableId   string            `json:"subscribableId"`
}

type UpdateTeamDiscussionCommentInput struct {
	Body             string  `json:"body"`
	BodyVersion      *string `json:"bodyVersion"`
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}

type UpdateTeamDiscussionInput struct {
	Body             *string `json:"body"`
	BodyVersion      *string `json:"bodyVersion"`
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
	Pinned           *bool   `json:"pinned"`
	Title            *string `json:"title"`
}

type UpdateTeamReviewAssignmentInput struct {
	Algorithm             *TeamReviewAssignmentAlgorithm `json:"algorithm"`
	ClientMutationId      *string                        `json:"clientMutationId"`
	Enabled               bool                           `json:"enabled"`
	ExcludedTeamMemberIds []string                       `json:"excludedTeamMemberIds"`
	Id                    string                         `json:"id"`
	NotifyTeam            *bool                          `json:"notifyTeam"`
	TeamMemberCount       *int                           `json:"teamMemberCount"`
}

type UpdateTopicsInput struct {
	ClientMutationId *string  `json:"clientMutationId"`
	RepositoryId     string   `json:"repositoryId"`
	TopicNames       []string `json:"topicNames"`
}

type UserStatusOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     UserStatusOrderField `json:"field"`
}

type VerifiableDomainOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     VerifiableDomainOrderField `json:"field"`
}

type VerifyVerifiableDomainInput struct {
	ClientMutationId *string `json:"clientMutationId"`
	Id               string  `json:"id"`
}
