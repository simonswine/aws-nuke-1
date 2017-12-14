package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDSInstance struct {
	svc *rds.RDS
	id  string
}

func (n *RDSNuke) ListInstances() ([]Resource, error) {
	params := &rds.DescribeDBInstancesInput{}
	resp, err := n.Service.DescribeDBInstances(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, instance := range resp.DBInstances {
		resources = append(resources, &RDSInstance{
			svc: n.Service,
			id:  *instance.DBInstanceIdentifier,
		})
	}

	return resources, nil
}

func (i *RDSInstance) Remove() error {
	params := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: &i.id,
		SkipFinalSnapshot:    aws.Bool(true),
	}

	_, err := i.svc.DeleteDBInstance(params)
	if err != nil {
		return err
	}

	return nil
}

func (i *RDSInstance) String() string {
	return i.id
}
