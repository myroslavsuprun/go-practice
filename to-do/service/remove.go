package service

func (s Service) Remove(id int) (string, error) {
	title, err := s.r.Remove(id)
	if err != nil {
		return title, err
	}

	return title, nil
}
