package ui

type SystemModel struct {
	CPUUsage float64
	RAMUsage float64
	isLoading bool
	Quitting bool
}

func NewSystemModel() SystemModel {
  return SystemModel{
		CPUUsage: 0.0,
		RAMUsage: 0.0,
		isLoading: true,
		Quitting: false,
	}
}
