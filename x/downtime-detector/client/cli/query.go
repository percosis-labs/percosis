package cli

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/percosis-labs/percosis/v16/x/downtime-detector/client/queryproto"
	"github.com/percosis-labs/percosis/v16/x/downtime-detector/types"
)

func GetQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, RecoveredSinceQueryCmd)

	return cmd
}

func RecoveredSinceQueryCmd() (*percocli.QueryDescriptor, *queryproto.RecoveredSinceDowntimeOfLengthRequest) {
	return &percocli.QueryDescriptor{
		Use:   "recovered-since downtime-duration recovery-duration",
		Short: "Queries if it has been at least <recovery-duration> since the chain was down for <downtime-duration>",
		Long: `{{.Short}}
downtime-duration is a duration, but is restricted to a smaller set. Heres a few from the set: 30s, 1m, 5m, 10m, 30m, 1h, 3 h, 6h, 12h, 24h, 36h, 48h]
{{.ExampleHeader}}
{{.CommandPrefix}} recovered-since 24h 30m`,
		CustomFieldParsers: map[string]percocli.CustomFieldParserFn{"Downtime": parseDowntimeDuration},
	}, &queryproto.RecoveredSinceDowntimeOfLengthRequest{}
}

func parseDowntimeDuration(arg string, _ *pflag.FlagSet) (any, percocli.FieldReadLocation, error) {
	dur, err := time.ParseDuration(arg)
	if err != nil {
		return nil, percocli.UsedArg, err
	}
	downtime, err := types.DowntimeByDuration(dur)
	return downtime, percocli.UsedArg, err
}
