package github

import "time"

type AuditLogOrderField int

const (
	AuditLogOrderField_CREATED_AT AuditLogOrderField = iota
)

type CheckAnnotationLevel int

const (
	CheckAnnotationLevel_FAILURE CheckAnnotationLevel = iota
	CheckAnnotationLevel_NOTICE
	CheckAnnotationLevel_WARNING
)

type CheckConclusionState int

const (
	CheckConclusionState_ACTION_REQUIRED CheckConclusionState = iota
	CheckConclusionState_CANCELLED
	CheckConclusionState_FAILURE
	CheckConclusionState_NEUTRAL
	CheckConclusionState_SKIPPED
	CheckConclusionState_STALE
	CheckConclusionState_STARTUP_FAILURE
	CheckConclusionState_SUCCESS
	CheckConclusionState_TIMED_OUT
)

type CheckRunType int

const (
	CheckRunType_ALL CheckRunType = iota
	CheckRunType_LATEST
)

type CheckStatusState int

const (
	CheckStatusState_COMPLETED CheckStatusState = iota
	CheckStatusState_IN_PROGRESS
	CheckStatusState_PENDING
	CheckStatusState_QUEUED
	CheckStatusState_REQUESTED
	CheckStatusState_WAITING
)

type CollaboratorAffiliation int

const (
	CollaboratorAffiliation_ALL CollaboratorAffiliation = iota
	CollaboratorAffiliation_DIRECT
	CollaboratorAffiliation_OUTSIDE
)

type CommentAuthorAssociation int

const (
	CommentAuthorAssociation_COLLABORATOR CommentAuthorAssociation = iota
	CommentAuthorAssociation_CONTRIBUTOR
	CommentAuthorAssociation_FIRST_TIMER
	CommentAuthorAssociation_FIRST_TIME_CONTRIBUTOR
	CommentAuthorAssociation_MANNEQUIN
	CommentAuthorAssociation_MEMBER
	CommentAuthorAssociation_NONE
	CommentAuthorAssociation_OWNER
)

type CommentCannotUpdateReason int

const (
	CommentCannotUpdateReason_ARCHIVED CommentCannotUpdateReason = iota
	CommentCannotUpdateReason_DENIED
	CommentCannotUpdateReason_INSUFFICIENT_ACCESS
	CommentCannotUpdateReason_LOCKED
	CommentCannotUpdateReason_LOGIN_REQUIRED
	CommentCannotUpdateReason_MAINTENANCE
	CommentCannotUpdateReason_VERIFIED_EMAIL_REQUIRED
)

type CommitContributionOrderField int

const (
	CommitContributionOrderField_COMMIT_COUNT CommitContributionOrderField = iota
	CommitContributionOrderField_OCCURRED_AT
)

type ContributionLevel int

const (
	ContributionLevel_FIRST_QUARTILE ContributionLevel = iota
	ContributionLevel_FOURTH_QUARTILE
	ContributionLevel_NONE
	ContributionLevel_SECOND_QUARTILE
	ContributionLevel_THIRD_QUARTILE
)

type DefaultRepositoryPermissionField int

const (
	DefaultRepositoryPermissionField_ADMIN DefaultRepositoryPermissionField = iota
	DefaultRepositoryPermissionField_NONE
	DefaultRepositoryPermissionField_READ
	DefaultRepositoryPermissionField_WRITE
)

type DeploymentOrderField int

const (
	DeploymentOrderField_CREATED_AT DeploymentOrderField = iota
)

type DeploymentProtectionRuleType int

const (
	DeploymentProtectionRuleType_REQUIRED_REVIEWERS DeploymentProtectionRuleType = iota
	DeploymentProtectionRuleType_WAIT_TIMER
)

type DeploymentReviewState int

const (
	DeploymentReviewState_APPROVED DeploymentReviewState = iota
	DeploymentReviewState_REJECTED
)

type DeploymentState int

const (
	DeploymentState_ABANDONED DeploymentState = iota
	DeploymentState_ACTIVE
	DeploymentState_DESTROYED
	DeploymentState_ERROR
	DeploymentState_FAILURE
	DeploymentState_INACTIVE
	DeploymentState_IN_PROGRESS
	DeploymentState_PENDING
	DeploymentState_QUEUED
	DeploymentState_WAITING
)

type DeploymentStatusState int

const (
	DeploymentStatusState_ERROR DeploymentStatusState = iota
	DeploymentStatusState_FAILURE
	DeploymentStatusState_INACTIVE
	DeploymentStatusState_IN_PROGRESS
	DeploymentStatusState_PENDING
	DeploymentStatusState_QUEUED
	DeploymentStatusState_SUCCESS
	DeploymentStatusState_WAITING
)

type DiffSide int

const (
	DiffSide_LEFT DiffSide = iota
	DiffSide_RIGHT
)

type DiscussionOrderField int

const (
	DiscussionOrderField_CREATED_AT DiscussionOrderField = iota
	DiscussionOrderField_UPDATED_AT
)

type EnterpriseAdministratorInvitationOrderField int

const (
	EnterpriseAdministratorInvitationOrderField_CREATED_AT EnterpriseAdministratorInvitationOrderField = iota
)

type EnterpriseAdministratorRole int

const (
	EnterpriseAdministratorRole_BILLING_MANAGER EnterpriseAdministratorRole = iota
	EnterpriseAdministratorRole_OWNER
)

type EnterpriseDefaultRepositoryPermissionSettingValue int

const (
	EnterpriseDefaultRepositoryPermissionSettingValue_ADMIN EnterpriseDefaultRepositoryPermissionSettingValue = iota
	EnterpriseDefaultRepositoryPermissionSettingValue_NONE
	EnterpriseDefaultRepositoryPermissionSettingValue_NO_POLICY
	EnterpriseDefaultRepositoryPermissionSettingValue_READ
	EnterpriseDefaultRepositoryPermissionSettingValue_WRITE
)

