<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form>
        <el-form-item label="站点名称">
          <el-input v-model="sysconfig.site_name" />
        </el-form-item>
      </el-form>

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
  name: 'Systemconfig',
  data() {
    return {
      sysconfig: {
        site_name: '忠告'
      }
    }
  }
}

</script>
