package from

import (
	"fmt"
	"github.com/fractalpal/eventflow/ef/cmd/to"
	"github.com/spf13/cobra"
)

type KafkaCmd struct {
	*cobra.Command
	fromBrokers []string
	fromTopic   string
}

func KafkaCommand() *KafkaCmd {
	kafkaCmd := &KafkaCmd{}
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "subscribe for events using kafka",
		Long:  `kafka from Long Description goes here.....`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.KafkaCmd PersistentPreRun with args: %v and %s and b: %v\n", args, kafkaCmd.fromTopic, kafkaCmd.fromBrokers)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.KafkaCmd PersistentPostRun with args: %v and %s and b: %v\n", args, kafkaCmd.fromTopic, kafkaCmd.fromBrokers)
		},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringSliceVarP(&kafkaCmd.fromBrokers, "from-brokers", "", []string{}, "from brokers hosts")
	cmd.PersistentFlags().StringVarP(&kafkaCmd.fromTopic, "from-topic", "", "", "from topic")
	cmd.AddCommand(to.Cmd())
	kafkaCmd.Command = cmd

	return kafkaCmd
}
