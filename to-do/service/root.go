package service

func New() Actions {
	return Actions{
		Get:       get,
		Add:       add,
		Completed: complete,
		Clear:     clear,
		Remove:    remove,
		GetToday:  getToday,
	}
}
