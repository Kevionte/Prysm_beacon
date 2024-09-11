package iface

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1/consensus-types/validator"
	"github.com/pkg/errors"
)

var ErrNotSupported = errors.New("endpoint not supported")

type ValidatorCount struct {
	Status string
	Count  uint64
}

// PrysmBeaconChainClient defines an interface required to implement all the prysm specific custom endpoints.
type PrysmBeaconChainClient interface {
	GetValidatorCount(context.Context, string, []validator.Status) ([]ValidatorCount, error)
}
