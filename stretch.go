package pdf

type stretch struct {
	container
}

func (s *stretch) messure(available size) sizePlan {
	m := s.container.messure(available)

	if m.pType == wrap {
		return m
	}

	return sizePlan{pType: m.pType, size: available}
}
