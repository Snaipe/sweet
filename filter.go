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

func walk(obj interface{}, fn func(interface{}) (bool, error)) (ok bool, err error) {
	ok, err = fn(obj)
	if !ok || err != nil {
		return
	}

	switch obj := obj.(type) {
	case map[string]interface{}:
		for _, v := range obj {
			ok, err = walk(v, fn)
			if !ok || err != nil {
				break
			}
		}
	case []interface{}:
		for i := range obj {
			ok, err = walk(obj[i], fn)
			if !ok || err != nil {
				break
			}
		}
	}
	return
}

func filterTree(root interface{}) (interface{}, error) {
	walk(root, func(val interface{}) (bool, error) {
		obj, ok := val.(map[string]interface{})
		if !ok {
			return true, nil
		}

		appID, ok := obj["app_id"]
		if !ok {
			return true, nil
		}

		// i3 doesn't have app_id
		delete(obj, "app_id")

		// only perform the fallback for native wayland windows
		if appID != nil {
			if class, ok := obj["class"]; !ok || class == nil {
				obj["class"] = appID
			}

			obj["window_properties"] = map[string]interface{}{
				"class":         obj["class"],
				"instance":      appID,
				"title":         obj["name"],
				"transient_for": nil,
			}
		}

		return true, nil
	})
	return root, nil
}
