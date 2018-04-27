package business

type Achievment struct {
	Id          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	// Max         int    `json: "max"`
	// Progress    int    `json: "progress"`
}

type Persistence interface {
	read() []Achievment
	save([]Achievment)
}

type AchievmentHandler struct {
	Pers Persistence
}

func (h *AchievmentHandler) Create(achiev Achievment) {
	achievs := h.Pers.read()
	achievs = append(achievs, achiev)
	h.Pers.save(achievs)
}

func (h *AchievmentHandler) Read() []Achievment {
	return h.Pers.read()
}
