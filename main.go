package YTOpener

import (
	"fmt"
	"os/exec"
	"time"
)

type ChannelID struct {
	ID   string
	NAME string
}

func GetName(channelID *ChannelID) string {
	return channelID.NAME
}

func GetFunctionClear(name string, char string) string {
	return name + " " + char

}

func GetUserReader(channelID *ChannelID) {

	fmt.Println("Enter Name: ")
	_, err := fmt.Scanln(channelID.NAME)

	time.Sleep(1 * time.Second)

	if err == nil {
		fmt.Println("Empty..")
	}

	exec.Command("cmd", GetFunctionClear("/C start", channelID.NAME))

}

func PrintChannelID(message string) {

	_, err := fmt.Println("ID: ", message)

	if err == nil {
		fmt.Println("EMPTY..")
	}
}

func GetChannelName(channelID *ChannelID) string {

	return channelID.ID

}
