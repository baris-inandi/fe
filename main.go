// Copyright (C) 2022 baris-inandi
//
// This file is part of fe.
//
// fe is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// fe is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with fe.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/baris-inandi/fe/fecli"
)

func main() {
	err := fecli.Cli.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
