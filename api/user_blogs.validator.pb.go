// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user_blogs.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Blog) Validate() error {
	return nil
}
func (this *FindFeedURLRequest) Validate() error {
	return nil
}
func (this *CreateUserBlogRequest) Validate() error {
	if this.Blog != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Blog); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Blog", err)
		}
	}
	return nil
}
func (this *UpdateUserBlogRequest) Validate() error {
	if this.Blog != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Blog); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Blog", err)
		}
	}
	return nil
}
func (this *DeleteUserBlogRequest) Validate() error {
	return nil
}
