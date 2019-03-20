package to

import (
	"fmt"
	"github.com/spf13/cobra"
)

type KafkaCmd struct {
	*cobra.Command
	toTopic   string
	toBrokers []string
}

func KafkaCommand() *KafkaCmd {
	kafkaCmd := &KafkaCmd{}
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "send events using kafka",
		Long:  `to kafka Long Description goes here.....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside to.KafkaCmd Run with args: %v and %s and b : %v\n", args, kafkaCmd.toTopic, kafkaCmd.toBrokers)
		},
		//PreRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Printf("Inside to.KafkaCmd PreRun with args: %v and %s\n", args, flags.Author)
		//},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringVarP(&kafkaCmd.toTopic, "to-topic", "", "", "to kafka topic")
	cmd.PersistentFlags().StringSliceVarP(&kafkaCmd.toBrokers, "to-brokers", "", nil, "to brokers hosts comma separated")
	kafkaCmd.Command = cmd

	return kafkaCmd
}
