# desloader
Desloader is a bulk entity loader for GCP Datastore working with YAML formatted schema and CSV formatted data source.

## Usage
```bash
 $ desloader --help
Usage:
  desloader [OPTIONS]

Application Options:
  --projectId= Target GCP project ID
  --schema=    A schema definition in YAML format
  --source=    A source file path in CSV format
  --kind=      Target kind on Datastore

Help Options:
  -h, --help       Show this help message
```

## Development
```bash
$ make
```

## Contribution
PRs accepted

## License
MIT © IzumiSy
