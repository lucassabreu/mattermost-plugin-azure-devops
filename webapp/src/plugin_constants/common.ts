export const pluginId = 'mattermost-plugin-azure-devops';

export const AzureDevops = 'Azure DevOps';
export const RightSidebarHeader = 'Azure DevOps';

export const MMCSRF = 'MMCSRF';
export const HeaderCSRFToken = 'X-CSRF-Token';

export const boardsEventTypeMap: Record<ServiceType, Partial<Record<EventType, string>>> = {
    board: {
        create: 'Work Item Created',
        update: 'Work Item Updated',
        delete: 'Work Item Deleted',
        comment: 'Work Item Commented On',
    },
    repos: {
        create: 'Pull Request Created',
        update: 'Pull Request Updated',
        comment: 'Pull Request Commented On',
        merge_attempt: 'Pull Request Merge Attempted',
        code_push: 'Code Pushed',
    }
};

export const defaultPage = 0;
export const defaultPerPageLimit = 10;

export const SubscriptionFilterCreatedBy = {
    me: 'me',
    anyone: 'anyone',
};
