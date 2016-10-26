// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ddl_test

import (
	"os"
	"testing"

	"github.com/ngaut/log"
	. "github.com/pingcap/check"
	"github.com/insionng/zenpress/libraries/pingcap/tidb"
	"github.com/insionng/zenpress/libraries/pingcap/tidb/ast"
	"github.com/insionng/zenpress/libraries/pingcap/tidb/kv"
)

func TestT(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testSuite{})

type testSuite struct {
	store       kv.Storage
	charsetInfo *ast.CharsetOpt
}

func (ts *testSuite) SetUpSuite(c *C) {
	store, err := tidb.NewStore(tidb.EngineGoLevelDBMemory)
	c.Assert(err, IsNil)
	ts.store = store
	ts.charsetInfo = &ast.CharsetOpt{
		Chs: "utf8",
		Col: "utf8_bin",
	}

}

func init() {
	logLevel := os.Getenv("log_level")
	log.SetLevelByString(logLevel)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}
