<template>
    <div>
        <Button id="btn2" @click="goback">Go Back</Button>


        <div>
            <Row>
                <Col offset="20">
                    <Button id="btn" type='primary' :disabled="startdisable" @click="startStressTest">START</Button>
                    <Button id="btn1" type='error' :disabled="enddisable" @click="StopStressTest">END</Button>
                </Col>
            </Row>

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
    </div>
</template>

<script>
    const pagerInit = {
        page: 1,
        page_size: 10,
        total: 0,
    }

    import {fetchStatData} from '@/apis/stat'
    import {StartStress, StopStress} from '@/apis/control'
    import {cols} from './helper.jsx'
    import _omit from 'lodash/omit'
    import VeLine from 'v-charts/lib/line.common'

    var dateFormat = require('dateformat');


    export default {
        name: "StatDataList",
        components: {VeLine},
        data() {
            return {
                startdisable: false,
                enddisable: true,
                chartData: {
                    columns: ['ts',],
                    rows: [
                        {'ts': '01-01',},
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
            goback() {
                this.$router.push({name: 'testconfList',})
            },

            startStressTest() {
                const id = this.$route.params.id
                StartStress(id)
                this.startdisable = true
                this.enddisable = false
            },
            StopStressTest() {
                const id = this.$route.params.id
                StopStress(id)

                this.startdisable = false
                this.enddisable = true
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
                    if (Math.abs(val.current_rps) < 1e-3) {
                        this.startdisable = false
                        this.enddisable = true
                    } else {
                        this.startdisable = true
                        this.enddisable = false
                    }
                })
                this.chartData.columns = ['ts', ...cols]
                this.chartData.rows.push(row)

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
