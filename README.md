# Megarac

Megarac is a command-line interface (CLI) tool designed to manage BMC operations. It allows users to perform various tasks on servers remotely by interfacing with the BMC.

Please note that Megarac is an independent project and is not affiliated with or endorsed by AMI (American Megatrends Inc.), the company that develops and manufactures BMCs. Any references to AMI or its products are purely for informational purposes.

## Installation

See [Releases page](https://github.com/MrZoidberg/megarac).

## Usage

Megarac provides a simple CLI to manage BMC operations. You can specify BMC details directly or use a profile configured with the required details.

`megarac [global options] command [command options] [arguments...]`

### Global options

`--use-ssl`: Use SSL to connect to the BMC. Default is true.

`--insecure`: Skip SSL certificate verification. Default is false.

`--help`, `-h`: Show help.

## Contributing

Contributions to Megarac are welcome! Please refer to the CONTRIBUTING.md file for guidelines on how to contribute to this project.

## License

Megarac is released under the MIT License. See the LICENSE file for more details.