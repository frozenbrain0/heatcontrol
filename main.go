package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	default:
		printHelp()
		os.Exit(1)
	case "list-profiles":
		listProfiles()
	case "set-profile":
		setAcpiProfile()
	}
}

func printHelp() {
	fmt.Println("heatcontrol v1.1")
	fmt.Println()
	fmt.Println("   list-profiles          - list all performance-profiles")
	fmt.Println("   set-profile <profile>  - set performance-profil")
	fmt.Println()
}

func listProfiles() { //Reads and formats /sys/firmware/acpi/platform_profile_choices and prints the results
	platformProfilesBytes, err := os.ReadFile("/sys/firmware/acpi/platform_profile_choices")
	if err != nil {
		fmt.Println("Error reading file /sys/firmware/acpi/platform_profile_choices: ", err)
		os.Exit(2)
	}
	platformProfiles := string(platformProfilesBytes)
	platformProfiles = strings.Replace(platformProfiles, " ", "\n   ", -1)
	fmt.Println("Available profiles:\n  ", platformProfiles)
}

func setAcpiProfile() { //validates the given profile and writes the string to /sys/firmware/acpi/platform_profile
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}
	platformProfilesBytes, err := os.ReadFile("/sys/firmware/acpi/platform_profile_choices")
	if err != nil {
		fmt.Println("Error reading file /sys/firmware/acpi/platform_profile_choices: ", err)
		os.Exit(2)
	}
	platformProfilesString := string(platformProfilesBytes)
	platformProfilesString = strings.ReplaceAll(platformProfilesString, "\n", "") // Remove new line for check
	platformProfiles := strings.Split(platformProfilesString, " ")
	for c := range platformProfiles {
		if os.Args[2] == platformProfiles[c] {

			msg, success := setProfile(os.Args[2])
			fmt.Println(msg)
			if success {
				os.Exit(0)
			} else {
				os.Exit(3)
			}

		}
	}
	fmt.Println("Error: profile not found!")

}

func setProfile(profile string) (string, bool) { //run the command to set the profile and check the exit code for errors
	wri := exec.Command("run0", "--unit=heatcontrol-change-profile", "bash", "-c", "echo "+profile+" > /sys/firmware/acpi/platform_profile")
	wri.Stdout = os.Stdout
	wri.Stderr = os.Stderr
	wri.Run()
	wri.Wait()
	if wri.ProcessState.ExitCode() != 0 {
		return "An error occurred! Try running \"run0 bash -c 'echo " + profile + " > /sys/firmware/acpi/platform_profile'\" manually.", false
	}
	return "Performance-profile set to " + profile, true
}
