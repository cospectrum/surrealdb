// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rixxdb

import (
	"io"

	"github.com/abcum/rixxdb"
	"github.com/abcum/surreal/kvs"
)

type DB struct {
	pntr *rixxdb.DB
}

func (db *DB) Begin(writable bool) (txn kvs.TX, err error) {
	var pntr *rixxdb.TX
	if pntr, err = db.pntr.Begin(writable); err != nil {
		err = &kvs.DBError{Err: err}
		return
	}
	return &TX{pntr: pntr}, err
}

func (db *DB) Import(r io.Reader) (err error) {
	return db.pntr.Load(r)
}

func (db *DB) Export(w io.Writer) (err error) {
	return db.pntr.Save(w)
}

func (db *DB) Close() (err error) {
	return db.pntr.Close()
}
