// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// S3Client is an autogenerated mock type for the S3Client type
type S3Client struct {
	mock.Mock
}

// Get provides a mock function with given fields: bucketName, objectName, fileName
func (_m *S3Client) Get(bucketName string, objectName string, fileName string) error {
	ret := _m.Called(bucketName, objectName, fileName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(bucketName, objectName, fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllObjectKeys provides a mock function with given fields: bucketName
func (_m *S3Client) GetAllObjectKeys(bucketName string) []string {
	ret := _m.Called(bucketName)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(bucketName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Put provides a mock function with given fields: bucketName, objectName, filePath
func (_m *S3Client) Put(bucketName string, objectName string, filePath string) (int64, error) {
	ret := _m.Called(bucketName, objectName, filePath)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string) int64); ok {
		r0 = rf(bucketName, objectName, filePath)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(bucketName, objectName, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewS3Client interface {
	mock.TestingT
	Cleanup(func())
}

// NewS3Client creates a new instance of S3Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewS3Client(t mockConstructorTestingTNewS3Client) *S3Client {
	mock := &S3Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
