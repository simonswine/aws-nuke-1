package resources

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2InternetGateway struct {
	svc *ec2.EC2
	id  *string
}

func ListEC2InternetGateways(sess *session.Session) ([]Resource, error) {
	resp, err := n.Service.DescribeInternetGateways(nil)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.InternetGateways {
		resources = append(resources, &EC2InternetGateway{
			svc: n.Service,
			id:  out.InternetGatewayId,
		})
	}

	return resources, nil
}

func (e *EC2InternetGateway) Remove() error {
	params := &ec2.DeleteInternetGatewayInput{
		InternetGatewayId: e.id,
	}

	_, err := e.svc.DeleteInternetGateway(params)
	if err != nil {
		return err
	}

	return nil
}

func (e *EC2InternetGateway) String() string {
	return *e.id
}
