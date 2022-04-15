package config

import (
	"reflect"

	"testing"
)

func TestReadConfigFile(t *testing.T) {
	tests := []struct {
		configFile string
		want       *Config
	}{
		{configFile: "test.yaml",
			want: &Config{
				URLs: []string{"https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json",
					"https://raw.githubusercontent.com/assignment132/assignment/main/google.json",
					"https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json"},
			}},
	}
	for _, tt := range tests {
		got, err := ReadConfigFile(tt.configFile)
		if err != nil {
			if !reflect.DeepEqual(got.URLs, tt.want.URLs) {
				t.Errorf("ReadConfigFile() = %v, want %v", got, tt.want)
			}
		}
	}
}
