package s3

import (
	"errors"
	"github.com/JesseleDuran/secure-graph-worker/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Setenv("AWS_BUCKET_NAME", "test_bucket")
	os.Exit(m.Run())
}

func Test_Download(t *testing.T) {
	t.Run("Download successful", func(t *testing.T) {
		mockClient := mocks.S3Client{}
		mockClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		s3 := NewS3Manager(&mockClient)
		exp := "downloads/test1"
		got, err := s3.Download("test2", "test1")

		assert.Nil(t, err)
		assert.Equal(t, exp, got)
	})
	t.Run("Download error", func(t *testing.T) {
		mockClient := mocks.S3Client{}
		mockClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("File not found"))
		s3 := NewS3Manager(&mockClient)
		exp := ""
		errExp := "[S3Manager:Download][s3 downloading][err: File not found]"
		got, err := s3.Download("test2", "test1")

		assert.NotNil(t, err)
		assert.Equal(t, errExp, err.Error())
		assert.Equal(t, exp, got)
	})
}

func Test_Upload(t *testing.T) {
	t.Run("Upload successful", func(t *testing.T) {
		mockClient := mocks.S3Client{}
		mockClient.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)
		s3 := NewS3Manager(&mockClient)
		err := s3.Upload("test2", "test1")
		assert.Nil(t, err)
	})
	t.Run("Upload error", func(t *testing.T) {
		mockClient := mocks.S3Client{}
		mockClient.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("Unable to upload"))
		s3 := NewS3Manager(&mockClient)
		errExp := "[S3Manager:Upload][s3 uploading][err: Unable to upload]"
		err := s3.Upload("test2", "test1")
		assert.NotNil(t, err)
		assert.Equal(t, errExp, err.Error())
	})
}
