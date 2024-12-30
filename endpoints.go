package main

import "strings"

type endpoint struct {
	Route  string
	Prefix string
}

func (e *endpoint) URL() string {
	return e.Prefix + e.Route
}

func (e *endpoint) URLWithParams(params map[string]string) string {
	url := e.Prefix + e.Route
	for k, v := range params {
		url = strings.Replace(url, "${"+k+"}", v, -1)
	}
	return url
}

var endpoints = map[string]endpoint{
	"activity": {
		Route:  "activity",
		Prefix: "/connect/transfers/",
	},
	"authenticate": {
		Route:  "authenticate",
		Prefix: "/connect/info/",
	},
	"droppedFiles": {
		Route:  "dropped-files",
		Prefix: "/connect/file/",
	},
	"getTransfer": {
		Route:  "info/${id}",
		Prefix: "/connect/transfers/",
	},
	"initDragDrop": {
		Route:  "initialize-drag-drop",
		Prefix: "/connect/file/",
	},
	"modifyTransfer": {
		Route:  "modify/${id}",
		Prefix: "/connect/transfers/",
	},
	"ping": {
		Route:  "ping",
		Prefix: "/connect/info/",
	},
	"readAsArrayBuffer": {
		Route:  "read-as-array-buffer/",
		Prefix: "/connect/file/",
	},
	"readChunkAsArrayBuffer": {
		Route:  "read-chunk-as-array-buffer/",
		Prefix: "/connect/file/",
	},
	"getChecksum": {
		Route:  "checksum/",
		Prefix: "/connect/file/",
	},
	"removeTransfer": {
		Route:  "remove/${id}",
		Prefix: "/connect/transfers/",
	},
	"resumeTransfer": {
		Route:  "resume/${id}",
		Prefix: "/connect/transfers/",
	},
	"showAbout": {
		Route:  "about",
		Prefix: "/connect/windows/",
	},
	"showDirectory": {
		Route:  "finder/${id}",
		Prefix: "/connect/windows/",
	},
	"showPreferences": {
		Route:  "preferences",
		Prefix: "/connect/windows/",
	},
	"showPreferencesPage": {
		Route:  "preferences/${id}",
		Prefix: "/connect/windows/",
	},
	"showSaveFileDialog": {
		Route:  "select-save-file-dialog/",
		Prefix: "/connect/windows/",
	},
	"showSelectFileDialog": {
		Route:  "select-open-file-dialog/",
		Prefix: "/connect/windows/",
	},
	"showSelectFolderDialog": {
		Route:  "select-open-folder-dialog/",
		Prefix: "/connect/windows/",
	},
	"showTransferManager": {
		Route:  "transfer-manager",
		Prefix: "/connect/windows/",
	},
	"showTransferMonitor": {
		Route:  "transfer-monitor/${id}",
		Prefix: "/connect/windows/",
	},
	"startTransfer": {
		Route:  "start",
		Prefix: "/connect/transfers/",
	},
	"stopTransfer": {
		Route:  "stop/${id}",
		Prefix: "/connect/transfers/",
	},
	"testSshPorts": {
		Route:  "ports",
		Prefix: "/connect/info/",
	},
	"version": {
		Route:  "version",
		Prefix: "/connect/info/",
	},
}
