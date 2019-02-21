// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/ProgrammingLab/prolab-accounts/api/type"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	if this.Department != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Department); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Department", err)
		}
	}
	return nil
}
func (this *ListUsersRequest) Validate() error {
	return nil
}
func (this *ListUsersResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *GetUserRequest) Validate() error {
	return nil
}
func (this *CreateUserRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *GetCurrentUserRequest) Validate() error {
	return nil
}
func (this *UpdateUserProfileRequest) Validate() error {
	if !(len(this.FullName) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("FullName", fmt.Errorf(`value '%v' must length be less than '128'`, this.FullName))
	}
	if !(len(this.Description) < 1024) {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must length be less than '1024'`, this.Description))
	}
	if !(this.Grade > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Grade", fmt.Errorf(`value '%v' must be greater than '0'`, this.Grade))
	}
	if !(this.Grade < 6) {
		return github_com_mwitkow_go_proto_validators.FieldError("Grade", fmt.Errorf(`value '%v' must be less than '6'`, this.Grade))
	}
	if !(this.RoleId > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleId", fmt.Errorf(`value '%v' must be greater than '-1'`, this.RoleId))
	}
	if !(len(this.TwitterScreenName) < 16) {
		return github_com_mwitkow_go_proto_validators.FieldError("TwitterScreenName", fmt.Errorf(`value '%v' must length be less than '16'`, this.TwitterScreenName))
	}
	if !(len(this.GithubUserName) < 40) {
		return github_com_mwitkow_go_proto_validators.FieldError("GithubUserName", fmt.Errorf(`value '%v' must length be less than '40'`, this.GithubUserName))
	}
	if !(this.DepartmentId > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("DepartmentId", fmt.Errorf(`value '%v' must be greater than '-1'`, this.DepartmentId))
	}
	if _, ok := ProfileScope_name[int32(this.ProfileScope)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("ProfileScope", fmt.Errorf(`value '%v' must be a valid ProfileScope field`, this.ProfileScope))
	}
	return nil
}
func (this *UpdatePasswordRequest) Validate() error {
	return nil
}
