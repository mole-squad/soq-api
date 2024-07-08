package selectinput

type SelectListOption struct {
	opt SelectOption
}

func (s SelectListOption) Label() string {
	return s.opt.Label()
}

func (s SelectListOption) Description() string {
	return ""
}

func (s SelectListOption) FilterValue() string {
	return s.opt.Label()
}
