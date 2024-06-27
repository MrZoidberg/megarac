package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Profile struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`

	UseSSL    *bool `json:"use_ssl"`
	InsureSsl *bool `json:"insecure_ssl"`
}

type Config struct {
	Profiles []*Profile
}

func newConfig() *Config {
	return &Config{}
}

var Cfg *Config

func Load() error {
	Cfg = newConfig()
	return Cfg.Load()
}

func (c *Config) AddProfile(profile *Profile) {
	c.Profiles = append(c.Profiles, profile)
}

func (c *Config) RemoveProfile(name string) {
	for i, p := range c.Profiles {
		if p.Name == name {
			c.Profiles = append(c.Profiles[:i], c.Profiles[i+1:]...)
			break
		}
	}
}

func (c *Config) GetProfile(name string) *Profile {
	for _, p := range c.Profiles {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// Save saves the config to disk to ~/.megarac/profiles/*.json
// Each profile is saved as a separate json file
// The directory structure is created if it does not exist
func (c *Config) Save() error {
	// Create the directory if it does not exist
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profilesPath := fmt.Sprintf("%s/.megarac/profiles", dirname)
	err = os.MkdirAll(profilesPath, 0770)
	if err != nil {
		return err
	}

	for _, profile := range c.Profiles {
		// Convert the profile to JSON
		profileJSON, err := json.Marshal(profile)
		if err != nil {
			return err
		}

		// Generate the file path
		filePath := fmt.Sprintf("%s/%s.json", profilesPath, profile.Name)

		// Write the JSON to the file with 0644 permissions
		err = os.WriteFile(filePath, profileJSON, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Load loads the config from disk from ~/.megarac/profiles/*.json
// Each profile is loaded from a separate json file
func (c *Config) Load() error {
	// Check if the directory exists
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profilesPath := fmt.Sprintf("%s/.megarac/profiles", dirname)
	_, err = os.Stat(profilesPath)
	if os.IsNotExist(err) {
		return nil
	}

	// Read the directory
	dir, err := os.ReadDir(profilesPath)
	if err != nil {
		return err
	}

	for _, file := range dir {
		// Read the file
		filePath := fmt.Sprintf("%s/%s", profilesPath, file.Name())
		profileJSON, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// Unmarshal the JSON into a Profile struct
		profile := &Profile{}
		err = json.Unmarshal(profileJSON, profile)
		if err != nil {
			return err
		}

		// Add the profile to the config
		c.AddProfile(profile)
	}

	return nil
}