type EnterpriseEnabledDisabledSettingValue int

const (
	EnterpriseEnabledDisabledSettingValue_DISABLED EnterpriseEnabledDisabledSettingValue = iota
	EnterpriseEnabledDisabledSettingValue_ENABLED
	EnterpriseEnabledDisabledSettingValue_NO_POLICY
)

type EnterpriseEnabledSettingValue int

const (
	EnterpriseEnabledSettingValue_ENABLED EnterpriseEnabledSettingValue = iota
	EnterpriseEnabledSettingValue_NO_POLICY
)

type EnterpriseMemberOrderField int

const (
	EnterpriseMemberOrderField_CREATED_AT EnterpriseMemberOrderField = iota
	EnterpriseMemberOrderField_LOGIN
)

type EnterpriseMembersCanCreateRepositoriesSettingValue int

const (
	EnterpriseMembersCanCreateRepositoriesSettingValue_ALL EnterpriseMembersCanCreateRepositoriesSettingValue = iota
	EnterpriseMembersCanCreateRepositoriesSettingValue_DISABLED
	EnterpriseMembersCanCreateRepositoriesSettingValue_NO_POLICY
	EnterpriseMembersCanCreateRepositoriesSettingValue_PRIVATE
	EnterpriseMembersCanCreateRepositoriesSettingValue_PUBLIC
)

type EnterpriseMembersCanMakePurchasesSettingValue int

const (
	EnterpriseMembersCanMakePurchasesSettingValue_DISABLED EnterpriseMembersCanMakePurchasesSettingValue = iota
	EnterpriseMembersCanMakePurchasesSettingValue_ENABLED
)

type EnterpriseServerInstallationOrderField int

const (
	EnterpriseServerInstallationOrderField_CREATED_AT EnterpriseServerInstallationOrderField = iota
	EnterpriseServerInstallationOrderField_CUSTOMER_NAME
	EnterpriseServerInstallationOrderField_HOST_NAME
)

type EnterpriseServerUserAccountEmailOrderField int

const (
	EnterpriseServerUserAccountEmailOrderField_EMAIL EnterpriseServerUserAccountEmailOrderField = iota
)

type EnterpriseServerUserAccountOrderField int

const (
	EnterpriseServerUserAccountOrderField_LOGIN EnterpriseServerUserAccountOrderField = iota
	EnterpriseServerUserAccountOrderField_REMOTE_CREATED_AT
)

type EnterpriseServerUserAccountsUploadOrderField int

const (
	EnterpriseServerUserAccountsUploadOrderField_CREATED_AT EnterpriseServerUserAccountsUploadOrderField = iota
)

type EnterpriseServerUserAccountsUploadSyncState int

const (
	EnterpriseServerUserAccountsUploadSyncState_FAILURE EnterpriseServerUserAccountsUploadSyncState = iota
	EnterpriseServerUserAccountsUploadSyncState_PENDING
	EnterpriseServerUserAccountsUploadSyncState_SUCCESS
)

type EnterpriseUserAccountMembershipRole int

const (
	EnterpriseUserAccountMembershipRole_MEMBER EnterpriseUserAccountMembershipRole = iota
	EnterpriseUserAccountMembershipRole_OWNER
)

type EnterpriseUserDeployment int

const (
	EnterpriseUserDeployment_CLOUD EnterpriseUserDeployment = iota
	EnterpriseUserDeployment_SERVER
)

type FileViewedState int

const (
	FileViewedState_DISMISSED FileViewedState = iota
	FileViewedState_UNVIEWED
	FileViewedState_VIEWED
)

type FundingPlatform int

const (
	FundingPlatform_COMMUNITY_BRIDGE FundingPlatform = iota
	FundingPlatform_CUSTOM
	FundingPlatform_GITHUB
	FundingPlatform_ISSUEHUNT
	FundingPlatform_KO_FI
	FundingPlatform_LIBERAPAY
	FundingPlatform_OPEN_COLLECTIVE
	FundingPlatform_OTECHIE
	FundingPlatform_PATREON
	FundingPlatform_TIDELIFT
)

type GistOrderField int

const (
	GistOrderField_CREATED_AT GistOrderField = iota
	GistOrderField_PUSHED_AT
	GistOrderField_UPDATED_AT
)

type GistPrivacy int

const (
	GistPrivacy_ALL GistPrivacy = iota
	GistPrivacy_PUBLIC
	GistPrivacy_SECRET
)

type GitSignatureState int

const (
	GitSignatureState_BAD_CERT GitSignatureState = iota
	GitSignatureState_BAD_EMAIL
	GitSignatureState_EXPIRED_KEY
	GitSignatureState_GPGVERIFY_ERROR
	GitSignatureState_GPGVERIFY_UNAVAILABLE
	GitSignatureState_INVALID
	GitSignatureState_MALFORMED_SIG
	GitSignatureState_NOT_SIGNING_KEY
	GitSignatureState_NO_USER
	GitSignatureState_OCSP_ERROR
	GitSignatureState_OCSP_PENDING
	GitSignatureState_OCSP_REVOKED
	GitSignatureState_UNKNOWN_KEY
	GitSignatureState_UNKNOWN_SIG_TYPE
	GitSignatureState_UNSIGNED
	GitSignatureState_UNVERIFIED_EMAIL
	GitSignatureState_VALID
)

type IdentityProviderConfigurationState int

const (
	IdentityProviderConfigurationState_CONFIGURED IdentityProviderConfigurationState = iota
	IdentityProviderConfigurationState_ENFORCED
	IdentityProviderConfigurationState_UNCONFIGURED
)

