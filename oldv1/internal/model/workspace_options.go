// Copyright 2020 Roger Chapman (rogchap). All Rights Reserved.

package model

import "github.com/therecipe/qt/core"

// WorkspaceOptions is a QObject that holds the options for the Workspace
type WorkspaceOptions struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string      `property:"addr"`
	_ bool        `property:"reflect"`
	_ bool        `property:"insecure"`
	_ bool        `property:"plaintext"`
	_ string      `property:"rootca"`
	_ string      `property:"clientcert"`
	_ string      `property:"clientkey"`
	_ *StringList `property:"protoListModel"`
	_ *StringList `property:"importListModel"`

	_ func() `slot:"clearProtoList"`
	_ func() `slot:"clearImportList"`
}

func (m *WorkspaceOptions) init() {
	m.SetReflect(false)
	m.SetPlaintext(false)
	m.SetInsecure(false)
	m.SetProtoListModel(NewStringList(nil))
	m.SetImportListModel(NewStringList(nil))

	m.ConnectClearProtoList(m.clearProtoList)
	m.ConnectClearImportList(m.clearImportList)
}

func (m *WorkspaceOptions) clearProtoList() {
	m.ProtoListModel().SetStringList(nil)
}

func (m *WorkspaceOptions) clearImportList() {
	m.ImportListModel().SetStringList(nil)
}
