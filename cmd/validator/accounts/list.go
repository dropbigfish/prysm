package accounts

import (
	"strings"

	"github.com/prysmaticlabs/prysm/v5/cmd"
	"github.com/prysmaticlabs/prysm/v5/cmd/validator/flags"
	"github.com/prysmaticlabs/prysm/v5/validator/accounts"
	"github.com/prysmaticlabs/prysm/v5/validator/client"
	"github.com/urfave/cli/v2"
)

func accountsList(c *cli.Context) error {
	w, km, err := walletWithKeymanager(c)
	if err != nil {
		return err
	}
	dialOpts := client.ConstructDialOptions(
		c.Int(cmd.GrpcMaxCallRecvMsgSizeFlag.Name),
		c.String(flags.CertFlag.Name),
		c.Uint(flags.GRPCRetriesFlag.Name),
		c.Duration(flags.GRPCRetryDelayFlag.Name),
	)
	grpcHeaders := strings.Split(c.String(flags.GRPCHeadersFlag.Name), ",")

	opts := []accounts.Option{
		accounts.WithWallet(w),
		accounts.WithKeymanager(km),
		accounts.WithGRPCDialOpts(dialOpts),
		accounts.WithBeaconRPCProvider(c.String(flags.BeaconRPCProviderFlag.Name)),
		accounts.WithBeaconRESTApiProvider(c.String(flags.BeaconRESTApiProviderFlag.Name)),
		accounts.WithGRPCHeaders(grpcHeaders),
	}
	if c.IsSet(flags.ShowPrivateKeysFlag.Name) {
		opts = append(opts, accounts.WithShowPrivateKeys())
	}
	if c.IsSet(flags.ListValidatorIndices.Name) {
		opts = append(opts, accounts.WithListValidatorIndices())
	}
	acc, err := accounts.NewCLIManager(opts...)
	if err != nil {
		return err
	}
	return acc.List(
		c.Context,
	)
}
