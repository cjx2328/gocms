<template>
  <div class="app-container">
    <div class="filter-container">
      <el-row>
        <el-col span="8">
          <el-tabs v-model="activeName"  >
            <el-tab-pane label="基础设置" name="baseConfig">
              <el-form ref="form" :model="sysconfig" label-width="80px">
                <el-form-item label="网站名称">
                  <el-input v-model="sysconfig.site_name" placeholder="请输入网站名称" />
                </el-form-item>
                <el-form-item label="网站地址">
                  <el-input v-model="sysconfig.site_host" placeholder="请输入网站的地址 例如: http://www.google.com"/>
                </el-form-item>
                <el-form-item label="网站logo">
                  <el-input type="hidden" v-model="sysconfig.site_log"></el-input>
                  <el-upload
                          class="upload-demo"
                          action="/api/systemconfig/uploadedfile"
                          :on-preview="handlePreview"
                          :on-remove="handleRemove"
                          :file-list="fileList"
                          :before-upload="onBeforeUpload"
                          :on-success="uploadFileSuccess"

                          list-type="picture">
                    <el-button size="small" type="primary">点击上传</el-button>
                    <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                  </el-upload>
                </el-form-item>
                <el-form-item label="备案号">
                  <el-input  v-model="sysconfig.site_icp" placeholder="请输入网站备案号"/>
                </el-form-item>
                <el-form-item label="copyright">
                  <el-input  v-model="sysconfig.site_copyright" placeholder="请输入网站copyright"/>
                </el-form-item>
                <el-form-item label="公司地址">
                  <el-input v-model="sysconfig.site_addresss" placeholder="请输入网站通信地址" />
                </el-form-item>
                <el-form-item label="联系电话">
                  <el-input v-model="sysconfig.site_mobile" placeholder="请输入联系方式"/>
                </el-form-item>
                <el-form-item label="邮箱账号">
                  <el-input v-model="sysconfig.site_email" placeholder="请输入E_mail"/>
                </el-form-item>
                <el-form-item label="统计代码">
                  <el-input type="textarea" v-model="sysconfig.site_countcode" placeholder="请输入统计代码：谷歌、友盟等"/>
                </el-form-item>
                <el-form-item label="客服代码">
                  <el-input type="textarea" v-model="sysconfig.site_servicecode" placeholder="请输入在线客服代码"/>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="savesysconfig()">保存</el-button>
                  <el-button>取消</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="SEO设置" name="seoConfig">
              <el-form ref="form" :model="sysconfig" label-width="80px">
                <el-form-item label="seo标题">
                  <el-input v-model="sysconfig.site_seo_title" placeholder="请输入seo标题" />
                </el-form-item>
                     <el-form-item label="seo关键字">
                  <el-input v-model="sysconfig.site_seo_keywords" placeholder="请输入seo关键字" />
                </el-form-item>
                     <el-form-item label="seo描述">
                  <el-input type="textarea" v-model="sysconfig.site_seo_desc" placeholder="请输入seo描述" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="savesysconfig()">保存</el-button>
                  <el-button>取消</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="其他设置" name="otherConfig">
              <el-form ref="form" :model="sysconfig" label-width="80px">
                <el-form-item label="验证码">

                  <el-switch
                          v-model="sysconfig.site_verify_code_switch"
                          active-value="1" inactive-value="0"
                          active-text="开启"
                          inactive-text="关闭">
                  </el-switch>


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
  name: 'Systemconfig',
  data() {
    return {
      sysconfig: {
        site_name: '',
        site_host: '',
        site_images: '',
        site_icp: '',
        site_copyright: '',
        site_addresss: '',
        site_mobile: '',
        site_email: '',
        site_countcode: '',
        site_servicecode: '',
        site_seo_title: '',
        site_seo_keywords: '',
        site_seo_desc: '',
        site_verify_code_switch: '',
      },
        verifydata:{},
      activeName: 'baseConfig',
        fileList: []

        , value1: true,
        value2: true
    }
  },
    computed:{



    },
    methods: {

        handleRemove(file, fileList) {
            console.log(file, fileList);
        },
        handlePreview(file) {
            console.log(file);
        },getList() {
            var _this = this;
            const token = this.$store.getters.token;
            this.listLoading = true;
            requestsysconfigs(token).then(response => {
               var _thisdata = _this.sysconfig;
               var _this_response_data = response.data;
               if(response.code !==20000){
                   this.$alert('获取数据失败，请稍后再试！', '获取系统配置数据失败', {
                       confirmButtonText: '确定',
                   });
               }else{
                   _this.verifydata =_this_response_data;
                   for(var key in _thisdata){
                      if(_this_response_data.hasOwnProperty(key)){
                          _thisdata[key] = _this_response_data[key].content;
                          if(key=="site_images"){
                              console.log(_this_response_data[key].content);
                              _this.fileList =  [{name: 'site_images', url: "http://localhost:8080/"+ _this_response_data[key].content}];
                          }

                      }
                   }
               }
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
              // console.log(JSON.stringify(newarrays));
              // console.log(JSON.stringify(senddata));
              // console.log( senddata );




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
        onBeforeUpload(file)
        {
            const isIMAGE = file.type === 'image/jpeg'||'image/gif'||'image/png';
            const isLt1M = file.size / 1024 / 1024 < 1;

            if (!isIMAGE) {
                this.$message.error('上传文件只能是图片格式!');
            }
            if (!isLt1M) {
                this.$message.error('上传文件大小不能超过 1MB!');
            }
            console.log(isIMAGE);
            console.log(isLt1M);
            console.log(file);
            return isIMAGE && isLt1M;
        },
        uploadFileSuccess(response, file, fileList) {
            var _this = this;
            console.log(response);
            console.log(file);
            if (response.code == 20000) {
                _this.sysconfig.site_images = response.data;
                console.log(file);
            } else {
                this.$message.error(response.message);//文件上传错误提示
            }
        }
    },
    created(){
      this.getList()
    },

}

</script>
