(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d20f3c6"],{b387:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("h2",{staticStyle:{"font-size":"20px",margin:"10px"}},[t._v(t._s(t.okMsg))]),t._l(t.failures,function(e,a){return n("Card",{key:a,staticStyle:{margin:"10px"}},[n("p",{attrs:{slot:"title"},slot:"title"},[n("span",{staticStyle:{color:"red"}},[t._v("API Name: ")]),t._v(" "+t._s(a))]),n("p",[n("span",{staticStyle:{color:"red"}},[t._v("Fail Reason")]),t._v(": "+t._s(e))])])})],2)},r=[],s=(n("96cf"),n("3b8d")),i=n("97af"),c=function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return i["a"].get("status",{params:t})},o={name:"HealthCheckList",data:function(){return{okMsg:"Everything is ok!",failures:{}}},computed:{},methods:{fetchHealthCheck:function(){var t=Object(s["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,c();case 2:e=t.sent,void 0!==e.failures&&(this.okMsg="Some apis not reachable!",this.failures=e.failures);case 4:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}()},mounted:function(){this.fetchHealthCheck()}},u=o,l=n("2877"),f=Object(l["a"])(u,a,r,!1,null,"6de60fe0",null);e["default"]=f.exports}}]);
//# sourceMappingURL=chunk-2d20f3c6.477bdf3c.js.map