type IpAllowListEnabledSettingValue int

const (
	IpAllowListEnabledSettingValue_DISABLED IpAllowListEnabledSettingValue = iota
	IpAllowListEnabledSettingValue_ENABLED
)

type IpAllowListEntryOrderField int

const (
	IpAllowListEntryOrderField_ALLOW_LIST_VALUE IpAllowListEntryOrderField = iota
	IpAllowListEntryOrderField_CREATED_AT
)

type IpAllowListForInstalledAppsEnabledSettingValue int

const (
	IpAllowListForInstalledAppsEnabledSettingValue_DISABLED IpAllowListForInstalledAppsEnabledSettingValue = iota
	IpAllowListForInstalledAppsEnabledSettingValue_ENABLED
)

type IssueCommentOrderField int

const (
	IssueCommentOrderField_UPDATED_AT IssueCommentOrderField = iota
)

type IssueOrderField int

const (
	IssueOrderField_COMMENTS IssueOrderField = iota
	IssueOrderField_CREATED_AT
	IssueOrderField_UPDATED_AT
)

type IssueState int

const (
	IssueState_CLOSED IssueState = iota
	IssueState_OPEN
)

type IssueTimelineItemsItemType int

const (
	IssueTimelineItemsItemType_ADDED_TO_PROJECT_EVENT IssueTimelineItemsItemType = iota
	IssueTimelineItemsItemType_ASSIGNED_EVENT
	IssueTimelineItemsItemType_CLOSED_EVENT
	IssueTimelineItemsItemType_COMMENT_DELETED_EVENT
	IssueTimelineItemsItemType_CONNECTED_EVENT
	IssueTimelineItemsItemType_CONVERTED_NOTE_TO_ISSUE_EVENT
	IssueTimelineItemsItemType_CROSS_REFERENCED_EVENT
	IssueTimelineItemsItemType_DEMILESTONED_EVENT
	IssueTimelineItemsItemType_DISCONNECTED_EVENT
	IssueTimelineItemsItemType_ISSUE_COMMENT
	IssueTimelineItemsItemType_LABELED_EVENT
	IssueTimelineItemsItemType_LOCKED_EVENT
	IssueTimelineItemsItemType_MARKED_AS_DUPLICATE_EVENT
	IssueTimelineItemsItemType_MENTIONED_EVENT
	IssueTimelineItemsItemType_MILESTONED_EVENT
	IssueTimelineItemsItemType_MOVED_COLUMNS_IN_PROJECT_EVENT
	IssueTimelineItemsItemType_PINNED_EVENT
	IssueTimelineItemsItemType_REFERENCED_EVENT
	IssueTimelineItemsItemType_REMOVED_FROM_PROJECT_EVENT
	IssueTimelineItemsItemType_RENAMED_TITLE_EVENT
	IssueTimelineItemsItemType_REOPENED_EVENT
	IssueTimelineItemsItemType_SUBSCRIBED_EVENT
	IssueTimelineItemsItemType_TRANSFERRED_EVENT
	IssueTimelineItemsItemType_UNASSIGNED_EVENT
	IssueTimelineItemsItemType_UNLABELED_EVENT
	IssueTimelineItemsItemType_UNLOCKED_EVENT
	IssueTimelineItemsItemType_UNMARKED_AS_DUPLICATE_EVENT
	IssueTimelineItemsItemType_UNPINNED_EVENT
	IssueTimelineItemsItemType_UNSUBSCRIBED_EVENT
	IssueTimelineItemsItemType_USER_BLOCKED_EVENT
)

type LabelOrderField int

const (
	LabelOrderField_CREATED_AT LabelOrderField = iota
	LabelOrderField_NAME
)

type LanguageOrderField int

const (
	LanguageOrderField_SIZE LanguageOrderField = iota
)

type LockReason int

const (
	LockReason_OFF_TOPIC LockReason = iota
	LockReason_RESOLVED
	LockReason_SPAM
	LockReason_TOO_HEATED
)

type MergeStateStatus int

const (
	MergeStateStatus_BEHIND MergeStateStatus = iota
	MergeStateStatus_BLOCKED
	MergeStateStatus_CLEAN
	MergeStateStatus_DIRTY
	MergeStateStatus_DRAFT
	MergeStateStatus_HAS_HOOKS
	MergeStateStatus_UNKNOWN
	MergeStateStatus_UNSTABLE
)

type MergeableState int

const (
	MergeableState_CONFLICTING MergeableState = iota
	MergeableState_MERGEABLE
	MergeableState_UNKNOWN
)

type MilestoneOrderField int

const (
	MilestoneOrderField_CREATED_AT MilestoneOrderField = iota
	MilestoneOrderField_DUE_DATE
	MilestoneOrderField_NUMBER
	MilestoneOrderField_UPDATED_AT
)

type MilestoneState int

const (
	MilestoneState_CLOSED MilestoneState = iota
	MilestoneState_OPEN
)

type NotificationRestrictionSettingValue int

const (
	NotificationRestrictionSettingValue_DISABLED NotificationRestrictionSettingValue = iota
	NotificationRestrictionSettingValue_ENABLED
)

type OauthApplicationCreateAuditEntryState int

const (
	OauthApplicationCreateAuditEntryState_ACTIVE OauthApplicationCreateAuditEntryState = iota
	OauthApplicationCreateAuditEntryState_PENDING_DELETION
	OauthApplicationCreateAuditEntryState_SUSPENDED
)

type OperationType int

const (
	OperationType_ACCESS OperationType = iota
	OperationType_AUTHENTICATION
	OperationType_CREATE
	OperationType_MODIFY
	OperationType_REMOVE
	OperationType_RESTORE
	OperationType_TRANSFER
)

