package resources

import "github.com/aws/aws-sdk-go/service/iam"

type IAMInstanceProfile struct {
	svc  *iam.IAM
	name string
}

func (n *IAMNuke) ListInstanceProfiles() ([]Resource, error) {
	resp, err := n.Service.ListInstanceProfiles(nil)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.InstanceProfiles {
		resources = append(resources, &IAMInstanceProfile{
			svc:  n.Service,
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
