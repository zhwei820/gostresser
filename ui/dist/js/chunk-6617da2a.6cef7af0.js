(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6617da2a"],{a064:function(t,e,n){"use strict";n.d(e,"c",function(){return a}),n.d(e,"d",function(){return u}),n.d(e,"b",function(){return i}),n.d(e,"e",function(){return c}),n.d(e,"a",function(){return o});var r=n("97af"),a=function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return r["a"].get("/oauth/servers",{params:t})},u=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return r["a"].get("/oauth/servers/".concat(t,"/"),{params:e})},i=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return r["a"].delete("/oauth/servers/".concat(t,"/"),{params:e})},c=function(t,e){return r["a"].put("/oauth/servers/".concat(t,"/"),e)},o=function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return r["a"].post("/oauth/servers/",t)}},c6ea:function(t,e,n){"use strict";n.r(e);var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("Button",{attrs:{id:"btn",type:"primary"},on:{click:t.createAuth}},[t._v("NEW")]),n("Table",{ref:"orderTable",staticClass:"ordertable",attrs:{columns:t.columns,data:t.tableData}})],1)},a=[],u=(n("96cf"),n("3b8d")),i=(n("7f7f"),n("cebc")),c=n("a064");function o(t){return[{title:"Auth Name",key:"name"},{title:"Operations",key:"action",width:250,align:"center",render:function(e,n){return e("div",[e("i-button",{attrs:{type:"primary"},on:{click:t.edit.bind(t,n.row)}},["Edit"]),e("i-button",{attrs:{type:"error"},on:{click:t.remove.bind(t,n.row)}},["Delete"])])}}]}var s=n("3eea"),h=n.n(s),f={page:1,page_size:10,total:0},l={name:"AuthList",data:function(){return{formData:{},tableData:[],filterform:{},pager:Object(i["a"])({},f)}},computed:{columns:function(){return o(this)}},methods:{edit:function(t){this.$router.push({name:"authDetail",params:{name:t.name}})},copy:function(t){this.$router.push({name:"authDetail",params:{name:t.name},query:{copy:"true"}})},createAuth:function(){this.$router.push({name:"authDetail",params:{name:"null"}})},remove:function(){var t=Object(u["a"])(regeneratorRuntime.mark(function t(e){var n=this;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:this.$Modal.confirm({title:"confirm",content:"confirm delete"+e.name+"?",width:350,onOk:function(){var t=Object(u["a"])(regeneratorRuntime.mark(function t(){return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(c["b"])(e.name,{});case 2:n.fetchAuth();case 3:case"end":return t.stop()}},t,this)}));function r(){return t.apply(this,arguments)}return r}()});case 1:case"end":return t.stop()}},t,this)}));function e(e){return t.apply(this,arguments)}return e}(),fetchAuth:function(){var t=Object(u["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(c["c"])(h()(this.pager,"total"));case 2:e=t.sent,this.tableData=e,this.pager.total=e.count;case 5:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}()},mounted:function(){this.fetchAuth()}},p=l,m=n("2877"),d=Object(m["a"])(p,r,a,!1,null,"34119916",null);e["default"]=d.exports}}]);
//# sourceMappingURL=chunk-6617da2a.6cef7af0.js.map