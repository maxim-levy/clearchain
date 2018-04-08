package commands

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/keys"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/spf13/cobra"
)

func GetExportPubCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "export-pub",
		RunE: exportPubCmd,
		Args: cobra.MinimumNArgs(1),
	}
}

func exportPubCmd(cmd *cobra.Command, args []string) error {
	keybase, err := keys.GetKeyBase()
	if err != nil {
		return err
	}
	for _, name := range args {
		key, err := keybase.Get(name)
		if err != nil {
			return fmt.Errorf("couldn't retrieve key by name %q: %v", name, err)
		}
		bz, err := key.PubKey.Wrap().MarshalJSON()
		if err != nil {
			return err
		}
		fmt.Println(string(bz))
	}
	return nil
}
