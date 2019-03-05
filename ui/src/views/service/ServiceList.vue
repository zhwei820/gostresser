<template>
  <div>
    <Button id="btn" type='primary' @click="createService">NEW</Button>

    <Table class="ordertable" ref="orderTable" :columns="columns" :data="tableData"></Table>

  </div>
</template>

<script>
  const pagerInit = {
    page: 1,
    page_size: 10,
    total: 0,
  }

  import {fetchService, deleteServiceByName} from '@/apis/service'
  import {cols} from './helper.jsx'
  import _omit from 'lodash/omit'

  export default {
    name: "ServiceList",
    data(){
      return {
        formData :{},
        tableData :[],
        filterform :{},
        pager :{...pagerInit},
      }
    },
    computed:{
      columns() {
        return cols(this)
      }
    },
    methods:{
      // handlePageSizeChange(pageSize) {
      //   this.pager.page = this.pager.page_size < pageSize ?
      //     Math.ceil(this.pager.page * (this.pager.page_size / pageSize)) : this.pager.page
      //   this.pager.page_size = pageSize
      //
      //   this.fetchService()
      // },
      //
      // handlePageChange(page) {
      //   this.pager.page = page
      //   this.fetchService()
      // },

      edit(row) {
        this.$router.push({name: 'serviceDetail', params: {name: row.name}})
      },

      copy(row) {
        this.$router.push({name: 'serviceDetail', params: {name: row.name}, query: {copy:'true'}})
      },

      createService() {
        this.$router.push({name: 'serviceDetail', params: {name: 'null'}})
      },

      async remove(row) {

        this.$Modal.confirm({
          title: 'confirm',
          content: 'confirm delete' + row.name + '?',
          width: 350,
          onOk: async () => {
            await deleteServiceByName(row.name, {})
            this.fetchService()
          },
        })
      },

      async fetchService() {
        // debugger
        const response = await fetchService(_omit(this.pager, 'total'))
        // @ts-ignore
        this.tableData = response
        // @ts-ignore
        this.pager.total = response.count
      },
    },
    mounted() {
      this.fetchService()
    }

  }
</script>

<style scoped>

</style>
