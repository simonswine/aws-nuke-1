package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func ListCloudWatchEventsRules(sess *session.Session) ([]Resource, error) {
	resp, err := n.Service.ListRules(nil)
	if err != nil {
		return nil, err
	}
	resources := make([]Resource, 0)
	for _, rule := range resp.Rules {
		resources = append(resources, &CloudWatchEventsRule{
			svc:  n.Service,
			name: rule.Name,
		})

	}
	return resources, nil
}

type CloudWatchEventsRule struct {
	svc  *cloudwatchevents.CloudWatchEvents
	name *string
}

func (rule *CloudWatchEventsRule) Remove() error {
	_, err := rule.svc.DeleteRule(&cloudwatchevents.DeleteRuleInput{
		Name: rule.name,
	})
	return err
}

func (rule *CloudWatchEventsRule) String() string {
	return fmt.Sprintf("Rule: %s", *rule.name)
}
