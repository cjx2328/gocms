<template>
  <div class="app-container">
    <div class="filter-container">

      <el-row>
        <el-button type="primary">添加关键词</el-button>

        <el-col :span="24">
          <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="关键词回复" name="replykeyword">

              <el-tabs  v-model="activeName" @tab-click="handleClick">
                <el-tab-pane label="回复文字" name="replytexts">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>
                <el-tab-pane label="回复图文" name="replypictexs">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>
                <el-tab-pane label="回复语音" name="replysound">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>
                <el-tab-pane label="回复图片" name="replypic">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>
                <el-tab-pane label="回复视频" name="replyvideos">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>
                <el-tab-pane label="回复音乐" name="replymusics">


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
                            label="关键词"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="回复内容">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="状态">
                    </el-table-column>


                    <el-table-column
                            prop="address"
                            label="操作">
                    </el-table-column>

                  </el-table>

                </el-tab-pane>



              </el-tabs>


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
    name: 'Weixinreplays',
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
