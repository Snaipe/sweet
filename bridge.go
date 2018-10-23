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
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
)

type Filter func(interface{}) (interface{}, error)

type Bridge struct {
	sway, i3 net.Conn
}

func (bridge Bridge) Close() {
	bridge.sway.Close()
	bridge.i3.Close()
}

var filters = map[MessageKind]Filter{
	ResponseTree: filterTree,
}

func (bridge Bridge) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Print(r)
		}
	}()
	defer bridge.Close()

	/* no need to filter inbound commands */
	go io.Copy(bridge.sway, bridge.i3)

	in := bufio.NewReader(bridge.sway)
	for {
		msg, err := ReadMessage(in)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if filter, ok := filters[msg.Kind]; ok {
			var object interface{}
			if err := json.Unmarshal(msg.Data, &object); err != nil {
				log.Print("could not unmarshal json payload:", err, "-- forwarding message.")
				goto forward
			}
			if object, err = filter(object); err != nil {
				panic(err)
			}
			newdata, err := json.Marshal(object)
			if err != nil {
				log.Print("could not marshal filtered json payload:", err, "-- forwarding message.")
				goto forward
			}
			msg.Data = newdata
		}

	forward:
		out := bufio.NewWriter(bridge.i3)
		if err := msg.WriteTo(out); err != nil {
			panic(err)
		}
		out.Flush()
	}
}