type OrderDirection int

const (
	OrderDirection_ASC OrderDirection = iota
	OrderDirection_DESC
)

type OrgAddMemberAuditEntryPermission int

const (
	OrgAddMemberAuditEntryPermission_ADMIN OrgAddMemberAuditEntryPermission = iota
	OrgAddMemberAuditEntryPermission_READ
)

type OrgCreateAuditEntryBillingPlan int

const (
	OrgCreateAuditEntryBillingPlan_BUSINESS OrgCreateAuditEntryBillingPlan = iota
	OrgCreateAuditEntryBillingPlan_BUSINESS_PLUS
	OrgCreateAuditEntryBillingPlan_FREE
	OrgCreateAuditEntryBillingPlan_TIERED_PER_SEAT
	OrgCreateAuditEntryBillingPlan_UNLIMITED
)

type OrgRemoveBillingManagerAuditEntryReason int

const (
	OrgRemoveBillingManagerAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING OrgRemoveBillingManagerAuditEntryReason = iota
	OrgRemoveBillingManagerAuditEntryReason_SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY
	OrgRemoveBillingManagerAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE
)

type OrgRemoveMemberAuditEntryMembershipType int

const (
	OrgRemoveMemberAuditEntryMembershipType_ADMIN OrgRemoveMemberAuditEntryMembershipType = iota
	OrgRemoveMemberAuditEntryMembershipType_BILLING_MANAGER
	OrgRemoveMemberAuditEntryMembershipType_DIRECT_MEMBER
	OrgRemoveMemberAuditEntryMembershipType_OUTSIDE_COLLABORATOR
	OrgRemoveMemberAuditEntryMembershipType_UNAFFILIATED
)

type OrgRemoveMemberAuditEntryReason int

const (
	OrgRemoveMemberAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING OrgRemoveMemberAuditEntryReason = iota
	OrgRemoveMemberAuditEntryReason_SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY
	OrgRemoveMemberAuditEntryReason_TWO_FACTOR_ACCOUNT_RECOVERY
	OrgRemoveMemberAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE
	OrgRemoveMemberAuditEntryReason_USER_ACCOUNT_DELETED
)

type OrgRemoveOutsideCollaboratorAuditEntryMembershipType int

const (
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_BILLING_MANAGER OrgRemoveOutsideCollaboratorAuditEntryMembershipType = iota
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_OUTSIDE_COLLABORATOR
	OrgRemoveOutsideCollaboratorAuditEntryMembershipType_UNAFFILIATED
)

type OrgRemoveOutsideCollaboratorAuditEntryReason int

const (
	OrgRemoveOutsideCollaboratorAuditEntryReason_SAML_EXTERNAL_IDENTITY_MISSING OrgRemoveOutsideCollaboratorAuditEntryReason = iota
	OrgRemoveOutsideCollaboratorAuditEntryReason_TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE
)

type OrgUpdateDefaultRepositoryPermissionAuditEntryPermission int

const (
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_ADMIN OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = iota
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_NONE
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_READ
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermission_WRITE
)

type OrgUpdateMemberAuditEntryPermission int

const (
	OrgUpdateMemberAuditEntryPermission_ADMIN OrgUpdateMemberAuditEntryPermission = iota
	OrgUpdateMemberAuditEntryPermission_READ
)

type OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility int

const (
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_ALL OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = iota
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_INTERNAL
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_NONE
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PRIVATE
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PRIVATE_INTERNAL
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC_INTERNAL
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility_PUBLIC_PRIVATE
)

type OrganizationInvitationRole int

const (
	OrganizationInvitationRole_ADMIN OrganizationInvitationRole = iota
	OrganizationInvitationRole_BILLING_MANAGER
	OrganizationInvitationRole_DIRECT_MEMBER
	OrganizationInvitationRole_REINSTATE
)

type OrganizationInvitationType int

const (
	OrganizationInvitationType_EMAIL OrganizationInvitationType = iota
	OrganizationInvitationType_USER
)

type OrganizationMemberRole int

const (
	OrganizationMemberRole_ADMIN OrganizationMemberRole = iota
	OrganizationMemberRole_MEMBER
)

type OrganizationMembersCanCreateRepositoriesSettingValue int

const (
	OrganizationMembersCanCreateRepositoriesSettingValue_ALL OrganizationMembersCanCreateRepositoriesSettingValue = iota
	OrganizationMembersCanCreateRepositoriesSettingValue_DISABLED
	OrganizationMembersCanCreateRepositoriesSettingValue_PRIVATE
)

type OrganizationOrderField int

const (
	OrganizationOrderField_CREATED_AT OrganizationOrderField = iota
	OrganizationOrderField_LOGIN
)

type PackageFileOrderField int

const (
	PackageFileOrderField_CREATED_AT PackageFileOrderField = iota
)

type PackageOrderField int

const (
	PackageOrderField_CREATED_AT PackageOrderField = iota
)

type PackageType int

const (
	PackageType_DEBIAN PackageType = iota
	PackageType_DOCKER
	PackageType_MAVEN
	PackageType_NPM
	PackageType_NUGET
	PackageType_PYPI
	PackageType_RUBYGEMS
)

type PackageVersionOrderField int

const (
	PackageVersionOrderField_CREATED_AT PackageVersionOrderField = iota
)

type PinnableItemType int

const (
	PinnableItemType_GIST PinnableItemType = iota
	PinnableItemType_ISSUE
	PinnableItemType_ORGANIZATION
	PinnableItemType_PROJECT
	PinnableItemType_PULL_REQUEST
	PinnableItemType_REPOSITORY
	PinnableItemType_TEAM
	PinnableItemType_USER
)

