package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Bios-Marcel/samp-launcher-terminal/internal/samp"
	"github.com/Southclaws/samp-servers-api/types"
	"github.com/rivo/tview"
	"gopkg.in/resty.v1"
)

const (
	columnLockedIndex = iota
	columnHostnameIndex
	columnPlayersIndex
	columnGamemodeIndex
	columnLanguageIndex
	columnVersionIndex
	columnAddressIndex
)

func main() {
	application := tview.NewApplication()

	flex := tview.NewFlex()
	root := tview.NewFrame(flex)
	root.SetBorder(true)
	root.SetTitle("SA-MP Launcher")
	root.SetBorderPadding(-1, -1, 0, 0)

	servers := tview.NewTable()
	servers.SetBorder(true)
	servers.SetSelectable(true, false)

	const lockedSymbol = "🔒"
	servers.SetCellSimple(0, columnLockedIndex, lockedSymbol)
	servers.SetCellSimple(0, columnHostnameIndex, "Hostname")
	servers.SetCellSimple(0, columnPlayersIndex, "Players")
	servers.SetCellSimple(0, columnGamemodeIndex, "Gamemode")
	servers.SetCellSimple(0, columnLanguageIndex, "Language")
	servers.SetCellSimple(0, columnVersionIndex, "Version")
	servers.SetCellSimple(0, columnAddressIndex, "Address")
	servers.SetFixed(2, 6)

	servers.SetSelectedFunc(func(row, column int) {
		selectedServerHostname := servers.GetCell(row, columnHostnameIndex).Text
		selectedServerAddress := servers.GetCell(row, columnAddressIndex).Text
		selectedServerPort := "7777"

		partsOfAddress := strings.Split(selectedServerAddress, ":")
		selectedServerAddress = partsOfAddress[0]
		if len(partsOfAddress) == 2 {
			selectedServerPort = partsOfAddress[1]
		}

		connectForm := tview.NewForm()
		connectForm.SetBorder(true)
		connectForm.SetTitle(fmt.Sprintf("Connecting to %s (%s:%s)", selectedServerHostname, selectedServerAddress, selectedServerPort))

		nameToUse := "" //TODO Retrieve name
		serverPasswordToUse := ""
		rconPasswordToUse := ""

		connectForm.AddInputField("Name", "PLACEHOLDER", 0,
			func(textToCheck string, lastChar rune) bool {
				//TODO Name validation
				return true
			}, func(text string) {
				nameToUse = text
			})

		connectForm.AddPasswordField("RCON Password", "", 0, '*', func(text string) {
			rconPasswordToUse = text
		})

		connectForm.AddPasswordField("Server Password", "", 0, '*', func(text string) {
			serverPasswordToUse = text
		})

		connectForm.AddButton("Connect", func() {
			connectionArguments := fmt.Sprintf("-h %s -p %s -n %s", selectedServerAddress, selectedServerPort, nameToUse)
			if len(serverPasswordToUse) != 0 {
				connectionArguments = fmt.Sprintf("%s -z %s", connectionArguments, serverPasswordToUse)
			}
			if len(rconPasswordToUse) != 0 {
				connectionArguments = fmt.Sprintf("%s -c %s", connectionArguments, rconPasswordToUse)
			}
			samp.LaunchSAMP(connectionArguments)

			application.SetRoot(flex, true)
		})

		connectForm.AddButton("Cancel", func() {
			//Do nothing
			application.SetRoot(flex, true)
		})

		application.SetRoot(connectForm, true)
	})

	flex.SetDirection(tview.FlexColumn)
	flex.AddItem(servers, 0, 1, true)

	response, requestError := resty.SetDebug(false).R().Get("https://api.samp-servers.net/v2/servers")
	if requestError != nil {
		dialog := tview.NewModal()
		dialog.SetText(fmt.Sprintf("Error loading servers (%s).", requestError.Error()))
		application.SetRoot(dialog, true)
		application.Run()
		return
	}

	var serverList []types.ServerCore
	json.Unmarshal([]byte(response.String()), &serverList)
	row := 1
	for _, server := range serverList {
		row = row + 1

		var lockedColumn *tview.TableCell
		if server.Password {
			lockedColumn = tview.NewTableCell(lockedSymbol)
		} else {
			lockedColumn = tview.NewTableCell("")
		}
		servers.SetCell(row, columnLockedIndex, lockedColumn)

		hostNameColumn := tview.NewTableCell(server.Hostname)
		servers.SetCell(row, columnHostnameIndex, hostNameColumn)

		playersColumn := tview.NewTableCell(fmt.Sprintf("%d/%d", server.Players, server.MaxPlayers))
		servers.SetCell(row, columnPlayersIndex, playersColumn)

		gamemodeColumn := tview.NewTableCell(server.Gamemode)
		servers.SetCell(row, columnGamemodeIndex, gamemodeColumn)

		languageColumn := tview.NewTableCell(server.Language)
		servers.SetCell(row, columnLanguageIndex, languageColumn)

		versionColumn := tview.NewTableCell(server.Version)
		servers.SetCell(row, columnVersionIndex, versionColumn)

		addressColumn := tview.NewTableCell(server.Address)
		servers.SetCell(row, columnAddressIndex, addressColumn)
	}

	application.SetRoot(root, true)
	if err := application.Run(); err != nil {
		panic(err)
	}
}
