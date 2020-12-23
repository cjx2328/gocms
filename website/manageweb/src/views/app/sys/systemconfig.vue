<template>
  <div class="app-container">
    <div class="filter-container">
      <el-row>
        <el-col span="8">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="基础设置" name="baseConfig">
              <el-form ref="form" :model="form" label-width="80px">
                <el-form-item label="网站名称">
                  <el-input v-model="sysconfig.site_name" placeholder="请输入网站名称" />
                </el-form-item>
                <el-form-item label="网站地址">
                  <el-input v-model="sysconfig.site_host" placeholder="请输入网站的地址 例如: http://www.google.com"/>
                </el-form-item>
                <el-form-item label="网站logo">
                  <el-upload
                          class="upload-demo"
                          action="https://jsonplaceholder.typicode.com/posts/"
                          :on-preview="handlePreview"
                          :on-remove="handleRemove"
                          :file-list="fileList"
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
                  <el-button type="primary" @click="savesysconfig">保存</el-button>
                  <el-button>取消</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="SEO设置" name="seoConfig">
              <el-form ref="form" :model="form" label-width="80px">
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
                  <el-button type="primary" @click="savesysconfig">保存</el-button>
                  <el-button>取消</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="其他设置" name="otherConfig">
              <el-form ref="form" :model="form" label-width="80px">
                <el-form-item label="验证码">

                  <el-switch
                          v-model="sysconfig.site_verify_code_switch"
                          active-text="开启"
                          inactive-text="关闭">
                  </el-switch>


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
import {requestsysconfigs} from "../../../api/app/sys/systemconfig";

export default {
  name: 'Systemconfig',
  data() {
    return {
      sysconfig: {
        site_name: 'sdfdsf',
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
      activeName: 'baseConfig',
        fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}]

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
               console.group("test")
               console.log("test")
               console.log(response.data.site_name.content);

               var _thisdata = _this.sysconfig;
               var _this_response_data = response.data;
                console.log(_this_response_data);
                console.log(response.code);
               if(response.code !==20000){

               }else{
                   console.log(_thisdata);
                   _thisdata.site_name                =_this_response_data.site_name.content;
                   _thisdata.site_host                =_this_response_data.site_host.content;
                   _thisdata.site_images              =_this_response_data.site_images.content;
                   _thisdata.site_icp                 =_this_response_data.site_icp.content;
                   _thisdata.site_copyright           =_this_response_data.site_copyright.content;
                   _thisdata.site_addresss            =_this_response_data.site_addresss.content;
                   _thisdata.site_mobile              =_this_response_data.site_mobile.content;
                   _thisdata.site_email               =_this_response_data.site_email.content;
                   _thisdata.site_countcode            =_this_response_data.site_countcode.content;
                   _thisdata.site_servicecode          =_this_response_data.site_servicecode.content;
                   _thisdata.site_seo_title           =_this_response_data.site_seo_title.content;
                   _thisdata.site_seo_keywords         =_this_response_data.site_seo_keywords.content;
                   _thisdata.site_seo_desc             =_this_response_data.site_seo_desc.content;
                   _thisdata.site_verify_code_switch  =_this_response_data.site_verify_code_switch.content;


               }

                this.listLoading = false
            })
        },

        savesysconfig(){

            this.loading = true;
            // var _this = this;
            // var _thissysconfig = this .sysconfig ;
            const sysconfigdataData = Object.assign({}, this.sysconfig)
             console.log(sysconfigdataData);

        }
    },
    created(){
      this.getList()
    }
}

</script>
