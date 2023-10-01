package service

func (s *Service) Complete(id int) (string, error) {
	title, err := s.r.Complete(id)
	if err != nil {
		return title, err
	}

	return title, nil

}
