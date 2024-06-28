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

Example:

```text
> go run . power --profile endor status
Power status for endor-bmc.lan: on
```

#### `sensors`

Get sensor readings from the BMC.

List sensors:
`megarac sensors list [--profile profile_name] [-all] [--find name]` or `megarac sensors list --host [hostname] --user [username] --password [password] [--insecure] [--use-ssl false] [-all] [--find name]`

Example:

```text
> go run . sensors --profile endor list --all
ID Name           Type                   Reading                Alert State
1  CPU0_TEMP      temperature            35 deg_c             active
2  DIMMG0_TEMP    temperature            32 deg_c             active
3  DIMMG1_TEMP    temperature            33 deg_c             active
4  MB_TEMP1       temperature            32 deg_c             active
5  MB_TEMP2       temperature            33 deg_c             active
6  CPU0_DTS       temperature            65 deg_c             active
7  VR_P0_TEMP     temperature            30 deg_c             active
8  VR_DIMMG0_TEMP temperature            33 deg_c             active
9  VR_DIMMG1_TEMP temperature            32 deg_c             active
10 M2_G0_AMB_TEMP temperature            37 deg_c             active
13 P_12V          voltage                12.09 volts          active
14 P_5V           voltage                5.0629 volts         active
15 P_3V3          voltage                3.3043 volts         active
16 P_5V_STBY      voltage                5.0629 volts         active
17 SOC_VDDCR      voltage                0.602 volts          active
18 P_VBAT         voltage                3.0528 volts         active
19 P0_VDDCR_CPU   voltage                0.553 volts          active
20 P0_VDD_18      voltage                1.8326 volts         active
21 P_1V0_AUX_LAN  voltage                0.987 volts          active
22 VR_P0_VOUT     voltage                0.784 volts          active
23 VR_DIMMG0_VOUT voltage                1.248 volts          active
24 VR_DIMMG1_VOUT voltage                1.264 volts          active
25 CPU0_FAN       fan                    1650 rpm             active
27 SYS_FAN2       fan                    900 rpm              active
28 SYS_FAN3       fan                    4800 rpm             active
29 SYS_FAN4       fan                    1200 rpm             active
31 SEL            event_logging_disabled 32768 unknown        inactive
32 CPU0_Status    processor              32896 unknown        inactive
33 PS1_Status     power_supply           32768 unknown        inactive
34 PS2_Status     power_supply           32768 unknown        inactive
36 Watchdog       watchdog_2             32768 Percent        inactive
```

### Output format

By default, Megarac outputs data in a human-readable format. You can use the `--format json` flag to output data in JSON format.
Please note that any errors or warnings will be displayed in the human-readable format.

## Contributing

Contributions to Megarac are welcome!

## License

Megarac is released under the MIT License. See the LICENSE file for more details.