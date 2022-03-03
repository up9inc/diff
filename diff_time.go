/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package diff

import (
	"reflect"
	"time"
)

func (d *Differ) diffTime(path []string, a, b reflect.Value) error {
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

	// Marshal and unmarshal time type will lose accuracy. Using unix nano to compare time type.
	au := ExportInterface(a).(time.Time).UnixNano()
	bu := ExportInterface(b).(time.Time).UnixNano()

	if au != bu {
		d.cl.Add(UPDATE, path, ExportInterface(a), ExportInterface(b))
	}

	return nil
}
