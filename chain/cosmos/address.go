package cosmos

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// AddressBuilder for Cosmos
type AddressBuilder struct {
}

// NewAddressBuilder creates a new Cosmos AddressBuilder
func NewAddressBuilder(asset xc.AssetConfig) (xc.AddressBuilder, error) {
	return AddressBuilder{}, errors.New("not implemented")
}

// GetAddressFromPublicKey returns an Address given a public key
func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	return xc.Address(""), errors.New("not implemented")
}

// GetAllPossibleAddressesFromPublicKey returns all PossubleAddress(es) given a public key
func (ab AddressBuilder) GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]xc.PossibleAddress, error) {
	address, err := ab.GetAddressFromPublicKey(publicKeyBytes)
	return []xc.PossibleAddress{
		{
			Address: address,
			Type:    xc.AddressTypeDefault,
		},
	}, err
}
