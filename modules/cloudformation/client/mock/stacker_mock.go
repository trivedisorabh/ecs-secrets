// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/awslabs/ecs-secrets/modules/cloudformation/client (interfaces: Stacker)

package mock_client

import (
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Stacker interface
type MockStacker struct {
	ctrl     *gomock.Controller
	recorder *_MockStackerRecorder
}

// Recorder for MockStacker (not exported)
type _MockStackerRecorder struct {
	mock *MockStacker
}

func NewMockStacker(ctrl *gomock.Controller) *MockStacker {
	mock := &MockStacker{ctrl: ctrl}
	mock.recorder = &_MockStackerRecorder{mock}
	return mock
}

func (_m *MockStacker) EXPECT() *_MockStackerRecorder {
	return _m.recorder
}

func (_m *MockStacker) CreateStack(_param0 string, _param1 string, _param2 string) (*cloudformation.Stack, error) {
	ret := _m.ctrl.Call(_m, "CreateStack", _param0, _param1, _param2)
	ret0, _ := ret[0].(*cloudformation.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStackerRecorder) CreateStack(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateStack", arg0, arg1, arg2)
}

func (_m *MockStacker) DeleteStack(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteStack", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStackerRecorder) DeleteStack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStack", arg0)
}
