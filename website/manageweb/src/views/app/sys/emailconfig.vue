<template>
  <div class="app-container">
    <div class="filter-container">
      <el-row>
        <el-col :span="12">
          <el-tabs v-model="activeName"  >
            <el-tab-pane label="基础设置" name="baseConfig">
              <el-form   :model="sysconfig" label-width="120px">
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
                  <el-input v-model="sysconfig.email_test_email" placeholder="请输入收件的测试邮箱" >
                    <el-button slot="append"  >发送测试邮件</el-button>
                  </el-input>
                </el-form-item>
                <el-form-item label="测试邮件内容">
                  <el-input v-model="sysconfig.email_test_contents" type="textarea" placeholder="请输入准备发生的测试内容" />
                  <span>【请勿测试违法违规内容】</span>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="savesysconfig()">保存</el-button>
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
import {requestsysconfigs ,updatesysconfigs} from "../../../api/app/sys/systemconfig";
export default {
  name: 'Emailconfig',
  data() {
    return {
      sysconfig: {
          email_host:"",
          email_SMTP_port:"",
          email_send_user:"",
          email_send_email_account:"",
          email_smtp_passwords:"",
          email_test_email:"",
          email_test_contents:"",


      },
      activeName: 'baseConfig',

    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    },
      getconfig(){
        var _this = this;
        const token = _this.$store.getters.token;
        _this.listLoading = true;
          requestsysconfigs(token).then(response=>{
              var _thisdata = _this.sysconfig;
              var _this_response_data = response.data;
              if(response.code !==20000){
                  this.$alert('获取数据失败，请稍后再试！', '获取系统配置数据失败', {
                      confirmButtonText: '确定',
                  });
              }else{
                  _this.verifydata =_this_response_data;
                  for(var key in _this_response_data){
                      if(_this_response_data.hasOwnProperty(key)){
                          _thisdata[key] = _this_response_data[key].content;
                      }
                  }
              }
              console.log(_this.sysconfig);
              this.listLoading = false
          })

      },
      savesysconfig(){
          this.loading = true;
          var _this = this;
          const sysconfigdataData = Object.assign({}, _this.sysconfig);
          var verifyolddata = _this.verifydata;
          var newarrays = {};
          var senddata = [];
          for (var key in sysconfigdataData){
              if(sysconfigdataData[key].length>0 ){
                  if(verifyolddata.hasOwnProperty(key)    ){
                      if(sysconfigdataData[key]!==verifyolddata[key].content){
                          newarrays["title"] =  key;
                          newarrays["content"] = sysconfigdataData[key];
                          senddata.push(newarrays);
                          newarrays={};
                      }
                  }else{
                      newarrays["title"] =  key;
                      newarrays["content"] = sysconfigdataData[key];
                      senddata.push(newarrays);
                      newarrays={};
                  }
              }
          }

          var  sendnewarrays = Object.assign({}, newarrays);


          if(senddata.length  > 0){
              updatesysconfigs( senddata).then(response=>{
                  var _this_response_data = response.data;
                  if(response.code !==20000){
                      this.$alert('获取数据失败，请稍后再试！', '获取系统配置数据失败', {
                          confirmButtonText: '确定',
                      });
                  }else{
                      this.$message({
                          message: '更新成功',
                          type: 'success'
                      });

                      _this.getList();
                  }


              });
          }else{
              this.$alert('需要上传的数据没有变化或者为空，请修改或者查证后再试！', '请确定数据正确性', {
                  confirmButtonText: '确定',
              });

          }

      },
  },
  created(){
      this.getconfig();
  }
}

</script>
