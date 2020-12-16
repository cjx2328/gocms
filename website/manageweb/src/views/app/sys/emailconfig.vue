<template>
  <div class="app-container">
    <div class="filter-container">
      <el-row>
        <el-col span="12">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="基础设置" name="baseConfig">
              <el-form ref="form" :model="form" label-width="120px">
                <el-form-item label="服务器">
                  <el-input v-model="sysconfig.email_host" placeholder="请输入服务器地址" />
                </el-form-item>
                <el-form-item label="SMTP端口">
                  <el-input v-model="sysconfig.email_SMTP_port" placeholder="请输SMTP端口号" />
                </el-form-item>

                <el-form-item label="发件人">
                  <el-input v-model="sysconfig.email_send_user" placeholder="请输入发件人" />
                </el-form-item>
                <el-form-item label="发件邮箱">
                  <el-input v-model="sysconfig.email_send_email_account" placeholder="请输入发件邮箱账号" />
                </el-form-item>
                <el-form-item label="身份码">
                  <el-input v-model="sysconfig.email_smtp_passwords" placeholder="请输入SMTP发件身份验证码" />
                </el-form-item>
                <el-form-item label="测试收件邮箱">
                  <el-input v-model="sysconfig.email_test_email" placeholder="请输入收件的测试邮箱" />
                </el-form-item>

                <el-form-item label="测试邮件内容">
                  <el-input v-model="sysconfig.email_test_contents" type="textarea" placeholder="请输入准备发生的测试内容" />
                  <span>【请勿测试违法违规内容】</span>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary" @click="savesysconfig">保存</el-button>
                  <el-button>取消</el-button>
                </el-form-item>
              </el-form>
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
  name: 'Emailconfig',
  data() {
    return {
      sysconfig: {
        site_name: '忠告'
      },
      activeName: 'baseConfig',
      fileList: [{ name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }],

      value1: true,
      value2: true
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    }
  }
}

</script>
