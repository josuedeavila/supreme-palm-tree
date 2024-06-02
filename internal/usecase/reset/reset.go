package reset

func (us *useCases) Reset() error {
	if err := us.eventRepository.DeleteAll(); err != nil {
		return err
	}
	if err := us.balaceRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}
