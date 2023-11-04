package main

import (
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"log"
)

var (
	omitempty       bool
	omitemptyPrefix string
)

var rootCmd = &cobra.Command{
	Use:     "protoc-gen-go-fiber",
	Short:   "Generate Go code for Fiber from Protocol Buffers",
	Version: release,
	Run: func(cmd *cobra.Command, args []string) {
		protogen.Options{
			ParamFunc: cmd.Flags().Set,
		}.Run(func(gen *protogen.Plugin) error {
			gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
			for _, f := range gen.Files {
				if !f.Generate {
					continue
				}
				generateFile(gen, f, omitempty, omitemptyPrefix)
			}
			return nil
		})
	},
}

func init() {
	rootCmd.Flags().BoolVar(&omitempty, "omitempty", true, "omit if google.api is empty")
	rootCmd.Flags().StringVar(&omitemptyPrefix, "omitempty_prefix", "", "omit if google.api is empty")
	rootCmd.AddCommand(addTagCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
