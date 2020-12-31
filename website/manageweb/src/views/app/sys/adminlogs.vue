<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>

        <el-col :span="24">

          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="管理员日志" name="showdb">

              <el-form :inline="true"  class="demo-form-inline">
                <el-form-item label="搜索">
                  <el-input placeholder="请输入关键字"></el-input>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary" @click="onSubmit">查询</el-button>
                </el-form-item>
              </el-form>

              <el-table
                      :data="tableData"
                      style="width: 100%">
                <el-table-column
                        prop="date"
                        label="编号"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="name"
                        label="用户"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="操作内容">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="访问地址">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="IP">
                </el-table-column>

                <el-table-column
                        prop="address"
                        label="浏览器">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="操作时间">
                </el-table-column>

                <el-table-column
                        prop="address"
                        label="操作">
                </el-table-column>

              </el-table>

            </el-tab-pane>

          </el-tabs>
        </el-col>

      </el-row>

    </div>
  </div>
</template>
<script>
import { requestAll as requestAllMenu, requestMenuButton } from '@/api/app/sys/menu'
import { requestList, requestDetail, requestUpdate, requestCreate, requestDelete, requestSetRole, requestRoleMenuIDList, requestAll } from '@/api/app/sys/role'
import waves from '@/directive/waves'
import Pagination from '@/components/Pagination'
import SelectTree from '@/components/TreeSelect'
import { checkAuthAdd, checkAuthDel, checkAuthView, checkAuthUpdate, checkAuthSetrolemenu } from '@/utils/permission'

export default {
  name: 'Adminlogs',
  data() {
    return {

      activeName: 'showdb'

        , value1: true,
        value2: true
    }
  },
    methods: {
        handleRemove(file, fileList) {
            console.log(file, fileList);
        },
        handlePreview(file) {
            console.log(file);
        }
    }
}

</script>
