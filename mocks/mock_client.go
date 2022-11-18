// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Brightscout/mattermost-plugin-azure-devops/server/plugin (interfaces: Client)

// Package mock_plugin is a generated GoMock package.
package mocks

import (
	serializers "github.com/Brightscout/mattermost-plugin-azure-devops/server/serializers"
	gomock "github.com/golang/mock/gomock"
	url "net/url"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CheckIfUserIsProjectAdmin mocks base method
func (m *MockClient) CheckIfUserIsProjectAdmin(arg0, arg1, arg2, arg3 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfUserIsProjectAdmin", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfUserIsProjectAdmin indicates an expected call of CheckIfUserIsProjectAdmin
func (mr *MockClientMockRecorder) CheckIfUserIsProjectAdmin(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfUserIsProjectAdmin", reflect.TypeOf((*MockClient)(nil).CheckIfUserIsProjectAdmin), arg0, arg1, arg2, arg3)
}

// CreateSubscription mocks base method
func (m *MockClient) CreateSubscription(arg0 *serializers.CreateSubscriptionRequestPayload, arg1 *serializers.ProjectDetails, arg2, arg3, arg4 string) (*serializers.SubscriptionValue, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*serializers.SubscriptionValue)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSubscription indicates an expected call of CreateSubscription
func (mr *MockClientMockRecorder) CreateSubscription(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockClient)(nil).CreateSubscription), arg0, arg1, arg2, arg3, arg4)
}

// CreateTask mocks base method
func (m *MockClient) CreateTask(arg0 *serializers.CreateTaskRequestPayload, arg1 string) (*serializers.TaskValue, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(*serializers.TaskValue)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockClientMockRecorder) CreateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockClient)(nil).CreateTask), arg0, arg1)
}

// DeleteSubscription mocks base method
func (m *MockClient) DeleteSubscription(arg0, arg1, arg2 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (mr *MockClientMockRecorder) DeleteSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockClient)(nil).DeleteSubscription), arg0, arg1, arg2)
}

// GenerateOAuthToken mocks base method
func (m *MockClient) GenerateOAuthToken(arg0 url.Values) (*serializers.OAuthSuccessResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateOAuthToken", arg0)
	ret0, _ := ret[0].(*serializers.OAuthSuccessResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateOAuthToken indicates an expected call of GenerateOAuthToken
func (mr *MockClientMockRecorder) GenerateOAuthToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateOAuthToken", reflect.TypeOf((*MockClient)(nil).GenerateOAuthToken), arg0)
}

// GetGitRepositories mocks base method
func (m *MockClient) GetGitRepositories(arg0, arg1, arg2 string) (*serializers.GitRepositoriesResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitRepositories", arg0, arg1, arg2)
	ret0, _ := ret[0].(*serializers.GitRepositoriesResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGitRepositories indicates an expected call of GetGitRepositories
func (mr *MockClientMockRecorder) GetGitRepositories(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitRepositories", reflect.TypeOf((*MockClient)(nil).GetGitRepositories), arg0, arg1, arg2)
}

// GetGitRepositoryBranches mocks base method
func (m *MockClient) GetGitRepositoryBranches(arg0, arg1, arg2, arg3 string) (*serializers.GitBranchesResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitRepositoryBranches", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*serializers.GitBranchesResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGitRepositoryBranches indicates an expected call of GetGitRepositoryBranches
func (mr *MockClientMockRecorder) GetGitRepositoryBranches(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitRepositoryBranches", reflect.TypeOf((*MockClient)(nil).GetGitRepositoryBranches), arg0, arg1, arg2, arg3)
}

// UpdatePipelineApprovalRequest mocks base method
func (m *MockClient) UpdatePipelineApprovalRequest(arg0 *serializers.PipelineApproveRequest, arg1, arg2, arg3 string, arg4 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineApprovalRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[1].(int)
	ret1, _ := ret[2].(error)
	return ret0, ret1
}

// GetPullRequest mocks base method
func (m *MockClient) GetPullRequest(arg0, arg1, arg2, arg3 string) (*serializers.PullRequest, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPullRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*serializers.PullRequest)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPullRequest indicates an expected call of GetPullRequest
func (mr *MockClientMockRecorder) GetPullRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequest", reflect.TypeOf((*MockClient)(nil).GetPullRequest), arg0, arg1, arg2, arg3)
}

// GetTask mocks base method
func (m *MockClient) GetTask(arg0, arg1, arg2, arg3 string) (*serializers.TaskValue, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*serializers.TaskValue)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTask indicates an expected call of GetTask
func (mr *MockClientMockRecorder) GetTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockClient)(nil).GetTask), arg0, arg1, arg2, arg3)
}

// Link mocks base method
func (m *MockClient) Link(arg0 *serializers.LinkRequestPayload, arg1 string) (*serializers.Project, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", arg0, arg1)
	ret0, _ := ret[0].(*serializers.Project)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Link indicates an expected call of Link
func (mr *MockClientMockRecorder) Link(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockClient)(nil).Link), arg0, arg1)
}
