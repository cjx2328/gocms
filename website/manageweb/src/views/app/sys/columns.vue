<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>

        <el-col :span="24">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="栏目列表" name="showdb">

              <el-button type="primary">添加栏目</el-button>

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
                        label="栏目名称【点击查看相关的文章】"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="归属模型">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="导航显示">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="排序">
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
  name: 'Columns',
  data() {
    return {
      sysconfig: {
        site_name: '忠告'
      },
      activeName: 'showdb',
        fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}]

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
