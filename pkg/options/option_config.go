package options

import (
	"encoding/json"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"k8sproxy/pkg/types"
	"k8sproxy/pkg/util"
	"os"
	"reflect"
	"unsafe"
)

var clientConfig *types.ClientConfig

type OptionConfig struct {
	Target       string
	DefaultValue any
	Description  string
}

func SetOptions(cmd *cobra.Command, flags *flag.FlagSet, optionStore any, config []OptionConfig) {
	cmd.Long = cmd.Short
	cmd.Flags().SortFlags = false
	cmd.InheritedFlags().SortFlags = false
	flags.SortFlags = false
	for _, c := range config {
		name := util.UnCapitalize(c.Target)
		field := reflect.ValueOf(optionStore).Elem().FieldByName(c.Target)
		switch c.DefaultValue.(type) {
		case string:
			fieldPtr := (*string)(unsafe.Pointer(field.UnsafeAddr()))
			defaultValue := c.DefaultValue.(string)
			if field.String() != "" {
				defaultValue = field.String()
			}
			flags.StringVar(fieldPtr, name, defaultValue, c.Description)

		case int:
			defaultValue := c.DefaultValue.(int)
			if field.Int() != 0 {
				defaultValue = int(field.Int())
			}
			fieldPtr := (*int)(unsafe.Pointer(field.UnsafeAddr()))
			flags.IntVar(fieldPtr, name, defaultValue, c.Description)
		case bool:
			defaultValue := c.DefaultValue.(bool)
			if field.Bool() {
				defaultValue = field.Bool()
			}
			fieldPtr := (*bool)(unsafe.Pointer(field.UnsafeAddr()))

			flags.BoolVar(fieldPtr, name, defaultValue, c.Description)

		}
	}
}

func LoadClientConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &clientConfig)
}

func getBaseURL() string {
	if clientConfig != nil && clientConfig.BaseURL != "" {
		return clientConfig.BaseURL
	}
	return "http://127.0.0.1:8080"
}
