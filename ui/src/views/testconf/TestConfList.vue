<template>
    <div>
        <Card>
            <h1 slot="title">Test Conf List</h1>

            <Button id="btn" type='primary' @click="createTestConf">NEW</Button>

            <Table class="ordertable" ref="orderTable" :columns="columns" :data="tableData"></Table>
        </Card>

    </div>
</template>

<script>
    const pagerInit = {
        page: 1,
        page_size: 10,
        total: 0,
    }

    import {deleteTestConfByID, fetchTestConf} from '@/apis/testconf'
    import {cols} from './helper.jsx'
    import _omit from 'lodash/omit'

    export default {
        name: "TestConfList",
        data() {
            return {
                formData: {},
                tableData: [],
                filterform: {},
                pager: {...pagerInit},
            }
        },
        computed: {
            columns() {
                return cols(this)
            }
        },
        methods: {

            edit(row) {
                this.$router.push({name: 'testconfDetail', params: {id: row._id}})
            },

            stress(row) {
                this.$router.push({name: 'statdata', params: {id: row._id}})
            },

            createTestConf() {
                this.$router.push({name: 'testconfDetail', params: {id: 'null'}})
            },

            async remove(row) {

                this.$Modal.confirm({
                    title: 'confirm',
                    content: 'confirm delete' + row._id + '?',
                    width: 350,
                    onOk: async () => {
                        await deleteTestConfByID(row._id, {})
                        this.fetchTestConf()
                    },
                })
            },

            async fetchTestConf() {
                // debugger
                const response = await fetchTestConf(_omit(this.pager, 'total'))
                // @ts-ignore
                this.tableData = response
                // @ts-ignore
                this.pager.total = response.count
            },
        },
        mounted() {
            this.fetchTestConf()
        }

    }
</script>

<style scoped>

</style>
