# Heatcontrol

This program uses the platform_profiles at /sys/firmware/acpi to allow the user to control the behavior of fans easily.

__NOTE: This might not be possible on your system if the files__ `/sys/firmware/acpi/platform_profile` __and__ `/sys/firmware/acpi/platform_profile_choices` __don't exist.__

### Installation

__This project requires systemd version >=256 because it uses run0 for gaining root permissions.__

Download one of the binaries from the releases, rename it to heatctl (or whatever you want) and put it into /usr/bin.

Alternatively, you can install Go and compile the project yourself by running

`$ git clone https://github.com/frozenbrain0/heatcontrol; cd heatcontrol; go build`

### Usage
`$ heatctl list-profiles` - list available profiles

`$ heatctl set-profile <profile name>` - set the profile

Note: The profile is reset to default on reboot

### Passwordless usage
If you want to use this without having to type in your password every time, you can set the SETUID bit for the file after you copied it to the desired location:

`$ sudo chown root:root <path to binary>` - set the owner to root or some user that has write access to the platform_profile file in sysfs

`$ sudo chmod u+s <path to binary>` - set the SETUID bit
