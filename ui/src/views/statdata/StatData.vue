<template>
    <div style="margin-top: 50px">
        <Button id="btn" type='primary' @click="startStressTest">START</Button>
        <Button id="btn2" type='primary' @click="startStressTest">END</Button>

        <Card>
            <h1 slot="title">Stress Data Stat</h1>

            <Table class="ordertable" ref="orderTable" :columns="columns" :data="tableData"></Table>
        </Card>

        <Card>
            <h1 slot="title">Stress Data Chart</h1>

            <div class="container">
                <ve-line :data="chartData"></ve-line>
            </div>
        </Card>
    </div>
</template>

<script>
    const pagerInit = {
        page: 1,
        page_size: 10,
        total: 0,
    }

    import {fetchStatData} from '@/apis/stat'
    import {cols} from './helper.jsx'
    import _omit from 'lodash/omit'
    import VeLine from 'v-charts/lib/line.common'


    export default {
        name: "StatDataList",
        components: {VeLine},
        data() {
            return {
                chartData: {
                    columns: ['ts', 'rps', 'test'],
                    rows: [
                        {'ts': '01-01', 'rps': 1231, 'test': 1931},
                        {'ts': '01-02', 'rps': 1223, 'test': 1933},
                        {'ts': '01-03', 'rps': 2123, 'test': 2933},
                        {'ts': '01-04', 'rps': 4123, 'test': 4933},
                        {'ts': '01-05', 'rps': 3123, 'test': 3933},
                        {'ts': '01-06', 'rps': 7123, 'test': 7933}
                    ]
                },
                formData: {},
                tableData: [],
                filterform: {},
                pager: {...pagerInit},
                loaded: false,
            }
        },
        computed: {
            columns() {
                return cols(this)
            }
        },
        methods: {
            startStressTest() {
            },

            async fetchStatData() {
                this.loaded = false

                // debugger
                const response = await fetchStatData(_omit(this.pager, 'total'))
                // @ts-ignore
                this.tableData = response.stats ? response.stats : []
                // @ts-ignore
                this.pager.total = response.count
                //

            },
        },
        mounted() {
            this.fetchStatData()
        }

    }
</script>

<style scoped>

</style>
