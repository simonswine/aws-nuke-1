package resources

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2Address struct {
	svc *ec2.EC2
	id  string
	ip  string
}

func init() {
	register("EC2Addresse", ListEC2Addresses)
}

func ListEC2Addresses(sess *session.Session) ([]Resource, error) {
	params := &ec2.DescribeAddressesInput{}
	resp, err := svc.DescribeAddresses(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.Addresses {
		resources = append(resources, &EC2Address{
			svc: svc,
			id:  *out.AllocationId,
			ip:  *out.PublicIp,
		})
	}

	return resources, nil
}

func (e *EC2Address) Remove() error {
	_, err := e.svc.ReleaseAddress(&ec2.ReleaseAddressInput{
		AllocationId: &e.id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *EC2Address) String() string {
	return e.ip
}
