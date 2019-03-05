<template>
    <Card>
        <p slot="title">{{title}}</p>

        <FormItem label="Listen Path" >
            <Col span="12">
                <Input v-model="proxy.listen_path" placeholder="Enter your name"></Input>
            </Col>
        </FormItem>


        <FormItem label="append_path" prop="append_path">

            <i-switch size="large" v-model="proxy.append_path">
                <span slot="open">ON</span>
                <span slot="close">OFF</span>
            </i-switch>
        </FormItem>


        <FormItem label="strip_path" prop="strip_path">

            <i-switch size="large" v-model="proxy.strip_path">
                <span slot="open">ON</span>
                <span slot="close">OFF</span>
            </i-switch>
        </FormItem>


        <FormItem label="methods" >
            <Col span="12">
                <Select v-model="proxy.methods" multiple>
                    <Option v-for="item in MethodsList" :value="item" :key="item">{{ item }}</Option>
                </Select>
            </Col>
        </FormItem>

        <FormItem label="Balancing" >
            <Col span="12">
                <Select v-model="proxy.upstreams.balancing" >
                    <Option v-for="item in BalanceList" :value="item" :key="item">{{ item }}</Option>
                </Select>
            </Col>
        </FormItem>

        <FormItem label="Target"
                  v-for="(item, ii) in proxy.upstreams.targets"
                  :key="ii"
                  :rules="{required: false, message: '', trigger: 'blur'}">
            <Row>
                <Col span="10">
                    <Input type="text" v-model="item.target" placeholder="Enter something..."></Input>
                </Col>
                <Col span="2" offset="1">
                    <span style="float: right; margin-right: 10px"> weight </span>
                </Col>
                <Col span="2">
                    <Input type="text" v-model="item.weight" placeholder="Enter something..."></Input>
                </Col>
                <Col span="2" offset="1">
                    <Button type="warning" @click="handleRemoveTarget(ii)">del</Button>
                </Col>
            </Row>
        </FormItem>
        <FormItem>
            <Row>
                <Col span="6">
                    <Button type="dashed" long @click="handleAddTarget" icon="md-add">Add target</Button>
                </Col>
            </Row>
        </FormItem>
    </Card>

</template>

<script>
  export default {
    name: "proxy",
    data(){
      return {
        proxy: {
          "preserve_host": false,
          "listen_path": "/example/*",
          "upstreams": {
            "balancing": "roundrobin",
            "targets": [
              {
                "target": "http://service1:8080/",
                "weight": 0
              }
            ]
          },
          "insecure_skip_verify": false,
          "strip_path": false,
          "append_path": false,
          "methods": [
            "GET"
          ],
          "hosts": [],
          "forwarding_timeouts": {
            "dial_timeout": "0s",
            "response_header_timeout": "0s"
          }
        },
        BalanceList: [
          'roundrobin',
          'weight',
        ],
        MethodsList: [
          'ALL',
          'GET',
          'POST',
          'PUT',
          'PATCH',
          'HEAD',
          'OPTIONS',
        ],
      }
    },
    props: {
      value: Object,
      title: String,
    },
    watch: {

      value: {
        deep: true,
        handler (nextForm) {
          this.proxy = nextForm
        },
      },
      proxy: {
        deep: true,
        handler (nextForm) {
          nextForm.upstreams.targets.forEach((val, ii) => {
            nextForm.upstreams.targets[ii].weight = parseInt(val.weight)
          })
          this.$emit('update:value', nextForm)

        },
      },
    },
    methods:{
      handleAddTarget() {
        this.proxy.upstreams.targets.push({
          target: '',
          weight: 0,
        });
      },

      handleRemoveTarget(ii) {
        if (this.proxy.upstreams.targets.length > 1) {
          this.proxy.upstreams.targets.splice(ii, 1)
        }else{
          this.$Message.warning('Can not delete last target!')
        }
      },
    },
    mounted () {
      // this.$set(this.proxy, this.value)
      // this.$nextTick(()=> {
      //   this.proxy = this.value
      // })
    },
  }
</script>

<style scoped>

</style>
