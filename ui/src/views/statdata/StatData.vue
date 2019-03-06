<template>
    <div style="margin-top: 50px">
        <Button id="btn" type='primary' @click="createTestConf">START</Button>
        <Button id="btn2" type='primary' @click="createTestConf">END</Button>

        <Card>
            <h1 slot="title">Stress Data Stat</h1>

            <Table class="ordertable" ref="orderTable" :columns="columns" :data="tableData"></Table>
        </Card>

        <Card>
            <h1 slot="title">Stress Data Chart</h1>

            <div class="container">
                <line-chart v-if="loaded" :chartdata="chartdata" :options="options"/>
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
    import LineChart from './LineChart.vue'

    export default {
        name: "StatDataList",
        components: {LineChart},
        data() {
            return {
                formData: {},
                tableData: [],
                filterform: {},
                pager: {...pagerInit},
                loaded: false,
                chartdata: null,
                options: {
                    responsive: true,
                    maintainAspectRatio: false
                },
            }
        },
        computed: {
            columns() {
                return cols(this)
            }
        },
        methods: {

            async fetchStatData() {
                this.loaded = false

                // debugger
                const response = await fetchStatData(_omit(this.pager, 'total'))
                // @ts-ignore
                this.tableData = response.stats
                // @ts-ignore
                this.pager.total = response.count

                try {
                    this.chartData = response.stats
                    this.loaded = true
                } catch (e) {
                    console.error(e)
                }


            },
        },
        mounted() {
            this.fetchStatData()
        }

    }
</script>

<style scoped>

</style>
