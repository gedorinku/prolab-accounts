// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roles.proto

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

func (this *Role) Validate() error {
	return nil
}
func (this *ListRolesRequest) Validate() error {
	return nil
}
func (this *ListRolesResponse) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	return nil
}
func (this *GetRoleRequest) Validate() error {
	return nil
}
