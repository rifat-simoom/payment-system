package card_info

import (
	"github.com/pkg/errors"
)

type CardInfo struct {
	cardNo string
}

type FactoryConfig struct {
}

func (f FactoryConfig) Validate() error {
	var err error

	return err
}

type Factory struct {
	// it's better to keep FactoryConfig as a private attribute,
	// thanks to that we are always sure that our configuration is not changed in the not allowed way
	fc FactoryConfig
}

func NewFactory(fc FactoryConfig) (Factory, error) {
	if err := fc.Validate(); err != nil {
		return Factory{}, errors.Wrap(err, "invalid config passed to factory")
	}

	return Factory{fc: fc}, nil
}

func MustNewFactory(fc FactoryConfig) Factory {
	f, err := NewFactory(fc)
	if err != nil {
		panic(err)
	}

	return f
}

func (f Factory) Config() FactoryConfig {
	return f.fc
}

func (f Factory) IsZero() bool {
	return f == Factory{}
}
