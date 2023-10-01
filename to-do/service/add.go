package service

func (s *Service) Add(title string) error {
	err := s.r.Add(title)
	if err != nil {
		return err
	}
	return nil

}
