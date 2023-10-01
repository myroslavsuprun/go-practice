package service

func (s *Service) Clear() error {
	err := s.r.Clear()
	if err != nil {
		return err
	}
	return nil
}
