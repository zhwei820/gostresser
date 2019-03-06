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
    var dateFormat = require('dateformat');


    export default {
        name: "StatDataList",
        components: {VeLine},
        data() {
            return {
                chartData: {
                    columns: ['ts', 'rps', 'test'],
                    rows: [
                        {'ts': '01-01', 'rps': 1231, 'test': 1931},
                    ]
                },
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
            startStressTest() {
            },

            async fetchStatData() {
                // debugger
                const response = await fetchStatData(_omit(this.pager, 'total'))
                // @ts-ignore
                this.tableData = response.stats ? response.stats : []
                // @ts-ignore
                this.pager.total = response.count
                //
                const cols = []

                var now = new Date();
                const row = {ts: dateFormat(now, "hh:MM:ss")}
                this.tableData.forEach((val) => {
                    cols.push(val.name)
                    row[val.name] = val.current_rps
                    console.log(val)

                })
                this.chartData.columns = ['ts', ...cols]
                this.chartData.rows.push(row)
                console.log(row)

            },
        },
        mounted() {
            this.fetchStatData()

            setInterval(() => {
                this.fetchStatData()
            }, 2000)
        }

    }
</script>

<style scoped>

</style>
