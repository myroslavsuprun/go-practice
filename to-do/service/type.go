package service

type Actions struct {
	Get       Get
	Add       Add
	Completed Complete
	Clear     Clear
	Remove    Remove
	GetToday  GetToday
}

type Clear = func() error

type Get = func() ([]TodoShow, error)

type Add = func(title string) error

type Complete = func(id int) (string, error)

type Remove = func(id int) (string, error)

type GetToday = func() ([]TodoShow, error)
