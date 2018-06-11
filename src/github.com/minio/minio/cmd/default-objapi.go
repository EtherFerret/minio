package cmd

import (
	"context"
	"github.com/minio/minio/pkg/acl"
	"github.com/minio/minio/pkg/lifecycle"
)

type DefaultObjectAPI struct {
}

// AccessControlPolicy
func (fs *DefaultObjectAPI) SetBucketAccessControlPolicy(ctx context.Context, bucket string, aclPolicy *acl.AccessControlPolicy) error {
	return NotImplemented{}
}

func (fs *DefaultObjectAPI) GetBucketAccessControlPolicy(ctx context.Context, bucket string) (*acl.AccessControlPolicy, error) {
	return nil, NotImplemented{}
}

func (fs *DefaultObjectAPI) SetObjectAccessControlPolicy(ctx context.Context, bucket, object string, aclPolicy *acl.AccessControlPolicy) error {
	return NotImplemented{}
}

func (fs *DefaultObjectAPI) GetObjectAccessControlPolicy(ctx context.Context, bucket, object string) (*acl.AccessControlPolicy, error) {
	return nil, NotImplemented{}
}

// Lifecycle
func (fs *DefaultObjectAPI) SetBucketLifecycle(ctx context.Context, bucket string, lifecycle *lifecycle.Lifecycle) error {
	return NotImplemented{}
}

func (fs *DefaultObjectAPI) GetBucketLifecycle(ctx context.Context, bucket string) (*lifecycle.Lifecycle, error) {
	return nil, NotImplemented{}
}

func (fs *DefaultObjectAPI) DeleteBucketLifecycle(ctx context.Context, bucket string) error {
	return NotImplemented{}
}

func (fs *DefaultObjectAPI) RestoreObject(context.Context, string, string, *lifecycle.RestoreRequest) error {
	return NotImplemented{}
}