type PinnedDiscussionGradient int

const (
	PinnedDiscussionGradient_BLUE_MINT PinnedDiscussionGradient = iota
	PinnedDiscussionGradient_BLUE_PURPLE
	PinnedDiscussionGradient_PINK_BLUE
	PinnedDiscussionGradient_PURPLE_CORAL
	PinnedDiscussionGradient_RED_ORANGE
)

type PinnedDiscussionPattern int

const (
	PinnedDiscussionPattern_CHEVRON_UP PinnedDiscussionPattern = iota
	PinnedDiscussionPattern_DOT
	PinnedDiscussionPattern_DOT_FILL
	PinnedDiscussionPattern_HEART_FILL
	PinnedDiscussionPattern_PLUS
	PinnedDiscussionPattern_ZAP
)

type ProjectCardArchivedState int

const (
	ProjectCardArchivedState_ARCHIVED ProjectCardArchivedState = iota
	ProjectCardArchivedState_NOT_ARCHIVED
)

type ProjectCardState int

const (
	ProjectCardState_CONTENT_ONLY ProjectCardState = iota
	ProjectCardState_NOTE_ONLY
	ProjectCardState_REDACTED
)

type ProjectColumnPurpose int

const (
	ProjectColumnPurpose_DONE ProjectColumnPurpose = iota
	ProjectColumnPurpose_IN_PROGRESS
	ProjectColumnPurpose_TODO
)

type ProjectOrderField int

const (
	ProjectOrderField_CREATED_AT ProjectOrderField = iota
	ProjectOrderField_NAME
	ProjectOrderField_UPDATED_AT
)

type ProjectState int

const (
	ProjectState_CLOSED ProjectState = iota
	ProjectState_OPEN
)

type ProjectTemplate int

const (
	ProjectTemplate_AUTOMATED_KANBAN_V2 ProjectTemplate = iota
	ProjectTemplate_AUTOMATED_REVIEWS_KANBAN
	ProjectTemplate_BASIC_KANBAN
	ProjectTemplate_BUG_TRIAGE
)

type PullRequestMergeMethod int

const (
	PullRequestMergeMethod_MERGE PullRequestMergeMethod = iota
	PullRequestMergeMethod_REBASE
	PullRequestMergeMethod_SQUASH
)

type PullRequestOrderField int

const (
	PullRequestOrderField_CREATED_AT PullRequestOrderField = iota
	PullRequestOrderField_UPDATED_AT
)

type PullRequestReviewCommentState int

const (
	PullRequestReviewCommentState_PENDING PullRequestReviewCommentState = iota
	PullRequestReviewCommentState_SUBMITTED
)

type PullRequestReviewDecision int

const (
	PullRequestReviewDecision_APPROVED PullRequestReviewDecision = iota
	PullRequestReviewDecision_CHANGES_REQUESTED
	PullRequestReviewDecision_REVIEW_REQUIRED
)

type PullRequestReviewEvent int

const (
	PullRequestReviewEvent_APPROVE PullRequestReviewEvent = iota
	PullRequestReviewEvent_COMMENT
	PullRequestReviewEvent_DISMISS
	PullRequestReviewEvent_REQUEST_CHANGES
)

type PullRequestReviewState int

const (
	PullRequestReviewState_APPROVED PullRequestReviewState = iota
	PullRequestReviewState_CHANGES_REQUESTED
	PullRequestReviewState_COMMENTED
	PullRequestReviewState_DISMISSED
	PullRequestReviewState_PENDING
)

type PullRequestState int

const (
	PullRequestState_CLOSED PullRequestState = iota
	PullRequestState_MERGED
	PullRequestState_OPEN
)

type PullRequestTimelineItemsItemType int

