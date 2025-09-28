package params

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/vincadian/arkh-blockchain/app"
)

// EncodingConfig specifies the encoding configuration for the application.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeTestEncodingConfig creates an EncodingConfig for testing.
func MakeTestEncodingConfig() EncodingConfig {
	encodingConfig := EncodingConfig{
		InterfaceRegistry: types.NewInterfaceRegistry(),
		Marshaler:         codec.NewProtoCodec(types.NewInterfaceRegistry()),
		TxConfig:          authtx.NewTxConfig(codec.NewProtoCodec(types.NewInterfaceRegistry()), authtx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}

	app.ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	app.ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
