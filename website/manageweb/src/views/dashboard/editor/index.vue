<template>
  <div class="dashboard-editor-container">
    <el-row  :gutter="20">
      <el-col span="8"    >

        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>待办事项</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div  :data="tableData"  class="text item">
            werewr
          </div>
        </el-card>

  <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>业务处理</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div  :data="tableData"  class="text item">
            werewr
          </div>
        </el-card>


  <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>网站访问情况</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div    class="text item">



            <el-table
                    :data="tableData"
                    style="width: 100%">
              <el-table-column
                      prop="userpv"
                      label="浏览量(PV)"
              >
              </el-table-column>
      <el-table-column
                      prop="useruv"
                      label="访客数(UV)"
              >
              </el-table-column>
      <el-table-column
                      prop="userip"
                      label="独立IP(IP)"
              >
              </el-table-column>


            </el-table>

          </div>
        </el-card>


      </el-col>
      <el-col span="16" >


        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>订单统计报表</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div     class="text item">



            <div ref="initCharts" :style="{height: '400px',width:'100%'}"  ></div>



          </div>
        </el-card>




        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>业务处理</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div  :data="tableData"  class="text item">
            <el-table
                    :data="tableData"
                    style="width: 100%">

              <el-table-column
                      prop="date"
                      label="排名"
              >
              </el-table-column>

              <el-table-column
                      prop="date"
                      label="商品名称"
              >
              </el-table-column>

              <el-table-column
                      prop="date"
                      label="数量"
              >
              </el-table-column>


            </el-table>
          </div>
        </el-card>












      </el-col>

    </el-row>
  </div>
</template>
<style>
  .el-col > .el-table {
    margin-top: 20px;
  }


</style>
<script>
import { mapGetters } from 'vuex'
import PanThumb from '@/components/PanThumb'
// import GithubCorner from '@/components/GithubCorner'

//全部引入
var echarts = require('echarts')


export default {

  name: 'DashboardEditor',
  components: { PanThumb },
  data() {
    return {
      emptyGif: 'https://wpimg.wallstcn.com/0e03b7da-db9e-4819-ba10-9016ddfdaed3' // '/resource/img/walk.gif'
    , imageurls :"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAK8AAABQCAYAAABvViW5AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyFpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNS1jMDE0IDc5LjE1MTQ4MSwgMjAxMy8wMy8xMy0xMjowOToxNSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIChXaW5kb3dzKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDozRERBRURENjYyMUExMUU4QUZDNUY4MzNGMzM0MjIwRiIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDozRERBRURENzYyMUExMUU4QUZDNUY4MzNGMzM0MjIwRiI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjNEREFFREQ0NjIxQTExRThBRkM1RjgzM0YzMzQyMjBGIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjNEREFFREQ1NjIxQTExRThBRkM1RjgzM0YzMzQyMjBGIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+7mX1cgAACwBJREFUeNrsXWuMG1cV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"
        ,
        tableData: [{
            todojobs: '您有23个订单需要处理',
            businessprodate: '王小虎',
            address: '上海市普陀区金沙江路 1518 弄'
        }, {
            todojobs: '您有45个企业需要认证核验',
            businessprodate: '王小虎',
            address: '上海市普陀区金沙江路 1517 弄'
        }, {
            todojobs: '您有235条消息需要处理',
            businessprodate: '王小虎',
            address: '上海市普陀区金沙江路 1519 弄'
        }, {
            todojobs: '您有23笔退款需要处理',
            businessprodate: '王小虎',
            address: '上海市普陀区金沙江路 1516 弄'
        }]

    }

  },
  computed: {
    ...mapGetters([
      'name',
      'avatar',
      'roles'
    ])
  },
    mounted() {
        this.initCharts();
    },
    methods: {
        initCharts() {
            this.chart = echarts.init(this.$refs.initCharts);
            this.setOptions();
        },
        setOptions() {
            this.chart.setOption({

                tooltip: {},
                xAxis: {
                    data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"]
                },
                yAxis: {},
                series: [{
                    name: '销量',
                    type: 'line',
                    data: [5, 20, 36, 10, 10, 20]
                }]
            })
        }
    },
    watch: {

    }
}

</script>

<style lang="scss" scoped>
  .emptyGif {
    display: block;
    width: 45%;
    margin: 0 auto;
  }

  .dashboard-editor-container {
    background-color: #e3e3e3;
    min-height: 100vh;
    padding: 10px 10px 0px;
    .pan-info-roles {
      font-size: 12px;
      font-weight: 700;
      color: #333;
      display: block;
    }
    .info-container {
      position: relative;
      margin-left: 190px;
      height: 150px;
      line-height: 200px;
      .display_name {
        font-size: 48px;
        line-height: 48px;
        color: #212121;
        position: absolute;
        top: 25px;
      }
    }
  }
</style>
