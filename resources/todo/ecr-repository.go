package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type ECRRepository struct {
	svc  *ecr.ECR
	name *string
}

func ListECRRepos(sess *session.Session) ([]Resource, error) {
	var params *ecr.DescribeRepositoriesInput
	var resp *ecr.DescribeRepositoriesOutput
	var resources []Resource
	var err error
	for moreRepositories := true; moreRepositories; {
		if resp == nil {
			params = &ecr.DescribeRepositoriesInput{
				MaxResults: aws.Int64(100),
			}
			moreRepositories = true
		} else {
			if resp.NextToken != nil {
				fmt.Printf("Next token\n")
				params = &ecr.DescribeRepositoriesInput{
					MaxResults: aws.Int64(100),
					NextToken:  resp.NextToken,
				}
				moreRepositories = true
			} else {
				moreRepositories = false
				continue
			}
		}
		resp, err = n.Service.DescribeRepositories(params)
		if err != nil {
			fmt.Printf("Error occured")
			return nil, err
		}
		for _, repository := range resp.Repositories {
			resources = append(resources, &ECRRepository{
				svc:  n.Service,
				name: repository.RepositoryName,
			})
		}
	}

	return resources, nil
}

func (r *ECRRepository) Filter() error {
	return nil
}

func (r *ECRRepository) Remove() error {

	params := &ecr.DeleteRepositoryInput{
		RepositoryName: r.name,
		Force:          aws.Bool(true),
	}
	_, err := r.svc.DeleteRepository(params)
	return err
}

func (r *ECRRepository) String() string {
	return fmt.Sprintf("Repository: %s", *r.name)
}
