package v2

type Pagination struct {
	Count   int `json:"count"`
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per"`
}

type Asset struct {
	AssetType string `json:"assetType"`
	Path      string `json:"path"`
}

type ItemType int
type RunVerifiedStatus int
type RunObsoleteFilter int
type VerifiedFilter int
type VideoState int
type VideoFilter int
type EmulatorFilter int
type AuditLogEventType string
type GameType int
type GameOrderType int
type ForumType int
type ResourceType int
type ValidTimer int
type VariableCategoryScope int
type VariableLevelScope int
type VariableDisplayMode int
type TicketQueue int
type TicketType int
type TicketStatus int
type GameModeratorLevel int
type SitePowerLevel int
type UserSocialConnection int
type TimeDirection int
type PlayerMatchMode int
type PermissionType int
type IconType int
type IconPosition int
type HomepageStreamType int
type GameSortType int
type TimeDisplayUnits int
type TimeReference int
type DateFormat int
type TimeFormat int
type NavbarColorType int
type ScrollType int
type PositionType int
type RepeatType int
type FitType int
type SupportPlanPeriod string
type DefaultViewType int
type ChallengeState int

const (
	ItemTypeUnknown        ItemType = 0
	ItemTypeComment        ItemType = 1
	ItemTypeRun            ItemType = 2
	ItemTypeGame           ItemType = 3
	ItemTypeGuide          ItemType = 4
	ItemTypeResource       ItemType = 5
	ItemTypeUser           ItemType = 6
	ItemTypeThread         ItemType = 7
	ItemTypeGameMod        ItemType = 8
	ItemTypeCategory       ItemType = 9
	ItemTypeLevel          ItemType = 10
	ItemTypeGameRequest    ItemType = 11
	ItemTypeTicket         ItemType = 22
	ItemTypeTicketNote     ItemType = 23
	ItemTypeNews           ItemType = 27
	ItemTypeGameBoostToken ItemType = 28
	ItemTypeGameBoost      ItemType = 29
	ItemTypeArticle        ItemType = 30
	ItemTypeUserFollower   ItemType = 31
	ItemTypeChallenge      ItemType = 32
	ItemTypeChallengeRun   ItemType = 33

	RunVerifiedStatusPending  RunVerifiedStatus = 0
	RunVerifiedStatusVerified RunVerifiedStatus = 1
	RunVerifiedStatusRejected RunVerifiedStatus = 2

	RunObsoleteFilterHidden    RunObsoleteFilter = 0
	RunObsoleteFilterShown     RunObsoleteFilter = 1
	RunObsoleteFilterExclusive RunObsoleteFilter = 2

	VerifiedFilterAwaiting VerifiedFilter = 0
	VerifiedFilterVerified VerifiedFilter = 1
	VerifiedFilterRejected VerifiedFilter = 2

	VideoStateUnknown   VideoState = 0
	VideoStateAtRisk    VideoState = 1
	VideoStateSafe      VideoState = 2
	VideoStateAbandoned VideoState = 3

	EmulatorFilterHidden    EmulatorFilter = 0
	EmulatorFilterShown     EmulatorFilter = 1
	EmulatorFilterExclusive EmulatorFilter = 2

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

	GameOrderTypeName           GameOrderType = 1
	GameOrderTypeNewestReleased GameOrderType = 2
	GameOrderTypeOldestReleased GameOrderType = 3
	GameOrderTypeMostActive     GameOrderType = 4
	GameOrderTypeLeastActive    GameOrderType = 5
	GameOrderTypeMostPlayers    GameOrderType = 6
	GameOrderTypeLeastPlayers   GameOrderType = 7
	GameOrderTypeMostRuns       GameOrderType = 8
	GameOrderTypeLeastRuns      GameOrderType = 9
	GameOrderTypeNewestAdded    GameOrderType = 10
	GameOrderTypeOldestAdded    GameOrderType = 11

	ForumTypeFrontPage ForumType = 0
	ForumTypeSupporter ForumType = 1
	ForumTypeGame      ForumType = 2

	ResourceTypeTool   ResourceType = 1
	ResourceTypeSave   ResourceType = 2
	ResourceTypeSplits ResourceType = 3
	ResourceTypePatch  ResourceType = 4

	ValidTimersRTA ValidTimer = 0
	ValidTimersLRT ValidTimer = 1
	ValidTimersIGT ValidTimer = 2

	VariableCategoryScopeAll    VariableCategoryScope = -1
	VariableCategoryScopeSingle VariableCategoryScope = 1

	VariableLevelScopeFullGameAllLevels VariableLevelScope = -2
	VariableLevelScopeAllLevels         VariableLevelScope = -1
	VariableLevelScopeFullGame          VariableLevelScope = 0
	VariableLevelScopeSingleLevel       VariableLevelScope = 1

	VariableDisplayModeAuto     VariableDisplayMode = 0
	VariableDisplayModeDropdown VariableDisplayMode = 1
	VariableDisplayModeButtons  VariableDisplayMode = 2

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

	SitePowerLevelBanned           SitePowerLevel = 0
	SitePowerLevelUser             SitePowerLevel = 1
	SitePowerLevelContentModerator SitePowerLevel = 2
	SitePowerLevelSiteModerator    SitePowerLevel = 3
	SitePowerLevelSiteAdmin        SitePowerLevel = 4

	UserSocialConnectionAskFM       UserSocialConnection = 1 // Deprecated Connection
	UserSocialConnectionBattlenet   UserSocialConnection = 2 // Deprecated Connection
	UserSocialConnectionBilibili    UserSocialConnection = 3
	UserSocialConnectionDeviantArt  UserSocialConnection = 4 // Deprecated Connection
	UserSocialConnectionDiscord     UserSocialConnection = 5
	UserSocialConnectionDouyu       UserSocialConnection = 6 // Deprecated Connection
	UserSocialConnectionDuolingo    UserSocialConnection = 7 // Deprecated Connection
	UserSocialConnectionFacebook    UserSocialConnection = 8
	UserSocialConnectionGooglePlus  UserSocialConnection = 9  // Deprecated Connection
	UserSocialConnectionGPodcasts   UserSocialConnection = 10 // Deprecated Connection
	UserSocialConnectionInstagram   UserSocialConnection = 11
	UserSocialConnectionITunes      UserSocialConnection = 12 // Deprecated Connection
	UserSocialConnectionMixer       UserSocialConnection = 13 // Deprecated Connection
	UserSocialConnectionMMRTA       UserSocialConnection = 14 // Deprecated Connection
	UserSocialConnectionNicoNico    UserSocialConnection = 15
	UserSocialConnectionPatreon     UserSocialConnection = 16 // Deprecated Connection
	UserSocialConnectionPinterest   UserSocialConnection = 17 // Deprecated Connection
	UserSocialConnectionReddit      UserSocialConnection = 18
	UserSocialConnectionSmashcast   UserSocialConnection = 19 // Deprecated Connection
	UserSocialConnectionSnapchat    UserSocialConnection = 20 // Deprecated Connection
	UserSocialConnectionSoundCloud  UserSocialConnection = 21 // Deprecated Connection
	UserSocialConnectionSplitsIO    UserSocialConnection = 22 // Deprecated Connection
	UserSocialConnectionSpotify     UserSocialConnection = 23 // Deprecated Connection
	UserSocialConnectionSpotifyShow UserSocialConnection = 24 // Deprecated Connection
	UserSocialConnectionSRL         UserSocialConnection = 25 // Deprecated Connection
	UserSocialConnectionSteam       UserSocialConnection = 26 // Deprecated Connection
	UserSocialConnectionStitcher    UserSocialConnection = 27 // Deprecated Connection
	UserSocialConnectionTumblr      UserSocialConnection = 28 // Deprecated Connection
	UserSocialConnectionTwitch      UserSocialConnection = 29
	UserSocialConnectionTwitter     UserSocialConnection = 30
	UserSocialConnectionWebsite     UserSocialConnection = 31
	UserSocialConnectionYouTube     UserSocialConnection = 32
	UserSocialConnectionZsr         UserSocialConnection = 33 // Deprecated Connection

	TimeDirectionAscending  TimeDirection = 0
	TimeDirectionDescending TimeDirection = 1

	PlayerMatchModeAllPlayersInOrder  PlayerMatchMode = 0
	PlayerMatchModeAllPlayersAnyOrder PlayerMatchMode = 1
	PlayerMatchModeAnyPlayersInOrder  PlayerMatchMode = 2
	PlayerMatchModeAnyPlayersAnyOrder PlayerMatchMode = 3

	PermissionTypeAll          PermissionType = 0
	PermissionTypeDisabled     PermissionType = 1
	PermissionTypeVerifiedHere PermissionType = 2
	PermissionTypeVerifiedAny  PermissionType = 3
	PermissionTypeModsOnly     PermissionType = 4

	IconTypeNone    IconType = 0
	IconTypeDefault IconType = 1
	IconTypeCustom  IconType = 2

	IconPositionBefore IconPosition = 0
	IconPositionAfter  IconPosition = 1

	HomepageStreamTypeMuted  HomepageStreamType = 0
	HomepageStreamTypePaused HomepageStreamType = 1
	HomepageStreamTypeHidden HomepageStreamType = 2

	GameSortTypeAlphabetical  GameSortType = 0
	GameSortTypeChronological GameSortType = 1
	GameSortTypeCustom        GameSortType = 2

	TimeDisplayUnitsExplicit TimeDisplayUnits = 0
	TimeDisplayUnitsColon    TimeDisplayUnits = 1

	TimeReferenceAbsolute TimeReference = 0
	TimeReferenceRelative TimeReference = 1

	DateFormatYearMonthDate DateFormat = 0
	DateFormatDateMonthYear DateFormat = 1
	DateFormatMonthDateYear DateFormat = 2

	TimeFormatHourMinute24       TimeFormat = 0
	TimeFormatHourMinuteSecond24 TimeFormat = 1
	TimeFormatHourMinute12       TimeFormat = 2
	TimeFormatHourMinuteSecond12 TimeFormat = 3

	NavbarColorTypePrimary NavbarColorType = 0
	NavbarColorTypePanel   NavbarColorType = 1

	ScrollTypeNone   ScrollType = 0
	ScrollTypeSlow   ScrollType = 1
	ScrollTypeMedium ScrollType = 2
	ScrollTypeFast   ScrollType = 3

	PositionTypeTopLeft     PositionType = 0
	PositionTypeTop         PositionType = 1
	PositionTypeTopRight    PositionType = 2
	PositionTypeLeft        PositionType = 3
	PositionTypeCenter      PositionType = 4
	PositionTypeRight       PositionType = 5
	PositionTypeBottomLeft  PositionType = 6
	PositionTypeBottom      PositionType = 7
	PositionTypeBottomRight PositionType = 8

	RepeatTypeNone       RepeatType = 0
	RepeatTypeHorizontal RepeatType = 1
	RepeatTypeVertical   RepeatType = 2
	RepeatTypeBoth       RepeatType = 3

	FitTypeOriginal FitType = 0
	FitTypeFit      FitType = 1

	SupportPlanPeriodMonthly SupportPlanPeriod = "monthly"
	SupportPlanPeriodYearly  SupportPlanPeriod = "yearly"

	DefaultViewTypeFullGame DefaultViewType = 0
	DefaultViewTypeLevels   DefaultViewType = 1

	ChallengeStateDraft     ChallengeState = 0
	ChallengeStatePublished ChallengeState = 1
	ChallengeStateFinalized ChallengeState = 2
)
