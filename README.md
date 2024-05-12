# ConfigKit

`ConfigKit` is a configuration management library for Go, designed to simplify loading and managing settings from JSON files with the option to override them via environment variables. Using generics, it accommodates any struct type, adapting to your specific configuration needs.

## Key Features

- **Generic Configuration Loader**: Directly load settings into any struct.
- **Environment Overrides**: Prioritize environment variables over file-based settings.
- **Reflective Mapping**: Dynamically associate environment variables with struct fields.

## Installation

Ensure Go is installed on your system, then run:

```bash
go get -u github.com/code-brew-lab/configkit
```

## Usage

Using `ConfigKit` involves three primary steps:

1. Define Your Configuration Struct**: Create a struct that represents the structure of your JSON configuration file.

```go
type Config struct {
    Port int    `json:"port"`
    Host string `json:"host"`
}
```

2. Initialize Settings: Instantiate a new ConfigKit settings object with the path to your JSON configuration file.

```go
settings := configkit.NewSettings[Config]("config.json")
```

 3. Load Configuration: Call the Load method to parse the JSON file and apply any environment variable overrides. This method returns the configuration object and any error encountered.

 ```go
 config, err := settings.Load()
if err != nil {
    log.Fatalf("Failed to load configuration: %s", err)
}
fmt.Printf("Starting server on %s:%d\n", config.Host, config.Port)
```

This sequence initializes the configuration management, loads the settings from a JSON file while overriding with environment variables if they exist, and uses the configuration data in your application.

## Contributing

Contributions to `ConfigKit` are highly appreciated. To contribute:

1. **Fork the Repository**: Start by forking the `ConfigKit` repository on GitHub.
2. **Create a Feature Branch**: Create a new branch in your forked repository for your feature or bug fix.
3. **Commit Your Changes**: Make your modifications in your feature branch, committing changes with clear, descriptive messages.
4. **Push to the Branch**: Push your changes to your repository.
5. **Submit a Pull Request**: Open a pull request to the main `ConfigKit` repository. Include a detailed description of your changes and any other relevant information.

Your contributions will be reviewed as soon as possible by the maintainers.