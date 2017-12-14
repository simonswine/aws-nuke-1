package resources

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2Subnet struct {
	svc *ec2.EC2
	id  *string
}

func ListEC2Subnets(sess *session.Session) ([]Resource, error) {
	params := &ec2.DescribeSubnetsInput{}
	resp, err := svc.DescribeSubnets(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.Subnets {
		resources = append(resources, &EC2Subnet{
			svc: svc,
			id:  out.SubnetId,
		})
	}

	return resources, nil
}

func (e *EC2Subnet) Remove() error {
	params := &ec2.DeleteSubnetInput{
		SubnetId: e.id,
	}

	_, err := e.svc.DeleteSubnet(params)
	if err != nil {
		return err
	}

	return nil
}

func (e *EC2Subnet) String() string {
	return *e.id
}
