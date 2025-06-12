package v2

type Pagination struct {
	Count   int `json:"count"`
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per"`
}

type RunCommentsMode int
type RunVerifiedStatus int
type RunObsoleteFilter int
type VideoFilter int
type AuditLogEventType string
type GameType int
type ForumType int
type ResourceType int
type ValidTimers int
type VariableCategoryScope int
type VariableLevelScope int
type TicketQueue int
type TicketType int
type TicketStatus int
type GameModeratorLevel int
type UserSocialConnection int

const (
	RunCommentsDisabled RunCommentsMode = 1

	RunVerifiedStatusPending  RunVerifiedStatus = 0
	RunVerifiedStatusVerified RunVerifiedStatus = 1
	RunVerifiedStatusRejected RunVerifiedStatus = 2

	RunObsoleteFilterHidden    RunObsoleteFilter = 0
	RunObsoleteFilterShown     RunObsoleteFilter = 1
	RunObsoleteFilterExclusive RunObsoleteFilter = 2

	VideoFilterOptional VideoFilter = 0
	VideoFilterRequired VideoFilter = 1
	VideoFilterMissing  VideoFilter = 2

	AuditLogEventTypeCategoryCreated      AuditLogEventType = "category-created"
	AuditLogEventTypeCategoryRemoved      AuditLogEventType = "category-removed"
	AuditLogEventTypeCategoryRestored     AuditLogEventType = "category-restored"
	AuditLogEventTypeCategoryUpdated      AuditLogEventType = "category-updated"
	AuditLogEventTypeCommentCreated       AuditLogEventType = "comment-created"
	AuditLogEventTypeCommentDeleted       AuditLogEventType = "comment-deleted"
	AuditLogEventTypeCommentUpdated       AuditLogEventType = "comment-updated"
	AuditLogEventTypeGameCoversUpdated    AuditLogEventType = "game-covers-updated"
	AuditLogEventTypeGameCreated          AuditLogEventType = "game-created"
	AuditLogEventTypeGameModeratorCreated AuditLogEventType = "game-moderator-created"
	AuditLogEventTypeGameModeratorRemoved AuditLogEventType = "game-moderator-removed"
	AuditLogEventTypeGameModeratorUpdated AuditLogEventType = "game-moderator-updated"
	AuditLogEventTypeGameNewsPostCreated  AuditLogEventType = "game-news-post-created"
	AuditLogEventTypeGameNewsPostEdited   AuditLogEventType = "game-news-post-edited"
	AuditLogEventTypeGameNewsPostRemoved  AuditLogEventType = "game-news-post-removed"
	AuditLogEventTypeGameRestored         AuditLogEventType = "game-restored"
	AuditLogEventTypeGameUpdated          AuditLogEventType = "game-updated"
	AuditLogEventTypeGameRequestReviewed  AuditLogEventType = "gamerequest-reviewed"
	AuditLogEventTypeLevelCreated         AuditLogEventType = "level-created"

	GameTypeRomHack           GameType = 1
	GameTypeModification      GameType = 2
	GameTypeFanGame           GameType = 3
	GameTypeWebGame           GameType = 4
	GameTypePrerelease        GameType = 5
	GameTypeMobile            GameType = 6
	GameTypeExpansionDLC      GameType = 7
	GameTypeCategoryExtension GameType = 8
	GameTypeMultipleGames     GameType = 9
	GameTypeMinigame          GameType = 10
	GameTypeServerMap         GameType = 11
	GameTypeHomebrew          GameType = 12

	ForumTypeFrontPage ForumType = 0
	ForumTypeSupporter ForumType = 1
	ForumTypeGame      ForumType = 2

	ResourceTypeTool   ResourceType = 1
	ResourceTypeSave   ResourceType = 2
	ResourceTypeSplits ResourceType = 3
	ResourceTypePatch  ResourceType = 4

	ValidTimersRTA ValidTimers = 0
	ValidTimersLRT ValidTimers = 1
	ValidTimersIGT ValidTimers = 2

	VariableCategoryScopeAll    VariableCategoryScope = -1
	VariableCategoryScopeSingle VariableCategoryScope = 1

	VariableLevelScopeFullGameAllLevels VariableLevelScope = -2
	VariableLevelScopeAllLevels         VariableLevelScope = -1
	VariableLevelScopeFullGame          VariableLevelScope = 0
	VariableLevelScopeSingleLevel       VariableLevelScope = 1

	TicketQueueGameRequests      TicketQueue = 1
	TicketQueueSeriesRequests    TicketQueue = 2
	TicketQueueModeratorReports  TicketQueue = 3
	TicketQueueMarathonRequests  TicketQueue = 4
	TicketQueueContentReports    TicketQueue = 5
	TicketQueueUserReports       TicketQueue = 6
	TicketQueueBugReports        TicketQueue = 7
	TicketQueueFrontPageRequests TicketQueue = 8
	TicketQueueFeedback          TicketQueue = 9
	TicketQueueStaffApplications TicketQueue = 10
	TicketQueueSupport           TicketQueue = 11
	TicketQueueContentRequests   TicketQueue = 12
	TicketQueueSupporter         TicketQueue = 13

	TicketTypeGameRequest        TicketType = 1
	TicketTypeSeriesRequest      TicketType = 2
	TicketTypeModRequest         TicketType = 3
	TicketTypeMarathonRequest    TicketType = 4
	TicketTypeContentReport      TicketType = 5
	TicketTypeUserReport         TicketType = 6
	TicketTypeBugReport          TicketType = 7
	TicketTypeFrontPageRequest   TicketType = 8
	TicketTypeFeedback           TicketType = 9
	TicketTypeStaffApplication   TicketType = 10
	TicketTypeOtherSupport       TicketType = 11
	TicketTypeGameTypeUpdate     TicketType = 12
	TicketTypeAddToSeriesRequest TicketType = 13
	TicketTypeAddPlatformRequest TicketType = 14
	TicketTypeOtherGameRequest   TicketType = 15
	TicketTypeSupporterHelp      TicketType = 16

	TicketStatusPending   TicketStatus = 0
	TicketStatusApproved  TicketStatus = 1
	TicketStatusDenied    TicketStatus = 2
	TicketStatusReviewing TicketStatus = 3
	TicketStatusWithdrawn TicketStatus = 4

	GameModeratorLevelVerifier       GameModeratorLevel = -1
	GameModeratorLevelModerator      GameModeratorLevel = 0
	GameModeratorLevelSuperModerator GameModeratorLevel = 1

	// Deprecated Connection
	UserSocialConnectionAskFM UserSocialConnection = 1
	// Deprecated Connection
	UserSocialConnectionBattlenet UserSocialConnection = 2
	UserSocialConnectionBilibili  UserSocialConnection = 3
	// Deprecated Connection
	UserSocialConnectionDeviantArt UserSocialConnection = 4
	UserSocialConnectionDiscord    UserSocialConnection = 5
	// Deprecated Connection
	UserSocialConnectionDouyu UserSocialConnection = 6
	// Deprecated Connection
	UserSocialConnectionDuolingo UserSocialConnection = 7
	UserSocialConnectionFacebook UserSocialConnection = 8
	// Deprecated Connection
	UserSocialConnectionGooglePlus UserSocialConnection = 9
	// Deprecated Connection
	UserSocialConnectionGPodcasts UserSocialConnection = 10
	UserSocialConnectionInstagram UserSocialConnection = 11
	// Deprecated Connection
	UserSocialConnectionITunes UserSocialConnection = 12
	// Deprecated Connection
	UserSocialConnectionMixer UserSocialConnection = 13
	// Deprecated Connection
	UserSocialConnectionMMRTA    UserSocialConnection = 14
	UserSocialConnectionNicoNico UserSocialConnection = 15
	// Deprecated Connection
	UserSocialConnectionPatreon UserSocialConnection = 16
	// Deprecated Connection
	UserSocialConnectionPinterest UserSocialConnection = 17
	UserSocialConnectionReddit    UserSocialConnection = 18
	// Deprecated Connection
	UserSocialConnectionSmashcast UserSocialConnection = 19
	// Deprecated Connection
	UserSocialConnectionSnapchat UserSocialConnection = 20
	// Deprecated Connection
	UserSocialConnectionSoundCloud UserSocialConnection = 21
	// Deprecated Connection
	UserSocialConnectionSplitsIO UserSocialConnection = 22
	// Deprecated Connection
	UserSocialConnectionSpotify UserSocialConnection = 23
	// Deprecated Connection
	UserSocialConnectionSpotifyShow UserSocialConnection = 24
	// Deprecated Connection
	UserSocialConnectionSRL UserSocialConnection = 25
	// Deprecated Connection
	UserSocialConnectionSteam UserSocialConnection = 26
	// Deprecated Connection
	UserSocialConnectionStitcher UserSocialConnection = 27
	// Deprecated Connection
	UserSocialConnectionTumblr  UserSocialConnection = 28
	UserSocialConnectionTwitch  UserSocialConnection = 29
	UserSocialConnectionTwitter UserSocialConnection = 30
	UserSocialConnectionWebsite UserSocialConnection = 31
	UserSocialConnectionYouTube UserSocialConnection = 32
	// Deprecated Connection
	UserSocialConnectionZsr UserSocialConnection = 33
)
