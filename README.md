
A CLI tool for interacting with DeepLake tensor database.
Complete documentation is available at https://docs.activeloop.ai

Usage:
  godeep [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  create-dataset Create a new dataset
  help           Help about any command

Flags:
      --config string         config file (default is ./config/default.yaml)
      --dataset-path string   Dataset path
  -h, --help                  help for godeep
      --org-id string         ActiveLoop organization ID
      --token string          ActiveLoop API token
  -v, --version               version for godeep

Use "godeep [command] --help" for more information about a command.


# Using environment variables
export ACTIVELOOP_TOKEN="your-token"
export ACTIVELOOP_ORG_ID="your-org"
export ACTIVELOOP_DATASET_PATH="hub://org/dataset"

# Create a dataset
./godeep create-dataset

# Using Docker
docker-compose up

# Using default config
./godeep create-dataset

# Using custom config
./godeep --config custom-config.yaml create-dataset

# Override config values
./godeep create-dataset --path "hub://custom/path" --tensor-db=true

# Using environment variables
export ACTIVELOOP_TOKEN="your-token"
export ACTIVELOOP_ORG_ID="your-org"


./godeep create-dataset

docker-compose down                    

docker-compose up --build
