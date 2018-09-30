package main

import (
	"encoding/json"
	"fmt"

	"github.com/Southclaws/samp-servers-api/types"
	"github.com/rivo/tview"
	"gopkg.in/resty.v1"
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
	servers.SetCellSimple(0, 0, lockedSymbol)
	servers.SetCellSimple(0, 1, "Hostname")
	servers.SetCellSimple(0, 2, "Players")
	servers.SetCellSimple(0, 3, "Gamemode")
	servers.SetCellSimple(0, 4, "Language")
	servers.SetCellSimple(0, 5, "Version")
	servers.SetCellSimple(0, 6, "Address")
	servers.SetFixed(2, 6)

	servers.SetSelectedFunc(func(row, column int) {
		connectForm := tview.NewForm()

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
			//TODO Connect
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
		servers.SetCell(row, 0, lockedColumn)

		hostNameColumn := tview.NewTableCell(server.Hostname)
		servers.SetCell(row, 1, hostNameColumn)

		playersColumn := tview.NewTableCell(fmt.Sprintf("%d/%d", server.Players, server.MaxPlayers))
		servers.SetCell(row, 2, playersColumn)

		gamemodeColumn := tview.NewTableCell(server.Gamemode)
		servers.SetCell(row, 3, gamemodeColumn)

		languageColumn := tview.NewTableCell(server.Language)
		servers.SetCell(row, 4, languageColumn)

		versionColumn := tview.NewTableCell(server.Version)
		servers.SetCell(row, 5, versionColumn)

		addressColumn := tview.NewTableCell(server.Address)
		servers.SetCell(row, 6, addressColumn)
	}

	application.SetRoot(root, true)
	if err := application.Run(); err != nil {
		panic(err)
	}
}
