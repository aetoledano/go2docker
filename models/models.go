package models

type DkrConfig struct {
	App struct {
		Name string
	}
	Go struct {
		Version string
	}
}

func (config *DkrConfig) Validate() error {
	return nil
}
