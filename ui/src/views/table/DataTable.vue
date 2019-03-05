<template>
  <div>
    <Button id="btn" type='primary' @click="createAuth">NEW</Button>

    <Table class="ordertable" ref="orderTable" :columns="columns" :data="tableData"></Table>

  </div>
</template>

<script>
  const pagerInit = {
    page: 1,
    page_size: 10,
    total: 0,
  }

  import {fetchAuth, deleteAuthByName} from '@/apis/auths'
  import {cols} from './helper.jsx'
  import _omit from 'lodash/omit'

  export default {
    name: "AuthList",
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

      edit(row) {
        this.$router.push({name: 'authDetail', params: {name: row.name}})
      },

      copy(row) {
        this.$router.push({name: 'authDetail', params: {name: row.name}, query: {copy:'true'}})
      },

      createAuth() {
        this.$router.push({name: 'authDetail', params: {name: 'null'}})
      },

      async remove(row) {

        this.$Modal.confirm({
          title: 'confirm',
          content: 'confirm delete' + row.name + '?',
          width: 350,
          onOk: async () => {
            await deleteAuthByName(row.name, {})
            this.fetchAuth()
          },
        })
      },

      async fetchAuth() {
        // debugger
        const response = await fetchAuth(_omit(this.pager, 'total'))
        // @ts-ignore
        this.tableData = response
        // @ts-ignore
        this.pager.total = response.count
      },
    },
    mounted() {
      this.fetchAuth()
    }

  }
</script>

<style scoped>

</style>
