package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env      string
	DB       DBConfig
	Auth     AuthConfig
	Server   ServerConfig
	Data     DataConfig
	External ExternalConfig
	App      AppConfig
	Debug    DebugConfig
}

func ReadConfigs(path string) (Config, error) {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

type DBConfig struct {
	Driver string
	URI    string
}

type ServerConfig struct {
	Port    int
	TLSPort int `yaml:"tlsPort"`
	Host    string
	Domain  string
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type (
	ExternalConfig struct {
		Telegram   TelegramConfig
		Sendinblue SendinblueConfig `yaml:"sendinblue"`
	}
	TelegramConfig struct {
		TelegramTo    int64  `yaml:"telegramTo"`
		TelegramToken string `yaml:"telegramToken"`
	}
	SendinblueConfig struct {
		Key string `yaml:"key"`
	}
)

type DebugConfig struct {
	LogDBQueries bool `yaml:"logDBQueries"`
}

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		UserPresetsPaths []string `yaml:"userPresetsPaths"`
		TestPresetsPaths []string `yaml:"testPresetsPaths"`
		TagPresetsPaths  []string `yaml:"tagPresetsPaths"`
	}
)

type (
	AuthConfig struct {
		Secret  string
		Exp     int
		Google  Google
		Github  Github
		Twitter Twitter
	}
	Google struct {
		Client   string
		Secret   string
		Callback string
	}
	Github struct {
		Client   string
		Secret   string
		Callback string
	}
	Twitter struct {
		Client   string
		Secret   string
		Callback string
	}
)

type AppConfig struct {
	SystemSummarySchedule   string `yaml:"systemSummarySchedule"`
	SystemSummaryTimeout    int    `yaml:"systemSummaryTimeout"`
	UpdateNormsSchedule     string `yaml:"updateNormsSchedule"`
	UpdateNormsTimeout      int    `yaml:"updateNormsTimeout"`
	UpdateMarksSchedule     string `yaml:"updateMarksSchedule"`
	UpdateMarksTimeout      int    `yaml:"updateMarksTimeout"`
	UpdateDurationsSchedule string `yaml:"updateDurationsSchedule"`
	UpdateDurationsTimeout  int    `yaml:"updateDurationsTimeout"`
	DefaultTimeLayout       string `yaml:"defaultTimeLayout"`
	Version                 string `yaml:"version"`
}
