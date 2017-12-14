package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2NATGateway struct {
	svc   *ec2.EC2
	id    string
	state string
}

func (n *EC2Nuke) ListNATGateways() ([]Resource, error) {
	params := &ec2.DescribeNatGatewaysInput{}
	resp, err := n.Service.DescribeNatGateways(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.NatGateways {
		resources = append(resources, &EC2NATGateway{
			svc:   n.Service,
			id:    *out.NatGatewayId,
			state: *out.State,
		})
	}

	return resources, nil
}

func (n *EC2NATGateway) Filter() error {
	if n.state == "deleted" {
		return fmt.Errorf("already deleted")
	}
	return nil
}

func (n *EC2NATGateway) Remove() error {
	params := &ec2.DeleteNatGatewayInput{
		NatGatewayId: &n.id,
	}

	_, err := n.svc.DeleteNatGateway(params)
	if err != nil {
		return err
	}

	return nil
}

func (n *EC2NATGateway) String() string {
	return n.id
}
