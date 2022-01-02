package main

import "cobraExample/cmd"
//Cobra Explanation: https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

func main() {
	//to add commands, cd to cobraExample Folder
	//type cobra add <command name>
	//cobra will create the <command name>.go in the cmd folder
	//after you will need to change the innards to handle the new command
	//once complete, to have the changes take effect run:
	//go install cobraExample


	//to make a command a sub command:
	//	in the init() of the command new command file:
	//		change: rootCmd.AddCommand(<command name>Cmd) to <parent>Cmd.AddCommand(<command name>Cmd)
	//do the same for flags
	cmd.Execute()
}
