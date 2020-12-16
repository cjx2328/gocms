<template>
  <div class="app-container">
    <div class="filter-container">
      <el-row>
        <el-col :span="24">

          <el-row>
            <el-col :span="12">
              数据库有45张表，共计24K
            </el-col>
            <el-col span="4">
              <el-button>备份</el-button>
              <el-button type="primary">优化</el-button>

            </el-col>
          </el-row>
        </el-col>
        </el-row>
      <el-row>

        <el-col :span="24">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="数据库表" name="showdb">

              <el-table
                      :data="tableData"
                      style="width: 100%">
                <el-table-column
                        prop="date"
                        label="数据库表"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="name"
                        label="记录条数"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="占用空间">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="类型">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="编码">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="创建时间">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="说明">
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
  name: 'Dbbackup',
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
