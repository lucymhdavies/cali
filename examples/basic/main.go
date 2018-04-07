package main

import "github.com/skybet/cali"

var (
	cli = cali.NewCli("example")
)

// For simplcity, just one init.
// In reality, for apps with many commands, you would split each command into its own file
func init() {
	cli.SetShort("Example CLI tool")
	cli.SetLong("A nice long description of what your tool actually does")

	cmdTerraform()
}

// cmdTerraform defines a Cali command for running Terraform
func cmdTerraform() {

	// New command
	command := cli.NewCommand("terraform [command]")

	// Short helptext. Shows up with: example --help
	command.SetShort("Run Terraform in an ephemeral container")

	// Long helptext. Shows up with: example terraform --help
	// This is where you'd outline detailed usage instructions
	command.SetLong(`Starts a container for Terraform and attempts to run it against your code. There are two choices for code source; a local mount, or directly from a git repo.

Examples:

  To build the contents of the current working directory using my_account as the AWS profile from the shared credentials file on this host.
  # cali terraform plan -p my_account

  Any addtional flags sent to the terraform command come after the --, e.g.
  # cali terraform plan -- -state=environments/test/terraform.tfstate -var-file=environments/test/terraform.tfvars
  # cali terraform -- plan -out plan.out
`)

	// Flags are pflags:
	// https://godoc.org/github.com/spf13/pflag
	//
	// e.g.
	// example terraform --profile=foo
	// example terraform -p foo
	//
	// In Cali, these are also Viper configs
	// https://godoc.org/github.com/spf13/viper
	//
	// e.g. in .example.yml (in $PWD or $HOME)
	// profile: foo

	command.Flags().StringP("profile", "p", "default", "Profile to use from the AWS shared credentials file")
	command.BindFlags()

	// Task is "the thing the command does"
	// Can either be a function or a string
	// Here we use a string, which Cali interprets as a Docker image
	// The Run function of this task will be essentially: docker run ... hashicorp/terraform:latest
	task := command.Task("hashicorp/terraform:latest")

	// SetInitFunc is optional, and runs before the Run function.
	task.SetInitFunc(func(t *cali.Task, args []string) {
		t.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("profile"))

		// TODO: explicitly bind ~/.aws, if creds aws
		// https://github.com/skybet/cali/issues/7
		// Relies on Cali understanding of Credentials, which I've not got an issue to track my ideas in yet
	})
}

func main() {
	cli.Start()
}
