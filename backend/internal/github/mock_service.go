// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package github is a generated GoMock package.
package github

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v43/github"
	oauth2 "golang.org/x/oauth2"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockService) DeleteFile(ctx context.Context, ghToken oauth2.Token, fileProps GitFileProps) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, ghToken, fileProps)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockServiceMockRecorder) DeleteFile(ctx, ghToken, fileProps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockService)(nil).DeleteFile), ctx, ghToken, fileProps)
}

// GetAuthCodeURL mocks base method.
func (m *MockService) GetAuthCodeURL(state string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthCodeURL", state)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuthCodeURL indicates an expected call of GetAuthCodeURL.
func (mr *MockServiceMockRecorder) GetAuthCodeURL(state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthCodeURL", reflect.TypeOf((*MockService)(nil).GetAuthCodeURL), state)
}

// GetFile mocks base method.
func (m *MockService) GetFile(ctx context.Context, ghToken oauth2.Token, fileProps GitFileProps) (GitFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, ghToken, fileProps)
	ret0, _ := ret[0].(GitFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockServiceMockRecorder) GetFile(ctx, ghToken, fileProps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockService)(nil).GetFile), ctx, ghToken, fileProps)
}

// GetRepos mocks base method.
func (m *MockService) GetRepos(ctx context.Context, ghToken oauth2.Token) ([]GitRepo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepos", ctx, ghToken)
	ret0, _ := ret[0].([]GitRepo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepos indicates an expected call of GetRepos.
func (mr *MockServiceMockRecorder) GetRepos(ctx, ghToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepos", reflect.TypeOf((*MockService)(nil).GetRepos), ctx, ghToken)
}

// GetToken mocks base method.
func (m *MockService) GetToken(ctx context.Context, code string) (oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, code)
	ret0, _ := ret[0].(oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockServiceMockRecorder) GetToken(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockService)(nil).GetToken), ctx, code)
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, ghToken oauth2.Token) (github.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, ghToken)
	ret0, _ := ret[0].(github.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, ghToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, ghToken)
}

// SaveFile mocks base method.
func (m *MockService) SaveFile(ctx context.Context, ghToken oauth2.Token, fileProps GitFileProps) (GitFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", ctx, ghToken, fileProps)
	ret0, _ := ret[0].(GitFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockServiceMockRecorder) SaveFile(ctx, ghToken, fileProps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockService)(nil).SaveFile), ctx, ghToken, fileProps)
}

// SearchFiles mocks base method.
func (m *MockService) SearchFiles(ctx context.Context, ghToken oauth2.Token, fileProps GitFileProps, query string, pageNo int) ([]GitFile, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFiles", ctx, ghToken, fileProps, query, pageNo)
	ret0, _ := ret[0].([]GitFile)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchFiles indicates an expected call of SearchFiles.
func (mr *MockServiceMockRecorder) SearchFiles(ctx, ghToken, fileProps, query, pageNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFiles", reflect.TypeOf((*MockService)(nil).SearchFiles), ctx, ghToken, fileProps, query, pageNo)
}
