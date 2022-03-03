/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package diff

import "reflect"

func (d *Differ) DiffString(path []string, a, b reflect.Value, parent interface{}) error {
	if a.Kind() == reflect.Invalid {
		d.cl.Add(CREATE, path, nil, ExportInterface(b))
		return nil
	}

	if b.Kind() == reflect.Invalid {
		d.cl.Add(DELETE, path, ExportInterface(a), nil)
		return nil
	}

	if a.Kind() != b.Kind() {
		return ErrTypeMismatch
	}

	if a.String() != b.String() {
		if a.CanInterface() {
			// If a and/or b is of a type that is an alias for String, store that type in changelog
			d.cl.Add(UPDATE, path, ExportInterface(a), ExportInterface(b), parent)
		} else {
			d.cl.Add(UPDATE, path, a.String(), b.String(), parent)
		}
	}

	return nil
}
