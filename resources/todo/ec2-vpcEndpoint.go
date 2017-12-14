package resources

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2VPCEndpoint struct {
	svc *ec2.EC2
	id  *string
}

func init() {
	register("EC2VPCEndpoint", ListEC2VPCEndpoints)
}

func ListEC2VPCEndpoints(sess *session.Session) ([]Resource, error) {
	resp, err := svc.DescribeVpcEndpoints(nil)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, vpcEndpoint := range resp.VpcEndpoints {
		resources = append(resources, &EC2VPCEndpoint{
			svc: svc,
			id:  vpcEndpoint.VpcEndpointId,
		})
	}

	return resources, nil
}

func (endpoint *EC2VPCEndpoint) Remove() error {
	params := &ec2.DeleteVpcEndpointsInput{
		VpcEndpointIds: []*string{endpoint.id},
	}

	_, err := endpoint.svc.DeleteVpcEndpoints(params)
	if err != nil {
		return err
	}

	return nil
}

func (endpoint *EC2VPCEndpoint) String() string {
	return *endpoint.id
}
