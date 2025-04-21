package main

import (
	"fmt"
	"github.com/samber/do"
)

type EngineService interface{}

func NewEngineService(i *do.Injector) (EngineService, error) {
	nuts := do.MustInvoke[NutsService](i)
	nuts.NutCount()
	return &engineServiceImplem{}, nil
}

type engineServiceImplem struct{}

// [Optional] Implements do.Healthcheckable.
func (c *engineServiceImplem) HealthCheck() error {
	return fmt.Errorf("engine broken")
}

func NewCarService(i *do.Injector) (*CarService, error) {
	// test out how a engine service would need either nuts or bolts
	engine := do.MustInvoke[EngineService](i)
	car := CarService{Engine: engine}
	return &car, nil
}

type CarService struct {
	Engine EngineService
}

func (c *CarService) Start() {
	println("car starting")
}

// [Optional] Implements do.Shutdownable.
func (c *CarService) Shutdown() error {
	println("car stopped")
	return nil
}

type NutsService interface {
	NutCount()
}

type nutsService struct{}
type premiumNutsService struct{}

func NewNutsService(i *do.Injector) (NutsService, error) {
	fmt.Println("providing nuts")
	return &premiumNutsService{}, nil
}

func (n *nutsService) NutCount() {
	fmt.Println("0 nuts to provide")
}

func (p *premiumNutsService) NutCount() {
	fmt.Println("0 premium nuts to provide the engine")
}

func main() {
	injector := do.New()

	// provides CarService
	do.Provide(injector, NewCarService)

	// provides EngineService
	do.Provide(injector, NewEngineService)

	// provides the NutsService
	do.Provide(injector, NewNutsService)

	car := do.MustInvoke[*CarService](injector)
	car.Start()
	// prints "car starting"

	do.HealthCheck[EngineService](injector)
	// returns "engine broken"

	// injector.ShutdownOnSIGTERM()    // will block until receiving sigterm signal
	injector.Shutdown()
	// prints "car stopped"
}