const (
	PullRequestTimelineItemsItemType_ADDED_TO_PROJECT_EVENT PullRequestTimelineItemsItemType = iota
	PullRequestTimelineItemsItemType_ASSIGNED_EVENT
	PullRequestTimelineItemsItemType_AUTOMATIC_BASE_CHANGE_FAILED_EVENT
	PullRequestTimelineItemsItemType_AUTOMATIC_BASE_CHANGE_SUCCEEDED_EVENT
	PullRequestTimelineItemsItemType_AUTO_MERGE_DISABLED_EVENT
	PullRequestTimelineItemsItemType_AUTO_MERGE_ENABLED_EVENT
	PullRequestTimelineItemsItemType_AUTO_REBASE_ENABLED_EVENT
	PullRequestTimelineItemsItemType_AUTO_SQUASH_ENABLED_EVENT
	PullRequestTimelineItemsItemType_BASE_REF_CHANGED_EVENT
	PullRequestTimelineItemsItemType_BASE_REF_DELETED_EVENT
	PullRequestTimelineItemsItemType_BASE_REF_FORCE_PUSHED_EVENT
	PullRequestTimelineItemsItemType_CLOSED_EVENT
	PullRequestTimelineItemsItemType_COMMENT_DELETED_EVENT
	PullRequestTimelineItemsItemType_CONNECTED_EVENT
	PullRequestTimelineItemsItemType_CONVERTED_NOTE_TO_ISSUE_EVENT
	PullRequestTimelineItemsItemType_CONVERT_TO_DRAFT_EVENT
	PullRequestTimelineItemsItemType_CROSS_REFERENCED_EVENT
	PullRequestTimelineItemsItemType_DEMILESTONED_EVENT
	PullRequestTimelineItemsItemType_DEPLOYED_EVENT
	PullRequestTimelineItemsItemType_DEPLOYMENT_ENVIRONMENT_CHANGED_EVENT
	PullRequestTimelineItemsItemType_DISCONNECTED_EVENT
	PullRequestTimelineItemsItemType_HEAD_REF_DELETED_EVENT
	PullRequestTimelineItemsItemType_HEAD_REF_FORCE_PUSHED_EVENT
	PullRequestTimelineItemsItemType_HEAD_REF_RESTORED_EVENT
	PullRequestTimelineItemsItemType_ISSUE_COMMENT
	PullRequestTimelineItemsItemType_LABELED_EVENT
	PullRequestTimelineItemsItemType_LOCKED_EVENT
	PullRequestTimelineItemsItemType_MARKED_AS_DUPLICATE_EVENT
	PullRequestTimelineItemsItemType_MENTIONED_EVENT
	PullRequestTimelineItemsItemType_MERGED_EVENT
	PullRequestTimelineItemsItemType_MILESTONED_EVENT
	PullRequestTimelineItemsItemType_MOVED_COLUMNS_IN_PROJECT_EVENT
	PullRequestTimelineItemsItemType_PINNED_EVENT
	PullRequestTimelineItemsItemType_PULL_REQUEST_COMMIT
	PullRequestTimelineItemsItemType_PULL_REQUEST_COMMIT_COMMENT_THREAD
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVIEW
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVIEW_THREAD
	PullRequestTimelineItemsItemType_PULL_REQUEST_REVISION_MARKER
	PullRequestTimelineItemsItemType_READY_FOR_REVIEW_EVENT
	PullRequestTimelineItemsItemType_REFERENCED_EVENT
	PullRequestTimelineItemsItemType_REMOVED_FROM_PROJECT_EVENT
	PullRequestTimelineItemsItemType_RENAMED_TITLE_EVENT
	PullRequestTimelineItemsItemType_REOPENED_EVENT
	PullRequestTimelineItemsItemType_REVIEW_DISMISSED_EVENT
	PullRequestTimelineItemsItemType_REVIEW_REQUESTED_EVENT
	PullRequestTimelineItemsItemType_REVIEW_REQUEST_REMOVED_EVENT
	PullRequestTimelineItemsItemType_SUBSCRIBED_EVENT
	PullRequestTimelineItemsItemType_TRANSFERRED_EVENT
	PullRequestTimelineItemsItemType_UNASSIGNED_EVENT
	PullRequestTimelineItemsItemType_UNLABELED_EVENT
	PullRequestTimelineItemsItemType_UNLOCKED_EVENT
	PullRequestTimelineItemsItemType_UNMARKED_AS_DUPLICATE_EVENT
	PullRequestTimelineItemsItemType_UNPINNED_EVENT
	PullRequestTimelineItemsItemType_UNSUBSCRIBED_EVENT
	PullRequestTimelineItemsItemType_USER_BLOCKED_EVENT
)

type PullRequestUpdateState int

const (
	PullRequestUpdateState_CLOSED PullRequestUpdateState = iota
	PullRequestUpdateState_OPEN
)

type ReactionContent int

const (
	ReactionContent_CONFUSED ReactionContent = iota
	ReactionContent_EYES
	ReactionContent_HEART
	ReactionContent_HOORAY
	ReactionContent_LAUGH
	ReactionContent_ROCKET
	ReactionContent_THUMBS_DOWN
	ReactionContent_THUMBS_UP
)

type ReactionOrderField int

const (
	ReactionOrderField_CREATED_AT ReactionOrderField = iota
)

type RefOrderField int

const (
	RefOrderField_ALPHABETICAL RefOrderField = iota
	RefOrderField_TAG_COMMIT_DATE
)

type ReleaseOrderField int

const (
	ReleaseOrderField_CREATED_AT ReleaseOrderField = iota
	ReleaseOrderField_NAME
)

type RepoAccessAuditEntryVisibility int

const (
	RepoAccessAuditEntryVisibility_INTERNAL RepoAccessAuditEntryVisibility = iota
	RepoAccessAuditEntryVisibility_PRIVATE
	RepoAccessAuditEntryVisibility_PUBLIC
)

type RepoAddMemberAuditEntryVisibility int

const (
	RepoAddMemberAuditEntryVisibility_INTERNAL RepoAddMemberAuditEntryVisibility = iota
	RepoAddMemberAuditEntryVisibility_PRIVATE
	RepoAddMemberAuditEntryVisibility_PUBLIC
)

type RepoArchivedAuditEntryVisibility int

const (
	RepoArchivedAuditEntryVisibility_INTERNAL RepoArchivedAuditEntryVisibility = iota
	RepoArchivedAuditEntryVisibility_PRIVATE
	RepoArchivedAuditEntryVisibility_PUBLIC
)

type RepoChangeMergeSettingAuditEntryMergeType int

const (
	RepoChangeMergeSettingAuditEntryMergeType_MERGE RepoChangeMergeSettingAuditEntryMergeType = iota
	RepoChangeMergeSettingAuditEntryMergeType_REBASE
	RepoChangeMergeSettingAuditEntryMergeType_SQUASH
)

type RepoCreateAuditEntryVisibility int

const (
	RepoCreateAuditEntryVisibility_INTERNAL RepoCreateAuditEntryVisibility = iota
	RepoCreateAuditEntryVisibility_PRIVATE
	RepoCreateAuditEntryVisibility_PUBLIC
)

type RepoDestroyAuditEntryVisibility int

const (
	RepoDestroyAuditEntryVisibility_INTERNAL RepoDestroyAuditEntryVisibility = iota
	RepoDestroyAuditEntryVisibility_PRIVATE
	RepoDestroyAuditEntryVisibility_PUBLIC
)

type RepoRemoveMemberAuditEntryVisibility int

