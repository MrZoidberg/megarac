# Megarac

Megarac is a command-line interface (CLI) tool designed to manage BMC operations. It allows users to perform various tasks on servers remotely by interfacing with the BMC.

Please note that Megarac is an independent project and is not affiliated with or endorsed by AMI (American Megatrends Inc.), the company that develops and manufactures BMCs. Any references to AMI or its products are purely for informational purposes.

## Installation

See [Releases page](https://github.com/MrZoidberg/megarac/releases).

## Verified BMCs

Megarac has been tested with the following motherboards:

- Gigabyte MZ31-AR0

Please write an [issue](https://github.com/MrZoidberg/megarac/issues) if you have tested Megarac with a different motherboard and it works.

## Usage

Megarac provides a simple CLI to manage BMC operations. You can specify BMC details directly or use a profile configured with the required details.

`megarac [global options] command [command options] [arguments...]`

### Global options

`--help`, `-h`: Show help.

### Commands

#### `configure`

Configure the profile with the BMC authentication details.
Profiles are stored in the user's home directory in the `.megarac/profiles` directory.

Add/update profile:
`megarac configure add --name [profile_name] --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false]`

List profiles:
`megarac configure list`

Remove profile:
`megarac configure remove --name [profile_name]`

#### `power`

Control the power state of the server.

Power on:
`megarac power on [--profile profile_name]` or `megarac power on --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false]`

Power off:
`megarac power off [--profile profile_name]` or `megarac power off --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false]`

Power status:
`megarac power status [--profile profile_name]` or `megarac power status --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false]`

#### `sensors`

Get sensor readings from the BMC.

List sensors:
`megarac sensors list [--profile profile_name] [-all] [--find name]` or `megarac sensors list --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false] [-all] [--find name]`

### Output format

By default, Megarac outputs data in a human-readable format. You can use the `--format json` flag to output data in JSON format.
Please note that any errors or warnings will be displayed in the human-readable format.

## Contributing

Contributions to Megarac are welcome! Please refer to the CONTRIBUTING.md file for guidelines on how to contribute to this project.

## License

Megarac is released under the MIT License. See the LICENSE file for more details.