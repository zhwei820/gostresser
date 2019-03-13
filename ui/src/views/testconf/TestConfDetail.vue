<template>
    <div>

        <Card>
            <h1 slot="title">Test Conf Detail</h1>

            <Form ref="formDynamic" :model="formDynamic" :label-width="100" style="width: 850px">

                <FormItem label="Name">
                    <Col span="12">
                        <Input v-model="formDynamic.name" placeholder="Enter your name"></Input>
                    </Col>
                </FormItem>

                <FormItem label="H2">
                    <i-switch v-model="formDynamic.h2" size="large">
                        <span slot="open">On</span>
                        <span slot="close">Off</span>
                    </i-switch>
                </FormItem>

                <FormItem label="Cpus">
                    <Col span="12">
                        <InputNumber v-model="formDynamic.cpus" placeholder="Enter your cpus"></InputNumber>
                    </Col>
                </FormItem>

                <FormItem label="Disable Compression">
                    <i-switch v-model="formDynamic.disablecompression" size="large">
                        <span slot="open">On</span>
                        <span slot="close">Off</span>
                    </i-switch>
                </FormItem>

                <FormItem label="Disable Keepalives">
                    <i-switch v-model="formDynamic.disablekeepalives" size="large">
                        <span slot="open">On</span>
                        <span slot="close">Off</span>
                    </i-switch>
                </FormItem>

                <FormItem label="Disable Redirects">
                    <i-switch v-model="formDynamic.disableredirects" size="large">
                        <span slot="open">On</span>
                        <span slot="close">Off</span>
                    </i-switch>
                </FormItem>


                <FormItem label="Proxy Addr">
                    <Col span="12">
                        <Input v-model="formDynamic.proxyaddr" placeholder="Enter your proxyaddr"></Input>
                    </Col>
                </FormItem>


                <FormItem label="Timeout (in seconds)">
                    <Col span="12">
                        <InputNumber v-model="formDynamic.timeout" placeholder="Enter your Timeout"></InputNumber>
                    </Col>
                </FormItem>


                <FormItem label="Duration (in seconds)">
                    <Col span="12">
                        <InputNumber v-model="formDynamic.duration" placeholder="Enter your Duration"></InputNumber>
                    </Col>
                </FormItem>


                <FormItem label="Concurrency">
                    <Col span="12">
                        <InputNumber v-model="formDynamic.concurrency"
                                placeholder="Enter your Concurrency"></InputNumber>
                    </Col>
                </FormItem>


                <Card>
                    <p slot="title">Request Conf List</p>

                    <Card v-for="(item, ii) in formDynamic.reqconfs" :key="ii">

                        <FormItem label="Url">
                            <Col span="12">
                                <Input v-model="item.url" placeholder="Enter your url"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="HTTP Method">
                            <Col span="12">
                                <Select v-model="item.method" placeholder="Enter your HTTP">
                                    <Option v-for="item in method_names" :value="item" :key="item">{{ item}}
                                    </Option>
                                </Select>
                            </Col>
                        </FormItem>

                        <FormItem label="Body">
                            <Col span="12">
                                <Input v-model="item.body" placeholder="Enter your body"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="Accept">
                            <Col span="12">
                                <Input v-model="item.accept" placeholder="Enter your accept"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="Contenttype">
                            <Col span="12">
                                <Input v-model="item.contenttype" placeholder="Enter your contenttype"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="Authheader">
                            <Col span="12">
                                <Input v-model="item.authheader" placeholder="Enter your authheader"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="Hostheader">
                            <Col span="12">
                                <Input v-model="item.hostheader" placeholder="Enter your hostheader"></Input>
                            </Col>
                        </FormItem>

                        <FormItem label="ReqFreq">
                            <Col span="12">
                                <InputNumber v-model="item.ReqFreq" placeholder="Enter your ReqFreq"></InputNumber>
                            </Col>
                        </FormItem>


                        <FormItem>
                            <Row>
                                <Col span="20">
                                    <Col span="2" offset="15">
                                        <Button type="warning" @click="handleRemoveUrlSettings(ii)">del</Button>
                                    </Col>
                                </Col>
                            </Row>
                        </FormItem>
                    </Card>

                    <FormItem>
                        <Row>
                            <Col span="6">
                                <Button type="dashed" long @click="handleAddUrlSettings" icon="md-add">Add Test Request
                                </Button>
                            </Col>
                        </Row>
                    </FormItem>

                </Card>


                <FormItem>
                    <Button type="primary" @click="handleSubmit('formDynamic')">Submit</Button>
                </FormItem>
            </Form>
        </Card>
    </div>
</template>
<script lang="jsx">

    import {createTestConf, fetchTestConfByID, putTestConf} from '@/apis/testconf'
    import InputForm from '@/components/InputForm'

    export default {
        name: "TestConfDetail",
        components: {
            // eslint-disable-next-line vue/no-unused-components
            InputForm: InputForm,
        },
        data() {
            return {
                method_names: [
                    'GET',
                    'POST',
                    'PUT',
                    'PATCH',
                    'OPTIONS',
                ],

                copy: '',
                formDynamic: {
                    "_id": "",
                    "name": "",
                    "h2": false,
                    "cpus": 4,
                    "disablecompression": false,
                    "disablekeepalives": false,
                    "disableredirects": false,
                    "proxyaddr": "",
                    "timeout": 10,
                    "duration": 30,
                    "concurrency": 20,
                    "reqconfs": [{
                        "url": "https://cn.bing.com/",
                        "method": "GET",
                        "headers": [],
                        "body": "",
                        "accept": "",
                        "contenttype": "",
                        "authheader": "",
                        "hostheader": "",
                        "ReqFreq": 100000
                    }]
                },
                authData: {},
            }
        },
        computed: {
            id() {
                if (!this.copy) {
                    const id = this.$route.params.id
                    return id === 'null' ? null : id
                }
                return null
            }
        },
        methods: {

            async handleSubmit(name) {
                const res = await this.createOrUpdateTestConf(this.formDynamic)
            },
            handleAddUrlSettings() {
                this.formDynamic.reqconfs.push({
                    "url": "https://www.baidu.com",
                    "method": "GET",
                    "headers": [],
                    "body": "",
                    "accept": "",
                    "contenttype": "",
                    "authheader": "",
                    "hostheader": "",
                    "ReqFreq": 0
                });
            },

            handleRemoveUrlSettings(ii) {
                if (this.formDynamic.reqconfs.length > 1) {
                    this.formDynamic.reqconfs.splice(ii, 1)
                } else {
                    this.$Message.warning('Can not delete last request!')
                }
            },

            async createOrUpdateTestConf(data) {
                this.errors = {}  //
                                  // https://github.com/iview/iview/blob/master/src/components/form/form-item.vue#L81
                try {
                    if (!this.id) {
                        await createTestConf(data)
                    } else {
                        await putTestConf(this.id, data)
                    }
                    this.$router.push({name: 'testconfList'})
                    this.$Message.info('Success !!')

                } catch (e) {
                    console.log(e)
                    this.errors = e.data || {}
                }
            },

            async fetchTestConfByID(name) {
                const response = await fetchTestConfByID(name)
                // @ts-ignore
                this.formDynamic = response
            },

        },
        created() {
            if (this.id) {
                this.fetchTestConfByID(this.id)
            }
        }
    }

</script>

<style>
    .ivu-card .ivu-card-bordered {
        margin-bottom: 15px;
    }
</style>
