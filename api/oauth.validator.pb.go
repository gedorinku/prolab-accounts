// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oauth.proto

package api_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ProgrammingLab/prolab-accounts/api/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StartOAuthLoginRequest) Validate() error {
	return nil
}
func (this *StartOAuthLoginResponse) Validate() error {
	return nil
}
func (this *OAuthLoginRequest) Validate() error {
	return nil
}
func (this *OAuthLoginResponse) Validate() error {
	return nil
}
func (this *StartOAuthConsentRequest) Validate() error {
	return nil
}
func (this *StartOAuthConsentResponse) Validate() error {
	if this.Client != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Client); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Client", err)
		}
	}
	return nil
}
func (this *OAuthConsentRequest) Validate() error {
	return nil
}
func (this *OAuthConsentResponse) Validate() error {
	return nil
}
