<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>

        <el-col :span="24">
          <el-button type="primary" @click="onSubmit">添加模板</el-button>

          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="模板管理" name="showdb">


              <el-button type="primary" @click="onSubmit">HTML</el-button>
              <el-button type="primary" @click="onSubmit">CSS</el-button>
              <el-button type="primary" @click="onSubmit">JS</el-button>
              <el-button type="primary" @click="onSubmit">媒体文件管理</el-button>

              <el-table
                      :data="tableData"
                      style="width: 100%">
                <el-table-column
                        prop="date"
                        label="文件名称"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="name"
                        label="文件大小"
                        width="180">
                </el-table-column>


                <el-table-column
                        prop="address"
                        label="修改日期">
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
  name: 'Templates',
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
