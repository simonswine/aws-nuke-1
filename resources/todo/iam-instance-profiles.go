package resources

import "github.com/aws/aws-sdk-go/service/iam"

type IAMInstanceProfile struct {
	svc  *iam.IAM
	name string
}

func init() {
	register("IAMInstanceProfile", ListIAMInstanceProfiles)
}

func ListIAMInstanceProfiles(sess *session.Session) ([]Resource, error) {
	resp, err := svc.ListInstanceProfiles(nil)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.InstanceProfiles {
		resources = append(resources, &IAMInstanceProfile{
			svc:  svc,
			name: *out.InstanceProfileName,
		})
	}

	return resources, nil
}

func (e *IAMInstanceProfile) Remove() error {
	_, err := e.svc.DeleteInstanceProfile(&iam.DeleteInstanceProfileInput{
		InstanceProfileName: &e.name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *IAMInstanceProfile) String() string {
	return e.name
}
