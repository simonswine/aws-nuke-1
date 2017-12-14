package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type IAMUserPolicyAttachment struct {
	svc        *iam.IAM
	policyArn  string
	policyName string
	roleName   string
}

func init() {
	register("IAMUserPolicyAttachment", ListIAMUserPolicyAttachments)
}

func ListIAMUserPolicyAttachments(sess *session.Session) ([]Resource, error) {
	resp, err := svc.ListUsers(nil)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, role := range resp.Users {
		resp, err := svc.ListAttachedUserPolicies(
			&iam.ListAttachedUserPoliciesInput{
				UserName: role.UserName,
			})
		if err != nil {
			return nil, err
		}

		for _, pol := range resp.AttachedPolicies {
			resources = append(resources, &IAMUserPolicyAttachment{
				svc:        svc,
				policyArn:  *pol.PolicyArn,
				policyName: *pol.PolicyName,
				roleName:   *role.UserName,
			})
		}
	}

	return resources, nil
}

func (e *IAMUserPolicyAttachment) Remove() error {
	_, err := e.svc.DetachUserPolicy(
		&iam.DetachUserPolicyInput{
			PolicyArn: &e.policyArn,
			UserName:  &e.roleName,
		})
	if err != nil {
		return err
	}

	return nil
}

func (e *IAMUserPolicyAttachment) String() string {
	return fmt.Sprintf("%s -> %s", e.roleName, e.policyName)
}
