package cmd

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rebuy-de/aws-nuke/resources"
)

// account / region / queue / item

type Path struct {
	Account  string
	Region   string
	Resource string
}

type ItemState int

const (
	ItemStateNew ItemState = iota
	ItemStatePending
	ItemStateWaiting
	ItemStateFailed
	ItemStateFiltered
	ItemStateFinished
)

type Item struct {
	Resource resources.Resource
	Path     Path
	Session  *session.Session

	State  ItemState
	Reason string
}

func (i *Item) Print() {
	switch i.State {
	case ItemStateNew:
		Log(i.Path, i.Resource, ReasonWaitPending, "would remove")
	case ItemStatePending:
		Log(i.Path, i.Resource, ReasonWaitPending, "triggered remove")
	case ItemStateWaiting:
		Log(i.Path, i.Resource, ReasonWaitPending, "waiting")
	case ItemStateFailed:
		Log(i.Path, i.Resource, ReasonError, i.Reason)
	case ItemStateFiltered:
		Log(i.Path, i.Resource, ReasonSkip, i.Reason)
	case ItemStateFinished:
		Log(i.Path, i.Resource, ReasonSuccess, "removed")
	}
}

func (i *Item) List() ([]resources.Resource, error) {
	listers := resources.GetListers()
	return listers[i.Path.Resource](i.Session)
}

type Queue []*Item

func (q Queue) CountTotal() int {
	return len(q)
}

func (q Queue) Count(states ...ItemState) int {
	count := 0
	for _, item := range q {
		for _, state := range states {
			if item.State == state {
				count = count + 1
				break
			}
		}
	}
	return count
}