const (
	RepoRemoveMemberAuditEntryVisibility_INTERNAL RepoRemoveMemberAuditEntryVisibility = iota
	RepoRemoveMemberAuditEntryVisibility_PRIVATE
	RepoRemoveMemberAuditEntryVisibility_PUBLIC
)

type ReportedContentClassifiers int

const (
	ReportedContentClassifiers_ABUSE ReportedContentClassifiers = iota
	ReportedContentClassifiers_DUPLICATE
	ReportedContentClassifiers_OFF_TOPIC
	ReportedContentClassifiers_OUTDATED
	ReportedContentClassifiers_RESOLVED
	ReportedContentClassifiers_SPAM
)

type RepositoryAffiliation int

const (
	RepositoryAffiliation_COLLABORATOR RepositoryAffiliation = iota
	RepositoryAffiliation_ORGANIZATION_MEMBER
	RepositoryAffiliation_OWNER
)

type RepositoryContributionType int

const (
	RepositoryContributionType_COMMIT RepositoryContributionType = iota
	RepositoryContributionType_ISSUE
	RepositoryContributionType_PULL_REQUEST
	RepositoryContributionType_PULL_REQUEST_REVIEW
	RepositoryContributionType_REPOSITORY
)

type RepositoryInteractionLimit int

const (
	RepositoryInteractionLimit_COLLABORATORS_ONLY RepositoryInteractionLimit = iota
	RepositoryInteractionLimit_CONTRIBUTORS_ONLY
	RepositoryInteractionLimit_EXISTING_USERS
	RepositoryInteractionLimit_NO_LIMIT
)

type RepositoryInteractionLimitExpiry int

const (
	RepositoryInteractionLimitExpiry_ONE_DAY RepositoryInteractionLimitExpiry = iota
	RepositoryInteractionLimitExpiry_ONE_MONTH
	RepositoryInteractionLimitExpiry_ONE_WEEK
	RepositoryInteractionLimitExpiry_SIX_MONTHS
	RepositoryInteractionLimitExpiry_THREE_DAYS
)

type RepositoryInteractionLimitOrigin int

const (
	RepositoryInteractionLimitOrigin_ORGANIZATION RepositoryInteractionLimitOrigin = iota
	RepositoryInteractionLimitOrigin_REPOSITORY
	RepositoryInteractionLimitOrigin_USER
)

type RepositoryInvitationOrderField int

const (
	RepositoryInvitationOrderField_CREATED_AT RepositoryInvitationOrderField = iota
	RepositoryInvitationOrderField_INVITEE_LOGIN
)

type RepositoryLockReason int

const (
	RepositoryLockReason_BILLING RepositoryLockReason = iota
	RepositoryLockReason_MIGRATING
	RepositoryLockReason_MOVING
	RepositoryLockReason_RENAME
)

type RepositoryOrderField int

const (
	RepositoryOrderField_CREATED_AT RepositoryOrderField = iota
	RepositoryOrderField_NAME
	RepositoryOrderField_PUSHED_AT
	RepositoryOrderField_STARGAZERS
	RepositoryOrderField_UPDATED_AT
)

type RepositoryPermission int

const (
	RepositoryPermission_ADMIN RepositoryPermission = iota
	RepositoryPermission_MAINTAIN
	RepositoryPermission_READ
	RepositoryPermission_TRIAGE
	RepositoryPermission_WRITE
)

type RepositoryPrivacy int

const (
	RepositoryPrivacy_PRIVATE RepositoryPrivacy = iota
	RepositoryPrivacy_PUBLIC
)

type RepositoryVisibility int

const (
	RepositoryVisibility_INTERNAL RepositoryVisibility = iota
	RepositoryVisibility_PRIVATE
	RepositoryVisibility_PUBLIC
)

type RequestableCheckStatusState int

const (
	RequestableCheckStatusState_COMPLETED RequestableCheckStatusState = iota
	RequestableCheckStatusState_IN_PROGRESS
	RequestableCheckStatusState_PENDING
	RequestableCheckStatusState_QUEUED
	RequestableCheckStatusState_WAITING
)

type SamlDigestAlgorithm int

const (
	SamlDigestAlgorithm_SHA1 SamlDigestAlgorithm = iota
	SamlDigestAlgorithm_SHA256
	SamlDigestAlgorithm_SHA384
	SamlDigestAlgorithm_SHA512
)

type SamlSignatureAlgorithm int

const (
	SamlSignatureAlgorithm_RSA_SHA1 SamlSignatureAlgorithm = iota
	SamlSignatureAlgorithm_RSA_SHA256
	SamlSignatureAlgorithm_RSA_SHA384
	SamlSignatureAlgorithm_RSA_SHA512
)

type SavedReplyOrderField int

const (
	SavedReplyOrderField_UPDATED_AT SavedReplyOrderField = iota
)

type SearchType int

const (
	SearchType_DISCUSSION SearchType = iota
	SearchType_ISSUE
	SearchType_REPOSITORY
	SearchType_USER
)

type SecurityAdvisoryEcosystem int

const (
	SecurityAdvisoryEcosystem_COMPOSER SecurityAdvisoryEcosystem = iota
	SecurityAdvisoryEcosystem_GO
	SecurityAdvisoryEcosystem_MAVEN
	SecurityAdvisoryEcosystem_NPM
	SecurityAdvisoryEcosystem_NUGET
	SecurityAdvisoryEcosystem_OTHER
	SecurityAdvisoryEcosystem_PIP
	SecurityAdvisoryEcosystem_RUBYGEMS
)

type SecurityAdvisoryIdentifierType int

const (
	SecurityAdvisoryIdentifierType_CVE SecurityAdvisoryIdentifierType = iota
	SecurityAdvisoryIdentifierType_GHSA
)

type SecurityAdvisoryOrderField int

