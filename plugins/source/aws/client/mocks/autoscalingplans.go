// Code generated by MockGen. DO NOT EDIT.
// Source: autoscalingplans.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	autoscalingplans "github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	gomock "github.com/golang/mock/gomock"
)

// MockAutoscalingplansClient is a mock of AutoscalingplansClient interface.
type MockAutoscalingplansClient struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalingplansClientMockRecorder
}

// MockAutoscalingplansClientMockRecorder is the mock recorder for MockAutoscalingplansClient.
type MockAutoscalingplansClientMockRecorder struct {
	mock *MockAutoscalingplansClient
}

// NewMockAutoscalingplansClient creates a new mock instance.
func NewMockAutoscalingplansClient(ctrl *gomock.Controller) *MockAutoscalingplansClient {
	mock := &MockAutoscalingplansClient{ctrl: ctrl}
	mock.recorder = &MockAutoscalingplansClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoscalingplansClient) EXPECT() *MockAutoscalingplansClientMockRecorder {
	return m.recorder
}

// DescribeScalingPlanResources mocks base method.
func (m *MockAutoscalingplansClient) DescribeScalingPlanResources(arg0 context.Context, arg1 *autoscalingplans.DescribeScalingPlanResourcesInput, arg2 ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPlanResources", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlanResources indicates an expected call of DescribeScalingPlanResources.
func (mr *MockAutoscalingplansClientMockRecorder) DescribeScalingPlanResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlanResources", reflect.TypeOf((*MockAutoscalingplansClient)(nil).DescribeScalingPlanResources), varargs...)
}

// DescribeScalingPlans mocks base method.
func (m *MockAutoscalingplansClient) DescribeScalingPlans(arg0 context.Context, arg1 *autoscalingplans.DescribeScalingPlansInput, arg2 ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPlans", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlans indicates an expected call of DescribeScalingPlans.
func (mr *MockAutoscalingplansClientMockRecorder) DescribeScalingPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlans", reflect.TypeOf((*MockAutoscalingplansClient)(nil).DescribeScalingPlans), varargs...)
}

// GetScalingPlanResourceForecastData mocks base method.
func (m *MockAutoscalingplansClient) GetScalingPlanResourceForecastData(arg0 context.Context, arg1 *autoscalingplans.GetScalingPlanResourceForecastDataInput, arg2 ...func(*autoscalingplans.Options)) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScalingPlanResourceForecastData", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScalingPlanResourceForecastData indicates an expected call of GetScalingPlanResourceForecastData.
func (mr *MockAutoscalingplansClientMockRecorder) GetScalingPlanResourceForecastData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScalingPlanResourceForecastData", reflect.TypeOf((*MockAutoscalingplansClient)(nil).GetScalingPlanResourceForecastData), varargs...)
}
