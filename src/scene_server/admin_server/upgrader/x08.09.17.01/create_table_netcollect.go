/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x08_09_17_01

import (
	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage"

	"gopkg.in/mgo.v2"
)

func createTable(db storage.DI, conf *upgrader.Config) (err error) {
	for tablename, indexs := range tables {
		exists, err := db.HasTable(tablename)
		if err != nil {
			return err
		}
		if !exists {
			if err = db.CreateTable(tablename); err != nil && !mgo.IsDup(err) {
				return err
			}
		}
		for index := range indexs {
			if err = db.Index(tablename, &indexs[index]); err != nil && !mgo.IsDup(err) {
				return err
			}
		}
	}
	return nil
}

var tables = map[string][]storage.Index{
	common.BKTableNameNetcollectDevice: []storage.Index{
		storage.Index{Name: "", Columns: []string{"device_id"}, Type: storage.INDEX_TYPE_BACKGROUP},
		storage.Index{Name: "", Columns: []string{"device_name"}, Type: storage.INDEX_TYPE_BACKGROUP},
		storage.Index{Name: "", Columns: []string{"bk_supplier_account"}, Type: storage.INDEX_TYPE_BACKGROUP},
	},

	common.BKTableNameNetcollectProperty: []storage.Index{
		storage.Index{Name: "", Columns: []string{"netcollect_property_id"}, Type: storage.INDEX_TYPE_BACKGROUP},
		storage.Index{Name: "", Columns: []string{"bk_supplier_account"}, Type: storage.INDEX_TYPE_BACKGROUP},
	}}