const (
	SecurityAdvisoryOrderField_PUBLISHED_AT SecurityAdvisoryOrderField = iota
	SecurityAdvisoryOrderField_UPDATED_AT
)

type SecurityAdvisorySeverity int

const (
	SecurityAdvisorySeverity_CRITICAL SecurityAdvisorySeverity = iota
	SecurityAdvisorySeverity_HIGH
	SecurityAdvisorySeverity_LOW
	SecurityAdvisorySeverity_MODERATE
)

type SecurityVulnerabilityOrderField int

const (
	SecurityVulnerabilityOrderField_UPDATED_AT SecurityVulnerabilityOrderField = iota
)

type SponsorableOrderField int

const (
	SponsorableOrderField_LOGIN SponsorableOrderField = iota
)

type SponsorsActivityAction int

const (
	SponsorsActivityAction_CANCELLED_SPONSORSHIP SponsorsActivityAction = iota
	SponsorsActivityAction_NEW_SPONSORSHIP
	SponsorsActivityAction_PENDING_CHANGE
	SponsorsActivityAction_REFUND
	SponsorsActivityAction_SPONSOR_MATCH_DISABLED
	SponsorsActivityAction_TIER_CHANGE
)

type SponsorsActivityOrderField int

const (
	SponsorsActivityOrderField_TIMESTAMP SponsorsActivityOrderField = iota
)

type SponsorsActivityPeriod int

const (
	SponsorsActivityPeriod_ALL SponsorsActivityPeriod = iota
	SponsorsActivityPeriod_DAY
	SponsorsActivityPeriod_MONTH
	SponsorsActivityPeriod_WEEK
)

type SponsorsGoalKind int

const (
	SponsorsGoalKind_MONTHLY_SPONSORSHIP_AMOUNT SponsorsGoalKind = iota
	SponsorsGoalKind_TOTAL_SPONSORS_COUNT
)

type SponsorsTierOrderField int

const (
	SponsorsTierOrderField_CREATED_AT SponsorsTierOrderField = iota
	SponsorsTierOrderField_MONTHLY_PRICE_IN_CENTS
)

type SponsorshipOrderField int

const (
	SponsorshipOrderField_CREATED_AT SponsorshipOrderField = iota
)

type SponsorshipPrivacy int

const (
	SponsorshipPrivacy_PRIVATE SponsorshipPrivacy = iota
	SponsorshipPrivacy_PUBLIC
)

type StarOrderField int

const (
	StarOrderField_STARRED_AT StarOrderField = iota
)

type StatusState int

const (
	StatusState_ERROR StatusState = iota
	StatusState_EXPECTED
	StatusState_FAILURE
	StatusState_PENDING
	StatusState_SUCCESS
)

type SubscriptionState int

const (
	SubscriptionState_IGNORED SubscriptionState = iota
	SubscriptionState_SUBSCRIBED
	SubscriptionState_UNSUBSCRIBED
)

type TeamDiscussionCommentOrderField int

const (
	TeamDiscussionCommentOrderField_NUMBER TeamDiscussionCommentOrderField = iota
)

type TeamDiscussionOrderField int

const (
	TeamDiscussionOrderField_CREATED_AT TeamDiscussionOrderField = iota
)

type TeamMemberOrderField int

const (
	TeamMemberOrderField_CREATED_AT TeamMemberOrderField = iota
	TeamMemberOrderField_LOGIN
)

type TeamMemberRole int

const (
	TeamMemberRole_MAINTAINER TeamMemberRole = iota
	TeamMemberRole_MEMBER
)

type TeamMembershipType int

const (
	TeamMembershipType_ALL TeamMembershipType = iota
	TeamMembershipType_CHILD_TEAM
	TeamMembershipType_IMMEDIATE
)

type TeamOrderField int

const (
	TeamOrderField_NAME TeamOrderField = iota
)

type TeamPrivacy int

const (
	TeamPrivacy_SECRET TeamPrivacy = iota
	TeamPrivacy_VISIBLE
)

type TeamRepositoryOrderField int

const (
	TeamRepositoryOrderField_CREATED_AT TeamRepositoryOrderField = iota
	TeamRepositoryOrderField_NAME
	TeamRepositoryOrderField_PERMISSION
	TeamRepositoryOrderField_PUSHED_AT
	TeamRepositoryOrderField_STARGAZERS
	TeamRepositoryOrderField_UPDATED_AT
)

type TeamReviewAssignmentAlgorithm int

const (
	TeamReviewAssignmentAlgorithm_LOAD_BALANCE TeamReviewAssignmentAlgorithm = iota
	TeamReviewAssignmentAlgorithm_ROUND_ROBIN
)

type TeamRole int

const (
	TeamRole_ADMIN TeamRole = iota
	TeamRole_MEMBER
)

type TopicSuggestionDeclineReason int

const (
	TopicSuggestionDeclineReason_NOT_RELEVANT TopicSuggestionDeclineReason = iota
	TopicSuggestionDeclineReason_PERSONAL_PREFERENCE
	TopicSuggestionDeclineReason_TOO_GENERAL
	TopicSuggestionDeclineReason_TOO_SPECIFIC
)

type UserBlockDuration int

const (
	UserBlockDuration_ONE_DAY UserBlockDuration = iota
	UserBlockDuration_ONE_MONTH
	UserBlockDuration_ONE_WEEK
	UserBlockDuration_PERMANENT
	UserBlockDuration_THREE_DAYS
)

type UserStatusOrderField int

const (
	UserStatusOrderField_UPDATED_AT UserStatusOrderField = iota
)

type VerifiableDomainOrderField int

const (
	VerifiableDomainOrderField_CREATED_AT VerifiableDomainOrderField = iota
	VerifiableDomainOrderField_DOMAIN
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
