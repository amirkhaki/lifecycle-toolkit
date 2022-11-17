package common

import "sigs.k8s.io/controller-runtime/pkg/client"

//go:generate moq -pkg fake --skip-ensure -out ./fake/listitem_mock.go . ListItem
// represents taskList, evaluationList, workloadInstanceList, appVersionList
type ListItem interface {
	GetItems() []client.Object
}

type ListItemWrapper struct {
	Obj ListItem
}

func NewListItemWrapperFromClientObjectList(object client.ObjectList) (*ListItemWrapper, error) {
	pi, ok := object.(ListItem)
	if !ok {
		return nil, ErrCannotWrapToListItem
	}
	return &ListItemWrapper{Obj: pi}, nil
}

func (pw ListItemWrapper) GetItems() []client.Object {
	return pw.Obj.GetItems()
}
