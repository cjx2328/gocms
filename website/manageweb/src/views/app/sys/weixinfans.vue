<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>
        <el-button type="primary">同步选中粉丝信息</el-button>
        <el-button type="primary">同步全部粉丝</el-button>
        <el-col :span="24">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="粉丝管理" name="showdb">

              <el-form :inline="true"  class="demo-form-inline">
                <el-form-item label="搜索">
                  <el-input placeholder="关键字"></el-input>
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
                        label="昵称"
                        width="180">
                </el-table-column>


                <el-table-column
                        prop="address"
                        label="关注时间">
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
  name: 'Weixinfans',
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
