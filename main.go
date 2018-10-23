/* sweet
 *
 * Copyright (C) 2018  Franklin "Snaipe" Mathieu <me@snai.pe>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	swaysock := os.Getenv("SWAYSOCK")
	if swaysock != "" {
		out, err := exec.Command("sway", "--get-socketpath").Output()
		if err != nil {
			log.Fatal(err)
		}
		swaysock = strings.TrimSpace(string(out))
	}

	var sockpath string
	if len(os.Args) <= 1 {
		sockpath = swaysock + ".i3"
	} else {
		sockpath = os.Args[1]
	}
	if sockpath == "" {
		log.Fatal("must specify a socket path (either by passing a parameter, or setting SWAYSOCK)")
	}

	os.Remove(sockpath)
	ln, err := net.Listen("unix", sockpath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("export I3SOCK='%s'\n", sockpath)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		sconn, err := net.Dial("unix", swaysock)
		if err != nil {
			log.Print(err)
			conn.Close()
			continue
		}

		go Bridge{sway: sconn, i3: conn}.Run()
	}
}
