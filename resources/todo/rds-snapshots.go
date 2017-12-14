package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDSSnapshot struct {
	svc        *rds.RDS
	identifier *string
	status     *string
}

func ListRDSSnapshots(sess *session.Session) ([]Resource, error) {
	params := &rds.DescribeDBSnapshotsInput{MaxRecords: aws.Int64(100)}
	resp, err := n.Service.DescribeDBSnapshots(params)
	if err != nil {
		return nil, err
	}
	var resources []Resource
	for _, snapshot := range resp.DBSnapshots {
		resources = append(resources, &RDSSnapshot{
			svc:        n.Service,
			identifier: snapshot.DBSnapshotIdentifier,
			status:     snapshot.Status,
		})

	}

	return resources, nil
}
func (i *RDSSnapshot) Remove() error {
	params := &rds.DeleteDBSnapshotInput{
		DBSnapshotIdentifier: i.identifier,
	}

	_, err := i.svc.DeleteDBSnapshot(params)
	if err != nil {
		return err
	}

	return nil
}

func (i *RDSSnapshot) String() string {
	return *i.identifier
}
