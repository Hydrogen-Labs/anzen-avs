package example_sf_module

import safetyfactorbase "anzen-avs/safety-factor/safety-factor-base"

type ExampleModule struct {
	*safetyfactorbase.BaseSFModule
}

func NewExampleModule(name string, id int64) *ExampleModule {
	module := &ExampleModule{}
	module.BaseSFModule = safetyfactorbase.NewBaseSFModule(name, id, module)
	return module
}

func (m *ExampleModule) QueryProfitFromCorruption() (float64, error) {
	// Implement specific logic for ExampleModule
	return 150.0, nil
}

func (m *ExampleModule) QueryCostOfCorruption() (float64, error) {
	// Implement specific logic for ExampleModule
	return 250.0, nil
}
