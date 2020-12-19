<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>

        <el-col :span="24">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="未读消息" name="moreart">


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
                        label="头像"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="昵称">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="消息">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="时间">
                </el-table-column>


                <el-table-column
                        prop="address"
                        label="操作">
                </el-table-column>

              </el-table>

            </el-tab-pane>
           <el-tab-pane label="已读消息" name="moreart">


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
                        label="头像"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="昵称">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="消息">
                </el-table-column>
                <el-table-column
                        prop="address"
                        label="时间">
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
  name: 'Weixinmsgmanager',
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
