package constants

const (
	// Bot configs
	BotUsername    = "azuredevops"
	BotDisplayName = "Azure DevOps"
	BotDescription = "A bot account created by the Azure DevOps plugin."

	// Plugin configs
	PluginID               = "mattermost-plugin-azure-devops"
	ChannelID              = "channel_id"
	HeaderMattermostUserID = "Mattermost-User-ID"

	// Command configs
	CommandTriggerName = "azuredevops"
	HelpText           = "###### Mattermost Azure DevOps Plugin - Slash Command Help\n" +
		"* `/azuredevops connect` - Connect your Mattermost account to your Azure DevOps account.\n" +
		"* `/azuredevops disconnect` - Disconnect your Mattermost account from your Azure DevOps account.\n" +
		"* `/azuredevops link [projectURL]` - Link your project to a current channel.\n" +
		"* `/azuredevops boards create [title] [description]` - Create a new task for your project.\n" +
		"* `/azuredevops boards subscription add` - Add a new Boards subscription for your linked projects.\n" +
		"* `/azuredevops boards subscription list [me or anyone] [all_channels]` - View Boards subscriptions.\n" +
		"* `/azuredevops boards subscription delete [subscription id]` - Unsubscribe a Boards subscription"
	InvalidCommand      = "Invalid command parameters. Please use `/azuredevops help` for more information."
	CommandHelp         = "help"
	CommandConnect      = "connect"
	CommandDisconnect   = "disconnect"
	CommandLink         = "link"
	CommandBoards       = "boards"
	CommandRepos        = "repos"
	CommandCreate       = "create"
	CommandSubscription = "subscription"
	CommandAdd          = "add"
	CommandList         = "list"
	CommandDelete       = "delete"

	// Regex to verify task link
	TaskLinkRegex = `http(s)?:\/\/dev.azure.com\/[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*\/[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*\/_workitems\/edit\/[1-9]+`

	// Regex to verify pull request link
	PullRequestLinkRegex = `http(s)?:\/\/dev.azure.com\/[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*\/[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*\/_git\/[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*\/pullrequest\/[1-9]+`

	// Azure API Versions
	CreateTaskAPIVersion = "7.1-preview.3"
	TasksIDAPIVersion    = "5.1"
	TasksAPIVersion      = "6.0"

	// Subscription constants
	PublisherID                           = "tfs"
	ConsumerID                            = "webHooks"
	ConsumerActionID                      = "httpRequest"
	SubscriptionEventPullRequestCreated   = "git.pullrequest.created"
	SubscriptionEventPullRequestUpdated   = "git.pullrequest.updated"
	SubscriptionEventPullRequestCommented = "ms.vss-code.git-pullrequest-comment-event"
	SubscriptionEventPullRequestMerged    = "git.pullrequest.merged"
	SubscriptionEventCodePushed           = "git.push"
	SubscriptionEventWorkItemCreated      = "workitem.created"
	SubscriptionEventWorkItemUpdated      = "workitem.updated"
	SubscriptionEventWorkItemDeleted      = "workitem.deleted"
	SubscriptionEventWorkItemCommented    = "workitem.commented"

	// Path params
	PathParamTeamID = "team_id"

	// URL query params constants
	QueryParamProject   = "project"
	QueryParamChannelID = "channel_id"
	QueryParamCreatedBy = "created_by"
	QueryParamPage      = "page"
	QueryParamPerPage   = "per_page"

	// Filters
	FilterCreatedByMe     = "me"
	FilterCreatedByAnyone = "anyone"
	FilterAllChannels     = "all_channels"

	DefaultPage         = 0
	DefaultPerPageLimit = 50

	// Authorization constants
	Bearer        = "Bearer"
	Authorization = "Authorization"

	GetTasksID  = "/%s/_apis/wit/wiql"
	GetTasks    = "/%s/_apis/wit/workitems"
	StaticFiles = "%s/plugins/%s/static/%s"

	PageQueryParam       = "$top"
	APIVersionQueryParam = "api-version"
	IDsQueryParam        = "ids"

	// Websocket events
	WSEventConnect             = "connect"
	WSEventDisconnect          = "disconnect"
	WSEventSubscriptionDeleted = "subscription_deleted"

	// Colors
	ReposIconColor  = "#d74f27"
	BoardsIconColor = "#53bba1"

	SubscriptionEventTypeDummy = "dummy"
	FileNameGitBranchIcon      = "git-branch-icon.svg"
	FileNameProjectIcon        = "project-icon.svg"
	FileNameReposIcon          = "repos-icon.svg"
	FileNameBoardsIcon         = "boards-icon.svg"
	IconColorBoards            = "#53bba1"
	IconColorRepos             = "#d74f27"

	SlackAttachmentAuthorNameRepos  = "Azure Repos"
	SlackAttachmentAuthorNameBoards = "Azure Boards"
)
