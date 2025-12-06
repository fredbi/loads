// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package loads

type formatKind uint8

const (
	formatKindJSON formatKind = iota
	formatKindYAML
)

func preprocessXOrder(kind formatKind, data []byte) ([]byte, error) {
	return nil, nil // TODO
}
